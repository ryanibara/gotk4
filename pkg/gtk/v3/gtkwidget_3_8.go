// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TickCallback: callback type for adding a function to update animations.
// See gtk_widget_add_tick_callback().
type TickCallback func(widget Widgetter, frameClock gdk.FrameClocker) (ok bool)
