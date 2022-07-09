// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk4_CellEditableIface_editing_done(void*);
// extern void _gotk4_gtk4_CellEditableIface_remove_widget(void*);
// extern void _gotk4_gtk4_CellEditableIface_start_editing(void*, void*);
// extern void _gotk4_gtk4_CellEditable_ConnectEditingDone(gpointer, guintptr);
// extern void _gotk4_gtk4_CellEditable_ConnectRemoveWidget(gpointer, guintptr);
import "C"

// GTypeCellEditable returns the GType for the type CellEditable.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCellEditable() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CellEditable").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCellEditable)
	return gtype
}

// CellEditableOverrider contains methods that are overridable.
type CellEditableOverrider interface {
	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
	// StartEditing begins editing on a cell_editable.
	//
	// The CellRenderer for the cell creates and returns a CellEditable from
	// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
	//
	// gtk_cell_editable_start_editing() can then set up cell_editable suitably
	// for editing a cell, e.g. making the Esc key emit
	// CellEditable::editing-done.
	//
	// Note that the cell_editable is created on-demand for the current edit;
	// its lifetime is temporary and does not persist across other edits and/or
	// cells.
	//
	// The function takes the following parameters:
	//
	//    - event (optional) that began the editing process, or NULL if editing
	//      was initiated programmatically.
	//
	StartEditing(event gdk.Eventer)
}

// CellEditable: interface for widgets that can be used for editing cells
//
// The CellEditable interface must be implemented for widgets to be usable to
// edit the contents of a TreeView cell. It provides a way to specify how
// temporary widgets should be configured for editing, get the new value, etc.
//
// CellEditable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type CellEditable struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*CellEditable)(nil)
)

// CellEditabler describes CellEditable's interface methods.
type CellEditabler interface {
	coreglib.Objector

	// EditingDone emits the CellEditable::editing-done signal.
	EditingDone()
	// RemoveWidget emits the CellEditable::remove-widget signal.
	RemoveWidget()
	// StartEditing begins editing on a cell_editable.
	StartEditing(event gdk.Eventer)

	// Editing-done: this signal is a sign for the cell renderer to update its
	// value from the cell_editable.
	ConnectEditingDone(func()) coreglib.SignalHandle
	// Remove-widget: this signal is meant to indicate that the cell is finished
	// editing, and the cell_editable widget is being removed and may
	// subsequently be destroyed.
	ConnectRemoveWidget(func()) coreglib.SignalHandle
}

var _ CellEditabler = (*CellEditable)(nil)

func ifaceInitCellEditabler(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Gtk", "CellEditableIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("editing_done"))) = unsafe.Pointer(C._gotk4_gtk4_CellEditableIface_editing_done)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("remove_widget"))) = unsafe.Pointer(C._gotk4_gtk4_CellEditableIface_remove_widget)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), pclass.StructFieldOffset("start_editing"))) = unsafe.Pointer(C._gotk4_gtk4_CellEditableIface_start_editing)
}

//export _gotk4_gtk4_CellEditableIface_editing_done
func _gotk4_gtk4_CellEditableIface_editing_done(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellEditableOverrider)

	iface.EditingDone()
}

//export _gotk4_gtk4_CellEditableIface_remove_widget
func _gotk4_gtk4_CellEditableIface_remove_widget(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellEditableOverrider)

	iface.RemoveWidget()
}

//export _gotk4_gtk4_CellEditableIface_start_editing
func _gotk4_gtk4_CellEditableIface_start_editing(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(CellEditableOverrider)

	var _event gdk.Eventer // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(gdk.Eventer)
				return ok
			})
			rv, ok := casted.(gdk.Eventer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Eventer")
			}
			_event = rv
		}
	}

	iface.StartEditing(_event)
}

func wrapCellEditable(obj *coreglib.Object) *CellEditable {
	return &CellEditable{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalCellEditable(p uintptr) (interface{}, error) {
	return wrapCellEditable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_CellEditable_ConnectEditingDone
func _gotk4_gtk4_CellEditable_ConnectEditingDone(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectEditingDone: this signal is a sign for the cell renderer to update its
// value from the cell_editable.
//
// Implementations of CellEditable are responsible for emitting this signal when
// they are done editing, e.g. Entry emits this signal when the user presses
// Enter. Typical things to do in a handler for ::editing-done are to capture
// the edited value, disconnect the cell_editable from signals on the
// CellRenderer, etc.
//
// gtk_cell_editable_editing_done() is a convenience method for emitting
// CellEditable::editing-done.
func (cellEditable *CellEditable) ConnectEditingDone(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cellEditable, "editing-done", false, unsafe.Pointer(C._gotk4_gtk4_CellEditable_ConnectEditingDone), f)
}

//export _gotk4_gtk4_CellEditable_ConnectRemoveWidget
func _gotk4_gtk4_CellEditable_ConnectRemoveWidget(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectRemoveWidget: this signal is meant to indicate that the cell is
// finished editing, and the cell_editable widget is being removed and may
// subsequently be destroyed.
//
// Implementations of CellEditable are responsible for emitting this signal when
// they are done editing. It must be emitted after the
// CellEditable::editing-done signal, to give the cell renderer a chance to
// update the cell's value before the widget is removed.
//
// gtk_cell_editable_remove_widget() is a convenience method for emitting
// CellEditable::remove-widget.
func (cellEditable *CellEditable) ConnectRemoveWidget(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(cellEditable, "remove-widget", false, unsafe.Pointer(C._gotk4_gtk4_CellEditable_ConnectRemoveWidget), f)
}

// EditingDone emits the CellEditable::editing-done signal.
func (cellEditable *CellEditable) EditingDone() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellEditable).Native()))

	runtime.KeepAlive(cellEditable)
}

// RemoveWidget emits the CellEditable::remove-widget signal.
func (cellEditable *CellEditable) RemoveWidget() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellEditable).Native()))

	runtime.KeepAlive(cellEditable)
}

// StartEditing begins editing on a cell_editable.
//
// The CellRenderer for the cell creates and returns a CellEditable from
// gtk_cell_renderer_start_editing(), configured for the CellRenderer type.
//
// gtk_cell_editable_start_editing() can then set up cell_editable suitably for
// editing a cell, e.g. making the Esc key emit CellEditable::editing-done.
//
// Note that the cell_editable is created on-demand for the current edit; its
// lifetime is temporary and does not persist across other edits and/or cells.
//
// The function takes the following parameters:
//
//    - event (optional) that began the editing process, or NULL if editing was
//      initiated programmatically.
//
func (cellEditable *CellEditable) StartEditing(event gdk.Eventer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cellEditable).Native()))
	if event != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(event).Native()))
	}

	runtime.KeepAlive(cellEditable)
	runtime.KeepAlive(event)
}
