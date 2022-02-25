// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_DrawingAreaClass_resize(GtkDrawingArea*, int, int);
// extern void _gotk4_gtk4_DrawingAreaDrawFunc(GtkDrawingArea*, cairo_t*, int, int, gpointer);
// extern void _gotk4_gtk4_DrawingArea_ConnectResize(gpointer, gint, gint, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkdrawingarea.go.
var GTypeDrawingArea = externglib.Type(C.gtk_drawing_area_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDrawingArea, F: marshalDrawingArea},
	})
}

// DrawingAreaDrawFunc: whenever drawing_area needs to redraw, this function
// will be called.
//
// This function should exclusively redraw the contents of the drawing area and
// must not call any widget functions that cause changes.
type DrawingAreaDrawFunc func(drawingArea *DrawingArea, cr *cairo.Context, width, height int)

//export _gotk4_gtk4_DrawingAreaDrawFunc
func _gotk4_gtk4_DrawingAreaDrawFunc(arg1 *C.GtkDrawingArea, arg2 *C.cairo_t, arg3 C.int, arg4 C.int, arg5 C.gpointer) {
	var fn DrawingAreaDrawFunc
	{
		v := gbox.Get(uintptr(arg5))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(DrawingAreaDrawFunc)
	}

	var _drawingArea *DrawingArea // out
	var _cr *cairo.Context        // out
	var _width int                // out
	var _height int               // out

	_drawingArea = wrapDrawingArea(externglib.Take(unsafe.Pointer(arg1)))
	_cr = cairo.WrapContext(uintptr(unsafe.Pointer(arg2)))
	C.cairo_reference(arg2)
	runtime.SetFinalizer(_cr, func(v *cairo.Context) {
		C.cairo_destroy((*C.cairo_t)(unsafe.Pointer(v.Native())))
	})
	_width = int(arg3)
	_height = int(arg4)

	fn(_drawingArea, _cr, _width, _height)
}

// DrawingAreaOverrider contains methods that are overridable.
type DrawingAreaOverrider interface {
	// The function takes the following parameters:
	//
	//    - width
	//    - height
	//
	Resize(width, height int)
}

// DrawingArea: GtkDrawingArea is a widget that allows drawing with cairo.
//
// !An example GtkDrawingArea (drawingarea.png)
//
// It’s essentially a blank widget; you can draw on it. After creating a drawing
// area, the application may want to connect to:
//
// - The gtk.Widget::realize signal to take any necessary actions when the
// widget is instantiated on a particular display. (Create GDK resources in
// response to this signal.)
//
// - The gtk.DrawingArea::resize signal to take any necessary actions when the
// widget changes size.
//
// - Call gtk.DrawingArea.SetDrawFunc() to handle redrawing the contents of the
// widget.
//
// The following code portion demonstrates using a drawing area to display a
// circle in the normal widget foreground color.
//
// Simple GtkDrawingArea usage
//
//    static void
//    draw_function (GtkDrawingArea *area,
//                   cairo_t        *cr,
//                   int             width,
//                   int             height,
//                   gpointer        data)
//    {
//      GdkRGBA color;
//      GtkStyleContext *context;
//
//      context = gtk_widget_get_style_context (GTK_WIDGET (area));
//
//      cairo_arc (cr,
//                 width / 2.0, height / 2.0,
//                 MIN (width, height) / 2.0,
//                 0, 2 * G_PI);
//
//      gtk_style_context_get_color (context,
//                                   &color);
//      gdk_cairo_set_source_rgba (cr, &color);
//
//      cairo_fill (cr);
//    }
//
//    int
//    main (int argc, char **argv)
//    {
//      gtk_init ();
//
//      GtkWidget *area = gtk_drawing_area_new ();
//      gtk_drawing_area_set_content_width (GTK_DRAWING_AREA (area), 100);
//      gtk_drawing_area_set_content_height (GTK_DRAWING_AREA (area), 100);
//      gtk_drawing_area_set_draw_func (GTK_DRAWING_AREA (area),
//                                      draw_function,
//                                      NULL, NULL);
//      return 0;
//    }
//
//
// The draw function is normally called when a drawing area first comes
// onscreen, or when it’s covered by another window and then uncovered. You can
// also force a redraw by adding to the “damage region” of the drawing area’s
// window using gtk.Widget.QueueDraw(). This will cause the drawing area to call
// the draw function again.
//
// The available routines for drawing are documented on the [GDK Drawing
// Primitives][gdk4-Cairo-Interaction] page and the cairo documentation.
//
// To receive mouse events on a drawing area, you will need to use event
// controllers. To receive keyboard events, you will need to set the “can-focus”
// property on the drawing area, and you should probably draw some user-visible
// indication that the drawing area is focused.
//
// If you need more complex control over your widget, you should consider
// creating your own GtkWidget subclass.
type DrawingArea struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*DrawingArea)(nil)
)

func classInitDrawingAreaer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkDrawingAreaClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkDrawingAreaClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Resize(width, height int) }); ok {
		pclass.resize = (*[0]byte)(C._gotk4_gtk4_DrawingAreaClass_resize)
	}
}

//export _gotk4_gtk4_DrawingAreaClass_resize
func _gotk4_gtk4_DrawingAreaClass_resize(arg0 *C.GtkDrawingArea, arg1 C.int, arg2 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Resize(width, height int) })

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	iface.Resize(_width, _height)
}

func wrapDrawingArea(obj *externglib.Object) *DrawingArea {
	return &DrawingArea{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	return wrapDrawingArea(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_DrawingArea_ConnectResize
func _gotk4_gtk4_DrawingArea_ConnectResize(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.guintptr) {
	var f func(width, height int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(width, height int))
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	f(_width, _height)
}

// ConnectResize: emitted once when the widget is realized, and then each time
// the widget is changed while realized.
//
// This is useful in order to keep state up to date with the widget size, like
// for instance a backing surface.
func (self *DrawingArea) ConnectResize(f func(width, height int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "resize", false, unsafe.Pointer(C._gotk4_gtk4_DrawingArea_ConnectResize), f)
}

// NewDrawingArea creates a new drawing area.
//
// The function returns the following values:
//
//    - drawingArea: new GtkDrawingArea.
//
func NewDrawingArea() *DrawingArea {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_drawing_area_new()

	var _drawingArea *DrawingArea // out

	_drawingArea = wrapDrawingArea(externglib.Take(unsafe.Pointer(_cret)))

	return _drawingArea
}

// ContentHeight retrieves the content height of the GtkDrawingArea.
//
// The function returns the following values:
//
//    - gint: height requested for content of the drawing area.
//
func (self *DrawingArea) ContentHeight() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_drawing_area_get_content_height(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ContentWidth retrieves the content width of the GtkDrawingArea.
//
// The function returns the following values:
//
//    - gint: width requested for content of the drawing area.
//
func (self *DrawingArea) ContentWidth() int {
	var _arg0 *C.GtkDrawingArea // out
	var _cret C.int             // in

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_drawing_area_get_content_width(_arg0)
	runtime.KeepAlive(self)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetContentHeight sets the desired height of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual height passed to your draw function is larger
// than the height set here. You can use gtk.Widget.SetVAlign() to avoid that.
//
// If the height is set to 0 (the default), the drawing area may disappear.
//
// The function takes the following parameters:
//
//    - height of contents.
//
func (self *DrawingArea) SetContentHeight(height int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(height)

	C.gtk_drawing_area_set_content_height(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(height)
}

// SetContentWidth sets the desired width of the contents of the drawing area.
//
// Note that because widgets may be allocated larger sizes than they requested,
// it is possible that the actual width passed to your draw function is larger
// than the width set here. You can use gtk.Widget.SetHAlign() to avoid that.
//
// If the width is set to 0 (the default), the drawing area may disappear.
//
// The function takes the following parameters:
//
//    - width of contents.
//
func (self *DrawingArea) SetContentWidth(width int) {
	var _arg0 *C.GtkDrawingArea // out
	var _arg1 C.int             // out

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.int(width)

	C.gtk_drawing_area_set_content_width(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(width)
}

// SetDrawFunc: setting a draw function is the main thing you want to do when
// using a drawing area.
//
// The draw function is called whenever GTK needs to draw the contents of the
// drawing area to the screen.
//
// The draw function will be called during the drawing stage of GTK. In the
// drawing stage it is not allowed to change properties of any GTK widgets or
// call any functions that would cause any properties to be changed. You should
// restrict yourself exclusively to drawing your contents in the draw function.
//
// If what you are drawing does change, call gtk.Widget.QueueDraw() on the
// drawing area. This will cause a redraw and will call draw_func again.
//
// The function takes the following parameters:
//
//    - drawFunc (optional): callback that lets you draw the drawing area's
//      contents.
//
func (self *DrawingArea) SetDrawFunc(drawFunc DrawingAreaDrawFunc) {
	var _arg0 *C.GtkDrawingArea        // out
	var _arg1 C.GtkDrawingAreaDrawFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkDrawingArea)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if drawFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_DrawingAreaDrawFunc)
		_arg2 = C.gpointer(gbox.Assign(drawFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_drawing_area_set_draw_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(self)
	runtime.KeepAlive(drawFunc)
}
