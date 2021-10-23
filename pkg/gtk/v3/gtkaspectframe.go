// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFramer},
	})
}

// AspectFrame is useful when you want pack a widget so that it can resize but
// always retains the same aspect ratio. For instance, one might be drawing a
// small preview of a larger image. AspectFrame derives from Frame, so it can
// draw a label and a frame around the child. The frame will be “shrink-wrapped”
// to the size of the child.
//
//
// CSS nodes
//
// GtkAspectFrame uses a CSS node with name frame.
type AspectFrame struct {
	Frame
}

func wrapAspectFrame(obj *externglib.Object) *AspectFrame {
	return &AspectFrame{
		Frame: Frame{
			Bin: Bin{
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
			},
		},
	}
}

func marshalAspectFramer(p uintptr) (interface{}, error) {
	return wrapAspectFrame(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewAspectFrame: create a new AspectFrame.
//
// The function takes the following parameters:
//
//    - label: label text.
//    - xalign: horizontal alignment of the child within the allocation of the
//    AspectFrame. This ranges from 0.0 (left aligned) to 1.0 (right aligned).
//    - yalign: vertical alignment of the child within the allocation of the
//    AspectFrame. This ranges from 0.0 (top aligned) to 1.0 (bottom aligned).
//    - ratio: desired aspect ratio.
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken
//    from the requistion of the child.
//
func NewAspectFrame(label string, xalign, yalign, ratio float32, obeyChild bool) *AspectFrame {
	var _arg1 *C.gchar     // out
	var _arg2 C.gfloat     // out
	var _arg3 C.gfloat     // out
	var _arg4 C.gfloat     // out
	var _arg5 C.gboolean   // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = C.gfloat(xalign)
	_arg3 = C.gfloat(yalign)
	_arg4 = C.gfloat(ratio)
	if obeyChild {
		_arg5 = C.TRUE
	}

	_cret = C.gtk_aspect_frame_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(label)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(ratio)
	runtime.KeepAlive(obeyChild)

	var _aspectFrame *AspectFrame // out

	_aspectFrame = wrapAspectFrame(externglib.Take(unsafe.Pointer(_cret)))

	return _aspectFrame
}

// Set parameters for an existing AspectFrame.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment of the child within the allocation of the
//    AspectFrame. This ranges from 0.0 (left aligned) to 1.0 (right aligned).
//    - yalign: vertical alignment of the child within the allocation of the
//    AspectFrame. This ranges from 0.0 (top aligned) to 1.0 (bottom aligned).
//    - ratio: desired aspect ratio.
//    - obeyChild: if TRUE, ratio is ignored, and the aspect ratio is taken
//    from the requistion of the child.
//
func (aspectFrame *AspectFrame) Set(xalign, yalign, ratio float32, obeyChild bool) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.gfloat          // out
	var _arg2 C.gfloat          // out
	var _arg3 C.gfloat          // out
	var _arg4 C.gboolean        // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(aspectFrame.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(ratio)
	if obeyChild {
		_arg4 = C.TRUE
	}

	C.gtk_aspect_frame_set(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(aspectFrame)
	runtime.KeepAlive(xalign)
	runtime.KeepAlive(yalign)
	runtime.KeepAlive(ratio)
	runtime.KeepAlive(obeyChild)
}
