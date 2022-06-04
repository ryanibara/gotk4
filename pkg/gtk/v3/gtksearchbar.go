// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtksearchbar.go.
var GTypeSearchBar = coreglib.Type(C.gtk_search_bar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeSearchBar, F: marshalSearchBar},
	})
}

// SearchBarOverrider contains methods that are overridable.
type SearchBarOverrider interface {
}

// SearchBar is a container made to have a search entry (possibly with
// additional connex widgets, such as drop-down menus, or buttons) built-in. The
// search bar would appear when a search is started through typing on the
// keyboard, or the application’s search mode is toggled on.
//
// For keyboard presses to start a search, events will need to be forwarded from
// the top-level window that contains the search bar. See
// gtk_search_bar_handle_event() for example code. Common shortcuts such as
// Ctrl+F should be handled as an application action, or through the menu items.
//
// You will also need to tell the search bar about which entry you are using as
// your search entry using gtk_search_bar_connect_entry(). The following example
// shows you how to create a more complex search entry.
//
//
// CSS nodes
//
// GtkSearchBar has a single CSS node with name searchbar.
//
//
// Creating a search bar
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/blob/gtk-3-24/examples/search-bar.c).
type SearchBar struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*SearchBar)(nil)
)

func classInitSearchBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSearchBar(obj *coreglib.Object) *SearchBar {
	return &SearchBar{
		Bin: Bin{
			Container: Container{
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
		},
	}
}

func marshalSearchBar(p uintptr) (interface{}, error) {
	return wrapSearchBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSearchBar creates a SearchBar. You will need to tell it about which widget
// is going to be your text entry using gtk_search_bar_connect_entry().
//
// The function returns the following values:
//
//    - searchBar: new SearchBar.
//
func NewSearchBar() *SearchBar {
	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("new_SearchBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _searchBar *SearchBar // out

	_searchBar = wrapSearchBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _searchBar
}

// ConnectEntry connects the Entry widget passed as the one to be used in this
// search bar. The entry should be a descendant of the search bar. This is only
// required if the entry isn’t the direct child of the search bar (as in our
// main example).
//
// The function takes the following parameters:
//
//    - entry: Entry.
//
func (bar *SearchBar) ConnectEntry(entry *Entry) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("connect_entry", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(entry)
}

// SearchMode returns whether the search mode is on or off.
//
// The function returns the following values:
//
//    - ok: whether search mode is toggled on.
//
func (bar *SearchBar) SearchMode() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("get_search_mode", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether the close button is shown.
//
// The function returns the following values:
//
//    - ok: whether the close button is shown.
//
func (bar *SearchBar) ShowCloseButton() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("get_show_close_button", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HandleEvent: this function should be called when the top-level window which
// contains the search bar received a key event.
//
// If the key event is handled by the search bar, the bar will be shown, the
// entry populated with the entered text and GDK_EVENT_STOP will be returned.
// The caller should ensure that events are not propagated further.
//
// If no entry has been connected to the search bar, using
// gtk_search_bar_connect_entry(), this function will return immediately with a
// warning.
//
// Showing the search bar on key presses
//
//    static gboolean
//    on_key_press_event (GtkWidget *widget,
//                        GdkEvent  *event,
//                        gpointer   user_data)
//    {
//      GtkSearchBar *bar = GTK_SEARCH_BAR (user_data);
//      return gtk_search_bar_handle_event (bar, event);
//    }
//
//    static void
//    create_toplevel (void)
//    {
//      GtkWidget *window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      GtkWindow *search_bar = gtk_search_bar_new ();
//
//     // Add more widgets to the window...
//
//      g_signal_connect (window,
//                       "key-press-event",
//                        G_CALLBACK (on_key_press_event),
//                        search_bar);
//    }.
//
// The function takes the following parameters:
//
//    - event containing key press events.
//
// The function returns the following values:
//
//    - ok: GDK_EVENT_STOP if the key press event resulted in text being entered
//      in the search entry (and revealing the search bar if necessary),
//      GDK_EVENT_PROPAGATE otherwise.
//
func (bar *SearchBar) HandleEvent(event *gdk.Event) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(event)))

	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("handle_event", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)
	runtime.KeepAlive(event)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetSearchMode switches the search mode on or off.
//
// The function takes the following parameters:
//
//    - searchMode: new state of the search mode.
//
func (bar *SearchBar) SetSearchMode(searchMode bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if searchMode {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("set_search_mode", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(searchMode)
}

// SetShowCloseButton shows or hides the close button. Applications that already
// have a “search” toggle button should not show a close button in their search
// bar, as it duplicates the role of the toggle button.
//
// The function takes the following parameters:
//
//    - visible: whether the close button will be shown or not.
//
func (bar *SearchBar) SetShowCloseButton(visible bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if visible {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("set_show_close_button", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(visible)
}
