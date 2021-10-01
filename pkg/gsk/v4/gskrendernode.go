// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/graphene"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gsk/gsk.h>
// void _gotk4_gsk4_ParseErrorFunc(GskParseLocation*, GskParseLocation*, GError*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_render_node_get_type()), F: marshalRenderNoder},
	})
}

// ParseErrorFunc: type of callback that is called when an error occurs during
// node deserialization.
type ParseErrorFunc func(start *ParseLocation, end *ParseLocation, err error)

//export _gotk4_gsk4_ParseErrorFunc
func _gotk4_gsk4_ParseErrorFunc(arg0 *C.GskParseLocation, arg1 *C.GskParseLocation, arg2 *C.GError, arg3 C.gpointer) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var start *ParseLocation // out
	var end *ParseLocation   // out
	var err error            // out

	start = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg0)))
	end = (*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	err = gerror.Take(unsafe.Pointer(arg2))

	fn := v.(ParseErrorFunc)
	fn(start, end, err)
}

// RenderNode: GskRenderNode is the basic block in a scene graph to be rendered
// using GskRenderer.
//
// Each node has a parent, except the top-level node; each node may have
// children nodes.
//
// Each node has an associated drawing surface, which has the size of the
// rectangle set when creating it.
//
// Render nodes are meant to be transient; once they have been associated to a
// gsk.Renderer it's safe to release any reference you have on them. All
// gsk.RenderNodes are immutable, you can only specify their properties during
// construction.
type RenderNode struct {
	*externglib.Object
}

// RenderNoder describes RenderNode's abstract methods.
type RenderNoder interface {
	externglib.Objector

	// Draw the contents of node to the given cairo context.
	Draw(cr *cairo.Context)
	// Bounds retrieves the boundaries of the node.
	Bounds() graphene.Rect
	// NodeType returns the type of the node.
	NodeType() RenderNodeType
	// Serialize serializes the node for later deserialization via
	// gsk_render_node_deserialize().
	Serialize() *glib.Bytes
	// WriteToFile: this function is equivalent to calling
	// gsk_render_node_serialize() followed by g_file_set_contents().
	WriteToFile(filename string) error
}

var _ RenderNoder = (*RenderNode)(nil)

func wrapRenderNode(obj *externglib.Object) *RenderNode {
	return &RenderNode{
		Object: obj,
	}
}

func marshalRenderNoder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRenderNode(obj), nil
}

// Draw the contents of node to the given cairo context.
//
// Typically, you'll use this function to implement fallback rendering of
// GskRenderNodes on an intermediate Cairo context, instead of using the drawing
// context associated to a GdkSurface's rendering buffer.
//
// For advanced nodes that cannot be supported using Cairo, in particular for
// nodes doing 3D operations, this function may fail.
func (node *RenderNode) Draw(cr *cairo.Context) {
	var _arg0 *C.GskRenderNode // out
	var _arg1 *C.cairo_t       // out

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	C.gsk_render_node_draw(_arg0, _arg1)
	runtime.KeepAlive(node)
	runtime.KeepAlive(cr)
}

// Bounds retrieves the boundaries of the node.
//
// The node will not draw outside of its boundaries.
func (node *RenderNode) Bounds() graphene.Rect {
	var _arg0 *C.GskRenderNode  // out
	var _arg1 C.graphene_rect_t // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))

	C.gsk_render_node_get_bounds(_arg0, &_arg1)
	runtime.KeepAlive(node)

	var _bounds graphene.Rect // out

	_bounds = *(*graphene.Rect)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _bounds
}

// NodeType returns the type of the node.
func (node *RenderNode) NodeType() RenderNodeType {
	var _arg0 *C.GskRenderNode    // out
	var _cret C.GskRenderNodeType // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))

	_cret = C.gsk_render_node_get_node_type(_arg0)
	runtime.KeepAlive(node)

	var _renderNodeType RenderNodeType // out

	_renderNodeType = RenderNodeType(_cret)

	return _renderNodeType
}

// Serialize serializes the node for later deserialization via
// gsk_render_node_deserialize(). No guarantees are made about the format used
// other than that the same version of GTK will be able to deserialize the
// result of a call to gsk_render_node_serialize() and
// gsk_render_node_deserialize() will correctly reject files it cannot open that
// were created with previous versions of GTK.
//
// The intended use of this functions is testing, benchmarking and debugging.
// The format is not meant as a permanent storage format.
func (node *RenderNode) Serialize() *glib.Bytes {
	var _arg0 *C.GskRenderNode // out
	var _cret *C.GBytes        // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))

	_cret = C.gsk_render_node_serialize(_arg0)
	runtime.KeepAlive(node)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// WriteToFile: this function is equivalent to calling
// gsk_render_node_serialize() followed by g_file_set_contents().
//
// See those two functions for details on the arguments.
//
// It is mostly intended for use inside a debugger to quickly dump a render node
// to a file for later inspection.
func (node *RenderNode) WriteToFile(filename string) error {
	var _arg0 *C.GskRenderNode // out
	var _arg1 *C.char          // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(node.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gsk_render_node_write_to_file(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(node)
	runtime.KeepAlive(filename)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// RenderNodeDeserialize loads data previously created via
// gsk_render_node_serialize().
//
// For a discussion of the supported format, see that function.
func RenderNodeDeserialize(bytes *glib.Bytes, errorFunc ParseErrorFunc) RenderNoder {
	var _arg1 *C.GBytes           // out
	var _arg2 C.GskParseErrorFunc // out
	var _arg3 C.gpointer
	var _cret *C.GskRenderNode // in

	_arg1 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))
	if errorFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gsk4_ParseErrorFunc)
		_arg3 = C.gpointer(gbox.Assign(errorFunc))
		defer gbox.Delete(uintptr(_arg3))
	}

	_cret = C.gsk_render_node_deserialize(_arg1, _arg2, _arg3)
	runtime.KeepAlive(bytes)
	runtime.KeepAlive(errorFunc)

	var _renderNode RenderNoder // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			rv, ok := (externglib.CastObject(object)).(RenderNoder)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gsk.RenderNoder")
			}
			_renderNode = rv
		}
	}

	return _renderNode
}

// ColorStop: color stop in a gradient node.
//
// An instance of this type is always passed by reference.
type ColorStop struct {
	*colorStop
}

// colorStop is the struct that's finalized.
type colorStop struct {
	native *C.GskColorStop
}

// Offset: offset of the color stop.
func (c *ColorStop) Offset() float32 {
	var v float32 // out
	v = float32(c.native.offset)
	return v
}

// Color: color at the given offset.
func (c *ColorStop) Color() gdk.RGBA {
	var v gdk.RGBA // out
	v = *(*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&c.native.color))))
	return v
}

// ParseLocation: location in a parse buffer.
//
// An instance of this type is always passed by reference.
type ParseLocation struct {
	*parseLocation
}

// parseLocation is the struct that's finalized.
type parseLocation struct {
	native *C.GskParseLocation
}

// NewParseLocation creates a new ParseLocation instance from the given
// fields.
func NewParseLocation(bytes, chars, lines, lineBytes, lineChars uint) ParseLocation {
	var f0 C.gsize // out
	f0 = C.gsize(bytes)
	var f1 C.gsize // out
	f1 = C.gsize(chars)
	var f2 C.gsize // out
	f2 = C.gsize(lines)
	var f3 C.gsize // out
	f3 = C.gsize(lineBytes)
	var f4 C.gsize // out
	f4 = C.gsize(lineChars)

	v := C.GskParseLocation{
		bytes:      f0,
		chars:      f1,
		lines:      f2,
		line_bytes: f3,
		line_chars: f4,
	}

	return *(*ParseLocation)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Bytes: offset of the location in the parse buffer, as bytes.
func (p *ParseLocation) Bytes() uint {
	var v uint // out
	v = uint(p.native.bytes)
	return v
}

// Chars: offset of the location in the parse buffer, as characters.
func (p *ParseLocation) Chars() uint {
	var v uint // out
	v = uint(p.native.chars)
	return v
}

// Lines: line of the location in the parse buffer.
func (p *ParseLocation) Lines() uint {
	var v uint // out
	v = uint(p.native.lines)
	return v
}

// LineBytes: position in the line, as bytes.
func (p *ParseLocation) LineBytes() uint {
	var v uint // out
	v = uint(p.native.line_bytes)
	return v
}

// LineChars: position in the line, as characters.
func (p *ParseLocation) LineChars() uint {
	var v uint // out
	v = uint(p.native.line_chars)
	return v
}

// Shadow: shadow parameters in a shadow node.
//
// An instance of this type is always passed by reference.
type Shadow struct {
	*shadow
}

// shadow is the struct that's finalized.
type shadow struct {
	native *C.GskShadow
}

// Color: color of the shadow.
func (s *Shadow) Color() gdk.RGBA {
	var v gdk.RGBA // out
	v = *(*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&s.native.color))))
	return v
}

// Dx: horizontal offset of the shadow.
func (s *Shadow) Dx() float32 {
	var v float32 // out
	v = float32(s.native.dx)
	return v
}

// Dy: vertical offset of the shadow.
func (s *Shadow) Dy() float32 {
	var v float32 // out
	v = float32(s.native.dy)
	return v
}

// Radius radius of the shadow.
func (s *Shadow) Radius() float32 {
	var v float32 // out
	v = float32(s.native.radius)
	return v
}
