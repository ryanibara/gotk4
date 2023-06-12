// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
	GTypeSpinner = coreglib.Type(C.gtk_spinner_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSpinner, F: marshalSpinner},
	})
}

// SpinnerOverrides contains methods that are overridable.
type SpinnerOverrides struct {
}

func defaultSpinnerOverrides(v *Spinner) SpinnerOverrides {
	return SpinnerOverrides{}
}

// Spinner widget displays an icon-size spinning animation. It is often used
// as an alternative to a ProgressBar for displaying indefinite activity,
// instead of actual progress.
//
// To start the animation, use gtk_spinner_start(), to stop it use
// gtk_spinner_stop().
//
// # CSS nodes
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

func init() {
	coreglib.RegisterClassInfo[*Spinner, *SpinnerClass, SpinnerOverrides](
		GTypeSpinner,
		initSpinnerClass,
		wrapSpinner,
		defaultSpinnerOverrides,
	)
}

func initSpinnerClass(gclass unsafe.Pointer, overrides SpinnerOverrides, classInitFunc func(*SpinnerClass)) {
	if classInitFunc != nil {
		class := (*SpinnerClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapSpinner(obj *coreglib.Object) *Spinner {
	return &Spinner{
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
	}
}

func marshalSpinner(p uintptr) (interface{}, error) {
	return wrapSpinner(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpinner returns a new spinner widget. Not yet started.
//
// The function returns the following values:
//
//   - spinner: new Spinner.
//
func NewSpinner() *Spinner {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_spinner_new()

	var _spinner *Spinner // out

	_spinner = wrapSpinner(coreglib.Take(unsafe.Pointer(_cret)))

	return _spinner
}

// Start starts the animation of the spinner.
func (spinner *Spinner) Start() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(coreglib.InternObject(spinner).Native()))

	C.gtk_spinner_start(_arg0)
	runtime.KeepAlive(spinner)
}

// Stop stops the animation of the spinner.
func (spinner *Spinner) Stop() {
	var _arg0 *C.GtkSpinner // out

	_arg0 = (*C.GtkSpinner)(unsafe.Pointer(coreglib.InternObject(spinner).Native()))

	C.gtk_spinner_stop(_arg0)
	runtime.KeepAlive(spinner)
}

// SpinnerClass: instance of this type is always passed by reference.
type SpinnerClass struct {
	*spinnerClass
}

// spinnerClass is the struct that's finalized.
type spinnerClass struct {
	native *C.GtkSpinnerClass
}

func (s *SpinnerClass) ParentClass() *WidgetClass {
	valptr := &s.native.parent_class
	var _v *WidgetClass // out
	_v = (*WidgetClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
