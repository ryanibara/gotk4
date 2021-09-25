package generators

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/pen"
	"github.com/diamondburned/gotk4/gir/girgen/strcases"
	"github.com/diamondburned/gotk4/gir/girgen/types"
	"github.com/diamondburned/gotk4/gir/girgen/types/typeconv"
)

var callbackTmpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	type {{ .GoName }} func{{ .GoTail }}

	//export {{ .CGoName }}
	func {{ .CGoName }}{{ .CGoTail }} {{ .Block }}
`)

// GenerateCallback generates a callback type declaration and handler into the
// given file generator.
func GenerateCallback(gen FileGeneratorWriter, callback *gir.Callback) bool {
	generator := NewCallbackGenerator(gen)
	if !generator.Use(callback) {
		return false
	}

	writer := FileWriterFromType(gen, callback)
	writer.Pen().WriteTmpl(callbackTmpl, &generator)
	file.ApplyHeader(writer, &generator)

	return true
}

// CallbackGenerator generates a callback in Go.
type CallbackGenerator struct {
	*gir.Callback
	GoName  string
	GoTail  string
	CGoName string
	CGoTail string
	Block   string

	Closure *int
	Destroy *int

	pen    *pen.BlockSections
	gen    FileGenerator
	header file.Header
}

// NewCallbackGenerator creates a new CallbackGenerator instance.
func NewCallbackGenerator(gen FileGenerator) CallbackGenerator {
	return CallbackGenerator{
		pen: pen.NewBlockSections(256, 1024, 4096, 128, 1024, 4096, 128),
		gen: gen,
	}
}

func (g *CallbackGenerator) Logln(lvl logger.Level, v ...interface{}) {
	g.gen.Logln(lvl, logger.Prefix(v, fmt.Sprintf("callback %s (C.%s):", g.Name, g.CIdentifier))...)
}

// Reset resets the callback generator.
func (g *CallbackGenerator) Reset() {
	g.pen.Reset()

	*g = CallbackGenerator{
		pen: g.pen,
		gen: g.gen,
	}
}

// Header returns the callback generator's current header.
func (g *CallbackGenerator) Header() *file.Header {
	return &g.header
}

// Use sets the callback generator to the given GIR callback.
func (g *CallbackGenerator) Use(cb *gir.Callback) bool {
	g.Reset()
	g.Callback = cb

	// We can't use the callback if it has no closure parameters.
	if cb.Parameters == nil || len(cb.Parameters.Parameters) == 0 {
		g.Logln(logger.Debug, "has no parameters at all")
		return false
	}

	if !cb.IsIntrospectable() || types.Filter(g.gen, cb.Name, cb.CIdentifier) {
		return false
	}

	// Don't generate destroy notifiers. It's an edge case that we handle
	// separately and mostly manually. There are also no good ways to detect
	// this.
	if strings.HasSuffix(cb.Name, "DestroyNotify") {
		g.Logln(logger.Debug, "skipping DestroyNotify-ish callback")
		return false
	}

	g.GoName = strcases.PascalToGo(cb.Name)
	g.CGoName = file.CallbackExportedName(g.gen.Namespace(), cb)

	if g.Closure = findClosure(cb.Parameters.Parameters); g.Closure == nil {
		g.Logln(logger.Debug, "skipping since no closure argument")
		return false
	}

	if g.CGoTail = g.cgoTail(); g.CGoTail == "" {
		return false
	}

	if !g.renderBlock() {
		return false
	}

	return true
}

// findClosure returns the closure number or nil.
func findClosure(params []gir.Parameter) *int {
	for _, param := range params {
		if param.Closure != nil {
			return param.Closure
		}
	}
	return nil
}

func callbackArg(i int) string {
	return "arg" + strconv.Itoa(i)
}

func (g *CallbackGenerator) cgoTail() string {
	cgotail := pen.NewJoints(", ", len(g.Parameters.Parameters))

	for i, param := range g.Parameters.Parameters {
		ctype := types.AnyTypeC(param.AnyType)
		if ctype == "" {
			g.Logln(logger.Debug, "anyTypeC parameter is empty")
			return "" // probably var_args
		}

		cgotype := types.AnyTypeCGo(param.AnyType)
		cgotail.Addf("%s %s", callbackArg(i), cgotype)
	}

	callTail := "(" + cgotail.Join() + ")"

	if !types.ReturnIsVoid(g.ReturnValue) {
		ctype := types.AnyTypeC(g.ReturnValue.AnyType)
		if ctype == "" {
			g.Logln(logger.Debug, "anyTypeC return is empty")
			return ""
		}

		callTail += fmt.Sprintf(" (cret %s)", types.AnyTypeCGo(g.ReturnValue.AnyType))
	}

	return callTail
}

func (g *CallbackGenerator) renderBlock() bool {
	defer g.pen.Reset()

	const (
		secPrefix = iota
		secInputPre
		secInputConv
		secFnCall
		secOutputPre
		secOutputConv
		secReturn
	)

	g.pen.Linef(secPrefix, "v := gbox.Get(uintptr(%s))", callbackArg(*g.Closure))
	g.pen.Linef(secPrefix, "if v == nil {")
	g.pen.Linef(secPrefix, "  panic(`callback not found`)")
	g.pen.Linef(secPrefix, "}")
	g.pen.EmptyLine(secPrefix)

	g.pen.Linef(secFnCall, "fn := v.(%s)", g.GoName)

	callbackValues := make([]typeconv.ConversionValue, 0, len(g.Parameters.Parameters)+2)

	for i, param := range g.Parameters.Parameters {
		// Skip generating the closure parameter.
		if param.Skip || i == *g.Closure {
			continue
		}

		// Reguess the parameter.
		param.Direction = types.GuessParameterOutput(&param)

		var in string
		var out string
		var dir typeconv.ConversionDirection

		switch param.Direction {
		case "in":
			in = callbackArg(i)
			out = strcases.SnakeToGo(false, param.Name)
			dir = typeconv.ConvertCToGo
		case "out":
			in = strcases.SnakeToGo(false, param.Name)
			out = callbackArg(i)
			dir = typeconv.ConvertGoToC
		default:
			// TODO: inout
			return false
		}

		value := typeconv.NewValue(in, out, i, dir, param)
		callbackValues = append(callbackValues, value)
	}

	var hasReturn bool
	if !types.ReturnIsVoid(g.ReturnValue) {
		hasReturn = true
		returnName := callable.ReturnName(&g.CallableAttrs)

		value := typeconv.NewReturnValue(returnName, "cret", typeconv.ConvertGoToC, *g.ReturnValue)
		callbackValues = append(callbackValues, value)
	}

	typ := &gir.TypeFindResult{
		NamespaceFindResult: g.gen.Namespace(),
		Type:                g.Callback,
	}

	convert := typeconv.NewConverter(g.gen, typ, callbackValues)
	convert.Callback = true
	convert.UseLogger(g)

	results := convert.ConvertAll()
	if results == nil {
		g.Logln(logger.Debug, "has no conversion")
		return false
	}

	file.ApplyHeader(g, convert)

	goCallArgs := pen.NewJoints(", ", len(results))
	goCallRets := pen.NewJoints(", ", len(results))

	goTypeArgs := pen.NewJoints(", ", len(results))
	goTypeRets := pen.NewJoints(", ", len(results))

	for _, result := range results {
		if result.Skip {
			continue
		}

		switch result.Direction {
		case typeconv.ConvertCToGo:
			goCallArgs.Add(result.Out.Call)
			goTypeArgs.Addf("%s %s", result.OutName, result.Out.Type)

			g.pen.Line(secInputPre, result.Out.Declare)
			g.pen.Line(secInputConv, result.Conversion)

		case typeconv.ConvertGoToC:
			goCallRets.Add(result.In.Call)
			goTypeRets.Addf("%s %s", result.InName, result.In.Type)

			// Out.Declare is declared in the function signature.
			// g.pen.Line(secOutputPre, result.Out.Declare)

			g.pen.Line(secOutputConv, result.Conversion)
		}
	}

	if goCallRets.Len() == 0 {
		g.pen.Linef(secFnCall, "fn(%s)", goCallArgs.Join())
	} else {
		g.pen.Linef(secFnCall, "%s := fn(%s)", goCallRets.Join(), goCallArgs.Join())
	}

	if hasReturn {
		g.pen.Linef(secReturn, "return cret")
	}

	g.Block = g.pen.String()

	g.GoTail = "(" + goTypeArgs.Join() + ")"
	if goTypeRets.Len() > 0 {
		g.GoTail += " (" + goTypeRets.Join() + ")"
	}

	// Only add the import now, since we know that the callback will be
	// generated.
	g.header.ImportCore("gbox")

	return true
}
