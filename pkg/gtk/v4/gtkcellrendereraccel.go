// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_CellRendererAccel_ConnectAccelCleared(gpointer, void*, guintptr);
import "C"

// GTypeCellRendererAccelMode returns the GType for the type CellRendererAccelMode.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellRendererAccelMode() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CellRendererAccelMode").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellRendererAccelMode)
	return gtype
}

// GTypeCellRendererAccel returns the GType for the type CellRendererAccel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellRendererAccel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CellRendererAccel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellRendererAccel)
	return gtype
}

// CellRendererAccelMode determines if the edited accelerators are GTK
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode C.gint

const (
	// CellRendererAccelModeGTK: GTK accelerators mode.
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// CellRendererAccelModeOther: other accelerator mode.
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CellRendererAccelMode.
func (c CellRendererAccelMode) String() string {
	switch c {
	case CellRendererAccelModeGTK:
		return "GTK"
	case CellRendererAccelModeOther:
		return "Other"
	default:
		return fmt.Sprintf("CellRendererAccelMode(%d)", c)
	}
}

// CellRendererAccel renders a keyboard accelerator in a cell
//
// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like Control + a). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
type CellRendererAccel struct {
	_ [0]func() // equal guard
	CellRendererText
}

var (
	_ CellRendererer = (*CellRendererAccel)(nil)
)

func wrapCellRendererAccel(obj *coreglib.Object) *CellRendererAccel {
	return &CellRendererAccel{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	return wrapCellRendererAccel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_CellRendererAccel_ConnectAccelCleared
func _gotk4_gtk4_CellRendererAccel_ConnectAccelCleared(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(pathString string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pathString string))
	}

	var _pathString string // out

	_pathString = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_pathString)
}

// ConnectAccelCleared gets emitted when the user has removed the accelerator.
func (v *CellRendererAccel) ConnectAccelCleared(f func(pathString string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "accel-cleared", false, unsafe.Pointer(C._gotk4_gtk4_CellRendererAccel_ConnectAccelCleared), f)
}

// NewCellRendererAccel creates a new CellRendererAccel.
//
// The function returns the following values:
//
//    - cellRendererAccel: new cell renderer.
//
func NewCellRendererAccel() *CellRendererAccel {
	_info := girepository.MustFind("Gtk", "CellRendererAccel")
	_gret := _info.InvokeClassMethod("new_CellRendererAccel", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _cellRendererAccel *CellRendererAccel // out

	_cellRendererAccel = wrapCellRendererAccel(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererAccel
}
