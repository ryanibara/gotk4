// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeAppChooser returns the GType for the type AppChooser.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAppChooser() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AppChooser").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAppChooser)
	return gtype
}

// AppChooser: GtkAppChooser is an interface for widgets which allow the user to
// choose an application.
//
// The main objects that implement this interface are gtk.AppChooserWidget,
// gtk.AppChooserDialog and gtk.AppChooserButton.
//
// Applications are represented by GIO GAppInfo objects here. GIO has a concept
// of recommended and fallback applications for a given content type.
// Recommended applications are those that claim to handle the content type
// itself, while fallback also includes applications that handle a more generic
// content type. GIO also knows the default and last-used application for a
// given content type. The GtkAppChooserWidget provides detailed control over
// whether the shown list of applications should include default, recommended or
// fallback applications.
//
// To obtain the application that has been selected in a GtkAppChooser, use
// gtk.AppChooser.GetAppInfo().
//
// AppChooser wraps an interface. This means the user can get the
// underlying type by calling Cast().
type AppChooser struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*AppChooser)(nil)
)

// AppChooserer describes AppChooser's interface methods.
type AppChooserer interface {
	coreglib.Objector

	// AppInfo returns the currently selected application.
	AppInfo() *gio.AppInfo
	// ContentType returns the content type for which the GtkAppChooser shows
	// applications.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

var _ AppChooserer = (*AppChooser)(nil)

func wrapAppChooser(obj *coreglib.Object) *AppChooser {
	return &AppChooser{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
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

func marshalAppChooser(p uintptr) (interface{}, error) {
	return wrapAppChooser(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppInfo returns the currently selected application.
//
// The function returns the following values:
//
//    - appInfo (optional): GAppInfo for the currently selected application, or
//      NULL if none is selected. Free with g_object_unref().
//
func (self *AppChooser) AppInfo() *gio.AppInfo {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooser")
	_gret := _info.InvokeIfaceMethod("get_app_info", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _appInfo *gio.AppInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
			_appInfo = &gio.AppInfo{
				Object: obj,
			}
		}
	}

	return _appInfo
}

// ContentType returns the content type for which the GtkAppChooser shows
// applications.
//
// The function returns the following values:
//
//    - utf8: content type of self. Free with g_free().
//
func (self *AppChooser) ContentType() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooser")
	_gret := _info.InvokeIfaceMethod("get_content_type", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret))))

	return _utf8
}

// Refresh reloads the list of applications.
func (self *AppChooser) Refresh() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "AppChooser")
	_info.InvokeIfaceMethod("refresh", _args[:], nil)

	runtime.KeepAlive(self)
}
