// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_menu_attribute_iter_get_type()), F: marshalMenuAttributeIterer},
		{T: externglib.Type(C.g_menu_link_iter_get_type()), F: marshalMenuLinkIterer},
		{T: externglib.Type(C.g_menu_model_get_type()), F: marshalMenuModeller},
	})
}

// MENU_ATTRIBUTE_ACTION: menu item attribute which holds the action name of the
// item. Action names are namespaced with an identifier for the action group in
// which the action resides. For example, "win." for window-specific actions and
// "app." for application-wide actions.
//
// See also g_menu_model_get_item_attribute() and g_menu_item_set_attribute().
const MENU_ATTRIBUTE_ACTION = "action"

// MENU_ATTRIBUTE_ACTION_NAMESPACE: menu item attribute that holds the namespace
// for all action names in menus that are linked from this item.
const MENU_ATTRIBUTE_ACTION_NAMESPACE = "action-namespace"

// MENU_ATTRIBUTE_ICON: menu item attribute which holds the icon of the item.
//
// The icon is stored in the format returned by g_icon_serialize().
//
// This attribute is intended only to represent 'noun' icons such as favicons
// for a webpage, or application icons. It should not be used for 'verbs' (ie:
// stock icons).
const MENU_ATTRIBUTE_ICON = "icon"

// MENU_ATTRIBUTE_LABEL: menu item attribute which holds the label of the item.
const MENU_ATTRIBUTE_LABEL = "label"

// MENU_ATTRIBUTE_TARGET: menu item attribute which holds the target with which
// the item's action will be activated.
//
// See also g_menu_item_set_action_and_target()
const MENU_ATTRIBUTE_TARGET = "target"

// MENU_LINK_SECTION: name of the link that associates a menu item with a
// section. The linked menu will usually be shown in place of the menu item,
// using the item's label as a header.
//
// See also g_menu_item_set_link().
const MENU_LINK_SECTION = "section"

// MENU_LINK_SUBMENU: name of the link that associates a menu item with a
// submenu.
//
// See also g_menu_item_set_link().
const MENU_LINK_SUBMENU = "submenu"

// MenuAttributeIterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuAttributeIterOverrider interface {
	// Next: this function combines g_menu_attribute_iter_next() with
	// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
	//
	// First the iterator is advanced to the next (possibly first) attribute. If
	// that fails, then FALSE is returned and there are no other effects.
	//
	// If successful, name and value are set to the name and value of the
	// attribute that has just been advanced to. At this point,
	// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value()
	// will return the same values again.
	//
	// The value returned in name remains valid for as long as the iterator
	// remains at the current position. The value returned in value must be
	// unreffed using g_variant_unref() when it is no longer in use.
	Next() (string, *glib.Variant, bool)
}

// MenuAttributeIter is an opaque structure type. You must access it using the
// functions below.
type MenuAttributeIter struct {
	*externglib.Object
}

// MenuAttributeIterer describes MenuAttributeIter's abstract methods.
type MenuAttributeIterer interface {
	externglib.Objector

	// Name gets the name of the attribute at the current iterator position, as
	// a string.
	Name() string
	// GetNext: this function combines g_menu_attribute_iter_next() with
	// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
	GetNext() (string, *glib.Variant, bool)
	// Value gets the value of the attribute at the current iterator position.
	Value() *glib.Variant
	// Next attempts to advance the iterator to the next (possibly first)
	// attribute.
	Next() bool
}

var _ MenuAttributeIterer = (*MenuAttributeIter)(nil)

func wrapMenuAttributeIter(obj *externglib.Object) *MenuAttributeIter {
	return &MenuAttributeIter{
		Object: obj,
	}
}

func marshalMenuAttributeIterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuAttributeIter(obj), nil
}

// Name gets the name of the attribute at the current iterator position, as a
// string.
//
// The iterator is not advanced.
func (iter *MenuAttributeIter) Name() string {
	var _arg0 *C.GMenuAttributeIter // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_attribute_iter_get_name(_arg0)
	runtime.KeepAlive(iter)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetNext: this function combines g_menu_attribute_iter_next() with
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) attribute. If
// that fails, then FALSE is returned and there are no other effects.
//
// If successful, name and value are set to the name and value of the attribute
// that has just been advanced to. At this point,
// g_menu_attribute_iter_get_name() and g_menu_attribute_iter_get_value() will
// return the same values again.
//
// The value returned in name remains valid for as long as the iterator remains
// at the current position. The value returned in value must be unreffed using
// g_variant_unref() when it is no longer in use.
func (iter *MenuAttributeIter) GetNext() (string, *glib.Variant, bool) {
	var _arg0 *C.GMenuAttributeIter // out
	var _arg1 *C.gchar              // in
	var _arg2 *C.GVariant           // in
	var _cret C.gboolean            // in

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_attribute_iter_get_next(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(iter)

	var _outName string      // out
	var _value *glib.Variant // out
	var _ok bool             // out

	if _arg1 != nil {
		_outName = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	if _arg2 != nil {
		_value = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_value)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}
	if _cret != 0 {
		_ok = true
	}

	return _outName, _value, _ok
}

// Value gets the value of the attribute at the current iterator position.
//
// The iterator is not advanced.
func (iter *MenuAttributeIter) Value() *glib.Variant {
	var _arg0 *C.GMenuAttributeIter // out
	var _cret *C.GVariant           // in

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_attribute_iter_get_value(_arg0)
	runtime.KeepAlive(iter)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_variant)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_variant_unref((*C.GVariant)(intern.C))
		},
	)

	return _variant
}

// Next attempts to advance the iterator to the next (possibly first) attribute.
//
// TRUE is returned on success, or FALSE if there are no more attributes.
//
// You must call this function when you first acquire the iterator to advance it
// to the first attribute (and determine if the first attribute exists at all).
func (iter *MenuAttributeIter) Next() bool {
	var _arg0 *C.GMenuAttributeIter // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GMenuAttributeIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_attribute_iter_next(_arg0)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenuLinkIterOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuLinkIterOverrider interface {
	// Next: this function combines g_menu_link_iter_next() with
	// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
	//
	// First the iterator is advanced to the next (possibly first) link. If that
	// fails, then FALSE is returned and there are no other effects.
	//
	// If successful, out_link and value are set to the name and Model of the
	// link that has just been advanced to. At this point,
	// g_menu_link_iter_get_name() and g_menu_link_iter_get_value() will return
	// the same values again.
	//
	// The value returned in out_link remains valid for as long as the iterator
	// remains at the current position. The value returned in value must be
	// unreffed using g_object_unref() when it is no longer in use.
	Next() (string, MenuModeller, bool)
}

// MenuLinkIter is an opaque structure type. You must access it using the
// functions below.
type MenuLinkIter struct {
	*externglib.Object
}

// MenuLinkIterer describes MenuLinkIter's abstract methods.
type MenuLinkIterer interface {
	externglib.Objector

	// Name gets the name of the link at the current iterator position.
	Name() string
	// GetNext: this function combines g_menu_link_iter_next() with
	// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
	GetNext() (string, MenuModeller, bool)
	// Value gets the linked Model at the current iterator position.
	Value() MenuModeller
	// Next attempts to advance the iterator to the next (possibly first) link.
	Next() bool
}

var _ MenuLinkIterer = (*MenuLinkIter)(nil)

func wrapMenuLinkIter(obj *externglib.Object) *MenuLinkIter {
	return &MenuLinkIter{
		Object: obj,
	}
}

func marshalMenuLinkIterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuLinkIter(obj), nil
}

// Name gets the name of the link at the current iterator position.
//
// The iterator is not advanced.
func (iter *MenuLinkIter) Name() string {
	var _arg0 *C.GMenuLinkIter // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_link_iter_get_name(_arg0)
	runtime.KeepAlive(iter)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// GetNext: this function combines g_menu_link_iter_next() with
// g_menu_link_iter_get_name() and g_menu_link_iter_get_value().
//
// First the iterator is advanced to the next (possibly first) link. If that
// fails, then FALSE is returned and there are no other effects.
//
// If successful, out_link and value are set to the name and Model of the link
// that has just been advanced to. At this point, g_menu_link_iter_get_name()
// and g_menu_link_iter_get_value() will return the same values again.
//
// The value returned in out_link remains valid for as long as the iterator
// remains at the current position. The value returned in value must be unreffed
// using g_object_unref() when it is no longer in use.
func (iter *MenuLinkIter) GetNext() (string, MenuModeller, bool) {
	var _arg0 *C.GMenuLinkIter // out
	var _arg1 *C.gchar         // in
	var _arg2 *C.GMenuModel    // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_link_iter_get_next(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(iter)

	var _outLink string     // out
	var _value MenuModeller // out
	var _ok bool            // out

	if _arg1 != nil {
		_outLink = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
	}
	if _arg2 != nil {
		{
			object := externglib.AssumeOwnership(unsafe.Pointer(_arg2))
			rv, ok := (externglib.CastObject(object)).(MenuModeller)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuModeller")
			}
			_value = rv
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _outLink, _value, _ok
}

// Value gets the linked Model at the current iterator position.
//
// The iterator is not advanced.
func (iter *MenuLinkIter) Value() MenuModeller {
	var _arg0 *C.GMenuLinkIter // out
	var _cret *C.GMenuModel    // in

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_link_iter_get_value(_arg0)
	runtime.KeepAlive(iter)

	var _menuModel MenuModeller // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(MenuModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuModeller")
		}
		_menuModel = rv
	}

	return _menuModel
}

// Next attempts to advance the iterator to the next (possibly first) link.
//
// TRUE is returned on success, or FALSE if there are no more links.
//
// You must call this function when you first acquire the iterator to advance it
// to the first link (and determine if the first link exists at all).
func (iter *MenuLinkIter) Next() bool {
	var _arg0 *C.GMenuLinkIter // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GMenuLinkIter)(unsafe.Pointer(iter.Native()))

	_cret = C.g_menu_link_iter_next(_arg0)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenuModelOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MenuModelOverrider interface {
	// ItemAttributeValue queries the item at position item_index in model for
	// the attribute specified by attribute.
	//
	// If expected_type is non-NULL then it specifies the expected type of the
	// attribute. If it is NULL then any type will be accepted.
	//
	// If the attribute exists and matches expected_type (or if the expected
	// type is unspecified) then the value is returned.
	//
	// If the attribute does not exist, or does not match the expected type then
	// NULL is returned.
	ItemAttributeValue(itemIndex int, attribute string, expectedType *glib.VariantType) *glib.Variant
	// ItemAttributes gets all the attributes associated with the item in the
	// menu model.
	ItemAttributes(itemIndex int) map[string]glib.Variant
	// ItemLink queries the item at position item_index in model for the link
	// specified by link.
	//
	// If the link exists, the linked Model is returned. If the link does not
	// exist, NULL is returned.
	ItemLink(itemIndex int, link string) MenuModeller
	// ItemLinks gets all the links associated with the item in the menu model.
	ItemLinks(itemIndex int) map[string]MenuModeller
	// NItems: query the number of items in model.
	NItems() int
	// IsMutable queries if model is mutable.
	//
	// An immutable Model will never emit the Model::items-changed signal.
	// Consumers of the model may make optimisations accordingly.
	IsMutable() bool
	// IterateItemAttributes creates a AttributeIter to iterate over the
	// attributes of the item at position item_index in model.
	//
	// You must free the iterator with g_object_unref() when you are done.
	IterateItemAttributes(itemIndex int) MenuAttributeIterer
	// IterateItemLinks creates a LinkIter to iterate over the links of the item
	// at position item_index in model.
	//
	// You must free the iterator with g_object_unref() when you are done.
	IterateItemLinks(itemIndex int) MenuLinkIterer
}

// MenuModel represents the contents of a menu -- an ordered list of menu items.
// The items are associated with actions, which can be activated through them.
// Items can be grouped in sections, and may have submenus associated with them.
// Both items and sections usually have some representation data, such as labels
// or icons. The type of the associated action (ie whether it is stateful, and
// what kind of state it has) can influence the representation of the item.
//
// The conceptual model of menus in Model is hierarchical: sections and submenus
// are again represented by Models. Menus themselves do not define their own
// roles. Rather, the role of a particular Model is defined by the item that
// references it (or, in the case of the 'root' menu, is defined by the context
// in which it is used).
//
// As an example, consider the visible portions of this menu:
//
//
// An example menu
//
// ! (menu-example.png)
//
// There are 8 "menus" visible in the screenshot: one menubar, two submenus and
// 5 sections:
//
// - the toplevel menubar (containing 4 items)
//
// - the View submenu (containing 3 sections)
//
// - the first section of the View submenu (containing 2 items)
//
// - the second section of the View submenu (containing 1 item)
//
// - the final section of the View submenu (containing 1 item)
//
// - the Highlight Mode submenu (containing 2 sections)
//
// - the Sources section (containing 2 items)
//
// - the Markup section (containing 2 items)
//
// The [example][menu-model] illustrates the conceptual connection between these
// 8 menus. Each large block in the figure represents a menu and the smaller
// blocks within the large block represent items in that menu. Some items
// contain references to other menus.
//
//
// A menu example
//
// ! (menu-model.png)
//
// Notice that the separators visible in the [example][menu-example] appear
// nowhere in the [menu model][menu-model]. This is because separators are not
// explicitly represented in the menu model. Instead, a separator is inserted
// between any two non-empty sections of a menu. Section items can have labels
// just like any other item. In that case, a display system may show a section
// header instead of a separator.
//
// The motivation for this abstract model of application controls is that modern
// user interfaces tend to make these controls available outside the
// application. Examples include global menus, jumplists, dash boards, etc. To
// support such uses, it is necessary to 'export' information about actions and
// their representation in menus, which is exactly what the [GActionGroup
// exporter][gio-GActionGroup-exporter] and the [GMenuModel
// exporter][gio-GMenuModel-exporter] do for Group and Model. The client-side
// counterparts to make use of the exported information are BusActionGroup and
// BusMenuModel.
//
// The API of Model is very generic, with iterators for the attributes and links
// of an item, see g_menu_model_iterate_item_attributes() and
// g_menu_model_iterate_item_links(). The 'standard' attributes and link types
// have predefined names: G_MENU_ATTRIBUTE_LABEL, G_MENU_ATTRIBUTE_ACTION,
// G_MENU_ATTRIBUTE_TARGET, G_MENU_LINK_SECTION and G_MENU_LINK_SUBMENU.
//
// Items in a Model represent active controls if they refer to an action that
// can get activated when the user interacts with the menu item. The reference
// to the action is encoded by the string id in the G_MENU_ATTRIBUTE_ACTION
// attribute. An action id uniquely identifies an action in an action group.
// Which action group(s) provide actions depends on the context in which the
// menu model is used. E.g. when the model is exported as the application menu
// of a Application, actions can be application-wide or window-specific (and
// thus come from two different action groups). By convention, the
// application-wide actions have names that start with "app.", while the names
// of window-specific actions start with "win.".
//
// While a wide variety of stateful actions is possible, the following is the
// minimum that is expected to be supported by all users of exported menu
// information:
//
// - an action with no parameter type and no state
//
// - an action with no parameter type and boolean state
//
// - an action with string parameter type and string state
//
//
// Stateless
//
// A stateless action typically corresponds to an ordinary menu item.
//
// Selecting such a menu item will activate the action (with no parameter).
//
//
// Boolean State
//
// An action with a boolean state will most typically be used with a "toggle" or
// "switch" menu item. The state can be set directly, but activating the action
// (with no parameter) results in the state being toggled.
//
// Selecting a toggle menu item will activate the action. The menu item should
// be rendered as "checked" when the state is true.
//
//
// String Parameter and State
//
// Actions with string parameters and state will most typically be used to
// represent an enumerated choice over the items available for a group of radio
// menu items. Activating the action with a string parameter is equivalent to
// setting that parameter as the state.
//
// Radio menu items, in addition to being associated with the action, will have
// a target value. Selecting that menu item will result in activation of the
// action with the target value as the parameter. The menu item should be
// rendered as "selected" when the state of the action is equal to the target
// value of the menu item.
type MenuModel struct {
	*externglib.Object
}

// MenuModeller describes MenuModel's abstract methods.
type MenuModeller interface {
	externglib.Objector

	// ItemAttributeValue queries the item at position item_index in model for
	// the attribute specified by attribute.
	ItemAttributeValue(itemIndex int, attribute string, expectedType *glib.VariantType) *glib.Variant
	// ItemLink queries the item at position item_index in model for the link
	// specified by link.
	ItemLink(itemIndex int, link string) MenuModeller
	// NItems: query the number of items in model.
	NItems() int
	// IsMutable queries if model is mutable.
	IsMutable() bool
	// ItemsChanged requests emission of the Model::items-changed signal on
	// model.
	ItemsChanged(position int, removed int, added int)
	// IterateItemAttributes creates a AttributeIter to iterate over the
	// attributes of the item at position item_index in model.
	IterateItemAttributes(itemIndex int) MenuAttributeIterer
	// IterateItemLinks creates a LinkIter to iterate over the links of the item
	// at position item_index in model.
	IterateItemLinks(itemIndex int) MenuLinkIterer
}

var _ MenuModeller = (*MenuModel)(nil)

func wrapMenuModel(obj *externglib.Object) *MenuModel {
	return &MenuModel{
		Object: obj,
	}
}

func marshalMenuModeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMenuModel(obj), nil
}

// ItemAttributeValue queries the item at position item_index in model for the
// attribute specified by attribute.
//
// If expected_type is non-NULL then it specifies the expected type of the
// attribute. If it is NULL then any type will be accepted.
//
// If the attribute exists and matches expected_type (or if the expected type is
// unspecified) then the value is returned.
//
// If the attribute does not exist, or does not match the expected type then
// NULL is returned.
func (model *MenuModel) ItemAttributeValue(itemIndex int, attribute string, expectedType *glib.VariantType) *glib.Variant {
	var _arg0 *C.GMenuModel   // out
	var _arg1 C.gint          // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.GVariantType // out
	var _cret *C.GVariant     // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.gint(itemIndex)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg2))
	if expectedType != nil {
		_arg3 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(expectedType)))
	}

	_cret = C.g_menu_model_get_item_attribute_value(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)
	runtime.KeepAlive(attribute)
	runtime.KeepAlive(expectedType)

	var _variant *glib.Variant // out

	if _cret != nil {
		_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_variant)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
	}

	return _variant
}

// ItemLink queries the item at position item_index in model for the link
// specified by link.
//
// If the link exists, the linked Model is returned. If the link does not exist,
// NULL is returned.
func (model *MenuModel) ItemLink(itemIndex int, link string) MenuModeller {
	var _arg0 *C.GMenuModel // out
	var _arg1 C.gint        // out
	var _arg2 *C.gchar      // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.gint(itemIndex)
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(link)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_menu_model_get_item_link(_arg0, _arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)
	runtime.KeepAlive(link)

	var _menuModel MenuModeller // out

	if _cret != nil {
		{
			object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			rv, ok := (externglib.CastObject(object)).(MenuModeller)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuModeller")
			}
			_menuModel = rv
		}
	}

	return _menuModel
}

// NItems: query the number of items in model.
func (model *MenuModel) NItems() int {
	var _arg0 *C.GMenuModel // out
	var _cret C.gint        // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.g_menu_model_get_n_items(_arg0)
	runtime.KeepAlive(model)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IsMutable queries if model is mutable.
//
// An immutable Model will never emit the Model::items-changed signal. Consumers
// of the model may make optimisations accordingly.
func (model *MenuModel) IsMutable() bool {
	var _arg0 *C.GMenuModel // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	_cret = C.g_menu_model_is_mutable(_arg0)
	runtime.KeepAlive(model)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ItemsChanged requests emission of the Model::items-changed signal on model.
//
// This function should never be called except by Model subclasses. Any other
// calls to this function will very likely lead to a violation of the interface
// of the model.
//
// The implementation should update its internal representation of the menu
// before emitting the signal. The implementation should further expect to
// receive queries about the new state of the menu (and particularly added menu
// items) while signal handlers are running.
//
// The implementation must dispatch this call directly from a mainloop entry and
// not in response to calls -- particularly those from the Model API. Said
// another way: the menu must not change while user code is running without
// returning to the mainloop.
func (model *MenuModel) ItemsChanged(position int, removed int, added int) {
	var _arg0 *C.GMenuModel // out
	var _arg1 C.gint        // out
	var _arg2 C.gint        // out
	var _arg3 C.gint        // out

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.gint(position)
	_arg2 = C.gint(removed)
	_arg3 = C.gint(added)

	C.g_menu_model_items_changed(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(model)
	runtime.KeepAlive(position)
	runtime.KeepAlive(removed)
	runtime.KeepAlive(added)
}

// IterateItemAttributes creates a AttributeIter to iterate over the attributes
// of the item at position item_index in model.
//
// You must free the iterator with g_object_unref() when you are done.
func (model *MenuModel) IterateItemAttributes(itemIndex int) MenuAttributeIterer {
	var _arg0 *C.GMenuModel         // out
	var _arg1 C.gint                // out
	var _cret *C.GMenuAttributeIter // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.gint(itemIndex)

	_cret = C.g_menu_model_iterate_item_attributes(_arg0, _arg1)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)

	var _menuAttributeIter MenuAttributeIterer // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(MenuAttributeIterer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuAttributeIterer")
		}
		_menuAttributeIter = rv
	}

	return _menuAttributeIter
}

// IterateItemLinks creates a LinkIter to iterate over the links of the item at
// position item_index in model.
//
// You must free the iterator with g_object_unref() when you are done.
func (model *MenuModel) IterateItemLinks(itemIndex int) MenuLinkIterer {
	var _arg0 *C.GMenuModel    // out
	var _arg1 C.gint           // out
	var _cret *C.GMenuLinkIter // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	_arg1 = C.gint(itemIndex)

	_cret = C.g_menu_model_iterate_item_links(_arg0, _arg1)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)

	var _menuLinkIter MenuLinkIterer // out

	{
		object := externglib.AssumeOwnership(unsafe.Pointer(_cret))
		rv, ok := (externglib.CastObject(object)).(MenuLinkIterer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.MenuLinkIterer")
		}
		_menuLinkIter = rv
	}

	return _menuLinkIter
}
