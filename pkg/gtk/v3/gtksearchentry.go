// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// void _gotk4_gtk3_SearchEntry_virtual_next_match(void* fnptr, GtkSearchEntry* arg0) {
//   ((void (*)(GtkSearchEntry*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_SearchEntry_virtual_previous_match(void* fnptr, GtkSearchEntry* arg0) {
//   ((void (*)(GtkSearchEntry*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_SearchEntry_virtual_search_changed(void* fnptr, GtkSearchEntry* arg0) {
//   ((void (*)(GtkSearchEntry*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_SearchEntry_virtual_stop_search(void* fnptr, GtkSearchEntry* arg0) {
//   ((void (*)(GtkSearchEntry*))(fnptr))(arg0);
// };
import "C"

func (entry *SearchEntry) nextMatch() {
	gclass := (*C.GtkSearchEntryClass)(coreglib.PeekParentClass(entry))
	fnarg := gclass.next_match

	var _arg0 *C.GtkSearchEntry // out

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	C._gotk4_gtk3_SearchEntry_virtual_next_match(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(entry)
}

func (entry *SearchEntry) previousMatch() {
	gclass := (*C.GtkSearchEntryClass)(coreglib.PeekParentClass(entry))
	fnarg := gclass.previous_match

	var _arg0 *C.GtkSearchEntry // out

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	C._gotk4_gtk3_SearchEntry_virtual_previous_match(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(entry)
}

func (entry *SearchEntry) searchChanged() {
	gclass := (*C.GtkSearchEntryClass)(coreglib.PeekParentClass(entry))
	fnarg := gclass.search_changed

	var _arg0 *C.GtkSearchEntry // out

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	C._gotk4_gtk3_SearchEntry_virtual_search_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(entry)
}

func (entry *SearchEntry) stopSearch() {
	gclass := (*C.GtkSearchEntryClass)(coreglib.PeekParentClass(entry))
	fnarg := gclass.stop_search

	var _arg0 *C.GtkSearchEntry // out

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	C._gotk4_gtk3_SearchEntry_virtual_stop_search(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(entry)
}

// SearchEntryClass: instance of this type is always passed by reference.
type SearchEntryClass struct {
	*searchEntryClass
}

// searchEntryClass is the struct that's finalized.
type searchEntryClass struct {
	native *C.GtkSearchEntryClass
}

func (s *SearchEntryClass) ParentClass() *EntryClass {
	valptr := &s.native.parent_class
	var _v *EntryClass // out
	_v = (*EntryClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
