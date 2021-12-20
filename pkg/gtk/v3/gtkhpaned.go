// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hpaned_get_type()), F: marshalHPanedder},
	})
}

// HPaned widget is a container widget with two children arranged horizontally.
// The division between the two panes is adjustable by the user by dragging a
// handle. See Paned for details.
//
// GtkHPaned has been deprecated, use Paned instead.
type HPaned struct {
	_ [0]func() // equal guard
	Paned
}

var (
	_ Containerer         = (*HPaned)(nil)
	_ externglib.Objector = (*HPaned)(nil)
)

func wrapHPaned(obj *externglib.Object) *HPaned {
	return &HPaned{
		Paned: Paned{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalHPanedder(p uintptr) (interface{}, error) {
	return wrapHPaned(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHPaned: create a new HPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function returns the following values:
//
//    - hPaned: new HPaned.
//
func NewHPaned() *HPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hpaned_new()

	var _hPaned *HPaned // out

	_hPaned = wrapHPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _hPaned
}
