// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_controls_get_type()), F: marshalWindowControls},
	})
}

// WindowControls: `GtkWindowControls` shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// `GtkWindowControls` only displays start or end side of the controls (see
// [property@Gtk.WindowControls:side]), so it's intended to be always used in
// pair with another `GtkWindowControls` for the opposite side, for example:
//
// “`xml <object class="GtkBox"> <child> <object class="GtkWindowControls">
// <property name="side">start</property> </object> </child>
//
//    ...
//
//    <child>
//      <object class="GtkWindowControls">
//        <property name="side">end</property>
//      </object>
//    </child>
//
// </object> “`
//
//
// CSS nodes
//
// “` windowcontrols ├── [image.icon] ├── [button.minimize] ├──
// [button.maximize] ╰── [button.close] “`
//
// A `GtkWindowControls`' CSS node is called windowcontrols. It contains
// subnodes corresponding to each title button. Which of the title buttons exist
// and where they are placed exactly depends on the desktop environment and
// [property@Gtk.WindowControls:decoration-layout] value.
//
// When [property@Gtk.WindowControls:empty] is true, it gets the .empty style
// class.
//
//
// Accessibility
//
// `GtkWindowControls` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// DecorationLayout gets the decoration layout of this `GtkWindowControls`.
	DecorationLayout() string
	// Empty gets whether the widget has any window buttons.
	Empty() bool
	// SetDecorationLayout sets the decoration layout for the title buttons.
	//
	// This overrides the [property@Gtk.Settings:gtk-decoration-layout] setting.
	//
	// The format of the string is button names, separated by commas. A colon
	// separates the buttons that should appear on the left from those on the
	// right. Recognized button names are minimize, maximize, close and icon
	// (the window icon).
	//
	// For example, “icon:minimize,maximize,close” specifies a icon on the left,
	// and minimize, maximize and close buttons on the right.
	//
	// If [property@Gtk.WindowControls:side] value is @GTK_PACK_START, @self
	// will display the part before the colon, otherwise after that.
	SetDecorationLayout(layout string)
	// SetSide determines which part of decoration layout the
	// `GtkWindowControls` uses.
	//
	// See [property@Gtk.WindowControls:decoration-layout].
	SetSide(side PackType)
}

// windowControls implements the WindowControls interface.
type windowControls struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ WindowControls = (*windowControls)(nil)

// WrapWindowControls wraps a GObject to the right type. It is
// primarily used internally.
func WrapWindowControls(obj *externglib.Object) WindowControls {
	return WindowControls{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalWindowControls(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindowControls(obj), nil
}

// DecorationLayout gets the decoration layout of this `GtkWindowControls`.
func (s windowControls) DecorationLayout() string {
	var _arg0 *C.GtkWindowControls

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))

	var _cret *C.char

	_cret = C.gtk_window_controls_get_decoration_layout(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Empty gets whether the widget has any window buttons.
func (s windowControls) Empty() bool {
	var _arg0 *C.GtkWindowControls

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean

	_cret = C.gtk_window_controls_get_empty(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetDecorationLayout sets the decoration layout for the title buttons.
//
// This overrides the [property@Gtk.Settings:gtk-decoration-layout] setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the
// right. Recognized button names are minimize, maximize, close and icon
// (the window icon).
//
// For example, “icon:minimize,maximize,close” specifies a icon on the left,
// and minimize, maximize and close buttons on the right.
//
// If [property@Gtk.WindowControls:side] value is @GTK_PACK_START, @self
// will display the part before the colon, otherwise after that.
func (s windowControls) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkWindowControls
	var _arg1 *C.char

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(layout))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_window_controls_set_decoration_layout(_arg0, _arg1)
}

// SetSide determines which part of decoration layout the
// `GtkWindowControls` uses.
//
// See [property@Gtk.WindowControls:decoration-layout].
func (s windowControls) SetSide(side PackType) {
	var _arg0 *C.GtkWindowControls
	var _arg1 C.GtkPackType

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkPackType)(side)

	C.gtk_window_controls_set_side(_arg0, _arg1)
}
