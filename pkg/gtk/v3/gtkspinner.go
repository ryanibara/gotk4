// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_spinner_get_type()), F: marshalSpinner},
	})
}

// Spinner: a GtkSpinner widget displays an icon-size spinning animation. It is
// often used as an alternative to a ProgressBar for displaying indefinite
// activity, instead of actual progress.
//
// To start the animation, use gtk_spinner_start(), to stop it use
// gtk_spinner_stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner interface {
	Widget

	StartSpinner()

	StopSpinner()
}

// spinner implements the Spinner class.
type spinner struct {
	Widget
}

// WrapSpinner wraps a GObject to the right type. It is
// primarily used internally.
func WrapSpinner(obj *externglib.Object) Spinner {
	return spinner{
		Widget: WrapWidget(obj),
	}
}

func marshalSpinner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSpinner(obj), nil
}

func NewSpinner() Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner Spinner // out

	_spinner = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Spinner)

	return _spinner
}

func (s spinner) StartSpinner() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(s.Native()))

	C.gtk_spinner_start(_arg0)
}

func (s spinner) StopSpinner() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(s.Native()))

	C.gtk_spinner_stop(_arg0)
}

func (b spinner) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b spinner) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b spinner) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b spinner) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b spinner) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b spinner) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b spinner) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b spinner) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b spinner) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b spinner) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
