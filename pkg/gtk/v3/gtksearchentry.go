// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_search_entry_get_type()), F: marshalSearchEntrier},
	})
}

// SearchEntryOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type SearchEntryOverrider interface {
	NextMatch()

	PreviousMatch()

	SearchChanged()

	StopSearch()
}

// SearchEntrier describes SearchEntry's methods.
type SearchEntrier interface {
	privateSearchEntry()
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
	Entry
}

var (
	_ SearchEntrier   = (*SearchEntry)(nil)
	_ gextras.Nativer = (*SearchEntry)(nil)
)

func wrapSearchEntry(obj *externglib.Object) SearchEntrier {
	return &SearchEntry{
		Entry: Entry{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
			CellEditable: CellEditable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
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

func marshalSearchEntrier(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSearchEntry(obj), nil
}

// NewSearchEntry creates a SearchEntry, with a find icon when the search field
// is empty, and a clear icon when it isn't.
func NewSearchEntry() *SearchEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_search_entry_new()

	var _searchEntry *SearchEntry // out

	_searchEntry = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*SearchEntry)

	return _searchEntry
}

func (*SearchEntry) privateSearchEntry() {}
