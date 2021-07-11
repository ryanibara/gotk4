// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_action_get_type()), F: marshalActioner},
	})
}

// ActionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ActionOverrider interface {
	// Activate emits the “activate” signal on the specified action, if it isn't
	// insensitive. This gets called by the proxy widgets when they get
	// activated.
	//
	// It can also be used to manually activate an action.
	//
	// Deprecated: Use g_action_group_activate_action() on a #GAction instead.
	Activate()

	ConnectProxy(proxy Widgetter)
	// CreateMenu: if @action provides a Menu widget as a submenu for the menu
	// item or the toolbar item it creates, this function returns an instance of
	// that menu.
	//
	// Deprecated: Use #GAction and Model instead, and create a Menu with
	// gtk_menu_new_from_model().
	CreateMenu() *Widget
	// CreateMenuItem creates a menu item widget that proxies for the given
	// action.
	//
	// Deprecated: Use g_menu_item_new() and associate it with a #GAction
	// instead.
	CreateMenuItem() *Widget
	// CreateToolItem creates a toolbar item widget that proxies for the given
	// action.
	//
	// Deprecated: Use a ToolItem and associate it with a #GAction using
	// gtk_actionable_set_action_name() instead.
	CreateToolItem() *Widget

	DisconnectProxy(proxy Widgetter)
}

// Actioner describes Action's methods.
type Actioner interface {
	// Activate emits the “activate” signal on the specified action, if it isn't
	// insensitive.
	Activate()
	// BlockActivate: disable activation signals from the action This is needed
	// when updating the state of your proxy Activatable widget could result in
	// calling gtk_action_activate(), this is a convenience function to avoid
	// recursing in those cases (updating toggle state for instance).
	BlockActivate()
	// ConnectAccelerator installs the accelerator for @action if @action has an
	// accel path and group.
	ConnectAccelerator()
	// CreateIcon: this function is intended for use by action implementations
	// to create icons displayed in the proxy widgets.
	CreateIcon(iconSize int) *Widget
	// CreateMenu: if @action provides a Menu widget as a submenu for the menu
	// item or the toolbar item it creates, this function returns an instance of
	// that menu.
	CreateMenu() *Widget
	// CreateMenuItem creates a menu item widget that proxies for the given
	// action.
	CreateMenuItem() *Widget
	// CreateToolItem creates a toolbar item widget that proxies for the given
	// action.
	CreateToolItem() *Widget
	// DisconnectAccelerator undoes the effect of one call to
	// gtk_action_connect_accelerator().
	DisconnectAccelerator()
	// AccelPath returns the accel path for this action.
	AccelPath() string
	// AlwaysShowImage returns whether @action's menu item proxies will always
	// show their image, if available.
	AlwaysShowImage() bool
	// GIcon gets the gicon of @action.
	GIcon() *gio.Icon
	// IconName gets the icon name of @action.
	IconName() string
	// IsImportant checks whether @action is important or not Deprecated: Use
	// #GAction instead, and control and monitor whether labels are shown
	// directly.
	IsImportant() bool
	// Label gets the label text of @action.
	Label() string
	// Name returns the name of the action.
	Name() string
	// Sensitive returns whether the action itself is sensitive.
	Sensitive() bool
	// ShortLabel gets the short label text of @action.
	ShortLabel() string
	// StockID gets the stock id of @action.
	StockID() string
	// Tooltip gets the tooltip text of @action.
	Tooltip() string
	// Visible returns whether the action itself is visible.
	Visible() bool
	// VisibleHorizontal checks whether @action is visible when horizontal
	// Deprecated: Use #GAction instead, and control and monitor the visibility
	// of associated widgets and menu items directly.
	VisibleHorizontal() bool
	// VisibleVertical checks whether @action is visible when horizontal
	// Deprecated: Use #GAction instead, and control and monitor the visibility
	// of associated widgets and menu items directly.
	VisibleVertical() bool
	// IsSensitive returns whether the action is effectively sensitive.
	IsSensitive() bool
	// IsVisible returns whether the action is effectively visible.
	IsVisible() bool
	// SetAccelGroup sets the AccelGroup in which the accelerator for this
	// action will be installed.
	SetAccelGroup(accelGroup AccelGrouper)
	// SetAccelPath sets the accel path for this action.
	SetAccelPath(accelPath string)
	// SetAlwaysShowImage sets whether @action's menu item proxies will ignore
	// the Settings:gtk-menu-images setting and always show their image, if
	// available.
	SetAlwaysShowImage(alwaysShow bool)
	// SetGIcon sets the icon of @action.
	SetGIcon(icon gio.Iconner)
	// SetIconName sets the icon name on @action Deprecated: Use #GAction
	// instead, and g_menu_item_set_icon() to set an icon on a Item associated
	// with a #GAction, or gtk_container_add() to add a Image to a Button.
	SetIconName(iconName string)
	// SetIsImportant sets whether the action is important, this attribute is
	// used primarily by toolbar items to decide whether to show a label or not.
	SetIsImportant(isImportant bool)
	// SetLabel sets the label of @action.
	SetLabel(label string)
	// SetSensitive sets the :sensitive property of the action to @sensitive.
	SetSensitive(sensitive bool)
	// SetShortLabel sets a shorter label text on @action.
	SetShortLabel(shortLabel string)
	// SetStockID sets the stock id on @action Deprecated: Use #GAction instead,
	// which has no equivalent of stock items.
	SetStockID(stockId string)
	// SetTooltip sets the tooltip text on @action Deprecated: Use #GAction
	// instead, and set tooltips on associated Actionable widgets with
	// gtk_widget_set_tooltip_text().
	SetTooltip(tooltip string)
	// SetVisible sets the :visible property of the action to @visible.
	SetVisible(visible bool)
	// SetVisibleHorizontal sets whether @action is visible when horizontal
	// Deprecated: Use #GAction instead, and control and monitor the visibility
	// of associated widgets and menu items directly.
	SetVisibleHorizontal(visibleHorizontal bool)
	// SetVisibleVertical sets whether @action is visible when vertical
	// Deprecated: Use #GAction instead, and control and monitor the visibility
	// of associated widgets and menu items directly.
	SetVisibleVertical(visibleVertical bool)
	// UnblockActivate: reenable activation signals from the action Deprecated:
	// Use g_simple_action_set_enabled() to enable the Action instead.
	UnblockActivate()
}

// Action: > In GTK+ 3.10, GtkAction has been deprecated. Use #GAction >
// instead, and associate actions with Actionable widgets. Use > Model for
// creating menus with gtk_menu_new_from_model().
//
// Actions represent operations that the user can be perform, along with some
// information how it should be presented in the interface. Each action provides
// methods to create icons, menu items and toolbar items representing itself.
//
// As well as the callback that is called when the action gets activated, the
// following also gets associated with the action:
//
// - a name (not translated, for path lookup)
//
// - a label (translated, for display)
//
// - an accelerator
//
// - whether label indicates a stock id
//
// - a tooltip (optional, translated)
//
// - a toolbar label (optional, shorter than label)
//
//    The action will also have some state information:
//
// - visible (shown/hidden)
//
// - sensitive (enabled/disabled)
//
// Apart from regular actions, there are [toggle actions][GtkToggleAction],
// which can be toggled between two states and [radio actions][GtkRadioAction],
// of which only one in a group can be in the “active” state. Other actions can
// be implemented as Action subclasses.
//
// Each action can have one or more proxy widgets. To act as an action proxy,
// widget needs to implement Activatable interface. Proxies mirror the state of
// the action and should change when the action’s state changes. Properties that
// are always mirrored by proxies are Action:sensitive and Action:visible.
// Action:gicon, Action:icon-name, Action:label, Action:short-label and
// Action:stock-id properties are only mirorred if proxy widget has
// Activatable:use-action-appearance property set to true.
//
// When the proxy is activated, it should activate its action.
type Action struct {
	*externglib.Object

	Buildable
}

var (
	_ Actioner        = (*Action)(nil)
	_ gextras.Nativer = (*Action)(nil)
)

func wrapAction(obj *externglib.Object) Actioner {
	return &Action{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalActioner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAction(obj), nil
}

// NewAction creates a new Action object. To add the action to a ActionGroup and
// set the accelerator for the action, call
// gtk_action_group_add_action_with_accel(). See the [UI Definition
// section][XML-UI] for information on allowed action names.
//
// Deprecated: Use #GAction instead, associating it to a widget with Actionable
// or creating a Menu with gtk_menu_new_from_model().
func NewAction(name string, label string, tooltip string, stockId string) *Action {
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _arg4 *C.gchar     // out
	var _cret *C.GtkAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))

	_cret = C.gtk_action_new(_arg1, _arg2, _arg3, _arg4)

	var _action *Action // out

	_action = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Action)

	return _action
}

// Activate emits the “activate” signal on the specified action, if it isn't
// insensitive. This gets called by the proxy widgets when they get activated.
//
// It can also be used to manually activate an action.
//
// Deprecated: Use g_action_group_activate_action() on a #GAction instead.
func (action *Action) Activate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_activate(_arg0)
}

// BlockActivate: disable activation signals from the action
//
// This is needed when updating the state of your proxy Activatable widget could
// result in calling gtk_action_activate(), this is a convenience function to
// avoid recursing in those cases (updating toggle state for instance).
//
// Deprecated: Use g_simple_action_set_enabled() to disable the Action instead.
func (action *Action) BlockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_block_activate(_arg0)
}

// ConnectAccelerator installs the accelerator for @action if @action has an
// accel path and group. See gtk_action_set_accel_path() and
// gtk_action_set_accel_group()
//
// Since multiple proxies may independently trigger the installation of the
// accelerator, the @action counts the number of times this function has been
// called and doesn’t remove the accelerator until
// gtk_action_disconnect_accelerator() has been called as many times.
//
// Deprecated: Use #GAction and the accelerator group on an associated Menu
// instead.
func (action *Action) ConnectAccelerator() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_connect_accelerator(_arg0)
}

// CreateIcon: this function is intended for use by action implementations to
// create icons displayed in the proxy widgets.
//
// Deprecated: Use g_menu_item_set_icon() to set an icon on a Item, or
// gtk_container_add() to add a Image to a Button.
func (action *Action) CreateIcon(iconSize int) *Widget {
	var _arg0 *C.GtkAction  // out
	var _arg1 C.GtkIconSize // out
	var _cret *C.GtkWidget  // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = C.GtkIconSize(iconSize)

	_cret = C.gtk_action_create_icon(_arg0, _arg1)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// CreateMenu: if @action provides a Menu widget as a submenu for the menu item
// or the toolbar item it creates, this function returns an instance of that
// menu.
//
// Deprecated: Use #GAction and Model instead, and create a Menu with
// gtk_menu_new_from_model().
func (action *Action) CreateMenu() *Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_create_menu(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// CreateMenuItem creates a menu item widget that proxies for the given action.
//
// Deprecated: Use g_menu_item_new() and associate it with a #GAction instead.
func (action *Action) CreateMenuItem() *Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_create_menu_item(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// CreateToolItem creates a toolbar item widget that proxies for the given
// action.
//
// Deprecated: Use a ToolItem and associate it with a #GAction using
// gtk_actionable_set_action_name() instead.
func (action *Action) CreateToolItem() *Widget {
	var _arg0 *C.GtkAction // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_create_tool_item(_arg0)

	var _widget *Widget // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Widget)

	return _widget
}

// DisconnectAccelerator undoes the effect of one call to
// gtk_action_connect_accelerator().
//
// Deprecated: Use #GAction and the accelerator group on an associated Menu
// instead.
func (action *Action) DisconnectAccelerator() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_disconnect_accelerator(_arg0)
}

// AccelPath returns the accel path for this action.
//
// Deprecated: Use #GAction and the accelerator path on an associated Menu
// instead.
func (action *Action) AccelPath() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_accel_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// AlwaysShowImage returns whether @action's menu item proxies will always show
// their image, if available.
//
// Deprecated: Use g_menu_item_get_attribute_value() on a Item instead.
func (action *Action) AlwaysShowImage() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_always_show_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GIcon gets the gicon of @action.
//
// Deprecated: Use #GAction instead, and g_menu_item_get_attribute_value() to
// get an icon from a Item associated with a #GAction.
func (action *Action) GIcon() *gio.Icon {
	var _arg0 *C.GtkAction // out
	var _cret *C.GIcon     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_gicon(_arg0)

	var _icon *gio.Icon // out

	_icon = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.Icon)

	return _icon
}

// IconName gets the icon name of @action.
//
// Deprecated: Use #GAction instead, and g_menu_item_get_attribute_value() to
// get an icon from a Item associated with a #GAction.
func (action *Action) IconName() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_icon_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsImportant checks whether @action is important or not
//
// Deprecated: Use #GAction instead, and control and monitor whether labels are
// shown directly.
func (action *Action) IsImportant() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_is_important(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label gets the label text of @action.
//
// Deprecated: Use #GAction instead, and get a label from a menu item with
// g_menu_item_get_attribute_value(). For Actionable widgets, use the
// widget-specific API to get a label.
func (action *Action) Label() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Name returns the name of the action.
//
// Deprecated: Use g_action_get_name() on a #GAction instead.
func (action *Action) Name() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Sensitive returns whether the action itself is sensitive. Note that this
// doesn’t necessarily mean effective sensitivity. See gtk_action_is_sensitive()
// for that.
//
// Deprecated: Use g_action_get_enabled() on a #GAction instead.
func (action *Action) Sensitive() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShortLabel gets the short label text of @action.
//
// Deprecated: Use #GAction instead, which has no equivalent of short labels.
func (action *Action) ShortLabel() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_short_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// StockID gets the stock id of @action.
//
// Deprecated: Use #GAction instead, which has no equivalent of stock items.
func (action *Action) StockID() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_stock_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Tooltip gets the tooltip text of @action.
//
// Deprecated: Use #GAction instead, and get tooltips from associated Actionable
// widgets with gtk_widget_get_tooltip_text().
func (action *Action) Tooltip() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_tooltip(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Visible returns whether the action itself is visible. Note that this doesn’t
// necessarily mean effective visibility. See gtk_action_is_sensitive() for
// that.
//
// Deprecated: Use #GAction instead, and control and monitor the state of
// Actionable widgets directly.
func (action *Action) Visible() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleHorizontal checks whether @action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
func (action *Action) VisibleHorizontal() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_visible_horizontal(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleVertical checks whether @action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
func (action *Action) VisibleVertical() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_get_visible_vertical(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSensitive returns whether the action is effectively sensitive.
//
// Deprecated: Use g_action_get_enabled() on a #GAction instead.
func (action *Action) IsSensitive() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_is_sensitive(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsVisible returns whether the action is effectively visible.
//
// Deprecated: Use #GAction instead, and control and monitor the state of
// Actionable widgets directly.
func (action *Action) IsVisible() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	_cret = C.gtk_action_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAccelGroup sets the AccelGroup in which the accelerator for this action
// will be installed.
//
// Deprecated: Use #GAction and the accelerator group on an associated Menu
// instead.
func (action *Action) SetAccelGroup(accelGroup AccelGrouper) {
	var _arg0 *C.GtkAction     // out
	var _arg1 *C.GtkAccelGroup // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer((accelGroup).(gextras.Nativer).Native()))

	C.gtk_action_set_accel_group(_arg0, _arg1)
}

// SetAccelPath sets the accel path for this action. All proxy widgets
// associated with the action will have this accel path, so that their
// accelerators are consistent.
//
// Note that @accel_path string will be stored in a #GQuark. Therefore, if you
// pass a static string, you can save some memory by interning it first with
// g_intern_static_string().
//
// Deprecated: Use #GAction and the accelerator path on an associated Menu
// instead.
func (action *Action) SetAccelPath(accelPath string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(accelPath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_accel_path(_arg0, _arg1)
}

// SetAlwaysShowImage sets whether @action's menu item proxies will ignore the
// Settings:gtk-menu-images setting and always show their image, if available.
//
// Use this if the menu item would be useless or hard to use without their
// image.
//
// Deprecated: Use g_menu_item_set_icon() on a Item instead, if the item should
// have an image.
func (action *Action) SetAlwaysShowImage(alwaysShow bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_always_show_image(_arg0, _arg1)
}

// SetGIcon sets the icon of @action.
//
// Deprecated: Use #GAction instead, and g_menu_item_set_icon() to set an icon
// on a Item associated with a #GAction, or gtk_container_add() to add a Image
// to a Button.
func (action *Action) SetGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer((icon).(gextras.Nativer).Native()))

	C.gtk_action_set_gicon(_arg0, _arg1)
}

// SetIconName sets the icon name on @action
//
// Deprecated: Use #GAction instead, and g_menu_item_set_icon() to set an icon
// on a Item associated with a #GAction, or gtk_container_add() to add a Image
// to a Button.
func (action *Action) SetIconName(iconName string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_icon_name(_arg0, _arg1)
}

// SetIsImportant sets whether the action is important, this attribute is used
// primarily by toolbar items to decide whether to show a label or not.
//
// Deprecated: Use #GAction instead, and control and monitor whether labels are
// shown directly.
func (action *Action) SetIsImportant(isImportant bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_is_important(_arg0, _arg1)
}

// SetLabel sets the label of @action.
//
// Deprecated: Use #GAction instead, and set a label on a menu item with
// g_menu_item_set_label(). For Actionable widgets, use the widget-specific API
// to set a label.
func (action *Action) SetLabel(label string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_label(_arg0, _arg1)
}

// SetSensitive sets the :sensitive property of the action to @sensitive. Note
// that this doesn’t necessarily mean effective sensitivity. See
// gtk_action_is_sensitive() for that.
//
// Deprecated: Use g_simple_action_set_enabled() on a Action instead.
func (action *Action) SetSensitive(sensitive bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if sensitive {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_sensitive(_arg0, _arg1)
}

// SetShortLabel sets a shorter label text on @action.
//
// Deprecated: Use #GAction instead, which has no equivalent of short labels.
func (action *Action) SetShortLabel(shortLabel string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(shortLabel))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_short_label(_arg0, _arg1)
}

// SetStockID sets the stock id on @action
//
// Deprecated: Use #GAction instead, which has no equivalent of stock items.
func (action *Action) SetStockID(stockId string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_stock_id(_arg0, _arg1)
}

// SetTooltip sets the tooltip text on @action
//
// Deprecated: Use #GAction instead, and set tooltips on associated Actionable
// widgets with gtk_widget_set_tooltip_text().
func (action *Action) SetTooltip(tooltip string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	_arg1 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_tooltip(_arg0, _arg1)
}

// SetVisible sets the :visible property of the action to @visible. Note that
// this doesn’t necessarily mean effective visibility. See
// gtk_action_is_visible() for that.
//
// Deprecated: Use #GAction instead, and control and monitor the state of
// Actionable widgets directly.
func (action *Action) SetVisible(visible bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible(_arg0, _arg1)
}

// SetVisibleHorizontal sets whether @action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
func (action *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_horizontal(_arg0, _arg1)
}

// SetVisibleVertical sets whether @action is visible when vertical
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
func (action *Action) SetVisibleVertical(visibleVertical bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_vertical(_arg0, _arg1)
}

// UnblockActivate: reenable activation signals from the action
//
// Deprecated: Use g_simple_action_set_enabled() to enable the Action instead.
func (action *Action) UnblockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_action_unblock_activate(_arg0)
}
