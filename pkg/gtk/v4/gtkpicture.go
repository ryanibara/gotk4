// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_picture_get_type()), F: marshalPicture},
	})
}

// Picture: the `GtkPicture` widget displays a `GdkPaintable`.
//
// !An example GtkPicture (picture.png)
//
// Many convenience functions are provided to make pictures simple to use. For
// example, if you want to load an image from a file, and then display it,
// there’s a convenience function to do this:
//
// “`c GtkWidget *widget = gtk_picture_new_for_filename ("myfile.png"); “`
//
// If the file isn’t loaded successfully, the picture will contain a “broken
// image” icon similar to that used in many web browsers. If you want to handle
// errors in loading the file yourself, for example by displaying an error
// message, then load the image with [ctor@Gdk.Texture.new_from_file], then
// create the `GtkPicture` with [ctor@Gtk.Picture.new_for_paintable].
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of `GResource` for details. In
// this case, [ctor@Gtk.Picture.new_for_resource] and
// [method@Gtk.Picture.set_resource] should be used.
//
// `GtkPicture` displays an image at its natural size. See [class@Gtk.Image] if
// you want to display a fixed-size image, such as an icon.
//
//
// Sizing the paintable
//
// You can influence how the paintable is displayed inside the `GtkPicture`. By
// turning off [property@Gtk.Picture:keep-aspect-ratio] you can allow the
// paintable to get stretched. [property@Gtk.Picture:can-shrink] can be unset to
// make sure that paintables are never made smaller than their ideal size - but
// be careful if you do not know the size of the paintable in use (like when
// displaying user-loaded images). This can easily cause the picture to grow
// larger than the screen. And [property@GtkWidget:halign] and
// [property@GtkWidget:valign] can be used to make sure the paintable doesn't
// fill all available space but is instead displayed at its original size.
//
//
// CSS nodes
//
// `GtkPicture` has a single CSS node with the name `picture`.
//
//
// Accessibility
//
// `GtkPicture` uses the `GTK_ACCESSIBLE_ROLE_IMG` role.
type Picture interface {
	Widget

	AlternativeText() string

	CanShrink() bool

	KeepAspectRatio() bool

	SetAlternativeTextPicture(alternativeText string)

	SetCanShrinkPicture(canShrink bool)

	SetFilenamePicture(filename string)

	SetKeepAspectRatioPicture(keepAspectRatio bool)

	SetPixbufPicture(pixbuf gdkpixbuf.Pixbuf)

	SetResourcePicture(resourcePath string)
}

// picture implements the Picture class.
type picture struct {
	Widget
}

// WrapPicture wraps a GObject to the right type. It is
// primarily used internally.
func WrapPicture(obj *externglib.Object) Picture {
	return picture{
		Widget: WrapWidget(obj),
	}
}

func marshalPicture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPicture(obj), nil
}

func NewPicture() Picture {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_picture_new()

	var _picture Picture // out

	_picture = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Picture)

	return _picture
}

func NewPictureForFilename(filename string) Picture {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_picture_new_for_filename(_arg1)

	var _picture Picture // out

	_picture = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Picture)

	return _picture
}

func NewPictureForPixbuf(pixbuf gdkpixbuf.Pixbuf) Picture {
	var _arg1 *C.GdkPixbuf // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	_cret = C.gtk_picture_new_for_pixbuf(_arg1)

	var _picture Picture // out

	_picture = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Picture)

	return _picture
}

func NewPictureForResource(resourcePath string) Picture {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_picture_new_for_resource(_arg1)

	var _picture Picture // out

	_picture = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Picture)

	return _picture
}

func (s picture) AlternativeText() string {
	var _arg0 *C.GtkPicture // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_picture_get_alternative_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s picture) CanShrink() bool {
	var _arg0 *C.GtkPicture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_picture_get_can_shrink(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s picture) KeepAspectRatio() bool {
	var _arg0 *C.GtkPicture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_picture_get_keep_aspect_ratio(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s picture) SetAlternativeTextPicture(alternativeText string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(alternativeText))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_picture_set_alternative_text(_arg0, _arg1)
}

func (s picture) SetCanShrinkPicture(canShrink bool) {
	var _arg0 *C.GtkPicture // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if canShrink {
		_arg1 = C.TRUE
	}

	C.gtk_picture_set_can_shrink(_arg0, _arg1)
}

func (s picture) SetFilenamePicture(filename string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_picture_set_filename(_arg0, _arg1)
}

func (s picture) SetKeepAspectRatioPicture(keepAspectRatio bool) {
	var _arg0 *C.GtkPicture // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if keepAspectRatio {
		_arg1 = C.TRUE
	}

	C.gtk_picture_set_keep_aspect_ratio(_arg0, _arg1)
}

func (s picture) SetPixbufPicture(pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.GdkPixbuf  // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_picture_set_pixbuf(_arg0, _arg1)
}

func (s picture) SetResourcePicture(resourcePath string) {
	var _arg0 *C.GtkPicture // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_picture_set_resource(_arg0, _arg1)
}

func (s picture) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s picture) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s picture) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s picture) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s picture) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s picture) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s picture) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b picture) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
