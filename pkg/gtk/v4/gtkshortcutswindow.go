// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_window_get_type()), F: marshalShortcutsWindower},
	})
}

// ShortcutsWindow: GtkShortcutsWindow shows information about the keyboard
// shortcuts and gestures of an application.
//
// The shortcuts can be grouped, and you can have multiple sections in this
// window, corresponding to the major modes of your application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with gtk.Builder, by
// populating a GtkShortcutsWindow with one or more GtkShortcutsSection objects,
// which contain GtkShortcutsGroups that in turn contain objects of class
// GtkShortcutsShortcut.
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
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a GtkShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a GtkShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://gitlab.gnome.org/GNOME/gtk/tree/master/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Widgetter           = (*ShortcutsWindow)(nil)
	_ externglib.Objector = (*ShortcutsWindow)(nil)
)

func wrapShortcutsWindow(obj *externglib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
			Object: obj,
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
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
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsWindower(p uintptr) (interface{}, error) {
	return wrapShortcutsWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose: emitted when the user uses a keybinding to close the window.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is the Escape key.
func (v *ShortcutsWindow) ConnectClose(f func()) externglib.SignalHandle {
	return v.Connect("close", f)
}

// ConnectSearch: emitted when the user uses a keybinding to start a search.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is Control-F.
func (v *ShortcutsWindow) ConnectSearch(f func()) externglib.SignalHandle {
	return v.Connect("search", f)
}
