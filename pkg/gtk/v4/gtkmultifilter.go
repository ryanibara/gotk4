// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_any_filter_get_type()), F: marshalAnyFilterer},
		{T: externglib.Type(C.gtk_every_filter_get_type()), F: marshalEveryFilterer},
		{T: externglib.Type(C.gtk_multi_filter_get_type()), F: marshalMultiFilterer},
	})
}

// AnyFilter: GtkAnyFilter matches an item when at least one of its filters
// matches.
//
// To add filters to a GtkAnyFilter, use gtk.MultiFilter.Append().
type AnyFilter struct {
	MultiFilter
}

func wrapAnyFilter(obj *externglib.Object) *AnyFilter {
	return &AnyFilter{
		MultiFilter: MultiFilter{
			Filter: Filter{
				Object: obj,
			},
			ListModel: gio.ListModel{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalAnyFilterer(p uintptr) (interface{}, error) {
	return wrapAnyFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAnyFilter creates a new empty "any" filter.
//
// Use gtk.MultiFilter.Append() to add filters to it.
//
// This filter matches an item if any of the filters added to it matches the
// item. In particular, this means that if no filter has been added to it, the
// filter matches no item.
func NewAnyFilter() *AnyFilter {
	var _cret *C.GtkAnyFilter // in

	_cret = C.gtk_any_filter_new()

	var _anyFilter *AnyFilter // out

	_anyFilter = wrapAnyFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _anyFilter
}

func (*AnyFilter) privateAnyFilter() {}

// EveryFilter: GtkEveryFilter matches an item when each of its filters matches.
//
// To add filters to a GtkEveryFilter, use gtk.MultiFilter.Append().
type EveryFilter struct {
	MultiFilter
}

func wrapEveryFilter(obj *externglib.Object) *EveryFilter {
	return &EveryFilter{
		MultiFilter: MultiFilter{
			Filter: Filter{
				Object: obj,
			},
			ListModel: gio.ListModel{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalEveryFilterer(p uintptr) (interface{}, error) {
	return wrapEveryFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEveryFilter creates a new empty "every" filter.
//
// Use gtk.MultiFilter.Append() to add filters to it.
//
// This filter matches an item if each of the filters added to it matches the
// item. In particular, this means that if no filter has been added to it, the
// filter matches every item.
func NewEveryFilter() *EveryFilter {
	var _cret *C.GtkEveryFilter // in

	_cret = C.gtk_every_filter_new()

	var _everyFilter *EveryFilter // out

	_everyFilter = wrapEveryFilter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _everyFilter
}

func (*EveryFilter) privateEveryFilter() {}

// MultiFilter: GtkMultiFilter is the base class for filters that combine
// multiple filters.
type MultiFilter struct {
	Filter

	gio.ListModel
	Buildable
	*externglib.Object
}

// MultiFilterer describes MultiFilter's abstract methods.
type MultiFilterer interface {
	externglib.Objector

	// Append adds a filter to self to use for matching.
	Append(filter *Filter)
	// Remove removes the filter at the given position from the list of filters
	// used by self.
	Remove(position uint)
}

var _ MultiFilterer = (*MultiFilter)(nil)

func wrapMultiFilter(obj *externglib.Object) *MultiFilter {
	return &MultiFilter{
		Filter: Filter{
			Object: obj,
		},
		ListModel: gio.ListModel{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalMultiFilterer(p uintptr) (interface{}, error) {
	return wrapMultiFilter(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Append adds a filter to self to use for matching.
//
// The function takes the following parameters:
//
//    - filter: new filter to use.
//
func (self *MultiFilter) Append(filter *Filter) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 *C.GtkFilter      // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkFilter)(unsafe.Pointer(filter.Native()))
	C.g_object_ref(C.gpointer(filter.Native()))

	C.gtk_multi_filter_append(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(filter)
}

// Remove removes the filter at the given position from the list of filters used
// by self.
//
// If position is larger than the number of filters, nothing happens and the
// function returns.
//
// The function takes the following parameters:
//
//    - position of filter to remove.
//
func (self *MultiFilter) Remove(position uint) {
	var _arg0 *C.GtkMultiFilter // out
	var _arg1 C.guint           // out

	_arg0 = (*C.GtkMultiFilter)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(position)

	C.gtk_multi_filter_remove(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(position)
}
