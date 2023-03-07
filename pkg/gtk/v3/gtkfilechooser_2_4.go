// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// FileChooserErrorQuark registers an error quark for FileChooser if necessary.
//
// The function returns the following values:
//
//    - quark: error quark used for FileChooser errors.
//
func FileChooserErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_file_chooser_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}