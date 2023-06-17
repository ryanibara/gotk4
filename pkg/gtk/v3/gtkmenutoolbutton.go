// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
// extern void _gotk4_gtk3_MenuToolButton_ConnectShowMenu(gpointer, guintptr);
// extern void _gotk4_gtk3_MenuToolButtonClass_show_menu(GtkMenuToolButton*);
// void _gotk4_gtk3_MenuToolButton_virtual_show_menu(void* fnptr, GtkMenuToolButton* arg0) {
//   ((void (*)(GtkMenuToolButton*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeMenuToolButton = coreglib.Type(C.gtk_menu_tool_button_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeMenuToolButton, F: marshalMenuToolButton},
	})
}

// MenuToolButtonOverrides contains methods that are overridable.
type MenuToolButtonOverrides struct {
	ShowMenu func()
}

func defaultMenuToolButtonOverrides(v *MenuToolButton) MenuToolButtonOverrides {
	return MenuToolButtonOverrides{
		ShowMenu: v.showMenu,
	}
}

// MenuToolButton is a ToolItem that contains a button and a small additional
// button with an arrow. When clicked, the arrow button pops up a dropdown menu.
//
// Use gtk_menu_tool_button_new() to create a new MenuToolButton.
//
// # GtkMenuToolButton as GtkBuildable
//
// The GtkMenuToolButton implementation of the GtkBuildable interface supports
// adding a menu by specifying “menu” as the “type” attribute of a <child>
// element.
//
// An example for a UI definition fragment with menus:
//
//    <object class="GtkMenuToolButton">
//      <child type="menu">
//        <object class="GtkMenu"/>
//      </child>
//    </object>.
type MenuToolButton struct {
	_ [0]func() // equal guard
	ToolButton
}

var (
	_ coreglib.Objector = (*MenuToolButton)(nil)
	_ Binner            = (*MenuToolButton)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*MenuToolButton, *MenuToolButtonClass, MenuToolButtonOverrides](
		GTypeMenuToolButton,
		initMenuToolButtonClass,
		wrapMenuToolButton,
		defaultMenuToolButtonOverrides,
	)
}

func initMenuToolButtonClass(gclass unsafe.Pointer, overrides MenuToolButtonOverrides, classInitFunc func(*MenuToolButtonClass)) {
	pclass := (*C.GtkMenuToolButtonClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeMenuToolButton))))

	if overrides.ShowMenu != nil {
		pclass.show_menu = (*[0]byte)(C._gotk4_gtk3_MenuToolButtonClass_show_menu)
	}

	if classInitFunc != nil {
		class := (*MenuToolButtonClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapMenuToolButton(obj *coreglib.Object) *MenuToolButton {
	return &MenuToolButton{
		ToolButton: ToolButton{
			ToolItem: ToolItem{
				Bin: Bin{
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
				},
				Object: obj,
				Activatable: Activatable{
					Object: obj,
				},
			},
			Object: obj,
			Actionable: Actionable{
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
		},
	}
}

func marshalMenuToolButton(p uintptr) (interface{}, error) {
	return wrapMenuToolButton(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectShowMenu signal is emitted before the menu is shown.
//
// It can be used to populate the menu on demand, using
// gtk_menu_tool_button_set_menu().
//
// Note that even if you populate the menu dynamically in this way, you must
// set an empty menu on the MenuToolButton beforehand, since the arrow is made
// insensitive if the menu is not set.
func (button *MenuToolButton) ConnectShowMenu(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(button, "show-menu", false, unsafe.Pointer(C._gotk4_gtk3_MenuToolButton_ConnectShowMenu), f)
}

// NewMenuToolButton creates a new MenuToolButton using icon_widget as icon and
// label as label.
//
// The function takes the following parameters:
//
//   - iconWidget (optional): widget that will be used as icon widget, or NULL.
//   - label (optional): string that will be used as label, or NULL.
//
// The function returns the following values:
//
//   - menuToolButton: new MenuToolButton.
//
func NewMenuToolButton(iconWidget Widgetter, label string) *MenuToolButton {
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	if iconWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(iconWidget).Native()))
	}
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_menu_tool_button_new(_arg1, _arg2)
	runtime.KeepAlive(iconWidget)
	runtime.KeepAlive(label)

	var _menuToolButton *MenuToolButton // out

	_menuToolButton = wrapMenuToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuToolButton
}

// NewMenuToolButtonFromStock creates a new MenuToolButton. The new
// MenuToolButton will contain an icon and label from the stock item indicated
// by stock_id.
//
// Deprecated: Use gtk_menu_tool_button_new() instead.
//
// The function takes the following parameters:
//
//   - stockId: name of a stock item.
//
// The function returns the following values:
//
//   - menuToolButton: new MenuToolButton.
//
func NewMenuToolButtonFromStock(stockId string) *MenuToolButton {
	var _arg1 *C.gchar       // out
	var _cret *C.GtkToolItem // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_menu_tool_button_new_from_stock(_arg1)
	runtime.KeepAlive(stockId)

	var _menuToolButton *MenuToolButton // out

	_menuToolButton = wrapMenuToolButton(coreglib.Take(unsafe.Pointer(_cret)))

	return _menuToolButton
}

// Menu gets the Menu associated with MenuToolButton.
//
// The function returns the following values:
//
//   - widget associated with MenuToolButton.
//
func (button *MenuToolButton) Menu() Widgetter {
	var _arg0 *C.GtkMenuToolButton // out
	var _cret *C.GtkWidget         // in

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	_cret = C.gtk_menu_tool_button_get_menu(_arg0)
	runtime.KeepAlive(button)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

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

	return _widget
}

// SetArrowTooltipMarkup sets the tooltip markup text to be used as tooltip for
// the arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text()
// for setting a tooltip on the whole MenuToolButton.
//
// The function takes the following parameters:
//
//   - markup text to be used as tooltip text for button’s arrow button.
//
func (button *MenuToolButton) SetArrowTooltipMarkup(markup string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_markup(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(markup)
}

// SetArrowTooltipText sets the tooltip text to be used as tooltip for the
// arrow button which pops up the menu. See gtk_tool_item_set_tooltip_text() for
// setting a tooltip on the whole MenuToolButton.
//
// The function takes the following parameters:
//
//   - text to be used as tooltip text for button’s arrow button.
//
func (button *MenuToolButton) SetArrowTooltipText(text string) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_menu_tool_button_set_arrow_tooltip_text(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(text)
}

// SetMenu sets the Menu that is popped up when the user clicks on the arrow.
// If menu is NULL, the arrow button becomes insensitive.
//
// The function takes the following parameters:
//
//   - menu associated with MenuToolButton.
//
func (button *MenuToolButton) SetMenu(menu Widgetter) {
	var _arg0 *C.GtkMenuToolButton // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(menu).Native()))

	C.gtk_menu_tool_button_set_menu(_arg0, _arg1)
	runtime.KeepAlive(button)
	runtime.KeepAlive(menu)
}

func (button *MenuToolButton) showMenu() {
	gclass := (*C.GtkMenuToolButtonClass)(coreglib.PeekParentClass(button))
	fnarg := gclass.show_menu

	var _arg0 *C.GtkMenuToolButton // out

	_arg0 = (*C.GtkMenuToolButton)(unsafe.Pointer(coreglib.InternObject(button).Native()))

	C._gotk4_gtk3_MenuToolButton_virtual_show_menu(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(button)
}

// MenuToolButtonClass: instance of this type is always passed by reference.
type MenuToolButtonClass struct {
	*menuToolButtonClass
}

// menuToolButtonClass is the struct that's finalized.
type menuToolButtonClass struct {
	native *C.GtkMenuToolButtonClass
}

// ParentClass: parent class.
func (m *MenuToolButtonClass) ParentClass() *ToolButtonClass {
	valptr := &m.native.parent_class
	var _v *ToolButtonClass // out
	_v = (*ToolButtonClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
