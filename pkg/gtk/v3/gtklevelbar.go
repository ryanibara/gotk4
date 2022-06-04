// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_LevelBarClass_offset_changed(void*, void*);
// extern void _gotk4_gtk3_LevelBar_ConnectOffsetChanged(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtklevelbar.go.
var GTypeLevelBar = coreglib.Type(C.gtk_level_bar_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeLevelBar, F: marshalLevelBar},
	})
}

// LEVEL_BAR_OFFSET_FULL: name used for the stock full offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_FULL = "full"

// LEVEL_BAR_OFFSET_HIGH: name used for the stock high offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_HIGH = "high"

// LEVEL_BAR_OFFSET_LOW: name used for the stock low offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_LOW = "low"

// LevelBarOverrider contains methods that are overridable.
type LevelBarOverrider interface {
	// The function takes the following parameters:
	//
	OffsetChanged(name string)
}

// LevelBar is a bar widget that can be used as a level indicator. Typical use
// cases are displaying the strength of a password, or showing the charge level
// of a battery.
//
// Use gtk_level_bar_set_value() to set the current value, and
// gtk_level_bar_add_offset_value() to set the value offsets at which the bar
// will be considered in a different state. GTK will add a few offsets by
// default on the level bar: K_LEVEL_BAR_OFFSET_LOW, K_LEVEL_BAR_OFFSET_HIGH and
// K_LEVEL_BAR_OFFSET_FULL, with values 0.25, 0.75 and 1.0 respectively.
//
// Note that it is your responsibility to update preexisting offsets when
// changing the minimum or maximum value. GTK+ will simply clamp them to the new
// range.
//
// Adding a custom offset on the bar
//
//    levelbar[.discrete]
//    ╰── trough
//        ├── block.filled.level-name
//        ┊
//        ├── block.empty
//        ┊
//
// GtkLevelBar has a main CSS node with name levelbar and one of the style
// classes .discrete or .continuous and a subnode with name trough. Below the
// trough node are a number of nodes with name block and style class .filled or
// .empty. In continuous mode, there is exactly one node of each, in discrete
// mode, the number of filled and unfilled nodes corresponds to blocks that are
// drawn. The block.filled nodes also get a style class .level-name
// corresponding to the level for the current value.
//
// In horizontal orientation, the nodes are always arranged from left to right,
// regardless of text direction.
type LevelBar struct {
	_ [0]func() // equal guard
	Widget

	*coreglib.Object
	Orientable
}

var (
	_ Widgetter         = (*LevelBar)(nil)
	_ coreglib.Objector = (*LevelBar)(nil)
)

func classInitLevelBarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkLevelBarClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkLevelBarClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ OffsetChanged(name string) }); ok {
		pclass.offset_changed = (*[0]byte)(C._gotk4_gtk3_LevelBarClass_offset_changed)
	}
}

//export _gotk4_gtk3_LevelBarClass_offset_changed
func _gotk4_gtk3_LevelBarClass_offset_changed(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ OffsetChanged(name string) })

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.OffsetChanged(_name)
}

func wrapLevelBar(obj *coreglib.Object) *LevelBar {
	return &LevelBar{
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
		Object: obj,
		Orientable: Orientable{
			Object: obj,
		},
	}
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	return wrapLevelBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_LevelBar_ConnectOffsetChanged
func _gotk4_gtk3_LevelBar_ConnectOffsetChanged(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(name string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(name string))
	}

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_name)
}

// ConnectOffsetChanged is emitted when an offset specified on the bar changes
// value as an effect to gtk_level_bar_add_offset_value() being called.
//
// The signal supports detailed connections; you can connect to the detailed
// signal "changed::x" in order to only receive callbacks when the value of
// offset "x" changes.
func (self *LevelBar) ConnectOffsetChanged(f func(name string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(self, "offset-changed", false, unsafe.Pointer(C._gotk4_gtk3_LevelBar_ConnectOffsetChanged), f)
}

// NewLevelBar creates a new LevelBar.
//
// The function returns the following values:
//
//    - levelBar: LevelBar.
//
func NewLevelBar() *LevelBar {
	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("new_LevelBar", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// NewLevelBarForInterval: utility constructor that creates a new LevelBar for
// the specified interval.
//
// The function takes the following parameters:
//
//    - minValue: positive value.
//    - maxValue: positive value.
//
// The function returns the following values:
//
//    - levelBar: LevelBar.
//
func NewLevelBarForInterval(minValue, maxValue float64) *LevelBar {
	var _args [2]girepository.Argument

	*(*C.gdouble)(unsafe.Pointer(&_args[0])) = C.gdouble(minValue)
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(maxValue)

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("new_LevelBar_for_interval", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(minValue)
	runtime.KeepAlive(maxValue)

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// AddOffsetValue adds a new offset marker on self at the position specified by
// value. When the bar value is in the interval topped by value (or between
// value and LevelBar:max-value in case the offset is the last one on the bar) a
// style class named level-name will be applied when rendering the level bar
// fill. If another offset marker named name exists, its value will be replaced
// by value.
//
// The function takes the following parameters:
//
//    - name of the new offset.
//    - value for the new offset.
//
func (self *LevelBar) AddOffsetValue(name string, value float64) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.gdouble)(unsafe.Pointer(&_args[2])) = C.gdouble(value)

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("add_offset_value", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// Inverted: return the value of the LevelBar:inverted property.
//
// The function returns the following values:
//
//    - ok: TRUE if the level bar is inverted.
//
func (self *LevelBar) Inverted() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_inverted", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// MaxValue returns the value of the LevelBar:max-value property.
//
// The function returns the following values:
//
//    - gdouble: positive value.
//
func (self *LevelBar) MaxValue() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_max_value", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// MinValue returns the value of the LevelBar:min-value property.
//
// The function returns the following values:
//
//    - gdouble: positive value.
//
func (self *LevelBar) MinValue() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_min_value", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// OffsetValue fetches the value specified for the offset marker name in self,
// returning TRUE in case an offset named name was found.
//
// The function takes the following parameters:
//
//    - name (optional) of an offset in the bar.
//
// The function returns the following values:
//
//    - value: location where to store the value.
//    - ok: TRUE if the specified offset is found.
//
func (self *LevelBar) OffsetValue(name string) (float64, bool) {
	var _args [2]girepository.Argument
	var _outs [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if name != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_offset_value", _args[:], _outs[:])
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)

	var _value float64 // out
	var _ok bool       // out

	_value = *(*float64)(unsafe.Pointer(_outs[0]))
	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _value, _ok
}

// Value returns the value of the LevelBar:value property.
//
// The function returns the following values:
//
//    - gdouble: value in the interval between LevelBar:min-value and
//      LevelBar:max-value.
//
func (self *LevelBar) Value() float64 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "LevelBar").InvokeMethod("get_value", _args[:], nil)
	_cret = *(*C.gdouble)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(*(*C.gdouble)(unsafe.Pointer(&_cret)))

	return _gdouble
}

// RemoveOffsetValue removes an offset marker previously added with
// gtk_level_bar_add_offset_value().
//
// The function takes the following parameters:
//
//    - name (optional) of an offset in the bar.
//
func (self *LevelBar) RemoveOffsetValue(name string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if name != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("remove_offset_value", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetInverted sets the value of the LevelBar:inverted property.
//
// The function takes the following parameters:
//
//    - inverted: TRUE to invert the level bar.
//
func (self *LevelBar) SetInverted(inverted bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if inverted {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_inverted", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(inverted)
}

// SetMaxValue sets the value of the LevelBar:max-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMaxValue(value float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(value)

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_max_value", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMinValue sets the value of the LevelBar:min-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMinValue(value float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(value)

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_min_value", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetValue sets the value of the LevelBar:value property.
//
// The function takes the following parameters:
//
//    - value in the interval between LevelBar:min-value and LevelBar:max-value.
//
func (self *LevelBar) SetValue(value float64) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.gdouble)(unsafe.Pointer(&_args[1])) = C.gdouble(value)

	girepository.MustFind("Gtk", "LevelBar").InvokeMethod("set_value", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}
