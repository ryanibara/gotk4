// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_pad_action_type_get_type()), F: marshalPadActionType},
		{T: externglib.Type(C.gtk_pad_controller_get_type()), F: marshalPadControllerer},
	})
}

// PadActionType: type of a pad action.
type PadActionType int

const (
	// Button: action is triggered by a pad button
	PadActionTypeButton PadActionType = iota
	// Ring: action is triggered by a pad ring
	PadActionTypeRing
	// Strip: action is triggered by a pad strip
	PadActionTypeStrip
)

func marshalPadActionType(p uintptr) (interface{}, error) {
	return PadActionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PadControllerer describes PadController's methods.
type PadControllerer interface {
	// SetActionEntries: convenience function to add a group of action entries
	// on @controller.
	SetActionEntries(entries []PadActionEntry)
}

// PadController: `GtkPadController` is an event controller for the pads found
// in drawing tablets.
//
// Pads are the collection of buttons and tactile sensors often found around the
// stylus-sensitive area.
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action. `GtkPadController` is provided to map those to `GAction`
// objects, thus letting the application give them a more semantic meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// GDK_SOURCE_TABLET_PAD devices feature multiple "modes". All these input
// elements have one current mode, which may determine the final action being
// triggered.
//
// Pad devices often divide buttons and sensors into groups. All elements in a
// group share the same current mode, but different groups may have different
// modes. See [method@Gdk.DevicePad.get_n_groups] and
// [method@Gdk.DevicePad.get_group_n_modes].
//
// Each of the actions that a given button/strip/ring performs for a given mode
// is defined by a [struct@Gtk.PadActionEntry]. It contains an action name that
// will be looked up in the given `GActionGroup` and activated whenever the
// specified input element and mode are triggered.
//
// A simple example of `GtkPadController` usage: Assigning button 1 in all modes
// and pad devices to an "invert-selection" action:
//
// “`c GtkPadActionEntry *pad_actions[] = { { GTK_PAD_ACTION_BUTTON, 1, -1,
// "Invert selection", "pad-actions.invert-selection" }, … };
//
// … action_group = g_simple_action_group_new (); action = g_simple_action_new
// ("pad-actions.invert-selection", NULL); g_signal_connect (action, "activate",
// on_invert_selection_activated, NULL); g_action_map_add_action (G_ACTION_MAP
// (action_group), action); … pad_controller = gtk_pad_controller_new
// (action_group, NULL); “`
//
// The actions belonging to rings/strips will be activated with a parameter of
// type G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it is
// required that those are made stateful and accepting this `GVariantType`.
type PadController struct {
	EventController
}

var (
	_ PadControllerer = (*PadController)(nil)
	_ gextras.Nativer = (*PadController)(nil)
)

func wrapPadController(obj *externglib.Object) PadControllerer {
	return &PadController{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalPadControllerer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPadController(obj), nil
}

// NewPadController creates a new `GtkPadController` that will associate events
// from @pad to actions.
//
// A nil pad may be provided so the controller manages all pad devices
// generically, it is discouraged to mix `GtkPadController` objects with nil and
// non-nil @pad argument on the same toplevel window, as execution order is not
// guaranteed.
//
// The `GtkPadController` is created with no mapped actions. In order to map pad
// events to actions, use [method@Gtk.PadController.set_action_entries] or
// [method@Gtk.PadController.set_action].
//
// Be aware that pad events will only be delivered to `GtkWindow`s, so adding a
// pad controller to any other type of widget will not have an effect.
func NewPadController(group gio.ActionGrouper, pad gdk.Devicer) *PadController {
	var _arg1 *C.GActionGroup     // out
	var _arg2 *C.GdkDevice        // out
	var _cret *C.GtkPadController // in

	_arg1 = (*C.GActionGroup)(unsafe.Pointer((group).(gextras.Nativer).Native()))
	_arg2 = (*C.GdkDevice)(unsafe.Pointer((pad).(gextras.Nativer).Native()))

	_cret = C.gtk_pad_controller_new(_arg1, _arg2)

	var _padController *PadController // out

	_padController = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PadController)

	return _padController
}

// SetActionEntries: convenience function to add a group of action entries on
// @controller.
//
// See [struct@Gtk.PadActionEntry] and [method@Gtk.PadController.set_action].
func (controller *PadController) SetActionEntries(entries []PadActionEntry) {
	var _arg0 *C.GtkPadController // out
	var _arg1 *C.GtkPadActionEntry
	var _arg2 C.int

	_arg0 = (*C.GtkPadController)(unsafe.Pointer(controller.Native()))
	_arg2 = C.int(len(entries))
	_arg1 = (*C.GtkPadActionEntry)(unsafe.Pointer(&entries[0]))

	C.gtk_pad_controller_set_action_entries(_arg0, _arg1, _arg2)
}

// PadActionEntry: struct defining a pad action entry.
type PadActionEntry struct {
	native C.GtkPadActionEntry
}

// Native returns the underlying C source pointer.
func (p *PadActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}
