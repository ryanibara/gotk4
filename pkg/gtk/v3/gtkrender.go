// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// RenderActivity renders an activity indicator (such as in Spinner). The state
// GTK_STATE_FLAG_CHECKED determines whether there is activity going on.
func RenderActivity(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_activity(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderArrow renders an arrow pointing to angle.
//
// Typical arrow rendering at 0, 1⁄2 π;, π; and 3⁄2 π:
//
// ! (arrows.png)
func RenderArrow(context *StyleContext, cr *cairo.Context, angle float64, x float64, y float64, size float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(angle)
	_arg4 = C.gdouble(x)
	_arg5 = C.gdouble(y)
	_arg6 = C.gdouble(size)

	C.gtk_render_arrow(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderBackground renders the background of an element.
//
// Typical background rendering, showing the effect of background-image,
// border-width and border-radius:
//
// ! (background.png)
func RenderBackground(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_background(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderBackgroundGetClip returns the area that will be affected (i.e. drawn
// to) when calling gtk_render_background() for the given context and rectangle.
func RenderBackgroundGetClip(context *StyleContext, x float64, y float64, width float64, height float64) gdk.Rectangle {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 C.gdouble          // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _outClip gdk.Rectangle

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = C.gdouble(x)
	_arg3 = C.gdouble(y)
	_arg4 = C.gdouble(width)
	_arg5 = C.gdouble(height)

	C.gtk_render_background_get_clip(_arg1, _arg2, _arg3, _arg4, _arg5, (*C.GdkRectangle)(unsafe.Pointer(&_outClip)))

	return _outClip
}

// RenderCheck renders a checkmark (as in a CheckButton).
//
// The GTK_STATE_FLAG_CHECKED state determines whether the check is on or off,
// and GTK_STATE_FLAG_INCONSISTENT determines whether it should be marked as
// undefined.
//
// Typical checkmark rendering:
//
// ! (checks.png)
func RenderCheck(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_check(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderExpander renders an expander (as used in TreeView and Expander) in the
// area defined by x, y, width, height. The state GTK_STATE_FLAG_CHECKED
// determines whether the expander is collapsed or expanded.
//
// Typical expander rendering:
//
// ! (expanders.png)
func RenderExpander(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_expander(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderExtension renders a extension (as in a Notebook tab) in the rectangle
// defined by x, y, width, height. The side where the extension connects to is
// defined by gap_side.
//
// Typical extension rendering:
//
// ! (extensions.png)
func RenderExtension(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkPositionType  // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkPositionType(gapSide)

	C.gtk_render_extension(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// RenderFocus renders a focus indicator on the rectangle determined by x, y,
// width, height.
//
// Typical focus rendering:
//
// ! (focus.png)
func RenderFocus(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_focus(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderFrame renders a frame around the rectangle defined by x, y, width,
// height.
//
// Examples of frame rendering, showing the effect of border-image,
// border-color, border-width, border-radius and junctions:
//
// ! (frames.png)
func RenderFrame(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_frame(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderFrameGap renders a frame around the rectangle defined by (x, y, width,
// height), leaving a gap on one side. xy0_gap and xy1_gap will mean X
// coordinates for GTK_POS_TOP and GTK_POS_BOTTOM gap sides, and Y coordinates
// for GTK_POS_LEFT and GTK_POS_RIGHT.
//
// Typical rendering of a frame with a gap:
//
// ! (frame-gap.png)
//
// Deprecated: Use gtk_render_frame() instead. Themes can create gaps by
// omitting borders via CSS.
func RenderFrameGap(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, gapSide PositionType, xy0Gap float64, xy1Gap float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkPositionType  // out
	var _arg8 C.gdouble          // out
	var _arg9 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkPositionType(gapSide)
	_arg8 = C.gdouble(xy0Gap)
	_arg9 = C.gdouble(xy1Gap)

	C.gtk_render_frame_gap(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7, _arg8, _arg9)
}

// RenderHandle renders a handle (as in HandleBox, Paned and Window’s resize
// grip), in the rectangle determined by x, y, width, height.
//
// Handles rendered for the paned and grip classes:
//
// ! (handles.png)
func RenderHandle(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_handle(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderIcon renders the icon in pixbuf at the specified x and y coordinates.
//
// This function will render the icon in pixbuf at exactly its size, regardless
// of scaling factors, which may not be appropriate when drawing on displays
// with high pixel densities.
//
// You probably want to use gtk_render_icon_surface() instead, if you already
// have a Cairo surface.
func RenderIcon(context *StyleContext, cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x float64, y float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 *C.GdkPixbuf       // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))
	_arg4 = C.gdouble(x)
	_arg5 = C.gdouble(y)

	C.gtk_render_icon(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderIconPixbuf renders the icon specified by source at the given size,
// returning the result in a pixbuf.
//
// Deprecated: Use gtk_icon_theme_load_icon() instead.
func RenderIconPixbuf(context *StyleContext, source *IconSource, size int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.GtkIconSource   // out
	var _arg3 C.GtkIconSize      // out
	var _cret *C.GdkPixbuf       // in

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkIconSource)(unsafe.Pointer(source))
	_arg3 = C.GtkIconSize(size)

	_cret = C.gtk_render_icon_pixbuf(_arg1, _arg2, _arg3)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			Icon: gio.Icon{
				Object: obj,
			},
		}
	}

	return _pixbuf
}

// RenderIconSurface renders the icon in surface at the specified x and y
// coordinates.
func RenderIconSurface(context *StyleContext, cr *cairo.Context, surface *cairo.Surface, x float64, y float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 *C.cairo_surface_t // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = (*C.cairo_surface_t)(unsafe.Pointer(surface))
	_arg4 = C.gdouble(x)
	_arg5 = C.gdouble(y)

	C.gtk_render_icon_surface(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderLayout renders layout on the coordinates x, y
func RenderLayout(context *StyleContext, cr *cairo.Context, x float64, y float64, layout *pango.Layout) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 *C.PangoLayout     // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))

	C.gtk_render_layout(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// RenderLine renders a line from (x0, y0) to (x1, y1).
func RenderLine(context *StyleContext, cr *cairo.Context, x0 float64, y0 float64, x1 float64, y1 float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x0)
	_arg4 = C.gdouble(y0)
	_arg5 = C.gdouble(x1)
	_arg6 = C.gdouble(y1)

	C.gtk_render_line(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderOption renders an option mark (as in a RadioButton), the
// GTK_STATE_FLAG_CHECKED state will determine whether the option is on or off,
// and GTK_STATE_FLAG_INCONSISTENT whether it should be marked as undefined.
//
// Typical option mark rendering:
//
// ! (options.png)
func RenderOption(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)

	C.gtk_render_option(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// RenderSlider renders a slider (as in Scale) in the rectangle defined by x, y,
// width, height. orientation defines whether the slider is vertical or
// horizontal.
//
// Typical slider rendering:
//
// ! (sliders.png)
func RenderSlider(context *StyleContext, cr *cairo.Context, x float64, y float64, width float64, height float64, orientation Orientation) {
	var _arg1 *C.GtkStyleContext // out
	var _arg2 *C.cairo_t         // out
	var _arg3 C.gdouble          // out
	var _arg4 C.gdouble          // out
	var _arg5 C.gdouble          // out
	var _arg6 C.gdouble          // out
	var _arg7 C.GtkOrientation   // out

	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.cairo_t)(unsafe.Pointer(cr))
	_arg3 = C.gdouble(x)
	_arg4 = C.gdouble(y)
	_arg5 = C.gdouble(width)
	_arg6 = C.gdouble(height)
	_arg7 = C.GtkOrientation(orientation)

	C.gtk_render_slider(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}
