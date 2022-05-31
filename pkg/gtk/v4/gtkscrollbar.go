// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkscrollbar.go.
var GTypeScrollbar = coreglib.Type(C.gtk_scrollbar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeScrollbar, F: marshalScrollbar},
	})
}

// Scrollbar: GtkScrollbar widget is a horizontal or vertical scrollbar.
//
// !An example GtkScrollbar (scrollbar.png)
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by gtk.Scrollbar.New. See [class.Gtk.Adjustment] for more details.
// The gtk.Adjustment:value field sets the position of the thumb and must be
// between gtk.Adjustment:lower and gtk.Adjustment:upper -
// gtk.Adjustment:page-size. The gtk.Adjustment:page-size represents the size of
// the visible scrollable area.
//
// The fields gtk.Adjustment:step-increment and gtk.Adjustment:page-increment
// fields are added to or subtracted from the gtk.Adjustment:value when the user
// asks to move by a step (using e.g. the cursor arrow keys) or by a page (using
// e.g. the Page Down/Up keys).
//
// CSS nodes
//
//    scrollbar
//    ╰── range[.fine-tune]
//        ╰── trough
//            ╰── slider
//
//
// GtkScrollbar has a main CSS node with name scrollbar and a subnode for its
// contents. The main node gets the .horizontal or .vertical style classes
// applied, depending on the scrollbar's orientation.
//
// The range node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// Other style classes that may be added to scrollbars inside gtk.ScrolledWindow
// include the positional classes (.left, .right, .top, .bottom) and style
// classes related to overlay scrolling (.overlay-indicator, .dragging,
// .hovering).
//
//
// Accessibility
//
// GtkScrollbar uses the GTK_ACCESSIBLE_ROLE_SCROLLBAR role.
type Scrollbar struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*Scrollbar)(nil)
	_ coreglib.Objector = (*Scrollbar)(nil)
)

func wrapScrollbar(obj *coreglib.Object) *Scrollbar {
	return &Scrollbar{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	return wrapScrollbar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Adjustment returns the scrollbar's adjustment.
//
// The function returns the following values:
//
//    - adjustment scrollbar's adjustment.
//
func (self *Scrollbar) Adjustment() *Adjustment {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Scrollbar").InvokeMethod("get_adjustment", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// SetAdjustment makes the scrollbar use the given adjustment.
//
// The function takes the following parameters:
//
//    - adjustment (optional) to set.
//
func (self *Scrollbar) SetAdjustment(adjustment *Adjustment) {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if adjustment != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Scrollbar").InvokeMethod("set_adjustment", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(adjustment)
}
