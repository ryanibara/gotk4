// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_widget_paintable_get_type()), F: marshalWidgetPaintable},
	})
}

// WidgetPaintable: `GtkWidgetPaintable` is a `GdkPaintable` that displays the
// contents of a widget.
//
// `GtkWidgetPaintable` will also take care of the widget not being in a state
// where it can be drawn (like when it isn't shown) and just draw nothing or
// where it does not have a size (like when it is hidden) and report no size in
// that case.
//
// Of course, `GtkWidgetPaintable` allows you to monitor widgets for size
// changes by emitting the [signal@Gdk.Paintable::invalidate-size] signal
// whenever the size of the widget changes as well as for visual changes by
// emitting the [signal@Gdk.Paintable::invalidate-contents] signal whenever the
// widget changes.
//
// You can use a `GtkWidgetPaintable` everywhere a `GdkPaintable` is allowed,
// including using it on a `GtkPicture` (or one of its parents) that it was set
// on itself via gtk_picture_set_paintable(). The paintable will take care of
// recursion when this happens. If you do this however, ensure that the
// [property@Gtk.Picture:can-shrink] property is set to true or you might end up
// with an infinitely growing widget.
type WidgetPaintable interface {
	gextras.Objector
	gdk.Paintable

	// Widget returns the widget that is observed or nil if none.
	Widget() Widget
	// SetWidget sets the widget that should be observed.
	SetWidget(widget Widget)
}

// widgetPaintable implements the WidgetPaintable interface.
type widgetPaintable struct {
	gextras.Objector
	gdk.Paintable
}

var _ WidgetPaintable = (*widgetPaintable)(nil)

// WrapWidgetPaintable wraps a GObject to the right type. It is
// primarily used internally.
func WrapWidgetPaintable(obj *externglib.Object) WidgetPaintable {
	return WidgetPaintable{
		Objector:      obj,
		gdk.Paintable: gdk.WrapPaintable(obj),
	}
}

func marshalWidgetPaintable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWidgetPaintable(obj), nil
}

// NewWidgetPaintable constructs a class WidgetPaintable.
func NewWidgetPaintable(widget Widget) WidgetPaintable {
	var _arg1 *C.GtkWidget

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _cret C.GtkWidgetPaintable

	cret = C.gtk_widget_paintable_new(_arg1)

	var _widgetPaintable WidgetPaintable

	_widgetPaintable = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(WidgetPaintable)

	return _widgetPaintable
}

// Widget returns the widget that is observed or nil if none.
func (s widgetPaintable) Widget() Widget {
	var _arg0 *C.GtkWidgetPaintable

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(s.Native()))

	var _cret *C.GtkWidget

	cret = C.gtk_widget_paintable_get_widget(_arg0)

	var _widget Widget

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetWidget sets the widget that should be observed.
func (s widgetPaintable) SetWidget(widget Widget) {
	var _arg0 *C.GtkWidgetPaintable
	var _arg1 *C.GtkWidget

	_arg0 = (*C.GtkWidgetPaintable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_widget_paintable_set_widget(_arg0, _arg1)
}