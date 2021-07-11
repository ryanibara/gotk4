// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooserer},
	})
}

// AppChooserer describes AppChooser's methods.
type AppChooserer interface {
	// AppInfo returns the currently selected application.
	AppInfo() *gio.AppInfo
	// ContentType returns the content type for which the `GtkAppChooser` shows
	// applications.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

// AppChooser: `GtkAppChooser` is an interface for widgets which allow the user
// to choose an application.
//
// The main objects that implement this interface are
// [class@Gtk.AppChooserWidget], [class@Gtk.AppChooserDialog] and
// [class@Gtk.AppChooserButton].
//
// Applications are represented by GIO `GAppInfo` objects here. GIO has a
// concept of recommended and fallback applications for a given content type.
// Recommended applications are those that claim to handle the content type
// itself, while fallback also includes applications that handle a more generic
// content type. GIO also knows the default and last-used application for a
// given content type. The `GtkAppChooserWidget` provides detailed control over
// whether the shown list of applications should include default, recommended or
// fallback applications.
//
// To obtain the application that has been selected in a `GtkAppChooser`, use
// [method@Gtk.AppChooser.get_app_info].
type AppChooser struct {
	Widget
}

var (
	_ AppChooserer    = (*AppChooser)(nil)
	_ gextras.Nativer = (*AppChooser)(nil)
)

func wrapAppChooser(obj *externglib.Object) AppChooserer {
	return &AppChooser{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalAppChooserer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooser(obj), nil
}

// AppInfo returns the currently selected application.
func (self *AppChooser) AppInfo() *gio.AppInfo {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.GAppInfo      // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_app_info(_arg0)

	var _appInfo *gio.AppInfo // out

	_appInfo = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*gio.AppInfo)

	return _appInfo
}

// ContentType returns the content type for which the `GtkAppChooser` shows
// applications.
func (self *AppChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Refresh reloads the list of applications.
func (self *AppChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_refresh(_arg0)
}
