// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_alignment_get_type()), F: marshalAlignment},
	})
}

// Alignment: the Alignment widget controls the alignment and size of its child
// widget. It has four settings: xscale, yscale, xalign, and yalign.
//
// The scale settings are used to specify how much the child widget should
// expand to fill the space allocated to the Alignment. The values can range
// from 0 (meaning the child doesn’t expand at all) to 1 (meaning the child
// expands to fill all of the available space).
//
// The align settings are used to place the child widget within the available
// area. The values range from 0 (top or left) to 1 (bottom or right). Of
// course, if the scale settings are both set to 1, the alignment settings have
// no effect.
//
// GtkAlignment has been deprecated in 3.14 and should not be used in
// newly-written code. The desired effect can be achieved by using the
// Widget:halign, Widget:valign and Widget:margin properties on the child
// widget.
type Alignment interface {
	Bin
	Buildable

	// Padding gets the padding on the different sides of the widget. See
	// gtk_alignment_set_padding ().
	Padding() (paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint)
	// Set sets the Alignment values.
	Set(xalign float32, yalign float32, xscale float32, yscale float32)
	// SetPadding sets the padding on the different sides of the widget. The
	// padding adds blank space to the sides of the widget. For instance, this
	// can be used to indent the child widget towards the right by adding
	// padding on the left.
	SetPadding(paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint)
}

// alignment implements the Alignment interface.
type alignment struct {
	Bin
	Buildable
}

var _ Alignment = (*alignment)(nil)

// WrapAlignment wraps a GObject to the right type. It is
// primarily used internally.
func WrapAlignment(obj *externglib.Object) Alignment {
	return Alignment{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAlignment(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAlignment(obj), nil
}

// NewAlignment constructs a class Alignment.
func NewAlignment(xalign float32, yalign float32, xscale float32, yscale float32) Alignment {
	var arg1 C.gfloat
	var arg2 C.gfloat
	var arg3 C.gfloat
	var arg4 C.gfloat

	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)
	arg3 = C.gfloat(xscale)
	arg4 = C.gfloat(yscale)

	var cret C.GtkAlignment
	var ret1 Alignment

	cret = C.gtk_alignment_new(xalign, yalign, xscale, yscale)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Alignment)

	return ret1
}

// Padding gets the padding on the different sides of the widget. See
// gtk_alignment_set_padding ().
func (a alignment) Padding() (paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint) {
	var arg0 *C.GtkAlignment

	arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))

	var arg1 C.guint
	var ret1 uint
	var arg2 C.guint
	var ret2 uint
	var arg3 C.guint
	var ret3 uint
	var arg4 C.guint
	var ret4 uint

	C.gtk_alignment_get_padding(arg0, &arg1, &arg2, &arg3, &arg4)

	*ret1 = C.guint(arg1)
	*ret2 = C.guint(arg2)
	*ret3 = C.guint(arg3)
	*ret4 = C.guint(arg4)

	return ret1, ret2, ret3, ret4
}

// Set sets the Alignment values.
func (a alignment) Set(xalign float32, yalign float32, xscale float32, yscale float32) {
	var arg0 *C.GtkAlignment
	var arg1 C.gfloat
	var arg2 C.gfloat
	var arg3 C.gfloat
	var arg4 C.gfloat

	arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))
	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)
	arg3 = C.gfloat(xscale)
	arg4 = C.gfloat(yscale)

	C.gtk_alignment_set(arg0, xalign, yalign, xscale, yscale)
}

// SetPadding sets the padding on the different sides of the widget. The
// padding adds blank space to the sides of the widget. For instance, this
// can be used to indent the child widget towards the right by adding
// padding on the left.
func (a alignment) SetPadding(paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint) {
	var arg0 *C.GtkAlignment
	var arg1 C.guint
	var arg2 C.guint
	var arg3 C.guint
	var arg4 C.guint

	arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))
	arg1 = C.guint(paddingTop)
	arg2 = C.guint(paddingBottom)
	arg3 = C.guint(paddingLeft)
	arg4 = C.guint(paddingRight)

	C.gtk_alignment_set_padding(arg0, paddingTop, paddingBottom, paddingLeft, paddingRight)
}

type AlignmentPrivate struct {
	native C.GtkAlignmentPrivate
}

// WrapAlignmentPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAlignmentPrivate(ptr unsafe.Pointer) *AlignmentPrivate {
	if ptr == nil {
		return nil
	}

	return (*AlignmentPrivate)(ptr)
}

func marshalAlignmentPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAlignmentPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AlignmentPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}
