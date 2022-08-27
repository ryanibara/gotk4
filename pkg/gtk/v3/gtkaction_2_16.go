// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// BlockActivate: disable activation signals from the action
//
// This is needed when updating the state of your proxy Activatable widget could
// result in calling gtk_action_activate(), this is a convenience function to
// avoid recursing in those cases (updating toggle state for instance).
//
// Deprecated: Use g_simple_action_set_enabled() to disable the Action instead.
func (action *Action) BlockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	C.gtk_action_block_activate(_arg0)
	runtime.KeepAlive(action)
}

// GIcon gets the gicon of action.
//
// Deprecated: Use #GAction instead, and g_menu_item_get_attribute_value() to
// get an icon from a Item associated with a #GAction.
//
// The function returns the following values:
//
//    - icon action’s #GIcon if one is set.
//
func (action *Action) GIcon() *gio.Icon {
	var _arg0 *C.GtkAction // out
	var _cret *C.GIcon     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_gicon(_arg0)
	runtime.KeepAlive(action)

	var _icon *gio.Icon // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_icon = &gio.Icon{
			Object: obj,
		}
	}

	return _icon
}

// IconName gets the icon name of action.
//
// Deprecated: Use #GAction instead, and g_menu_item_get_attribute_value() to
// get an icon from a Item associated with a #GAction.
//
// The function returns the following values:
//
//    - utf8: icon name.
//
func (action *Action) IconName() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_icon_name(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// IsImportant checks whether action is important or not
//
// Deprecated: Use #GAction instead, and control and monitor whether labels are
// shown directly.
//
// The function returns the following values:
//
//    - ok: whether action is important.
//
func (action *Action) IsImportant() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_is_important(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label gets the label text of action.
//
// Deprecated: Use #GAction instead, and get a label from a menu item with
// g_menu_item_get_attribute_value(). For Actionable widgets, use the
// widget-specific API to get a label.
//
// The function returns the following values:
//
//    - utf8: label text.
//
func (action *Action) Label() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_label(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ShortLabel gets the short label text of action.
//
// Deprecated: Use #GAction instead, which has no equivalent of short labels.
//
// The function returns the following values:
//
//    - utf8: short label text.
//
func (action *Action) ShortLabel() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_short_label(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// StockID gets the stock id of action.
//
// Deprecated: Use #GAction instead, which has no equivalent of stock items.
//
// The function returns the following values:
//
//    - utf8: stock id.
//
func (action *Action) StockID() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_stock_id(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Tooltip gets the tooltip text of action.
//
// Deprecated: Use #GAction instead, and get tooltips from associated Actionable
// widgets with gtk_widget_get_tooltip_text().
//
// The function returns the following values:
//
//    - utf8: tooltip text.
//
func (action *Action) Tooltip() string {
	var _arg0 *C.GtkAction // out
	var _cret *C.gchar     // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_tooltip(_arg0)
	runtime.KeepAlive(action)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// VisibleHorizontal checks whether action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
//
// The function returns the following values:
//
//    - ok: whether action is visible when horizontal.
//
func (action *Action) VisibleHorizontal() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_visible_horizontal(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleVertical checks whether action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
//
// The function returns the following values:
//
//    - ok: whether action is visible when horizontal.
//
func (action *Action) VisibleVertical() bool {
	var _arg0 *C.GtkAction // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	_cret = C.gtk_action_get_visible_vertical(_arg0)
	runtime.KeepAlive(action)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetGIcon sets the icon of action.
//
// Deprecated: Use #GAction instead, and g_menu_item_set_icon() to set an icon
// on a Item associated with a #GAction, or gtk_container_add() to add a Image
// to a Button.
//
// The function takes the following parameters:
//
//    - icon to set.
//
func (action *Action) SetGIcon(icon gio.Iconner) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	C.gtk_action_set_gicon(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(icon)
}

// SetIconName sets the icon name on action
//
// Deprecated: Use #GAction instead, and g_menu_item_set_icon() to set an icon
// on a Item associated with a #GAction, or gtk_container_add() to add a Image
// to a Button.
//
// The function takes the following parameters:
//
//    - iconName: icon name to set.
//
func (action *Action) SetIconName(iconName string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_icon_name(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(iconName)
}

// SetIsImportant sets whether the action is important, this attribute is used
// primarily by toolbar items to decide whether to show a label or not.
//
// Deprecated: Use #GAction instead, and control and monitor whether labels are
// shown directly.
//
// The function takes the following parameters:
//
//    - isImportant: TRUE to make the action important.
//
func (action *Action) SetIsImportant(isImportant bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if isImportant {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_is_important(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(isImportant)
}

// SetLabel sets the label of action.
//
// Deprecated: Use #GAction instead, and set a label on a menu item with
// g_menu_item_set_label(). For Actionable widgets, use the widget-specific API
// to set a label.
//
// The function takes the following parameters:
//
//    - label text to set.
//
func (action *Action) SetLabel(label string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_label(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(label)
}

// SetShortLabel sets a shorter label text on action.
//
// Deprecated: Use #GAction instead, which has no equivalent of short labels.
//
// The function takes the following parameters:
//
//    - shortLabel: label text to set.
//
func (action *Action) SetShortLabel(shortLabel string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(shortLabel)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_short_label(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(shortLabel)
}

// SetStockID sets the stock id on action
//
// Deprecated: Use #GAction instead, which has no equivalent of stock items.
//
// The function takes the following parameters:
//
//    - stockId: stock id.
//
func (action *Action) SetStockID(stockId string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_stock_id(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(stockId)
}

// SetTooltip sets the tooltip text on action
//
// Deprecated: Use #GAction instead, and set tooltips on associated Actionable
// widgets with gtk_widget_set_tooltip_text().
//
// The function takes the following parameters:
//
//    - tooltip text.
//
func (action *Action) SetTooltip(tooltip string) {
	var _arg0 *C.GtkAction // out
	var _arg1 *C.gchar     // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(tooltip)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_action_set_tooltip(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(tooltip)
}

// SetVisibleHorizontal sets whether action is visible when horizontal
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
//
// The function takes the following parameters:
//
//    - visibleHorizontal: whether the action is visible horizontally.
//
func (action *Action) SetVisibleHorizontal(visibleHorizontal bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if visibleHorizontal {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_horizontal(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(visibleHorizontal)
}

// SetVisibleVertical sets whether action is visible when vertical
//
// Deprecated: Use #GAction instead, and control and monitor the visibility of
// associated widgets and menu items directly.
//
// The function takes the following parameters:
//
//    - visibleVertical: whether the action is visible vertically.
//
func (action *Action) SetVisibleVertical(visibleVertical bool) {
	var _arg0 *C.GtkAction // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))
	if visibleVertical {
		_arg1 = C.TRUE
	}

	C.gtk_action_set_visible_vertical(_arg0, _arg1)
	runtime.KeepAlive(action)
	runtime.KeepAlive(visibleVertical)
}

// UnblockActivate: reenable activation signals from the action
//
// Deprecated: Use g_simple_action_set_enabled() to enable the Action instead.
func (action *Action) UnblockActivate() {
	var _arg0 *C.GtkAction // out

	_arg0 = (*C.GtkAction)(unsafe.Pointer(coreglib.InternObject(action).Native()))

	C.gtk_action_unblock_activate(_arg0)
	runtime.KeepAlive(action)
}
