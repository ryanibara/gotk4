// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_PopoverClass_closed(GtkPopover*);
// extern void _gotk4_gtk3_Popover_ConnectClosed(gpointer, guintptr);
import "C"

// glib.Type values for gtkpopover.go.
var GTypePopover = externglib.Type(C.gtk_popover_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypePopover, F: marshalPopover},
	})
}

// PopoverOverrider contains methods that are overridable.
type PopoverOverrider interface {
	Closed()
}

// Popover is a bubble-like context window, primarily meant to provide
// context-dependent information or options. Popovers are attached to a widget,
// passed at construction time on gtk_popover_new(), or updated afterwards
// through gtk_popover_set_relative_to(), by default they will point to the
// whole widget area, although this behavior can be changed through
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
type Popover struct {
	_ [0]func() // equal guard
	Bin
}

var (
	_ Binner = (*Popover)(nil)
)

func classInitPopoverer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkPopoverClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkPopoverClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Closed() }); ok {
		pclass.closed = (*[0]byte)(C._gotk4_gtk3_PopoverClass_closed)
	}
}

//export _gotk4_gtk3_PopoverClass_closed
func _gotk4_gtk3_PopoverClass_closed(arg0 *C.GtkPopover) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closed() })

	iface.Closed()
}

func wrapPopover(obj *externglib.Object) *Popover {
	return &Popover{
		Bin: Bin{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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

func marshalPopover(p uintptr) (interface{}, error) {
	return wrapPopover(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Popover_ConnectClosed
func _gotk4_gtk3_Popover_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClosed: this signal is emitted when the popover is dismissed either
// through API or user interaction.
func (popover *Popover) ConnectClosed(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(popover, "closed", false, unsafe.Pointer(C._gotk4_gtk3_Popover_ConnectClosed), f)
}

// NewPopover creates a new popover to point to relative_to.
//
// The function takes the following parameters:
//
//    - relativeTo (optional) the popover is related to.
//
// The function returns the following values:
//
//    - popover: new Popover.
//
func NewPopover(relativeTo Widgetter) *Popover {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkWidget // in

	if relativeTo != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	}

	_cret = C.gtk_popover_new(_arg1)
	runtime.KeepAlive(relativeTo)

	var _popover *Popover // out

	_popover = wrapPopover(externglib.Take(unsafe.Pointer(_cret)))

	return _popover
}

// NewPopoverFromModel creates a Popover and populates it according to model.
// The popover is pointed to the relative_to widget.
//
// The created buttons are connected to actions found in the ApplicationWindow
// to which the popover belongs - typically by means of being attached to a
// widget that is contained within the ApplicationWindows widget hierarchy.
//
// Actions can also be added using gtk_widget_insert_action_group() on the menus
// attach widget or on any of its parent widgets.
//
// The function takes the following parameters:
//
//    - relativeTo (optional) the popover is related to.
//    - model: Model.
//
// The function returns the following values:
//
//    - popover: new Popover.
//
func NewPopoverFromModel(relativeTo Widgetter, model gio.MenuModeller) *Popover {
	var _arg1 *C.GtkWidget  // out
	var _arg2 *C.GMenuModel // out
	var _cret *C.GtkWidget  // in

	if relativeTo != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	}
	_arg2 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.gtk_popover_new_from_model(_arg1, _arg2)
	runtime.KeepAlive(relativeTo)
	runtime.KeepAlive(model)

	var _popover *Popover // out

	_popover = wrapPopover(externglib.Take(unsafe.Pointer(_cret)))

	return _popover
}

// BindModel establishes a binding between a Popover and a Model.
//
// The contents of popover are removed and then refilled with menu items
// according to model. When model changes, popover is updated. Calling this
// function twice on popover with different model will cause the first binding
// to be replaced with a binding to the new model. If model is NULL then any
// previous binding is undone and all children are removed.
//
// If action_namespace is non-NULL then the effect is as if all actions
// mentioned in the model have their names prefixed with the namespace, plus a
// dot. For example, if the action “quit” is mentioned and action_namespace is
// “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values on
// the created menu items. If you want to use an action group other than “app”
// and “win”, or if you want to use a MenuShell outside of a ApplicationWindow,
// then you will need to attach your own action group to the widget hierarchy
// using gtk_widget_insert_action_group(). As an example, if you created a group
// with a “quit” action and inserted it with the name “mygroup” then you would
// use the action name “mygroup.quit” in your Model.
//
// The function takes the following parameters:
//
//    - model (optional) to bind to or NULL to remove binding.
//    - actionNamespace (optional): namespace for actions in model.
//
func (popover *Popover) BindModel(model gio.MenuModeller, actionNamespace string) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GMenuModel // out
	var _arg2 *C.gchar      // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	}
	if actionNamespace != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(actionNamespace)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_popover_bind_model(_arg0, _arg1, _arg2)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(model)
	runtime.KeepAlive(actionNamespace)
}

// ConstrainTo returns the constraint for placing this popover. See
// gtk_popover_set_constrain_to().
//
// The function returns the following values:
//
//    - popoverConstraint: constraint for placing this popover.
//
func (popover *Popover) ConstrainTo() PopoverConstraint {
	var _arg0 *C.GtkPopover          // out
	var _cret C.GtkPopoverConstraint // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_constrain_to(_arg0)
	runtime.KeepAlive(popover)

	var _popoverConstraint PopoverConstraint // out

	_popoverConstraint = PopoverConstraint(_cret)

	return _popoverConstraint
}

// DefaultWidget gets the widget that should be set as the default while the
// popover is shown.
//
// The function returns the following values:
//
//    - widget (optional): default widget, or NULL if there is none.
//
func (popover *Popover) DefaultWidget() Widgetter {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_default_widget(_arg0)
	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// Modal returns whether the popover is modal, see gtk_popover_set_modal to see
// the implications of this.
//
// The function returns the following values:
//
//    - ok if popover is modal.
//
func (popover *Popover) Modal() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_modal(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PointingTo: if a rectangle to point to has been set, this function will
// return TRUE and fill in rect with such rectangle, otherwise it will return
// FALSE and fill in rect with the attached widget coordinates.
//
// The function returns the following values:
//
//    - rect: location to store the rectangle.
//    - ok: TRUE if a rectangle to point to was set.
//
func (popover *Popover) PointingTo() (*gdk.Rectangle, bool) {
	var _arg0 *C.GtkPopover  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_pointing_to(_arg0, &_arg1)
	runtime.KeepAlive(popover)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// Position returns the preferred position of popover.
//
// The function returns the following values:
//
//    - positionType: preferred position.
//
func (popover *Popover) Position() PositionType {
	var _arg0 *C.GtkPopover     // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_position(_arg0)
	runtime.KeepAlive(popover)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// RelativeTo returns the widget popover is currently attached to.
//
// The function returns the following values:
//
//    - widget: Widget.
//
func (popover *Popover) RelativeTo() Widgetter {
	var _arg0 *C.GtkPopover // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_relative_to(_arg0)
	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
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

// TransitionsEnabled returns whether show/hide transitions are enabled on this
// popover.
//
// Deprecated: You can show or hide the popover without transitions using
// gtk_widget_show() and gtk_widget_hide() while gtk_popover_popup() and
// gtk_popover_popdown() will use transitions.
//
// The function returns the following values:
//
//    - ok if the show and hide transitions of the given popover are enabled, LSE
//      otherwise.
//
func (popover *Popover) TransitionsEnabled() bool {
	var _arg0 *C.GtkPopover // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	_cret = C.gtk_popover_get_transitions_enabled(_arg0)
	runtime.KeepAlive(popover)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown pops popover down.This is different than a gtk_widget_hide() call in
// that it shows the popover with a transition. If you want to hide the popover
// without a transition, use gtk_widget_hide().
func (popover *Popover) Popdown() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_popdown(_arg0)
	runtime.KeepAlive(popover)
}

// Popup pops popover up. This is different than a gtk_widget_show() call in
// that it shows the popover with a transition. If you want to show the popover
// without a transition, use gtk_widget_show().
func (popover *Popover) Popup() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))

	C.gtk_popover_popup(_arg0)
	runtime.KeepAlive(popover)
}

// SetConstrainTo sets a constraint for positioning this popover.
//
// Note that not all platforms support placing popovers freely, and may already
// impose constraints.
//
// The function takes the following parameters:
//
//    - constraint: new constraint.
//
func (popover *Popover) SetConstrainTo(constraint PopoverConstraint) {
	var _arg0 *C.GtkPopover          // out
	var _arg1 C.GtkPopoverConstraint // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = C.GtkPopoverConstraint(constraint)

	C.gtk_popover_set_constrain_to(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(constraint)
}

// SetDefaultWidget sets the widget that should be set as default widget while
// the popover is shown (see gtk_window_set_default()). Popover remembers the
// previous default widget and reestablishes it when the popover is dismissed.
//
// The function takes the following parameters:
//
//    - widget (optional): new default widget, or NULL.
//
func (popover *Popover) SetDefaultWidget(widget Widgetter) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_popover_set_default_widget(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(widget)
}

// SetModal sets whether popover is modal, a modal popover will grab all input
// within the toplevel and grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the popover
// and ungrab input.
//
// The function takes the following parameters:
//
//    - modal to make popover claim all input within the toplevel.
//
func (popover *Popover) SetModal(modal bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if modal {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_modal(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(modal)
}

// SetPointingTo sets the rectangle that popover will point to, in the
// coordinate space of the widget popover is attached to, see
// gtk_popover_set_relative_to().
//
// The function takes the following parameters:
//
//    - rect: rectangle to point to.
//
func (popover *Popover) SetPointingTo(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect)))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(rect)
}

// SetPosition sets the preferred position for popover to appear. If the popover
// is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of space
// (eg. if close to the window edges), the Popover may choose to appear on the
// opposite side.
//
// The function takes the following parameters:
//
//    - position: preferred popover position.
//
func (popover *Popover) SetPosition(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	_arg1 = C.GtkPositionType(position)

	C.gtk_popover_set_position(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(position)
}

// SetRelativeTo sets a new widget to be attached to popover. If popover is
// visible, the position will be updated.
//
// Note: the ownership of popovers is always given to their relative_to widget,
// so if relative_to is set to NULL on an attached popover, it will be detached
// from its previous widget, and consequently destroyed unless extra references
// are kept.
//
// The function takes the following parameters:
//
//    - relativeTo (optional): Widget.
//
func (popover *Popover) SetRelativeTo(relativeTo Widgetter) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if relativeTo != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(relativeTo.Native()))
	}

	C.gtk_popover_set_relative_to(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(relativeTo)
}

// SetTransitionsEnabled sets whether show/hide transitions are enabled on this
// popover
//
// Deprecated: You can show or hide the popover without transitions using
// gtk_widget_show() and gtk_widget_hide() while gtk_popover_popup() and
// gtk_popover_popdown() will use transitions.
//
// The function takes the following parameters:
//
//    - transitionsEnabled: whether transitions are enabled.
//
func (popover *Popover) SetTransitionsEnabled(transitionsEnabled bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(popover.Native()))
	if transitionsEnabled {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_transitions_enabled(_arg0, _arg1)
	runtime.KeepAlive(popover)
	runtime.KeepAlive(transitionsEnabled)
}
