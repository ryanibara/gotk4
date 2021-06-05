// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_multi_sorter_get_type()), F: marshalMultiSorter},
	})
}

// MultiSorter: gtkMultiSorter combines multiple sorters by trying them in turn.
// If the first sorter compares two items as equal, the second is tried next,
// and so on.
type MultiSorter interface {
	Sorter
	gio.ListModel
	Buildable

	// Append: add @sorter to @self to use for sorting at the end. @self will
	// consult all existing sorters before it will sort with the given @sorter.
	Append(sorter Sorter)
	// Remove removes the sorter at the given @position from the list of sorter
	// used by @self.
	//
	// If @position is larger than the number of sorters, nothing happens.
	Remove(position uint)
}

// multiSorter implements the MultiSorter interface.
type multiSorter struct {
	Sorter
	gio.ListModel
	Buildable
}

var _ MultiSorter = (*multiSorter)(nil)

// WrapMultiSorter wraps a GObject to the right type. It is
// primarily used internally.
func WrapMultiSorter(obj *externglib.Object) MultiSorter {
	return MultiSorter{
		Sorter:        WrapSorter(obj),
		gio.ListModel: gio.WrapListModel(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalMultiSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiSorter(obj), nil
}

// NewMultiSorter constructs a class MultiSorter.
func NewMultiSorter() MultiSorter {
	var cret C.GtkMultiSorter
	var ret1 MultiSorter

	cret = C.gtk_multi_sorter_new()

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MultiSorter)

	return ret1
}

// Append: add @sorter to @self to use for sorting at the end. @self will
// consult all existing sorters before it will sort with the given @sorter.
func (s multiSorter) Append(sorter Sorter) {
	var arg0 *C.GtkMultiSorter
	var arg1 *C.GtkSorter

	arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_multi_sorter_append(arg0, sorter)
}

// Remove removes the sorter at the given @position from the list of sorter
// used by @self.
//
// If @position is larger than the number of sorters, nothing happens.
func (s multiSorter) Remove(position uint) {
	var arg0 *C.GtkMultiSorter
	var arg1 C.guint

	arg0 = (*C.GtkMultiSorter)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)

	C.gtk_multi_sorter_remove(arg0, position)
}
