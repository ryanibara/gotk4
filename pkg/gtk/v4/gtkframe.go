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
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFramer},
	})
}

// FrameOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type FrameOverrider interface {
	ComputeChildAllocation(allocation *Allocation)
}

// Frame: GtkFrame is a widget that surrounds its child with a decorative frame
// and an optional label.
//
// !An example GtkFrame (frame.png)
//
// If present, the label is drawn inside the top edge of the frame. The
// horizontal position of the label can be controlled with
// gtk.Frame.SetLabelAlign().
//
// GtkFrame clips its child. You can use this to add rounded corners to widgets,
// but be aware that it also cuts off shadows.
//
//
// GtkFrame as GtkBuildable
//
// The GtkFrame implementation of the GtkBuildable interface supports placing a
// child in the label position by specifying “label” as the “type” attribute of
// a <child> element. A normal content child can be specified without specifying
// a <child> type attribute.
//
// An example of a UI definition fragment with GtkFrame:
//
//    <object class="GtkFrame">
//      <child type="label">
//        <object class="GtkLabel" id="frame_label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame_content"/>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    frame
//    ├── <label widget>
//    ╰── <child>
//
//
// GtkFrame has a main CSS node with name “frame”, which is used to draw the
// visible border. You can set the appearance of the border using CSS properties
// like “border-style” on this node.
type Frame struct {
	Widget
}

func wrapFrame(obj *externglib.Object) *Frame {
	return &Frame{
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

func marshalFramer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapFrame(obj), nil
}

// NewFrame creates a new GtkFrame, with optional label label.
//
// If label is NULL, the label is omitted.
func NewFrame(label string) *Frame {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_frame_new(_arg1)
	runtime.KeepAlive(label)

	var _frame *Frame // out

	_frame = wrapFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _frame
}

// Child gets the child widget of frame.
func (frame *Frame) Child() Widgetter {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_child(_arg0)
	runtime.KeepAlive(frame)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Label returns the frame labels text.
//
// If the frame's label widget is not a GtkLabel, NULL is returned.
func (frame *Frame) Label() string {
	var _arg0 *C.GtkFrame // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_label(_arg0)
	runtime.KeepAlive(frame)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelAlign retrieves the X alignment of the frame’s label.
func (frame *Frame) LabelAlign() float32 {
	var _arg0 *C.GtkFrame // out
	var _cret C.float     // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_label_align(_arg0)
	runtime.KeepAlive(frame)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// LabelWidget retrieves the label widget for the frame.
func (frame *Frame) LabelWidget() Widgetter {
	var _arg0 *C.GtkFrame  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))

	_cret = C.gtk_frame_get_label_widget(_arg0)
	runtime.KeepAlive(frame)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Widgetter)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// SetChild sets the child widget of frame.
func (frame *Frame) SetChild(child Widgetter) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_frame_set_child(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(child)
}

// SetLabel creates a new GtkLabel with the label and sets it as the frame's
// label widget.
func (frame *Frame) SetLabel(label string) {
	var _arg0 *C.GtkFrame // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	if label != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_frame_set_label(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(label)
}

// SetLabelAlign sets the X alignment of the frame widget’s label.
//
// The default value for a newly created frame is 0.0.
func (frame *Frame) SetLabelAlign(xalign float32) {
	var _arg0 *C.GtkFrame // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	_arg1 = C.float(xalign)

	C.gtk_frame_set_label_align(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(xalign)
}

// SetLabelWidget sets the label widget for the frame.
//
// This is the widget that will appear embedded in the top edge of the frame as
// a title.
func (frame *Frame) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkFrame  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkFrame)(unsafe.Pointer(frame.Native()))
	if labelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))
	}

	C.gtk_frame_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(frame)
	runtime.KeepAlive(labelWidget)
}
