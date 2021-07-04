// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
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
		{T: externglib.Type(C.gtk_offscreen_window_get_type()), F: marshalOffscreenWindow},
	})
}

// OffscreenWindow: gtkOffscreenWindow is strictly intended to be used for
// obtaining snapshots of widgets that are not part of a normal widget
// hierarchy. Since OffscreenWindow is a toplevel widget you cannot obtain
// snapshots of a full window with it since you cannot pack a toplevel widget in
// another toplevel.
//
// The idea is to take a widget and manually set the state of it, add it to a
// GtkOffscreenWindow and then retrieve the snapshot as a #cairo_surface_t or
// Pixbuf.
//
// GtkOffscreenWindow derives from Window only as an implementation detail.
// Applications should not use any API specific to Window to operate on this
// object. It should be treated as a Bin that has no parent widget.
//
// When contained offscreen widgets are redrawn, GtkOffscreenWindow will emit a
// Widget::damage-event signal.
type OffscreenWindow interface {
	Window

	Pixbuf() gdkpixbuf.Pixbuf

	Surface() *cairo.Surface
}

// offscreenWindow implements the OffscreenWindow class.
type offscreenWindow struct {
	Window
}

// WrapOffscreenWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapOffscreenWindow(obj *externglib.Object) OffscreenWindow {
	return offscreenWindow{
		Window: WrapWindow(obj),
	}
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOffscreenWindow(obj), nil
}

func NewOffscreenWindow() OffscreenWindow {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_offscreen_window_new()

	var _offscreenWindow OffscreenWindow // out

	_offscreenWindow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(OffscreenWindow)

	return _offscreenWindow
}

func (o offscreenWindow) Pixbuf() gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkOffscreenWindow // out
	var _cret *C.GdkPixbuf          // in

	_arg0 = (*C.GtkOffscreenWindow)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_offscreen_window_get_pixbuf(_arg0)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (o offscreenWindow) Surface() *cairo.Surface {
	var _arg0 *C.GtkOffscreenWindow // out
	var _cret *C.cairo_surface_t    // in

	_arg0 = (*C.GtkOffscreenWindow)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_offscreen_window_get_surface(_arg0)

	var _surface *cairo.Surface // out

	_surface = (*cairo.Surface)(unsafe.Pointer(_cret))

	return _surface
}

func (b offscreenWindow) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b offscreenWindow) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b offscreenWindow) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b offscreenWindow) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b offscreenWindow) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b offscreenWindow) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b offscreenWindow) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b offscreenWindow) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b offscreenWindow) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b offscreenWindow) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
