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
import "C"

// glib.Type values for gtksorter.go.
var (
	GTypeSorterChange = coreglib.Type(C.gtk_sorter_change_get_type())
	GTypeSorterOrder  = coreglib.Type(C.gtk_sorter_order_get_type())
	GTypeSorter       = coreglib.Type(C.gtk_sorter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSorterChange, F: marshalSorterChange},
		{T: GTypeSorterOrder, F: marshalSorterOrder},
		{T: GTypeSorter, F: marshalSorter},
	})
}

// SorterChange describes changes in a sorter in more detail and allows users to
// optimize resorting.
type SorterChange C.gint

const (
	// SorterChangeDifferent: sorter change cannot be described by any of the
	// other enumeration values.
	SorterChangeDifferent SorterChange = iota
	// SorterChangeInverted: sort order was inverted. Comparisons that returned
	// GTK_ORDERING_SMALLER now return GTK_ORDERING_LARGER and vice versa. Other
	// comparisons return the same values as before.
	SorterChangeInverted
	// SorterChangeLessStrict: sorter is less strict: Comparisons may now return
	// GTK_ORDERING_EQUAL that did not do so before.
	SorterChangeLessStrict
	// SorterChangeMoreStrict: sorter is more strict: Comparisons that did
	// return GTK_ORDERING_EQUAL may not do so anymore.
	SorterChangeMoreStrict
)

func marshalSorterChange(p uintptr) (interface{}, error) {
	return SorterChange(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SorterChange.
func (s SorterChange) String() string {
	switch s {
	case SorterChangeDifferent:
		return "Different"
	case SorterChangeInverted:
		return "Inverted"
	case SorterChangeLessStrict:
		return "LessStrict"
	case SorterChangeMoreStrict:
		return "MoreStrict"
	default:
		return fmt.Sprintf("SorterChange(%d)", s)
	}
}

// SorterOrder describes the type of order that a GtkSorter may produce.
type SorterOrder C.gint

const (
	// SorterOrderPartial: partial order. Any Ordering is possible.
	SorterOrderPartial SorterOrder = iota
	// SorterOrderNone: no order, all elements are considered equal.
	// gtk_sorter_compare() will only return GTK_ORDERING_EQUAL.
	SorterOrderNone
	// SorterOrderTotal: total order. gtk_sorter_compare() will only return
	// GTK_ORDERING_EQUAL if an item is compared with itself. Two different
	// items will never cause this value to be returned.
	SorterOrderTotal
)

func marshalSorterOrder(p uintptr) (interface{}, error) {
	return SorterOrder(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SorterOrder.
func (s SorterOrder) String() string {
	switch s {
	case SorterOrderPartial:
		return "Partial"
	case SorterOrderNone:
		return "None"
	case SorterOrderTotal:
		return "Total"
	default:
		return fmt.Sprintf("SorterOrder(%d)", s)
	}
}

// SorterOverrider contains methods that are overridable.
type SorterOverrider interface {
}

// Sorter: GtkSorter is an object to describe sorting criteria.
//
//
// Its primary user is gtk.SortListModel
//
// The model will use a sorter to determine the order in which its items should
// appear by calling gtk.Sorter.Compare() for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime. In that
// case, they will emit the gtk.Sorter::changed signal to notify that the sort
// order is no longer valid and should be updated by calling
// gtk_sorter_compare() again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. gtk.ColumnView has built-in support for sorting lists via the
// gtk.ColumnViewColumn:sorter property, where the user can change the sorting
// by clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// GtkSorter and provide one's own sorter.
type Sorter struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Sorter)(nil)
)

func classInitSorterer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSorter(obj *coreglib.Object) *Sorter {
	return &Sorter{
		Object: obj,
	}
}

func marshalSorter(p uintptr) (interface{}, error) {
	return wrapSorter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
