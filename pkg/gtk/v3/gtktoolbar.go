// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_ToolbarClass_popup_context_menu(GtkToolbar*, gint, gint, gint);
// extern gboolean _gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd(gpointer, gboolean, guintptr);
// extern gboolean _gotk4_gtk3_Toolbar_ConnectPopupContextMenu(gpointer, gint, gint, gint, guintptr);
// extern void _gotk4_gtk3_ToolbarClass_orientation_changed(GtkToolbar*, GtkOrientation);
// extern void _gotk4_gtk3_ToolbarClass_style_changed(GtkToolbar*, GtkToolbarStyle);
// extern void _gotk4_gtk3_Toolbar_ConnectOrientationChanged(gpointer, GtkOrientation, guintptr);
// extern void _gotk4_gtk3_Toolbar_ConnectStyleChanged(gpointer, GtkToolbarStyle, guintptr);
import "C"

// glib.Type values for gtktoolbar.go.
var (
	GTypeToolbarSpaceStyle = externglib.Type(C.gtk_toolbar_space_style_get_type())
	GTypeToolbar           = externglib.Type(C.gtk_toolbar_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeToolbarSpaceStyle, F: marshalToolbarSpaceStyle},
		{T: GTypeToolbar, F: marshalToolbar},
	})
}

// ToolbarSpaceStyle: whether spacers are vertical lines or just blank.
//
// Deprecated: since version 3.20.
type ToolbarSpaceStyle C.gint

const (
	// ToolbarSpaceEmpty: use blank spacers.
	ToolbarSpaceEmpty ToolbarSpaceStyle = iota
	// ToolbarSpaceLine: use vertical lines for spacers.
	ToolbarSpaceLine
)

func marshalToolbarSpaceStyle(p uintptr) (interface{}, error) {
	return ToolbarSpaceStyle(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ToolbarSpaceStyle.
func (t ToolbarSpaceStyle) String() string {
	switch t {
	case ToolbarSpaceEmpty:
		return "Empty"
	case ToolbarSpaceLine:
		return "Line"
	default:
		return fmt.Sprintf("ToolbarSpaceStyle(%d)", t)
	}
}

// ToolbarOverrider contains methods that are overridable.
type ToolbarOverrider interface {
	// The function takes the following parameters:
	//
	OrientationChanged(orientation Orientation)
	// The function takes the following parameters:
	//
	//    - x
	//    - y
	//    - buttonNumber
	//
	// The function returns the following values:
	//
	PopupContextMenu(x, y, buttonNumber int) bool
	// The function takes the following parameters:
	//
	StyleChanged(style ToolbarStyle)
}

// Toolbar: toolbar is created with a call to gtk_toolbar_new().
//
// A toolbar can contain instances of a subclass of ToolItem. To add a ToolItem
// to the a toolbar, use gtk_toolbar_insert(). To remove an item from the
// toolbar use gtk_container_remove(). To add a button to the toolbar, add an
// instance of ToolButton.
//
// Toolbar items can be visually grouped by adding instances of
// SeparatorToolItem to the toolbar. If the GtkToolbar child property “expand”
// is UE and the property SeparatorToolItem:draw is set to LSE, the effect is to
// force all following items to the end of the toolbar.
//
// By default, a toolbar can be shrunk, upon which it will add an arrow button
// to show an overflow menu offering access to any ToolItem child that has a
// proxy menu item. To disable this and request enough size for all children,
// call gtk_toolbar_set_show_arrow() to set Toolbar:show-arrow to FALSE.
//
// Creating a context menu for the toolbar can be done by connecting to the
// Toolbar::popup-context-menu signal.
//
//
// CSS nodes
//
// GtkToolbar has a single CSS node with name toolbar.
type Toolbar struct {
	_ [0]func() // equal guard
	Container

	*externglib.Object
	atk.ImplementorIface
	externglib.InitiallyUnowned
	Buildable
	Orientable
	ToolShell
	Widget
}

var (
	_ Containerer         = (*Toolbar)(nil)
	_ externglib.Objector = (*Toolbar)(nil)
	_ Widgetter           = (*Toolbar)(nil)
)

func classInitToolbarrer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkToolbarClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkToolbarClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ OrientationChanged(orientation Orientation) }); ok {
		pclass.orientation_changed = (*[0]byte)(C._gotk4_gtk3_ToolbarClass_orientation_changed)
	}

	if _, ok := goval.(interface {
		PopupContextMenu(x, y, buttonNumber int) bool
	}); ok {
		pclass.popup_context_menu = (*[0]byte)(C._gotk4_gtk3_ToolbarClass_popup_context_menu)
	}

	if _, ok := goval.(interface{ StyleChanged(style ToolbarStyle) }); ok {
		pclass.style_changed = (*[0]byte)(C._gotk4_gtk3_ToolbarClass_style_changed)
	}
}

//export _gotk4_gtk3_ToolbarClass_orientation_changed
func _gotk4_gtk3_ToolbarClass_orientation_changed(arg0 *C.GtkToolbar, arg1 C.GtkOrientation) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ OrientationChanged(orientation Orientation) })

	var _orientation Orientation // out

	_orientation = Orientation(arg1)

	iface.OrientationChanged(_orientation)
}

//export _gotk4_gtk3_ToolbarClass_popup_context_menu
func _gotk4_gtk3_ToolbarClass_popup_context_menu(arg0 *C.GtkToolbar, arg1 C.gint, arg2 C.gint, arg3 C.gint) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		PopupContextMenu(x, y, buttonNumber int) bool
	})

	var _x int            // out
	var _y int            // out
	var _buttonNumber int // out

	_x = int(arg1)
	_y = int(arg2)
	_buttonNumber = int(arg3)

	ok := iface.PopupContextMenu(_x, _y, _buttonNumber)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_ToolbarClass_style_changed
func _gotk4_gtk3_ToolbarClass_style_changed(arg0 *C.GtkToolbar, arg1 C.GtkToolbarStyle) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ StyleChanged(style ToolbarStyle) })

	var _style ToolbarStyle // out

	_style = ToolbarStyle(arg1)

	iface.StyleChanged(_style)
}

func wrapToolbar(obj *externglib.Object) *Toolbar {
	return &Toolbar{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
		Object: obj,
		ImplementorIface: atk.ImplementorIface{
			Object: obj,
		},
		InitiallyUnowned: externglib.InitiallyUnowned{
			Object: obj,
		},
		Buildable: Buildable{
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		ToolShell: ToolShell{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
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
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
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
	}
}

func marshalToolbar(p uintptr) (interface{}, error) {
	return wrapToolbar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd
func _gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd(arg0 C.gpointer, arg1 C.gboolean, arg2 C.guintptr) (cret C.gboolean) {
	var f func(focusHome bool) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(focusHome bool) (ok bool))
	}

	var _focusHome bool // out

	if arg1 != 0 {
		_focusHome = true
	}

	ok := f(_focusHome)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectFocusHomeOrEnd: keybinding signal used internally by GTK+. This signal
// can't be used in application code.
func (toolbar *Toolbar) ConnectFocusHomeOrEnd(f func(focusHome bool) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(toolbar, "focus-home-or-end", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectFocusHomeOrEnd), f)
}

//export _gotk4_gtk3_Toolbar_ConnectOrientationChanged
func _gotk4_gtk3_Toolbar_ConnectOrientationChanged(arg0 C.gpointer, arg1 C.GtkOrientation, arg2 C.guintptr) {
	var f func(orientation Orientation)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(orientation Orientation))
	}

	var _orientation Orientation // out

	_orientation = Orientation(arg1)

	f(_orientation)
}

// ConnectOrientationChanged: emitted when the orientation of the toolbar
// changes.
func (toolbar *Toolbar) ConnectOrientationChanged(f func(orientation Orientation)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(toolbar, "orientation-changed", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectOrientationChanged), f)
}

//export _gotk4_gtk3_Toolbar_ConnectPopupContextMenu
func _gotk4_gtk3_Toolbar_ConnectPopupContextMenu(arg0 C.gpointer, arg1 C.gint, arg2 C.gint, arg3 C.gint, arg4 C.guintptr) (cret C.gboolean) {
	var f func(x, y, button int) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(x, y, button int) (ok bool))
	}

	var _x int      // out
	var _y int      // out
	var _button int // out

	_x = int(arg1)
	_y = int(arg2)
	_button = int(arg3)

	ok := f(_x, _y, _button)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectPopupContextMenu: emitted when the user right-clicks the toolbar or
// uses the keybinding to display a popup menu.
//
// Application developers should handle this signal if they want to display a
// context menu on the toolbar. The context-menu should appear at the
// coordinates given by x and y. The mouse button number is given by the button
// parameter. If the menu was popped up using the keybaord, button is -1.
func (toolbar *Toolbar) ConnectPopupContextMenu(f func(x, y, button int) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(toolbar, "popup-context-menu", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectPopupContextMenu), f)
}

//export _gotk4_gtk3_Toolbar_ConnectStyleChanged
func _gotk4_gtk3_Toolbar_ConnectStyleChanged(arg0 C.gpointer, arg1 C.GtkToolbarStyle, arg2 C.guintptr) {
	var f func(style ToolbarStyle)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(style ToolbarStyle))
	}

	var _style ToolbarStyle // out

	_style = ToolbarStyle(arg1)

	f(_style)
}

// ConnectStyleChanged: emitted when the style of the toolbar changes.
func (toolbar *Toolbar) ConnectStyleChanged(f func(style ToolbarStyle)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(toolbar, "style-changed", false, unsafe.Pointer(C._gotk4_gtk3_Toolbar_ConnectStyleChanged), f)
}

// NewToolbar creates a new toolbar.
//
// The function returns the following values:
//
//    - toolbar: newly-created toolbar.
//
func NewToolbar() *Toolbar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toolbar_new()

	var _toolbar *Toolbar // out

	_toolbar = wrapToolbar(externglib.Take(unsafe.Pointer(_cret)))

	return _toolbar
}

// DropIndex returns the position corresponding to the indicated point on
// toolbar. This is useful when dragging items to the toolbar: this function
// returns the position a new item should be inserted.
//
// x and y are in toolbar coordinates.
//
// The function takes the following parameters:
//
//    - x coordinate of a point on the toolbar.
//    - y coordinate of a point on the toolbar.
//
// The function returns the following values:
//
//    - gint: position corresponding to the point (x, y) on the toolbar.
//
func (toolbar *Toolbar) DropIndex(x, y int) int {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.gint        // out
	var _arg2 C.gint        // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_toolbar_get_drop_index(_arg0, _arg1, _arg2)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IconSize retrieves the icon size for the toolbar. See
// gtk_toolbar_set_icon_size().
//
// The function returns the following values:
//
//    - iconSize: current icon size for the icons on the toolbar.
//
func (toolbar *Toolbar) IconSize() IconSize {
	var _arg0 *C.GtkToolbar // out
	var _cret C.GtkIconSize // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	_cret = C.gtk_toolbar_get_icon_size(_arg0)
	runtime.KeepAlive(toolbar)

	var _iconSize IconSize // out

	_iconSize = IconSize(_cret)

	return _iconSize
}

// ItemIndex returns the position of item on the toolbar, starting from 0. It is
// an error if item is not a child of the toolbar.
//
// The function takes the following parameters:
//
//    - item that is a child of toolbar.
//
// The function returns the following values:
//
//    - gint: position of item on the toolbar.
//
func (toolbar *Toolbar) ItemIndex(item *ToolItem) int {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(externglib.InternObject(item).Native()))

	_cret = C.gtk_toolbar_get_item_index(_arg0, _arg1)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(item)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NItems returns the number of items on the toolbar.
//
// The function returns the following values:
//
//    - gint: number of items on the toolbar.
//
func (toolbar *Toolbar) NItems() int {
	var _arg0 *C.GtkToolbar // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	_cret = C.gtk_toolbar_get_n_items(_arg0)
	runtime.KeepAlive(toolbar)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthItem returns the n'th item on toolbar, or NULL if the toolbar does not
// contain an n'th item.
//
// The function takes the following parameters:
//
//    - n on the toolbar.
//
// The function returns the following values:
//
//    - toolItem (optional): n'th ToolItem on toolbar, or NULL if there isn’t an
//      n'th item.
//
func (toolbar *Toolbar) NthItem(n int) *ToolItem {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 C.gint         // out
	var _cret *C.GtkToolItem // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = C.gint(n)

	_cret = C.gtk_toolbar_get_nth_item(_arg0, _arg1)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(n)

	var _toolItem *ToolItem // out

	if _cret != nil {
		_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _toolItem
}

// ReliefStyle returns the relief style of buttons on toolbar. See
// gtk_button_set_relief().
//
// The function returns the following values:
//
//    - reliefStyle: relief style of buttons on toolbar.
//
func (toolbar *Toolbar) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolbar    // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	_cret = C.gtk_toolbar_get_relief_style(_arg0)
	runtime.KeepAlive(toolbar)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// ShowArrow returns whether the toolbar has an overflow menu. See
// gtk_toolbar_set_show_arrow().
//
// The function returns the following values:
//
//    - ok: TRUE if the toolbar has an overflow menu.
//
func (toolbar *Toolbar) ShowArrow() bool {
	var _arg0 *C.GtkToolbar // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	_cret = C.gtk_toolbar_get_show_arrow(_arg0)
	runtime.KeepAlive(toolbar)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Style retrieves whether the toolbar has text, icons, or both . See
// gtk_toolbar_set_style().
//
// The function returns the following values:
//
//    - toolbarStyle: current style of toolbar.
//
func (toolbar *Toolbar) Style() ToolbarStyle {
	var _arg0 *C.GtkToolbar     // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	_cret = C.gtk_toolbar_get_style(_arg0)
	runtime.KeepAlive(toolbar)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// Insert a ToolItem into the toolbar at position pos. If pos is 0 the item is
// prepended to the start of the toolbar. If pos is negative, the item is
// appended to the end of the toolbar.
//
// The function takes the following parameters:
//
//    - item: ToolItem.
//    - pos: position of the new item.
//
func (toolbar *Toolbar) Insert(item *ToolItem, pos int) {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(externglib.InternObject(item).Native()))
	_arg2 = C.gint(pos)

	C.gtk_toolbar_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(item)
	runtime.KeepAlive(pos)
}

// SetDropHighlightItem highlights toolbar to give an idea of what it would look
// like if item was added to toolbar at the position indicated by index_. If
// item is NULL, highlighting is turned off. In that case index_ is ignored.
//
// The tool_item passed to this function must not be part of any widget
// hierarchy. When an item is set as drop highlight item it can not added to any
// widget hierarchy or used as highlight item for another toolbar.
//
// The function takes the following parameters:
//
//    - toolItem (optional) or NULL to turn of highlighting.
//    - index_: position on toolbar.
//
func (toolbar *Toolbar) SetDropHighlightItem(toolItem *ToolItem, index_ int) {
	var _arg0 *C.GtkToolbar  // out
	var _arg1 *C.GtkToolItem // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	if toolItem != nil {
		_arg1 = (*C.GtkToolItem)(unsafe.Pointer(externglib.InternObject(toolItem).Native()))
	}
	_arg2 = C.gint(index_)

	C.gtk_toolbar_set_drop_highlight_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(index_)
}

// SetIconSize: this function sets the size of stock icons in the toolbar. You
// can call it both before you add the icons and after they’ve been added. The
// size you set will override user preferences for the default icon size.
//
// This should only be used for special-purpose toolbars, normal application
// toolbars should respect the user preferences for the size of icons.
//
// The function takes the following parameters:
//
//    - iconSize that stock icons in the toolbar shall have.
//
func (toolbar *Toolbar) SetIconSize(iconSize IconSize) {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.GtkIconSize // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = C.GtkIconSize(iconSize)

	C.gtk_toolbar_set_icon_size(_arg0, _arg1)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(iconSize)
}

// SetShowArrow sets whether to show an overflow menu when toolbar isn’t
// allocated enough size to show all of its items. If TRUE, items which can’t
// fit in toolbar, and which have a proxy menu item set by
// gtk_tool_item_set_proxy_menu_item() or ToolItem::create-menu-proxy, will be
// available in an overflow menu, which can be opened by an added arrow button.
// If FALSE, toolbar will request enough size to fit all of its child items
// without any overflow.
//
// The function takes the following parameters:
//
//    - showArrow: whether to show an overflow menu.
//
func (toolbar *Toolbar) SetShowArrow(showArrow bool) {
	var _arg0 *C.GtkToolbar // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	if showArrow {
		_arg1 = C.TRUE
	}

	C.gtk_toolbar_set_show_arrow(_arg0, _arg1)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(showArrow)
}

// SetStyle alters the view of toolbar to display either icons only, text only,
// or both.
//
// The function takes the following parameters:
//
//    - style: new style for toolbar.
//
func (toolbar *Toolbar) SetStyle(style ToolbarStyle) {
	var _arg0 *C.GtkToolbar     // out
	var _arg1 C.GtkToolbarStyle // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))
	_arg1 = C.GtkToolbarStyle(style)

	C.gtk_toolbar_set_style(_arg0, _arg1)
	runtime.KeepAlive(toolbar)
	runtime.KeepAlive(style)
}

// UnsetIconSize unsets toolbar icon size set with gtk_toolbar_set_icon_size(),
// so that user preferences will be used to determine the icon size.
func (toolbar *Toolbar) UnsetIconSize() {
	var _arg0 *C.GtkToolbar // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	C.gtk_toolbar_unset_icon_size(_arg0)
	runtime.KeepAlive(toolbar)
}

// UnsetStyle unsets a toolbar style set with gtk_toolbar_set_style(), so that
// user preferences will be used to determine the toolbar style.
func (toolbar *Toolbar) UnsetStyle() {
	var _arg0 *C.GtkToolbar // out

	_arg0 = (*C.GtkToolbar)(unsafe.Pointer(externglib.InternObject(toolbar).Native()))

	C.gtk_toolbar_unset_style(_arg0)
	runtime.KeepAlive(toolbar)
}
