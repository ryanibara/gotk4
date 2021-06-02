// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_tab_array_get_type()), F: marshalTabArray},
	})
}

// TabArray: a `PangoTabArray` contains an array of tab stops.
//
// `PangoTabArray` can be used to set tab stops in a `PangoLayout`. Each tab
// stop has an alignment and a position.
type TabArray struct {
	native C.PangoTabArray
}

// WrapTabArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTabArray(ptr unsafe.Pointer) *TabArray {
	if ptr == nil {
		return nil
	}

	return (*TabArray)(ptr)
}

func marshalTabArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTabArray(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TabArray) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// NewTabArray constructs a struct TabArray.
func NewTabArray(initialSize int, positionsInPixels bool) *TabArray {
	var arg1 C.gint
	var arg2 C.gboolean

	arg1 = C.gint(initialSize)
	if positionsInPixels {
		arg2 = C.TRUE
	}

	ret := C.pango_tab_array_new(arg1, arg2)

	var ret0 *TabArray

	{
		ret0 = WrapTabArray(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *TabArray) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Copy copies a `PangoTabArray`.
func (s *TabArray) Copy() *TabArray {
	var arg0 *C.PangoTabArray

	arg0 = (*C.PangoTabArray)(s.Native())

	ret := C.pango_tab_array_copy(arg0)

	var ret0 *TabArray

	{
		ret0 = WrapTabArray(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *TabArray) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Free frees a tab array and associated resources.
func (t *TabArray) Free() {
	var arg0 *C.PangoTabArray

	arg0 = (*C.PangoTabArray)(t.Native())

	C.pango_tab_array_free(arg0)
}

// PositionsInPixels returns true if the tab positions are in pixels, false if
// they are in Pango units.
func (t *TabArray) PositionsInPixels() bool {
	var arg0 *C.PangoTabArray

	arg0 = (*C.PangoTabArray)(t.Native())

	ret := C.pango_tab_array_get_positions_in_pixels(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != C.false

	return ret0
}

// Size gets the number of tab stops in @tab_array.
func (t *TabArray) Size() int {
	var arg0 *C.PangoTabArray

	arg0 = (*C.PangoTabArray)(t.Native())

	ret := C.pango_tab_array_get_size(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// Tab gets the alignment and position of a tab stop.
func (t *TabArray) Tab(tabIndex int) (alignment TabAlign, location int) {
	var arg0 *C.PangoTabArray
	var arg1 C.gint
	var arg2 *C.PangoTabAlign // out
	var arg3 *C.gint          // out

	arg0 = (*C.PangoTabArray)(t.Native())
	arg1 = C.gint(tabIndex)

	C.pango_tab_array_get_tab(arg0, arg1, &arg2, &arg3)

	var ret0 *TabAlign
	var ret1 int

	ret0 = (*TabAlign)(arg2)

	ret1 = int(arg3)

	return ret0, ret1
}

// Tabs: if non-nil, @alignments and @locations are filled with allocated
// arrays.
//
// The arrays are of length [method@Pango.TabArray.get_size]. You must free the
// returned array.
func (t *TabArray) Tabs() (alignments *TabAlign, locations []int) {
	var arg0 *C.PangoTabArray
	var arg1 **C.PangoTabAlign // out
	var arg2 **C.gint          // out

	arg0 = (*C.PangoTabArray)(t.Native())

	C.pango_tab_array_get_tabs(arg0, &arg1, &arg2)

	var ret0 **TabAlign
	var ret1 []int

	ret0 = (**TabAlign)(arg1)

	return ret0, ret1
}

// Resize resizes a tab array.
//
// You must subsequently initialize any tabs that were added as a result of
// growing the array.
func (t *TabArray) Resize(newSize int) {
	var arg0 *C.PangoTabArray
	var arg1 C.gint

	arg0 = (*C.PangoTabArray)(t.Native())
	arg1 = C.gint(newSize)

	C.pango_tab_array_resize(arg0, arg1)
}

// SetTab sets the alignment and location of a tab stop.
//
// @alignment must always be PANGO_TAB_LEFT in the current implementation.
func (t *TabArray) SetTab(tabIndex int, alignment TabAlign, location int) {
	var arg0 *C.PangoTabArray
	var arg1 C.gint
	var arg2 C.PangoTabAlign
	var arg3 C.gint

	arg0 = (*C.PangoTabArray)(t.Native())
	arg1 = C.gint(tabIndex)
	arg2 = (C.PangoTabAlign)(alignment)
	arg3 = C.gint(location)

	C.pango_tab_array_set_tab(arg0, arg1, arg2, arg3)
}
