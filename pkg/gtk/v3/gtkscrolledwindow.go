// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_corner_type_get_type()), F: marshalCornerType},
		{T: externglib.Type(C.gtk_policy_type_get_type()), F: marshalPolicyType},
		{T: externglib.Type(C.gtk_scrolled_window_get_type()), F: marshalScrolledWindow},
	})
}

// CornerType specifies which corner a child widget should be placed in when
// packed into a ScrolledWindow. This is effectively the opposite of where the
// scroll bars are placed.
type CornerType int

const (
	// CornerTypeTopLeft: place the scrollbars on the right and bottom of the
	// widget (default behaviour).
	CornerTypeTopLeft CornerType = 0
	// CornerTypeBottomLeft: place the scrollbars on the top and right of the
	// widget.
	CornerTypeBottomLeft CornerType = 1
	// CornerTypeTopRight: place the scrollbars on the left and bottom of the
	// widget.
	CornerTypeTopRight CornerType = 2
	// CornerTypeBottomRight: place the scrollbars on the top and left of the
	// widget.
	CornerTypeBottomRight CornerType = 3
)

func marshalCornerType(p uintptr) (interface{}, error) {
	return CornerType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PolicyType determines how the size should be computed to achieve the one of
// the visibility mode for the scrollbars.
type PolicyType int

const (
	// PolicyTypeAlways: the scrollbar is always visible. The view size is
	// independent of the content.
	PolicyTypeAlways PolicyType = 0
	// PolicyTypeAutomatic: the scrollbar will appear and disappear as
	// necessary. For example, when all of a TreeView can not be seen.
	PolicyTypeAutomatic PolicyType = 1
	// PolicyTypeNever: the scrollbar should never appear. In this mode the
	// content determines the size.
	PolicyTypeNever PolicyType = 2
	// PolicyTypeExternal: don't show a scrollbar, but don't force the size to
	// follow the content. This can be used e.g. to make multiple scrolled
	// windows share a scrollbar. Since: 3.16
	PolicyTypeExternal PolicyType = 3
)

func marshalPolicyType(p uintptr) (interface{}, error) {
	return PolicyType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrolledWindow: gtkScrolledWindow is a container that accepts a single child
// widget, makes that child scrollable using either internally added scrollbars
// or externally associated adjustments, and optionally draws a frame around the
// child.
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
type ScrolledWindow interface {
	Bin
	Buildable

	// AddWithViewport: used to add children without native scrolling
	// capabilities. This is simply a convenience function; it is equivalent to
	// adding the unscrollable child to a viewport, then adding the viewport to
	// the scrolled window. If a child has native scrolling, use
	// gtk_container_add() instead of this function.
	//
	// The viewport scrolls the child by moving its Window, and takes the size
	// of the child to be the size of its toplevel Window. This will be very
	// wrong for most widgets that support native scrolling; for example, if you
	// add a widget such as TreeView with a viewport, the whole widget will
	// scroll, including the column headings. Thus, widgets with native
	// scrolling support should not be used with the Viewport proxy.
	//
	// A widget supports scrolling natively if it implements the Scrollable
	// interface.
	AddWithViewport(child Widget)
	// CaptureButtonPress: return whether button presses are captured during
	// kinetic scrolling. See gtk_scrolled_window_set_capture_button_press().
	CaptureButtonPress() bool
	// HAdjustment returns the horizontal scrollbar’s adjustment, used to
	// connect the horizontal scrollbar to the child widget’s horizontal scroll
	// functionality.
	HAdjustment() Adjustment
	// HScrollbar returns the horizontal scrollbar of @scrolled_window.
	HScrollbar() Widget
	// KineticScrolling returns the specified kinetic scrolling behavior.
	KineticScrolling() bool
	// MaxContentHeight returns the maximum content height set.
	MaxContentHeight() int
	// MaxContentWidth returns the maximum content width set.
	MaxContentWidth() int
	// MinContentHeight gets the minimal content height of @scrolled_window, or
	// -1 if not set.
	MinContentHeight() int
	// MinContentWidth gets the minimum content width of @scrolled_window, or -1
	// if not set.
	MinContentWidth() int
	// OverlayScrolling returns whether overlay scrolling is enabled for this
	// scrolled window.
	OverlayScrolling() bool
	// Placement gets the placement of the contents with respect to the
	// scrollbars for the scrolled window. See
	// gtk_scrolled_window_set_placement().
	Placement() CornerType
	// Policy retrieves the current policy values for the horizontal and
	// vertical scrollbars. See gtk_scrolled_window_set_policy().
	Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType)
	// PropagateNaturalHeight reports whether the natural height of the child
	// will be calculated and propagated through the scrolled window’s requested
	// natural height.
	PropagateNaturalHeight() bool
	// PropagateNaturalWidth reports whether the natural width of the child will
	// be calculated and propagated through the scrolled window’s requested
	// natural width.
	PropagateNaturalWidth() bool
	// ShadowType gets the shadow type of the scrolled window. See
	// gtk_scrolled_window_set_shadow_type().
	ShadowType() ShadowType
	// VAdjustment returns the vertical scrollbar’s adjustment, used to connect
	// the vertical scrollbar to the child widget’s vertical scroll
	// functionality.
	VAdjustment() Adjustment
	// VScrollbar returns the vertical scrollbar of @scrolled_window.
	VScrollbar() Widget
	// SetCaptureButtonPress changes the behaviour of @scrolled_window with
	// regard to the initial event that possibly starts kinetic scrolling. When
	// @capture_button_press is set to true, the event is captured by the
	// scrolled window, and then later replayed if it is meant to go to the
	// child widget.
	//
	// This should be enabled if any child widgets perform non-reversible
	// actions on Widget::button-press-event. If they don't, and handle
	// additionally handle Widget::grab-broken-event, it might be better to set
	// @capture_button_press to false.
	//
	// This setting only has an effect if kinetic scrolling is enabled.
	SetCaptureButtonPress(captureButtonPress bool)
	// SetHAdjustment sets the Adjustment for the horizontal scrollbar.
	SetHAdjustment(hadjustment Adjustment)
	// SetKineticScrolling turns kinetic scrolling on or off. Kinetic scrolling
	// only applies to devices with source GDK_SOURCE_TOUCHSCREEN.
	SetKineticScrolling(kineticScrolling bool)
	// SetMaxContentHeight sets the maximum height that @scrolled_window should
	// keep visible. The @scrolled_window will grow up to this height before it
	// starts scrolling the content.
	//
	// It is a programming error to set the maximum content height to a value
	// smaller than ScrolledWindow:min-content-height.
	SetMaxContentHeight(height int)
	// SetMaxContentWidth sets the maximum width that @scrolled_window should
	// keep visible. The @scrolled_window will grow up to this width before it
	// starts scrolling the content.
	//
	// It is a programming error to set the maximum content width to a value
	// smaller than ScrolledWindow:min-content-width.
	SetMaxContentWidth(width int)
	// SetMinContentHeight sets the minimum height that @scrolled_window should
	// keep visible. Note that this can and (usually will) be smaller than the
	// minimum size of the content.
	//
	// It is a programming error to set the minimum content height to a value
	// greater than ScrolledWindow:max-content-height.
	SetMinContentHeight(height int)
	// SetMinContentWidth sets the minimum width that @scrolled_window should
	// keep visible. Note that this can and (usually will) be smaller than the
	// minimum size of the content.
	//
	// It is a programming error to set the minimum content width to a value
	// greater than ScrolledWindow:max-content-width.
	SetMinContentWidth(width int)
	// SetOverlayScrolling enables or disables overlay scrolling for this
	// scrolled window.
	SetOverlayScrolling(overlayScrolling bool)
	// SetPlacement sets the placement of the contents with respect to the
	// scrollbars for the scrolled window.
	//
	// The default is GTK_CORNER_TOP_LEFT, meaning the child is in the top left,
	// with the scrollbars underneath and to the right. Other values in
	// CornerType are GTK_CORNER_TOP_RIGHT, GTK_CORNER_BOTTOM_LEFT, and
	// GTK_CORNER_BOTTOM_RIGHT.
	//
	// See also gtk_scrolled_window_get_placement() and
	// gtk_scrolled_window_unset_placement().
	SetPlacement(windowPlacement CornerType)
	// SetPolicy sets the scrollbar policy for the horizontal and vertical
	// scrollbars.
	//
	// The policy determines when the scrollbar should appear; it is a value
	// from the PolicyType enumeration. If GTK_POLICY_ALWAYS, the scrollbar is
	// always present; if GTK_POLICY_NEVER, the scrollbar is never present; if
	// GTK_POLICY_AUTOMATIC, the scrollbar is present only if needed (that is,
	// if the slider part of the bar would be smaller than the trough — the
	// display is larger than the page size).
	SetPolicy(hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType)
	// SetPropagateNaturalHeight sets whether the natural height of the child
	// should be calculated and propagated through the scrolled window’s
	// requested natural height.
	SetPropagateNaturalHeight(propagate bool)
	// SetPropagateNaturalWidth sets whether the natural width of the child
	// should be calculated and propagated through the scrolled window’s
	// requested natural width.
	SetPropagateNaturalWidth(propagate bool)
	// SetShadowType changes the type of shadow drawn around the contents of
	// @scrolled_window.
	SetShadowType(typ ShadowType)
	// SetVAdjustment sets the Adjustment for the vertical scrollbar.
	SetVAdjustment(vadjustment Adjustment)
	// UnsetPlacement unsets the placement of the contents with respect to the
	// scrollbars for the scrolled window. If no window placement is set for a
	// scrolled window, it defaults to GTK_CORNER_TOP_LEFT.
	//
	// See also gtk_scrolled_window_set_placement() and
	// gtk_scrolled_window_get_placement().
	UnsetPlacement()
}

// scrolledWindow implements the ScrolledWindow class.
type scrolledWindow struct {
	Bin
	Buildable
}

var _ ScrolledWindow = (*scrolledWindow)(nil)

// WrapScrolledWindow wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrolledWindow(obj *externglib.Object) ScrolledWindow {
	return scrolledWindow{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrolledWindow(obj), nil
}

// NewScrolledWindow constructs a class ScrolledWindow.
func NewScrolledWindow(hadjustment Adjustment, vadjustment Adjustment) ScrolledWindow {
	var _arg1 *C.GtkAdjustment    // out
	var _arg2 *C.GtkAdjustment    // out
	var _cret C.GtkScrolledWindow // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_scrolled_window_new(_arg1, _arg2)

	var _scrolledWindow ScrolledWindow // out

	_scrolledWindow = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(ScrolledWindow)

	return _scrolledWindow
}

// AddWithViewport: used to add children without native scrolling
// capabilities. This is simply a convenience function; it is equivalent to
// adding the unscrollable child to a viewport, then adding the viewport to
// the scrolled window. If a child has native scrolling, use
// gtk_container_add() instead of this function.
//
// The viewport scrolls the child by moving its Window, and takes the size
// of the child to be the size of its toplevel Window. This will be very
// wrong for most widgets that support native scrolling; for example, if you
// add a widget such as TreeView with a viewport, the whole widget will
// scroll, including the column headings. Thus, widgets with native
// scrolling support should not be used with the Viewport proxy.
//
// A widget supports scrolling natively if it implements the Scrollable
// interface.
func (s scrolledWindow) AddWithViewport(child Widget) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_scrolled_window_add_with_viewport(_arg0, _arg1)
}

// CaptureButtonPress: return whether button presses are captured during
// kinetic scrolling. See gtk_scrolled_window_set_capture_button_press().
func (s scrolledWindow) CaptureButtonPress() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_capture_button_press(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HAdjustment returns the horizontal scrollbar’s adjustment, used to
// connect the horizontal scrollbar to the child widget’s horizontal scroll
// functionality.
func (s scrolledWindow) HAdjustment() Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Adjustment)

	return _adjustment
}

// HScrollbar returns the horizontal scrollbar of @scrolled_window.
func (s scrolledWindow) HScrollbar() Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_hscrollbar(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// KineticScrolling returns the specified kinetic scrolling behavior.
func (s scrolledWindow) KineticScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_kinetic_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxContentHeight returns the maximum content height set.
func (s scrolledWindow) MaxContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MaxContentWidth returns the maximum content width set.
func (s scrolledWindow) MaxContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_max_content_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MinContentHeight gets the minimal content height of @scrolled_window, or
// -1 if not set.
func (s scrolledWindow) MinContentHeight() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MinContentWidth gets the minimum content width of @scrolled_window, or -1
// if not set.
func (s scrolledWindow) MinContentWidth() int {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_min_content_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// OverlayScrolling returns whether overlay scrolling is enabled for this
// scrolled window.
func (s scrolledWindow) OverlayScrolling() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_overlay_scrolling(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Placement gets the placement of the contents with respect to the
// scrollbars for the scrolled window. See
// gtk_scrolled_window_set_placement().
func (s scrolledWindow) Placement() CornerType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkCornerType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_placement(_arg0)

	var _cornerType CornerType // out

	_cornerType = CornerType(_cret)

	return _cornerType
}

// Policy retrieves the current policy values for the horizontal and
// vertical scrollbars. See gtk_scrolled_window_set_policy().
func (s scrolledWindow) Policy() (hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // in
	var _arg2 C.GtkPolicyType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	C.gtk_scrolled_window_get_policy(_arg0, &_arg1, &_arg2)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	_hscrollbarPolicy = PolicyType(_arg1)
	_vscrollbarPolicy = PolicyType(_arg2)

	return _hscrollbarPolicy, _vscrollbarPolicy
}

// PropagateNaturalHeight reports whether the natural height of the child
// will be calculated and propagated through the scrolled window’s requested
// natural height.
func (s scrolledWindow) PropagateNaturalHeight() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_height(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PropagateNaturalWidth reports whether the natural width of the child will
// be calculated and propagated through the scrolled window’s requested
// natural width.
func (s scrolledWindow) PropagateNaturalWidth() bool {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_propagate_natural_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShadowType gets the shadow type of the scrolled window. See
// gtk_scrolled_window_set_shadow_type().
func (s scrolledWindow) ShadowType() ShadowType {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret C.GtkShadowType      // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = ShadowType(_cret)

	return _shadowType
}

// VAdjustment returns the vertical scrollbar’s adjustment, used to connect
// the vertical scrollbar to the child widget’s vertical scroll
// functionality.
func (s scrolledWindow) VAdjustment() Adjustment {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkAdjustment     // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Adjustment)

	return _adjustment
}

// VScrollbar returns the vertical scrollbar of @scrolled_window.
func (s scrolledWindow) VScrollbar() Widget {
	var _arg0 *C.GtkScrolledWindow // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrolled_window_get_vscrollbar(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// SetCaptureButtonPress changes the behaviour of @scrolled_window with
// regard to the initial event that possibly starts kinetic scrolling. When
// @capture_button_press is set to true, the event is captured by the
// scrolled window, and then later replayed if it is meant to go to the
// child widget.
//
// This should be enabled if any child widgets perform non-reversible
// actions on Widget::button-press-event. If they don't, and handle
// additionally handle Widget::grab-broken-event, it might be better to set
// @capture_button_press to false.
//
// This setting only has an effect if kinetic scrolling is enabled.
func (s scrolledWindow) SetCaptureButtonPress(captureButtonPress bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if captureButtonPress {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_capture_button_press(_arg0, _arg1)
}

// SetHAdjustment sets the Adjustment for the horizontal scrollbar.
func (s scrolledWindow) SetHAdjustment(hadjustment Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))

	C.gtk_scrolled_window_set_hadjustment(_arg0, _arg1)
}

// SetKineticScrolling turns kinetic scrolling on or off. Kinetic scrolling
// only applies to devices with source GDK_SOURCE_TOUCHSCREEN.
func (s scrolledWindow) SetKineticScrolling(kineticScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if kineticScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_kinetic_scrolling(_arg0, _arg1)
}

// SetMaxContentHeight sets the maximum height that @scrolled_window should
// keep visible. The @scrolled_window will grow up to this height before it
// starts scrolling the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than ScrolledWindow:min-content-height.
func (s scrolledWindow) SetMaxContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.gint)(height)

	C.gtk_scrolled_window_set_max_content_height(_arg0, _arg1)
}

// SetMaxContentWidth sets the maximum width that @scrolled_window should
// keep visible. The @scrolled_window will grow up to this width before it
// starts scrolling the content.
//
// It is a programming error to set the maximum content width to a value
// smaller than ScrolledWindow:min-content-width.
func (s scrolledWindow) SetMaxContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.gint)(width)

	C.gtk_scrolled_window_set_max_content_width(_arg0, _arg1)
}

// SetMinContentHeight sets the minimum height that @scrolled_window should
// keep visible. Note that this can and (usually will) be smaller than the
// minimum size of the content.
//
// It is a programming error to set the minimum content height to a value
// greater than ScrolledWindow:max-content-height.
func (s scrolledWindow) SetMinContentHeight(height int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.gint)(height)

	C.gtk_scrolled_window_set_min_content_height(_arg0, _arg1)
}

// SetMinContentWidth sets the minimum width that @scrolled_window should
// keep visible. Note that this can and (usually will) be smaller than the
// minimum size of the content.
//
// It is a programming error to set the minimum content width to a value
// greater than ScrolledWindow:max-content-width.
func (s scrolledWindow) SetMinContentWidth(width int) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.gint)(width)

	C.gtk_scrolled_window_set_min_content_width(_arg0, _arg1)
}

// SetOverlayScrolling enables or disables overlay scrolling for this
// scrolled window.
func (s scrolledWindow) SetOverlayScrolling(overlayScrolling bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if overlayScrolling {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_overlay_scrolling(_arg0, _arg1)
}

// SetPlacement sets the placement of the contents with respect to the
// scrollbars for the scrolled window.
//
// The default is GTK_CORNER_TOP_LEFT, meaning the child is in the top left,
// with the scrollbars underneath and to the right. Other values in
// CornerType are GTK_CORNER_TOP_RIGHT, GTK_CORNER_BOTTOM_LEFT, and
// GTK_CORNER_BOTTOM_RIGHT.
//
// See also gtk_scrolled_window_get_placement() and
// gtk_scrolled_window_unset_placement().
func (s scrolledWindow) SetPlacement(windowPlacement CornerType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkCornerType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkCornerType)(windowPlacement)

	C.gtk_scrolled_window_set_placement(_arg0, _arg1)
}

// SetPolicy sets the scrollbar policy for the horizontal and vertical
// scrollbars.
//
// The policy determines when the scrollbar should appear; it is a value
// from the PolicyType enumeration. If GTK_POLICY_ALWAYS, the scrollbar is
// always present; if GTK_POLICY_NEVER, the scrollbar is never present; if
// GTK_POLICY_AUTOMATIC, the scrollbar is present only if needed (that is,
// if the slider part of the bar would be smaller than the trough — the
// display is larger than the page size).
func (s scrolledWindow) SetPolicy(hscrollbarPolicy PolicyType, vscrollbarPolicy PolicyType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkPolicyType      // out
	var _arg2 C.GtkPolicyType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkPolicyType)(hscrollbarPolicy)
	_arg2 = (C.GtkPolicyType)(vscrollbarPolicy)

	C.gtk_scrolled_window_set_policy(_arg0, _arg1, _arg2)
}

// SetPropagateNaturalHeight sets whether the natural height of the child
// should be calculated and propagated through the scrolled window’s
// requested natural height.
func (s scrolledWindow) SetPropagateNaturalHeight(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_height(_arg0, _arg1)
}

// SetPropagateNaturalWidth sets whether the natural width of the child
// should be calculated and propagated through the scrolled window’s
// requested natural width.
func (s scrolledWindow) SetPropagateNaturalWidth(propagate bool) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	if propagate {
		_arg1 = C.TRUE
	}

	C.gtk_scrolled_window_set_propagate_natural_width(_arg0, _arg1)
}

// SetShadowType changes the type of shadow drawn around the contents of
// @scrolled_window.
func (s scrolledWindow) SetShadowType(typ ShadowType) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 C.GtkShadowType      // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkShadowType)(typ)

	C.gtk_scrolled_window_set_shadow_type(_arg0, _arg1)
}

// SetVAdjustment sets the Adjustment for the vertical scrollbar.
func (s scrolledWindow) SetVAdjustment(vadjustment Adjustment) {
	var _arg0 *C.GtkScrolledWindow // out
	var _arg1 *C.GtkAdjustment     // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	C.gtk_scrolled_window_set_vadjustment(_arg0, _arg1)
}

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars for the scrolled window. If no window placement is set for a
// scrolled window, it defaults to GTK_CORNER_TOP_LEFT.
//
// See also gtk_scrolled_window_set_placement() and
// gtk_scrolled_window_get_placement().
func (s scrolledWindow) UnsetPlacement() {
	var _arg0 *C.GtkScrolledWindow // out

	_arg0 = (*C.GtkScrolledWindow)(unsafe.Pointer(s.Native()))

	C.gtk_scrolled_window_unset_placement(_arg0)
}
