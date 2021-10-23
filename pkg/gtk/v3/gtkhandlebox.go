// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_handle_box_get_type()), F: marshalHandleBoxer},
	})
}

// HandleBoxOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HandleBoxOverrider interface {
	ChildAttached(child Widgetter)
	ChildDetached(child Widgetter)
}

// HandleBox widget allows a portion of a window to be "torn off". It is a bin
// widget which displays its child and a handle that the user can drag to tear
// off a separate window (the “float window”) containing the child widget. A
// thin “ghost” is drawn in the original location of the handlebox. By dragging
// the separate window back to its original location, it can be reattached.
//
// When reattaching, the ghost and float window, must be aligned along one of
// the edges, the “snap edge”. This either can be specified by the application
// programmer explicitly, or GTK+ will pick a reasonable default based on the
// handle position.
//
// To make detaching and reattaching the handlebox as minimally confusing as
// possible to the user, it is important to set the snap edge so that the snap
// edge does not move when the handlebox is deattached. For instance, if the
// handlebox is packed at the bottom of a VBox, then when the handlebox is
// detached, the bottom edge of the handlebox's allocation will remain fixed as
// the height of the handlebox shrinks, so the snap edge should be set to
// GTK_POS_BOTTOM.
//
// > HandleBox has been deprecated. It is very specialized, lacks features > to
// make it useful and most importantly does not fit well into modern >
// application design. Do not use it. There is no replacement.
type HandleBox struct {
	Bin
}

func wrapHandleBox(obj *externglib.Object) *HandleBox {
	return &HandleBox{
		Bin: Bin{
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
					Object: obj,
				},
			},
		},
	}
}

func marshalHandleBoxer(p uintptr) (interface{}, error) {
	return wrapHandleBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHandleBox: create a new handle box.
//
// Deprecated: HandleBox has been deprecated.
func NewHandleBox() *HandleBox {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_handle_box_new()

	var _handleBox *HandleBox // out

	_handleBox = wrapHandleBox(externglib.Take(unsafe.Pointer(_cret)))

	return _handleBox
}

// ChildDetached: whether the handlebox’s child is currently detached.
//
// Deprecated: HandleBox has been deprecated.
func (handleBox *HandleBox) ChildDetached() bool {
	var _arg0 *C.GtkHandleBox // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))

	_cret = C.gtk_handle_box_get_child_detached(_arg0)
	runtime.KeepAlive(handleBox)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HandlePosition gets the handle position of the handle box. See
// gtk_handle_box_set_handle_position().
//
// Deprecated: HandleBox has been deprecated.
func (handleBox *HandleBox) HandlePosition() PositionType {
	var _arg0 *C.GtkHandleBox   // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))

	_cret = C.gtk_handle_box_get_handle_position(_arg0)
	runtime.KeepAlive(handleBox)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// ShadowType gets the type of shadow drawn around the handle box. See
// gtk_handle_box_set_shadow_type().
//
// Deprecated: HandleBox has been deprecated.
func (handleBox *HandleBox) ShadowType() ShadowType {
	var _arg0 *C.GtkHandleBox // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))

	_cret = C.gtk_handle_box_get_shadow_type(_arg0)
	runtime.KeepAlive(handleBox)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// SnapEdge gets the edge used for determining reattachment of the handle box.
// See gtk_handle_box_set_snap_edge().
//
// Deprecated: HandleBox has been deprecated.
func (handleBox *HandleBox) SnapEdge() PositionType {
	var _arg0 *C.GtkHandleBox   // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))

	_cret = C.gtk_handle_box_get_snap_edge(_arg0)
	runtime.KeepAlive(handleBox)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// SetHandlePosition sets the side of the handlebox where the handle is drawn.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - position: side of the handlebox where the handle should be drawn.
//
func (handleBox *HandleBox) SetHandlePosition(position PositionType) {
	var _arg0 *C.GtkHandleBox   // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_handle_box_set_handle_position(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(position)
}

// SetShadowType sets the type of shadow to be drawn around the border of the
// handle box.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - typ: shadow type.
//
func (handleBox *HandleBox) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkHandleBox // out
	var _arg1 C.GtkShadowType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))
	_arg1 = C.GtkShadowType(typ)

	C.gtk_handle_box_set_shadow_type(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(typ)
}

// SetSnapEdge sets the snap edge of a handlebox. The snap edge is the edge of
// the detached child that must be aligned with the corresponding edge of the
// “ghost” left behind when the child was detached to reattach the torn-off
// window. Usually, the snap edge should be chosen so that it stays in the same
// place on the screen when the handlebox is torn off.
//
// If the snap edge is not set, then an appropriate value will be guessed from
// the handle position. If the handle position is GTK_POS_RIGHT or GTK_POS_LEFT,
// then the snap edge will be GTK_POS_TOP, otherwise it will be GTK_POS_LEFT.
//
// Deprecated: HandleBox has been deprecated.
//
// The function takes the following parameters:
//
//    - edge: snap edge, or -1 to unset the value; in which case GTK+ will try
//    to guess an appropriate value in the future.
//
func (handleBox *HandleBox) SetSnapEdge(edge PositionType) {
	var _arg0 *C.GtkHandleBox   // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkHandleBox)(unsafe.Pointer(handleBox.Native()))
	_arg1 = C.GtkPositionType(edge)

	C.gtk_handle_box_set_snap_edge(_arg0, _arg1)
	runtime.KeepAlive(handleBox)
	runtime.KeepAlive(edge)
}

// ConnectChildAttached: this signal is emitted when the contents of the
// handlebox are reattached to the main window.
func (handleBox *HandleBox) ConnectChildAttached(f func(widget Widgetter)) externglib.SignalHandle {
	return handleBox.Connect("child-attached", f)
}

// ConnectChildDetached: this signal is emitted when the contents of the
// handlebox are detached from the main window.
func (handleBox *HandleBox) ConnectChildDetached(f func(widget Widgetter)) externglib.SignalHandle {
	return handleBox.Connect("child-detached", f)
}
