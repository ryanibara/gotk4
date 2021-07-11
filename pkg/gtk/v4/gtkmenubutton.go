// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_button_get_type()), F: marshalMenuButtonner},
	})
}

// MenuButtonCreatePopupFunc: user-provided callback function to create a popup
// for a `GtkMenuButton` on demand.
//
// This function is called when the popup of @menu_button is shown, but none has
// been provided via [method@Gtk.MenuButton.set_popover] or
// [method@Gtk.MenuButton.set_menu_model].
type MenuButtonCreatePopupFunc func(menuButton *MenuButton, userData interface{})

//export gotk4_MenuButtonCreatePopupFunc
func gotk4_MenuButtonCreatePopupFunc(arg0 *C.GtkMenuButton, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var menuButton *MenuButton // out
	var userData interface{}   // out

	menuButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*MenuButton)
	userData = box.Get(uintptr(arg1))

	fn := v.(MenuButtonCreatePopupFunc)
	fn(menuButton, userData)
}

// MenuButtonner describes MenuButton's methods.
type MenuButtonner interface {
	// Direction returns the direction the popup will be pointing at when popped
	// up.
	Direction() ArrowType
	// HasFrame returns whether the button has a frame.
	HasFrame() bool
	// IconName gets the name of the icon shown in the button.
	IconName() string
	// Label gets the label shown in the button
	Label() string
	// MenuModel returns the `GMenuModel` used to generate the popup.
	MenuModel() *gio.MenuModel
	// Popover returns the `GtkPopover` that pops out of the button.
	Popover() *Popover
	// UseUnderline returns whether an embedded underline in the text indicates
	// a mnemonic.
	UseUnderline() bool
	// Popdown dismiss the menu.
	Popdown()
	// Popup: pop up the menu.
	Popup()
	// SetHasFrame sets the style of the button.
	SetHasFrame(hasFrame bool)
	// SetIconName sets the name of an icon to show inside the menu button.
	SetIconName(iconName string)
	// SetLabel sets the label to show inside the menu button.
	SetLabel(label string)
	// SetMenuModel sets the `GMenuModel` from which the popup will be
	// constructed.
	SetMenuModel(menuModel gio.MenuModeller)
	// SetPopover sets the `GtkPopover` that will be popped up when the
	// @menu_button is clicked.
	SetPopover(popover Widgetter)
	// SetUseUnderline: if true, an underline in the text indicates a mnemonic.
	SetUseUnderline(useUnderline bool)
}

// MenuButton: `GtkMenuButton` widget is used to display a popup when clicked.
//
// !An example GtkMenuButton (menu-button.png)
//
// This popup can be provided either as a `GtkPopover` or as an abstract
// `GMenuModel`.
//
// The `GtkMenuButton` widget can show either an icon (set with the
// [property@Gtk.MenuButton:icon-name] property) or a label (set with the
// [property@Gtk.MenuButton:label] property). If neither is explicitly set, a
// [class@Gtk.Image] is automatically created, using an arrow image oriented
// according to [property@Gtk.MenuButton:direction] or the generic
// “open-menu-symbolic” icon if the direction is not set.
//
// The positioning of the popup is determined by the
// [property@Gtk.MenuButton:direction] property of the menu button.
//
// For menus, the [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign]
// properties of the menu are also taken into account. For example, when the
// direction is GTK_ARROW_DOWN and the horizontal alignment is GTK_ALIGN_START,
// the menu will be positioned below the button, with the starting edge
// (depending on the text direction) of the menu aligned with the starting edge
// of the button. If there is not enough space below the button, the menu is
// popped up above the button instead. If the alignment would move part of the
// menu offscreen, it is “pushed in”.
//
// | | start | center | end | | - | --- | --- | --- | | **down** | !
// (down-start.png) | ! (down-center.png) | ! (down-end.png) | | **up** | !
// (up-start.png) | ! (up-center.png) | ! (up-end.png) | | **left** | !
// (left-start.png) | ! (left-center.png) | ! (left-end.png) | | **right** | !
// (right-start.png) | ! (right-center.png) | ! (right-end.png) |
//
//
// CSS nodes
//
// “` menubutton ╰── button.toggle ╰── <content> ╰── [arrow] “`
//
// `GtkMenuButton` has a single CSS node with name `menubutton` which contains a
// `button` node with a `.toggle` style class.
//
// Inside the toggle button content, there is an `arrow` node for the indicator,
// which will carry one of the `.none`, `.up`, `.down`, `.left` or `.right`
// style classes to indicate the direction that the menu will appear in. The CSS
// is expected to provide a suitable image for each of these cases using the
// `-gtk-icon-source` property.
//
// Optionally, the `menubutton` node can carry the `.circular` style class to
// request a round appearance.
//
//
// Accessibility
//
// `GtkMenuButton` uses the K_ACCESSIBLE_ROLE_BUTTON role.
type MenuButton struct {
	Widget
}

var (
	_ MenuButtonner   = (*MenuButton)(nil)
	_ gextras.Nativer = (*MenuButton)(nil)
)

func wrapMenuButton(obj *externglib.Object) MenuButtonner {
	return &MenuButton{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
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
	}
}

func marshalMenuButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuButton(obj), nil
}

// NewMenuButton creates a new `GtkMenuButton` widget with downwards-pointing
// arrow as the only child.
//
// You can replace the child widget with another `GtkWidget` should you wish to.
func NewMenuButton() *MenuButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_menu_button_new()

	var _menuButton *MenuButton // out

	_menuButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MenuButton)

	return _menuButton
}

// Direction returns the direction the popup will be pointing at when popped up.
func (menuButton *MenuButton) Direction() ArrowType {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.GtkArrowType   // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_direction(_arg0)

	var _arrowType ArrowType // out

	_arrowType = (ArrowType)(_cret)

	return _arrowType
}

// HasFrame returns whether the button has a frame.
func (menuButton *MenuButton) HasFrame() bool {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconName gets the name of the icon shown in the button.
func (menuButton *MenuButton) IconName() string {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Label gets the label shown in the button
func (menuButton *MenuButton) Label() string {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.char          // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// MenuModel returns the `GMenuModel` used to generate the popup.
func (menuButton *MenuButton) MenuModel() *gio.MenuModel {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GMenuModel    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_menu_model(_arg0)

	var _menuModel *gio.MenuModel // out

	_menuModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.MenuModel)

	return _menuModel
}

// Popover returns the `GtkPopover` that pops out of the button.
//
// If the button is not using a `GtkPopover`, this function returns nil.
func (menuButton *MenuButton) Popover() *Popover {
	var _arg0 *C.GtkMenuButton // out
	var _cret *C.GtkPopover    // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_popover(_arg0)

	var _popover *Popover // out

	_popover = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Popover)

	return _popover
}

// UseUnderline returns whether an embedded underline in the text indicates a
// mnemonic.
func (menuButton *MenuButton) UseUnderline() bool {
	var _arg0 *C.GtkMenuButton // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	_cret = C.gtk_menu_button_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Popdown dismiss the menu.
func (menuButton *MenuButton) Popdown() {
	var _arg0 *C.GtkMenuButton // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	C.gtk_menu_button_popdown(_arg0)
}

// Popup: pop up the menu.
func (menuButton *MenuButton) Popup() {
	var _arg0 *C.GtkMenuButton // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))

	C.gtk_menu_button_popup(_arg0)
}

// SetHasFrame sets the style of the button.
func (menuButton *MenuButton) SetHasFrame(hasFrame bool) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if hasFrame {
		_arg1 = C.TRUE
	}

	C.gtk_menu_button_set_has_frame(_arg0, _arg1)
}

// SetIconName sets the name of an icon to show inside the menu button.
func (menuButton *MenuButton) SetIconName(iconName string) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_button_set_icon_name(_arg0, _arg1)
}

// SetLabel sets the label to show inside the menu button.
func (menuButton *MenuButton) SetLabel(label string) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.char          // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_button_set_label(_arg0, _arg1)
}

// SetMenuModel sets the `GMenuModel` from which the popup will be constructed.
//
// If @menu_model is nil, the button is disabled.
//
// A [class@Gtk.Popover] will be created from the menu model with
// [ctor@Gtk.PopoverMenu.new_from_model]. Actions will be connected as
// documented for this function.
//
// If [property@Gtk.MenuButton:popover] is already set, it will be dissociated
// from the @menu_button, and the property is set to nil.
func (menuButton *MenuButton) SetMenuModel(menuModel gio.MenuModeller) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GMenuModel    // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer((menuModel).(gextras.Nativer).Native()))

	C.gtk_menu_button_set_menu_model(_arg0, _arg1)
}

// SetPopover sets the `GtkPopover` that will be popped up when the @menu_button
// is clicked.
//
// If @popover is nil, the button is disabled.
//
// If [property@Gtk.MenuButton:menu-model] is set, the menu model is dissociated
// from the @menu_button, and the property is set to nil.
func (menuButton *MenuButton) SetPopover(popover Widgetter) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((popover).(gextras.Nativer).Native()))

	C.gtk_menu_button_set_popover(_arg0, _arg1)
}

// SetUseUnderline: if true, an underline in the text indicates a mnemonic.
func (menuButton *MenuButton) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkMenuButton // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkMenuButton)(unsafe.Pointer(menuButton.Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_menu_button_set_use_underline(_arg0, _arg1)
}
