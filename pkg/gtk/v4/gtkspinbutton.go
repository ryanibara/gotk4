// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_spin_button_update_policy_get_type()), F: marshalSpinButtonUpdatePolicy},
		{T: externglib.Type(C.gtk_spin_type_get_type()), F: marshalSpinType},
		{T: externglib.Type(C.gtk_spin_button_get_type()), F: marshalSpinButtonner},
	})
}

// INPUT_ERROR: constant to return from a signal handler for the ::input signal
// in case of conversion failure.
//
// See gtk.SpinButton::input.
const INPUT_ERROR = -1

// SpinButtonUpdatePolicy determines whether the spin button displays values
// outside the adjustment bounds.
//
// See gtk.SpinButton.SetUpdatePolicy().
type SpinButtonUpdatePolicy C.gint

const (
	// UpdateAlways: when refreshing your SpinButton, the value is always
	// displayed.
	UpdateAlways SpinButtonUpdatePolicy = iota
	// UpdateIfValid: when refreshing your SpinButton, the value is only
	// displayed if it is valid within the bounds of the spin button's
	// adjustment.
	UpdateIfValid
)

func marshalSpinButtonUpdatePolicy(p uintptr) (interface{}, error) {
	return SpinButtonUpdatePolicy(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinButtonUpdatePolicy.
func (s SpinButtonUpdatePolicy) String() string {
	switch s {
	case UpdateAlways:
		return "Always"
	case UpdateIfValid:
		return "IfValid"
	default:
		return fmt.Sprintf("SpinButtonUpdatePolicy(%d)", s)
	}
}

// SpinType values of the GtkSpinType enumeration are used to specify the change
// to make in gtk_spin_button_spin().
type SpinType C.gint

const (
	// SpinStepForward: increment by the adjustments step increment.
	SpinStepForward SpinType = iota
	// SpinStepBackward: decrement by the adjustments step increment.
	SpinStepBackward
	// SpinPageForward: increment by the adjustments page increment.
	SpinPageForward
	// SpinPageBackward: decrement by the adjustments page increment.
	SpinPageBackward
	// SpinHome: go to the adjustments lower bound.
	SpinHome
	// SpinEnd: go to the adjustments upper bound.
	SpinEnd
	// SpinUserDefined: change by a specified amount.
	SpinUserDefined
)

func marshalSpinType(p uintptr) (interface{}, error) {
	return SpinType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for SpinType.
func (s SpinType) String() string {
	switch s {
	case SpinStepForward:
		return "StepForward"
	case SpinStepBackward:
		return "StepBackward"
	case SpinPageForward:
		return "PageForward"
	case SpinPageBackward:
		return "PageBackward"
	case SpinHome:
		return "Home"
	case SpinEnd:
		return "End"
	case SpinUserDefined:
		return "UserDefined"
	default:
		return fmt.Sprintf("SpinType(%d)", s)
	}
}

// SpinButton: GtkSpinButton is an ideal way to allow the user to set the value
// of some attribute.
//
// !An example GtkSpinButton (spinbutton.png)
//
// Rather than having to directly type a number into a GtkEntry, GtkSpinButton
// allows the user to click on one of two arrows to increment or decrement the
// displayed value. A value can still be typed in, with the bonus that it can be
// checked to ensure it is in a given range.
//
// The main properties of a GtkSpinButton are through an adjustment. See the
// gtk.Adjustment documentation for more details about an adjustment's
// properties.
//
// Note that GtkSpinButton will by default make its entry large enough to
// accommodate the lower and upper bounds of the adjustment. If this is not
// desired, the automatic sizing can be turned off by explicitly setting
// gtk.Editable:width-chars to a value != -1.
//
// Using a GtkSpinButton to get an integer
//
//    // Provides a function to retrieve an integer value from a GtkSpinButton
//    // and creates a spin button to model percentage values.
//
//    int
//    grab_int_value (GtkSpinButton *button,
//                    gpointer       user_data)
//    {
//      return gtk_spin_button_get_value_as_int (button);
//    }
//
//    void
//    create_integer_spin_button (void)
//    {
//
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (50.0, 0.0, 100.0, 1.0, 5.0, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with no decimal places
//      button = gtk_spin_button_new (adjustment, 1.0, 0);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// Using a GtkSpinButton to get a floating point value
//
//    // Provides a function to retrieve a floating point value from a
//    // GtkSpinButton, and creates a high precision spin button.
//
//    float
//    grab_float_value (GtkSpinButton *button,
//                      gpointer       user_data)
//    {
//      return gtk_spin_button_get_value (button);
//    }
//
//    void
//    create_floating_spin_button (void)
//    {
//      GtkWidget *window, *button;
//      GtkAdjustment *adjustment;
//
//      adjustment = gtk_adjustment_new (2.500, 0.0, 5.0, 0.001, 0.1, 0.0);
//
//      window = gtk_window_new ();
//
//      // creates the spinbutton, with three decimal places
//      button = gtk_spin_button_new (adjustment, 0.001, 3);
//      gtk_window_set_child (GTK_WINDOW (window), button);
//
//      gtk_widget_show (window);
//    }
//
//
// CSS nodes
//
//    spinbutton.horizontal
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ├── button.down
//    ╰── button.up
//
//
//
//
//    spinbutton.vertical
//    ├── button.up
//    ├── text
//    │    ├── undershoot.left
//    │    ╰── undershoot.right
//    ╰── button.down
//
//
// GtkSpinButtons main CSS node has the name spinbutton. It creates subnodes for
// the entry and the two buttons, with these names. The button nodes have the
// style classes .up and .down. The GtkText subnodes (if present) are put below
// the text node. The orientation of the spin button is reflected in the
// .vertical or .horizontal style class on the main node.
//
//
// Accessiblity
//
// GtkSpinButton uses the GTK_ACCESSIBLE_ROLE_SPIN_BUTTON role.
type SpinButton struct {
	Widget

	CellEditable
	Editable
	Orientable
	*externglib.Object
}

func wrapSpinButton(obj *externglib.Object) *SpinButton {
	return &SpinButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		CellEditable: CellEditable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSpinButtonner(p uintptr) (interface{}, error) {
	return wrapSpinButton(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSpinButton creates a new GtkSpinButton.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment that this spin button should use, or NULL.
//    - climbRate specifies by how much the rate of change in the value will
//    accelerate if you continue to hold down an up/down button or arrow key.
//    - digits: number of decimal places to display.
//
func NewSpinButton(adjustment *Adjustment, climbRate float64, digits uint) *SpinButton {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.double         // out
	var _arg3 C.guint          // out
	var _cret *C.GtkWidget     // in

	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}
	_arg2 = C.double(climbRate)
	_arg3 = C.guint(digits)

	_cret = C.gtk_spin_button_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(externglib.Take(unsafe.Pointer(_cret)))

	return _spinButton
}

// NewSpinButtonWithRange creates a new GtkSpinButton with the given properties.
//
// This is a convenience constructor that allows creation of a numeric
// GtkSpinButton without manually creating an adjustment. The value is initially
// set to the minimum value and a page increment of 10 * step is the default.
// The precision of the spin button is equivalent to the precision of step.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk.SpinButton.SetDigits() to correct it.
//
// The function takes the following parameters:
//
//    - min: minimum allowable value.
//    - max: maximum allowable value.
//    - step: increment added or subtracted by spinning the widget.
//
func NewSpinButtonWithRange(min, max, step float64) *SpinButton {
	var _arg1 C.double     // out
	var _arg2 C.double     // out
	var _arg3 C.double     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.double(min)
	_arg2 = C.double(max)
	_arg3 = C.double(step)

	_cret = C.gtk_spin_button_new_with_range(_arg1, _arg2, _arg3)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _spinButton *SpinButton // out

	_spinButton = wrapSpinButton(externglib.Take(unsafe.Pointer(_cret)))

	return _spinButton
}

// Configure changes the properties of an existing spin button.
//
// The adjustment, climb rate, and number of decimal places are updated
// accordingly.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment to replace the spin button’s existing
//    adjustment, or NULL to leave its current adjustment unchanged.
//    - climbRate: new climb rate.
//    - digits: number of decimal places to display in the spin button.
//
func (spinButton *SpinButton) Configure(adjustment *Adjustment, climbRate float64, digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out
	var _arg2 C.double         // out
	var _arg3 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	if adjustment != nil {
		_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))
	}
	_arg2 = C.double(climbRate)
	_arg3 = C.guint(digits)

	C.gtk_spin_button_configure(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
	runtime.KeepAlive(climbRate)
	runtime.KeepAlive(digits)
}

// Adjustment: get the adjustment associated with a GtkSpinButton.
func (spinButton *SpinButton) Adjustment() *Adjustment {
	var _arg0 *C.GtkSpinButton // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_adjustment(_arg0)
	runtime.KeepAlive(spinButton)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// ClimbRate returns the acceleration rate for repeated changes.
func (spinButton *SpinButton) ClimbRate() float64 {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_climb_rate(_arg0)
	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Digits fetches the precision of spin_button.
func (spinButton *SpinButton) Digits() uint {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.guint          // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_digits(_arg0)
	runtime.KeepAlive(spinButton)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Increments gets the current step and page the increments used by spin_button.
//
// See gtk.SpinButton.SetIncrements().
func (spinButton *SpinButton) Increments() (step float64, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // in
	var _arg2 C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	C.gtk_spin_button_get_increments(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(spinButton)

	var _step float64 // out
	var _page float64 // out

	_step = float64(_arg1)
	_page = float64(_arg2)

	return _step, _page
}

// Numeric returns whether non-numeric text can be typed into the spin button.
func (spinButton *SpinButton) Numeric() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_numeric(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Range gets the range allowed for spin_button.
//
// See gtk.SpinButton.SetRange().
func (spinButton *SpinButton) Range() (min float64, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // in
	var _arg2 C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	C.gtk_spin_button_get_range(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(spinButton)

	var _min float64 // out
	var _max float64 // out

	_min = float64(_arg1)
	_max = float64(_arg2)

	return _min, _max
}

// SnapToTicks returns whether the values are corrected to the nearest step.
func (spinButton *SpinButton) SnapToTicks() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_snap_to_ticks(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UpdatePolicy gets the update behavior of a spin button.
//
// See gtk.SpinButton.SetUpdatePolicy().
func (spinButton *SpinButton) UpdatePolicy() SpinButtonUpdatePolicy {
	var _arg0 *C.GtkSpinButton            // out
	var _cret C.GtkSpinButtonUpdatePolicy // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_update_policy(_arg0)
	runtime.KeepAlive(spinButton)

	var _spinButtonUpdatePolicy SpinButtonUpdatePolicy // out

	_spinButtonUpdatePolicy = SpinButtonUpdatePolicy(_cret)

	return _spinButtonUpdatePolicy
}

// Value: get the value in the spin_button.
func (spinButton *SpinButton) Value() float64 {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.double         // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_value(_arg0)
	runtime.KeepAlive(spinButton)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// ValueAsInt: get the value spin_button represented as an integer.
func (spinButton *SpinButton) ValueAsInt() int {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.int            // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_value_as_int(_arg0)
	runtime.KeepAlive(spinButton)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Wrap returns whether the spin button’s value wraps around to the opposite
// limit when the upper or lower limit of the range is exceeded.
func (spinButton *SpinButton) Wrap() bool {
	var _arg0 *C.GtkSpinButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	_cret = C.gtk_spin_button_get_wrap(_arg0)
	runtime.KeepAlive(spinButton)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAdjustment replaces the GtkAdjustment associated with spin_button.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment to replace the existing adjustment.
//
func (spinButton *SpinButton) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_spin_button_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(adjustment)
}

// SetClimbRate sets the acceleration rate for repeated changes when you hold
// down a button or key.
//
// The function takes the following parameters:
//
//    - climbRate: rate of acceleration, must be >= 0.
//
func (spinButton *SpinButton) SetClimbRate(climbRate float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.double(climbRate)

	C.gtk_spin_button_set_climb_rate(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(climbRate)
}

// SetDigits: set the precision to be displayed by spin_button.
//
// Up to 20 digit precision is allowed.
//
// The function takes the following parameters:
//
//    - digits: number of digits after the decimal point to be displayed for
//    the spin button’s value.
//
func (spinButton *SpinButton) SetDigits(digits uint) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.guint(digits)

	C.gtk_spin_button_set_digits(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(digits)
}

// SetIncrements sets the step and page increments for spin_button.
//
// This affects how quickly the value changes when the spin button’s arrows are
// activated.
//
// The function takes the following parameters:
//
//    - step: increment applied for a button 1 press.
//    - page: increment applied for a button 2 press.
//
func (spinButton *SpinButton) SetIncrements(step, page float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.double(step)
	_arg2 = C.double(page)

	C.gtk_spin_button_set_increments(_arg0, _arg1, _arg2)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetNumeric sets the flag that determines if non-numeric text can be typed
// into the spin button.
//
// The function takes the following parameters:
//
//    - numeric: flag indicating if only numeric entry is allowed.
//
func (spinButton *SpinButton) SetNumeric(numeric bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	if numeric {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_numeric(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(numeric)
}

// SetRange sets the minimum and maximum allowable values for spin_button.
//
// If the current value is outside this range, it will be adjusted to fit within
// the range, otherwise it will remain unchanged.
//
// The function takes the following parameters:
//
//    - min: minimum allowable value.
//    - max: maximum allowable value.
//
func (spinButton *SpinButton) SetRange(min, max float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.double(min)
	_arg2 = C.double(max)

	C.gtk_spin_button_set_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

// SetSnapToTicks sets the policy as to whether values are corrected to the
// nearest step increment when a spin button is activated after providing an
// invalid value.
//
// The function takes the following parameters:
//
//    - snapToTicks: flag indicating if invalid values should be corrected.
//
func (spinButton *SpinButton) SetSnapToTicks(snapToTicks bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	if snapToTicks {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_snap_to_ticks(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(snapToTicks)
}

// SetUpdatePolicy sets the update behavior of a spin button.
//
// This determines whether the spin button is always updated or only when a
// valid value is set.
//
// The function takes the following parameters:
//
//    - policy: GtkSpinButtonUpdatePolicy value.
//
func (spinButton *SpinButton) SetUpdatePolicy(policy SpinButtonUpdatePolicy) {
	var _arg0 *C.GtkSpinButton            // out
	var _arg1 C.GtkSpinButtonUpdatePolicy // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.GtkSpinButtonUpdatePolicy(policy)

	C.gtk_spin_button_set_update_policy(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(policy)
}

// SetValue sets the value of spin_button.
//
// The function takes the following parameters:
//
//    - value: new value.
//
func (spinButton *SpinButton) SetValue(value float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.double(value)

	C.gtk_spin_button_set_value(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(value)
}

// SetWrap sets the flag that determines if a spin button value wraps around to
// the opposite limit when the upper or lower limit of the range is exceeded.
//
// The function takes the following parameters:
//
//    - wrap: flag indicating if wrapping behavior is performed.
//
func (spinButton *SpinButton) SetWrap(wrap bool) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_spin_button_set_wrap(_arg0, _arg1)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(wrap)
}

// Spin: increment or decrement a spin button’s value in a specified direction
// by a specified amount.
//
// The function takes the following parameters:
//
//    - direction: GtkSpinType indicating the direction to spin.
//    - increment: step increment to apply in the specified direction.
//
func (spinButton *SpinButton) Spin(direction SpinType, increment float64) {
	var _arg0 *C.GtkSpinButton // out
	var _arg1 C.GtkSpinType    // out
	var _arg2 C.double         // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))
	_arg1 = C.GtkSpinType(direction)
	_arg2 = C.double(increment)

	C.gtk_spin_button_spin(_arg0, _arg1, _arg2)
	runtime.KeepAlive(spinButton)
	runtime.KeepAlive(direction)
	runtime.KeepAlive(increment)
}

// Update: manually force an update of the spin button.
func (spinButton *SpinButton) Update() {
	var _arg0 *C.GtkSpinButton // out

	_arg0 = (*C.GtkSpinButton)(unsafe.Pointer(spinButton.Native()))

	C.gtk_spin_button_update(_arg0)
	runtime.KeepAlive(spinButton)
}

// ConnectChangeValue: emitted when the user initiates a value change.
//
// This is a keybinding signal (class.SignalAction.html).
//
// Applications should not connect to it, but may emit it with
// g_signal_emit_by_name() if they need to control the cursor programmatically.
//
// The default bindings for this signal are Up/Down and PageUp/PageDown.
func (spinButton *SpinButton) ConnectChangeValue(f func(scroll ScrollType)) externglib.SignalHandle {
	return spinButton.Connect("change-value", f)
}

// ConnectOutput: emitted to tweak the formatting of the value for display.
//
//    // show leading zeros
//    static gboolean
//    on_output (GtkSpinButton *spin,
//               gpointer       data)
//    {
//       GtkAdjustment *adjustment;
//       char *text;
//       int value;
//
//       adjustment = gtk_spin_button_get_adjustment (spin);
//       value = (int)gtk_adjustment_get_value (adjustment);
//       text = g_strdup_printf ("02d", value);
//       gtk_spin_button_set_text (spin, text):
//       g_free (text);
//
//       return TRUE;
//    }.
func (spinButton *SpinButton) ConnectOutput(f func() bool) externglib.SignalHandle {
	return spinButton.Connect("output", f)
}

// ConnectValueChanged: emitted when the value is changed.
//
// Also see the gtk.SpinButton::output signal.
func (spinButton *SpinButton) ConnectValueChanged(f func()) externglib.SignalHandle {
	return spinButton.Connect("value-changed", f)
}

// ConnectWrapped: emitted right after the spinbutton wraps from its maximum to
// its minimum value or vice-versa.
func (spinButton *SpinButton) ConnectWrapped(f func()) externglib.SignalHandle {
	return spinButton.Connect("wrapped", f)
}
