// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_overlay_layout_get_type()), F: marshalOverlayLayout},
		{T: externglib.Type(C.gtk_overlay_layout_child_get_type()), F: marshalOverlayLayoutChild},
	})
}

// OverlayLayout: `GtkOverlayLayout` is the layout manager used by `GtkOverlay`.
//
// It places widgets as overlays on top of the main child.
//
// This is not a reusable layout manager, since it expects its widget to be a
// `GtkOverlay`. It only listed here so that its layout properties get
// documented.
type OverlayLayout interface {
	LayoutManager
}

// overlayLayout implements the OverlayLayout class.
type overlayLayout struct {
	LayoutManager
}

// WrapOverlayLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapOverlayLayout(obj *externglib.Object) OverlayLayout {
	return overlayLayout{
		LayoutManager: WrapLayoutManager(obj),
	}
}

func marshalOverlayLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOverlayLayout(obj), nil
}

func NewOverlayLayout() OverlayLayout {
	var _cret *C.GtkLayoutManager // in

	_cret = C.gtk_overlay_layout_new()

	var _overlayLayout OverlayLayout // out

	_overlayLayout = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(OverlayLayout)

	return _overlayLayout
}

// OverlayLayoutChild: `GtkLayoutChild` subclass for children in a
// `GtkOverlayLayout`.
type OverlayLayoutChild interface {
	LayoutChild

	ClipOverlay() bool

	Measure() bool

	SetClipOverlayOverlayLayoutChild(clipOverlay bool)

	SetMeasureOverlayLayoutChild(measure bool)
}

// overlayLayoutChild implements the OverlayLayoutChild class.
type overlayLayoutChild struct {
	LayoutChild
}

// WrapOverlayLayoutChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapOverlayLayoutChild(obj *externglib.Object) OverlayLayoutChild {
	return overlayLayoutChild{
		LayoutChild: WrapLayoutChild(obj),
	}
}

func marshalOverlayLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOverlayLayoutChild(obj), nil
}

func (c overlayLayoutChild) ClipOverlay() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_overlay_layout_child_get_clip_overlay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c overlayLayoutChild) Measure() bool {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_overlay_layout_child_get_measure(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c overlayLayoutChild) SetClipOverlayOverlayLayoutChild(clipOverlay bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(c.Native()))
	if clipOverlay {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_clip_overlay(_arg0, _arg1)
}

func (c overlayLayoutChild) SetMeasureOverlayLayoutChild(measure bool) {
	var _arg0 *C.GtkOverlayLayoutChild // out
	var _arg1 C.gboolean               // out

	_arg0 = (*C.GtkOverlayLayoutChild)(unsafe.Pointer(c.Native()))
	if measure {
		_arg1 = C.TRUE
	}

	C.gtk_overlay_layout_child_set_measure(_arg0, _arg1)
}
