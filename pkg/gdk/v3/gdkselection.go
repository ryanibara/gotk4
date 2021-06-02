// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gdk-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <gdk/gdk.h>
import "C"

// SelectionConvert retrieves the contents of a selection in a given form.
func SelectionConvert(requestor Window, selection Atom, target Atom, time_ uint32) {
	var arg1 *C.GdkWindow
	var arg2 C.GdkAtom
	var arg3 C.GdkAtom
	var arg4 C.guint32

	arg1 = (*C.GdkWindow)(requestor.Native())
	arg2 = (C.GdkAtom)(selection.Native())
	arg3 = (C.GdkAtom)(target.Native())
	arg4 = C.guint32(time_)

	C.gdk_selection_convert(arg1, arg2, arg3, arg4)
}

// SelectionOwnerGet determines the owner of the given selection.
func SelectionOwnerGet(selection Atom) Window {
	var arg1 C.GdkAtom

	arg1 = (C.GdkAtom)(selection.Native())

	ret := C.gdk_selection_owner_get(arg1)

	var ret0 Window

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Window)

	return ret0
}

// SelectionOwnerGetForDisplay: determine the owner of the given selection.
//
// Note that the return value may be owned by a different process if a foreign
// window was previously created for that window, but a new foreign window will
// never be created by this call.
func SelectionOwnerGetForDisplay(display Display, selection Atom) Window {
	var arg1 *C.GdkDisplay
	var arg2 C.GdkAtom

	arg1 = (*C.GdkDisplay)(display.Native())
	arg2 = (C.GdkAtom)(selection.Native())

	ret := C.gdk_selection_owner_get_for_display(arg1, arg2)

	var ret0 Window

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Window)

	return ret0
}

// SelectionOwnerSet sets the owner of the given selection.
func SelectionOwnerSet(owner Window, selection Atom, time_ uint32, sendEvent bool) bool {
	var arg1 *C.GdkWindow
	var arg2 C.GdkAtom
	var arg3 C.guint32
	var arg4 C.gboolean

	arg1 = (*C.GdkWindow)(owner.Native())
	arg2 = (C.GdkAtom)(selection.Native())
	arg3 = C.guint32(time_)
	if sendEvent {
		arg4 = C.TRUE
	}

	ret := C.gdk_selection_owner_set(arg1, arg2, arg3, arg4)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// SelectionOwnerSetForDisplay sets the Window @owner as the current owner of
// the selection @selection.
func SelectionOwnerSetForDisplay(display Display, owner Window, selection Atom, time_ uint32, sendEvent bool) bool {
	var arg1 *C.GdkDisplay
	var arg2 *C.GdkWindow
	var arg3 C.GdkAtom
	var arg4 C.guint32
	var arg5 C.gboolean

	arg1 = (*C.GdkDisplay)(display.Native())
	arg2 = (*C.GdkWindow)(owner.Native())
	arg3 = (C.GdkAtom)(selection.Native())
	arg4 = C.guint32(time_)
	if sendEvent {
		arg5 = C.TRUE
	}

	ret := C.gdk_selection_owner_set_for_display(arg1, arg2, arg3, arg4, arg5)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// SelectionPropertyGet retrieves selection data that was stored by the
// selection data in response to a call to gdk_selection_convert(). This
// function will not be used by applications, who should use the Clipboard API
// instead.
func SelectionPropertyGet(requestor Window, data byte, propType *Atom, propFormat int) int {
	var arg1 *C.GdkWindow
	var arg2 **C.guchar
	var arg3 *C.GdkAtom
	var arg4 *C.gint

	arg1 = (*C.GdkWindow)(requestor.Native())
	arg2 = (**C.guchar)(data)
	arg3 = (*C.GdkAtom)(propType.Native())
	arg4 = (*C.gint)(propFormat)

	ret := C.gdk_selection_property_get(arg1, arg2, arg3, arg4)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// SelectionSendNotify sends a response to SelectionRequest event.
func SelectionSendNotify(requestor Window, selection Atom, target Atom, property Atom, time_ uint32) {
	var arg1 *C.GdkWindow
	var arg2 C.GdkAtom
	var arg3 C.GdkAtom
	var arg4 C.GdkAtom
	var arg5 C.guint32

	arg1 = (*C.GdkWindow)(requestor.Native())
	arg2 = (C.GdkAtom)(selection.Native())
	arg3 = (C.GdkAtom)(target.Native())
	arg4 = (C.GdkAtom)(property.Native())
	arg5 = C.guint32(time_)

	C.gdk_selection_send_notify(arg1, arg2, arg3, arg4, arg5)
}

// SelectionSendNotifyForDisplay: send a response to SelectionRequest event.
func SelectionSendNotifyForDisplay(display Display, requestor Window, selection Atom, target Atom, property Atom, time_ uint32) {
	var arg1 *C.GdkDisplay
	var arg2 *C.GdkWindow
	var arg3 C.GdkAtom
	var arg4 C.GdkAtom
	var arg5 C.GdkAtom
	var arg6 C.guint32

	arg1 = (*C.GdkDisplay)(display.Native())
	arg2 = (*C.GdkWindow)(requestor.Native())
	arg3 = (C.GdkAtom)(selection.Native())
	arg4 = (C.GdkAtom)(target.Native())
	arg5 = (C.GdkAtom)(property.Native())
	arg6 = C.guint32(time_)

	C.gdk_selection_send_notify_for_display(arg1, arg2, arg3, arg4, arg5, arg6)
}
