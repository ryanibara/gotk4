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
// extern void _gotk4_gtk4_CellRendererAccel_ConnectAccelCleared(gpointer, gchar*, guintptr);
import "C"

// glib.Type values for gtkcellrendereraccel.go.
var (
	GTypeCellRendererAccelMode = coreglib.Type(C.gtk_cell_renderer_accel_mode_get_type())
	GTypeCellRendererAccel     = coreglib.Type(C.gtk_cell_renderer_accel_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellRendererAccelMode, F: marshalCellRendererAccelMode},
		{T: GTypeCellRendererAccel, F: marshalCellRendererAccel},
	})
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
func _gotk4_gtk4_CellRendererAccel_ConnectAccelCleared(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
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
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "CellRendererAccel").InvokeMethod("new_CellRendererAccel", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _cellRendererAccel *CellRendererAccel // out

	_cellRendererAccel = wrapCellRendererAccel(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererAccel
}
