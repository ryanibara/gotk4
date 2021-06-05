// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tool_item_group_get_type()), F: marshalToolItemGroup},
	})
}

// ToolItemGroup: a ToolItemGroup is used together with ToolPalette to add
// ToolItems to a palette like container with different categories and drag and
// drop support.
//
//
// CSS nodes
//
// GtkToolItemGroup has a single CSS node named toolitemgroup.
type ToolItemGroup interface {
	Container
	Buildable
	ToolShell

	// Collapsed gets whether @group is collapsed or expanded.
	Collapsed() bool
	// DropItem gets the tool item at position (x, y).
	DropItem(x int, y int) ToolItem
	// Ellipsize gets the ellipsization mode of @group.
	Ellipsize() pango.EllipsizeMode
	// HeaderRelief gets the relief mode of the header button of @group.
	HeaderRelief() ReliefStyle
	// ItemPosition gets the position of @item in @group as index.
	ItemPosition(item ToolItem) int
	// Label gets the label of @group.
	Label() string
	// LabelWidget gets the label widget of @group. See
	// gtk_tool_item_group_set_label_widget().
	LabelWidget() Widget
	// NItems gets the number of tool items in @group.
	NItems() uint
	// NthItem gets the tool item at @index in group.
	NthItem(index uint) ToolItem
	// Insert inserts @item at @position in the list of children of @group.
	Insert(item ToolItem, position int)
	// SetCollapsed sets whether the @group should be collapsed or expanded.
	SetCollapsed(collapsed bool)
	// SetEllipsize sets the ellipsization mode which should be used by labels
	// in @group.
	SetEllipsize(ellipsize pango.EllipsizeMode)
	// SetHeaderRelief: set the button relief of the group header. See
	// gtk_button_set_relief() for details.
	SetHeaderRelief(style ReliefStyle)
	// SetItemPosition sets the position of @item in the list of children of
	// @group.
	SetItemPosition(item ToolItem, position int)
	// SetLabel sets the label of the tool item group. The label is displayed in
	// the header of the group.
	SetLabel(label string)
	// SetLabelWidget sets the label of the tool item group. The label widget is
	// displayed in the header of the group, in place of the usual label.
	SetLabelWidget(labelWidget Widget)
}

// toolItemGroup implements the ToolItemGroup interface.
type toolItemGroup struct {
	Container
	Buildable
	ToolShell
}

var _ ToolItemGroup = (*toolItemGroup)(nil)

// WrapToolItemGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolItemGroup(obj *externglib.Object) ToolItemGroup {
	return ToolItemGroup{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
		ToolShell: WrapToolShell(obj),
	}
}

func marshalToolItemGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolItemGroup(obj), nil
}

// NewToolItemGroup constructs a class ToolItemGroup.
func NewToolItemGroup(label string) ToolItemGroup {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkToolItemGroup
	var ret1 ToolItemGroup

	cret = C.gtk_tool_item_group_new(label)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToolItemGroup)

	return ret1
}

// Collapsed gets whether @group is collapsed or expanded.
func (g toolItemGroup) Collapsed() bool {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_tool_item_group_get_collapsed(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// DropItem gets the tool item at position (x, y).
func (g toolItemGroup) DropItem(x int, y int) ToolItem {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = C.gint(x)
	arg2 = C.gint(y)

	var cret *C.GtkToolItem
	var ret1 ToolItem

	cret = C.gtk_tool_item_group_get_drop_item(arg0, x, y)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToolItem)

	return ret1
}

// Ellipsize gets the ellipsization mode of @group.
func (g toolItemGroup) Ellipsize() pango.EllipsizeMode {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret C.PangoEllipsizeMode
	var ret1 pango.EllipsizeMode

	cret = C.gtk_tool_item_group_get_ellipsize(arg0)

	ret1 = pango.EllipsizeMode(cret)

	return ret1
}

// HeaderRelief gets the relief mode of the header button of @group.
func (g toolItemGroup) HeaderRelief() ReliefStyle {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret C.GtkReliefStyle
	var ret1 ReliefStyle

	cret = C.gtk_tool_item_group_get_header_relief(arg0)

	ret1 = ReliefStyle(cret)

	return ret1
}

// ItemPosition gets the position of @item in @group as index.
func (g toolItemGroup) ItemPosition(item ToolItem) int {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_tool_item_group_get_item_position(arg0, item)

	ret1 = C.gint(cret)

	return ret1
}

// Label gets the label of @group.
func (g toolItemGroup) Label() string {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_tool_item_group_get_label(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// LabelWidget gets the label widget of @group. See
// gtk_tool_item_group_set_label_widget().
func (g toolItemGroup) LabelWidget() Widget {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_tool_item_group_get_label_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// NItems gets the number of tool items in @group.
func (g toolItemGroup) NItems() uint {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gtk_tool_item_group_get_n_items(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// NthItem gets the tool item at @index in group.
func (g toolItemGroup) NthItem(index uint) ToolItem {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.guint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = C.guint(index)

	var cret *C.GtkToolItem
	var ret1 ToolItem

	cret = C.gtk_tool_item_group_get_nth_item(arg0, index)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ToolItem)

	return ret1
}

// Insert inserts @item at @position in the list of children of @group.
func (g toolItemGroup) Insert(item ToolItem, position int) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	arg2 = C.gint(position)

	C.gtk_tool_item_group_insert(arg0, item, position)
}

// SetCollapsed sets whether the @group should be collapsed or expanded.
func (g toolItemGroup) SetCollapsed(collapsed bool) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.gboolean

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	if collapsed {
		arg1 = C.gboolean(1)
	}

	C.gtk_tool_item_group_set_collapsed(arg0, collapsed)
}

// SetEllipsize sets the ellipsization mode which should be used by labels
// in @group.
func (g toolItemGroup) SetEllipsize(ellipsize pango.EllipsizeMode) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.PangoEllipsizeMode

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (C.PangoEllipsizeMode)(ellipsize)

	C.gtk_tool_item_group_set_ellipsize(arg0, ellipsize)
}

// SetHeaderRelief: set the button relief of the group header. See
// gtk_button_set_relief() for details.
func (g toolItemGroup) SetHeaderRelief(style ReliefStyle) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.GtkReliefStyle

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (C.GtkReliefStyle)(style)

	C.gtk_tool_item_group_set_header_relief(arg0, style)
}

// SetItemPosition sets the position of @item in the list of children of
// @group.
func (g toolItemGroup) SetItemPosition(item ToolItem, position int) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	arg2 = C.gint(position)

	C.gtk_tool_item_group_set_item_position(arg0, item, position)
}

// SetLabel sets the label of the tool item group. The label is displayed in
// the header of the group.
func (g toolItemGroup) SetLabel(label string) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.gchar

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_item_group_set_label(arg0, label)
}

// SetLabelWidget sets the label of the tool item group. The label widget is
// displayed in the header of the group, in place of the usual label.
func (g toolItemGroup) SetLabelWidget(labelWidget Widget) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_item_group_set_label_widget(arg0, labelWidget)
}

type ToolItemGroupPrivate struct {
	native C.GtkToolItemGroupPrivate
}

// WrapToolItemGroupPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToolItemGroupPrivate(ptr unsafe.Pointer) *ToolItemGroupPrivate {
	if ptr == nil {
		return nil
	}

	return (*ToolItemGroupPrivate)(ptr)
}

func marshalToolItemGroupPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapToolItemGroupPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ToolItemGroupPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
