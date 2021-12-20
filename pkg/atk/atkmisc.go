// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_misc_get_type()), F: marshalMiscer},
	})
}

// MiscOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MiscOverrider interface {
	// ThreadsEnter: take the thread mutex for the GUI toolkit, if one exists.
	// (This method is implemented by the toolkit ATK implementation layer; for
	// instance, for GTK+, GAIL implements this via GDK_THREADS_ENTER).
	//
	// Deprecated: Since 2.12.
	ThreadsEnter()
	// ThreadsLeave: release the thread mutex for the GUI toolkit, if one
	// exists. This method, and atk_misc_threads_enter, are needed in some
	// situations by threaded application code which services ATK requests,
	// since fulfilling ATK requests often requires calling into the GUI
	// toolkit. If a long-running or potentially blocking call takes place
	// inside such a block, it should be bracketed by
	// atk_misc_threads_leave/atk_misc_threads_enter calls. (This method is
	// implemented by the toolkit ATK implementation layer; for instance, for
	// GTK+, GAIL implements this via GDK_THREADS_LEAVE).
	//
	// Deprecated: Since 2.12.
	ThreadsLeave()
}

// Misc: set of utility functions for thread locking. This interface and all his
// related methods are deprecated since 2.12.
type Misc struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Misc)(nil)
)

func wrapMisc(obj *externglib.Object) *Misc {
	return &Misc{
		Object: obj,
	}
}

func marshalMiscer(p uintptr) (interface{}, error) {
	return wrapMisc(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ThreadsEnter: take the thread mutex for the GUI toolkit, if one exists. (This
// method is implemented by the toolkit ATK implementation layer; for instance,
// for GTK+, GAIL implements this via GDK_THREADS_ENTER).
//
// Deprecated: Since 2.12.
func (misc *Misc) ThreadsEnter() {
	var _arg0 *C.AtkMisc // out

	_arg0 = (*C.AtkMisc)(unsafe.Pointer(misc.Native()))

	C.atk_misc_threads_enter(_arg0)
	runtime.KeepAlive(misc)
}

// ThreadsLeave: release the thread mutex for the GUI toolkit, if one exists.
// This method, and atk_misc_threads_enter, are needed in some situations by
// threaded application code which services ATK requests, since fulfilling ATK
// requests often requires calling into the GUI toolkit. If a long-running or
// potentially blocking call takes place inside such a block, it should be
// bracketed by atk_misc_threads_leave/atk_misc_threads_enter calls. (This
// method is implemented by the toolkit ATK implementation layer; for instance,
// for GTK+, GAIL implements this via GDK_THREADS_LEAVE).
//
// Deprecated: Since 2.12.
func (misc *Misc) ThreadsLeave() {
	var _arg0 *C.AtkMisc // out

	_arg0 = (*C.AtkMisc)(unsafe.Pointer(misc.Native()))

	C.atk_misc_threads_leave(_arg0)
	runtime.KeepAlive(misc)
}

// MiscGetInstance: obtain the singleton instance of AtkMisc for this
// application.
//
// Deprecated: Since 2.12.
//
// The function returns the following values:
//
//    - misc: singleton instance of AtkMisc for this application.
//
func MiscGetInstance() *Misc {
	var _cret *C.AtkMisc // in

	_cret = C.atk_misc_get_instance()

	var _misc *Misc // out

	_misc = wrapMisc(externglib.Take(unsafe.Pointer(_cret)))

	return _misc
}
