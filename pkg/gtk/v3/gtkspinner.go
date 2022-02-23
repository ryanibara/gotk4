// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// glib.Type values for gtkspinner.go.
var GTypeSpinner = externglib.Type(C.gtk_spinner_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeSpinner, F: marshalSpinner},
	})
}

// SpinnerOverrider contains methods that are overridable.
type SpinnerOverrider interface {
}

// Spinner widget displays an icon-size spinning animation. It is often used as
// an alternative to a ProgressBar for displaying indefinite activity, instead
// of actual progress.
//
// To start the animation, use gtk_spinner_start(), to stop it use
// gtk_spinner_stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*Spinner)(nil)
)

func classInitSpinnerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSpinner(obj *externglib.Object) *Spinner {
	return &Spinner{
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
	}
}

func marshalSpinner(p uintptr) (interface{}, error) {
	return wrapSpinner(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpinner returns a new spinner widget. Not yet started.
//
// The function returns the following values:
//
//    - spinner: new Spinner.
//
func NewSpinner() *Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner *Spinner // out

	_spinner = wrapSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _spinner
}

// Start starts the animation of the spinner.
func (spinner *Spinner) Start() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_start(_arg0)
	runtime.KeepAlive(spinner)
}

// Stop stops the animation of the spinner.
func (spinner *Spinner) Stop() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	C.gtk_spinner_stop(_arg0)
	runtime.KeepAlive(spinner)
}
