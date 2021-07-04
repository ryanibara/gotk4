// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_media_controls_get_type()), F: marshalMediaControls},
	})
}

// MediaControls: `GtkMediaControls` is a widget to show controls for a video.
//
// !An example GtkMediaControls (media-controls.png)
//
// Usually, `GtkMediaControls` is used as part of [class@Gtk.Video].
type MediaControls interface {
	Widget

	MediaStream() MediaStream

	SetMediaStreamMediaControls(stream MediaStream)
}

// mediaControls implements the MediaControls class.
type mediaControls struct {
	Widget
}

// WrapMediaControls wraps a GObject to the right type. It is
// primarily used internally.
func WrapMediaControls(obj *externglib.Object) MediaControls {
	return mediaControls{
		Widget: WrapWidget(obj),
	}
}

func marshalMediaControls(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaControls(obj), nil
}

func NewMediaControls(stream MediaStream) MediaControls {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(stream.Native()))

	_cret = C.gtk_media_controls_new(_arg1)

	var _mediaControls MediaControls // out

	_mediaControls = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MediaControls)

	return _mediaControls
}

func (c mediaControls) MediaStream() MediaStream {
	var _arg0 *C.GtkMediaControls // out
	var _cret *C.GtkMediaStream   // in

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_media_controls_get_media_stream(_arg0)

	var _mediaStream MediaStream // out

	_mediaStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(MediaStream)

	return _mediaStream
}

func (c mediaControls) SetMediaStreamMediaControls(stream MediaStream) {
	var _arg0 *C.GtkMediaControls // out
	var _arg1 *C.GtkMediaStream   // out

	_arg0 = (*C.GtkMediaControls)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer(stream.Native()))

	C.gtk_media_controls_set_media_stream(_arg0, _arg1)
}

func (s mediaControls) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s mediaControls) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s mediaControls) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s mediaControls) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s mediaControls) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s mediaControls) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s mediaControls) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b mediaControls) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}
