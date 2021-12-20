// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

type HeaderBarAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible
}

var (
	_ externglib.Objector = (*HeaderBarAccessible)(nil)
)

func wrapHeaderBarAccessible(obj *externglib.Object) *HeaderBarAccessible {
	return &HeaderBarAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
	}
}
