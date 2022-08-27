// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewModelButton creates a new GtkModelButton.
//
// The function returns the following values:
//
//    - modelButton: newly created ModelButton widget.
//
func NewModelButton() *ModelButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_model_button_new()

	var _modelButton *ModelButton // out

	_modelButton = wrapModelButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _modelButton
}
