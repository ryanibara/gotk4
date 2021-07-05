// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_image_menu_item_get_type()), F: marshalImageMenuItem},
	})
}

// ImageMenuItem is a menu item which has an icon next to the text label.
//
// This is functionally equivalent to:
//
//      GtkWidget *box = gtk_box_new (GTK_ORIENTATION_HORIZONTAL, 6);
//      GtkWidget *icon = gtk_image_new_from_icon_name ("folder-music-symbolic", GTK_ICON_SIZE_MENU);
//      GtkWidget *label = gtk_accel_label_new ("Music");
//      GtkWidget *menu_item = gtk_menu_item_new ();
//      GtkAccelGroup *accel_group = gtk_accel_group_new ();
//
//      gtk_container_add (GTK_CONTAINER (box), icon);
//
//      gtk_label_set_use_underline (GTK_LABEL (label), TRUE);
//      gtk_label_set_xalign (GTK_LABEL (label), 0.0);
//
//      gtk_widget_add_accelerator (menu_item, "activate", accel_group,
//                                  GDK_KEY_m, GDK_CONTROL_MASK, GTK_ACCEL_VISIBLE);
//      gtk_accel_label_set_accel_widget (GTK_ACCEL_LABEL (label), menu_item);
//
//      gtk_box_pack_end (GTK_BOX (box), label, TRUE, TRUE, 0);
//
//      gtk_container_add (GTK_CONTAINER (menu_item), box);
//
//      gtk_widget_show_all (menu_item);
type ImageMenuItem interface {
	MenuItem

	// AsActionable casts the class to the Actionable interface.
	AsActionable() Actionable
	// AsActivatable casts the class to the Activatable interface.
	AsActivatable() Activatable
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// AlwaysShowImage returns whether the menu item will ignore the
	// Settings:gtk-menu-images setting and always show the image, if available.
	//
	// Deprecated: since version 3.10.
	AlwaysShowImage() bool
	// Image gets the widget that is currently set as the image of
	// @image_menu_item. See gtk_image_menu_item_set_image().
	//
	// Deprecated: since version 3.10.
	Image() Widget
	// UseStock checks whether the label set in the menuitem is used as a stock
	// id to select the stock item for the item.
	//
	// Deprecated: since version 3.10.
	UseStock() bool
	// SetAccelGroupImageMenuItem specifies an @accel_group to add the menu
	// items accelerator to (this only applies to stock items so a stock item
	// must already be set, make sure to call
	// gtk_image_menu_item_set_use_stock() and gtk_menu_item_set_label() with a
	// valid stock item first).
	//
	// If you want this menu item to have changeable accelerators then you
	// shouldnt need this (see gtk_image_menu_item_new_from_stock()).
	//
	// Deprecated: since version 3.10.
	SetAccelGroupImageMenuItem(accelGroup AccelGroup)
	// SetAlwaysShowImageImageMenuItem: if true, the menu item will ignore the
	// Settings:gtk-menu-images setting and always show the image, if available.
	//
	// Use this property if the menuitem would be useless or hard to use without
	// the image.
	//
	// Deprecated: since version 3.10.
	SetAlwaysShowImageImageMenuItem(alwaysShow bool)
	// SetImageImageMenuItem sets the image of @image_menu_item to the given
	// widget. Note that it depends on the show-menu-images setting whether the
	// image will be displayed or not.
	//
	// Deprecated: since version 3.10.
	SetImageImageMenuItem(image Widget)
	// SetUseStockImageMenuItem: if true, the label set in the menuitem is used
	// as a stock id to select the stock item for the item.
	//
	// Deprecated: since version 3.10.
	SetUseStockImageMenuItem(useStock bool)
}

// imageMenuItem implements the ImageMenuItem class.
type imageMenuItem struct {
	MenuItem
}

// WrapImageMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapImageMenuItem(obj *externglib.Object) ImageMenuItem {
	return imageMenuItem{
		MenuItem: WrapMenuItem(obj),
	}
}

func marshalImageMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapImageMenuItem(obj), nil
}

// NewImageMenuItem creates a new ImageMenuItem with an empty label.
//
// Deprecated: since version 3.10.
func NewImageMenuItem() ImageMenuItem {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_image_menu_item_new()

	var _imageMenuItem ImageMenuItem // out

	_imageMenuItem = WrapImageMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemFromStock creates a new ImageMenuItem containing the image
// and text from a stock item. Some stock ids have preprocessor macros like
// K_STOCK_OK and K_STOCK_APPLY.
//
// If you want this menu item to have changeable accelerators, then pass in nil
// for accel_group. Next call gtk_menu_item_set_accel_path() with an appropriate
// path for the menu item, use gtk_stock_lookup() to look up the standard
// accelerator for the stock item, and if one is found, call
// gtk_accel_map_add_entry() to register it.
//
// Deprecated: since version 3.10.
func NewImageMenuItemFromStock(stockId string, accelGroup AccelGroup) ImageMenuItem {
	var _arg1 *C.gchar         // out
	var _arg2 *C.GtkAccelGroup // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	_cret = C.gtk_image_menu_item_new_from_stock(_arg1, _arg2)

	var _imageMenuItem ImageMenuItem // out

	_imageMenuItem = WrapImageMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemWithLabel creates a new ImageMenuItem containing a label.
//
// Deprecated: since version 3.10.
func NewImageMenuItemWithLabel(label string) ImageMenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_menu_item_new_with_label(_arg1)

	var _imageMenuItem ImageMenuItem // out

	_imageMenuItem = WrapImageMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

// NewImageMenuItemWithMnemonic creates a new ImageMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in @label indicate the mnemonic for the menu item.
//
// Deprecated: since version 3.10.
func NewImageMenuItemWithMnemonic(label string) ImageMenuItem {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_image_menu_item_new_with_mnemonic(_arg1)

	var _imageMenuItem ImageMenuItem // out

	_imageMenuItem = WrapImageMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _imageMenuItem
}

func (i imageMenuItem) AlwaysShowImage() bool {
	var _arg0 *C.GtkImageMenuItem // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_menu_item_get_always_show_image(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i imageMenuItem) Image() Widget {
	var _arg0 *C.GtkImageMenuItem // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_menu_item_get_image(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (i imageMenuItem) UseStock() bool {
	var _arg0 *C.GtkImageMenuItem // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))

	_cret = C.gtk_image_menu_item_get_use_stock(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (i imageMenuItem) SetAccelGroupImageMenuItem(accelGroup AccelGroup) {
	var _arg0 *C.GtkImageMenuItem // out
	var _arg1 *C.GtkAccelGroup    // out

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkAccelGroup)(unsafe.Pointer(accelGroup.Native()))

	C.gtk_image_menu_item_set_accel_group(_arg0, _arg1)
}

func (i imageMenuItem) SetAlwaysShowImageImageMenuItem(alwaysShow bool) {
	var _arg0 *C.GtkImageMenuItem // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	if alwaysShow {
		_arg1 = C.TRUE
	}

	C.gtk_image_menu_item_set_always_show_image(_arg0, _arg1)
}

func (i imageMenuItem) SetImageImageMenuItem(image Widget) {
	var _arg0 *C.GtkImageMenuItem // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(image.Native()))

	C.gtk_image_menu_item_set_image(_arg0, _arg1)
}

func (i imageMenuItem) SetUseStockImageMenuItem(useStock bool) {
	var _arg0 *C.GtkImageMenuItem // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkImageMenuItem)(unsafe.Pointer(i.Native()))
	if useStock {
		_arg1 = C.TRUE
	}

	C.gtk_image_menu_item_set_use_stock(_arg0, _arg1)
}

func (i imageMenuItem) AsActionable() Actionable {
	return WrapActionable(gextras.InternObject(i))
}

func (i imageMenuItem) AsActivatable() Activatable {
	return WrapActivatable(gextras.InternObject(i))
}

func (i imageMenuItem) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(i))
}
