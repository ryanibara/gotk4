// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// Popover: gtkPopover is a bubble-like context window, primarily meant to
// provide context-dependent information or options. Popovers are attached to a
// widget, passed at construction time on gtk_popover_new(), or updated
// afterwards through gtk_popover_set_relative_to(), by default they will point
// to the whole widget area, although this behavior can be changed through
// gtk_popover_set_pointing_to().
//
// The position of a popover relative to the widget it is attached to can also
// be changed through gtk_popover_set_position().
//
// By default, Popover performs a GTK+ grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Esc key being
// pressed). If no such modal behavior is desired on a popover,
// gtk_popover_set_modal() may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// GtkPopover is often used to replace menus. To facilitate this, it supports
// being populated from a Model, using gtk_popover_new_from_model(). In addition
// to all the regular menu model features, this function supports rendering
// sections in the model in a more compact form, as a row of icon buttons
// instead of menu items.
//
// To use this rendering, set the ”display-hint” attribute of the section to
// ”horizontal-buttons” and set the icons of your items with the ”verb-icon”
// attribute.
//
//    <section>
//      <attribute name="display-hint">horizontal-buttons</attribute>
//      <item>
//        <attribute name="label">Cut</attribute>
//        <attribute name="action">app.cut</attribute>
//        <attribute name="verb-icon">edit-cut-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Copy</attribute>
//        <attribute name="action">app.copy</attribute>
//        <attribute name="verb-icon">edit-copy-symbolic</attribute>
//      </item>
//      <item>
//        <attribute name="label">Paste</attribute>
//        <attribute name="action">app.paste</attribute>
//        <attribute name="verb-icon">edit-paste-symbolic</attribute>
//      </item>
//    </section>
//
//
// CSS nodes
//
// GtkPopover has a single css node called popover. It always gets the
// .background style class and it gets the .menu style class if it is menu-like
// (e.g. PopoverMenu or created using gtk_popover_new_from_model().
//
// Particular uses of GtkPopover, such as touch selection popups or magnifiers
// in Entry or TextView get style classes like .touch-selection or .magnifier to
// differentiate from plain popovers.
type Popover interface {
	Bin
	Buildable

	// BindModel establishes a binding between a Popover and a Model.
	//
	// The contents of @popover are removed and then refilled with menu items
	// according to @model. When @model changes, @popover is updated. Calling
	// this function twice on @popover with different @model will cause the
	// first binding to be replaced with a binding to the new model. If @model
	// is nil then any previous binding is undone and all children are removed.
	//
	// If @action_namespace is non-nil then the effect is as if all actions
	// mentioned in the @model have their names prefixed with the namespace,
	// plus a dot. For example, if the action “quit” is mentioned and
	// @action_namespace is “app” then the effective action name is “app.quit”.
	//
	// This function uses Actionable to define the action name and target values
	// on the created menu items. If you want to use an action group other than
	// “app” and “win”, or if you want to use a MenuShell outside of a
	// ApplicationWindow, then you will need to attach your own action group to
	// the widget hierarchy using gtk_widget_insert_action_group(). As an
	// example, if you created a group with a “quit” action and inserted it with
	// the name “mygroup” then you would use the action name “mygroup.quit” in
	// your Model.
	BindModel(model gio.MenuModel, actionNamespace string)
	// ConstrainTo returns the constraint for placing this popover. See
	// gtk_popover_set_constrain_to().
	ConstrainTo() PopoverConstraint
	// DefaultWidget gets the widget that should be set as the default while the
	// popover is shown.
	DefaultWidget() Widget
	// Modal returns whether the popover is modal, see gtk_popover_set_modal to
	// see the implications of this.
	Modal() bool
	// PointingTo: if a rectangle to point to has been set, this function will
	// return true and fill in @rect with such rectangle, otherwise it will
	// return false and fill in @rect with the attached widget coordinates.
	PointingTo() (rect gdk.Rectangle, ok bool)
	// Position returns the preferred position of @popover.
	Position() PositionType
	// RelativeTo returns the widget @popover is currently attached to
	RelativeTo() Widget
	// TransitionsEnabled returns whether show/hide transitions are enabled on
	// this popover.
	TransitionsEnabled() bool
	// Popdown pops @popover down.This is different than a gtk_widget_hide()
	// call in that it shows the popover with a transition. If you want to hide
	// the popover without a transition, use gtk_widget_hide().
	Popdown()
	// Popup pops @popover up. This is different than a gtk_widget_show() call
	// in that it shows the popover with a transition. If you want to show the
	// popover without a transition, use gtk_widget_show().
	Popup()
	// SetConstrainTo sets a constraint for positioning this popover.
	//
	// Note that not all platforms support placing popovers freely, and may
	// already impose constraints.
	SetConstrainTo(constraint PopoverConstraint)
	// SetDefaultWidget sets the widget that should be set as default widget
	// while the popover is shown (see gtk_window_set_default()). Popover
	// remembers the previous default widget and reestablishes it when the
	// popover is dismissed.
	SetDefaultWidget(widget Widget)
	// SetModal sets whether @popover is modal, a modal popover will grab all
	// input within the toplevel and grab the keyboard focus on it when being
	// displayed. Clicking outside the popover area or pressing Esc will dismiss
	// the popover and ungrab input.
	SetModal(modal bool)
	// SetPointingTo sets the rectangle that @popover will point to, in the
	// coordinate space of the widget @popover is attached to, see
	// gtk_popover_set_relative_to().
	SetPointingTo(rect *gdk.Rectangle)
	// SetPosition sets the preferred position for @popover to appear. If the
	// @popover is currently visible, it will be immediately updated.
	//
	// This preference will be respected where possible, although on lack of
	// space (eg. if close to the window edges), the Popover may choose to
	// appear on the opposite side
	SetPosition(position PositionType)
	// SetRelativeTo sets a new widget to be attached to @popover. If @popover
	// is visible, the position will be updated.
	//
	// Note: the ownership of popovers is always given to their @relative_to
	// widget, so if @relative_to is set to nil on an attached @popover, it will
	// be detached from its previous widget, and consequently destroyed unless
	// extra references are kept.
	SetRelativeTo(relativeTo Widget)
	// SetTransitionsEnabled sets whether show/hide transitions are enabled on
	// this popover
	SetTransitionsEnabled(transitionsEnabled bool)
}

// popover implements the Popover interface.
type popover struct {
	Bin
	Buildable
}

var _ Popover = (*popover)(nil)

// WrapPopover wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopover(obj *externglib.Object) Popover {
	return Popover{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopover(obj), nil
}

// NewPopover constructs a class Popover.
func NewPopover(relativeTo Widget) Popover {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	var cret C.GtkPopover
	var ret1 Popover

	cret = C.gtk_popover_new(relativeTo)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Popover)

	return ret1
}

// NewPopoverFromModel constructs a class Popover.
func NewPopoverFromModel(relativeTo Widget, model gio.MenuModel) Popover {
	var arg1 *C.GtkWidget
	var arg2 *C.GMenuModel

	arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	arg2 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	var cret C.GtkPopover
	var ret1 Popover

	cret = C.gtk_popover_new_from_model(relativeTo, model)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Popover)

	return ret1
}

// BindModel establishes a binding between a Popover and a Model.
//
// The contents of @popover are removed and then refilled with menu items
// according to @model. When @model changes, @popover is updated. Calling
// this function twice on @popover with different @model will cause the
// first binding to be replaced with a binding to the new model. If @model
// is nil then any previous binding is undone and all children are removed.
//
// If @action_namespace is non-nil then the effect is as if all actions
// mentioned in the @model have their names prefixed with the namespace,
// plus a dot. For example, if the action “quit” is mentioned and
// @action_namespace is “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values
// on the created menu items. If you want to use an action group other than
// “app” and “win”, or if you want to use a MenuShell outside of a
// ApplicationWindow, then you will need to attach your own action group to
// the widget hierarchy using gtk_widget_insert_action_group(). As an
// example, if you created a group with a “quit” action and inserted it with
// the name “mygroup” then you would use the action name “mygroup.quit” in
// your Model.
func (p popover) BindModel(model gio.MenuModel, actionNamespace string) {
	var arg0 *C.GtkPopover
	var arg1 *C.GMenuModel
	var arg2 *C.gchar

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.gchar)(C.CString(actionNamespace))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_popover_bind_model(arg0, model, actionNamespace)
}

// ConstrainTo returns the constraint for placing this popover. See
// gtk_popover_set_constrain_to().
func (p popover) ConstrainTo() PopoverConstraint {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.GtkPopoverConstraint
	var ret1 PopoverConstraint

	cret = C.gtk_popover_get_constrain_to(arg0)

	ret1 = PopoverConstraint(cret)

	return ret1
}

// DefaultWidget gets the widget that should be set as the default while the
// popover is shown.
func (p popover) DefaultWidget() Widget {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_popover_get_default_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Modal returns whether the popover is modal, see gtk_popover_set_modal to
// see the implications of this.
func (p popover) Modal() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_popover_get_modal(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// PointingTo: if a rectangle to point to has been set, this function will
// return true and fill in @rect with such rectangle, otherwise it will
// return false and fill in @rect with the attached widget coordinates.
func (p popover) PointingTo() (rect gdk.Rectangle, ok bool) {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var arg1 C.GdkRectangle
	var ret1 *gdk.Rectangle
	var cret C.gboolean
	var ret2 bool

	cret = C.gtk_popover_get_pointing_to(arg0, &arg1)

	*ret1 = gdk.WrapRectangle(unsafe.Pointer(arg1))
	ret2 = C.bool(cret) != C.false

	return ret1, ret2
}

// Position returns the preferred position of @popover.
func (p popover) Position() PositionType {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.GtkPositionType
	var ret1 PositionType

	cret = C.gtk_popover_get_position(arg0)

	ret1 = PositionType(cret)

	return ret1
}

// RelativeTo returns the widget @popover is currently attached to
func (p popover) RelativeTo() Widget {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_popover_get_relative_to(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// TransitionsEnabled returns whether show/hide transitions are enabled on
// this popover.
func (p popover) TransitionsEnabled() bool {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_popover_get_transitions_enabled(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Popdown pops @popover down.This is different than a gtk_widget_hide()
// call in that it shows the popover with a transition. If you want to hide
// the popover without a transition, use gtk_widget_hide().
func (p popover) Popdown() {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(arg0)
}

// Popup pops @popover up. This is different than a gtk_widget_show() call
// in that it shows the popover with a transition. If you want to show the
// popover without a transition, use gtk_widget_show().
func (p popover) Popup() {
	var arg0 *C.GtkPopover

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(arg0)
}

// SetConstrainTo sets a constraint for positioning this popover.
//
// Note that not all platforms support placing popovers freely, and may
// already impose constraints.
func (p popover) SetConstrainTo(constraint PopoverConstraint) {
	var arg0 *C.GtkPopover
	var arg1 C.GtkPopoverConstraint

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (C.GtkPopoverConstraint)(constraint)

	C.gtk_popover_set_constrain_to(arg0, constraint)
}

// SetDefaultWidget sets the widget that should be set as default widget
// while the popover is shown (see gtk_window_set_default()). Popover
// remembers the previous default widget and reestablishes it when the
// popover is dismissed.
func (p popover) SetDefaultWidget(widget Widget) {
	var arg0 *C.GtkPopover
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(arg0, widget)
}

// SetModal sets whether @popover is modal, a modal popover will grab all
// input within the toplevel and grab the keyboard focus on it when being
// displayed. Clicking outside the popover area or pressing Esc will dismiss
// the popover and ungrab input.
func (p popover) SetModal(modal bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if modal {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_modal(arg0, modal)
}

// SetPointingTo sets the rectangle that @popover will point to, in the
// coordinate space of the widget @popover is attached to, see
// gtk_popover_set_relative_to().
func (p popover) SetPointingTo(rect *gdk.Rectangle) {
	var arg0 *C.GtkPopover
	var arg1 *C.GdkRectangle

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_popover_set_pointing_to(arg0, rect)
}

// SetPosition sets the preferred position for @popover to appear. If the
// @popover is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of
// space (eg. if close to the window edges), the Popover may choose to
// appear on the opposite side
func (p popover) SetPosition(position PositionType) {
	var arg0 *C.GtkPopover
	var arg1 C.GtkPositionType

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (C.GtkPositionType)(position)

	C.gtk_popover_set_position(arg0, position)
}

// SetRelativeTo sets a new widget to be attached to @popover. If @popover
// is visible, the position will be updated.
//
// Note: the ownership of popovers is always given to their @relative_to
// widget, so if @relative_to is set to nil on an attached @popover, it will
// be detached from its previous widget, and consequently destroyed unless
// extra references are kept.
func (p popover) SetRelativeTo(relativeTo Widget) {
	var arg0 *C.GtkPopover
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))

	C.gtk_popover_set_relative_to(arg0, relativeTo)
}

// SetTransitionsEnabled sets whether show/hide transitions are enabled on
// this popover
func (p popover) SetTransitionsEnabled(transitionsEnabled bool) {
	var arg0 *C.GtkPopover
	var arg1 C.gboolean

	arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if transitionsEnabled {
		arg1 = C.gboolean(1)
	}

	C.gtk_popover_set_transitions_enabled(arg0, transitionsEnabled)
}

type PopoverPrivate struct {
	native C.GtkPopoverPrivate
}

// WrapPopoverPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPopoverPrivate(ptr unsafe.Pointer) *PopoverPrivate {
	if ptr == nil {
		return nil
	}

	return (*PopoverPrivate)(ptr)
}

func marshalPopoverPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPopoverPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PopoverPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
