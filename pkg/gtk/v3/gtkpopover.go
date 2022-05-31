// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_PopoverClass_closed(void*);
// extern void _gotk4_gtk3_Popover_ConnectClosed(gpointer, guintptr);
import "C"

// glib.Type values for gtkpopover.go.
var GTypePopover = coreglib.Type(C.gtk_popover_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
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
func _gotk4_gtk3_PopoverClass_closed(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Closed() })

	iface.Closed()
}

func wrapPopover(obj *coreglib.Object) *Popover {
	return &Popover{
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

func marshalPopover(p uintptr) (interface{}, error) {
	return wrapPopover(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Popover_ConnectClosed
func _gotk4_gtk3_Popover_ConnectClosed(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
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
func (popover *Popover) ConnectClosed(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(popover, "closed", false, unsafe.Pointer(C._gotk4_gtk3_Popover_ConnectClosed), f)
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
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if relativeTo != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(relativeTo).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("new_Popover", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(relativeTo)

	var _popover *Popover // out

	_popover = wrapPopover(coreglib.Take(unsafe.Pointer(_cret)))

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	if relativeTo != nil {
		_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(relativeTo).Native()))
	}
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("new_Popover_from_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(relativeTo)
	runtime.KeepAlive(model)

	var _popover *Popover // out

	_popover = wrapPopover(coreglib.Take(unsafe.Pointer(_cret)))

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
	var _args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if model != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}
	if actionNamespace != "" {
		_arg2 = (*C.void)(unsafe.Pointer(C.CString(actionNamespace)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1
	*(**C.void)(unsafe.Pointer(&_args[2])) = _arg2

	girepository.MustFind("Gtk", "Popover").InvokeMethod("bind_model", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(model)
	runtime.KeepAlive(actionNamespace)
}

// DefaultWidget gets the widget that should be set as the default while the
// popover is shown.
//
// The function returns the following values:
//
//    - widget (optional): default widget, or NULL if there is none.
//
func (popover *Popover) DefaultWidget() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_default_widget", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _widget Widgetter // out

	if _cret != nil {
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

// Modal returns whether the popover is modal, see gtk_popover_set_modal to see
// the implications of this.
//
// The function returns the following values:
//
//    - ok if popover is modal.
//
func (popover *Popover) Modal() bool {
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_modal", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument
	var _outs [1]girepository.Argument
	var _arg0 *C.void    // out
	var _out0 *C.void    // in
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_pointing_to", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out
	_out0 = *(**C.void)(unsafe.Pointer(&_outs[0]))

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer(_out0)))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// RelativeTo returns the widget popover is currently attached to.
//
// The function returns the following values:
//
//    - widget: Widget.
//
func (popover *Popover) RelativeTo() Widgetter {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_relative_to", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(popover)

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
	var _args [1]girepository.Argument
	var _arg0 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Popover").InvokeMethod("get_transitions_enabled", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

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
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Popover").InvokeMethod("popdown", _args[:], nil)

	runtime.KeepAlive(popover)
}

// Popup pops popover up. This is different than a gtk_widget_show() call in
// that it shows the popover with a transition. If you want to show the popover
// without a transition, use gtk_widget_show().
func (popover *Popover) Popup() {
	var _args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0

	girepository.MustFind("Gtk", "Popover").InvokeMethod("popup", _args[:], nil)

	runtime.KeepAlive(popover)
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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if widget != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(widget).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_default_widget", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if modal {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_modal", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	_arg1 = (*C.void)(gextras.StructNative(unsafe.Pointer(rect)))

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_pointing_to", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(rect)
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
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if relativeTo != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(relativeTo).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_relative_to", _args[:], nil)

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
	var _args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(popover).Native()))
	if transitionsEnabled {
		_arg1 = C.TRUE
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(*C.gboolean)(unsafe.Pointer(&_args[1])) = _arg1

	girepository.MustFind("Gtk", "Popover").InvokeMethod("set_transitions_enabled", _args[:], nil)

	runtime.KeepAlive(popover)
	runtime.KeepAlive(transitionsEnabled)
}
