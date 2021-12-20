// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_context_get_type()), F: marshalPrintContexter},
	})
}

// PrintContext: GtkPrintContext encapsulates context information that is
// required when drawing pages for printing.
//
// This includes the cairo context and important parameters like page size and
// resolution. It also lets you easily create pango.Layout and pango.Context
// objects that match the font metrics of the cairo surface.
//
// GtkPrintContext objects get passed to the gtk.PrintOperation::begin-print,
// gtk.PrintOperation::end-print, gtk.PrintOperation::request-page-setup and
// gtk.PrintOperation::draw-page signals on the gtk.PrintOperation object.
//
// Using GtkPrintContext in a ::draw-page callback
//
//    static void
//    draw_page (GtkPrintOperation *operation,
//               GtkPrintContext   *context,
//               int                page_nr)
//    {
//      cairo_t *cr;
//      PangoLayout *layout;
//      PangoFontDescription *desc;
//
//      cr = gtk_print_context_get_cairo_context (context);
//
//      // Draw a red rectangle, as wide as the paper (inside the margins)
//      cairo_set_source_rgb (cr, 1.0, 0, 0);
//      cairo_rectangle (cr, 0, 0, gtk_print_context_get_width (context), 50);
//
//      cairo_fill (cr);
//
//      // Draw some lines
//      cairo_move_to (cr, 20, 10);
//      cairo_line_to (cr, 40, 20);
//      cairo_arc (cr, 60, 60, 20, 0, M_PI);
//      cairo_line_to (cr, 80, 20);
//
//      cairo_set_source_rgb (cr, 0, 0, 0);
//      cairo_set_line_width (cr, 5);
//      cairo_set_line_cap (cr, CAIRO_LINE_CAP_ROUND);
//      cairo_set_line_join (cr, CAIRO_LINE_JOIN_ROUND);
//
//      cairo_stroke (cr);
//
//      // Draw some text
//      layout = gtk_print_context_create_pango_layout (context);
//      pango_layout_set_text (layout, "Hello World! Printing is easy", -1);
//      desc = pango_font_description_from_string ("sans 28");
//      pango_layout_set_font_description (layout, desc);
//      pango_font_description_free (desc);
//
//      cairo_move_to (cr, 30, 20);
//      pango_cairo_layout_path (cr, layout);
//
//      // Font Outline
//      cairo_set_source_rgb (cr, 0.93, 1.0, 0.47);
//      cairo_set_line_width (cr, 0.5);
//      cairo_stroke_preserve (cr);
//
//      // Font Fill
//      cairo_set_source_rgb (cr, 0, 0.0, 1.0);
//      cairo_fill (cr);
//
//      g_object_unref (layout);
//    }.
type PrintContext struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*PrintContext)(nil)
)

func wrapPrintContext(obj *externglib.Object) *PrintContext {
	return &PrintContext{
		Object: obj,
	}
}

func marshalPrintContexter(p uintptr) (interface{}, error) {
	return wrapPrintContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CreatePangoContext creates a new PangoContext that can be used with the
// GtkPrintContext.
//
// The function returns the following values:
//
//    - ret: new Pango context for context.
//
func (context *PrintContext) CreatePangoContext() *pango.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoContext    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_create_pango_context(_arg0)
	runtime.KeepAlive(context)

	var _ret *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_ret = &pango.Context{
			Object: obj,
		}
	}

	return _ret
}

// CreatePangoLayout creates a new PangoLayout that is suitable for use with the
// GtkPrintContext.
//
// The function returns the following values:
//
//    - layout: new Pango layout for context.
//
func (context *PrintContext) CreatePangoLayout() *pango.Layout {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoLayout     // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_create_pango_layout(_arg0)
	runtime.KeepAlive(context)

	var _layout *pango.Layout // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// CairoContext obtains the cairo context that is associated with the
// GtkPrintContext.
//
// The function returns the following values:
//
//    - ret: cairo context of context.
//
func (context *PrintContext) CairoContext() *cairo.Context {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.cairo_t         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_cairo_context(_arg0)
	runtime.KeepAlive(context)

	var _ret *cairo.Context // out

	_ret = cairo.WrapContext(uintptr(unsafe.Pointer(_cret)))
	C.cairo_reference(_cret)
	runtime.SetFinalizer(_ret, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})

	return _ret
}

// DPIX obtains the horizontal resolution of the GtkPrintContext, in dots per
// inch.
//
// The function returns the following values:
//
//    - gdouble: horizontal resolution of context.
//
func (context *PrintContext) DPIX() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_dpi_x(_arg0)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DPIY obtains the vertical resolution of the GtkPrintContext, in dots per
// inch.
//
// The function returns the following values:
//
//    - gdouble: vertical resolution of context.
//
func (context *PrintContext) DPIY() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_dpi_y(_arg0)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// HardMargins obtains the hardware printer margins of the GtkPrintContext, in
// units.
//
// The function returns the following values:
//
//    - top hardware printer margin.
//    - bottom hardware printer margin.
//    - left hardware printer margin.
//    - right hardware printer margin.
//    - ok: TRUE if the hard margins were retrieved.
//
func (context *PrintContext) HardMargins() (top float64, bottom float64, left float64, right float64, ok bool) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 C.double           // in
	var _arg2 C.double           // in
	var _arg3 C.double           // in
	var _arg4 C.double           // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_hard_margins(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(context)

	var _top float64    // out
	var _bottom float64 // out
	var _left float64   // out
	var _right float64  // out
	var _ok bool        // out

	_top = float64(_arg1)
	_bottom = float64(_arg2)
	_left = float64(_arg3)
	_right = float64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _top, _bottom, _left, _right, _ok
}

// Height obtains the height of the GtkPrintContext, in pixels.
//
// The function returns the following values:
//
//    - gdouble: height of context.
//
func (context *PrintContext) Height() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_height(_arg0)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PageSetup obtains the GtkPageSetup that determines the page dimensions of the
// GtkPrintContext.
//
// The function returns the following values:
//
//    - pageSetup: page setup of context.
//
func (context *PrintContext) PageSetup() *PageSetup {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.GtkPageSetup    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_page_setup(_arg0)
	runtime.KeepAlive(context)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(externglib.Take(unsafe.Pointer(_cret)))

	return _pageSetup
}

// PangoFontmap returns a PangoFontMap that is suitable for use with the
// GtkPrintContext.
//
// The function returns the following values:
//
//    - fontMap: font map of context.
//
func (context *PrintContext) PangoFontmap() pango.FontMapper {
	var _arg0 *C.GtkPrintContext // out
	var _cret *C.PangoFontMap    // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_pango_fontmap(_arg0)
	runtime.KeepAlive(context)

	var _fontMap pango.FontMapper // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type pango.FontMapper is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(pango.FontMapper)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not pango.FontMapper")
		}
		_fontMap = rv
	}

	return _fontMap
}

// Width obtains the width of the GtkPrintContext, in pixels.
//
// The function returns the following values:
//
//    - gdouble: width of context.
//
func (context *PrintContext) Width() float64 {
	var _arg0 *C.GtkPrintContext // out
	var _cret C.double           // in

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))

	_cret = C.gtk_print_context_get_width(_arg0)
	runtime.KeepAlive(context)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetCairoContext sets a new cairo context on a print context.
//
// This function is intended to be used when implementing an internal print
// preview, it is not needed for printing, since GTK itself creates a suitable
// cairo context in that case.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - dpiX: horizontal resolution to use with cr.
//    - dpiY: vertical resolution to use with cr.
//
func (context *PrintContext) SetCairoContext(cr *cairo.Context, dpiX, dpiY float64) {
	var _arg0 *C.GtkPrintContext // out
	var _arg1 *C.cairo_t         // out
	var _arg2 C.double           // out
	var _arg3 C.double           // out

	_arg0 = (*C.GtkPrintContext)(unsafe.Pointer(context.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = C.double(dpiX)
	_arg3 = C.double(dpiY)

	C.gtk_print_context_set_cairo_context(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(dpiX)
	runtime.KeepAlive(dpiY)
}
