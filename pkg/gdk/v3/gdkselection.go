// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// SelectionConvert retrieves the contents of a selection in a given form.
func SelectionConvert(requestor Window, selection Atom, target Atom, time_ uint32) {
	var _arg1 *C.GdkWindow
	var _arg2 C.GdkAtom
	var _arg3 C.GdkAtom
	var _arg4 C.guint32

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	_arg3 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	_arg4 = C.guint32(time_)

	C.gdk_selection_convert(_arg1, _arg2, _arg3, _arg4)
}

// SelectionOwnerSet sets the owner of the given selection.
func SelectionOwnerSet(owner Window, selection Atom, time_ uint32, sendEvent bool) bool {
	var _arg1 *C.GdkWindow
	var _arg2 C.GdkAtom
	var _arg3 C.guint32
	var _arg4 C.gboolean

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(owner.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	_arg3 = C.guint32(time_)
	if sendEvent {
		_arg4 = C.gboolean(1)
	}

	var _cret C.gboolean

	_cret = C.gdk_selection_owner_set(_arg1, _arg2, _arg3, _arg4)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SelectionOwnerSetForDisplay sets the Window @owner as the current owner of
// the selection @selection.
func SelectionOwnerSetForDisplay(display Display, owner Window, selection Atom, time_ uint32, sendEvent bool) bool {
	var _arg1 *C.GdkDisplay
	var _arg2 *C.GdkWindow
	var _arg3 C.GdkAtom
	var _arg4 C.guint32
	var _arg5 C.gboolean

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(owner.Native()))
	_arg3 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	_arg4 = C.guint32(time_)
	if sendEvent {
		_arg5 = C.gboolean(1)
	}

	var _cret C.gboolean

	_cret = C.gdk_selection_owner_set_for_display(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SelectionPropertyGet retrieves selection data that was stored by the
// selection data in response to a call to gdk_selection_convert(). This
// function will not be used by applications, who should use the Clipboard API
// instead.
func SelectionPropertyGet(requestor Window, data **byte, propType *Atom, propFormat *int) int {
	var _arg1 *C.GdkWindow
	var _arg2 **C.guchar
	var _arg3 *C.GdkAtom
	var _arg4 *C.gint

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	_arg2 = **C.guchar(data)
	_arg3 = (*C.GdkAtom)(unsafe.Pointer(propType.Native()))
	_arg4 = *C.gint(propFormat)

	var _cret C.gint

	_cret = C.gdk_selection_property_get(_arg1, _arg2, _arg3, _arg4)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// SelectionSendNotify sends a response to SelectionRequest event.
func SelectionSendNotify(requestor Window, selection Atom, target Atom, property Atom, time_ uint32) {
	var _arg1 *C.GdkWindow
	var _arg2 C.GdkAtom
	var _arg3 C.GdkAtom
	var _arg4 C.GdkAtom
	var _arg5 C.guint32

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	_arg2 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	_arg3 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	_arg4 = (C.GdkAtom)(unsafe.Pointer(property.Native()))
	_arg5 = C.guint32(time_)

	C.gdk_selection_send_notify(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// SelectionSendNotifyForDisplay: send a response to SelectionRequest event.
func SelectionSendNotifyForDisplay(display Display, requestor Window, selection Atom, target Atom, property Atom, time_ uint32) {
	var _arg1 *C.GdkDisplay
	var _arg2 *C.GdkWindow
	var _arg3 C.GdkAtom
	var _arg4 C.GdkAtom
	var _arg5 C.GdkAtom
	var _arg6 C.guint32

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(requestor.Native()))
	_arg3 = (C.GdkAtom)(unsafe.Pointer(selection.Native()))
	_arg4 = (C.GdkAtom)(unsafe.Pointer(target.Native()))
	_arg5 = (C.GdkAtom)(unsafe.Pointer(property.Native()))
	_arg6 = C.guint32(time_)

	C.gdk_selection_send_notify_for_display(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}
