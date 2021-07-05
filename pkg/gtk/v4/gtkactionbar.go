// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_action_bar_get_type()), F: marshalActionBar},
	})
}

// ActionBar: `GtkActionBar` is designed to present contextual actions.
//
// !An example GtkActionBar (action-bar.png)
//
// It is expected to be displayed below the content and expand horizontally to
// fill the area.
//
// It allows placing children at the start or the end. In addition, it contains
// an internal centered box which is centered with respect to the full width of
// the box, even if the children at either side take up different amounts of
// space.
//
//
// CSS nodes
//
// `GtkActionBar` has a single CSS node with name actionbar.
type ActionBar interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget

	// CenterWidget retrieves the center bar widget of the bar.
	CenterWidget() Widget
	// Revealed gets whether the contents of the action bar are revealed.
	Revealed() bool
	// PackEndActionBar adds @child to @action_bar, packed with reference to the
	// end of the @action_bar.
	PackEndActionBar(child Widget)
	// PackStartActionBar adds @child to @action_bar, packed with reference to
	// the start of the @action_bar.
	PackStartActionBar(child Widget)
	// RemoveActionBar removes a child from @action_bar.
	RemoveActionBar(child Widget)
	// SetCenterWidgetActionBar sets the center widget for the `GtkActionBar`.
	SetCenterWidgetActionBar(centerWidget Widget)
	// SetRevealedActionBar reveals or conceals the content of the action bar.
	//
	// Note: this does not show or hide @action_bar in the
	// [property@Gtk.Widget:visible] sense, so revealing has no effect if the
	// action bar is hidden.
	SetRevealedActionBar(revealed bool)
}

// actionBar implements the ActionBar class.
type actionBar struct {
	Widget
}

// WrapActionBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapActionBar(obj *externglib.Object) ActionBar {
	return actionBar{
		Widget: WrapWidget(obj),
	}
}

func marshalActionBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActionBar(obj), nil
}

// NewActionBar creates a new `GtkActionBar` widget.
func NewActionBar() ActionBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_action_bar_new()

	var _actionBar ActionBar // out

	_actionBar = WrapActionBar(externglib.Take(unsafe.Pointer(_cret)))

	return _actionBar
}

func (a actionBar) CenterWidget() Widget {
	var _arg0 *C.GtkActionBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_bar_get_center_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a actionBar) Revealed() bool {
	var _arg0 *C.GtkActionBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_action_bar_get_revealed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a actionBar) PackEndActionBar(child Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_end(_arg0, _arg1)
}

func (a actionBar) PackStartActionBar(child Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_pack_start(_arg0, _arg1)
}

func (a actionBar) RemoveActionBar(child Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_action_bar_remove(_arg0, _arg1)
}

func (a actionBar) SetCenterWidgetActionBar(centerWidget Widget) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(centerWidget.Native()))

	C.gtk_action_bar_set_center_widget(_arg0, _arg1)
}

func (a actionBar) SetRevealedActionBar(revealed bool) {
	var _arg0 *C.GtkActionBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkActionBar)(unsafe.Pointer(a.Native()))
	if revealed {
		_arg1 = C.TRUE
	}

	C.gtk_action_bar_set_revealed(_arg0, _arg1)
}

func (a actionBar) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(a))
}

func (a actionBar) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(a))
}

func (a actionBar) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(a))
}
