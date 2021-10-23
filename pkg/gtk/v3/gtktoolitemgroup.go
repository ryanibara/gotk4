// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
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
		{T: externglib.Type(C.gtk_tool_item_group_get_type()), F: marshalToolItemGrouper},
	})
}

// ToolItemGroup is used together with ToolPalette to add ToolItems to a palette
// like container with different categories and drag and drop support.
//
//
// CSS nodes
//
// GtkToolItemGroup has a single CSS node named toolitemgroup.
type ToolItemGroup struct {
	Container

	ToolShell
	*externglib.Object
}

func wrapToolItemGroup(obj *externglib.Object) *ToolItemGroup {
	return &ToolItemGroup{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
		ToolShell: ToolShell{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalToolItemGrouper(p uintptr) (interface{}, error) {
	return wrapToolItemGroup(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewToolItemGroup creates a new tool item group with label label.
//
// The function takes the following parameters:
//
//    - label of the new group.
//
func NewToolItemGroup(label string) *ToolItemGroup {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_tool_item_group_new(_arg1)
	runtime.KeepAlive(label)

	var _toolItemGroup *ToolItemGroup // out

	_toolItemGroup = wrapToolItemGroup(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItemGroup
}

// Collapsed gets whether group is collapsed or expanded.
func (group *ToolItemGroup) Collapsed() bool {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_collapsed(_arg0)
	runtime.KeepAlive(group)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DropItem gets the tool item at position (x, y).
//
// The function takes the following parameters:
//
//    - x position.
//    - y position.
//
func (group *ToolItemGroup) DropItem(x, y int) *ToolItem {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.gint              // out
	var _arg2 C.gint              // out
	var _cret *C.GtkToolItem      // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_tool_item_group_get_drop_item(_arg0, _arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// Ellipsize gets the ellipsization mode of group.
func (group *ToolItemGroup) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkToolItemGroup  // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_ellipsize(_arg0)
	runtime.KeepAlive(group)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// HeaderRelief gets the relief mode of the header button of group.
func (group *ToolItemGroup) HeaderRelief() ReliefStyle {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.GtkReliefStyle    // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_header_relief(_arg0)
	runtime.KeepAlive(group)

	var _reliefStyle ReliefStyle // out

	_reliefStyle = ReliefStyle(_cret)

	return _reliefStyle
}

// ItemPosition gets the position of item in group as index.
//
// The function takes the following parameters:
//
//    - item: ToolItem.
//
func (group *ToolItemGroup) ItemPosition(item *ToolItem) int {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _cret C.gint              // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	_cret = C.gtk_tool_item_group_get_item_position(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(item)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Label gets the label of group.
func (group *ToolItemGroup) Label() string {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret *C.gchar            // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_label(_arg0)
	runtime.KeepAlive(group)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// LabelWidget gets the label widget of group. See
// gtk_tool_item_group_set_label_widget().
func (group *ToolItemGroup) LabelWidget() Widgetter {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_label_widget(_arg0)
	runtime.KeepAlive(group)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// NItems gets the number of tool items in group.
func (group *ToolItemGroup) NItems() uint {
	var _arg0 *C.GtkToolItemGroup // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_tool_item_group_get_n_items(_arg0)
	runtime.KeepAlive(group)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// NthItem gets the tool item at index in group.
//
// The function takes the following parameters:
//
//    - index: index.
//
func (group *ToolItemGroup) NthItem(index uint) *ToolItem {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.guint             // out
	var _cret *C.GtkToolItem      // in

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.guint(index)

	_cret = C.gtk_tool_item_group_get_nth_item(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(index)

	var _toolItem *ToolItem // out

	_toolItem = wrapToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _toolItem
}

// Insert inserts item at position in the list of children of group.
//
// The function takes the following parameters:
//
//    - item to insert into group.
//    - position of item in group, starting with 0. The position -1 means end
//    of list.
//
func (group *ToolItemGroup) Insert(item *ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_insert(_arg0, _arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(item)
	runtime.KeepAlive(position)
}

// SetCollapsed sets whether the group should be collapsed or expanded.
//
// The function takes the following parameters:
//
//    - collapsed: whether the group should be collapsed or expanded.
//
func (group *ToolItemGroup) SetCollapsed(collapsed bool) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if collapsed {
		_arg1 = C.TRUE
	}

	C.gtk_tool_item_group_set_collapsed(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(collapsed)
}

// SetEllipsize sets the ellipsization mode which should be used by labels in
// group.
//
// The function takes the following parameters:
//
//    - ellipsize labels in group should use.
//
func (group *ToolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	var _arg0 *C.GtkToolItemGroup  // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.PangoEllipsizeMode(ellipsize)

	C.gtk_tool_item_group_set_ellipsize(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(ellipsize)
}

// SetHeaderRelief: set the button relief of the group header. See
// gtk_button_set_relief() for details.
//
// The function takes the following parameters:
//
//    - style: ReliefStyle.
//
func (group *ToolItemGroup) SetHeaderRelief(style ReliefStyle) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 C.GtkReliefStyle    // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = C.GtkReliefStyle(style)

	C.gtk_tool_item_group_set_header_relief(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(style)
}

// SetItemPosition sets the position of item in the list of children of group.
//
// The function takes the following parameters:
//
//    - item to move to a new position, should be a child of group.
//    - position: new position of item in group, starting with 0. The position
//    -1 means end of list.
//
func (group *ToolItemGroup) SetItemPosition(item *ToolItem, position int) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkToolItem      // out
	var _arg2 C.gint              // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	_arg2 = C.gint(position)

	C.gtk_tool_item_group_set_item_position(_arg0, _arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(item)
	runtime.KeepAlive(position)
}

// SetLabel sets the label of the tool item group. The label is displayed in the
// header of the group.
//
// The function takes the following parameters:
//
//    - label: new human-readable label of of the group.
//
func (group *ToolItemGroup) SetLabel(label string) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tool_item_group_set_label(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)
}

// SetLabelWidget sets the label of the tool item group. The label widget is
// displayed in the header of the group, in place of the usual label.
//
// The function takes the following parameters:
//
//    - labelWidget: widget to be displayed in place of the usual label.
//
func (group *ToolItemGroup) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkToolItemGroup // out
	var _arg1 *C.GtkWidget        // out

	_arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_item_group_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(group)
	runtime.KeepAlive(labelWidget)
}
