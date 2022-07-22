// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ShortcutsWindowClass_close(GtkShortcutsWindow*);
// extern void _gotk4_gtk3_ShortcutsWindowClass_search(GtkShortcutsWindow*);
// extern void _gotk4_gtk3_ShortcutsWindow_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk3_ShortcutsWindow_ConnectSearch(gpointer, guintptr);
import "C"

// GType values.
var (
	GTypeShortcutsWindow = coreglib.Type(C.gtk_shortcuts_window_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeShortcutsWindow, F: marshalShortcutsWindow},
	})
}

// ShortcutsWindowOverrider contains methods that are overridable.
type ShortcutsWindowOverrider interface {
	Close()
	Search()
}

// ShortcutsWindow shows brief information about the keyboard shortcuts and
// gestures of an application. The shortcuts can be grouped, and you can have
// multiple sections in this window, corresponding to the major modes of your
// application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with GtkBuilder, by
// populating a ShortcutsWindow with one or more ShortcutsSection objects, which
// contain ShortcutsGroups that in turn contain objects of class
// ShortcutsShortcut.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a ShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a ShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*ShortcutsWindow)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:        GTypeShortcutsWindow,
		GoType:       reflect.TypeOf((*ShortcutsWindow)(nil)),
		InitClass:    initClassShortcutsWindow,
		ClassSize:    uint32(unsafe.Sizeof(C.GtkShortcutsWindow{})),
		InstanceSize: uint32(unsafe.Sizeof(C.GtkShortcutsWindowClass{})),
	})
}

func initClassShortcutsWindow(gclass unsafe.Pointer, goval any) {

	pclass := (*C.GtkShortcutsWindowClass)(unsafe.Pointer(gclass))

	if _, ok := goval.(interface{ Close() }); ok {
		pclass.close = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_close)
	}

	if _, ok := goval.(interface{ Search() }); ok {
		pclass.search = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_search)
	}
}

//export _gotk4_gtk3_ShortcutsWindowClass_close
func _gotk4_gtk3_ShortcutsWindowClass_close(arg0 *C.GtkShortcutsWindow) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk3_ShortcutsWindowClass_search
func _gotk4_gtk3_ShortcutsWindowClass_search(arg0 *C.GtkShortcutsWindow) {
	goval := coreglib.GoObjectFromInstance(unsafe.Pointer(arg0))
	iface := goval.(interface{ Search() })

	iface.Search()
}

func wrapShortcutsWindow(obj *coreglib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{
		Window: Window{
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
		},
	}
}

func marshalShortcutsWindow(p uintptr) (interface{}, error) {
	return wrapShortcutsWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_ShortcutsWindow_ConnectClose
func _gotk4_gtk3_ShortcutsWindow_ConnectClose(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectClose signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to close the window.
//
// The default binding for this signal is the Escape key.
func (self *ShortcutsWindow) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "close", false, unsafe.Pointer(C._gotk4_gtk3_ShortcutsWindow_ConnectClose), f)
}

//export _gotk4_gtk3_ShortcutsWindow_ConnectSearch
func _gotk4_gtk3_ShortcutsWindow_ConnectSearch(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectSearch signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to start a search.
//
// The default binding for this signal is Control-F.
func (self *ShortcutsWindow) ConnectSearch(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "search", false, unsafe.Pointer(C._gotk4_gtk3_ShortcutsWindow_ConnectSearch), f)
}
