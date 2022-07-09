// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypePadActionType returns the GType for the type PadActionType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePadActionType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PadActionType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPadActionType)
	return gtype
}

// GTypePadController returns the GType for the type PadController.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypePadController() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "PadController").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalPadController)
	return gtype
}

// PadActionType: type of a pad action.
type PadActionType C.gint

const (
	// PadActionButton: action is triggered by a pad button.
	PadActionButton PadActionType = iota
	// PadActionRing: action is triggered by a pad ring.
	PadActionRing
	// PadActionStrip: action is triggered by a pad strip.
	PadActionStrip
)

func marshalPadActionType(p uintptr) (interface{}, error) {
	return PadActionType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for PadActionType.
func (p PadActionType) String() string {
	switch p {
	case PadActionButton:
		return "Button"
	case PadActionRing:
		return "Ring"
	case PadActionStrip:
		return "Strip"
	default:
		return fmt.Sprintf("PadActionType(%d)", p)
	}
}

// PadController is an event controller for the pads found in drawing tablets
// (The collection of buttons and tactile sensors often found around the
// stylus-sensitive area).
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action, this event controller is provided to map those to #GAction
// objects, thus letting the application give those a more semantic meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// GDK_SOURCE_TABLET_PAD devices feature multiple "modes", all these input
// elements have one current mode, which may determine the final action being
// triggered. Pad devices often divide buttons and sensors into groups, all
// elements in a group share the same current mode, but different groups may
// have different modes. See gdk_device_pad_get_n_groups() and
// gdk_device_pad_get_group_n_modes().
//
// Each of the actions that a given button/strip/ring performs for a given mode
// is defined by PadActionEntry, it contains an action name that will be looked
// up in the given Group and activated whenever the specified input element and
// mode are triggered.
//
// A simple example of PadController usage, assigning button 1 in all modes and
// pad devices to an "invert-selection" action:
//
//      GtkPadActionEntry *pad_actions[] = {
//        { GTK_PAD_ACTION_BUTTON, 1, -1, "Invert selection", "pad-actions.invert-selection" },
//        …
//      };
//
//      …
//      action_group = g_simple_action_group_new ();
//      action = g_simple_action_new ("pad-actions.invert-selection", NULL);
//      g_signal_connect (action, "activate", on_invert_selection_activated, NULL);
//      g_action_map_add_action (G_ACTION_MAP (action_group), action);
//      …
//      pad_controller = gtk_pad_controller_new (window, action_group, NULL);
//
// The actions belonging to rings/strips will be activated with a parameter of
// type G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it is
// required that those are made stateful and accepting this Type.
type PadController struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*PadController)(nil)
)

func wrapPadController(obj *coreglib.Object) *PadController {
	return &PadController{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalPadController(p uintptr) (interface{}, error) {
	return wrapPadController(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewPadController creates a new PadController that will associate events from
// pad to actions. A NULL pad may be provided so the controller manages all pad
// devices generically, it is discouraged to mix PadController objects with NULL
// and non-NULL pad argument on the same window, as execution order is not
// guaranteed.
//
// The PadController is created with no mapped actions. In order to map pad
// events to actions, use gtk_pad_controller_set_action_entries() or
// gtk_pad_controller_set_action().
//
// The function takes the following parameters:
//
//    - window: Window.
//    - group to trigger actions from.
//    - pad (optional): GDK_SOURCE_TABLET_PAD device, or NULL to handle all pads.
//
// The function returns the following values:
//
//    - padController: newly created PadController.
//
func NewPadController(window *Window, group gio.ActionGrouper, pad gdk.Devicer) *PadController {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	if pad != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pad).Native()))
	}

	_info := girepository.MustFind("Gtk", "PadController")
	_gret := _info.InvokeClassMethod("new_PadController", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(window)
	runtime.KeepAlive(group)
	runtime.KeepAlive(pad)

	var _padController *PadController // out

	_padController = wrapPadController(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _padController
}

// PadActionEntry: struct defining a pad action entry.
//
// An instance of this type is always passed by reference.
type PadActionEntry struct {
	*padActionEntry
}

// padActionEntry is the struct that's finalized.
type padActionEntry struct {
	native unsafe.Pointer
}
