// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_alignment_get_type()), F: marshalAlignmenter},
	})
}

// Alignment widget controls the alignment and size of its child widget. It has
// four settings: xscale, yscale, xalign, and yalign.
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
type Alignment struct {
	Bin

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Binner = (*Alignment)(nil)
)

func wrapAlignment(obj *externglib.Object) *Alignment {
	return &Alignment{
		Bin: Bin{
			Container: Container{
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
			},
		},
	}
}

func marshalAlignmenter(p uintptr) (interface{}, error) {
	return wrapAlignment(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAlignment creates a new Alignment.
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment of the child widget, from 0 (left) to 1
//      (right).
//    - yalign: vertical alignment of the child widget, from 0 (top) to 1
//      (bottom).
//    - xscale: amount that the child widget expands horizontally to fill up
//      unused space, from 0 to 1. A value of 0 indicates that the child widget
//      should never expand. A value of 1 indicates that the child widget will
//      expand to fill all of the space allocated for the Alignment.
//    - yscale: amount that the child widget expands vertically to fill up unused
//      space, from 0 to 1. The values are similar to xscale.
//
// The function returns the following values:
//
//    - alignment: new Alignment.
//
func NewAlignment(xalign, yalign, xscale, yscale float32) *Alignment {
	var _arg1 C.gfloat     // out
	var _arg2 C.gfloat     // out
	var _arg3 C.gfloat     // out
	var _arg4 C.gfloat     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(xscale)
	_arg4 = C.gfloat(yscale)

	_cret = C.gtk_alignment_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(xscale)
	runtime.KeepAlive(yscale)

	var _alignment *Alignment // out

	_alignment = wrapAlignment(externglib.Take(unsafe.Pointer(_cret)))

	return _alignment
}

// Padding gets the padding on the different sides of the widget. See
// gtk_alignment_set_padding ().
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function returns the following values:
//
//    - paddingTop (optional): location to store the padding for the top of the
//      widget, or NULL.
//    - paddingBottom (optional): location to store the padding for the bottom of
//      the widget, or NULL.
//    - paddingLeft (optional): location to store the padding for the left of the
//      widget, or NULL.
//    - paddingRight (optional): location to store the padding for the right of
//      the widget, or NULL.
//
func (alignment *Alignment) Padding() (paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint) {
	var _arg0 *C.GtkAlignment // out
	var _arg1 C.guint         // in
	var _arg2 C.guint         // in
	var _arg3 C.guint         // in
	var _arg4 C.guint         // in

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(alignment.Native()))

	C.gtk_alignment_get_padding(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(alignment)

	var _paddingTop uint    // out
	var _paddingBottom uint // out
	var _paddingLeft uint   // out
	var _paddingRight uint  // out

	_paddingTop = uint(_arg1)
	_paddingBottom = uint(_arg2)
	_paddingLeft = uint(_arg3)
	_paddingRight = uint(_arg4)

	return _paddingTop, _paddingBottom, _paddingLeft, _paddingRight
}

// Set sets the Alignment values.
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment of the child widget, from 0 (left) to 1
//      (right).
//    - yalign: vertical alignment of the child widget, from 0 (top) to 1
//      (bottom).
//    - xscale: amount that the child widget expands horizontally to fill up
//      unused space, from 0 to 1. A value of 0 indicates that the child widget
//      should never expand. A value of 1 indicates that the child widget will
//      expand to fill all of the space allocated for the Alignment.
//    - yscale: amount that the child widget expands vertically to fill up unused
//      space, from 0 to 1. The values are similar to xscale.
//
func (alignment *Alignment) Set(xalign, yalign, xscale, yscale float32) {
	var _arg0 *C.GtkAlignment // out
	var _arg1 C.gfloat        // out
	var _arg2 C.gfloat        // out
	var _arg3 C.gfloat        // out
	var _arg4 C.gfloat        // out

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(alignment.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(xscale)
	_arg4 = C.gfloat(yscale)

	C.gtk_alignment_set(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(alignment)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(xscale)
	runtime.KeepAlive(yscale)
}

// SetPadding sets the padding on the different sides of the widget. The padding
// adds blank space to the sides of the widget. For instance, this can be used
// to indent the child widget towards the right by adding padding on the left.
//
// Deprecated: Use Widget alignment and margin properties.
//
// The function takes the following parameters:
//
//    - paddingTop: padding at the top of the widget.
//    - paddingBottom: padding at the bottom of the widget.
//    - paddingLeft: padding at the left of the widget.
//    - paddingRight: padding at the right of the widget.
//
func (alignment *Alignment) SetPadding(paddingTop, paddingBottom, paddingLeft, paddingRight uint) {
	var _arg0 *C.GtkAlignment // out
	var _arg1 C.guint         // out
	var _arg2 C.guint         // out
	var _arg3 C.guint         // out
	var _arg4 C.guint         // out

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(alignment.Native()))
	_arg1 = C.guint(paddingTop)
	_arg2 = C.guint(paddingBottom)
	_arg3 = C.guint(paddingLeft)
	_arg4 = C.guint(paddingRight)

	C.gtk_alignment_set_padding(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(alignment)
	runtime.KeepAlive(paddingTop)
	runtime.KeepAlive(paddingBottom)
	runtime.KeepAlive(paddingLeft)
	runtime.KeepAlive(paddingRight)
}
