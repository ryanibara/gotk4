// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_item_get_type()), F: marshalToolItemmer},
	})
}

// ToolItemOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ToolItemOverrider interface {
	// The function returns the following values:
	//
	CreateMenuProxy() bool
	// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
	// tool_item. Toolbar and other ToolShell implementations use this function
	// to notify children, when some aspect of their configuration changes.
	ToolbarReconfigured()
}

// ToolItem are widgets that can appear on a toolbar. To create a toolbar item
// that contain something else than a button, use gtk_tool_item_new(). Use
// gtk_container_add() to add a child widget to the tool item.
//
// For toolbar items that contain buttons, see the ToolButton, ToggleToolButton
// and RadioToolButton classes.
//
// See the Toolbar class for a description of the toolbar widget, and ToolShell
// for a description of the tool shell interface.
type ToolItem struct {
	Bin

	*externglib.Object
	Activatable

	_ [0]func()     // equal guard
	_ [0]sync.Mutex // copy guard
}

var (
	_ Binner              = (*ToolItem)(nil)
	_ externglib.Objector = (*ToolItem)(nil)
)

func wrapToolItem(obj *externglib.Object) *ToolItem {
	return &ToolItem{
		Bin: Bin{
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
		},
		Object: obj,
		Activatable: Activatable{
			Object: obj,
		},
	}
}

func marshalToolItemmer(p uintptr) (interface{}, error) {
	return wrapToolItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectCreateMenuProxy: this signal is emitted when the toolbar needs
// information from tool_item about whether the item should appear in the
// toolbar overflow menu. In response the tool item should either
//
// - call gtk_tool_item_set_proxy_menu_item() with a NULL pointer and return
// TRUE to indicate that the item should not appear in the overflow menu
//
// - call gtk_tool_item_set_proxy_menu_item() with a new menu item and return
// TRUE, or
//
// - return FALSE to indicate that the signal was not handled by the item. This
// means that the item will not appear in the overflow menu unless a later
// handler installs a menu item.
//
// The toolbar may cache the result of this signal. When the tool item changes
// how it will respond to this signal it must call gtk_tool_item_rebuild_menu()
// to invalidate the cache and ensure that the toolbar rebuilds its overflow
// menu.
func (toolItem *ToolItem) ConnectCreateMenuProxy(f func() bool) externglib.SignalHandle {
	return toolItem.Connect("create-menu-proxy", f)
}

// ConnectToolbarReconfigured: this signal is emitted when some property of the
// toolbar that the item is a child of changes. For custom subclasses of
// ToolItem, the default handler of this signal use the functions
//
// - gtk_tool_shell_get_orientation()
//
// - gtk_tool_shell_get_style()
//
// - gtk_tool_shell_get_icon_size()
//
// - gtk_tool_shell_get_relief_style() to find out what the toolbar should look
// like and change themselves accordingly.
func (toolItem *ToolItem) ConnectToolbarReconfigured(f func()) externglib.SignalHandle {
	return toolItem.Connect("toolbar-reconfigured", f)
}

// NewToolItem creates a new ToolItem.
//
// The function returns the following values:
//
//    - toolItem: new ToolItem.
//
func NewToolItem() *ToolItem {
	var _cret *C.GtkToolItem // in

	_cret = C.gtk_tool_item_new()

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// EllipsizeMode returns the ellipsize mode used for tool_item. Custom
// subclasses of ToolItem should call this function to find out how text should
// be ellipsized.
//
// The function returns the following values:
//
//    - ellipsizeMode indicating how text in tool_item should be ellipsized.
//
func (toolItem *ToolItem) EllipsizeMode() pango.EllipsizeMode {
	var _arg0 *C.GtkToolItem       // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_ellipsize_mode(_arg0)
	runtime.KeepAlive(toolItem)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// Expand returns whether tool_item is allocated extra space. See
// gtk_tool_item_set_expand().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is allocated extra space.
//
func (toolItem *ToolItem) Expand() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_expand(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous returns whether tool_item is the same size as other homogeneous
// items. See gtk_tool_item_set_homogeneous().
//
// The function returns the following values:
//
//    - ok: TRUE if the item is the same size as other homogeneous items.
//
func (toolItem *ToolItem) Homogeneous() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_homogeneous(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IconSize returns the icon size used for tool_item. Custom subclasses of
// ToolItem should call this function to find out what size icons they should
// use.
//
// The function returns the following values:
//
//    - gint indicating the icon size used for tool_item.
//
func (toolItem *ToolItem) IconSize() int {
	var _arg0 *C.GtkToolItem // out
	var _cret C.GtkIconSize  // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_icon_size(_arg0)
	runtime.KeepAlive(toolItem)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsImportant returns whether tool_item is considered important. See
// gtk_tool_item_set_is_important().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is considered important.
//
func (toolItem *ToolItem) IsImportant() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_is_important(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Orientation returns the orientation used for tool_item. Custom subclasses of
// ToolItem should call this function to find out what size icons they should
// use.
//
// The function returns the following values:
//
//    - orientation indicating the orientation used for tool_item.
//
func (toolItem *ToolItem) Orientation() Orientation {
	var _arg0 *C.GtkToolItem   // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_orientation(_arg0)
	runtime.KeepAlive(toolItem)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// ProxyMenuItem: if menu_item_id matches the string passed to
// gtk_tool_item_set_proxy_menu_item() return the corresponding MenuItem.
//
// Custom subclasses of ToolItem should use this function to update their menu
// item when the ToolItem changes. That the menu_item_ids must match ensures
// that a ToolItem will not inadvertently change a menu item that they did not
// create.
//
// The function takes the following parameters:
//
//    - menuItemId: string used to identify the menu item.
//
// The function returns the following values:
//
//    - widget (optional) passed to gtk_tool_item_set_proxy_menu_item(), if the
//      menu_item_ids match.
//
func (toolItem *ToolItem) ProxyMenuItem(menuItemId string) Widgetter {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(menuItemId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_tool_item_get_proxy_menu_item(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(menuItemId)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.Cast()
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// ReliefStyle returns the relief style of tool_item. See
// gtk_button_set_relief(). Custom subclasses of ToolItem should call this
// function in the handler of the ToolItem::toolbar_reconfigured signal to find
// out the relief style of buttons.
//
// The function returns the following values:
//
//    - reliefStyle indicating the relief style used for tool_item.
//
func (toolItem *ToolItem) ReliefStyle() ReliefStyle {
	var _arg0 *C.GtkToolItem   // out
	var _cret C.GtkReliefStyle // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_relief_style(_arg0)
	runtime.KeepAlive(toolItem)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// TextAlignment returns the text alignment used for tool_item. Custom
// subclasses of ToolItem should call this function to find out how text should
// be aligned.
//
// The function returns the following values:
//
//    - gfloat indicating the horizontal text alignment used for tool_item.
//
func (toolItem *ToolItem) TextAlignment() float32 {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gfloat       // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_text_alignment(_arg0)
	runtime.KeepAlive(toolItem)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// TextOrientation returns the text orientation used for tool_item. Custom
// subclasses of ToolItem should call this function to find out how text should
// be orientated.
//
// The function returns the following values:
//
//    - orientation indicating the text orientation used for tool_item.
//
func (toolItem *ToolItem) TextOrientation() Orientation {
	var _arg0 *C.GtkToolItem   // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_text_orientation(_arg0)
	runtime.KeepAlive(toolItem)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// TextSizeGroup returns the size group used for labels in tool_item. Custom
// subclasses of ToolItem should call this function and use the size group for
// labels.
//
// The function returns the following values:
//
//    - sizeGroup: SizeGroup.
//
func (toolItem *ToolItem) TextSizeGroup() *SizeGroup {
	var _arg0 *C.GtkToolItem  // out
	var _cret *C.GtkSizeGroup // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_text_size_group(_arg0)
	runtime.KeepAlive(toolItem)

	var _sizeGroup *SizeGroup // out

	_sizeGroup = wrapSizeGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _sizeGroup
}

// ToolbarStyle returns the toolbar style used for tool_item. Custom subclasses
// of ToolItem should call this function in the handler of the
// GtkToolItem::toolbar_reconfigured signal to find out in what style the
// toolbar is displayed and change themselves accordingly
//
// Possibilities are:
//
// - GTK_TOOLBAR_BOTH, meaning the tool item should show both an icon and a
// label, stacked vertically
//
// - GTK_TOOLBAR_ICONS, meaning the toolbar shows only icons
//
// - GTK_TOOLBAR_TEXT, meaning the tool item should only show text
//
// - GTK_TOOLBAR_BOTH_HORIZ, meaning the tool item should show both an icon and
// a label, arranged horizontally.
//
// The function returns the following values:
//
//    - toolbarStyle indicating the toolbar style used for tool_item.
//
func (toolItem *ToolItem) ToolbarStyle() ToolbarStyle {
	var _arg0 *C.GtkToolItem    // out
	var _cret C.GtkToolbarStyle // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_toolbar_style(_arg0)
	runtime.KeepAlive(toolItem)

	var _toolbarStyle ToolbarStyle // out

	_toolbarStyle = ToolbarStyle(_cret)

	return _toolbarStyle
}

// UseDragWindow returns whether tool_item has a drag window. See
// gtk_tool_item_set_use_drag_window().
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item uses a drag window.
//
func (toolItem *ToolItem) UseDragWindow() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_use_drag_window(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleHorizontal returns whether the tool_item is visible on toolbars that
// are docked horizontally.
//
// The function returns the following values:
//
//    - ok: TRUE if tool_item is visible on toolbars that are docked
//      horizontally.
//
func (toolItem *ToolItem) VisibleHorizontal() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_visible_horizontal(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleVertical returns whether tool_item is visible when the toolbar is
// docked vertically. See gtk_tool_item_set_visible_vertical().
//
// The function returns the following values:
//
//    - ok: whether tool_item is visible when the toolbar is docked vertically.
//
func (toolItem *ToolItem) VisibleVertical() bool {
	var _arg0 *C.GtkToolItem // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_get_visible_vertical(_arg0)
	runtime.KeepAlive(toolItem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RebuildMenu: calling this function signals to the toolbar that the overflow
// menu item for tool_item has changed. If the overflow menu is visible when
// this function it called, the menu will be rebuilt.
//
// The function must be called when the tool item changes what it will do in
// response to the ToolItem::create-menu-proxy signal.
func (toolItem *ToolItem) RebuildMenu() {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	C.gtk_tool_item_rebuild_menu(_arg0)
	runtime.KeepAlive(toolItem)
}

// RetrieveProxyMenuItem returns the MenuItem that was last set by
// gtk_tool_item_set_proxy_menu_item(), ie. the MenuItem that is going to appear
// in the overflow menu.
//
// The function returns the following values:
//
//    - widget that is going to appear in the overflow menu for tool_item.
//
func (toolItem *ToolItem) RetrieveProxyMenuItem() Widgetter {
	var _arg0 *C.GtkToolItem // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	_cret = C.gtk_tool_item_retrieve_proxy_menu_item(_arg0)
	runtime.KeepAlive(toolItem)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.Cast()
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("object of type " + reflect.TypeOf(casted).String() + " (" + object.TypeFromInstance().String() + ") is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// SetExpand sets whether tool_item is allocated extra space when there is more
// room on the toolbar then needed for the items. The effect is that the item
// gets bigger when the toolbar gets bigger and smaller when the toolbar gets
// smaller.
//
// The function takes the following parameters:
//
//    - expand: whether tool_item is allocated extra space.
//
func (toolItem *ToolItem) SetExpand(expand bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_expand(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(expand)
}

// SetHomogeneous sets whether tool_item is to be allocated the same size as
// other homogeneous items. The effect is that all homogeneous items will have
// the same width as the widest of the items.
//
// The function takes the following parameters:
//
//    - homogeneous: whether tool_item is the same size as other homogeneous
//      items.
//
func (toolItem *ToolItem) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_homogeneous(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(homogeneous)
}

// SetIsImportant sets whether tool_item should be considered important. The
// ToolButton class uses this property to determine whether to show or hide its
// label when the toolbar style is GTK_TOOLBAR_BOTH_HORIZ. The result is that
// only tool buttons with the “is_important” property set have labels, an effect
// known as “priority text”.
//
// The function takes the following parameters:
//
//    - isImportant: whether the tool item should be considered important.
//
func (toolItem *ToolItem) SetIsImportant(isImportant bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_is_important(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(isImportant)
}

// SetProxyMenuItem sets the MenuItem used in the toolbar overflow menu. The
// menu_item_id is used to identify the caller of this function and should also
// be used with gtk_tool_item_get_proxy_menu_item().
//
// See also ToolItem::create-menu-proxy.
//
// The function takes the following parameters:
//
//    - menuItemId: string used to identify menu_item.
//    - menuItem (optional) to use in the overflow menu, or NULL.
//
func (toolItem *ToolItem) SetProxyMenuItem(menuItemId string, menuItem Widgetter) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(menuItemId)))
	defer C.free(unsafe.Pointer(_arg1))
	if menuItem != nil {
		_arg2 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))
	}

	C.gtk_tool_item_set_proxy_menu_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(menuItemId)
	runtime.KeepAlive(menuItem)
}

// SetTooltipMarkup sets the markup text to be displayed as tooltip on the item.
// See gtk_widget_set_tooltip_markup().
//
// The function takes the following parameters:
//
//    - markup text to be used as tooltip for tool_item.
//
func (toolItem *ToolItem) SetTooltipMarkup(markup string) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_set_tooltip_markup(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(markup)
}

// SetTooltipText sets the text to be displayed as tooltip on the item. See
// gtk_widget_set_tooltip_text().
//
// The function takes the following parameters:
//
//    - text to be used as tooltip for tool_item.
//
func (toolItem *ToolItem) SetTooltipText(text string) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_set_tooltip_text(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(text)
}

// SetUseDragWindow sets whether tool_item has a drag window. When TRUE the
// toolitem can be used as a drag source through gtk_drag_source_set(). When
// tool_item has a drag window it will intercept all events, even those that
// would otherwise be sent to a child of tool_item.
//
// The function takes the following parameters:
//
//    - useDragWindow: whether tool_item has a drag window.
//
func (toolItem *ToolItem) SetUseDragWindow(useDragWindow bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if useDragWindow {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_use_drag_window(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(useDragWindow)
}

// SetVisibleHorizontal sets whether tool_item is visible when the toolbar is
// docked horizontally.
//
// The function takes the following parameters:
//
//    - visibleHorizontal: whether tool_item is visible when in horizontal mode.
//
func (toolItem *ToolItem) SetVisibleHorizontal(visibleHorizontal bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_visible_horizontal(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(visibleHorizontal)
}

// SetVisibleVertical sets whether tool_item is visible when the toolbar is
// docked vertically. Some tool items, such as text entries, are too wide to be
// useful on a vertically docked toolbar. If visible_vertical is FALSE tool_item
// will not appear on toolbars that are docked vertically.
//
// The function takes the following parameters:
//
//    - visibleVertical: whether tool_item is visible when the toolbar is in
//      vertical mode.
//
func (toolItem *ToolItem) SetVisibleVertical(visibleVertical bool) {
	var _arg0 *C.GtkToolItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_set_visible_vertical(_arg0, _arg1)
	runtime.KeepAlive(toolItem)
	runtime.KeepAlive(visibleVertical)
}

// ToolbarReconfigured emits the signal ToolItem::toolbar_reconfigured on
// tool_item. Toolbar and other ToolShell implementations use this function to
// notify children, when some aspect of their configuration changes.
func (toolItem *ToolItem) ToolbarReconfigured() {
	var _arg0 *C.GtkToolItem // out

	_arg0 = (*C.GtkToolItem)(unsafe.Pointer(toolItem.Native()))

	C.gtk_tool_item_toolbar_reconfigured(_arg0)
	runtime.KeepAlive(toolItem)
}
