// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeOffscreenWindow returns the GType for the type OffscreenWindow.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeOffscreenWindow() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "OffscreenWindow").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalOffscreenWindow)
	return gtype
}

// OffscreenWindowOverrider contains methods that are overridable.
type OffscreenWindowOverrider interface {
}

// OffscreenWindow is strictly intended to be used for obtaining snapshots of
// widgets that are not part of a normal widget hierarchy. Since OffscreenWindow
// is a toplevel widget you cannot obtain snapshots of a full window with it
// since you cannot pack a toplevel widget in another toplevel.
//
// The idea is to take a widget and manually set the state of it, add it to a
// GtkOffscreenWindow and then retrieve the snapshot as a #cairo_surface_t or
// Pixbuf.
//
// GtkOffscreenWindow derives from Window only as an implementation detail.
// Applications should not use any API specific to Window to operate on this
// object. It should be treated as a Bin that has no parent widget.
//
// When contained offscreen widgets are redrawn, GtkOffscreenWindow will emit a
// Widget::damage-event signal.
type OffscreenWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*OffscreenWindow)(nil)
)

func classInitOffscreenWindower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapOffscreenWindow(obj *coreglib.Object) *OffscreenWindow {
	return &OffscreenWindow{
		Window: Window{
			Bin: Bin{
				Container: Container{
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
				},
			},
		},
	}
}

func marshalOffscreenWindow(p uintptr) (interface{}, error) {
	return wrapOffscreenWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewOffscreenWindow creates a toplevel container widget that is used to
// retrieve snapshots of widgets without showing them on the screen.
//
// The function returns the following values:
//
//    - offscreenWindow: pointer to a Widget.
//
func NewOffscreenWindow() *OffscreenWindow {
	_gret := girepository.MustFind("Gtk", "OffscreenWindow").InvokeMethod("new_OffscreenWindow", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _offscreenWindow *OffscreenWindow // out

	_offscreenWindow = wrapOffscreenWindow(coreglib.Take(unsafe.Pointer(_cret)))

	return _offscreenWindow
}

// Pixbuf retrieves a snapshot of the contained widget in the form of a Pixbuf.
// This is a new pixbuf with a reference count of 1, and the application should
// unreference it once it is no longer needed.
//
// The function returns the following values:
//
//    - pixbuf (optional) pointer, or NULL.
//
func (offscreen *OffscreenWindow) Pixbuf() *gdkpixbuf.Pixbuf {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(offscreen).Native()))

	_gret := girepository.MustFind("Gtk", "OffscreenWindow").InvokeMethod("get_pixbuf", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(offscreen)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// Surface retrieves a snapshot of the contained widget in the form of a
// #cairo_surface_t. If you need to keep this around over window resizes then
// you should add a reference to it.
//
// The function returns the following values:
//
//    - surface (optional) pointer to the offscreen surface, or NULL.
//
func (offscreen *OffscreenWindow) Surface() *cairo.Surface {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(offscreen).Native()))

	_gret := girepository.MustFind("Gtk", "OffscreenWindow").InvokeMethod("get_surface", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(offscreen)

	var _surface *cairo.Surface // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(_cret)))
		C.cairo_surface_reference(_cret)
		runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
			C.cairo_surface_destroy((*C.void)(unsafe.Pointer(v.Native())))
		})
	}

	return _surface
}
