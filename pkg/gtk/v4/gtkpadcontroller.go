// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_pad_controller_get_type()), F: marshalPadController},
	})
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
type PadController interface {
	EventController

	// SetAction adds an individual action to @controller.
	//
	// This action will only be activated if the given button/ring/strip number
	// in @index is interacted while the current mode is @mode. -1 may be used
	// for simple cases, so the action is triggered on all modes.
	//
	// The given @label should be considered user-visible, so
	// internationalization rules apply. Some windowing systems may be able to
	// use those for user feedback.
	SetAction(typ PadActionType, index int, mode int, label string, actionName string)
	// SetActionEntries: a convenience function to add a group of action entries
	// on @controller.
	//
	// See [struct@Gtk.PadActionEntry] and
	// [method@Gtk.PadController.set_action].
	SetActionEntries(entries []PadActionEntry)
}

// padController implements the PadController interface.
type padController struct {
	EventController
}

var _ PadController = (*padController)(nil)

// WrapPadController wraps a GObject to the right type. It is
// primarily used internally.
func WrapPadController(obj *externglib.Object) PadController {
	return PadController{
		EventController: WrapEventController(obj),
	}
}

func marshalPadController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPadController(obj), nil
}

// NewPadController constructs a class PadController.
func NewPadController(group gio.ActionGroup, pad gdk.Device) PadController {
	var arg1 *C.GActionGroup
	var arg2 *C.GdkDevice

	arg1 = (*C.GActionGroup)(group.Native())
	arg2 = (*C.GdkDevice)(pad.Native())

	ret := C.gtk_pad_controller_new(arg1, arg2)

	var ret0 PadController

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(PadController)

	return ret0
}

// SetAction adds an individual action to @controller.
//
// This action will only be activated if the given button/ring/strip number
// in @index is interacted while the current mode is @mode. -1 may be used
// for simple cases, so the action is triggered on all modes.
//
// The given @label should be considered user-visible, so
// internationalization rules apply. Some windowing systems may be able to
// use those for user feedback.
func (c padController) SetAction(typ PadActionType, index int, mode int, label string, actionName string) {
	var arg0 *C.GtkPadController
	var arg1 C.GtkPadActionType
	var arg2 C.int
	var arg3 C.int
	var arg4 *C.char
	var arg5 *C.char

	arg0 = (*C.GtkPadController)(c.Native())
	arg1 = (C.GtkPadActionType)(typ)
	arg2 = C.int(index)
	arg3 = C.int(mode)
	arg4 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(arg5))

	C.gtk_pad_controller_set_action(arg0, arg1, arg2, arg3, arg4, arg5)
}

// SetActionEntries: a convenience function to add a group of action entries
// on @controller.
//
// See [struct@Gtk.PadActionEntry] and
// [method@Gtk.PadController.set_action].
func (c padController) SetActionEntries(entries []PadActionEntry) {
	var arg0 *C.GtkPadController
	var arg1 *C.GtkPadActionEntry
	var arg2 C.int

	arg0 = (*C.GtkPadController)(c.Native())
	{
		var dst []C.GtkPadActionEntry
		ptr := C.malloc(C.sizeof_GtkPadActionEntry * len(entries))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(entries)
		sliceHeader.Cap = len(entries)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(entries); i++ {
			src := entries[i]
			dst[i] = (C.GtkPadActionEntry)(src.Native())
		}

		arg1 = (*C.GtkPadActionEntry)(unsafe.Pointer(ptr))
		arg2 = len(entries)
	}

	C.gtk_pad_controller_set_action_entries(arg0, arg1, arg2)
}

// PadActionEntry: struct defining a pad action entry.
type PadActionEntry struct {
	native C.GtkPadActionEntry
}

// WrapPadActionEntry wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPadActionEntry(ptr unsafe.Pointer) *PadActionEntry {
	if ptr == nil {
		return nil
	}

	return (*PadActionEntry)(ptr)
}

func marshalPadActionEntry(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPadActionEntry(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PadActionEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Type gets the field inside the struct.
func (p *PadActionEntry) Type() PadActionType {
	var ret PadActionType
	ret = PadActionType(p.native._type)
	return ret
}

// Index gets the field inside the struct.
func (p *PadActionEntry) Index() int {
	var ret int
	ret = int(p.native.index)
	return ret
}

// Mode gets the field inside the struct.
func (p *PadActionEntry) Mode() int {
	var ret int
	ret = int(p.native.mode)
	return ret
}

// Label gets the field inside the struct.
func (p *PadActionEntry) Label() string {
	var ret string
	ret = C.GoString(p.native.label)
	return ret
}

// ActionName gets the field inside the struct.
func (p *PadActionEntry) ActionName() string {
	var ret string
	ret = C.GoString(p.native.action_name)
	return ret
}
