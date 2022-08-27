// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
import "C"

// SetChild sets child as accessible child of plug and plug as accessible parent
// of child. child can be NULL.
//
// In some cases, one can not use the AtkPlug type directly as accessible object
// for the toplevel widget of the application. For instance in the gtk case,
// GtkPlugAccessible can not inherit both from GtkWindowAccessible and from
// AtkPlug. In such a case, one can create, in addition to the standard
// accessible object for the toplevel widget, an AtkPlug object, and make the
// former the child of the latter by calling atk_plug_set_child().
//
// The function takes the following parameters:
//
//    - child to be set as accessible child of plug.
//
func (plug *Plug) SetChild(child *AtkObject) {
	var _arg0 *C.AtkPlug   // out
	var _arg1 *C.AtkObject // out

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(coreglib.InternObject(plug).Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.atk_plug_set_child(_arg0, _arg1)
	runtime.KeepAlive(plug)
	runtime.KeepAlive(child)
}
