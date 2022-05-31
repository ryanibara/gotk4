// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkpadcontroller.go.
var (
	GTypePadActionType = coreglib.Type(C.gtk_pad_action_type_get_type())
	GTypePadController = coreglib.Type(C.gtk_pad_controller_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypePadActionType, F: marshalPadActionType},
		{T: GTypePadController, F: marshalPadController},
	})
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

// PadControllerOverrider contains methods that are overridable.
type PadControllerOverrider interface {
}

// PadController: GtkPadController is an event controller for the pads found in
// drawing tablets.
//
// Pads are the collection of buttons and tactile sensors often found around the
// stylus-sensitive area.
//
// These buttons and sensors have no implicit meaning, and by default they
// perform no action. GtkPadController is provided to map those to GAction
// objects, thus letting the application give them a more semantic meaning.
//
// Buttons and sensors are not constrained to triggering a single action, some
// GDK_SOURCE_TABLET_PAD devices feature multiple "modes". All these input
// elements have one current mode, which may determine the final action being
// triggered.
//
// Pad devices often divide buttons and sensors into groups. All elements in a
// group share the same current mode, but different groups may have different
// modes. See gdk.DevicePad.GetNGroups() and gdk.DevicePad.GetGroupNModes().
//
// Each of the actions that a given button/strip/ring performs for a given mode
// is defined by a gtk.PadActionEntry. It contains an action name that will be
// looked up in the given GActionGroup and activated whenever the specified
// input element and mode are triggered.
//
// A simple example of GtkPadController usage: Assigning button 1 in all modes
// and pad devices to an "invert-selection" action:
//
//    GtkPadActionEntry *pad_actions[] = {
//      { GTK_PAD_ACTION_BUTTON, 1, -1, "Invert selection", "pad-actions.invert-selection" },
//      …
//    };
//
//    …
//    action_group = g_simple_action_group_new ();
//    action = g_simple_action_new ("pad-actions.invert-selection", NULL);
//    g_signal_connect (action, "activate", on_invert_selection_activated, NULL);
//    g_action_map_add_action (G_ACTION_MAP (action_group), action);
//    …
//    pad_controller = gtk_pad_controller_new (action_group, NULL);
//
//
// The actions belonging to rings/strips will be activated with a parameter of
// type G_VARIANT_TYPE_DOUBLE bearing the value of the given axis, it is
// required that those are made stateful and accepting this GVariantType.
type PadController struct {
	_ [0]func() // equal guard
	EventController
}

var (
	_ EventControllerer = (*PadController)(nil)
)

func classInitPadControllerer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

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

// NewPadController creates a new GtkPadController that will associate events
// from pad to actions.
//
// A NULL pad may be provided so the controller manages all pad devices
// generically, it is discouraged to mix GtkPadController objects with NULL and
// non-NULL pad argument on the same toplevel window, as execution order is not
// guaranteed.
//
// The GtkPadController is created with no mapped actions. In order to map pad
// events to actions, use gtk.PadController.SetActionEntries() or
// gtk.PadController.SetAction().
//
// Be aware that pad events will only be delivered to GtkWindows, so adding a
// pad controller to any other type of widget will not have an effect.
//
// The function takes the following parameters:
//
//    - group: GActionGroup to trigger actions from.
//    - pad (optional): GDK_SOURCE_TABLET_PAD device, or NULL to handle all pads.
//
// The function returns the following values:
//
//    - padController: newly created GtkPadController.
//
func NewPadController(group gio.ActionGrouper, pad gdk.Devicer) *PadController {
	var _args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(group).Native()))
	if pad != nil {
		_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pad).Native()))
	}

	*(**C.void)(unsafe.Pointer(&_args[0])) = _arg0
	*(**C.void)(unsafe.Pointer(&_args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "PadController").InvokeMethod("new_PadController", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

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
	native *C.GtkPadActionEntry
}
