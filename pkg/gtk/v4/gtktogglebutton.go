// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_ToggleButtonClass_toggled(GtkToggleButton*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButtonner},
	})
}

// ToggleButtonOverrider contains methods that are overridable.
type ToggleButtonOverrider interface {
	// Toggled emits the ::toggled signal on the GtkToggleButton.
	//
	// There is no good reason for an application ever to call this function.
	Toggled()
}

// ToggleButton: GtkToggleButton is a button which remains “pressed-in” when
// clicked.
//
// Clicking again will cause the toggle button to return to its normal state.
//
// A toggle button is created by calling either gtk.ToggleButton.New or
// gtk.ToggleButton.NewWithLabel. If using the former, it is advisable to pack a
// widget, (such as a GtkLabel and/or a GtkImage), into the toggle button’s
// container. (See gtk.Button for more information).
//
// The state of a GtkToggleButton can be set specifically using
// gtk.ToggleButton.SetActive(), and retrieved using
// gtk.ToggleButton.GetActive().
//
// To simply switch the state of a toggle button, use
// gtk.ToggleButton.Toggled().
//
//
// Grouping
//
// Toggle buttons can be grouped together, to form mutually exclusive groups -
// only one of the buttons can be toggled at a time, and toggling another one
// will switch the currently toggled one off.
//
// To add a GtkToggleButton to a group, use gtk.ToggleButton.SetGroup().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain GtkButton, it gets the .toggle style class.
//
// Creating two GtkToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data)
//    {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void)
//    {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new ();
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_box_append (GTK_BOX (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_box_append (GTK_BOX (box), toggle2);
//
//      gtk_window_set_child (GTK_WINDOW (window), box);
//      gtk_widget_show (window);
//    }.
type ToggleButton struct {
	_ [0]func() // equal guard
	Button
}

var (
	_ Widgetter           = (*ToggleButton)(nil)
	_ externglib.Objector = (*ToggleButton)(nil)
)

func classInitToggleButtonner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkToggleButtonClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkToggleButtonClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Toggled() }); ok {
		pclass.toggled = (*[0]byte)(C._gotk4_gtk4_ToggleButtonClass_toggled)
	}
}

//export _gotk4_gtk4_ToggleButtonClass_toggled
func _gotk4_gtk4_ToggleButtonClass_toggled(arg0 *C.GtkToggleButton) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Toggled() })

	iface.Toggled()
}

func wrapToggleButton(obj *externglib.Object) *ToggleButton {
	return &ToggleButton{
		Button: Button{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
			Object: obj,
			Actionable: Actionable{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
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
			},
		},
	}
}

func marshalToggleButtonner(p uintptr) (interface{}, error) {
	return wrapToggleButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectToggled: emitted whenever the GtkToggleButton's state is changed.
func (toggleButton *ToggleButton) ConnectToggled(f func()) externglib.SignalHandle {
	return toggleButton.Connect("toggled", externglib.GeneratedClosure{Func: f})
}

// NewToggleButton creates a new toggle button.
//
// A widget should be packed into the button, as in gtk.Button.New.
//
// The function returns the following values:
//
//    - toggleButton: new toggle button.
//
func NewToggleButton() *ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithLabel creates a new toggle button with a text label.
//
// The function takes the following parameters:
//
//    - label: string containing the message to be placed in the toggle button.
//
// The function returns the following values:
//
//    - toggleButton: new toggle button.
//
func NewToggleButtonWithLabel(label string) *ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// NewToggleButtonWithMnemonic creates a new GtkToggleButton containing a label.
//
// The label will be created using gtk.Label.NewWithMnemonic, so underscores in
// label indicate the mnemonic for the button.
//
// The function takes the following parameters:
//
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - toggleButton: new GtkToggleButton.
//
func NewToggleButtonWithMnemonic(label string) *ToggleButton {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _toggleButton *ToggleButton // out

	_toggleButton = wrapToggleButton(externglib.Take(unsafe.Pointer(_cret)))

	return _toggleButton
}

// Active queries a GtkToggleButton and returns its current state.
//
// Returns TRUE if the toggle button is pressed in and FALSE if it is raised.
//
// The function returns the following values:
//
//    - ok: whether the button is pressed.
//
func (toggleButton *ToggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)
	runtime.KeepAlive(toggleButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetActive sets the status of the toggle button.
//
// Set to TRUE if you want the GtkToggleButton to be “pressed in”, and FALSE to
// raise it.
//
// If the status of the button changes, this action causes the
// gtktogglebutton::toggled signal to be emitted.
//
// The function takes the following parameters:
//
//    - isActive: TRUE or FALSE.
//
func (toggleButton *ToggleButton) SetActive(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(isActive)
}

// SetGroup adds self to the group of group.
//
// In a group of multiple toggle buttons, only one button can be active at a
// time.
//
// Setting up groups in a cycle leads to undefined behavior.
//
// Note that the same effect can be achieved via the gtk.Actionable API, by
// using the same action with parameter type and state type 's' for all buttons
// in the group, and giving each button its own target value.
//
// The function takes the following parameters:
//
//    - group (optional): another GtkToggleButton to form a group with.
//
func (toggleButton *ToggleButton) SetGroup(group *ToggleButton) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))
	if group != nil {
		_arg1 = (*C.GtkToggleButton)(unsafe.Pointer(group.Native()))
	}

	C.gtk_toggle_button_set_group(_arg0, _arg1)
	runtime.KeepAlive(toggleButton)
	runtime.KeepAlive(group)
}

// Toggled emits the ::toggled signal on the GtkToggleButton.
//
// There is no good reason for an application ever to call this function.
func (toggleButton *ToggleButton) Toggled() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(toggleButton.Native()))

	C.gtk_toggle_button_toggled(_arg0)
	runtime.KeepAlive(toggleButton)
}
