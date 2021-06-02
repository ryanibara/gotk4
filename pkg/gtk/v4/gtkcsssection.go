// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSection defines a part of a CSS document.
//
// Because sections are nested into one another, you can use
// gtk_css_section_get_parent() to get the containing region.
type CSSSection struct {
	native C.GtkCssSection
}

// WrapCSSSection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCSSSection(ptr unsafe.Pointer) *CSSSection {
	if ptr == nil {
		return nil
	}

	return (*CSSSection)(ptr)
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCSSSection(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// NewCSSSection constructs a struct CSSSection.
func NewCSSSection(file gio.File, start *CSSLocation, end *CSSLocation) *CSSSection {
	var arg1 *C.GFile
	var arg2 *C.GtkCssLocation
	var arg3 *C.GtkCssLocation

	arg1 = (*C.GFile)(file.Native())
	arg2 = (*C.GtkCssLocation)(start.Native())
	arg3 = (*C.GtkCssLocation)(end.Native())

	ret := C.gtk_css_section_new(arg1, arg2, arg3)

	var ret0 *CSSSection

	{
		ret0 = WrapCSSSection(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *CSSSection) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// EndLocation returns the location in the CSS document where this section ends.
func (s *CSSSection) EndLocation() *CSSLocation {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_get_end_location(arg0)

	var ret0 *CSSLocation

	{
		ret0 = WrapCSSLocation(unsafe.Pointer(ret))
	}

	return ret0
}

// File gets the file that @section was parsed from.
//
// If no such file exists, for example because the CSS was loaded via
// [method@Gtk.CssProvider.load_from_data], then `NULL` is returned.
func (s *CSSSection) File() gio.File {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_get_file(arg0)

	var ret0 gio.File

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gio.File)

	return ret0
}

// Parent gets the parent section for the given `section`.
//
// The parent section is the section that contains this `section`. A special
// case are sections of type `GTK_CSS_SECTION_DOCUMEN`T. Their parent will
// either be `NULL` if they are the original CSS document that was loaded by
// [method@Gtk.CssProvider.load_from_file] or a section of type
// `GTK_CSS_SECTION_IMPORT` if it was loaded with an `@import` rule from a
// different file.
func (s *CSSSection) Parent() *CSSSection {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_get_parent(arg0)

	var ret0 *CSSSection

	{
		ret0 = WrapCSSSection(unsafe.Pointer(ret))
	}

	return ret0
}

// StartLocation returns the location in the CSS document where this section
// starts.
func (s *CSSSection) StartLocation() *CSSLocation {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_get_start_location(arg0)

	var ret0 *CSSLocation

	{
		ret0 = WrapCSSLocation(unsafe.Pointer(ret))
	}

	return ret0
}

// Print prints the `section` into `string` in a human-readable form.
//
// This is a form like `gtk.css:32:1-23` to denote line 32, characters 1 to 23
// in the file `gtk.css`.
func (s *CSSSection) Print(string *glib.String) {
	var arg0 *C.GtkCssSection
	var arg1 *C.GString

	arg0 = (*C.GtkCssSection)(s.Native())
	arg1 = (*C.GString)(string.Native())

	C.gtk_css_section_print(arg0, arg1)
}

// Ref increments the reference count on `section`.
func (s *CSSSection) Ref() *CSSSection {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_ref(arg0)

	var ret0 *CSSSection

	{
		ret0 = WrapCSSSection(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *CSSSection) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// String prints the section into a human-readable text form using
// [method@Gtk.CssSection.print].
func (s *CSSSection) String() string {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	ret := C.gtk_css_section_to_string(arg0)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// Unref decrements the reference count on `section`, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Unref() {
	var arg0 *C.GtkCssSection

	arg0 = (*C.GtkCssSection)(s.Native())

	C.gtk_css_section_unref(arg0)
}
