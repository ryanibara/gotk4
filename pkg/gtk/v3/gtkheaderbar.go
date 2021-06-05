// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_header_bar_get_type()), F: marshalHeaderBar},
	})
}

// HeaderBar: gtkHeaderBar is similar to a horizontal Box. It allows children to
// be placed at the start or the end. In addition, it allows a title and
// subtitle to be displayed. The title will be centered with respect to the
// width of the box, even if the children at either side take up different
// amounts of space. The height of the titlebar will be set to provide
// sufficient space for the subtitle, even if none is currently set. If a
// subtitle is not needed, the space reservation can be turned off with
// gtk_header_bar_set_has_subtitle().
//
// GtkHeaderBar can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, GtkHeaderBar is the natural choice for use as the custom
// titlebar widget of a Window (see gtk_window_set_titlebar()), as it gives
// features typical of titlebars while allowing the addition of child widgets.
type HeaderBar interface {
	Container
	Buildable

	// CustomTitle retrieves the custom title widget of the header. See
	// gtk_header_bar_set_custom_title().
	CustomTitle() Widget
	// DecorationLayout gets the decoration layout set with
	// gtk_header_bar_set_decoration_layout().
	DecorationLayout() string
	// HasSubtitle retrieves whether the header bar reserves space for a
	// subtitle, regardless if one is currently set or not.
	HasSubtitle() bool
	// ShowCloseButton returns whether this header bar shows the standard window
	// decorations.
	ShowCloseButton() bool
	// Subtitle retrieves the subtitle of the header. See
	// gtk_header_bar_set_subtitle().
	Subtitle() string
	// Title retrieves the title of the header. See gtk_header_bar_set_title().
	Title() string
	// PackEnd adds @child to @bar, packed with reference to the end of the
	// @bar.
	PackEnd(child Widget)
	// PackStart adds @child to @bar, packed with reference to the start of the
	// @bar.
	PackStart(child Widget)
	// SetCustomTitle sets a custom title for the HeaderBar.
	//
	// The title should help a user identify the current view. This supersedes
	// any title set by gtk_header_bar_set_title() or
	// gtk_header_bar_set_subtitle(). To achieve the same style as the builtin
	// title and subtitle, use the “title” and “subtitle” style classes.
	//
	// You should set the custom title to nil, for the header title label to be
	// visible again.
	SetCustomTitle(titleWidget Widget)
	// SetDecorationLayout sets the decoration layout for this header bar,
	// overriding the Settings:gtk-decoration-layout setting.
	//
	// There can be valid reasons for overriding the setting, such as a header
	// bar design that does not allow for buttons to take room on the right, or
	// only offers room for a single close button. Split header bars are another
	// example for overriding the setting.
	//
	// The format of the string is button names, separated by commas. A colon
	// separates the buttons that should appear on the left from those on the
	// right. Recognized button names are minimize, maximize, close, icon (the
	// window icon) and menu (a menu button for the fallback app menu).
	//
	// For example, “menu:minimize,maximize,close” specifies a menu on the left,
	// and minimize, maximize and close buttons on the right.
	SetDecorationLayout(layout string)
	// SetHasSubtitle sets whether the header bar should reserve space for a
	// subtitle, even if none is currently set.
	SetHasSubtitle(setting bool)
	// SetShowCloseButton sets whether this header bar shows the standard window
	// decorations, including close, maximize, and minimize.
	SetShowCloseButton(setting bool)
	// SetSubtitle sets the subtitle of the HeaderBar. The title should give a
	// user an additional detail to help him identify the current view.
	//
	// Note that GtkHeaderBar by default reserves room for the subtitle, even if
	// none is currently set. If this is not desired, set the
	// HeaderBar:has-subtitle property to false.
	SetSubtitle(subtitle string)
	// SetTitle sets the title of the HeaderBar. The title should help a user
	// identify the current view. A good title should not include the
	// application name.
	SetTitle(title string)
}

// headerBar implements the HeaderBar interface.
type headerBar struct {
	Container
	Buildable
}

var _ HeaderBar = (*headerBar)(nil)

// WrapHeaderBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapHeaderBar(obj *externglib.Object) HeaderBar {
	return HeaderBar{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHeaderBar(obj), nil
}

// NewHeaderBar constructs a class HeaderBar.
func NewHeaderBar() HeaderBar {
	var cret C.GtkHeaderBar
	var ret1 HeaderBar

	cret = C.gtk_header_bar_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(HeaderBar)

	return ret1
}

// CustomTitle retrieves the custom title widget of the header. See
// gtk_header_bar_set_custom_title().
func (b headerBar) CustomTitle() Widget {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_header_bar_get_custom_title(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// DecorationLayout gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
func (b headerBar) DecorationLayout() string {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_header_bar_get_decoration_layout(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// HasSubtitle retrieves whether the header bar reserves space for a
// subtitle, regardless if one is currently set or not.
func (b headerBar) HasSubtitle() bool {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_header_bar_get_has_subtitle(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// ShowCloseButton returns whether this header bar shows the standard window
// decorations.
func (b headerBar) ShowCloseButton() bool {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_header_bar_get_show_close_button(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Subtitle retrieves the subtitle of the header. See
// gtk_header_bar_set_subtitle().
func (b headerBar) Subtitle() string {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_header_bar_get_subtitle(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Title retrieves the title of the header. See gtk_header_bar_set_title().
func (b headerBar) Title() string {
	var arg0 *C.GtkHeaderBar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_header_bar_get_title(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// PackEnd adds @child to @bar, packed with reference to the end of the
// @bar.
func (b headerBar) PackEnd(child Widget) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_end(arg0, child)
}

// PackStart adds @child to @bar, packed with reference to the start of the
// @bar.
func (b headerBar) PackStart(child Widget) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_start(arg0, child)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes
// any title set by gtk_header_bar_set_title() or
// gtk_header_bar_set_subtitle(). To achieve the same style as the builtin
// title and subtitle, use the “title” and “subtitle” style classes.
//
// You should set the custom title to nil, for the header title label to be
// visible again.
func (b headerBar) SetCustomTitle(titleWidget Widget) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))

	C.gtk_header_bar_set_custom_title(arg0, titleWidget)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header
// bar design that does not allow for buttons to take room on the right, or
// only offers room for a single close button. Split header bars are another
// example for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the
// right. Recognized button names are minimize, maximize, close, icon (the
// window icon) and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left,
// and minimize, maximize and close buttons on the right.
func (b headerBar) SetDecorationLayout(layout string) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.gchar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(layout))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_header_bar_set_decoration_layout(arg0, layout)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
func (b headerBar) SetHasSubtitle(setting bool) {
	var arg0 *C.GtkHeaderBar
	var arg1 C.gboolean

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_header_bar_set_has_subtitle(arg0, setting)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
func (b headerBar) SetShowCloseButton(setting bool) {
	var arg0 *C.GtkHeaderBar
	var arg1 C.gboolean

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_header_bar_set_show_close_button(arg0, setting)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a
// user an additional detail to help him identify the current view.
//
// Note that GtkHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the
// HeaderBar:has-subtitle property to false.
func (b headerBar) SetSubtitle(subtitle string) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.gchar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(subtitle))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_header_bar_set_subtitle(arg0, subtitle)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the
// application name.
func (b headerBar) SetTitle(title string) {
	var arg0 *C.GtkHeaderBar
	var arg1 *C.gchar

	arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_header_bar_set_title(arg0, title)
}

type HeaderBarPrivate struct {
	native C.GtkHeaderBarPrivate
}

// WrapHeaderBarPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapHeaderBarPrivate(ptr unsafe.Pointer) *HeaderBarPrivate {
	if ptr == nil {
		return nil
	}

	return (*HeaderBarPrivate)(ptr)
}

func marshalHeaderBarPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapHeaderBarPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (h *HeaderBarPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&h.native)
}
