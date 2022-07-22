// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeHSeparator = coreglib.Type(C.gtk_hseparator_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHSeparator, F: marshalHSeparator},
	})
}

// HSeparatorOverrider contains methods that are overridable.
type HSeparatorOverrider interface {
}

// HSeparator widget is a horizontal separator, used to group the widgets within
// a window. It displays a horizontal line with a shadow to make it appear
// sunken into the interface.
//
// > The HSeparator widget is not used as a separator within menus. > To create
// a separator in a menu create an empty SeparatorMenuItem > widget using
// gtk_separator_menu_item_new() and add it to the menu with >
// gtk_menu_shell_append().
//
// GtkHSeparator has been deprecated, use Separator instead.
type HSeparator struct {
	_ [0]func() // equal guard
	Separator
}

var (
	_ Widgetter         = (*HSeparator)(nil)
	_ coreglib.Objector = (*HSeparator)(nil)
)

func initClassHSeparator(gclass unsafe.Pointer, goval any) {
}

func wrapHSeparator(obj *coreglib.Object) *HSeparator {
	return &HSeparator{
		Separator: Separator{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalHSeparator(p uintptr) (interface{}, error) {
	return wrapHSeparator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHSeparator creates a new HSeparator.
//
// Deprecated: Use gtk_separator_new() with GTK_ORIENTATION_HORIZONTAL instead.
//
// The function returns the following values:
//
//    - hSeparator: new HSeparator.
//
func NewHSeparator() *HSeparator {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_hseparator_new()

	var _hSeparator *HSeparator // out

	_hSeparator = wrapHSeparator(coreglib.Take(unsafe.Pointer(_cret)))

	return _hSeparator
}
