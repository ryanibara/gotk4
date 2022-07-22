// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GType values.
var (
	GTypeHeaderBar = coreglib.Type(C.gtk_header_bar_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeHeaderBar, F: marshalHeaderBar},
	})
}

// HeaderBarOverrider contains methods that are overridable.
type HeaderBarOverrider interface {
}

// HeaderBar is similar to a horizontal Box. It allows children to be placed at
// the start or the end. In addition, it allows a title and subtitle to be
// displayed. The title will be centered with respect to the width of the box,
// even if the children at either side take up different amounts of space. The
// height of the titlebar will be set to provide sufficient space for the
// subtitle, even if none is currently set. If a subtitle is not needed, the
// space reservation can be turned off with gtk_header_bar_set_has_subtitle().
//
// GtkHeaderBar can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, GtkHeaderBar is the natural choice for use as the custom
// titlebar widget of a Window (see gtk_window_set_titlebar()), as it gives
// features typical of titlebars while allowing the addition of child widgets.
type HeaderBar struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*HeaderBar)(nil)
)

func init() {
	coreglib.RegisterClassInfo(coreglib.ClassTypeInfo{
		GType:     GTypeHeaderBar,
		GoType:    reflect.TypeOf((*HeaderBar)(nil)),
		InitClass: initClassHeaderBar,
	})
}

func initClassHeaderBar(gclass unsafe.Pointer, goval any) {
	if goval, ok := goval.(interface{ InitHeaderBar(*HeaderBarClass) }); ok {
		klass := (*HeaderBarClass)(gextras.NewStructNative(gclass))
		goval.InitHeaderBar(klass)
	}
}

func wrapHeaderBar(obj *coreglib.Object) *HeaderBar {
	return &HeaderBar{
		Container: Container{
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
		},
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	return wrapHeaderBar(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewHeaderBar creates a new HeaderBar widget.
//
// The function returns the following values:
//
//    - headerBar: new HeaderBar.
//
func NewHeaderBar() *HeaderBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_header_bar_new()

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(coreglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// CustomTitle retrieves the custom title widget of the header. See
// gtk_header_bar_set_custom_title().
//
// The function returns the following values:
//
//    - widget (optional): custom title widget of the header, or NULL if none has
//      been set explicitly.
//
func (bar *HeaderBar) CustomTitle() Widgetter {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_custom_title(_arg0)
	runtime.KeepAlive(bar)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// DecorationLayout gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
//
// The function returns the following values:
//
//    - utf8: decoration layout.
//
func (bar *HeaderBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)
	runtime.KeepAlive(bar)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// HasSubtitle retrieves whether the header bar reserves space for a subtitle,
// regardless if one is currently set or not.
//
// The function returns the following values:
//
//    - ok: TRUE if the header bar reserves space for a subtitle.
//
func (bar *HeaderBar) HasSubtitle() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_has_subtitle(_arg0)
	runtime.KeepAlive(bar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether this header bar shows the standard window
// decorations.
//
// The function returns the following values:
//
//    - ok: TRUE if the decorations are shown.
//
func (bar *HeaderBar) ShowCloseButton() bool {
	var _arg0 *C.GtkHeaderBar // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_show_close_button(_arg0)
	runtime.KeepAlive(bar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Subtitle retrieves the subtitle of the header. See
// gtk_header_bar_set_subtitle().
//
// The function returns the following values:
//
//    - utf8 (optional): subtitle of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (bar *HeaderBar) Subtitle() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_subtitle(_arg0)
	runtime.KeepAlive(bar)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Title retrieves the title of the header. See gtk_header_bar_set_title().
//
// The function returns the following values:
//
//    - utf8 (optional): title of the header, or NULL if none has been set
//      explicitly. The returned string is owned by the widget and must not be
//      modified or freed.
//
func (bar *HeaderBar) Title() string {
	var _arg0 *C.GtkHeaderBar // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))

	_cret = C.gtk_header_bar_get_title(_arg0)
	runtime.KeepAlive(bar)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// PackEnd adds child to bar, packed with reference to the end of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackEnd(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_header_bar_pack_end(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// PackStart adds child to bar, packed with reference to the start of the bar.
//
// The function takes the following parameters:
//
//    - child to be added to bar.
//
func (bar *HeaderBar) PackStart(child Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_header_bar_pack_start(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(child)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes any
// title set by gtk_header_bar_set_title() or gtk_header_bar_set_subtitle(). To
// achieve the same style as the builtin title and subtitle, use the “title” and
// “subtitle” style classes.
//
// You should set the custom title to NULL, for the header title label to be
// visible again.
//
// The function takes the following parameters:
//
//    - titleWidget (optional): custom widget to use for a title.
//
func (bar *HeaderBar) SetCustomTitle(titleWidget Widgetter) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if titleWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(titleWidget).Native()))
	}

	C.gtk_header_bar_set_custom_title(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(titleWidget)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header bar
// design that does not allow for buttons to take room on the right, or only
// offers room for a single close button. Split header bars are another example
// for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the right.
// Recognized button names are minimize, maximize, close, icon (the window icon)
// and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left, and
// minimize, maximize and close buttons on the right.
//
// The function takes the following parameters:
//
//    - layout (optional): decoration layout, or NULL to unset the layout.
//
func (bar *HeaderBar) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if layout != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(layout)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(layout)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
//
// The function takes the following parameters:
//
//    - setting: TRUE to reserve space for a subtitle.
//
func (bar *HeaderBar) SetHasSubtitle(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_has_subtitle(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
//
// The function takes the following parameters:
//
//    - setting: TRUE to show standard window decorations.
//
func (bar *HeaderBar) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_header_bar_set_show_close_button(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(setting)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a user
// an additional detail to help him identify the current view.
//
// Note that GtkHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the HeaderBar:has-subtitle
// property to FALSE.
//
// The function takes the following parameters:
//
//    - subtitle (optional): subtitle, or NULL.
//
func (bar *HeaderBar) SetSubtitle(subtitle string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if subtitle != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(subtitle)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_subtitle(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(subtitle)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the application
// name.
//
// The function takes the following parameters:
//
//    - title (optional): title, or NULL.
//
func (bar *HeaderBar) SetTitle(title string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(coreglib.InternObject(bar).Native()))
	if title != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_header_bar_set_title(_arg0, _arg1)
	runtime.KeepAlive(bar)
	runtime.KeepAlive(title)
}

// HeaderBarClass: instance of this type is always passed by reference.
type HeaderBarClass struct {
	*headerBarClass
}

// headerBarClass is the struct that's finalized.
type headerBarClass struct {
	native *C.GtkHeaderBarClass
}

func (h *HeaderBarClass) ParentClass() *ContainerClass {
	valptr := &h.native.parent_class
	var v *ContainerClass // out
	v = (*ContainerClass)(gextras.NewStructNative(unsafe.Pointer((&*valptr))))
	return v
}
