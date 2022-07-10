// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_gtk3_ThemingEngineClass_render_activity(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_arrow(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_background(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_check(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_expander(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_focus(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_frame(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_handle(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_icon(void*, void*, GdkPixbuf*, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_icon_surface(void*, void*, void*, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_layout(void*, void*, gdouble, gdouble, void*);
// extern void _gotk4_gtk3_ThemingEngineClass_render_line(void*, void*, gdouble, gdouble, gdouble, gdouble);
// extern void _gotk4_gtk3_ThemingEngineClass_render_option(void*, void*, gdouble, gdouble, gdouble, gdouble);
import "C"

// GTypeThemingEngine returns the GType for the type ThemingEngine.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeThemingEngine() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ThemingEngine").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalThemingEngine)
	return gtype
}

// ThemingEngineOverrider contains methods that are overridable.
type ThemingEngineOverrider interface {
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderActivity(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - angle
	//    - x
	//    - y
	//    - size
	//
	RenderArrow(cr *cairo.Context, angle, x, y, size float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderBackground(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderCheck(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderExpander(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFocus(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFrame(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderHandle(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - pixbuf
	//    - x
	//    - y
	//
	RenderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - surface
	//    - x
	//    - y
	//
	RenderIconSurface(cr *cairo.Context, surface *cairo.Surface, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - layout
	//
	RenderLayout(cr *cairo.Context, x, y float64, layout *pango.Layout)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x0
	//    - y0
	//    - x1
	//    - y1
	//
	RenderLine(cr *cairo.Context, x0, y0, x1, y1 float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderOption(cr *cairo.Context, x, y, width, height float64)
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ThemingEngine)(nil)
)

func classInitThemingEnginer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gtk", "ThemingEngineClass")

	if _, ok := goval.(interface {
		RenderActivity(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_activity")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_activity)
	}

	if _, ok := goval.(interface {
		RenderArrow(cr *cairo.Context, angle, x, y, size float64)
	}); ok {
		o := pclass.StructFieldOffset("render_arrow")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_arrow)
	}

	if _, ok := goval.(interface {
		RenderBackground(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_background")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_background)
	}

	if _, ok := goval.(interface {
		RenderCheck(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_check")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_check)
	}

	if _, ok := goval.(interface {
		RenderExpander(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_expander")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_expander)
	}

	if _, ok := goval.(interface {
		RenderFocus(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_focus")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_focus)
	}

	if _, ok := goval.(interface {
		RenderFrame(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_frame")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_frame)
	}

	if _, ok := goval.(interface {
		RenderHandle(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_handle")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_handle)
	}

	if _, ok := goval.(interface {
		RenderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64)
	}); ok {
		o := pclass.StructFieldOffset("render_icon")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_icon)
	}

	if _, ok := goval.(interface {
		RenderIconSurface(cr *cairo.Context, surface *cairo.Surface, x, y float64)
	}); ok {
		o := pclass.StructFieldOffset("render_icon_surface")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_icon_surface)
	}

	if _, ok := goval.(interface {
		RenderLayout(cr *cairo.Context, x, y float64, layout *pango.Layout)
	}); ok {
		o := pclass.StructFieldOffset("render_layout")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_layout)
	}

	if _, ok := goval.(interface {
		RenderLine(cr *cairo.Context, x0, y0, x1, y1 float64)
	}); ok {
		o := pclass.StructFieldOffset("render_line")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_line)
	}

	if _, ok := goval.(interface {
		RenderOption(cr *cairo.Context, x, y, width, height float64)
	}); ok {
		o := pclass.StructFieldOffset("render_option")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_ThemingEngineClass_render_option)
	}
}

//export _gotk4_gtk3_ThemingEngineClass_render_activity
func _gotk4_gtk3_ThemingEngineClass_render_activity(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderActivity(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderActivity(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_arrow
func _gotk4_gtk3_ThemingEngineClass_render_arrow(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderArrow(cr *cairo.Context, angle, x, y, size float64)
	})

	var _cr *cairo.Context // out
	var _angle float64     // out
	var _x float64         // out
	var _y float64         // out
	var _size float64      // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_angle = float64(arg2)
	_x = float64(arg3)
	_y = float64(arg4)
	_size = float64(arg5)

	iface.RenderArrow(_cr, _angle, _x, _y, _size)
}

//export _gotk4_gtk3_ThemingEngineClass_render_background
func _gotk4_gtk3_ThemingEngineClass_render_background(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderBackground(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderBackground(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_check
func _gotk4_gtk3_ThemingEngineClass_render_check(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderCheck(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderCheck(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_expander
func _gotk4_gtk3_ThemingEngineClass_render_expander(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderExpander(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderExpander(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_focus
func _gotk4_gtk3_ThemingEngineClass_render_focus(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderFocus(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderFocus(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_frame
func _gotk4_gtk3_ThemingEngineClass_render_frame(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderFrame(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderFrame(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_handle
func _gotk4_gtk3_ThemingEngineClass_render_handle(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderHandle(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderHandle(_cr, _x, _y, _width, _height)
}

//export _gotk4_gtk3_ThemingEngineClass_render_icon
func _gotk4_gtk3_ThemingEngineClass_render_icon(arg0 *C.void, arg1 *C.void, arg2 *C.GdkPixbuf, arg3 C.gdouble, arg4 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64)
	})

	var _cr *cairo.Context        // out
	var _pixbuf *gdkpixbuf.Pixbuf // out
	var _x float64                // out
	var _y float64                // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	{
		obj := coreglib.Take(unsafe.Pointer(arg2))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}
	_x = float64(arg3)
	_y = float64(arg4)

	iface.RenderIcon(_cr, _pixbuf, _x, _y)
}

//export _gotk4_gtk3_ThemingEngineClass_render_icon_surface
func _gotk4_gtk3_ThemingEngineClass_render_icon_surface(arg0 *C.void, arg1 *C.void, arg2 *C.void, arg3 C.gdouble, arg4 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderIconSurface(cr *cairo.Context, surface *cairo.Surface, x, y float64)
	})

	var _cr *cairo.Context      // out
	var _surface *cairo.Surface // out
	var _x float64              // out
	var _y float64              // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_surface = cairo.WrapSurface(uintptr(unsafe.Pointer(arg2)))
	C.cairo_surface_reference(arg2)
	runtime.SetFinalizer(_surface, func(v *cairo.Surface) {
		C.cairo_surface_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg3)
	_y = float64(arg4)

	iface.RenderIconSurface(_cr, _surface, _x, _y)
}

//export _gotk4_gtk3_ThemingEngineClass_render_layout
func _gotk4_gtk3_ThemingEngineClass_render_layout(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderLayout(cr *cairo.Context, x, y float64, layout *pango.Layout)
	})

	var _cr *cairo.Context    // out
	var _x float64            // out
	var _y float64            // out
	var _layout *pango.Layout // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	{
		obj := coreglib.Take(unsafe.Pointer(arg4))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	iface.RenderLayout(_cr, _x, _y, _layout)
}

//export _gotk4_gtk3_ThemingEngineClass_render_line
func _gotk4_gtk3_ThemingEngineClass_render_line(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderLine(cr *cairo.Context, x0, y0, x1, y1 float64)
	})

	var _cr *cairo.Context // out
	var _x0 float64        // out
	var _y0 float64        // out
	var _x1 float64        // out
	var _y1 float64        // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x0 = float64(arg2)
	_y0 = float64(arg3)
	_x1 = float64(arg4)
	_y1 = float64(arg5)

	iface.RenderLine(_cr, _x0, _y0, _x1, _y1)
}

//export _gotk4_gtk3_ThemingEngineClass_render_option
func _gotk4_gtk3_ThemingEngineClass_render_option(arg0 *C.void, arg1 *C.void, arg2 C.gdouble, arg3 C.gdouble, arg4 C.gdouble, arg5 C.gdouble) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		RenderOption(cr *cairo.Context, x, y, width, height float64)
	})

	var _cr *cairo.Context // out
	var _x float64         // out
	var _y float64         // out
	var _width float64     // out
	var _height float64    // out

	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg1)))
	C.cairo_reference(arg1)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.void)(unsafe.Pointer(v.Native())))
	})
	_x = float64(arg2)
	_y = float64(arg3)
	_width = float64(arg4)
	_height = float64(arg5)

	iface.RenderOption(_cr, _x, _y, _width, _height)
}

func wrapThemingEngine(obj *coreglib.Object) *ThemingEngine {
	return &ThemingEngine{
		Object: obj,
	}
}

func marshalThemingEngine(p uintptr) (interface{}, error) {
	return wrapThemingEngine(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Path returns the widget path used for style matching.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - widgetPath: WidgetPath.
//
func (engine *ThemingEngine) Path() *WidgetPath {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_gret := _info.InvokeClassMethod("get_path", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(engine)

	var _widgetPath *WidgetPath // out

	_widgetPath = (*WidgetPath)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	C.gtk_widget_path_ref(*(**C.void)(unsafe.Pointer(&_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_widgetPath)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("Gtk", "WidgetPath").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _widgetPath
}

// Screen returns the Screen to which engine currently rendering to.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - screen (optional) or NULL.
//
func (engine *ThemingEngine) Screen() *gdk.Screen {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_gret := _info.InvokeClassMethod("get_screen", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(engine)

	var _screen *gdk.Screen // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret))))
			_screen = &gdk.Screen{
				Object: obj,
			}
		}
	}

	return _screen
}

// StyleProperty gets the value for a widget style property.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - propertyName: name of the widget style property.
//
// The function returns the following values:
//
//    - value: return location for the property value, free with g_value_unset()
//      after use.
//
func (engine *ThemingEngine) StyleProperty(propertyName string) coreglib.Value {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_info.InvokeClassMethod("get_style_property", _args[:], _outs[:])

	runtime.KeepAlive(engine)
	runtime.KeepAlive(propertyName)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer(*(**C.GValue)(unsafe.Pointer(&_outs[0]))))

	return _value
}

// HasClass returns TRUE if the currently rendered contents have defined the
// given class name.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - styleClass class name to look up.
//
// The function returns the following values:
//
//    - ok: TRUE if engine has class_name defined.
//
func (engine *ThemingEngine) HasClass(styleClass string) bool {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(styleClass)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_gret := _info.InvokeClassMethod("has_class", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(engine)
	runtime.KeepAlive(styleClass)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HasRegion returns TRUE if the currently rendered contents have the region
// defined. If flags_return is not NULL, it is set to the flags affecting the
// region.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - styleRegion: region name.
//
// The function returns the following values:
//
//    - flags (optional): return location for region flags.
//    - ok: TRUE if region is defined.
//
func (engine *ThemingEngine) HasRegion(styleRegion string) (RegionFlags, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(styleRegion)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_gret := _info.InvokeClassMethod("has_region", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(engine)
	runtime.KeepAlive(styleRegion)

	var _flags RegionFlags // out
	var _ok bool           // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_flags = *(*RegionFlags)(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0]))))
	}
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _flags, _ok
}

// LookupColor looks up and resolves a color name in the current style’s color
// map.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - colorName: color name to lookup.
//
// The function returns the following values:
//
//    - color: return location for the looked up color.
//    - ok: TRUE if color_name was found and resolved, FALSE otherwise.
//
func (engine *ThemingEngine) LookupColor(colorName string) (*gdk.RGBA, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(engine).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(colorName)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ThemingEngine")
	_gret := _info.InvokeClassMethod("lookup_color", _args[:], _outs[:])
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(engine)
	runtime.KeepAlive(colorName)

	var _color *gdk.RGBA // out
	var _ok bool         // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0])))))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _color, _ok
}

// ThemingEngineLoad loads and initializes a theming engine module from the
// standard directories.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - name: theme engine name to load.
//
// The function returns the following values:
//
//    - themingEngine (optional): theming engine, or NULL if the engine name
//      doesn’t exist.
//
func ThemingEngineLoad(name string) *ThemingEngine {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "load")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)

	var _themingEngine *ThemingEngine // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_themingEngine = wrapThemingEngine(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _themingEngine
}
