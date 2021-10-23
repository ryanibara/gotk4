// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
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
		{T: externglib.Type(C.gtk_button_box_style_get_type()), F: marshalButtonBoxStyle},
		{T: externglib.Type(C.gtk_button_box_get_type()), F: marshalButtonBoxer},
	})
}

// ButtonBoxStyle: used to dictate the style that a ButtonBox uses to layout the
// buttons it contains.
type ButtonBoxStyle C.gint

const (
	// ButtonboxSpread buttons are evenly spread across the box.
	ButtonboxSpread ButtonBoxStyle = 1
	// ButtonboxEdge buttons are placed at the edges of the box.
	ButtonboxEdge ButtonBoxStyle = 2
	// ButtonboxStart buttons are grouped towards the start of the box, (on the
	// left for a HBox, or the top for a VBox).
	ButtonboxStart ButtonBoxStyle = 3
	// ButtonboxEnd buttons are grouped towards the end of the box, (on the
	// right for a HBox, or the bottom for a VBox).
	ButtonboxEnd ButtonBoxStyle = 4
	// ButtonboxCenter buttons are centered in the box. Since 2.12.
	ButtonboxCenter ButtonBoxStyle = 5
	// ButtonboxExpand buttons expand to fill the box. This entails giving
	// buttons a "linked" appearance, making button sizes homogeneous, and
	// setting spacing to 0 (same as calling gtk_box_set_homogeneous() and
	// gtk_box_set_spacing() manually). Since 3.12.
	ButtonboxExpand ButtonBoxStyle = 6
)

func marshalButtonBoxStyle(p uintptr) (interface{}, error) {
	return ButtonBoxStyle(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ButtonBoxStyle.
func (b ButtonBoxStyle) String() string {
	switch b {
	case ButtonboxSpread:
		return "Spread"
	case ButtonboxEdge:
		return "Edge"
	case ButtonboxStart:
		return "Start"
	case ButtonboxEnd:
		return "End"
	case ButtonboxCenter:
		return "Center"
	case ButtonboxExpand:
		return "Expand"
	default:
		return fmt.Sprintf("ButtonBoxStyle(%d)", b)
	}
}

type ButtonBox struct {
	Box
}

func wrapButtonBox(obj *externglib.Object) *ButtonBox {
	return &ButtonBox{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalButtonBoxer(p uintptr) (interface{}, error) {
	return wrapButtonBox(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewButtonBox creates a new ButtonBox.
//
// The function takes the following parameters:
//
//    - orientation box's orientation.
//
func NewButtonBox(orientation Orientation) *ButtonBox {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_button_box_new(_arg1)
	runtime.KeepAlive(orientation)

	var _buttonBox *ButtonBox // out

	_buttonBox = wrapButtonBox(externglib.Take(unsafe.Pointer(_cret)))

	return _buttonBox
}

// ChildNonHomogeneous returns whether the child is exempted from homogenous
// sizing.
//
// The function takes the following parameters:
//
//    - child of widget.
//
func (widget *ButtonBox) ChildNonHomogeneous(child Widgetter) bool {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_button_box_get_child_non_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ChildSecondary returns whether child should appear in a secondary group of
// children.
//
// The function takes the following parameters:
//
//    - child of widget.
//
func (widget *ButtonBox) ChildSecondary(child Widgetter) bool {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_button_box_get_child_secondary(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(child)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout retrieves the method being used to arrange the buttons in a button
// box.
func (widget *ButtonBox) Layout() ButtonBoxStyle {
	var _arg0 *C.GtkButtonBox     // out
	var _cret C.GtkButtonBoxStyle // in

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_button_box_get_layout(_arg0)
	runtime.KeepAlive(widget)

	var _buttonBoxStyle ButtonBoxStyle // out

	_buttonBoxStyle = ButtonBoxStyle(_cret)

	return _buttonBoxStyle
}

// SetChildNonHomogeneous sets whether the child is exempted from homogeous
// sizing.
//
// The function takes the following parameters:
//
//    - child of widget.
//    - nonHomogeneous: new value.
//
func (widget *ButtonBox) SetChildNonHomogeneous(child Widgetter, nonHomogeneous bool) {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if nonHomogeneous {
		_arg2 = C.TRUE
	}

	C.gtk_button_box_set_child_non_homogeneous(_arg0, _arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(child)
	runtime.KeepAlive(nonHomogeneous)
}

// SetChildSecondary sets whether child should appear in a secondary group of
// children. A typical use of a secondary child is the help button in a dialog.
//
// This group appears after the other children if the style is
// GTK_BUTTONBOX_START, GTK_BUTTONBOX_SPREAD or GTK_BUTTONBOX_EDGE, and before
// the other children if the style is GTK_BUTTONBOX_END. For horizontal button
// boxes, the definition of before/after depends on direction of the widget (see
// gtk_widget_set_direction()). If the style is GTK_BUTTONBOX_START or
// GTK_BUTTONBOX_END, then the secondary children are aligned at the other end
// of the button box from the main children. For the other styles, they appear
// immediately next to the main children.
//
// The function takes the following parameters:
//
//    - child of widget.
//    - isSecondary: if TRUE, the child appears in a secondary group of the
//    button box.
//
func (widget *ButtonBox) SetChildSecondary(child Widgetter, isSecondary bool) {
	var _arg0 *C.GtkButtonBox // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if isSecondary {
		_arg2 = C.TRUE
	}

	C.gtk_button_box_set_child_secondary(_arg0, _arg1, _arg2)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(child)
	runtime.KeepAlive(isSecondary)
}

// SetLayout changes the way buttons are arranged in their container.
//
// The function takes the following parameters:
//
//    - layoutStyle: new layout style.
//
func (widget *ButtonBox) SetLayout(layoutStyle ButtonBoxStyle) {
	var _arg0 *C.GtkButtonBox     // out
	var _arg1 C.GtkButtonBoxStyle // out

	_arg0 = (*C.GtkButtonBox)(unsafe.Pointer(widget.Native()))
	_arg1 = C.GtkButtonBoxStyle(layoutStyle)

	C.gtk_button_box_set_layout(_arg0, _arg1)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(layoutStyle)
}
