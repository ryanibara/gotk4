// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_path_priority_type_get_type()), F: marshalPathPriorityType},
		{T: externglib.Type(C.gtk_path_type_get_type()), F: marshalPathType},
		{T: externglib.Type(C.gtk_rc_token_type_get_type()), F: marshalRCTokenType},
		{T: externglib.Type(C.gtk_rc_flags_get_type()), F: marshalRCFlags},
		{T: externglib.Type(C.gtk_rc_style_get_type()), F: marshalRCStyler},
	})
}

const PATH_PRIO_MASK = 15

// PathPriorityType priorities for path lookups. See also
// gtk_binding_set_add_path().
//
// Deprecated: since version 3.0.
type PathPriorityType int

const (
	// PathPrioLowest: deprecated.
	PathPrioLowest PathPriorityType = 0
	// PathPrioGTK: deprecated.
	PathPrioGTK PathPriorityType = 4
	// PathPrioApplication: deprecated.
	PathPrioApplication PathPriorityType = 8
	// PathPrioTheme: deprecated.
	PathPrioTheme PathPriorityType = 10
	// PathPrioRC: deprecated.
	PathPrioRC PathPriorityType = 12
	// PathPrioHighest: deprecated.
	PathPrioHighest PathPriorityType = 15
)

func marshalPathPriorityType(p uintptr) (interface{}, error) {
	return PathPriorityType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PathPriorityType.
func (p PathPriorityType) String() string {
	switch p {
	case PathPrioLowest:
		return "Lowest"
	case PathPrioGTK:
		return "GTK"
	case PathPrioApplication:
		return "Application"
	case PathPrioTheme:
		return "Theme"
	case PathPrioRC:
		return "RC"
	case PathPrioHighest:
		return "Highest"
	default:
		return fmt.Sprintf("PathPriorityType(%d)", p)
	}
}

// PathType: widget path types. See also gtk_binding_set_add_path().
//
// Deprecated: since version 3.0.
type PathType int

const (
	// PathWidget: deprecated.
	PathWidget PathType = iota
	// PathWidgetClass: deprecated.
	PathWidgetClass
	// PathClass: deprecated.
	PathClass
)

func marshalPathType(p uintptr) (interface{}, error) {
	return PathType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PathType.
func (p PathType) String() string {
	switch p {
	case PathWidget:
		return "Widget"
	case PathWidgetClass:
		return "WidgetClass"
	case PathClass:
		return "Class"
	default:
		return fmt.Sprintf("PathType(%d)", p)
	}
}

// RCTokenType enumeration represents the tokens in the RC file. It is exposed
// so that theme engines can reuse these tokens when parsing the theme-engine
// specific portions of a RC file.
//
// Deprecated: Use CssProvider instead.
type RCTokenType int

const (
	// RCTokenInvalid: deprecated.
	RCTokenInvalid RCTokenType = 270
	// RCTokenInclude: deprecated.
	RCTokenInclude RCTokenType = 271
	// RCTokenNormal: deprecated.
	RCTokenNormal RCTokenType = 272
	// RCTokenActive: deprecated.
	RCTokenActive RCTokenType = 273
	// RCTokenPrelight: deprecated.
	RCTokenPrelight RCTokenType = 274
	// RCTokenSelected: deprecated.
	RCTokenSelected RCTokenType = 275
	// RCTokenInsensitive: deprecated.
	RCTokenInsensitive RCTokenType = 276
	// RCTokenFg: deprecated.
	RCTokenFg RCTokenType = 277
	// RCTokenBg: deprecated.
	RCTokenBg RCTokenType = 278
	// RCTokenText: deprecated.
	RCTokenText RCTokenType = 279
	// RCTokenBase: deprecated.
	RCTokenBase RCTokenType = 280
	// RCTokenXthickness: deprecated.
	RCTokenXthickness RCTokenType = 281
	// RCTokenYthickness: deprecated.
	RCTokenYthickness RCTokenType = 282
	// RCTokenFont: deprecated.
	RCTokenFont RCTokenType = 283
	// RCTokenFontset: deprecated.
	RCTokenFontset RCTokenType = 284
	// RCTokenFontName: deprecated.
	RCTokenFontName RCTokenType = 285
	// RCTokenBgPixmap: deprecated.
	RCTokenBgPixmap RCTokenType = 286
	// RCTokenPixmapPath: deprecated.
	RCTokenPixmapPath RCTokenType = 287
	// RCTokenStyle: deprecated.
	RCTokenStyle RCTokenType = 288
	// RCTokenBinding: deprecated.
	RCTokenBinding RCTokenType = 289
	// RCTokenBind: deprecated.
	RCTokenBind RCTokenType = 290
	// RCTokenWidget: deprecated.
	RCTokenWidget RCTokenType = 291
	// RCTokenWidgetClass: deprecated.
	RCTokenWidgetClass RCTokenType = 292
	// RCTokenClass: deprecated.
	RCTokenClass RCTokenType = 293
	// RCTokenLowest: deprecated.
	RCTokenLowest RCTokenType = 294
	// RCTokenGTK: deprecated.
	RCTokenGTK RCTokenType = 295
	// RCTokenApplication: deprecated.
	RCTokenApplication RCTokenType = 296
	// RCTokenTheme: deprecated.
	RCTokenTheme RCTokenType = 297
	// RCTokenRC: deprecated.
	RCTokenRC RCTokenType = 298
	// RCTokenHighest: deprecated.
	RCTokenHighest RCTokenType = 299
	// RCTokenEngine: deprecated.
	RCTokenEngine RCTokenType = 300
	// RCTokenModulePath: deprecated.
	RCTokenModulePath RCTokenType = 301
	// RCTokenImModulePath: deprecated.
	RCTokenImModulePath RCTokenType = 302
	// RCTokenImModuleFile: deprecated.
	RCTokenImModuleFile RCTokenType = 303
	// RCTokenStock: deprecated.
	RCTokenStock RCTokenType = 304
	// RCTokenLTR: deprecated.
	RCTokenLTR RCTokenType = 305
	// RCTokenRTL: deprecated.
	RCTokenRTL RCTokenType = 306
	// RCTokenColor: deprecated.
	RCTokenColor RCTokenType = 307
	// RCTokenUnbind: deprecated.
	RCTokenUnbind RCTokenType = 308
	// RCTokenLast: deprecated.
	RCTokenLast RCTokenType = 309
)

func marshalRCTokenType(p uintptr) (interface{}, error) {
	return RCTokenType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for RCTokenType.
func (r RCTokenType) String() string {
	switch r {
	case RCTokenInvalid:
		return "Invalid"
	case RCTokenInclude:
		return "Include"
	case RCTokenNormal:
		return "Normal"
	case RCTokenActive:
		return "Active"
	case RCTokenPrelight:
		return "Prelight"
	case RCTokenSelected:
		return "Selected"
	case RCTokenInsensitive:
		return "Insensitive"
	case RCTokenFg:
		return "Fg"
	case RCTokenBg:
		return "Bg"
	case RCTokenText:
		return "Text"
	case RCTokenBase:
		return "Base"
	case RCTokenXthickness:
		return "Xthickness"
	case RCTokenYthickness:
		return "Ythickness"
	case RCTokenFont:
		return "Font"
	case RCTokenFontset:
		return "Fontset"
	case RCTokenFontName:
		return "FontName"
	case RCTokenBgPixmap:
		return "BgPixmap"
	case RCTokenPixmapPath:
		return "PixmapPath"
	case RCTokenStyle:
		return "Style"
	case RCTokenBinding:
		return "Binding"
	case RCTokenBind:
		return "Bind"
	case RCTokenWidget:
		return "Widget"
	case RCTokenWidgetClass:
		return "WidgetClass"
	case RCTokenClass:
		return "Class"
	case RCTokenLowest:
		return "Lowest"
	case RCTokenGTK:
		return "GTK"
	case RCTokenApplication:
		return "Application"
	case RCTokenTheme:
		return "Theme"
	case RCTokenRC:
		return "RC"
	case RCTokenHighest:
		return "Highest"
	case RCTokenEngine:
		return "Engine"
	case RCTokenModulePath:
		return "ModulePath"
	case RCTokenImModulePath:
		return "ImModulePath"
	case RCTokenImModuleFile:
		return "ImModuleFile"
	case RCTokenStock:
		return "Stock"
	case RCTokenLTR:
		return "LTR"
	case RCTokenRTL:
		return "RTL"
	case RCTokenColor:
		return "Color"
	case RCTokenUnbind:
		return "Unbind"
	case RCTokenLast:
		return "Last"
	default:
		return fmt.Sprintf("RCTokenType(%d)", r)
	}
}

// RCFlags: deprecated.
type RCFlags int

const (
	// RCFg: deprecated.
	RCFg RCFlags = 0b1
	// RCBg: deprecated.
	RCBg RCFlags = 0b10
	// RCText: deprecated.
	RCText RCFlags = 0b100
	// RCBase: deprecated.
	RCBase RCFlags = 0b1000
)

func marshalRCFlags(p uintptr) (interface{}, error) {
	return RCFlags(C.g_value_get_flags((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for RCFlags.
func (r RCFlags) String() string {
	if r == 0 {
		return "RCFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(23)

	for r != 0 {
		next := r & (r - 1)
		bit := r - next

		switch bit {
		case RCFg:
			builder.WriteString("Fg|")
		case RCBg:
			builder.WriteString("Bg|")
		case RCText:
			builder.WriteString("Text|")
		case RCBase:
			builder.WriteString("Base|")
		default:
			builder.WriteString(fmt.Sprintf("RCFlags(0b%b)|", bit))
		}

		r = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if r contains other.
func (r RCFlags) Has(other RCFlags) bool {
	return (r & other) == other
}

// RCAddDefaultFile adds a file to the list of files to be parsed at the end of
// gtk_init().
//
// Deprecated: Use StyleContext with a custom StyleProvider instead.
func RCAddDefaultFile(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_add_default_file(_arg1)
	runtime.KeepAlive(filename)
}

// RCFindModuleInPath searches for a theme engine in the GTK+ search path. This
// function is not useful for applications and should not be used.
//
// Deprecated: Use CssProvider instead.
func RCFindModuleInPath(moduleFile string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(moduleFile)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_rc_find_module_in_path(_arg1)
	runtime.KeepAlive(moduleFile)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCFindPixmapInPath looks up a file in pixmap path for the specified Settings.
// If the file is not found, it outputs a warning message using g_warning() and
// returns NULL.
//
// Deprecated: Use CssProvider instead.
func RCFindPixmapInPath(settings *Settings, scanner *glib.Scanner, pixmapFile string) string {
	var _arg1 *C.GtkSettings // out
	var _arg2 *C.GScanner    // out
	var _arg3 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(pixmapFile)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_rc_find_pixmap_in_path(_arg1, _arg2, _arg3)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(scanner)
	runtime.KeepAlive(pixmapFile)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetDefaultFiles retrieves the current list of RC files that will be parsed
// at the end of gtk_init().
//
// Deprecated: Use StyleContext instead.
func RCGetDefaultFiles() []string {
	var _cret **C.gchar // in

	_cret = C.gtk_rc_get_default_files()

	var _filenames []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _filenames
}

// RCGetImModuleFile obtains the path to the IM modules file. See the
// documentation of the GTK_IM_MODULE_FILE environment variable for more
// details.
//
// Deprecated: Use CssProvider instead.
func RCGetImModuleFile() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_im_module_file()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetImModulePath obtains the path in which to look for IM modules. See the
// documentation of the GTK_PATH environment variable for more details about
// looking up modules. This function is useful solely for utilities supplied
// with GTK+ and should not be used by applications under normal circumstances.
//
// Deprecated: Use CssProvider instead.
func RCGetImModulePath() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_im_module_path()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetModuleDir returns a directory in which GTK+ looks for theme engines. For
// full information about the search for theme engines, see the docs for
// GTK_PATH in [Running GTK+ Applications][gtk-running].
//
// Deprecated: Use CssProvider instead.
func RCGetModuleDir() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_module_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// RCGetStyle finds all matching RC styles for a given widget, composites them
// together, and then creates a Style representing the composite appearance.
// (GTK+ actually keeps a cache of previously created styles, so a new style may
// not be created.)
//
// Deprecated: Use StyleContext instead.
func RCGetStyle(widget Widgetter) *Style {
	var _arg1 *C.GtkWidget // out
	var _cret *C.GtkStyle  // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_rc_get_style(_arg1)
	runtime.KeepAlive(widget)

	var _style *Style // out

	_style = wrapStyle(externglib.Take(unsafe.Pointer(_cret)))

	return _style
}

// RCGetStyleByPaths creates up a Style from styles defined in a RC file by
// providing the raw components used in matching. This function may be useful
// when creating pseudo-widgets that should be themed like widgets but don’t
// actually have corresponding GTK+ widgets. An example of this would be items
// inside a GNOME canvas widget.
//
// The action of gtk_rc_get_style() is similar to:
//
//    gtk_widget_path (widget, NULL, &path, NULL);
//    gtk_widget_class_path (widget, NULL, &class_path, NULL);
//    gtk_rc_get_style_by_paths (gtk_widget_get_settings (widget),
//                               path, class_path,
//                               G_OBJECT_TYPE (widget));
//
// Deprecated: Use StyleContext instead.
func RCGetStyleByPaths(settings *Settings, widgetPath string, classPath string, typ externglib.Type) *Style {
	var _arg1 *C.GtkSettings // out
	var _arg2 *C.char        // out
	var _arg3 *C.char        // out
	var _arg4 C.GType        // out
	var _cret *C.GtkStyle    // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	if widgetPath != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(widgetPath)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	if classPath != "" {
		_arg3 = (*C.char)(unsafe.Pointer(C.CString(classPath)))
		defer C.free(unsafe.Pointer(_arg3))
	}
	_arg4 = C.GType(typ)

	_cret = C.gtk_rc_get_style_by_paths(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(widgetPath)
	runtime.KeepAlive(classPath)
	runtime.KeepAlive(typ)

	var _style *Style // out

	if _cret != nil {
		_style = wrapStyle(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _style
}

// RCGetThemeDir returns the standard directory in which themes should be
// installed. (GTK+ does not actually use this directory itself.)
//
// Deprecated: Use CssProvider instead.
func RCGetThemeDir() string {
	var _cret *C.gchar // in

	_cret = C.gtk_rc_get_theme_dir()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// RCParse parses a given resource file.
//
// Deprecated: Use CssProvider instead.
func RCParse(filename string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse(_arg1)
	runtime.KeepAlive(filename)
}

// RCParseColor parses a color in the format expected in a RC file.
//
// Note that theme engines should use gtk_rc_parse_color_full() in order to
// support symbolic colors.
//
// Deprecated: Use CssProvider instead.
func RCParseColor(scanner *glib.Scanner) (gdk.Color, uint) {
	var _arg1 *C.GScanner // out
	var _arg2 C.GdkColor  // in
	var _cret C.guint     // in

	_arg1 = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))

	_cret = C.gtk_rc_parse_color(_arg1, &_arg2)
	runtime.KeepAlive(scanner)

	var _color gdk.Color // out
	var _guint uint      // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	_guint = uint(_cret)

	return _color, _guint
}

// RCParseColorFull parses a color in the format expected in a RC file. If style
// is not NULL, it will be consulted to resolve references to symbolic colors.
//
// Deprecated: Use CssProvider instead.
func RCParseColorFull(scanner *glib.Scanner, style *RCStyle) (gdk.Color, uint) {
	var _arg1 *C.GScanner   // out
	var _arg2 *C.GtkRcStyle // out
	var _arg3 C.GdkColor    // in
	var _cret C.guint       // in

	_arg1 = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))
	if style != nil {
		_arg2 = (*C.GtkRcStyle)(unsafe.Pointer(style.Native()))
	}

	_cret = C.gtk_rc_parse_color_full(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(scanner)
	runtime.KeepAlive(style)

	var _color gdk.Color // out
	var _guint uint      // out

	_color = *(*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	_guint = uint(_cret)

	return _color, _guint
}

// RCParseState parses a StateType variable from the format expected in a RC
// file.
//
// Deprecated: Use CssProvider instead.
func RCParseState(scanner *glib.Scanner) (StateType, uint) {
	var _arg1 *C.GScanner    // out
	var _arg2 C.GtkStateType // in
	var _cret C.guint        // in

	_arg1 = (*C.GScanner)(gextras.StructNative(unsafe.Pointer(scanner)))

	_cret = C.gtk_rc_parse_state(_arg1, &_arg2)
	runtime.KeepAlive(scanner)

	var _state StateType // out
	var _guint uint      // out

	_state = StateType(_arg2)
	_guint = uint(_cret)

	return _state, _guint
}

// RCParseString parses resource information directly from a string.
//
// Deprecated: Use CssProvider instead.
func RCParseString(rcString string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(rcString)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_rc_parse_string(_arg1)
	runtime.KeepAlive(rcString)
}

// RCReparseAll: if the modification time on any previously read file for the
// default Settings has changed, discard all style information and then reread
// all previously read RC files.
//
// Deprecated: Use CssProvider instead.
func RCReparseAll() bool {
	var _cret C.gboolean // in

	_cret = C.gtk_rc_reparse_all()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RCReparseAllForSettings: if the modification time on any previously read file
// for the given Settings has changed, discard all style information and then
// reread all previously read RC files.
//
// Deprecated: Use CssProvider instead.
func RCReparseAllForSettings(settings *Settings, forceLoad bool) bool {
	var _arg1 *C.GtkSettings // out
	var _arg2 C.gboolean     // out
	var _cret C.gboolean     // in

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))
	if forceLoad {
		_arg2 = C.TRUE
	}

	_cret = C.gtk_rc_reparse_all_for_settings(_arg1, _arg2)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(forceLoad)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RCResetStyles: this function recomputes the styles for all widgets that use a
// particular Settings object. (There is one Settings object per Screen, see
// gtk_settings_get_for_screen()); It is useful when some global parameter has
// changed that affects the appearance of all widgets, because when a widget
// gets a new style, it will both redraw and recompute any cached information
// about its appearance. As an example, it is used when the default font size
// set by the operating system changes. Note that this function doesn’t affect
// widgets that have a style set explicitly on them with gtk_widget_set_style().
//
// Deprecated: Use CssProvider instead.
func RCResetStyles(settings *Settings) {
	var _arg1 *C.GtkSettings // out

	_arg1 = (*C.GtkSettings)(unsafe.Pointer(settings.Native()))

	C.gtk_rc_reset_styles(_arg1)
	runtime.KeepAlive(settings)
}

// RCSetDefaultFiles sets the list of files that GTK+ will read at the end of
// gtk_init().
//
// Deprecated: Use StyleContext with a custom StyleProvider instead.
func RCSetDefaultFiles(filenames []string) {
	var _arg1 **C.gchar // out

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(filenames)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(filenames)+1)
			var zero *C.gchar
			out[len(filenames)] = zero
			for i := range filenames {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(filenames[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_rc_set_default_files(_arg1)
	runtime.KeepAlive(filenames)
}

// RCStyleOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RCStyleOverrider interface {
	Merge(src *RCStyle)
	Parse(settings *Settings, scanner *glib.Scanner) uint
}

// RCStyle is used to represent a set of information about the appearance of a
// widget. This can later be composited together with other RcStyle-struct<!--
// -->s to form a Style.
type RCStyle struct {
	*externglib.Object
}

func wrapRCStyle(obj *externglib.Object) *RCStyle {
	return &RCStyle{
		Object: obj,
	}
}

func marshalRCStyler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRCStyle(obj), nil
}

// NewRCStyle creates a new RcStyle with no fields set and a reference count of
// 1.
//
// Deprecated: Use CssProvider instead.
func NewRCStyle() *RCStyle {
	var _cret *C.GtkRcStyle // in

	_cret = C.gtk_rc_style_new()

	var _rcStyle *RCStyle // out

	_rcStyle = wrapRCStyle(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _rcStyle
}

// Copy makes a copy of the specified RcStyle. This function will correctly copy
// an RC style that is a member of a class derived from RcStyle.
//
// Deprecated: Use CssProvider instead.
func (orig *RCStyle) Copy() *RCStyle {
	var _arg0 *C.GtkRcStyle // out
	var _cret *C.GtkRcStyle // in

	_arg0 = (*C.GtkRcStyle)(unsafe.Pointer(orig.Native()))

	_cret = C.gtk_rc_style_copy(_arg0)
	runtime.KeepAlive(orig)

	var _rcStyle *RCStyle // out

	_rcStyle = wrapRCStyle(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _rcStyle
}

// RCProperty: deprecated
//
// An instance of this type is always passed by reference.
type RCProperty struct {
	*rcProperty
}

// rcProperty is the struct that's finalized.
type rcProperty struct {
	native *C.GtkRcProperty
}

// TypeName: quark-ified type identifier.
func (r *RCProperty) TypeName() glib.Quark {
	var v glib.Quark // out
	v = uint32(r.native.type_name)
	return v
}

// PropertyName: quark-ified property identifier like “GtkScrollbar::spacing”.
func (r *RCProperty) PropertyName() glib.Quark {
	var v glib.Quark // out
	v = uint32(r.native.property_name)
	return v
}

// Origin: field similar to one found in SettingsValue.
func (r *RCProperty) Origin() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(r.native.origin)))
	return v
}
