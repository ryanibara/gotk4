// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_expander_get_type()), F: marshalTreeExpander},
	})
}

// TreeExpander: `GtkTreeExpander` is a widget that provides an expander for a
// list.
//
// It is typically placed as a bottommost child into a `GtkListView` to allow
// users to expand and collapse children in a list with a
// [class@Gtk.TreeListModel]. `GtkTreeExpander` provides the common UI elements,
// gestures and keybindings for this purpose.
//
// On top of this, the "listitem.expand", "listitem.collapse" and
// "listitem.toggle-expand" actions are provided to allow adding custom UI for
// managing expanded state.
//
// The `GtkTreeListModel` must be set to not be passthrough. Then it will
// provide [class@Gtk.TreeListRow] items which can be set via
// [method@Gtk.TreeExpander.set_list_row] on the expander. The expander will
// then watch that row item automatically. [method@Gtk.TreeExpander.set_child]
// sets the widget that displays the actual row contents.
//
//
// CSS nodes
//
// “` treeexpander ├── [indent]* ├── [expander] ╰── <child> “`
//
// `GtkTreeExpander` has zero or one CSS nodes with the name "expander" that
// should display the expander icon. The node will be `:checked` when it is
// expanded. If the node is not expandable, an "indent" node will be displayed
// instead.
//
// For every level of depth, another "indent" node is prepended.
//
//
// Accessibility
//
// `GtkTreeExpander` uses the GTK_ACCESSIBLE_ROLE_GROUP role. The expander icon
// is represented as a GTK_ACCESSIBLE_ROLE_BUTTON, labelled by the expander's
// child, and toggling it will change the GTK_ACCESSIBLE_STATE_EXPANDED state.
type TreeExpander interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Child gets the child widget displayed by @self.
	Child() Widget
	// Item forwards the item set on the `GtkTreeListRow` that @self is
	// managing.
	//
	// This call is essentially equivalent to calling:
	//
	// “`c gtk_tree_list_row_get_item (gtk_tree_expander_get_list_row (@self));
	// “`
	Item() gextras.Objector
	// ListRow gets the list row managed by @self.
	ListRow() TreeListRow
	// SetChild sets the content widget to display.
	SetChild(child Widget)
	// SetListRow sets the tree list row that this expander should manage.
	SetListRow(listRow TreeListRow)
}

// treeExpander implements the TreeExpander interface.
type treeExpander struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ TreeExpander = (*treeExpander)(nil)

// WrapTreeExpander wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeExpander(obj *externglib.Object) TreeExpander {
	return TreeExpander{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalTreeExpander(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeExpander(obj), nil
}

// NewTreeExpander constructs a class TreeExpander.
func NewTreeExpander() TreeExpander {
	ret := C.gtk_tree_expander_new()

	var ret0 TreeExpander

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(TreeExpander)

	return ret0
}

// Child gets the child widget displayed by @self.
func (s treeExpander) Child() Widget {
	var arg0 *C.GtkTreeExpander

	arg0 = (*C.GtkTreeExpander)(s.Native())

	ret := C.gtk_tree_expander_get_child(arg0)

	var ret0 Widget

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(Widget)

	return ret0
}

// Item forwards the item set on the `GtkTreeListRow` that @self is
// managing.
//
// This call is essentially equivalent to calling:
//
// “`c gtk_tree_list_row_get_item (gtk_tree_expander_get_list_row (@self));
// “`
func (s treeExpander) Item() gextras.Objector {
	var arg0 *C.GtkTreeExpander

	arg0 = (*C.GtkTreeExpander)(s.Native())

	ret := C.gtk_tree_expander_get_item(arg0)

	var ret0 gextras.Objector

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(gextras.Objector)

	return ret0
}

// ListRow gets the list row managed by @self.
func (s treeExpander) ListRow() TreeListRow {
	var arg0 *C.GtkTreeExpander

	arg0 = (*C.GtkTreeExpander)(s.Native())

	ret := C.gtk_tree_expander_get_list_row(arg0)

	var ret0 TreeListRow

	ret0 = gextras.CastObject(externglib.Take(unsafe.Pointer(ret.Native()))).(TreeListRow)

	return ret0
}

// SetChild sets the content widget to display.
func (s treeExpander) SetChild(child Widget) {
	var arg0 *C.GtkTreeExpander
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkTreeExpander)(s.Native())
	arg1 = (*C.GtkWidget)(child.Native())

	C.gtk_tree_expander_set_child(arg0, arg1)
}

// SetListRow sets the tree list row that this expander should manage.
func (s treeExpander) SetListRow(listRow TreeListRow) {
	var arg0 *C.GtkTreeExpander
	var arg1 *C.GtkTreeListRow

	arg0 = (*C.GtkTreeExpander)(s.Native())
	arg1 = (*C.GtkTreeListRow)(listRow.Native())

	C.gtk_tree_expander_set_list_row(arg0, arg1)
}
