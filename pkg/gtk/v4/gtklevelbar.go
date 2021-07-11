// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_level_bar_get_type()), F: marshalLevelBarrer},
	})
}

// LevelBarrer describes LevelBar's methods.
type LevelBarrer interface {
	// AddOffsetValue adds a new offset marker on @self at the position
	// specified by @value.
	AddOffsetValue(name string, value float64)
	// Inverted returns whether the levelbar is inverted.
	Inverted() bool
	// MaxValue returns the `max-value` of the `GtkLevelBar`.
	MaxValue() float64
	// MinValue returns the `min-value of the `GtkLevelBar`.
	MinValue() float64
	// Mode returns the `mode` of the `GtkLevelBar`.
	Mode() LevelBarMode
	// OffsetValue fetches the value specified for the offset marker @name in
	// @self.
	OffsetValue(name string) (float64, bool)
	// Value returns the `value` of the `GtkLevelBar`.
	Value() float64
	// RemoveOffsetValue removes an offset marker from a `GtkLevelBar`.
	RemoveOffsetValue(name string)
	// SetInverted sets whether the `GtkLevelBar` is inverted.
	SetInverted(inverted bool)
	// SetMaxValue sets the `max-value` of the `GtkLevelBar`.
	SetMaxValue(value float64)
	// SetMinValue sets the `min-value` of the `GtkLevelBar`.
	SetMinValue(value float64)
	// SetValue sets the value of the `GtkLevelBar`.
	SetValue(value float64)
}

// LevelBar: `GtkLevelBar` is a widget that can be used as a level indicator.
//
// Typical use cases are displaying the strength of a password, or showing the
// charge level of a battery.
//
// !An example GtkLevelBar (levelbar.png)
//
// Use [method@Gtk.LevelBar.set_value] to set the current value, and
// [method@Gtk.LevelBar.add_offset_value] to set the value offsets at which the
// bar will be considered in a different state. GTK will add a few offsets by
// default on the level bar: GTK_LEVEL_BAR_OFFSET_LOW, GTK_LEVEL_BAR_OFFSET_HIGH
// and GTK_LEVEL_BAR_OFFSET_FULL, with values 0.25, 0.75 and 1.0 respectively.
//
// Note that it is your responsibility to update preexisting offsets when
// changing the minimum or maximum value. GTK will simply clamp them to the new
// range.
//
//
// Adding a custom offset on the bar
//
// “`c static GtkWidget * create_level_bar (void) { GtkWidget *widget;
// GtkLevelBar *bar;
//
//    widget = gtk_level_bar_new ();
//    bar = GTK_LEVEL_BAR (widget);
//
//    // This changes the value of the default low offset
//
//    gtk_level_bar_add_offset_value (bar,
//                                    GTK_LEVEL_BAR_OFFSET_LOW,
//                                    0.10);
//
//    // This adds a new offset to the bar; the application will
//    // be able to change its color CSS like this:
//    //
//    // levelbar block.my-offset {
//    //   background-color: magenta;
//    //   border-style: solid;
//    //   border-color: black;
//    //   border-style: 1px;
//    // }
//
//    gtk_level_bar_add_offset_value (bar, "my-offset", 0.60);
//
//    return widget;
//
// } “`
//
// The default interval of values is between zero and one, but it’s possible to
// modify the interval using [method@Gtk.LevelBar.set_min_value] and
// [method@Gtk.LevelBar.set_max_value]. The value will be always drawn in
// proportion to the admissible interval, i.e. a value of 15 with a specified
// interval between 10 and 20 is equivalent to a value of 0.5 with an interval
// between 0 and 1. When K_LEVEL_BAR_MODE_DISCRETE is used, the bar level is
// rendered as a finite number of separated blocks instead of a single one. The
// number of blocks that will be rendered is equal to the number of units
// specified by the admissible interval.
//
// For instance, to build a bar rendered with five blocks, it’s sufficient to
// set the minimum value to 0 and the maximum value to 5 after changing the
// indicator mode to discrete.
//
//
// GtkLevelBar as GtkBuildable
//
// The `GtkLevelBar` implementation of the `GtkBuildable` interface supports a
// custom <offsets> element, which can contain any number of <offset> elements,
// each of which must have name and value attributes.
//
//
// CSS nodes
//
// “` levelbar[.discrete] ╰── trough ├── block.filled.level-name ┊ ├──
// block.empty ┊ “`
//
// `GtkLevelBar` has a main CSS node with name levelbar and one of the style
// classes .discrete or .continuous and a subnode with name trough. Below the
// trough node are a number of nodes with name block and style class .filled or
// .empty. In continuous mode, there is exactly one node of each, in discrete
// mode, the number of filled and unfilled nodes corresponds to blocks that are
// drawn. The block.filled nodes also get a style class .level-name
// corresponding to the level for the current value.
//
// In horizontal orientation, the nodes are always arranged from left to right,
// regardless of text direction.
//
//
// Accessibility
//
// `GtkLevelBar` uses the K_ACCESSIBLE_ROLE_METER role.
type LevelBar struct {
	Widget

	Orientable
}

var (
	_ LevelBarrer     = (*LevelBar)(nil)
	_ gextras.Nativer = (*LevelBar)(nil)
)

func wrapLevelBar(obj *externglib.Object) LevelBarrer {
	return &LevelBar{
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
		},
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalLevelBarrer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLevelBar(obj), nil
}

// NewLevelBar creates a new `GtkLevelBar`.
func NewLevelBar() *LevelBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_level_bar_new()

	var _levelBar *LevelBar // out

	_levelBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LevelBar)

	return _levelBar
}

// NewLevelBarForInterval creates a new `GtkLevelBar` for the specified
// interval.
func NewLevelBarForInterval(minValue float64, maxValue float64) *LevelBar {
	var _arg1 C.double     // out
	var _arg2 C.double     // out
	var _cret *C.GtkWidget // in

	_arg1 = C.double(minValue)
	_arg2 = C.double(maxValue)

	_cret = C.gtk_level_bar_new_for_interval(_arg1, _arg2)

	var _levelBar *LevelBar // out

	_levelBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LevelBar)

	return _levelBar
}

// Native implements gextras.Nativer. It returns the underlying GObject
// field.
func (v *LevelBar) Native() uintptr {
	return v.Widget.InitiallyUnowned.Object.Native()
}

// AddOffsetValue adds a new offset marker on @self at the position specified by
// @value.
//
// When the bar value is in the interval topped by @value (or between @value and
// [property@Gtk.LevelBar:max-value] in case the offset is the last one on the
// bar) a style class named `level-`@name will be applied when rendering the
// level bar fill.
//
// If another offset marker named @name exists, its value will be replaced by
// @value.
func (self *LevelBar) AddOffsetValue(name string, value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.char        // out
	var _arg2 C.double       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)

	C.gtk_level_bar_add_offset_value(_arg0, _arg1, _arg2)
}

// Inverted returns whether the levelbar is inverted.
func (self *LevelBar) Inverted() bool {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxValue returns the `max-value` of the `GtkLevelBar`.
func (self *LevelBar) MaxValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.double       // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_max_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinValue returns the `min-value of the `GtkLevelBar`.
func (self *LevelBar) MinValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.double       // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_min_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Mode returns the `mode` of the `GtkLevelBar`.
func (self *LevelBar) Mode() LevelBarMode {
	var _arg0 *C.GtkLevelBar    // out
	var _cret C.GtkLevelBarMode // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_mode(_arg0)

	var _levelBarMode LevelBarMode // out

	_levelBarMode = (LevelBarMode)(_cret)

	return _levelBarMode
}

// OffsetValue fetches the value specified for the offset marker @name in @self.
func (self *LevelBar) OffsetValue(name string) (float64, bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.char        // out
	var _arg2 C.double       // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_level_bar_get_offset_value(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// Value returns the `value` of the `GtkLevelBar`.
func (self *LevelBar) Value() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.double       // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RemoveOffsetValue removes an offset marker from a `GtkLevelBar`.
//
// The marker must have been previously added with
// [method@Gtk.LevelBar.add_offset_value].
func (self *LevelBar) RemoveOffsetValue(name string) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_level_bar_remove_offset_value(_arg0, _arg1)
}

// SetInverted sets whether the `GtkLevelBar` is inverted.
func (self *LevelBar) SetInverted(inverted bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_level_bar_set_inverted(_arg0, _arg1)
}

// SetMaxValue sets the `max-value` of the `GtkLevelBar`.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (self *LevelBar) SetMaxValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(value)

	C.gtk_level_bar_set_max_value(_arg0, _arg1)
}

// SetMinValue sets the `min-value` of the `GtkLevelBar`.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (self *LevelBar) SetMinValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(value)

	C.gtk_level_bar_set_min_value(_arg0, _arg1)
}

// SetValue sets the value of the `GtkLevelBar`.
func (self *LevelBar) SetValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.double       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.double(value)

	C.gtk_level_bar_set_value(_arg0, _arg1)
}
