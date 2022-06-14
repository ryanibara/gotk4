// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkshortcutsgroup.go.
var GTypeShortcutsGroup = coreglib.Type(C.gtk_shortcuts_group_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeShortcutsGroup, F: marshalShortcutsGroup},
	})
}

// ShortcutsGroup represents a group of related keyboard shortcuts or gestures.
// The group has a title. It may optionally be associated with a view of the
// application, which can be used to show only relevant shortcuts depending on
// the application context.
//
// This widget is only meant to be used with ShortcutsWindow.
type ShortcutsGroup struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ShortcutsGroup)(nil)
	_ coreglib.Objector = (*ShortcutsGroup)(nil)
)

func wrapShortcutsGroup(obj *coreglib.Object) *ShortcutsGroup {
	return &ShortcutsGroup{
		Box: Box{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsGroup(p uintptr) (interface{}, error) {
	return wrapShortcutsGroup(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
