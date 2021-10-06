// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_spinner_get_type()), F: marshalSpinnerer},
	})
}

// Spinner: GtkSpinner widget displays an icon-size spinning animation.
//
// It is often used as an alternative to a gtk.ProgressBar for displaying
// indefinite activity, instead of actual progress.
//
// !An example GtkSpinner (spinner.png)
//
// To start the animation, use gtk.Spinner.Start(), to stop it use
// gtk.Spinner.Stop().
//
//
// CSS nodes
//
// GtkSpinner has a single CSS node with the name spinner. When the animation is
// active, the :checked pseudoclass is added to this node.
type Spinner struct {
	Widget
}

func wrapSpinner(obj *externglib.Object) *Spinner {
	return &Spinner{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalSpinnerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSpinner(obj), nil
}

// NewSpinner returns a new spinner widget. Not yet started.
func NewSpinner() *Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner *Spinner // out

	_spinner = wrapSpinner(externglib.Take(unsafe.Pointer(_cret)))

	return _spinner
}

// Spinning returns whether the spinner is spinning.
func (spinner *Spinner) Spinning() bool {
	var _arg0 *C.GtkSpinner // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))

	_cret = C.gtk_spinner_get_spinning(_arg0)
	runtime.KeepAlive(spinner)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSpinning sets the activity of the spinner.
//
// The function takes the following parameters:
//
//    - spinning: whether the spinner should be spinning.
//
func (spinner *Spinner) SetSpinning(spinning bool) {
	var _arg0 *C.GtkSpinner // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(spinner.Native()))
	if spinning {
		_arg1 = C.TRUE
	}

	C.gtk_spinner_set_spinning(_arg0, _arg1)
	runtime.KeepAlive(spinner)
	runtime.KeepAlive(spinning)
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
