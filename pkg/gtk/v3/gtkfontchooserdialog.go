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

// FontChooserDialogClass: instance of this type is always passed by reference.
type FontChooserDialogClass struct {
	*fontChooserDialogClass
}

// fontChooserDialogClass is the struct that's finalized.
type fontChooserDialogClass struct {
	native *C.GtkFontChooserDialogClass
}

// ParentClass: parent class.
func (f *FontChooserDialogClass) ParentClass() *DialogClass {
	valptr := &f.native.parent_class
	var _v *DialogClass // out
	_v = (*DialogClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
