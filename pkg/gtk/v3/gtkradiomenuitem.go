// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_RadioMenuItemClass_group_changed(GtkRadioMenuItem*);
// extern void _gotk4_gtk3_RadioMenuItem_ConnectGroupChanged(gpointer, guintptr);
import "C"

// glib.Type values for gtkradiomenuitem.go.
var GTypeRadioMenuItem = externglib.Type(C.gtk_radio_menu_item_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeRadioMenuItem, F: marshalRadioMenuItem},
	})
}

// RadioMenuItemOverrider contains methods that are overridable.
type RadioMenuItemOverrider interface {
	GroupChanged()
}

// RadioMenuItem: radio menu item is a check menu item that belongs to a group.
// At each instant exactly one of the radio menu items from a group is selected.
//
// The group list does not need to be freed, as each RadioMenuItem will remove
// itself and its list item when it is destroyed.
//
// The correct way to create a group of radio menu items is approximatively
// this:
//
// How to create a group of radio menu items.
//
//    menuitem
//    ├── radio.left
//    ╰── <child>
//
// GtkRadioMenuItem has a main CSS node with name menuitem, and a subnode with
// name radio, which gets the .left or .right style class.
type RadioMenuItem struct {
	_ [0]func() // equal guard
	CheckMenuItem
}

var (
	_ Binner              = (*RadioMenuItem)(nil)
	_ externglib.Objector = (*RadioMenuItem)(nil)
)

func classInitRadioMenuItemmer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkRadioMenuItemClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkRadioMenuItemClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ GroupChanged() }); ok {
		pclass.group_changed = (*[0]byte)(C._gotk4_gtk3_RadioMenuItemClass_group_changed)
	}
}

//export _gotk4_gtk3_RadioMenuItemClass_group_changed
func _gotk4_gtk3_RadioMenuItemClass_group_changed(arg0 *C.GtkRadioMenuItem) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ GroupChanged() })

	iface.GroupChanged()
}

func wrapRadioMenuItem(obj *externglib.Object) *RadioMenuItem {
	return &RadioMenuItem{
		CheckMenuItem: CheckMenuItem{
			MenuItem: MenuItem{
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
				Actionable: Actionable{
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
				Activatable: Activatable{
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	return wrapRadioMenuItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_RadioMenuItem_ConnectGroupChanged
func _gotk4_gtk3_RadioMenuItem_ConnectGroupChanged(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (radioMenuItem *RadioMenuItem) ConnectGroupChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(radioMenuItem, "group-changed", false, unsafe.Pointer(C._gotk4_gtk3_RadioMenuItem_ConnectGroupChanged), f)
}

// NewRadioMenuItem creates a new RadioMenuItem.
//
// The function takes the following parameters:
//
//    - group (optional) to which the radio menu item is to be attached, or NULL.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItem(group []*RadioMenuItem) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	_cret = C.gtk_radio_menu_item_new(_arg1)
	runtime.KeepAlive(group)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemFromWidget creates a new RadioMenuItem adding it to the same
// group as group.
//
// The function takes the following parameters:
//
//    - group (optional): existing RadioMenuItem.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItemFromWidget(group *RadioMenuItem) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(group).Native()))
	}

	_cret = C.gtk_radio_menu_item_new_from_widget(_arg1)
	runtime.KeepAlive(group)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithLabel creates a new RadioMenuItem whose child is a simple
// Label.
//
// The function takes the following parameters:
//
//    - group (optional): group the radio menu item is inside, or NULL.
//    - label: text for the label.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItemWithLabel(group []*RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_menu_item_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithLabelFromWidget creates a new GtkRadioMenuItem whose
// child is a simple GtkLabel. The new RadioMenuItem is added to the same group
// as group.
//
// The function takes the following parameters:
//
//    - group (optional): existing RadioMenuItem.
//    - label (optional): text for the label.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItemWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _arg2 *C.gchar            // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(group).Native()))
	}
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_radio_menu_item_new_with_label_from_widget(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithMnemonic creates a new RadioMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
//
// The function takes the following parameters:
//
//    - group (optional): group the radio menu item is inside, or NULL.
//    - label: text of the button, with an underscore in front of the mnemonic
//      character.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItemWithMnemonic(group []*RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_menu_item_new_with_mnemonic(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithMnemonicFromWidget creates a new GtkRadioMenuItem
// containing a label. The label will be created using
// gtk_label_new_with_mnemonic(), so underscores in label indicate the mnemonic
// for the menu item.
//
// The new RadioMenuItem is added to the same group as group.
//
// The function takes the following parameters:
//
//    - group (optional): existing RadioMenuItem.
//    - label (optional): text of the button, with an underscore in front of the
//      mnemonic character.
//
// The function returns the following values:
//
//    - radioMenuItem: new RadioMenuItem.
//
func NewRadioMenuItemWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _arg2 *C.gchar            // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(group).Native()))
	}
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_radio_menu_item_new_with_mnemonic_from_widget(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// Group returns the group to which the radio menu item belongs, as a #GList of
// RadioMenuItem. The list belongs to GTK+ and should not be freed.
//
// The function returns the following values:
//
//    - sList: group of radio_menu_item.
//
func (radioMenuItem *RadioMenuItem) Group() []*RadioMenuItem {
	var _arg0 *C.GtkRadioMenuItem // out
	var _cret *C.GSList           // in

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(radioMenuItem).Native()))

	_cret = C.gtk_radio_menu_item_get_group(_arg0)
	runtime.KeepAlive(radioMenuItem)

	var _sList []*RadioMenuItem // out

	_sList = make([]*RadioMenuItem, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkRadioMenuItem)(v)
		var dst *RadioMenuItem // out
		dst = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// JoinGroup joins a RadioMenuItem object to the group of another RadioMenuItem
// object.
//
// This function should be used by language bindings to avoid the memory
// manangement of the opaque List of gtk_radio_menu_item_get_group() and
// gtk_radio_menu_item_set_group().
//
// A common way to set up a group of RadioMenuItem instances is:
//
//      GtkRadioMenuItem *last_item = NULL;
//
//      while ( ...more items to add... )
//        {
//          GtkRadioMenuItem *radio_item;
//
//          radio_item = gtk_radio_menu_item_new (...);
//
//          gtk_radio_menu_item_join_group (radio_item, last_item);
//          last_item = radio_item;
//        }.
//
// The function takes the following parameters:
//
//    - groupSource (optional) whose group we are joining, or NULL to remove the
//      radio_menu_item from its current group.
//
func (radioMenuItem *RadioMenuItem) JoinGroup(groupSource *RadioMenuItem) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GtkRadioMenuItem // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(radioMenuItem).Native()))
	if groupSource != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(groupSource).Native()))
	}

	C.gtk_radio_menu_item_join_group(_arg0, _arg1)
	runtime.KeepAlive(radioMenuItem)
	runtime.KeepAlive(groupSource)
}

// SetGroup sets the group of a radio menu item, or changes it.
//
// The function takes the following parameters:
//
//    - group (optional): new group, or NULL.
//
func (radioMenuItem *RadioMenuItem) SetGroup(group []*RadioMenuItem) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GSList           // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(radioMenuItem).Native()))
	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer(externglib.InternObject(src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	C.gtk_radio_menu_item_set_group(_arg0, _arg1)
	runtime.KeepAlive(radioMenuItem)
	runtime.KeepAlive(group)
}
