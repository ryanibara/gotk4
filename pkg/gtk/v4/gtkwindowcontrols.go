// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

// GType values.
var (
	GTypeWindowControls = coreglib.Type(C.gtk_window_controls_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeWindowControls, F: marshalWindowControls},
	})
}

// WindowControlsOverrider contains methods that are overridable.
type WindowControlsOverrider interface {
}

// WindowControls: GtkWindowControls shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// GtkWindowControls only displays start or end side of the controls (see
// gtk.WindowControls:side), so it's intended to be always used in pair with
// another GtkWindowControls for the opposite side, for example:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">start</property>
//        </object>
//      </child>
//
//      ...
//
//      <child>
//        <object class="GtkWindowControls">
//          <property name="side">end</property>
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    windowcontrols
//    ├── [image.icon]
//    ├── [button.minimize]
//    ├── [button.maximize]
//    ╰── [button.close]
//
//
// A GtkWindowControls' CSS node is called windowcontrols. It contains subnodes
// corresponding to each title button. Which of the title buttons exist and
// where they are placed exactly depends on the desktop environment and
// gtk.WindowControls:decoration-layout value.
//
// When gtk.WindowControls:empty is TRUE, it gets the .empty style class.
//
//
// Accessibility
//
// GtkWindowControls uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*WindowControls)(nil)
)

func classInitWindowControlser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapWindowControls(obj *coreglib.Object) *WindowControls {
	return &WindowControls{
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

func marshalWindowControls(p uintptr) (interface{}, error) {
	return wrapWindowControls(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewWindowControls creates a new GtkWindowControls.
//
// The function takes the following parameters:
//
//    - side: side.
//
// The function returns the following values:
//
//    - windowControls: new GtkWindowControls.
//
func NewWindowControls(side PackType) *WindowControls {
	var _arg1 C.GtkPackType // out
	var _cret *C.GtkWidget  // in

	_arg1 = C.GtkPackType(side)

	_cret = C.gtk_window_controls_new(_arg1)
	runtime.KeepAlive(side)

	var _windowControls *WindowControls // out

	_windowControls = wrapWindowControls(coreglib.Take(unsafe.Pointer(_cret)))

	return _windowControls
}

// DecorationLayout gets the decoration layout of this GtkWindowControls.
//
// The function returns the following values:
//
//    - utf8 (optional): decoration layout or NULL if it is unset.
//
func (self *WindowControls) DecorationLayout() string {
	var _arg0 *C.GtkWindowControls // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_window_controls_get_decoration_layout(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Empty gets whether the widget has any window buttons.
//
// The function returns the following values:
//
//    - ok: TRUE if the widget has window buttons, otherwise FALSE.
//
func (self *WindowControls) Empty() bool {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_window_controls_get_empty(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Side gets the side to which this GtkWindowControls instance belongs.
//
// The function returns the following values:
//
//    - packType: side.
//
func (self *WindowControls) Side() PackType {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.GtkPackType        // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_cret = C.gtk_window_controls_get_side(_arg0)
	runtime.KeepAlive(self)

	var _packType PackType // out

	_packType = PackType(_cret)

	return _packType
}

// SetDecorationLayout sets the decoration layout for the title buttons.
//
// This overrides the gtk.Settings:gtk-decoration-layout setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close and icon (the window
// icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon on the left, and
// minimize, maximize and close buttons on the right.
//
// If gtk.WindowControls:side value is GTK_PACK_START, self will display the
// part before the colon, otherwise after that.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (self *WindowControls) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if layout != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_window_controls_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(layout)
}

// SetSide determines which part of decoration layout the GtkWindowControls
// uses.
//
// See gtk.WindowControls:decoration-layout.
//
// The function takes the following parameters:
//
//    - side: side.
//
func (self *WindowControls) SetSide(side PackType) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 C.GtkPackType        // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.GtkPackType(side)

	C.gtk_window_controls_set_side(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(side)
}
