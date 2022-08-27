// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_IconTheme_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gtk3_IconThemeClass_changed(GtkIconTheme*);
// void _gotk4_gtk3_IconTheme_virtual_changed(void* fnptr, GtkIconTheme* arg0) {
//   ((void (*)(GtkIconTheme*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeIconThemeError  = coreglib.Type(C.gtk_icon_theme_error_get_type())
	GTypeIconLookupFlags = coreglib.Type(C.gtk_icon_lookup_flags_get_type())
	GTypeIconInfo        = coreglib.Type(C.gtk_icon_info_get_type())
	GTypeIconTheme       = coreglib.Type(C.gtk_icon_theme_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeIconThemeError, F: marshalIconThemeError},
		coreglib.TypeMarshaler{T: GTypeIconLookupFlags, F: marshalIconLookupFlags},
		coreglib.TypeMarshaler{T: GTypeIconInfo, F: marshalIconInfo},
		coreglib.TypeMarshaler{T: GTypeIconTheme, F: marshalIconTheme},
	})
}

// IconThemeError: error codes for GtkIconTheme operations.
type IconThemeError C.gint

const (
	// IconThemeNotFound: icon specified does not exist in the theme.
	IconThemeNotFound IconThemeError = iota
	// IconThemeFailed: unspecified error occurred.
	IconThemeFailed
)

func marshalIconThemeError(p uintptr) (interface{}, error) {
	return IconThemeError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for IconThemeError.
func (i IconThemeError) String() string {
	switch i {
	case IconThemeNotFound:
		return "NotFound"
	case IconThemeFailed:
		return "Failed"
	default:
		return fmt.Sprintf("IconThemeError(%d)", i)
	}
}

// IconLookupFlags: used to specify options for gtk_icon_theme_lookup_icon().
type IconLookupFlags C.guint

const (
	// IconLookupNoSVG: never get SVG icons, even if gdk-pixbuf supports them.
	// Cannot be used together with GTK_ICON_LOOKUP_FORCE_SVG.
	IconLookupNoSVG IconLookupFlags = 0b1
	// IconLookupForceSVG: get SVG icons, even if gdk-pixbuf doesn’t support
	// them. Cannot be used together with GTK_ICON_LOOKUP_NO_SVG.
	IconLookupForceSVG IconLookupFlags = 0b10
	// IconLookupUseBuiltin: when passed to gtk_icon_theme_lookup_icon()
	// includes builtin icons as well as files. For a builtin icon,
	// gtk_icon_info_get_filename() is NULL and you need to call
	// gtk_icon_info_get_builtin_pixbuf().
	IconLookupUseBuiltin IconLookupFlags = 0b100
	// IconLookupGenericFallback: try to shorten icon name at '-' characters
	// before looking at inherited themes. This flag is only supported in
	// functions that take a single icon name. For more general fallback, see
	// gtk_icon_theme_choose_icon(). Since 2.12.
	IconLookupGenericFallback IconLookupFlags = 0b1000
	// IconLookupForceSize always get the icon scaled to the requested size.
	// Since 2.14.
	IconLookupForceSize IconLookupFlags = 0b10000
	// IconLookupForceRegular: try to always load regular icons, even when
	// symbolic icon names are given. Since 3.14.
	IconLookupForceRegular IconLookupFlags = 0b100000
	// IconLookupForceSymbolic: try to always load symbolic icons, even when
	// regular icon names are given. Since 3.14.
	IconLookupForceSymbolic IconLookupFlags = 0b1000000
	// IconLookupDirLTR: try to load a variant of the icon for left-to-right
	// text direction. Since 3.14.
	IconLookupDirLTR IconLookupFlags = 0b10000000
	// IconLookupDirRTL: try to load a variant of the icon for right-to-left
	// text direction. Since 3.14.
	IconLookupDirRTL IconLookupFlags = 0b100000000
)

func marshalIconLookupFlags(p uintptr) (interface{}, error) {
	return IconLookupFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for IconLookupFlags.
func (i IconLookupFlags) String() string {
	if i == 0 {
		return "IconLookupFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(182)

	for i != 0 {
		next := i & (i - 1)
		bit := i - next

		switch bit {
		case IconLookupNoSVG:
			builder.WriteString("NoSVG|")
		case IconLookupForceSVG:
			builder.WriteString("ForceSVG|")
		case IconLookupUseBuiltin:
			builder.WriteString("UseBuiltin|")
		case IconLookupGenericFallback:
			builder.WriteString("GenericFallback|")
		case IconLookupForceSize:
			builder.WriteString("ForceSize|")
		case IconLookupForceRegular:
			builder.WriteString("ForceRegular|")
		case IconLookupForceSymbolic:
			builder.WriteString("ForceSymbolic|")
		case IconLookupDirLTR:
			builder.WriteString("DirLTR|")
		case IconLookupDirRTL:
			builder.WriteString("DirRTL|")
		default:
			builder.WriteString(fmt.Sprintf("IconLookupFlags(0b%b)|", bit))
		}

		i = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if i contains other.
func (i IconLookupFlags) Has(other IconLookupFlags) bool {
	return (i & other) == other
}

// IconInfo contains information found when looking up an icon in an icon theme.
type IconInfo struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*IconInfo)(nil)
)

func wrapIconInfo(obj *coreglib.Object) *IconInfo {
	return &IconInfo{
		Object: obj,
	}
}

func marshalIconInfo(p uintptr) (interface{}, error) {
	return wrapIconInfo(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// IconThemeOverrides contains methods that are overridable.
type IconThemeOverrides struct {
	Changed func()
}

func defaultIconThemeOverrides(v *IconTheme) IconThemeOverrides {
	return IconThemeOverrides{
		Changed: v.changed,
	}
}

// IconTheme provides a facility for looking up icons by name and size. The main
// reason for using a name rather than simply providing a filename is to allow
// different icons to be used depending on what “icon theme” is selected by the
// user. The operation of icon themes on Linux and Unix follows the Icon Theme
// Specification (http://www.freedesktop.org/Standards/icon-theme-spec) There is
// a fallback icon theme, named hicolor, where applications should install their
// icons, but additional icon themes can be installed as operating system
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
//      }.
type IconTheme struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*IconTheme)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*IconTheme, *IconThemeClass, IconThemeOverrides](
		GTypeIconTheme,
		initIconThemeClass,
		wrapIconTheme,
		defaultIconThemeOverrides,
	)
}

func initIconThemeClass(gclass unsafe.Pointer, overrides IconThemeOverrides, classInitFunc func(*IconThemeClass)) {
	pclass := (*C.GtkIconThemeClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeIconTheme))))

	if overrides.Changed != nil {
		pclass.changed = (*[0]byte)(C._gotk4_gtk3_IconThemeClass_changed)
	}

	if classInitFunc != nil {
		class := (*IconThemeClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapIconTheme(obj *coreglib.Object) *IconTheme {
	return &IconTheme{
		Object: obj,
	}
}

func marshalIconTheme(p uintptr) (interface{}, error) {
	return wrapIconTheme(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChanged is emitted when the current icon theme is switched or GTK+
// detects that a change has occurred in the contents of the current icon theme.
func (iconTheme *IconTheme) ConnectChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(iconTheme, "changed", false, unsafe.Pointer(C._gotk4_gtk3_IconTheme_ConnectChanged), f)
}

func (iconTheme *IconTheme) changed() {
	gclass := (*C.GtkIconThemeClass)(coreglib.PeekParentClass(iconTheme))
	fnarg := gclass.changed

	var _arg0 *C.GtkIconTheme // out

	_arg0 = (*C.GtkIconTheme)(unsafe.Pointer(coreglib.InternObject(iconTheme).Native()))

	C._gotk4_gtk3_IconTheme_virtual_changed(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(iconTheme)
}

// IconThemeClass: instance of this type is always passed by reference.
type IconThemeClass struct {
	*iconThemeClass
}

// iconThemeClass is the struct that's finalized.
type iconThemeClass struct {
	native *C.GtkIconThemeClass
}
