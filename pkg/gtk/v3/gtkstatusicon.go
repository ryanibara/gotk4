// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_status_icon_get_type()), F: marshalStatusIcon},
	})
}

type StatusIconPrivate struct {
	native C.GtkStatusIconPrivate
}

// WrapStatusIconPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStatusIconPrivate(ptr unsafe.Pointer) *StatusIconPrivate {
	if ptr == nil {
		return nil
	}

	return (*StatusIconPrivate)(ptr)
}

func marshalStatusIconPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStatusIconPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StatusIconPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// StatusIcon: the “system tray” or notification area is normally used for
// transient icons that indicate some special state. For example, a system tray
// icon might appear to tell the user that they have new mail, or have an
// incoming instant message, or something along those lines. The basic idea is
// that creating an icon in the notification area is less annoying than popping
// up a dialog.
//
// A StatusIcon object can be used to display an icon in a “system tray”. The
// icon can have a tooltip, and the user can interact with it by activating it
// or popping up a context menu.
//
// It is very important to notice that status icons depend on the existence of a
// notification area being available to the user; you should not use status
// icons as the only way to convey critical information regarding your
// application, as the notification area may not exist on the user's
// environment, or may have been removed. You should always check that a status
// icon has been embedded into a notification area by using
// gtk_status_icon_is_embedded(), and gracefully recover if the function returns
// false.
//
// On X11, the implementation follows the FreeDesktop System Tray Specification
// (http://www.freedesktop.org/wiki/Specifications/systemtray-spec).
// Implementations of the “tray” side of this specification can be found e.g. in
// the GNOME 2 and KDE panel applications.
//
// Note that a GtkStatusIcon is not a widget, but just a #GObject. Making it a
// widget would be impractical, since the system tray on Windows doesn’t allow
// to embed arbitrary widgets.
//
// GtkStatusIcon has been deprecated in 3.14. You should consider using
// notifications or more modern platform-specific APIs instead. GLib provides
// the #GNotification API which works well with Application on multiple
// platforms and environments, and should be the preferred mechanism to notify
// the users of transient status updates. See this HowDoI
// (https://wiki.gnome.org/HowDoI/GNotification) for code examples.
type StatusIcon interface {
	gextras.Objector

	// Geometry obtains information about the location of the status icon on
	// screen. This information can be used to e.g. position popups like
	// notification bubbles.
	//
	// See gtk_status_icon_position_menu() for a more convenient alternative for
	// positioning menus.
	//
	// Note that some platforms do not allow GTK+ to provide this information,
	// and even on platforms that do allow it, the information is not reliable
	// unless the status icon is embedded in a notification area, see
	// gtk_status_icon_is_embedded().
	Geometry() (screen gdk.Screen, area gdk.Rectangle, orientation Orientation, ok bool)
	// GIcon retrieves the #GIcon being displayed by the StatusIcon. The storage
	// type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
	// gtk_status_icon_get_storage_type()). The caller of this function does not
	// own a reference to the returned #GIcon.
	//
	// If this function fails, @icon is left unchanged;
	GIcon() gio.Icon
	// HasTooltip returns the current value of the has-tooltip property. See
	// StatusIcon:has-tooltip for more information.
	HasTooltip() bool
	// IconName gets the name of the icon being displayed by the StatusIcon. The
	// storage type of the status icon must be GTK_IMAGE_EMPTY or
	// GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()). The
	// returned string is owned by the StatusIcon and should not be freed or
	// modified.
	IconName() string
	// Pixbuf gets the Pixbuf being displayed by the StatusIcon. The storage
	// type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
	// gtk_status_icon_get_storage_type()). The caller of this function does not
	// own a reference to the returned pixbuf.
	Pixbuf() gdkpixbuf.Pixbuf
	// Screen returns the Screen associated with @status_icon.
	Screen() gdk.Screen
	// Size gets the size in pixels that is available for the image. Stock icons
	// and named icons adapt their size automatically if the size of the
	// notification area changes. For other storage types, the size-changed
	// signal can be used to react to size changes.
	//
	// Note that the returned size is only meaningful while the status icon is
	// embedded (see gtk_status_icon_is_embedded()).
	Size() int
	// Stock gets the id of the stock icon being displayed by the StatusIcon.
	// The storage type of the status icon must be GTK_IMAGE_EMPTY or
	// GTK_IMAGE_STOCK (see gtk_status_icon_get_storage_type()). The returned
	// string is owned by the StatusIcon and should not be freed or modified.
	Stock() string
	// StorageType gets the type of representation being used by the StatusIcon
	// to store image data. If the StatusIcon has no image data, the return
	// value will be GTK_IMAGE_EMPTY.
	StorageType() ImageType
	// Title gets the title of this tray icon. See gtk_status_icon_set_title().
	Title() string
	// TooltipMarkup gets the contents of the tooltip for @status_icon.
	TooltipMarkup() string
	// TooltipText gets the contents of the tooltip for @status_icon.
	TooltipText() string
	// Visible returns whether the status icon is visible or not. Note that
	// being visible does not guarantee that the user can actually see the icon,
	// see also gtk_status_icon_is_embedded().
	Visible() bool
	// X11WindowID: this function is only useful on the X11/freedesktop.org
	// platform.
	//
	// It returns a window ID for the widget in the underlying status icon
	// implementation. This is useful for the Galago notification service, which
	// can send a window ID in the protocol in order for the server to position
	// notification windows pointing to a status icon reliably.
	//
	// This function is not intended for other use cases which are more likely
	// to be met by one of the non-X11 specific methods, such as
	// gtk_status_icon_position_menu().
	X11WindowID() uint32
	// IsEmbedded returns whether the status icon is embedded in a notification
	// area.
	IsEmbedded() bool
	// SetFromFile makes @status_icon display the file @filename. See
	// gtk_status_icon_new_from_file() for details.
	SetFromFile(filename string)
	// SetFromGIcon makes @status_icon display the #GIcon. See
	// gtk_status_icon_new_from_gicon() for details.
	SetFromGIcon(icon gio.Icon)
	// SetFromIconName makes @status_icon display the icon named @icon_name from
	// the current icon theme. See gtk_status_icon_new_from_icon_name() for
	// details.
	SetFromIconName(iconName string)
	// SetFromPixbuf makes @status_icon display @pixbuf. See
	// gtk_status_icon_new_from_pixbuf() for details.
	SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf)
	// SetFromStock makes @status_icon display the stock icon with the id
	// @stock_id. See gtk_status_icon_new_from_stock() for details.
	SetFromStock(stockID string)
	// SetHasTooltip sets the has-tooltip property on @status_icon to
	// @has_tooltip. See StatusIcon:has-tooltip for more information.
	SetHasTooltip(hasTooltip bool)
	// SetName sets the name of this tray icon. This should be a string
	// identifying this icon. It is may be used for sorting the icons in the
	// tray and will not be shown to the user.
	SetName(name string)
	// SetScreen sets the Screen where @status_icon is displayed; if the icon is
	// already mapped, it will be unmapped, and then remapped on the new screen.
	SetScreen(screen gdk.Screen)
	// SetTitle sets the title of this tray icon. This should be a short,
	// human-readable, localized string describing the tray icon. It may be used
	// by tools like screen readers to render the tray icon.
	SetTitle(title string)
	// SetTooltipMarkup sets @markup as the contents of the tooltip, which is
	// marked up with the [Pango text markup language][PangoMarkupFormat].
	//
	// This function will take care of setting StatusIcon:has-tooltip to true
	// and of the default handler for the StatusIcon::query-tooltip signal.
	//
	// See also the StatusIcon:tooltip-markup property and
	// gtk_tooltip_set_markup().
	SetTooltipMarkup(markup string)
	// SetTooltipText sets @text as the contents of the tooltip.
	//
	// This function will take care of setting StatusIcon:has-tooltip to true
	// and of the default handler for the StatusIcon::query-tooltip signal.
	//
	// See also the StatusIcon:tooltip-text property and gtk_tooltip_set_text().
	SetTooltipText(text string)
	// SetVisible shows or hides a status icon.
	SetVisible(visible bool)
}

// statusIcon implements the StatusIcon interface.
type statusIcon struct {
	gextras.Objector
}

var _ StatusIcon = (*statusIcon)(nil)

// WrapStatusIcon wraps a GObject to the right type. It is
// primarily used internally.
func WrapStatusIcon(obj *externglib.Object) StatusIcon {
	return StatusIcon{
		Objector: obj,
	}
}

func marshalStatusIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStatusIcon(obj), nil
}

// NewStatusIcon constructs a class StatusIcon.
func NewStatusIcon() StatusIcon {
	ret := C.gtk_status_icon_new()

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// NewStatusIconFromFile constructs a class StatusIcon.
func NewStatusIconFromFile(filename string) StatusIcon {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_status_icon_new_from_file(arg1)

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// NewStatusIconFromGIcon constructs a class StatusIcon.
func NewStatusIconFromGIcon(icon gio.Icon) StatusIcon {
	var arg1 *C.GIcon

	arg1 = (*C.GIcon)(icon.Native())

	ret := C.gtk_status_icon_new_from_gicon(arg1)

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// NewStatusIconFromIconName constructs a class StatusIcon.
func NewStatusIconFromIconName(iconName string) StatusIcon {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_status_icon_new_from_icon_name(arg1)

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// NewStatusIconFromPixbuf constructs a class StatusIcon.
func NewStatusIconFromPixbuf(pixbuf gdkpixbuf.Pixbuf) StatusIcon {
	var arg1 *C.GdkPixbuf

	arg1 = (*C.GdkPixbuf)(pixbuf.Native())

	ret := C.gtk_status_icon_new_from_pixbuf(arg1)

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// NewStatusIconFromStock constructs a class StatusIcon.
func NewStatusIconFromStock(stockID string) StatusIcon {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_status_icon_new_from_stock(arg1)

	var ret0 StatusIcon

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StatusIcon)

	return ret0
}

// Geometry obtains information about the location of the status icon on
// screen. This information can be used to e.g. position popups like
// notification bubbles.
//
// See gtk_status_icon_position_menu() for a more convenient alternative for
// positioning menus.
//
// Note that some platforms do not allow GTK+ to provide this information,
// and even on platforms that do allow it, the information is not reliable
// unless the status icon is embedded in a notification area, see
// gtk_status_icon_is_embedded().
func (s statusIcon) Geometry() (screen gdk.Screen, area gdk.Rectangle, orientation Orientation, ok bool) {
	var arg0 *C.GtkStatusIcon
	var arg1 **C.GdkScreen     // out
	var arg2 *C.GdkRectangle   // out
	var arg3 *C.GtkOrientation // out

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_geometry(arg0, &arg1, &arg2, &arg3)

	var ret0 gdk.Screen
	var ret1 *gdk.Rectangle
	var ret2 *Orientation
	var ret3 bool

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1.Native()))).(gdk.Screen)

	{
		ret1 = gdk.WrapRectangle(unsafe.Pointer(arg2))
	}

	ret2 = (*Orientation)(arg3)

	ret3 = C.bool(ret) != 0

	return ret0, ret1, ret2, ret3
}

// GIcon retrieves the #GIcon being displayed by the StatusIcon. The storage
// type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_GICON (see
// gtk_status_icon_get_storage_type()). The caller of this function does not
// own a reference to the returned #GIcon.
//
// If this function fails, @icon is left unchanged;
func (s statusIcon) GIcon() gio.Icon {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_gicon(arg0)

	var ret0 gio.Icon

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gio.Icon)

	return ret0
}

// HasTooltip returns the current value of the has-tooltip property. See
// StatusIcon:has-tooltip for more information.
func (s statusIcon) HasTooltip() bool {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_has_tooltip(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// IconName gets the name of the icon being displayed by the StatusIcon. The
// storage type of the status icon must be GTK_IMAGE_EMPTY or
// GTK_IMAGE_ICON_NAME (see gtk_status_icon_get_storage_type()). The
// returned string is owned by the StatusIcon and should not be freed or
// modified.
func (s statusIcon) IconName() string {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_icon_name(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Pixbuf gets the Pixbuf being displayed by the StatusIcon. The storage
// type of the status icon must be GTK_IMAGE_EMPTY or GTK_IMAGE_PIXBUF (see
// gtk_status_icon_get_storage_type()). The caller of this function does not
// own a reference to the returned pixbuf.
func (s statusIcon) Pixbuf() gdkpixbuf.Pixbuf {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_pixbuf(arg0)

	var ret0 gdkpixbuf.Pixbuf

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gdkpixbuf.Pixbuf)

	return ret0
}

// Screen returns the Screen associated with @status_icon.
func (s statusIcon) Screen() gdk.Screen {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_screen(arg0)

	var ret0 gdk.Screen

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(gdk.Screen)

	return ret0
}

// Size gets the size in pixels that is available for the image. Stock icons
// and named icons adapt their size automatically if the size of the
// notification area changes. For other storage types, the size-changed
// signal can be used to react to size changes.
//
// Note that the returned size is only meaningful while the status icon is
// embedded (see gtk_status_icon_is_embedded()).
func (s statusIcon) Size() int {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_size(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// Stock gets the id of the stock icon being displayed by the StatusIcon.
// The storage type of the status icon must be GTK_IMAGE_EMPTY or
// GTK_IMAGE_STOCK (see gtk_status_icon_get_storage_type()). The returned
// string is owned by the StatusIcon and should not be freed or modified.
func (s statusIcon) Stock() string {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_stock(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// StorageType gets the type of representation being used by the StatusIcon
// to store image data. If the StatusIcon has no image data, the return
// value will be GTK_IMAGE_EMPTY.
func (s statusIcon) StorageType() ImageType {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_storage_type(arg0)

	var ret0 ImageType

	ret0 = ImageType(ret)

	return ret0
}

// Title gets the title of this tray icon. See gtk_status_icon_set_title().
func (s statusIcon) Title() string {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_title(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// TooltipMarkup gets the contents of the tooltip for @status_icon.
func (s statusIcon) TooltipMarkup() string {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_tooltip_markup(arg0)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// TooltipText gets the contents of the tooltip for @status_icon.
func (s statusIcon) TooltipText() string {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_tooltip_text(arg0)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// Visible returns whether the status icon is visible or not. Note that
// being visible does not guarantee that the user can actually see the icon,
// see also gtk_status_icon_is_embedded().
func (s statusIcon) Visible() bool {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_visible(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// X11WindowID: this function is only useful on the X11/freedesktop.org
// platform.
//
// It returns a window ID for the widget in the underlying status icon
// implementation. This is useful for the Galago notification service, which
// can send a window ID in the protocol in order for the server to position
// notification windows pointing to a status icon reliably.
//
// This function is not intended for other use cases which are more likely
// to be met by one of the non-X11 specific methods, such as
// gtk_status_icon_position_menu().
func (s statusIcon) X11WindowID() uint32 {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_get_x11_window_id(arg0)

	var ret0 uint32

	ret0 = uint32(ret)

	return ret0
}

// IsEmbedded returns whether the status icon is embedded in a notification
// area.
func (s statusIcon) IsEmbedded() bool {
	var arg0 *C.GtkStatusIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())

	ret := C.gtk_status_icon_is_embedded(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// SetFromFile makes @status_icon display the file @filename. See
// gtk_status_icon_new_from_file() for details.
func (s statusIcon) SetFromFile(filename string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_from_file(arg0, arg1)
}

// SetFromGIcon makes @status_icon display the #GIcon. See
// gtk_status_icon_new_from_gicon() for details.
func (s statusIcon) SetFromGIcon(icon gio.Icon) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.GIcon

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.GIcon)(icon.Native())

	C.gtk_status_icon_set_from_gicon(arg0, arg1)
}

// SetFromIconName makes @status_icon display the icon named @icon_name from
// the current icon theme. See gtk_status_icon_new_from_icon_name() for
// details.
func (s statusIcon) SetFromIconName(iconName string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_from_icon_name(arg0, arg1)
}

// SetFromPixbuf makes @status_icon display @pixbuf. See
// gtk_status_icon_new_from_pixbuf() for details.
func (s statusIcon) SetFromPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.GdkPixbuf)(pixbuf.Native())

	C.gtk_status_icon_set_from_pixbuf(arg0, arg1)
}

// SetFromStock makes @status_icon display the stock icon with the id
// @stock_id. See gtk_status_icon_new_from_stock() for details.
func (s statusIcon) SetFromStock(stockID string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(stockID))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_from_stock(arg0, arg1)
}

// SetHasTooltip sets the has-tooltip property on @status_icon to
// @has_tooltip. See StatusIcon:has-tooltip for more information.
func (s statusIcon) SetHasTooltip(hasTooltip bool) {
	var arg0 *C.GtkStatusIcon
	var arg1 C.gboolean

	arg0 = (*C.GtkStatusIcon)(s.Native())
	if hasTooltip {
		arg1 = C.TRUE
	}

	C.gtk_status_icon_set_has_tooltip(arg0, arg1)
}

// SetName sets the name of this tray icon. This should be a string
// identifying this icon. It is may be used for sorting the icons in the
// tray and will not be shown to the user.
func (s statusIcon) SetName(name string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_name(arg0, arg1)
}

// SetScreen sets the Screen where @status_icon is displayed; if the icon is
// already mapped, it will be unmapped, and then remapped on the new screen.
func (s statusIcon) SetScreen(screen gdk.Screen) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.GdkScreen

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.GdkScreen)(screen.Native())

	C.gtk_status_icon_set_screen(arg0, arg1)
}

// SetTitle sets the title of this tray icon. This should be a short,
// human-readable, localized string describing the tray icon. It may be used
// by tools like screen readers to render the tray icon.
func (s statusIcon) SetTitle(title string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_title(arg0, arg1)
}

// SetTooltipMarkup sets @markup as the contents of the tooltip, which is
// marked up with the [Pango text markup language][PangoMarkupFormat].
//
// This function will take care of setting StatusIcon:has-tooltip to true
// and of the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-markup property and
// gtk_tooltip_set_markup().
func (s statusIcon) SetTooltipMarkup(markup string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(markup))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_tooltip_markup(arg0, arg1)
}

// SetTooltipText sets @text as the contents of the tooltip.
//
// This function will take care of setting StatusIcon:has-tooltip to true
// and of the default handler for the StatusIcon::query-tooltip signal.
//
// See also the StatusIcon:tooltip-text property and gtk_tooltip_set_text().
func (s statusIcon) SetTooltipText(text string) {
	var arg0 *C.GtkStatusIcon
	var arg1 *C.gchar

	arg0 = (*C.GtkStatusIcon)(s.Native())
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_status_icon_set_tooltip_text(arg0, arg1)
}

// SetVisible shows or hides a status icon.
func (s statusIcon) SetVisible(visible bool) {
	var arg0 *C.GtkStatusIcon
	var arg1 C.gboolean

	arg0 = (*C.GtkStatusIcon)(s.Native())
	if visible {
		arg1 = C.TRUE
	}

	C.gtk_status_icon_set_visible(arg0, arg1)
}
