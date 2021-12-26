// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_theming_engine_get_type()), F: marshalThemingEnginer},
	})
}

// ThemingEngineOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ThemingEngineOverrider interface {
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderActivity(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - angle
	//    - x
	//    - y
	//    - size
	//
	RenderArrow(cr *cairo.Context, angle, x, y, size float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderBackground(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderCheck(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderExpander(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - gapSide
	//
	RenderExtension(cr *cairo.Context, x, y, width, height float64, gapSide PositionType)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFocus(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderFrame(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - gapSide
	//    - xy0Gap
	//    - xy1Gap
	//
	RenderFrameGap(cr *cairo.Context, x, y, width, height float64, gapSide PositionType, xy0Gap, xy1Gap float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderHandle(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - pixbuf
	//    - x
	//    - y
	//
	RenderIcon(cr *cairo.Context, pixbuf *gdkpixbuf.Pixbuf, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - surface
	//    - x
	//    - y
	//
	RenderIconSurface(cr *cairo.Context, surface *cairo.Surface, x, y float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - layout
	//
	RenderLayout(cr *cairo.Context, x, y float64, layout *pango.Layout)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x0
	//    - y0
	//    - x1
	//    - y1
	//
	RenderLine(cr *cairo.Context, x0, y0, x1, y1 float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//
	RenderOption(cr *cairo.Context, x, y, width, height float64)
	// The function takes the following parameters:
	//
	//    - cr
	//    - x
	//    - y
	//    - width
	//    - height
	//    - orientation
	//
	RenderSlider(cr *cairo.Context, x, y, width, height float64, orientation Orientation)
}

// ThemingEngine was the object used for rendering themed content in GTK+
// widgets. It used to allow overriding GTK+'s default implementation of
// rendering functions by allowing engines to be loaded as modules.
//
// ThemingEngine has been deprecated in GTK+ 3.14 and will be ignored for
// rendering. The advancements in CSS theming are good enough to allow themers
// to achieve their goals without the need to modify source code.
type ThemingEngine struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*ThemingEngine)(nil)
)

func wrapThemingEngine(obj *externglib.Object) *ThemingEngine {
	return &ThemingEngine{
		Object: obj,
	}
}

func marshalThemingEnginer(p uintptr) (interface{}, error) {
	return wrapThemingEngine(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BackgroundColor gets the background color for a given state.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the color for.
//
// The function returns the following values:
//
//    - color: return value for the background color.
//
func (engine *ThemingEngine) BackgroundColor(state StateFlags) *gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_background_color(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _color
}

// Border gets the border for a given state as a Border.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the border for.
//
// The function returns the following values:
//
//    - border: return value for the border settings.
//
func (engine *ThemingEngine) Border(state StateFlags) *Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _border *Border // out

	_border = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _border
}

// BorderColor gets the border color for a given state.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the color for.
//
// The function returns the following values:
//
//    - color: return value for the border color.
//
func (engine *ThemingEngine) BorderColor(state StateFlags) *gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_border_color(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _color
}

// Color gets the foreground color for a given state.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the color for.
//
// The function returns the following values:
//
//    - color: return value for the foreground color.
//
func (engine *ThemingEngine) Color(state StateFlags) *gdk.RGBA {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GdkRGBA           // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_color(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _color *gdk.RGBA // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _color
}

// Direction returns the widget direction used for rendering.
//
// Deprecated: Use gtk_theming_engine_get_state() and check for
// K_STATE_FLAG_DIR_LTR and K_STATE_FLAG_DIR_RTL instead.
//
// The function returns the following values:
//
//    - textDirection: widget direction.
//
func (engine *ThemingEngine) Direction() TextDirection {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkTextDirection  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_direction(_arg0)
	runtime.KeepAlive(engine)

	var _textDirection TextDirection // out

	_textDirection = TextDirection(_cret)

	return _textDirection
}

// Font returns the font description for a given state.
//
// Deprecated: Use gtk_theming_engine_get().
//
// The function takes the following parameters:
//
//    - state to retrieve the font for.
//
// The function returns the following values:
//
//    - fontDescription for the given state. This object is owned by GTK+ and
//      should not be freed.
//
func (engine *ThemingEngine) Font(state StateFlags) *pango.FontDescription {
	var _arg0 *C.GtkThemingEngine     // out
	var _arg1 C.GtkStateFlags         // out
	var _cret *C.PangoFontDescription // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	_cret = C.gtk_theming_engine_get_font(_arg0, _arg1)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _fontDescription *pango.FontDescription // out

	_fontDescription = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _fontDescription
}

// JunctionSides returns the widget direction used for rendering.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - junctionSides: widget direction.
//
func (engine *ThemingEngine) JunctionSides() JunctionSides {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkJunctionSides  // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_junction_sides(_arg0)
	runtime.KeepAlive(engine)

	var _junctionSides JunctionSides // out

	_junctionSides = JunctionSides(_cret)

	return _junctionSides
}

// Margin gets the margin for a given state as a Border.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the border for.
//
// The function returns the following values:
//
//    - margin: return value for the margin settings.
//
func (engine *ThemingEngine) Margin(state StateFlags) *Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_margin(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _margin *Border // out

	_margin = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _margin
}

// Padding gets the padding for a given state as a Border.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - state to retrieve the padding for.
//
// The function returns the following values:
//
//    - padding: return value for the padding settings.
//
func (engine *ThemingEngine) Padding(state StateFlags) *Border {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateFlags     // out
	var _arg2 C.GtkBorder         // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_padding(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _padding *Border // out

	_padding = (*Border)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _padding
}

// Path returns the widget path used for style matching.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - widgetPath: WidgetPath.
//
func (engine *ThemingEngine) Path() *WidgetPath {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GtkWidgetPath    // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_path(_arg0)
	runtime.KeepAlive(engine)

	var _widgetPath *WidgetPath // out

	_widgetPath = (*WidgetPath)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gtk_widget_path_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_widgetPath)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_widget_path_unref((*C.GtkWidgetPath)(intern.C))
		},
	)

	return _widgetPath
}

// Property gets a property value as retrieved from the style settings that
// apply to the currently rendered element.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - property name.
//    - state to retrieve the value for.
//
// The function returns the following values:
//
//    - value: return location for the property value, you must free this memory
//      using g_value_unset() once you are done with it.
//
func (engine *ThemingEngine) Property(property string, state StateFlags) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkStateFlags     // out
	var _arg3 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(property)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkStateFlags(state)

	C.gtk_theming_engine_get_property(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(property)
	runtime.KeepAlive(state)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))
	runtime.SetFinalizer(_value, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(unsafe.Pointer(v.Native())))
	})

	return _value
}

// Screen returns the Screen to which engine currently rendering to.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - screen (optional) or NULL.
//
func (engine *ThemingEngine) Screen() *gdk.Screen {
	var _arg0 *C.GtkThemingEngine // out
	var _cret *C.GdkScreen        // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_screen(_arg0)
	runtime.KeepAlive(engine)

	var _screen *gdk.Screen // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_screen = &gdk.Screen{
				Object: obj,
			}
		}
	}

	return _screen
}

// State returns the state used when rendering.
//
// Deprecated: since version 3.14.
//
// The function returns the following values:
//
//    - stateFlags: state flags.
//
func (engine *ThemingEngine) State() StateFlags {
	var _arg0 *C.GtkThemingEngine // out
	var _cret C.GtkStateFlags     // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))

	_cret = C.gtk_theming_engine_get_state(_arg0)
	runtime.KeepAlive(engine)

	var _stateFlags StateFlags // out

	_stateFlags = StateFlags(_cret)

	return _stateFlags
}

// StyleProperty gets the value for a widget style property.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - propertyName: name of the widget style property.
//
// The function returns the following values:
//
//    - value: return location for the property value, free with g_value_unset()
//      after use.
//
func (engine *ThemingEngine) StyleProperty(propertyName string) externglib.Value {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GValue            // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_theming_engine_get_style_property(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(propertyName)

	var _value externglib.Value // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg2)))

	return _value
}

// HasClass returns TRUE if the currently rendered contents have defined the
// given class name.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - styleClass class name to look up.
//
// The function returns the following values:
//
//    - ok: TRUE if engine has class_name defined.
//
func (engine *ThemingEngine) HasClass(styleClass string) bool {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleClass)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_has_class(_arg0, _arg1)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(styleClass)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasRegion returns TRUE if the currently rendered contents have the region
// defined. If flags_return is not NULL, it is set to the flags affecting the
// region.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - styleRegion: region name.
//
// The function returns the following values:
//
//    - flags (optional): return location for region flags.
//    - ok: TRUE if region is defined.
//
func (engine *ThemingEngine) HasRegion(styleRegion string) (RegionFlags, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GtkRegionFlags    // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(styleRegion)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_has_region(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(styleRegion)

	var _flags RegionFlags // out
	var _ok bool           // out

	_flags = RegionFlags(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _flags, _ok
}

// LookupColor looks up and resolves a color name in the current style’s color
// map.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - colorName: color name to lookup.
//
// The function returns the following values:
//
//    - color: return location for the looked up color.
//    - ok: TRUE if color_name was found and resolved, FALSE otherwise.
//
func (engine *ThemingEngine) LookupColor(colorName string) (*gdk.RGBA, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 *C.gchar            // out
	var _arg2 C.GdkRGBA           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(colorName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_lookup_color(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(colorName)

	var _color *gdk.RGBA // out
	var _ok bool         // out

	_color = (*gdk.RGBA)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// StateIsRunning returns TRUE if there is a transition animation running for
// the current region (see gtk_style_context_push_animatable_region()).
//
// If progress is not NULL, the animation progress will be returned there, 0.0
// means the state is closest to being FALSE, while 1.0 means it’s closest to
// being TRUE. This means transition animations will run from 0 to 1 when state
// is being set to TRUE and from 1 to 0 when it’s being set to FALSE.
//
// Deprecated: Always returns FALSE.
//
// The function takes the following parameters:
//
//    - state: widget state.
//
// The function returns the following values:
//
//    - progress: return location for the transition progress.
//    - ok: TRUE if there is a running transition animation for state.
//
func (engine *ThemingEngine) StateIsRunning(state StateType) (float64, bool) {
	var _arg0 *C.GtkThemingEngine // out
	var _arg1 C.GtkStateType      // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkThemingEngine)(unsafe.Pointer(engine.Native()))
	_arg1 = C.GtkStateType(state)

	_cret = C.gtk_theming_engine_state_is_running(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(engine)
	runtime.KeepAlive(state)

	var _progress float64 // out
	var _ok bool          // out

	_progress = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _progress, _ok
}

// ThemingEngineLoad loads and initializes a theming engine module from the
// standard directories.
//
// Deprecated: since version 3.14.
//
// The function takes the following parameters:
//
//    - name: theme engine name to load.
//
// The function returns the following values:
//
//    - themingEngine (optional): theming engine, or NULL if the engine name
//      doesn’t exist.
//
func ThemingEngineLoad(name string) *ThemingEngine {
	var _arg1 *C.gchar            // out
	var _cret *C.GtkThemingEngine // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_theming_engine_load(_arg1)
	runtime.KeepAlive(name)

	var _themingEngine *ThemingEngine // out

	if _cret != nil {
		_themingEngine = wrapThemingEngine(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _themingEngine
}
