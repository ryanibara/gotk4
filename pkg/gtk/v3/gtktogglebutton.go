// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButton},
	})
}

// ToggleButton: a ToggleButton is a Button which will remain “pressed-in” when
// clicked. Clicking again will cause the toggle button to return to its normal
// state.
//
// A toggle button is created by calling either gtk_toggle_button_new() or
// gtk_toggle_button_new_with_label(). If using the former, it is advisable to
// pack a widget, (such as a Label and/or a Image), into the toggle button’s
// container. (See Button for more information).
//
// The state of a ToggleButton can be set specifically using
// gtk_toggle_button_set_active(), and retrieved using
// gtk_toggle_button_get_active().
//
// To simply switch the state of a toggle button, use
// gtk_toggle_button_toggled().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .toggle style class.
//
// Creating two ToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data) {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void) {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      // Makes this toggle button invisible
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle1),
//                                  TRUE);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle2),
//                                  FALSE);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle2);
//
//      gtk_container_add (GTK_CONTAINER (window), box);
//      gtk_widget_show_all (window);
//    }
type ToggleButton interface {
	Button
	Actionable
	Activatable
	Buildable

	// Active queries a ToggleButton and returns its current state. Returns true
	// if the toggle button is pressed in and false if it is raised.
	Active() bool
	// Inconsistent gets the value set by gtk_toggle_button_set_inconsistent().
	Inconsistent() bool
	// Mode retrieves whether the button is displayed as a separate indicator
	// and label. See gtk_toggle_button_set_mode().
	Mode() bool
	// SetActive sets the status of the toggle button. Set to true if you want
	// the GtkToggleButton to be “pressed in”, and false to raise it. This
	// action causes the ToggleButton::toggled signal and the Button::clicked
	// signal to be emitted.
	SetActive(isActive bool)
	// SetInconsistent: if the user has selected a range of elements (such as
	// some text or spreadsheet cells) that are affected by a toggle button, and
	// the current values in that range are inconsistent, you may want to
	// display the toggle in an “in between” state. This function turns on “in
	// between” display. Normally you would turn off the inconsistent state
	// again if the user toggles the toggle button. This has to be done
	// manually, gtk_toggle_button_set_inconsistent() only affects visual
	// appearance, it doesn’t affect the semantics of the button.
	SetInconsistent(setting bool)
	// SetMode sets whether the button is displayed as a separate indicator and
	// label. You can call this function on a checkbutton or a radiobutton with
	// @draw_indicator = false to make the button look like a normal button.
	//
	// This can be used to create linked strip of buttons that work like a
	// StackSwitcher.
	//
	// This function only affects instances of classes like CheckButton and
	// RadioButton that derive from ToggleButton, not instances of ToggleButton
	// itself.
	SetMode(drawIndicator bool)
	// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There
	// is no good reason for an application ever to call this function.
	Toggled()
}

// toggleButton implements the ToggleButton interface.
type toggleButton struct {
	Button
	Actionable
	Activatable
	Buildable
}

var _ ToggleButton = (*toggleButton)(nil)

// WrapToggleButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleButton(obj *externglib.Object) ToggleButton {
	return ToggleButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleButton(obj), nil
}

// NewToggleButton constructs a class ToggleButton.
func NewToggleButton() ToggleButton {
	var cret C.GtkToggleButton
	var ret1 ToggleButton

	cret = C.gtk_toggle_button_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToggleButton)

	return ret1
}

// NewToggleButtonWithLabel constructs a class ToggleButton.
func NewToggleButtonWithLabel(label string) ToggleButton {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkToggleButton
	var ret1 ToggleButton

	cret = C.gtk_toggle_button_new_with_label(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToggleButton)

	return ret1
}

// NewToggleButtonWithMnemonic constructs a class ToggleButton.
func NewToggleButtonWithMnemonic(label string) ToggleButton {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkToggleButton
	var ret1 ToggleButton

	cret = C.gtk_toggle_button_new_with_mnemonic(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToggleButton)

	return ret1
}

// Active queries a ToggleButton and returns its current state. Returns true
// if the toggle button is pressed in and false if it is raised.
func (t toggleButton) Active() bool {
	var arg0 *C.GtkToggleButton

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_toggle_button_get_active(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Inconsistent gets the value set by gtk_toggle_button_set_inconsistent().
func (t toggleButton) Inconsistent() bool {
	var arg0 *C.GtkToggleButton

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_toggle_button_get_inconsistent(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Mode retrieves whether the button is displayed as a separate indicator
// and label. See gtk_toggle_button_set_mode().
func (t toggleButton) Mode() bool {
	var arg0 *C.GtkToggleButton

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_toggle_button_get_mode(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetActive sets the status of the toggle button. Set to true if you want
// the GtkToggleButton to be “pressed in”, and false to raise it. This
// action causes the ToggleButton::toggled signal and the Button::clicked
// signal to be emitted.
func (t toggleButton) SetActive(isActive bool) {
	var arg0 *C.GtkToggleButton
	var arg1 C.gboolean

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if isActive {
		arg1 = C.gboolean(1)
	}

	C.gtk_toggle_button_set_active(arg0, isActive)
}

// SetInconsistent: if the user has selected a range of elements (such as
// some text or spreadsheet cells) that are affected by a toggle button, and
// the current values in that range are inconsistent, you may want to
// display the toggle in an “in between” state. This function turns on “in
// between” display. Normally you would turn off the inconsistent state
// again if the user toggles the toggle button. This has to be done
// manually, gtk_toggle_button_set_inconsistent() only affects visual
// appearance, it doesn’t affect the semantics of the button.
func (t toggleButton) SetInconsistent(setting bool) {
	var arg0 *C.GtkToggleButton
	var arg1 C.gboolean

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_toggle_button_set_inconsistent(arg0, setting)
}

// SetMode sets whether the button is displayed as a separate indicator and
// label. You can call this function on a checkbutton or a radiobutton with
// @draw_indicator = false to make the button look like a normal button.
//
// This can be used to create linked strip of buttons that work like a
// StackSwitcher.
//
// This function only affects instances of classes like CheckButton and
// RadioButton that derive from ToggleButton, not instances of ToggleButton
// itself.
func (t toggleButton) SetMode(drawIndicator bool) {
	var arg0 *C.GtkToggleButton
	var arg1 C.gboolean

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if drawIndicator {
		arg1 = C.gboolean(1)
	}

	C.gtk_toggle_button_set_mode(arg0, drawIndicator)
}

// Toggled emits the ToggleButton::toggled signal on the ToggleButton. There
// is no good reason for an application ever to call this function.
func (t toggleButton) Toggled() {
	var arg0 *C.GtkToggleButton

	arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	C.gtk_toggle_button_toggled(arg0)
}

type ToggleButtonPrivate struct {
	native C.GtkToggleButtonPrivate
}

// WrapToggleButtonPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToggleButtonPrivate(ptr unsafe.Pointer) *ToggleButtonPrivate {
	if ptr == nil {
		return nil
	}

	return (*ToggleButtonPrivate)(ptr)
}

func marshalToggleButtonPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapToggleButtonPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ToggleButtonPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
