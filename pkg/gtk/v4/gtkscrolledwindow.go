// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
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
// packed into a GtkScrolledWindow.
//
// This is effectively the opposite of where the scroll bars are placed.
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
	// example, when all of a GtkTreeView can not be seen.
	PolicyAutomatic
	// PolicyNever: scrollbar should never appear. In this mode the content
	// determines the size.
	PolicyNever
	// PolicyExternal: don't show a scrollbar, but don't force the size to
	// follow the content. This can be used e.g. to make multiple scrolled
	// windows share a scrollbar.
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

// ScrolledWindow: GtkScrolledWindow is a container that makes its child
// scrollable.
//
// It does so using either internally added scrollbars or externally associated
// adjustments, and optionally draws a frame around the child.
//
// Widgets with native scrolling support, i.e. those whose classes implement the
// gtk.Scrollable interface, are added directly. For other types of widget, the
// class gtk.Viewport acts as an adaptor, giving scrollability to other widgets.
// gtk.ScrolledWindow.SetChild() intelligently accounts for whether or not the
// added child is a GtkScrollable. If it isn’t, then it wraps the child in a
// GtkViewport. Therefore, you can just add any child widget and not worry about
// the details.
//
// If gtk.ScrolledWindow.SetChild() has added a GtkViewport for you, you can
// remove both your added child widget from the GtkViewport, and the GtkViewport
// from the GtkScrolledWindow, like this:
//
//    GtkWidget *scrolled_window = gtk_scrolled_window_new ();
//    GtkWidget *child_widget = gtk_button_new ();
//
//    // GtkButton is not a GtkScrollable, so GtkScrolledWindow will automatically
//    // add a GtkViewport.
//    gtk_box_append (GTK_BOX (scrolled_window), child_widget);
//
//    // Either of these will result in child_widget being unparented:
//    gtk_box_remove (GTK_BOX (scrolled_window), child_widget);
//    // or
//    gtk_box_remove (GTK_BOX (scrolled_window),
//                          gtk_bin_get_child (GTK_BIN (scrolled_window)));
//
//
// Unless gtk.ScrolledWindow:hscrollbar-policy and
// gtk.ScrolledWindow:vscrollbar-policy are GTK_POLICY_NEVER or
// GTK_POLICY_EXTERNAL, GtkScrolledWindow adds internal GtkScrollbar widgets
// around its child. The scroll position of the child, and if applicable the
// scrollbars, is controlled by the gtk.ScrolledWindow:hadjustment and
// gtk.ScrolledWindow:vadjustment that are associated with the
// GtkScrolledWindow. See the docs on gtk.Scrollbar for the details, but note
// that the “step_increment” and “page_increment” fields are only effective if
// the policy causes scrollbars to be present.
//
// If a GtkScrolledWindow doesn’t behave quite as you would like, or doesn’t
// have exactly the right layout, it’s very possible to set up your own
// scrolling with GtkScrollbar and for example a GtkGrid.
//
//
// Touch support
//
// GtkScrolledWindow has built-in support for touch devices. When a touchscreen
// is used, swiping will move the scrolled window, and will expose 'kinetic'
// behavior. This can be turned off with the
// gtk.ScrolledWindow:kinetic-scrolling property if it is undesired.
//
// GtkScrolledWindow also displays visual 'overshoot' indication when the
// content is pulled beyond the end, and this situation can be captured with the
// gtk.ScrolledWindow::edge-overshot signal.
//
// If no mouse device is present, the scrollbars will overlaid as narrow,
// auto-hiding indicators over the content. If traditional scrollbars are
// desired although no mouse is present, this behaviour can be turned off with
// the gtk.ScrolledWindow:overlay-scrolling property.
//
//
// CSS nodes
//
// GtkScrolledWindow has a main CSS node with name scrolledwindow. It gets a
// .frame style class added when gtk.ScrolledWindow:has-frame is TRUE.
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
//
//
// Accessibility
//
// GtkScrolledWindow uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type ScrolledWindow struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*ScrolledWindow)(nil)
)

func wrapScrolledWindow(obj *coreglib.Object) *ScrolledWindow {
	return &ScrolledWindow{
		Widget: Widget{
			InitiallyUnowned: coreglib.InitiallyUnowned{
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

func marshalScrolledWindow(p uintptr) (interface{}, error) {
	return wrapScrolledWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewScrolledWindow creates a new scrolled window.
//
// The function returns the following values:
//
//    - scrolledWindow: new scrolled window.
//
func NewScrolledWindow() *ScrolledWindow {
	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("new_ScrolledWindow", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _scrolledWindow *ScrolledWindow // out

	_scrolledWindow = wrapScrolledWindow(coreglib.Take(unsafe.Pointer(_cret)))

	return _scrolledWindow
}

// Child gets the child widget of scrolled_window.
//
// The function returns the following values:
//
//    - widget (optional): child widget of scrolled_window.
//
func (scrolledWindow *ScrolledWindow) Child() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_child", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(_cret)

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
	}

	return _widget
}

// HAdjustment returns the horizontal scrollbar’s adjustment.
//
// This is the adjustment used to connect the horizontal scrollbar to the child
// widget’s horizontal scroll functionality.
//
// The function returns the following values:
//
//    - adjustment: horizontal GtkAdjustment.
//
func (scrolledWindow *ScrolledWindow) HAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_hadjustment", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// HasFrame gets whether the scrolled window draws a frame.
//
// The function returns the following values:
//
//    - ok: TRUE if the scrolled_window has a frame.
//
func (scrolledWindow *ScrolledWindow) HasFrame() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_has_frame", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_hscrollbar", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_kinetic_scrolling", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_max_content_height", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_max_content_width", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// MinContentHeight gets the minimal content height of scrolled_window.
//
// The function returns the following values:
//
//    - gint: minimal content height.
//
func (scrolledWindow *ScrolledWindow) MinContentHeight() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_min_content_height", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// MinContentWidth gets the minimum content width of scrolled_window.
//
// The function returns the following values:
//
//    - gint: minimum content width.
//
func (scrolledWindow *ScrolledWindow) MinContentWidth() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_min_content_width", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_overlay_scrolling", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// Policy retrieves the current policy values for the horizontal and vertical
// scrollbars.
//
// See gtk.ScrolledWindow.SetPolicy().
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

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_policy", _args[:], _outs[:])

	runtime.KeepAlive(scrolledWindow)

	var _hscrollbarPolicy PolicyType // out
	var _vscrollbarPolicy PolicyType // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_hscrollbarPolicy = *(*PolicyType)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_vscrollbarPolicy = *(*PolicyType)(unsafe.Pointer(_outs[1]))
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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_propagate_natural_height", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_propagate_natural_width", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// VAdjustment returns the vertical scrollbar’s adjustment.
//
// This is the adjustment used to connect the vertical scrollbar to the child
// widget’s vertical scroll functionality.
//
// The function returns the following values:
//
//    - adjustment: vertical GtkAdjustment.
//
func (scrolledWindow *ScrolledWindow) VAdjustment() *Adjustment {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_vadjustment", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(coreglib.Take(unsafe.Pointer(_cret)))

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

	_gret := girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("get_vscrollbar", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scrolledWindow)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
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

// SetChild sets the child widget of scrolled_window.
//
// The function takes the following parameters:
//
//    - child (optional) widget.
//
func (scrolledWindow *ScrolledWindow) SetChild(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if child != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	}

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_child", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(child)
}

// SetHAdjustment sets the GtkAdjustment for the horizontal scrollbar.
//
// The function takes the following parameters:
//
//    - hadjustment (optional): GtkAdjustment to use, or NULL to create a new
//      one.
//
func (scrolledWindow *ScrolledWindow) SetHAdjustment(hadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if hadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(hadjustment).Native()))
	}

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_hadjustment", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hadjustment)
}

// SetHasFrame changes the frame drawn around the contents of scrolled_window.
//
// The function takes the following parameters:
//
//    - hasFrame: whether to draw a frame around scrolled window contents.
//
func (scrolledWindow *ScrolledWindow) SetHasFrame(hasFrame bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if hasFrame {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_has_frame", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(hasFrame)
}

// SetKineticScrolling turns kinetic scrolling on or off.
//
// Kinetic scrolling only applies to devices with source GDK_SOURCE_TOUCHSCREEN.
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

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_kinetic_scrolling", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(kineticScrolling)
}

// SetMaxContentHeight sets the maximum height that scrolled_window should keep
// visible.
//
// The scrolled_window will grow up to this height before it starts scrolling
// the content.
//
// It is a programming error to set the maximum content height to a value
// smaller than gtk.ScrolledWindow:min-content-height.
//
// The function takes the following parameters:
//
//    - height: maximum content height.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentHeight(height int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(height)

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_max_content_height", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMaxContentWidth sets the maximum width that scrolled_window should keep
// visible.
//
// The scrolled_window will grow up to this width before it starts scrolling the
// content.
//
// It is a programming error to set the maximum content width to a value smaller
// than gtk.ScrolledWindow:min-content-width.
//
// The function takes the following parameters:
//
//    - width: maximum content width.
//
func (scrolledWindow *ScrolledWindow) SetMaxContentWidth(width int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_max_content_width", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(width)
}

// SetMinContentHeight sets the minimum height that scrolled_window should keep
// visible.
//
// Note that this can and (usually will) be smaller than the minimum size of the
// content.
//
// It is a programming error to set the minimum content height to a value
// greater than gtk.ScrolledWindow:max-content-height.
//
// The function takes the following parameters:
//
//    - height: minimal content height.
//
func (scrolledWindow *ScrolledWindow) SetMinContentHeight(height int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(height)

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_min_content_height", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(height)
}

// SetMinContentWidth sets the minimum width that scrolled_window should keep
// visible.
//
// Note that this can and (usually will) be smaller than the minimum size of the
// content.
//
// It is a programming error to set the minimum content width to a value greater
// than gtk.ScrolledWindow:max-content-width.
//
// The function takes the following parameters:
//
//    - width: minimal content width.
//
func (scrolledWindow *ScrolledWindow) SetMinContentWidth(width int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_min_content_width", _args[:], nil)

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

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_overlay_scrolling", _args[:], nil)

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

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_propagate_natural_height", _args[:], nil)

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

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_propagate_natural_width", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(propagate)
}

// SetVAdjustment sets the GtkAdjustment for the vertical scrollbar.
//
// The function takes the following parameters:
//
//    - vadjustment (optional): GtkAdjustment to use, or NULL to create a new
//      one.
//
func (scrolledWindow *ScrolledWindow) SetVAdjustment(vadjustment *Adjustment) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))
	if vadjustment != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(vadjustment).Native()))
	}

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("set_vadjustment", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
	runtime.KeepAlive(vadjustment)
}

// UnsetPlacement unsets the placement of the contents with respect to the
// scrollbars.
//
// If no window placement is set for a scrolled window, it defaults to
// GTK_CORNER_TOP_LEFT.
func (scrolledWindow *ScrolledWindow) UnsetPlacement() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scrolledWindow).Native()))

	girepository.MustFind("Gtk", "ScrolledWindow").InvokeMethod("unset_placement", _args[:], nil)

	runtime.KeepAlive(scrolledWindow)
}
