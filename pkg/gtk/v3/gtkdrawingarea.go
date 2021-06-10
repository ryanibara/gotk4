// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_drawing_area_get_type()), F: marshalDrawingArea},
	})
}

// DrawingArea: the DrawingArea widget is used for creating custom user
// interface elements. It’s essentially a blank widget; you can draw on it.
// After creating a drawing area, the application may want to connect to:
//
// - Mouse and button press signals to respond to input from the user. (Use
// gtk_widget_add_events() to enable events you wish to receive.)
//
// - The Widget::realize signal to take any necessary actions when the widget is
// instantiated on a particular display. (Create GDK resources in response to
// this signal.)
//
// - The Widget::size-allocate signal to take any necessary actions when the
// widget changes size.
//
// - The Widget::draw signal to handle redrawing the contents of the widget.
//
// The following code portion demonstrates using a drawing area to display a
// circle in the normal widget foreground color.
//
// Note that GDK automatically clears the exposed area before sending the expose
// event, and that drawing is implicitly clipped to the exposed area. If you
// want to have a theme-provided background, you need to call
// gtk_render_background() in your ::draw method.
//
// Simple GtkDrawingArea usage
//
//    gboolean
//    draw_callback (GtkWidget *widget, cairo_t *cr, gpointer data)
//    {
//      guint width, height;
//      GdkRGBA color;
//      GtkStyleContext *context;
//
//      context = gtk_widget_get_style_context (widget);
//
//      width = gtk_widget_get_allocated_width (widget);
//      height = gtk_widget_get_allocated_height (widget);
//
//      gtk_render_background (context, cr, 0, 0, width, height);
//
//      cairo_arc (cr,
//                 width / 2.0, height / 2.0,
//                 MIN (width, height) / 2.0,
//                 0, 2 * G_PI);
//
//      gtk_style_context_get_color (context,
//                                   gtk_style_context_get_state (context),
//                                   &color);
//      gdk_cairo_set_source_rgba (cr, &color);
//
//      cairo_fill (cr);
//
//     return FALSE;
//    }
//    [...]
//      GtkWidget *drawing_area = gtk_drawing_area_new ();
//      gtk_widget_set_size_request (drawing_area, 100, 100);
//      g_signal_connect (G_OBJECT (drawing_area), "draw",
//                        G_CALLBACK (draw_callback), NULL);
//
// Draw signals are normally delivered when a drawing area first comes onscreen,
// or when it’s covered by another window and then uncovered. You can also force
// an expose event by adding to the “damage region” of the drawing area’s
// window; gtk_widget_queue_draw_area() and gdk_window_invalidate_rect() are
// equally good ways to do this. You’ll then get a draw signal for the invalid
// region.
//
// The available routines for drawing are documented on the [GDK Drawing
// Primitives][gdk3-Cairo-Interaction] page and the cairo documentation.
//
// To receive mouse events on a drawing area, you will need to enable them with
// gtk_widget_add_events(). To receive keyboard events, you will need to set the
// “can-focus” property on the drawing area, and you should probably draw some
// user-visible indication that the drawing area is focused. Use
// gtk_widget_has_focus() in your expose event handler to decide whether to draw
// the focus indicator. See gtk_render_focus() for one way to draw focus.
type DrawingArea interface {
	Widget
	Buildable
}

// drawingArea implements the DrawingArea interface.
type drawingArea struct {
	Widget
	Buildable
}

var _ DrawingArea = (*drawingArea)(nil)

// WrapDrawingArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawingArea(obj *externglib.Object) DrawingArea {
	return DrawingArea{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalDrawingArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawingArea(obj), nil
}
