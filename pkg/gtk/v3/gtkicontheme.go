// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_icon_info_get_type()), F: marshalIconInfo},
		{T: externglib.Type(C.gtk_icon_theme_get_type()), F: marshalIconTheme},
	})
}

// IconInfo contains information found when looking up an icon in an icon theme.
type IconInfo interface {
	gextras.Objector

	// Free: free a IconInfo and associated information
	Free()
	// BaseScale gets the base scale for the icon. The base scale is a scale for
	// the icon that was specified by the icon theme creator. For instance an
	// icon drawn for a high-dpi screen with window scale 2 for a base size of
	// 32 will be 64 pixels tall and have a base scale of 2.
	BaseScale() int
	// BaseSize gets the base size for the icon. The base size is a size for the
	// icon that was specified by the icon theme creator. This may be different
	// than the actual size of image; an example of this is small emblem icons
	// that can be attached to a larger icon. These icons will be given the same
	// base size as the larger icons to which they are attached.
	//
	// Note that for scaled icons the base size does not include the base scale.
	BaseSize() int
	// DisplayName: this function is deprecated and always returns nil.
	DisplayName() string
	// EmbeddedRect: this function is deprecated and always returns false.
	EmbeddedRect() (gdk.Rectangle, bool)
	// Filename gets the filename for the icon. If the
	// GTK_ICON_LOOKUP_USE_BUILTIN flag was passed to
	// gtk_icon_theme_lookup_icon(), there may be no filename if a builtin icon
	// is returned; in this case, you should use
	// gtk_icon_info_get_builtin_pixbuf().
	Filename() *string
	// IsSymbolic checks if the icon is symbolic or not. This currently uses
	// only the file name and not the file contents for determining this. This
	// behaviour may change in the future.
	IsSymbolic() bool
	// LoadIconAsync: asynchronously load, render and scale an icon previously
	// looked up from the icon theme using gtk_icon_theme_lookup_icon().
	//
	// For more details, see gtk_icon_info_load_icon() which is the synchronous
	// version of this call.
	LoadIconAsync(cancellable gio.Cancellable, callback gio.AsyncReadyCallback)
	// LoadSymbolicAsync: asynchronously load, render and scale a symbolic icon
	// previously looked up from the icon theme using
	// gtk_icon_theme_lookup_icon().
	//
	// For more details, see gtk_icon_info_load_symbolic() which is the
	// synchronous version of this call.
	LoadSymbolicAsync(fg *gdk.RGBA, successColor *gdk.RGBA, warningColor *gdk.RGBA, errorColor *gdk.RGBA, cancellable gio.Cancellable, callback gio.AsyncReadyCallback)
	// LoadSymbolicForContextAsync: asynchronously load, render and scale a
	// symbolic icon previously looked up from the icon theme using
	// gtk_icon_theme_lookup_icon().
	//
	// For more details, see gtk_icon_info_load_symbolic_for_context() which is
	// the synchronous version of this call.
	LoadSymbolicForContextAsync(context StyleContext, cancellable gio.Cancellable, callback gio.AsyncReadyCallback)
	// SetRawCoordinates sets whether the coordinates returned by
	// gtk_icon_info_get_embedded_rect() and gtk_icon_info_get_attach_points()
	// should be returned in their original form as specified in the icon theme,
	// instead of scaled appropriately for the pixbuf returned by
	// gtk_icon_info_load_icon().
	//
	// Raw coordinates are somewhat strange; they are specified to be with
	// respect to the unscaled pixmap for PNG and XPM icons, but for SVG icons,
	// they are in a 1000x1000 coordinate space that is scaled to the final size
	// of the icon. You can determine if the icon is an SVG icon by using
	// gtk_icon_info_get_filename(), and seeing if it is non-nil and ends in
	// “.svg”.
	//
	// This function is provided primarily to allow compatibility wrappers for
	// older API's, and is not expected to be useful for applications.
	SetRawCoordinates(rawCoordinates bool)
}

// iconInfo implements the IconInfo interface.
type iconInfo struct {
	gextras.Objector
}

var _ IconInfo = (*iconInfo)(nil)

// WrapIconInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapIconInfo(obj *externglib.Object) IconInfo {
	return IconInfo{
		Objector: obj,
	}
}

func marshalIconInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIconInfo(obj), nil
}

// Free: free a IconInfo and associated information
func (i iconInfo) Free() {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	C.gtk_icon_info_free(_arg0)
}

// BaseScale gets the base scale for the icon. The base scale is a scale for
// the icon that was specified by the icon theme creator. For instance an
// icon drawn for a high-dpi screen with window scale 2 for a base size of
// 32 will be 64 pixels tall and have a base scale of 2.
func (i iconInfo) BaseScale() int {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gint

	_cret = C.gtk_icon_info_get_base_scale(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// BaseSize gets the base size for the icon. The base size is a size for the
// icon that was specified by the icon theme creator. This may be different
// than the actual size of image; an example of this is small emblem icons
// that can be attached to a larger icon. These icons will be given the same
// base size as the larger icons to which they are attached.
//
// Note that for scaled icons the base size does not include the base scale.
func (i iconInfo) BaseSize() int {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gint

	_cret = C.gtk_icon_info_get_base_size(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// DisplayName: this function is deprecated and always returns nil.
func (i iconInfo) DisplayName() string {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.gchar

	_cret = C.gtk_icon_info_get_display_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// EmbeddedRect: this function is deprecated and always returns false.
func (i iconInfo) EmbeddedRect() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _rectangle gdk.Rectangle
	var _cret C.gboolean

	_cret = C.gtk_icon_info_get_embedded_rect(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_rectangle)))

	var _ok bool

	if _cret {
		_ok = true
	}

	return _rectangle, _ok
}

// Filename gets the filename for the icon. If the
// GTK_ICON_LOOKUP_USE_BUILTIN flag was passed to
// gtk_icon_theme_lookup_icon(), there may be no filename if a builtin icon
// is returned; in this case, you should use
// gtk_icon_info_get_builtin_pixbuf().
func (i iconInfo) Filename() *string {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.gchar

	_cret = C.gtk_icon_info_get_filename(_arg0)

	var _filename *string

	_filename = C.GoString(_cret)

	return _filename
}

// IsSymbolic checks if the icon is symbolic or not. This currently uses
// only the file name and not the file contents for determining this. This
// behaviour may change in the future.
func (i iconInfo) IsSymbolic() bool {
	var _arg0 *C.GtkIconInfo

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	_cret = C.gtk_icon_info_is_symbolic(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// LoadIconAsync: asynchronously load, render and scale an icon previously
// looked up from the icon theme using gtk_icon_theme_lookup_icon().
//
// For more details, see gtk_icon_info_load_icon() which is the synchronous
// version of this call.
func (i iconInfo) LoadIconAsync(cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkIconInfo
	var _arg1 *C.GCancellable
	var _arg2 C.GAsyncReadyCallback
	var _arg3 C.gpointer

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.gtk_icon_info_load_icon_async(_arg0, _arg1, _arg2, _arg3)
}

// LoadSymbolicAsync: asynchronously load, render and scale a symbolic icon
// previously looked up from the icon theme using
// gtk_icon_theme_lookup_icon().
//
// For more details, see gtk_icon_info_load_symbolic() which is the
// synchronous version of this call.
func (i iconInfo) LoadSymbolicAsync(fg *gdk.RGBA, successColor *gdk.RGBA, warningColor *gdk.RGBA, errorColor *gdk.RGBA, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkIconInfo
	var _arg1 *C.GdkRGBA
	var _arg2 *C.GdkRGBA
	var _arg3 *C.GdkRGBA
	var _arg4 *C.GdkRGBA
	var _arg5 *C.GCancellable
	var _arg6 C.GAsyncReadyCallback
	var _arg7 C.gpointer

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkRGBA)(unsafe.Pointer(fg.Native()))
	_arg2 = (*C.GdkRGBA)(unsafe.Pointer(successColor.Native()))
	_arg3 = (*C.GdkRGBA)(unsafe.Pointer(warningColor.Native()))
	_arg4 = (*C.GdkRGBA)(unsafe.Pointer(errorColor.Native()))
	_arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg6 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg7 = C.gpointer(box.Assign(callback))

	C.gtk_icon_info_load_symbolic_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// LoadSymbolicForContextAsync: asynchronously load, render and scale a
// symbolic icon previously looked up from the icon theme using
// gtk_icon_theme_lookup_icon().
//
// For more details, see gtk_icon_info_load_symbolic_for_context() which is
// the synchronous version of this call.
func (i iconInfo) LoadSymbolicForContextAsync(context StyleContext, cancellable gio.Cancellable, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GtkIconInfo
	var _arg1 *C.GtkStyleContext
	var _arg2 *C.GCancellable
	var _arg3 C.GAsyncReadyCallback
	var _arg4 C.gpointer

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.gtk_icon_info_load_symbolic_for_context_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetRawCoordinates sets whether the coordinates returned by
// gtk_icon_info_get_embedded_rect() and gtk_icon_info_get_attach_points()
// should be returned in their original form as specified in the icon theme,
// instead of scaled appropriately for the pixbuf returned by
// gtk_icon_info_load_icon().
//
// Raw coordinates are somewhat strange; they are specified to be with
// respect to the unscaled pixmap for PNG and XPM icons, but for SVG icons,
// they are in a 1000x1000 coordinate space that is scaled to the final size
// of the icon. You can determine if the icon is an SVG icon by using
// gtk_icon_info_get_filename(), and seeing if it is non-nil and ends in
// “.svg”.
//
// This function is provided primarily to allow compatibility wrappers for
// older API's, and is not expected to be useful for applications.
func (i iconInfo) SetRawCoordinates(rawCoordinates bool) {
	var _arg0 *C.GtkIconInfo
	var _arg1 C.gboolean

	_arg0 = (*C.GtkIconInfo)(unsafe.Pointer(i.Native()))
	if rawCoordinates {
		_arg1 = C.gboolean(1)
	}

	C.gtk_icon_info_set_raw_coordinates(_arg0, _arg1)
}

// IconTheme provides a facility for looking up icons by name and size. The main
// reason for using a name rather than simply providing a filename is to allow
// different icons to be used depending on what “icon theme” is selected by the
// user. The operation of icon themes on Linux and Unix follows the Icon Theme
// Specification (http://www.freedesktop.org/Standards/icon-theme-spec) There is
// a fallback icon theme, named `hicolor`, where applications should install
// their icons, but additional icon themes can be installed as operating system
// vendors and users choose.
//
// Named icons are similar to the deprecated [Stock Items][gtkstock], and the
// distinction between the two may be a bit confusing. A few things to keep in
// mind:
//
// - Stock images usually are used in conjunction with [Stock Items][gtkstock],
// such as GTK_STOCK_OK or GTK_STOCK_OPEN. Named icons are easier to set up and
// therefore are more useful for new icons that an application wants to add,
// such as application icons or window icons.
//
// - Stock images can only be loaded at the symbolic sizes defined by the
// IconSize enumeration, or by custom sizes defined by gtk_icon_size_register(),
// while named icons are more flexible and any pixel size can be specified.
//
// - Because stock images are closely tied to stock items, and thus to actions
// in the user interface, stock images may come in multiple variants for
// different widget states or writing directions.
//
// A good rule of thumb is that if there is a stock image for what you want to
// use, use it, otherwise use a named icon. It turns out that internally stock
// images are generally defined in terms of one or more named icons. (An example
// of the more than one case is icons that depend on writing direction;
// GTK_STOCK_GO_FORWARD uses the two themed icons “gtk-stock-go-forward-ltr” and
// “gtk-stock-go-forward-rtl”.)
//
// In many cases, named themes are used indirectly, via Image or stock items,
// rather than directly, but looking up icons directly is also simple. The
// IconTheme object acts as a database of all the icons in the current theme.
// You can create new IconTheme objects, but it’s much more efficient to use the
// standard icon theme for the Screen so that the icon information is shared
// with other people looking up icons.
//
//    GError *error = NULL;
//    GtkIconTheme *icon_theme;
//    GdkPixbuf *pixbuf;
//
//    icon_theme = gtk_icon_theme_get_default ();
//    pixbuf = gtk_icon_theme_load_icon (icon_theme,
//                                       "my-icon-name", // icon name
//                                       48, // icon size
//                                       0,  // flags
//                                       &error);
//    if (!pixbuf)
//      {
//        g_warning ("Couldn’t load icon: s", error->message);
//        g_error_free (error);
//      }
//    else
//      {
//        // Use the pixbuf
//        g_object_unref (pixbuf);
//      }
type IconTheme interface {
	gextras.Objector

	// AddResourcePath adds a resource path that will be looked at when looking
	// for icons, similar to search paths.
	//
	// This function should be used to make application-specific icons available
	// as part of the icon theme.
	//
	// The resources are considered as part of the hicolor icon theme and must
	// be located in subdirectories that are defined in the hicolor icon theme,
	// such as `@path/16x16/actions/run.png`. Icons that are directly placed in
	// the resource path instead of a subdirectory are also considered as
	// ultimate fallback.
	AddResourcePath(path string)
	// AppendSearchPath appends a directory to the search path. See
	// gtk_icon_theme_set_search_path().
	AppendSearchPath(path *string)
	// ExampleIconName gets the name of an icon that is representative of the
	// current theme (for instance, to use when presenting a list of themes to
	// the user.)
	ExampleIconName() string
	// IconSizes returns an array of integers describing the sizes at which the
	// icon is available without scaling. A size of -1 means that the icon is
	// available in a scalable format. The array is zero-terminated.
	IconSizes(iconName string) []int
	// SearchPath gets the current search path. See
	// gtk_icon_theme_set_search_path().
	SearchPath() []*string
	// HasIcon checks whether an icon theme includes an icon for a particular
	// name.
	HasIcon(iconName string) bool
	// PrependSearchPath prepends a directory to the search path. See
	// gtk_icon_theme_set_search_path().
	PrependSearchPath(path *string)
	// RescanIfNeeded checks to see if the icon theme has changed; if it has,
	// any currently cached information is discarded and will be reloaded next
	// time @icon_theme is accessed.
	RescanIfNeeded() bool
	// SetCustomTheme sets the name of the icon theme that the IconTheme object
	// uses overriding system configuration. This function cannot be called on
	// the icon theme objects returned from gtk_icon_theme_get_default() and
	// gtk_icon_theme_get_for_screen().
	SetCustomTheme(themeName string)
	// SetScreen sets the screen for an icon theme; the screen is used to track
	// the user’s currently configured icon theme, which might be different for
	// different screens.
	SetScreen(screen gdk.Screen)
	// SetSearchPath sets the search path for the icon theme object. When
	// looking for an icon theme, GTK+ will search for a subdirectory of one or
	// more of the directories in @path with the same name as the icon theme
	// containing an index.theme file. (Themes from multiple of the path
	// elements are combined to allow themes to be extended by adding icons in
	// the user’s home directory.)
	//
	// In addition if an icon found isn’t found either in the current icon theme
	// or the default icon theme, and an image file with the right name is found
	// directly in one of the elements of @path, then that image will be used
	// for the icon name. (This is legacy feature, and new icons should be put
	// into the fallback icon theme, which is called hicolor, rather than
	// directly on the icon path.)
	SetSearchPath(path []*string)
}

// iconTheme implements the IconTheme interface.
type iconTheme struct {
	gextras.Objector
}

var _ IconTheme = (*iconTheme)(nil)

// WrapIconTheme wraps a GObject to the right type. It is
// primarily used internally.
func WrapIconTheme(obj *externglib.Object) IconTheme {
	return IconTheme{
		Objector: obj,
	}
}

func marshalIconTheme(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIconTheme(obj), nil
}

// AddResourcePath adds a resource path that will be looked at when looking
// for icons, similar to search paths.
//
// This function should be used to make application-specific icons available
// as part of the icon theme.
//
// The resources are considered as part of the hicolor icon theme and must
// be located in subdirectories that are defined in the hicolor icon theme,
// such as `@path/16x16/actions/run.png`. Icons that are directly placed in
// the resource path instead of a subdirectory are also considered as
// ultimate fallback.
func (i iconTheme) AddResourcePath(path string) {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_add_resource_path(_arg0, _arg1)
}

// AppendSearchPath appends a directory to the search path. See
// gtk_icon_theme_set_search_path().
func (i iconTheme) AppendSearchPath(path *string) {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_append_search_path(_arg0, _arg1)
}

// ExampleIconName gets the name of an icon that is representative of the
// current theme (for instance, to use when presenting a list of themes to
// the user.)
func (i iconTheme) ExampleIconName() string {
	var _arg0 *C.GtkIconTheme

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))

	var _cret *C.char

	_cret = C.gtk_icon_theme_get_example_icon_name(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// IconSizes returns an array of integers describing the sizes at which the
// icon is available without scaling. A size of -1 means that the icon is
// available in a scalable format. The array is zero-terminated.
func (i iconTheme) IconSizes(iconName string) []int {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gint

	_cret = C.gtk_icon_theme_get_icon_sizes(_arg0, _arg1)

	var _gints []int

	{
		var length int
		for p := _cret; *p != 0; p = (*C.gint)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []C.gint
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(length))

		_gints = make([]int, length)
		for i := range src {
			_gints = (int)(_cret)
		}
	}

	return _gints
}

// SearchPath gets the current search path. See
// gtk_icon_theme_set_search_path().
func (i iconTheme) SearchPath() []*string {
	var _arg0 *C.GtkIconTheme

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))

	var _arg1 **C.gchar
	var _arg2 *C.gint

	C.gtk_icon_theme_get_search_path(_arg0, &_arg1, &_arg2)

	var _path []*string

	{
		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_arg1), int(_arg2))

		_path = make([]*string, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_path = C.GoString(_arg1)
			defer C.free(unsafe.Pointer(_arg1))
		}
	}

	return _path
}

// HasIcon checks whether an icon theme includes an icon for a particular
// name.
func (i iconTheme) HasIcon(iconName string) bool {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean

	_cret = C.gtk_icon_theme_has_icon(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// PrependSearchPath prepends a directory to the search path. See
// gtk_icon_theme_set_search_path().
func (i iconTheme) PrependSearchPath(path *string) {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_prepend_search_path(_arg0, _arg1)
}

// RescanIfNeeded checks to see if the icon theme has changed; if it has,
// any currently cached information is discarded and will be reloaded next
// time @icon_theme is accessed.
func (i iconTheme) RescanIfNeeded() bool {
	var _arg0 *C.GtkIconTheme

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean

	_cret = C.gtk_icon_theme_rescan_if_needed(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetCustomTheme sets the name of the icon theme that the IconTheme object
// uses overriding system configuration. This function cannot be called on
// the icon theme objects returned from gtk_icon_theme_get_default() and
// gtk_icon_theme_get_for_screen().
func (i iconTheme) SetCustomTheme(themeName string) {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.gchar

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(themeName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_icon_theme_set_custom_theme(_arg0, _arg1)
}

// SetScreen sets the screen for an icon theme; the screen is used to track
// the user’s currently configured icon theme, which might be different for
// different screens.
func (i iconTheme) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkIconTheme
	var _arg1 *C.GdkScreen

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_icon_theme_set_screen(_arg0, _arg1)
}

// SetSearchPath sets the search path for the icon theme object. When
// looking for an icon theme, GTK+ will search for a subdirectory of one or
// more of the directories in @path with the same name as the icon theme
// containing an index.theme file. (Themes from multiple of the path
// elements are combined to allow themes to be extended by adding icons in
// the user’s home directory.)
//
// In addition if an icon found isn’t found either in the current icon theme
// or the default icon theme, and an image file with the right name is found
// directly in one of the elements of @path, then that image will be used
// for the icon name. (This is legacy feature, and new icons should be put
// into the fallback icon theme, which is called hicolor, rather than
// directly on the icon path.)
func (i iconTheme) SetSearchPath(path []*string) {
	var _arg0 *C.GtkIconTheme
	var _arg1 **C.gchar
	var _arg2 C.gint

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(i.Native()))
	_arg2 = C.gint(len(path))
	_arg1 = (**C.gchar)(C.malloc(len(path) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(_arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&out), unsafe.Pointer(_arg1), int(len(path)))

		for i := range path {
			_arg1 = (*C.gchar)(C.CString(path))
			defer C.free(unsafe.Pointer(_arg1))
		}
	}

	C.gtk_icon_theme_set_search_path(_arg0, _arg1, _arg2)
}
