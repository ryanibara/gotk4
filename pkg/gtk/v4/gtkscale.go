// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gtk4_ScaleClass_get_layout_offsets(GtkScale*, int*, int*);
// extern char* _gotk4_gtk4_ScaleFormatValueFunc(GtkScale*, double, gpointer);
// void _gotk4_gtk4_Scale_virtual_get_layout_offsets(void* fnptr, GtkScale* arg0, int* arg1, int* arg2) {
//   ((void (*)(GtkScale*, int*, int*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeScale = coreglib.Type(C.gtk_scale_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScale, F: marshalScale},
	})
}

type ScaleFormatValueFunc func(scale *Scale, value float64) (utf8 string)

// ScaleOverrides contains methods that are overridable.
type ScaleOverrides struct {
	// LayoutOffsets obtains the coordinates where the scale will draw the
	// PangoLayout representing the text in the scale.
	//
	// Remember when using the PangoLayout function you need to convert to and
	// from pixels using PANGO_PIXELS() or PANGO_SCALE.
	//
	// If the gtkscale:draw-value property is FALSE, the return values are
	// undefined.
	//
	// The function returns the following values:
	//
	//    - x (optional): location to store X offset of layout, or NULL.
	//    - y (optional): location to store Y offset of layout, or NULL.
	//
	LayoutOffsets func() (x, y int)
}

func defaultScaleOverrides(v *Scale) ScaleOverrides {
	return ScaleOverrides{
		LayoutOffsets: v.layoutOffsets,
	}
}

// Scale: GtkScale is a slider control used to select a numeric value.
//
// !An example GtkScale (scales.png)
//
// To use it, you’ll probably want to investigate the methods on its base class,
// gtkrange, in addition to the methods for GtkScale itself. To set the value of
// a scale, you would normally use gtk.Range.SetValue(). To detect changes to
// the value, you would normally use the gtk.Range::value-changed signal.
//
// Note that using the same upper and lower bounds for the GtkScale (through the
// GtkRange methods) will hide the slider itself. This is useful for
// applications that want to show an undeterminate value on the scale, without
// changing the layout of the application (such as movie or music players).
//
//
// GtkScale as GtkBuildable
//
// GtkScale supports a custom <marks> element, which can contain multiple
// <mark\> elements. The “value” and “position” attributes have the same meaning
// as gtk.Scale.AddMark() parameters of the same name. If the element is not
// empty, its content is taken as the markup to show at the mark. It can be
// translated with the usual ”translatable” and “context” attributes.
//
// CSS nodes
//
//    scale[.fine-tune][.marks-before][.marks-after]
//    ├── [value][.top][.right][.bottom][.left]
//    ├── marks.top
//    │   ├── mark
//    │   ┊    ├── [label]
//    │   ┊    ╰── indicator
//    ┊   ┊
//    │   ╰── mark
//    ├── marks.bottom
//    │   ├── mark
//    │   ┊    ├── indicator
//    │   ┊    ╰── [label]
//    ┊   ┊
//    │   ╰── mark
//    ╰── trough
//        ├── [fill]
//        ├── [highlight]
//        ╰── slider
//
//
// GtkScale has a main CSS node with name scale and a subnode for its contents,
// with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scale is in
// 'fine-tuning' mode.
//
// If the scale has an origin (see gtk.Scale.SetHasOrigin()), there is a subnode
// with name highlight below the trough node that is used for rendering the
// highlighted part of the trough.
//
// If the scale is showing a fill level (see gtk.Range.SetShowFillLevel()),
// there is a subnode with name fill below the trough node that is used for
// rendering the filled in part of the trough.
//
// If marks are present, there is a marks subnode before or after the trough
// node, below which each mark gets a node with name mark. The marks nodes get
// either the .top or .bottom style class.
//
// The mark node has a subnode named indicator. If the mark has text, it also
// has a subnode named label. When the mark is either above or left of the
// scale, the label subnode is the first when present. Otherwise, the indicator
// subnode is the first.
//
// The main CSS node gets the 'marks-before' and/or 'marks-after' style classes
// added depending on what marks are present.
//
// If the scale is displaying the value (see gtk.Scale:draw-value), there is
// subnode with name value. This node will get the .top or .bottom style classes
// similar to the marks node.
//
//
// Accessibility
//
// GtkScale uses the GTK_ACCESSIBLE_ROLE_SLIDER role.
type Scale struct {
	_ [0]func() // equal guard
	Range
}

var (
	_ Widgetter         = (*Scale)(nil)
	_ coreglib.Objector = (*Scale)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*Scale, *ScaleClass, ScaleOverrides](
		GTypeScale,
		initScaleClass,
		wrapScale,
		defaultScaleOverrides,
	)
}

func initScaleClass(gclass unsafe.Pointer, overrides ScaleOverrides, classInitFunc func(*ScaleClass)) {
	pclass := (*C.GtkScaleClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeScale))))

	if overrides.LayoutOffsets != nil {
		pclass.get_layout_offsets = (*[0]byte)(C._gotk4_gtk4_ScaleClass_get_layout_offsets)
	}

	if classInitFunc != nil {
		class := (*ScaleClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapScale(obj *coreglib.Object) *Scale {
	return &Scale{
		Range: Range{
			Widget: Widget{
				InitiallyUnowned: coreglib.InitiallyUnowned{
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
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalScale(p uintptr) (interface{}, error) {
	return wrapScale(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewScale creates a new GtkScale.
//
// The function takes the following parameters:
//
//    - orientation scale’s orientation.
//    - adjustment (optional): gtk.Adjustment which sets the range of the scale,
//      or NULL to create a new adjustment.
//
// The function returns the following values:
//
//    - scale: new GtkScale.
//
func NewScale(orientation Orientation, adjustment *Adjustment) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	if adjustment != nil {
		_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(coreglib.InternObject(adjustment).Native()))
	}

	_cret = C.gtk_scale_new(_arg1, _arg2)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(adjustment)

	var _scale *Scale // out

	_scale = wrapScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _scale
}

// NewScaleWithRange creates a new scale widget with a range from min to max.
//
// The returns scale will have the given orientation and will let the user input
// a number between min and max (including min and max) with the increment step.
// step must be nonzero; it’s the distance the slider moves when using the arrow
// keys to adjust the scale value.
//
// Note that the way in which the precision is derived works best if step is a
// power of ten. If the resulting precision is not suitable for your needs, use
// gtk.Scale.SetDigits() to correct it.
//
// The function takes the following parameters:
//
//    - orientation scale’s orientation.
//    - min: minimum value.
//    - max: maximum value.
//    - step increment (tick size) used with keyboard shortcuts.
//
// The function returns the following values:
//
//    - scale: new Scale.
//
func NewScaleWithRange(orientation Orientation, min, max, step float64) *Scale {
	var _arg1 C.GtkOrientation // out
	var _arg2 C.double         // out
	var _arg3 C.double         // out
	var _arg4 C.double         // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = C.double(min)
	_arg3 = C.double(max)
	_arg4 = C.double(step)

	_cret = C.gtk_scale_new_with_range(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(orientation)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
	runtime.KeepAlive(step)

	var _scale *Scale // out

	_scale = wrapScale(coreglib.Take(unsafe.Pointer(_cret)))

	return _scale
}

// AddMark adds a mark at value.
//
// A mark is indicated visually by drawing a tick mark next to the scale, and
// GTK makes it easy for the user to position the scale exactly at the marks
// value.
//
// If markup is not NULL, text is shown next to the tick mark.
//
// To remove marks from a scale, use gtk.Scale.ClearMarks().
//
// The function takes the following parameters:
//
//    - value at which the mark is placed, must be between the lower and upper
//      limits of the scales’ adjustment.
//    - position: where to draw the mark. For a horizontal scale, K_POS_TOP and
//      GTK_POS_LEFT are drawn above the scale, anything else below. For a
//      vertical scale, K_POS_LEFT and GTK_POS_TOP are drawn to the left of the
//      scale, anything else to the right.
//    - markup (optional): text to be shown at the mark, using Pango markup, or
//      NULL.
//
func (scale *Scale) AddMark(value float64, position PositionType, markup string) {
	var _arg0 *C.GtkScale       // out
	var _arg1 C.double          // out
	var _arg2 C.GtkPositionType // out
	var _arg3 *C.char           // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	_arg1 = C.double(value)
	_arg2 = C.GtkPositionType(position)
	if markup != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(markup)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.gtk_scale_add_mark(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(value)
	runtime.KeepAlive(position)
	runtime.KeepAlive(markup)
}

// ClearMarks removes any marks that have been added.
func (scale *Scale) ClearMarks() {
	var _arg0 *C.GtkScale // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	C.gtk_scale_clear_marks(_arg0)
	runtime.KeepAlive(scale)
}

// Digits gets the number of decimal places that are displayed in the value.
//
// The function returns the following values:
//
//    - gint: number of decimal places that are displayed.
//
func (scale *Scale) Digits() int {
	var _arg0 *C.GtkScale // out
	var _cret C.int       // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	_cret = C.gtk_scale_get_digits(_arg0)
	runtime.KeepAlive(scale)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DrawValue returns whether the current value is displayed as a string next to
// the slider.
//
// The function returns the following values:
//
//    - ok: whether the current value is displayed as a string.
//
func (scale *Scale) DrawValue() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	_cret = C.gtk_scale_get_draw_value(_arg0)
	runtime.KeepAlive(scale)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasOrigin returns whether the scale has an origin.
//
// The function returns the following values:
//
//    - ok: TRUE if the scale has an origin.
//
func (scale *Scale) HasOrigin() bool {
	var _arg0 *C.GtkScale // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	_cret = C.gtk_scale_get_has_origin(_arg0)
	runtime.KeepAlive(scale)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Layout gets the PangoLayout used to display the scale.
//
// The returned object is owned by the scale so does not need to be freed by the
// caller.
//
// The function returns the following values:
//
//    - layout (optional): pango.Layout for this scale, or NULL if the
//      gtkscale:draw-value property is FALSE.
//
func (scale *Scale) Layout() *pango.Layout {
	var _arg0 *C.GtkScale    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	_cret = C.gtk_scale_get_layout(_arg0)
	runtime.KeepAlive(scale)

	var _layout *pango.Layout // out

	if _cret != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_layout = &pango.Layout{
				Object: obj,
			}
		}
	}

	return _layout
}

// LayoutOffsets obtains the coordinates where the scale will draw the
// PangoLayout representing the text in the scale.
//
// Remember when using the PangoLayout function you need to convert to and from
// pixels using PANGO_PIXELS() or PANGO_SCALE.
//
// If the gtkscale:draw-value property is FALSE, the return values are
// undefined.
//
// The function returns the following values:
//
//    - x (optional): location to store X offset of layout, or NULL.
//    - y (optional): location to store Y offset of layout, or NULL.
//
func (scale *Scale) LayoutOffsets() (x, y int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	C.gtk_scale_get_layout_offsets(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(scale)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// ValuePos gets the position in which the current value is displayed.
//
// The function returns the following values:
//
//    - positionType: position in which the current value is displayed.
//
func (scale *Scale) ValuePos() PositionType {
	var _arg0 *C.GtkScale       // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	_cret = C.gtk_scale_get_value_pos(_arg0)
	runtime.KeepAlive(scale)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// SetDigits sets the number of decimal places that are displayed in the value.
//
// Also causes the value of the adjustment to be rounded to this number of
// digits, so the retrieved value matches the displayed one, if
// gtkscale:draw-value is TRUE when the value changes. If you want to enforce
// rounding the value when gtkscale:draw-value is FALSE, you can set
// gtkrange:round-digits instead.
//
// Note that rounding to a small number of digits can interfere with the smooth
// autoscrolling that is built into GtkScale. As an alternative, you can use
// gtk.Scale.SetFormatValueFunc() to format the displayed value yourself.
//
// The function takes the following parameters:
//
//    - digits: number of decimal places to display, e.g. use 1 to display 1.0, 2
//      to display 1.00, etc.
//
func (scale *Scale) SetDigits(digits int) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	_arg1 = C.int(digits)

	C.gtk_scale_set_digits(_arg0, _arg1)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(digits)
}

// SetDrawValue specifies whether the current value is displayed as a string
// next to the slider.
//
// The function takes the following parameters:
//
//    - drawValue: TRUE to draw the value.
//
func (scale *Scale) SetDrawValue(drawValue bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	if drawValue {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_draw_value(_arg0, _arg1)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(drawValue)
}

// SetFormatValueFunc: func allows you to change how the scale value is
// displayed.
//
// The given function will return an allocated string representing value. That
// string will then be used to display the scale's value.
//
// If LL is passed as func, the value will be displayed on its own, rounded
// according to the value of the gtkscale:digits property.
//
// The function takes the following parameters:
//
//    - fn (optional): function that formats the value.
//
func (scale *Scale) SetFormatValueFunc(fn ScaleFormatValueFunc) {
	var _arg0 *C.GtkScale               // out
	var _arg1 C.GtkScaleFormatValueFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	if fn != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk4_ScaleFormatValueFunc)
		_arg2 = C.gpointer(gbox.Assign(fn))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_scale_set_format_value_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(fn)
}

// SetHasOrigin sets whether the scale has an origin.
//
// If gtkscale:has-origin is set to TRUE (the default), the scale will highlight
// the part of the trough between the origin (bottom or left side) and the
// current value.
//
// The function takes the following parameters:
//
//    - hasOrigin: TRUE if the scale has an origin.
//
func (scale *Scale) SetHasOrigin(hasOrigin bool) {
	var _arg0 *C.GtkScale // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	if hasOrigin {
		_arg1 = C.TRUE
	}

	C.gtk_scale_set_has_origin(_arg0, _arg1)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(hasOrigin)
}

// SetValuePos sets the position in which the current value is displayed.
//
// The function takes the following parameters:
//
//    - pos: position in which the current value is displayed.
//
func (scale *Scale) SetValuePos(pos PositionType) {
	var _arg0 *C.GtkScale       // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))
	_arg1 = C.GtkPositionType(pos)

	C.gtk_scale_set_value_pos(_arg0, _arg1)
	runtime.KeepAlive(scale)
	runtime.KeepAlive(pos)
}

// layoutOffsets obtains the coordinates where the scale will draw the
// PangoLayout representing the text in the scale.
//
// Remember when using the PangoLayout function you need to convert to and from
// pixels using PANGO_PIXELS() or PANGO_SCALE.
//
// If the gtkscale:draw-value property is FALSE, the return values are
// undefined.
//
// The function returns the following values:
//
//    - x (optional): location to store X offset of layout, or NULL.
//    - y (optional): location to store Y offset of layout, or NULL.
//
func (scale *Scale) layoutOffsets() (x, y int) {
	gclass := (*C.GtkScaleClass)(coreglib.PeekParentClass(scale))
	fnarg := gclass.get_layout_offsets

	var _arg0 *C.GtkScale // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkScale)(unsafe.Pointer(coreglib.InternObject(scale).Native()))

	C._gotk4_gtk4_Scale_virtual_get_layout_offsets(unsafe.Pointer(fnarg), _arg0, &_arg1, &_arg2)
	runtime.KeepAlive(scale)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// ScaleClass: instance of this type is always passed by reference.
type ScaleClass struct {
	*scaleClass
}

// scaleClass is the struct that's finalized.
type scaleClass struct {
	native *C.GtkScaleClass
}

func (s *ScaleClass) ParentClass() *RangeClass {
	valptr := &s.native.parent_class
	var _v *RangeClass // out
	_v = (*RangeClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
