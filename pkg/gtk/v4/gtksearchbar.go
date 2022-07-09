// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeSearchBar returns the GType for the type SearchBar.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeSearchBar() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "SearchBar").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalSearchBar)
	return gtype
}

// SearchBar: GtkSearchBar is a container made to have a search entry.
//
// !An example GtkSearchBar (search-bar.png)
//
// It can also contain additional widgets, such as drop-down menus, or buttons.
// The search bar would appear when a search is started through typing on the
// keyboard, or the application’s search mode is toggled on.
//
// For keyboard presses to start a search, the search bar must be told of a
// widget to capture key events from through
// gtk.SearchBar.SetKeyCaptureWidget(). This widget will typically be the
// top-level window, or a parent container of the search bar. Common shortcuts
// such as Ctrl+F should be handled as an application action, or through the
// menu items.
//
// You will also need to tell the search bar about which entry you are using as
// your search entry using gtk.SearchBar.ConnectEntry().
//
//
// Creating a search bar
//
// The following example shows you how to create a more complex search entry.
//
// A simple example
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/examples/search-bar.c)
//
// CSS nodes
//
//    searchbar
//    ╰── revealer
//        ╰── box
//             ├── [child]
//             ╰── [button.close]
//
//
// GtkSearchBar has a main CSS node with name searchbar. It has a child node
// with name revealer that contains a node with name box. The box node contains
// both the CSS node of the child widget as well as an optional button node
// which gets the .close style class applied.
//
//
// Accessibility
//
// GtkSearchBar uses the GTK_ACCESSIBLE_ROLE_SEARCH role.
type SearchBar struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*SearchBar)(nil)
)

func wrapSearchBar(obj *coreglib.Object) *SearchBar {
	return &SearchBar{
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

func marshalSearchBar(p uintptr) (interface{}, error) {
	return wrapSearchBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSearchBar creates a GtkSearchBar.
//
// You will need to tell it about which widget is going to be your text entry
// using gtk.SearchBar.ConnectEntry().
//
// The function returns the following values:
//
//    - searchBar: new GtkSearchBar.
//
func NewSearchBar() *SearchBar {
	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("new_SearchBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _searchBar *SearchBar // out

	_searchBar = wrapSearchBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _searchBar
}

// ConnectEntry connects the `GtkEditable widget passed as the one to be used in
// this search bar.
//
// The entry should be a descendant of the search bar. Calling this function
// manually is only required if the entry isn’t the direct child of the search
// bar (as in our main example).
//
// The function takes the following parameters:
//
//    - entry: GtkEditable.
//
func (bar *SearchBar) ConnectEntry(entry Editabler) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(entry).Native()))

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("connect_entry", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(entry)
}

// Child gets the child widget of bar.
//
// The function returns the following values:
//
//    - widget (optional): child widget of bar.
//
func (bar *SearchBar) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// KeyCaptureWidget gets the widget that bar is capturing key events from.
//
// The function returns the following values:
//
//    - widget: key capture widget.
//
func (bar *SearchBar) KeyCaptureWidget() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_gret := girepository.MustFind("Gtk", "SearchBar").InvokeMethod("get_key_capture_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(bar)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
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

// SetChild sets the child widget of bar.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (bar *SearchBar) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// SetKeyCaptureWidget sets widget as the widget that bar will capture key
// events from.
//
// If key events are handled by the search bar, the bar will be shown, and the
// entry populated with the entered text.
//
// Note that despite the name of this function, the events are only 'captured'
// in the bubble phase, which means that editable child widgets of widget will
// receive text input before it gets captured. If that is not desired, you can
// capture and forward the events yourself with
// gtk.EventControllerKey.Forward().
//
// The function takes the following parameters:
//
//    - widget (optional): GtkWidget.
//
func (bar *SearchBar) SetKeyCaptureWidget(widget Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if widget != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	girepository.MustFind("Gtk", "SearchBar").InvokeMethod("set_key_capture_widget", _args[:], nil)

	runtime.KeepAlive(bar)
	runtime.KeepAlive(widget)
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

// SetShowCloseButton shows or hides the close button.
//
// Applications that already have a “search” toggle button should not show a
// close button in their search bar, as it duplicates the role of the toggle
// button.
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
