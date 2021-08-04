// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PangoContextGet creates a Context for the default GDK screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the default screen; if these options change it
// will not be updated. Using gtk_widget_get_pango_context() is more convenient
// if you want to keep a context around and track changes to the screen’s font
// rendering settings.
func PangoContextGet() *pango.Context {
	var _cret *C.PangoContext // in

	_cret = C.gdk_pango_context_get()

	var _context *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}

// PangoContextGetForDisplay creates a Context for display.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the display; if these options change it will not
// be updated. Using gtk_widget_get_pango_context() is more convenient if you
// want to keep a context around and track changes to the font rendering
// settings.
func PangoContextGetForDisplay(display *Display) *pango.Context {
	var _arg1 *C.GdkDisplay   // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_pango_context_get_for_display(_arg1)

	var _context *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}

// PangoContextGetForScreen creates a Context for screen.
//
// The context must be freed when you’re finished with it.
//
// When using GTK+, normally you should use gtk_widget_get_pango_context()
// instead of this function, to get the appropriate context for the widget you
// intend to render text onto.
//
// The newly created context will have the default font options (see
// #cairo_font_options_t) for the screen; if these options change it will not be
// updated. Using gtk_widget_get_pango_context() is more convenient if you want
// to keep a context around and track changes to the screen’s font rendering
// settings.
func PangoContextGetForScreen(screen *Screen) *pango.Context {
	var _arg1 *C.GdkScreen    // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gdk_pango_context_get_for_screen(_arg1)

	var _context *pango.Context // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}
