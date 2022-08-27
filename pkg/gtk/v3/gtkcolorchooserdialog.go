// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ColorChooserDialogClass: instance of this type is always passed by reference.
type ColorChooserDialogClass struct {
	*colorChooserDialogClass
}

// colorChooserDialogClass is the struct that's finalized.
type colorChooserDialogClass struct {
	native *C.GtkColorChooserDialogClass
}

func (c *ColorChooserDialogClass) ParentClass() *DialogClass {
	valptr := &c.native.parent_class
	var _v *DialogClass // out
	_v = (*DialogClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
