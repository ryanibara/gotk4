// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_search_entry_get_type()), F: marshalSearchEntrier},
	})
}

// SearchEntrier describes SearchEntry's methods.
type SearchEntrier interface {
	// KeyCaptureWidget gets the widget that @entry is capturing key events
	// from.
	KeyCaptureWidget() *Widget
	// SetKeyCaptureWidget sets @widget as the widget that @entry will capture
	// key events from.
	SetKeyCaptureWidget(widget Widgeter)
}

// SearchEntry: `GtkSearchEntry` is an entry widget that has been tailored for
// use as a search entry.
//
// The main API for interacting with a `GtkSearchEntry` as entry is the
// `GtkEditable` interface.
//
// !An example GtkSearchEntry (search-entry.png)
//
// It will show an inactive symbolic “find” icon when the search entry is empty,
// and a symbolic “clear” icon when there is text. Clicking on the “clear” icon
// will empty the search entry.
//
// To make filtering appear more reactive, it is a good idea to not react to
// every change in the entry text immediately, but only after a short delay. To
// support this, `GtkSearchEntry` emits the
// [signal@Gtk.SearchEntry::search-changed] signal which can be used instead of
// the [signal@Gtk.Editable::changed] signal.
//
// The [signal@Gtk.SearchEntry::previous-match],
// [signal@Gtk.SearchEntry::next-match] and
// [signal@Gtk.SearchEntry::stop-search] signals can be used to implement moving
// between search results and ending the search.
//
// Often, `GtkSearchEntry` will be fed events by means of being placed inside a
// [class@Gtk.SearchBar]. If that is not the case, you can use
// [method@Gtk.SearchEntry.set_key_capture_widget] to let it capture key input
// from another widget.
//
// `GtkSearchEntry` provides only minimal API and should be used with the
// [iface@Gtk.Editable] API.
//
//
// CSS Nodes
//
// “` entry.search ╰── text “`
//
// `GtkSearchEntry` has a single CSS node with name entry that carries a
// `.search` style class, and the text node is a child of that.
//
//
// Accessibility
//
// `GtkSearchEntry` uses the GTK_ACCESSIBLE_ROLE_SEARCH_BOX role.
type SearchEntry struct {
	Widget

	Editable
}

var (
	_ SearchEntrier   = (*SearchEntry)(nil)
	_ gextras.Nativer = (*SearchEntry)(nil)
)

func wrapSearchEntry(obj *externglib.Object) *SearchEntry {
	return &SearchEntry{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
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
		},
	}
}

func marshalSearchEntrier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSearchEntry(obj), nil
}

// NewSearchEntry creates a `GtkSearchEntry`.
func NewSearchEntry() *SearchEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_search_entry_new()

	var _searchEntry *SearchEntry // out

	_searchEntry = wrapSearchEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _searchEntry
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *SearchEntry) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// KeyCaptureWidget gets the widget that @entry is capturing key events from.
func (entry *SearchEntry) KeyCaptureWidget() *Widget {
	var _arg0 *C.GtkSearchEntry // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_search_entry_get_key_capture_widget(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// SetKeyCaptureWidget sets @widget as the widget that @entry will capture key
// events from.
//
// Key events are consumed by the search entry to start or continue a search.
//
// If the entry is part of a `GtkSearchBar`, it is preferable to call
// [method@Gtk.SearchBar.set_key_capture_widget] instead, which will reveal the
// entry in addition to triggering the search entry.
//
// Note that despite the name of this function, the events are only 'captured'
// in the bubble phase, which means that editable child widgets of @widget will
// receive text input before it gets captured. If that is not desired, you can
// capture and forward the events yourself with
// [method@Gtk.EventControllerKey.forward].
func (entry *SearchEntry) SetKeyCaptureWidget(widget Widgeter) {
	var _arg0 *C.GtkSearchEntry // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.GtkSearchEntry)(unsafe.Pointer(entry.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((widget).(gextras.Nativer).Native()))

	C.gtk_search_entry_set_key_capture_widget(_arg0, _arg1)
}
