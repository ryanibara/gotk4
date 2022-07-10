// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeCornerType returns the GType for the type CornerType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCornerType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "CornerType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCornerType)
	return gtype
}

// GTypePolicyType returns the GType for the type PolicyType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePolicyType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PolicyType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPolicyType)
	return gtype
}

// GTypeScrolledWindow returns the GType for the type ScrolledWindow.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeScrolledWindow() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ScrolledWindow").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalScrolledWindow)
	return gtype
}

// CornerType specifies which corner a child widget should be placed in when
// packed into a ScrolledWindow. This is effectively the opposite of where the
// scroll bars are placed.
type CornerType C.gint

const (
	// CornerTopLeft: place the scrollbars on the right and bottom of the widget
	// (default behaviour).
	CornerTopLeft CornerType = iota
	// CornerBottomLeft: place the scrollbars on the top and right of the
	// widget.
	CornerBottomLeft
	// CornerTopRight: place the scrollbars on the left and bottom of the
	// widget.
	CornerTopRight
	// CornerBottomRight: place the scrollbars on the top and left of the
	// widget.
	CornerBottomRight
)

func marshalCornerType(p uintptr) (interface{}, error) {
	return CornerType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CornerType.
func (c CornerType) String() string {
	switch c {
	case CornerTopLeft:
		return "TopLeft"
	case CornerBottomLeft:
		return "BottomLeft"
	case CornerTopRight:
		return "TopRight"
	case CornerBottomRight:
		return "BottomRight"
	default:
		return fmt.Sprintf("CornerType(%d)", c)
	}
}

// PolicyType determines how the size should be computed to achieve the one of
// the visibility mode for the scrollbars.
type PolicyType C.gint

const (
	// PolicyAlways: scrollbar is always visible. The view size is independent
	// of the content.
	PolicyAlways PolicyType = iota
	// PolicyAutomatic: scrollbar will appear and disappear as necessary. For
	// example, when all of a TreeView can not be seen.
	PolicyAutomatic
	// PolicyNever: scrollbar should never appear. In this mode the content
	// determines the size.
	PolicyNever
	// PolicyExternal: don't show a scrollbar, but don't force the size to
	// follow the content. This can be used e.g. to make multiple scrolled
	// windows share a scrollbar. Since: 3.16.
	PolicyExternal
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	return PolicyType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PolicyType.
func (p PolicyType) String() string {
	switch p {
	case PolicyAlways:
		return "Always"
	case PolicyAutomatic:
		return "Automatic"
	case PolicyNever:
		return "Never"
	case PolicyExternal:
		return "External"
	default:
		return fmt.Sprintf("PolicyType(%d)", p)
	}
}

// ScrolledWindowOverrider contains methods that are overridable.
type ScrolledWindowOverrider interface {
}

// ScrolledWindow is a container that accepts a single child widget, makes that
// child scrollable using either internally added scrollbars or externally
// associated adjustments, and optionally draws a frame around the child.
//
// Widgets with native scrolling support, i.e. those whose classes implement the
// Scrollable interface, are added directly. For other types of widget, the
// class Viewport acts as an adaptor, giving scrollability to other widgets.
// GtkScrolledWindow’s implementation of gtk_container_add() intelligently
// accounts for whether or not the added child is a Scrollable. If it isn’t,
// ScrolledWindow wraps the child in a Viewport and adds that for you.
// Therefore, you can just add any child widget and not worry about the details.
//
// If gtk_container_add() has added a Viewport for you, you can remove both your
// added child widget from the Viewport, and the Viewport from the
// GtkScrolledWindow, like this:
//
//    GtkWidget *scrolled_window = gtk_scrolled_window_new (NULL, NULL);
//    GtkWidget *child_widget = gtk_button_new ();
//
//    // GtkButton is not a GtkScrollable, so GtkScrolledWindow will automatically
//    // add a GtkViewport.
//    gtk_container_add (GTK_CONTAINER (scrolled_window),
//                       child_widget);
//
//    // Either of these will result in child_widget being unparented:
//    gtk_container_remove (GTK_CONTAINER (scrolled_window),
//                          child_widget);
//    // or
//    gtk_container_remove (GTK_CONTAINER (scrolled_window),
//                          gtk_bin_get_child (GTK_BIN (scrolled_window)));
//
// Unless ScrolledWindow:policy is GTK_POLICY_NEVER or GTK_POLICY_EXTERNAL,
// GtkScrolledWindow adds internal Scrollbar widgets around its child. The
// scroll position of the child, and if applicable the scrollbars, is controlled
// by the ScrolledWindow:hadjustment and ScrolledWindow:vadjustment that are
// associated with the GtkScrolledWindow. See the docs on Scrollbar for the
// details, but note that the “step_increment” and “page_increment” fields are
// only effective if the policy causes scrollbars to be present.
//
// If a GtkScrolledWindow doesn’t behave quite as you would like, or doesn’t
// have exactly the right layout, it’s very possible to set up your own
// scrolling with Scrollbar and for example a Grid.
//
//
// Touch support
//
// GtkScrolledWindow has built-in support for touch devices. When a touchscreen
// is used, swiping will move the scrolled window, and will expose 'kinetic'
// behavior. This can be turned off with the ScrolledWindow:kinetic-scrolling
// property if it is undesired.
//
// GtkScrolledWindow also displays visual 'overshoot' indication when the
// content is pulled beyond the end, and this situation can be captured with the
// ScrolledWindow::edge-overshot signal.
//
// If no mouse device is present, the scrollbars will overlayed as narrow,
// auto-hiding indicators over the content. If traditional scrollbars are
// desired although no mouse is present, this behaviour can be turned off with
// the ScrolledWindow:overlay-scrolling property.
//
//
// CSS nodes
//
// GtkScrolledWindow has a main CSS node with name scrolledwindow.
//
// It uses subnodes with names overshoot and undershoot to draw the overflow and
// underflow indications. These nodes get the .left, .right, .top or .bottom
// style class added depending on where the indication is drawn.
//
// GtkScrolledWindow also sets the positional style classes (.left, .right,
// .top, .bottom) and style classes related to overlay scrolling
// (.overlay-indicator, .dragging, .hovering) on its scrollbars.
//
// If both scrollbars are visible, the area where they meet is drawn with a
// subnode named junction.
type ScrolledWindow struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*ScrolledWindow)(nil)
)

func classInitScrolledWindower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapScrolledWindow(obj *coreglib.Object) *ScrolledWindow {
	return &ScrolledWindow{
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
	}
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	return wrapScrolledWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewScrolledWindow creates a new scrolled window.
//
// The two arguments are the scrolled window’s adjustments; these will be shared
// with the scrollbars and the child widget to keep the bars in sync with the
// child. Usually you want to pass NULL for the adjustments, which will cause
// the scrolled window to create them for you.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): horizontal adjustment.
//    - vadjustment (optional): vertical adjustment.
//
// The function returns the following values:
//
//    - scrolledWindow: new scrolled window.
//
func NewScrolledWindow(hadjustment, vadjustment *Adjustment) *ScrolledWindow {
	var _args [2]girepository.Argument

	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("new_ScrolledWindow", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(hadjustment)
	runtime.KeepAlive(vadjustment)

	var _scrolledWindow *ScrolledWindow // out

	_scrolledWindow = wrapScrolledWindow(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _scrolledWindow
}

// AddWithViewport: used to add children without native scrolling capabilities.
// This is simply a convenience function; it is equivalent to adding the
// unscrollable child to a viewport, then adding the viewport to the scrolled
// window. If a child has native scrolling, use gtk_container_add() instead of
// this function.
//
// The viewport scrolls the child by moving its Window, and takes the size of
// the child to be the size of its toplevel Window. This will be very wrong for
// most widgets that support native scrolling; for example, if you add a widget
// such as TreeView with a viewport, the whole widget will scroll, including the
// column headings. Thus, widgets with native scrolling support should not be
// used with the Viewport proxy.
//
// A widget supports scrolling natively if it implements the Scrollable
// interface.
//
// Deprecated: gtk_container_add() will automatically add a Viewport if the
// child doesn’t implement Scrollable.
//
// The function takes the following parameters:
//
//    - child: widget you want to scroll.
//
func (scrolledWindow *ScrolledWindow) AddWithViewport(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("add_with_viewport", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(child)
}

// CaptureButtonPress: return whether button presses are captured during kinetic
// scrolling. See gtk_scrolled_window_set_capture_button_press().
//
// The function returns the following values:
//
//    - ok: TRUE if button presses are captured during kinetic scrolling.
//
func (scrolledWindow *ScrolledWindow) CaptureButtonPress() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_capture_button_press", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// HAdjustment returns the horizontal scrollbar’s adjustment, used to connect
// the horizontal scrollbar to the child widget’s horizontal scroll
// functionality.
//
// The function returns the following values:
//
//    - adjustment: horizontal Adjustment.
//
func (scrolledWindow *ScrolledWindow) HAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_hadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// HScrollbar returns the horizontal scrollbar of scrolled_window.
//
// The function returns the following values:
//
//    - widget: horizontal scrollbar of the scrolled window.
//
func (scrolledWindow *ScrolledWindow) HScrollbar() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_hscrollbar", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// KineticScrolling returns the specified kinetic scrolling behavior.
//
// The function returns the following values:
//
//    - ok: scrolling behavior flags.
//
func (scrolledWindow *ScrolledWindow) KineticScrolling() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_kinetic_scrolling", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MaxContentHeight returns the maximum content height set.
//
// The function returns the following values:
//
//    - gint: maximum content height, or -1.
//
func (scrolledWindow *ScrolledWindow) MaxContentHeight() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_max_content_height", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MaxContentWidth returns the maximum content width set.
//
// The function returns the following values:
//
//    - gint: maximum content width, or -1.
//
func (scrolledWindow *ScrolledWindow) MaxContentWidth() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_max_content_width", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MinContentHeight gets the minimal content height of scrolled_window, or -1 if
// not set.
//
// The function returns the following values:
//
//    - gint: minimal content height.
//
func (scrolledWindow *ScrolledWindow) MinContentHeight() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_min_content_height", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// MinContentWidth gets the minimum content width of scrolled_window, or -1 if
// not set.
//
// The function returns the following values:
//
//    - gint: minimum content width.
//
func (scrolledWindow *ScrolledWindow) MinContentWidth() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_min_content_width", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// OverlayScrolling returns whether overlay scrolling is enabled for this
// scrolled window.
//
// The function returns the following values:
//
//    - ok: TRUE if overlay scrolling is enabled.
//
func (scrolledWindow *ScrolledWindow) OverlayScrolling() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_overlay_scrolling", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Policy retrieves the current policy values for the horizontal and vertical
// scrollbars. See gtk_scrolled_window_set_policy().
//
// The function returns the following values:
//
//    - hscrollbarPolicy (optional): location to store the policy for the
//      horizontal scrollbar, or NULL.
//    - vscrollbarPolicy (optional): location to store the policy for the
//      vertical scrollbar, or NULL.
//
func (scrolledWindow *ScrolledWindow) Policy() (hscrollbarPolicy, vscrollbarPolicy PolicyType) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("get_policy", _args[:], _outs[:])

	runtime.KeepAlive(scrolledWindow)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_hscrollbarPolicy = *(*PolicyType)(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[0]))))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_vscrollbarPolicy = *(*PolicyType)(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_outs[1]))))
	}

	return _hscrollbarPolicy, _vscrollbarPolicy
}

// PropagateNaturalHeight reports whether the natural height of the child will
// be calculated and propagated through the scrolled window’s requested natural
// height.
//
// The function returns the following values:
//
//    - ok: whether natural height propagation is enabled.
//
func (scrolledWindow *ScrolledWindow) PropagateNaturalHeight() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_propagate_natural_height", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PropagateNaturalWidth reports whether the natural width of the child will be
// calculated and propagated through the scrolled window’s requested natural
// width.
//
// The function returns the following values:
//
//    - ok: whether natural width propagation is enabled.
//
func (scrolledWindow *ScrolledWindow) PropagateNaturalWidth() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_propagate_natural_width", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VAdjustment returns the vertical scrollbar’s adjustment, used to connect the
// vertical scrollbar to the child widget’s vertical scroll functionality.
//
// The function returns the following values:
//
//    - adjustment: vertical Adjustment.
//
func (scrolledWindow *ScrolledWindow) VAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_vadjustment", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _adjustment
}

// VScrollbar returns the vertical scrollbar of scrolled_window.
//
// The function returns the following values:
//
//    - widget: vertical scrollbar of the scrolled window.
//
func (scrolledWindow *ScrolledWindow) VScrollbar() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_gret := _info.InvokeClassMethod("get_vscrollbar", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetCaptureButtonPress changes the behaviour of scrolled_window with regard to
// the initial event that possibly starts kinetic scrolling. When
// capture_button_press is set to TRUE, the event is captured by the scrolled
// window, and then later replayed if it is meant to go to the child widget.
//
// This should be enabled if any child widgets perform non-reversible actions on
// Widget::button-press-event. If they don't, and handle additionally handle
// Widget::grab-broken-event, it might be better to set capture_button_press to
// FALSE.
//
// This setting only has an effect if kinetic scrolling is enabled.
//
// The function takes the following parameters:
//
//    - captureButtonPress: TRUE to capture button presses.
//
func (scrolledWindow *ScrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if captureButtonPress {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_capture_button_press", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(captureButtonPress)
}

// SetHAdjustment sets the Adjustment for the horizontal scrollbar.
//
// The function takes the following parameters:
//
//    - hadjustment (optional) to use, or NULL to create a new one.
//
func (scrolledWindow *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_hadjustment", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hadjustment)
}

// SetKineticScrolling turns kinetic scrolling on or off. Kinetic scrolling only
// applies to devices with source GDK_SOURCE_TOUCHSCREEN.
//
// The function takes the following parameters:
//
//    - kineticScrolling: TRUE to enable kinetic scrolling.
//
func (scrolledWindow *ScrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if kineticScrolling {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_kinetic_scrolling", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(kineticScrolling)
}

// SetMaxContentHeight sets the maximum height that scrolled_window should keep
// visible. The scrolled_window will grow up to this height before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than ScrolledWindow:min-content-height.
//
// The function takes the following parameters:
//
//    - height: maximum content height.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentHeight(height int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(height)

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_max_content_height", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMaxContentWidth sets the maximum width that scrolled_window should keep
// visible. The scrolled_window will grow up to this width before it starts
// scrolling the content.
//
// It is a programming error to set the maximum content width to a value smaller
// than ScrolledWindow:min-content-width.
//
// The function takes the following parameters:
//
//    - width: maximum content width.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentWidth(width int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(width)

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_max_content_width", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(width)
}

// SetMinContentHeight sets the minimum height that scrolled_window should keep
// visible. Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content height to a value
// greater than ScrolledWindow:max-content-height.
//
// The function takes the following parameters:
//
//    - height: minimal content height.
//
func (scrolledWindow *ScrolledWindow) SetMinContentHeight(height int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(height)

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_min_content_height", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMinContentWidth sets the minimum width that scrolled_window should keep
// visible. Note that this can and (usually will) be smaller than the minimum
// size of the content.
//
// It is a programming error to set the minimum content width to a value greater
// than ScrolledWindow:max-content-width.
//
// The function takes the following parameters:
//
//    - width: minimal content width.
//
func (scrolledWindow *ScrolledWindow) SetMinContentWidth(width int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(width)

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_min_content_width", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(width)
}

// SetOverlayScrolling enables or disables overlay scrolling for this scrolled
// window.
//
// The function takes the following parameters:
//
//    - overlayScrolling: whether to enable overlay scrolling.
//
func (scrolledWindow *ScrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if overlayScrolling {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_overlay_scrolling", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(overlayScrolling)
}

// SetPropagateNaturalHeight sets whether the natural height of the child should
// be calculated and propagated through the scrolled window’s requested natural
// height.
//
// The function takes the following parameters:
//
//    - propagate: whether to propagate natural height.
//
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if propagate {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_propagate_natural_height", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(propagate)
}

// SetPropagateNaturalWidth sets whether the natural width of the child should
// be calculated and propagated through the scrolled window’s requested natural
// width.
//
// The function takes the following parameters:
//
//    - propagate: whether to propagate natural width.
//
func (scrolledWindow *ScrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if propagate {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_propagate_natural_width", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(propagate)
}

// SetVAdjustment sets the Adjustment for the vertical scrollbar.
//
// The function takes the following parameters:
//
//    - vadjustment (optional) to use, or NULL to create a new one.
//
func (scrolledWindow *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("set_vadjustment", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(vadjustment)
}

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars for the scrolled window. If no window placement is set for a
// scrolled window, it defaults to GTK_CORNER_TOP_LEFT.
//
// See also gtk_scrolled_window_set_placement() and
// gtk_scrolled_window_get_placement().
func (scrolledWindow *ScrolledWindow) UnsetPlacement() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_info := girepository.MustFind("Gtk", "ScrolledWindow")
	_info.InvokeClassMethod("unset_placement", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
}
