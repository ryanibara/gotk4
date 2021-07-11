// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_widget_get_type()), F: marshalAppChooserWidgetter},
	})
}

// AppChooserWidgetOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type AppChooserWidgetOverrider interface {
	//
	ApplicationActivated(appInfo gio.AppInfor)
	//
	ApplicationSelected(appInfo gio.AppInfor)
	//
	PopulatePopup(menu Menuer, appInfo gio.AppInfor)
}

// AppChooserWidgetter describes AppChooserWidget's methods.
type AppChooserWidgetter interface {
	// DefaultText returns the text that is shown if there are not applications
	// that can handle the content type.
	DefaultText() string
	// ShowAll returns the current value of the AppChooserWidget:show-all
	// property.
	ShowAll() bool
	// ShowDefault returns the current value of the
	// AppChooserWidget:show-default property.
	ShowDefault() bool
	// ShowFallback returns the current value of the
	// AppChooserWidget:show-fallback property.
	ShowFallback() bool
	// ShowOther returns the current value of the AppChooserWidget:show-other
	// property.
	ShowOther() bool
	// ShowRecommended returns the current value of the
	// AppChooserWidget:show-recommended property.
	ShowRecommended() bool
	// SetDefaultText sets the text that is shown if there are not applications
	// that can handle the content type.
	SetDefaultText(text string)
	// SetShowAll sets whether the app chooser should show all applications in a
	// flat list.
	SetShowAll(setting bool)
	// SetShowDefault sets whether the app chooser should show the default
	// handler for the content type in a separate section.
	SetShowDefault(setting bool)
	// SetShowFallback sets whether the app chooser should show related
	// applications for the content type in a separate section.
	SetShowFallback(setting bool)
	// SetShowOther sets whether the app chooser should show applications which
	// are unrelated to the content type.
	SetShowOther(setting bool)
	// SetShowRecommended sets whether the app chooser should show recommended
	// applications for the content type in a separate section.
	SetShowRecommended(setting bool)
}

// AppChooserWidget is a widget for selecting applications. It is the main
// building block for AppChooserDialog. Most applications only need to use the
// latter; but you can use this widget as part of a larger widget if you have
// special needs.
//
// AppChooserWidget offers detailed control over what applications are shown,
// using the AppChooserWidget:show-default, AppChooserWidget:show-recommended,
// AppChooserWidget:show-fallback, AppChooserWidget:show-other and
// AppChooserWidget:show-all properties. See the AppChooser documentation for
// more information about these groups of applications.
//
// To keep track of the selected application, use the
// AppChooserWidget::application-selected and
// AppChooserWidget::application-activated signals.
//
//
// CSS nodes
//
// GtkAppChooserWidget has a single CSS node with name appchooser.
type AppChooserWidget struct {
	Box

	AppChooser
}

var (
	_ AppChooserWidgetter = (*AppChooserWidget)(nil)
	_ gextras.Nativer     = (*AppChooserWidget)(nil)
)

func wrapAppChooserWidget(obj *externglib.Object) AppChooserWidgetter {
	return &AppChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
		},
		AppChooser: AppChooser{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalAppChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserWidget(obj), nil
}

// NewAppChooserWidget creates a new AppChooserWidget for applications that can
// handle content of the given type.
func NewAppChooserWidget(contentType string) *AppChooserWidget {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_widget_new(_arg1)

	var _appChooserWidget *AppChooserWidget // out

	_appChooserWidget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AppChooserWidget)

	return _appChooserWidget
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *AppChooserWidget) Native() uintptr {
	return v.Box.Container.Widget.InitiallyUnowned.Object.Native()
}

// DefaultText returns the text that is shown if there are not applications that
// can handle the content type.
func (self *AppChooserWidget) DefaultText() string {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_default_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ShowAll returns the current value of the AppChooserWidget:show-all property.
func (self *AppChooserWidget) ShowAll() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_all(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefault returns the current value of the AppChooserWidget:show-default
// property.
func (self *AppChooserWidget) ShowDefault() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowFallback returns the current value of the AppChooserWidget:show-fallback
// property.
func (self *AppChooserWidget) ShowFallback() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowOther returns the current value of the AppChooserWidget:show-other
// property.
func (self *AppChooserWidget) ShowOther() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_other(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecommended returns the current value of the
// AppChooserWidget:show-recommended property.
func (self *AppChooserWidget) ShowRecommended() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_recommended(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultText sets the text that is shown if there are not applications that
// can handle the content type.
func (self *AppChooserWidget) SetDefaultText(text string) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_widget_set_default_text(_arg0, _arg1)
}

// SetShowAll sets whether the app chooser should show all applications in a
// flat list.
func (self *AppChooserWidget) SetShowAll(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_all(_arg0, _arg1)
}

// SetShowDefault sets whether the app chooser should show the default handler
// for the content type in a separate section.
func (self *AppChooserWidget) SetShowDefault(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_default(_arg0, _arg1)
}

// SetShowFallback sets whether the app chooser should show related applications
// for the content type in a separate section.
func (self *AppChooserWidget) SetShowFallback(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_fallback(_arg0, _arg1)
}

// SetShowOther sets whether the app chooser should show applications which are
// unrelated to the content type.
func (self *AppChooserWidget) SetShowOther(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_other(_arg0, _arg1)
}

// SetShowRecommended sets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (self *AppChooserWidget) SetShowRecommended(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_recommended(_arg0, _arg1)
}
