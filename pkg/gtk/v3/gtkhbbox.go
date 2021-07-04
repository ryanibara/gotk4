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
		{T: externglib.Type(C.gtk_hbutton_box_get_type()), F: marshalHButtonBox},
	})
}

type HButtonBox interface {
	ButtonBox
}

// hButtonBox implements the HButtonBox class.
type hButtonBox struct {
	ButtonBox
}

// WrapHButtonBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapHButtonBox(obj *externglib.Object) HButtonBox {
	return hButtonBox{
		ButtonBox: WrapButtonBox(obj),
	}
}

func marshalHButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHButtonBox(obj), nil
}

func NewHButtonBox() HButtonBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hbutton_box_new()

	var _hButtonBox HButtonBox // out

	_hButtonBox = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(HButtonBox)

	return _hButtonBox
}

func (b hButtonBox) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b hButtonBox) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b hButtonBox) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b hButtonBox) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b hButtonBox) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b hButtonBox) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b hButtonBox) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b hButtonBox) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b hButtonBox) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b hButtonBox) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o hButtonBox) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o hButtonBox) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
