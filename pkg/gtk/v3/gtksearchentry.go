// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_SearchEntryClass_next_match(void*);
// extern void _gotk4_gtk3_SearchEntryClass_previous_match(void*);
// extern void _gotk4_gtk3_SearchEntryClass_search_changed(void*);
// extern void _gotk4_gtk3_SearchEntryClass_stop_search(void*);
// extern void _gotk4_gtk3_SearchEntry_ConnectNextMatch(gpointer, guintptr);
// extern void _gotk4_gtk3_SearchEntry_ConnectPreviousMatch(gpointer, guintptr);
// extern void _gotk4_gtk3_SearchEntry_ConnectSearchChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_SearchEntry_ConnectStopSearch(gpointer, guintptr);
import "C"

// glib.Type values for gtksearchentry.go.
var GTypeSearchEntry = coreglib.Type(C.gtk_search_entry_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSearchEntry, F: marshalSearchEntry},
	})
}

// SearchEntryOverrider contains methods that are overridable.
type SearchEntryOverrider interface {
	NextMatch()
	PreviousMatch()
	SearchChanged()
	StopSearch()
}

// SearchEntry is a subclass of Entry that has been tailored for use as a search
// entry.
//
// It will show an inactive symbolic “find” icon when the search entry is empty,
// and a symbolic “clear” icon when there is text. Clicking on the “clear” icon
// will empty the search entry.
//
// Note that the search/clear icon is shown using a secondary icon, and thus
// does not work if you are using the secondary icon position for some other
// purpose.
//
// To make filtering appear more reactive, it is a good idea to not react to
// every change in the entry text immediately, but only after a short delay. To
// support this, SearchEntry emits the SearchEntry::search-changed signal which
// can be used instead of the Editable::changed signal.
//
// The SearchEntry::previous-match, SearchEntry::next-match and
// SearchEntry::stop-search signals can be used to implement moving between
// search results and ending the search.
//
// Often, GtkSearchEntry will be fed events by means of being placed inside a
// SearchBar. If that is not the case, you can use
// gtk_search_entry_handle_event() to pass events.
type SearchEntry struct {
	_ [0]func() // equal guard
	Entry
}

var (
	_ Widgetter         = (*SearchEntry)(nil)
	_ coreglib.Objector = (*SearchEntry)(nil)
)

func classInitSearchEntrier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkSearchEntryClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkSearchEntryClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ NextMatch() }); ok {
		pclass.next_match = (*[0]byte)(C._gotk4_gtk3_SearchEntryClass_next_match)
	}

	if _, ok := goval.(interface{ PreviousMatch() }); ok {
		pclass.previous_match = (*[0]byte)(C._gotk4_gtk3_SearchEntryClass_previous_match)
	}

	if _, ok := goval.(interface{ SearchChanged() }); ok {
		pclass.search_changed = (*[0]byte)(C._gotk4_gtk3_SearchEntryClass_search_changed)
	}

	if _, ok := goval.(interface{ StopSearch() }); ok {
		pclass.stop_search = (*[0]byte)(C._gotk4_gtk3_SearchEntryClass_stop_search)
	}
}

//export _gotk4_gtk3_SearchEntryClass_next_match
func _gotk4_gtk3_SearchEntryClass_next_match(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NextMatch() })

	iface.NextMatch()
}

//export _gotk4_gtk3_SearchEntryClass_previous_match
func _gotk4_gtk3_SearchEntryClass_previous_match(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ PreviousMatch() })

	iface.PreviousMatch()
}

//export _gotk4_gtk3_SearchEntryClass_search_changed
func _gotk4_gtk3_SearchEntryClass_search_changed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ SearchChanged() })

	iface.SearchChanged()
}

//export _gotk4_gtk3_SearchEntryClass_stop_search
func _gotk4_gtk3_SearchEntryClass_stop_search(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ StopSearch() })

	iface.StopSearch()
}

func wrapSearchEntry(obj *coreglib.Object) *SearchEntry {
	return &SearchEntry{
		Entry: Entry{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			Object: obj,
			CellEditable: CellEditable{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Editable: Editable{
				Object: obj,
			},
		},
	}
}

func marshalSearchEntry(p uintptr) (interface{}, error) {
	return wrapSearchEntry(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_SearchEntry_ConnectNextMatch
func _gotk4_gtk3_SearchEntry_ConnectNextMatch(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectNextMatch signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user initiates a move to the next match for the current
// search string.
//
// Applications should connect to it, to implement moving between matches.
//
// The default bindings for this signal is Ctrl-g.
func (entry *SearchEntry) ConnectNextMatch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "next-match", false, unsafe.Pointer(C._gotk4_gtk3_SearchEntry_ConnectNextMatch), f)
}

//export _gotk4_gtk3_SearchEntry_ConnectPreviousMatch
func _gotk4_gtk3_SearchEntry_ConnectPreviousMatch(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectPreviousMatch signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user initiates a move to the previous match for the
// current search string.
//
// Applications should connect to it, to implement moving between matches.
//
// The default bindings for this signal is Ctrl-Shift-g.
func (entry *SearchEntry) ConnectPreviousMatch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "previous-match", false, unsafe.Pointer(C._gotk4_gtk3_SearchEntry_ConnectPreviousMatch), f)
}

//export _gotk4_gtk3_SearchEntry_ConnectSearchChanged
func _gotk4_gtk3_SearchEntry_ConnectSearchChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectSearchChanged signal is emitted with a short delay of 150 milliseconds
// after the last change to the entry text.
func (entry *SearchEntry) ConnectSearchChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "search-changed", false, unsafe.Pointer(C._gotk4_gtk3_SearchEntry_ConnectSearchChanged), f)
}

//export _gotk4_gtk3_SearchEntry_ConnectStopSearch
func _gotk4_gtk3_SearchEntry_ConnectStopSearch(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectStopSearch signal is a [keybinding signal][GtkBindingSignal] which
// gets emitted when the user stops a search via keyboard input.
//
// Applications should connect to it, to implement hiding the search entry in
// this case.
//
// The default bindings for this signal is Escape.
func (entry *SearchEntry) ConnectStopSearch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(entry, "stop-search", false, unsafe.Pointer(C._gotk4_gtk3_SearchEntry_ConnectStopSearch), f)
}

// NewSearchEntry creates a SearchEntry, with a find icon when the search field
// is empty, and a clear icon when it isn't.
//
// The function returns the following values:
//
//    - searchEntry: new SearchEntry.
//
func NewSearchEntry() *SearchEntry {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "SearchEntry").InvokeMethod("new_SearchEntry", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _searchEntry *SearchEntry // out

	_searchEntry = wrapSearchEntry(coreglib.Take(unsafe.Pointer(_cret)))

	return _searchEntry
}

// HandleEvent: this function should be called when the top-level window which
// contains the search entry received a key event. If the entry is part of a
// SearchBar, it is preferable to call gtk_search_bar_handle_event() instead,
// which will reveal the entry in addition to passing the event to this
// function.
//
// If the key event is handled by the search entry and starts or continues a
// search, GDK_EVENT_STOP will be returned. The caller should ensure that the
// entry is shown in this case, and not propagate the event further.
//
// The function takes the following parameters:
//
//    - event: key event.
//
// The function returns the following values:
//
//    - ok: GDK_EVENT_STOP if the key press event resulted in a search beginning
//      or continuing, GDK_EVENT_PROPAGATE otherwise.
//
func (entry *SearchEntry) HandleEvent(event *gdk.Event) bool {
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(event)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "SearchEntry").InvokeMethod("handle_event", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(entry)
	runtime.KeepAlive(event)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
