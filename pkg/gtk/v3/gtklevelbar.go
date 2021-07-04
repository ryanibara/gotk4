// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_level_bar_get_type()), F: marshalLevelBar},
	})
}

// LevelBar: the LevelBar is a bar widget that can be used as a level indicator.
// Typical use cases are displaying the strength of a password, or showing the
// charge level of a battery.
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
type LevelBar interface {
	Widget
	Orientable

	AddOffsetValueLevelBar(name string, value float64)

	Inverted() bool

	MaxValue() float64

	MinValue() float64

	Mode() LevelBarMode

	OffsetValue(name string) (float64, bool)

	Value() float64

	RemoveOffsetValueLevelBar(name string)

	SetInvertedLevelBar(inverted bool)

	SetMaxValueLevelBar(value float64)

	SetMinValueLevelBar(value float64)

	SetModeLevelBar(mode LevelBarMode)

	SetValueLevelBar(value float64)
}

// levelBar implements the LevelBar class.
type levelBar struct {
	Widget
}

// WrapLevelBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapLevelBar(obj *externglib.Object) LevelBar {
	return levelBar{
		Widget: WrapWidget(obj),
	}
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLevelBar(obj), nil
}

func NewLevelBar() LevelBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_level_bar_new()

	var _levelBar LevelBar // out

	_levelBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(LevelBar)

	return _levelBar
}

func NewLevelBarForInterval(minValue float64, maxValue float64) LevelBar {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(minValue)
	_arg2 = C.gdouble(maxValue)

	_cret = C.gtk_level_bar_new_for_interval(_arg1, _arg2)

	var _levelBar LevelBar // out

	_levelBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(LevelBar)

	return _levelBar
}

func (s levelBar) AddOffsetValueLevelBar(name string, value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	C.gtk_level_bar_add_offset_value(_arg0, _arg1, _arg2)
}

func (s levelBar) Inverted() bool {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s levelBar) MaxValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_max_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s levelBar) MinValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_min_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s levelBar) Mode() LevelBarMode {
	var _arg0 *C.GtkLevelBar    // out
	var _cret C.GtkLevelBarMode // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_mode(_arg0)

	var _levelBarMode LevelBarMode // out

	_levelBarMode = LevelBarMode(_cret)

	return _levelBarMode
}

func (s levelBar) OffsetValue(name string) (float64, bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
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

func (s levelBar) Value() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (s levelBar) RemoveOffsetValueLevelBar(name string) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_level_bar_remove_offset_value(_arg0, _arg1)
}

func (s levelBar) SetInvertedLevelBar(inverted bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_level_bar_set_inverted(_arg0, _arg1)
}

func (s levelBar) SetMaxValueLevelBar(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_max_value(_arg0, _arg1)
}

func (s levelBar) SetMinValueLevelBar(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_min_value(_arg0, _arg1)
}

func (s levelBar) SetModeLevelBar(mode LevelBarMode) {
	var _arg0 *C.GtkLevelBar    // out
	var _arg1 C.GtkLevelBarMode // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkLevelBarMode(mode)

	C.gtk_level_bar_set_mode(_arg0, _arg1)
}

func (s levelBar) SetValueLevelBar(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_value(_arg0, _arg1)
}

func (b levelBar) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b levelBar) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b levelBar) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b levelBar) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b levelBar) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b levelBar) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b levelBar) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b levelBar) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b levelBar) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b levelBar) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o levelBar) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o levelBar) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}
