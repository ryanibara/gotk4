// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*CSSSection)(unsafe.Pointer(b)), nil
}

// NewCSSSection constructs a struct CSSSection.
func NewCSSSection(file gio.Filer, start *CSSLocation, end *CSSLocation) *CSSSection {
	var _arg1 *C.GFile          // out
	var _arg2 *C.GtkCssLocation // out
	var _arg3 *C.GtkCssLocation // out
	var _cret *C.GtkCssSection  // in

	_arg1 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))
	_arg2 = (*C.GtkCssLocation)(unsafe.Pointer(start))
	_arg3 = (*C.GtkCssLocation)(unsafe.Pointer(end))

	_cret = C.gtk_css_section_new(_arg1, _arg2, _arg3)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// EndLocation returns the location in the CSS document where this section ends.
func (section *CSSSection) EndLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection  // out
	var _cret *C.GtkCssLocation // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_get_end_location(_arg0)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(unsafe.Pointer(_cret))

	return _cssLocation
}

// File gets the file that @section was parsed from.
//
// If no such file exists, for example because the CSS was loaded via
// [method@Gtk.CssProvider.load_from_data], then `NULL` is returned.
func (section *CSSSection) File() *gio.File {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GFile         // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_get_file(_arg0)

	var _file *gio.File // out

	_file = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.File)

	return _file
}

// Parent gets the parent section for the given `section`.
//
// The parent section is the section that contains this `section`. A special
// case are sections of type `GTK_CSS_SECTION_DOCUMEN`T. Their parent will
// either be `NULL` if they are the original CSS document that was loaded by
// [method@Gtk.CssProvider.load_from_file] or a section of type
// `GTK_CSS_SECTION_IMPORT` if it was loaded with an `@import` rule from a
// different file.
func (section *CSSSection) Parent() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_get_parent(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// StartLocation returns the location in the CSS document where this section
// starts.
func (section *CSSSection) StartLocation() *CSSLocation {
	var _arg0 *C.GtkCssSection  // out
	var _cret *C.GtkCssLocation // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_get_start_location(_arg0)

	var _cssLocation *CSSLocation // out

	_cssLocation = (*CSSLocation)(unsafe.Pointer(_cret))

	return _cssLocation
}

// Ref increments the reference count on `section`.
func (section *CSSSection) ref() *CSSSection {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.GtkCssSection // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_ref(_arg0)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// String prints the section into a human-readable text form using
// [method@Gtk.CssSection.print].
func (section *CSSSection) String() string {
	var _arg0 *C.GtkCssSection // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	_cret = C.gtk_css_section_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref decrements the reference count on `section`, freeing the structure if
// the reference count reaches 0.
func (section *CSSSection) unref() {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(section))

	C.gtk_css_section_unref(_arg0)
}
