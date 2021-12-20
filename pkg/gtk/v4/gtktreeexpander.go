// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"runtime"
	"sync"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_expander_get_type()), F: marshalTreeExpanderer},
	})
}

// TreeExpander: GtkTreeExpander is a widget that provides an expander for a
// list.
//
// It is typically placed as a bottommost child into a GtkListView to allow
// users to expand and collapse children in a list with a gtk.TreeListModel.
// GtkTreeExpander provides the common UI elements, gestures and keybindings for
// this purpose.
//
// On top of this, the "listitem.expand", "listitem.collapse" and
// "listitem.toggle-expand" actions are provided to allow adding custom UI for
// managing expanded state.
//
// The GtkTreeListModel must be set to not be passthrough. Then it will provide
// gtk.TreeListRow items which can be set via gtk.TreeExpander.SetListRow() on
// the expander. The expander will then watch that row item automatically.
// gtk.TreeExpander.SetChild() sets the widget that displays the actual row
// contents.
//
// CSS nodes
//
//    treeexpander
//    ├── [indent]*
//    ├── [expander]
//    ╰── <child>
//
//
// GtkTreeExpander has zero or one CSS nodes with the name "expander" that
// should display the expander icon. The node will be :checked when it is
// expanded. If the node is not expandable, an "indent" node will be displayed
// instead.
//
// For every level of depth, another "indent" node is prepended.
//
//
// Accessibility
//
// GtkTreeExpander uses the GTK_ACCESSIBLE_ROLE_GROUP role. The expander icon is
// represented as a GTK_ACCESSIBLE_ROLE_BUTTON, labelled by the expander's
// child, and toggling it will change the GTK_ACCESSIBLE_STATE_EXPANDED state.
type TreeExpander struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*TreeExpander)(nil)
)

func wrapTreeExpander(obj *externglib.Object) *TreeExpander {
	return &TreeExpander{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalTreeExpanderer(p uintptr) (interface{}, error) {
	return wrapTreeExpander(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTreeExpander creates a new GtkTreeExpander.
//
// The function returns the following values:
//
//    - treeExpander: new GtkTreeExpander.
//
func NewTreeExpander() *TreeExpander {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_tree_expander_new()

	var _treeExpander *TreeExpander // out

	_treeExpander = wrapTreeExpander(externglib.Take(unsafe.Pointer(_cret)))

	return _treeExpander
}

// Child gets the child widget displayed by self.
//
// The function returns the following values:
//
//    - widget (optional): child displayed by self.
//
func (self *TreeExpander) Child() Widgetter {
	var _arg0 *C.GtkTreeExpander // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_child(_arg0)
	runtime.KeepAlive(self)

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

// Item forwards the item set on the GtkTreeListRow that self is managing.
//
// This call is essentially equivalent to calling:
//
//    gtk_tree_list_row_get_item (gtk_tree_expander_get_list_row (self));.
//
// The function returns the following values:
//
//    - object (optional): item of the row.
//
func (self *TreeExpander) Item() *externglib.Object {
	var _arg0 *C.GtkTreeExpander // out
	var _cret C.gpointer         // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_item(_arg0)
	runtime.KeepAlive(self)

	var _object *externglib.Object // out

	_object = externglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// ListRow gets the list row managed by self.
//
// The function returns the following values:
//
//    - treeListRow (optional): list row displayed by self.
//
func (self *TreeExpander) ListRow() *TreeListRow {
	var _arg0 *C.GtkTreeExpander // out
	var _cret *C.GtkTreeListRow  // in

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_tree_expander_get_list_row(_arg0)
	runtime.KeepAlive(self)

	var _treeListRow *TreeListRow // out

	if _cret != nil {
		_treeListRow = wrapTreeListRow(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _treeListRow
}

// SetChild sets the content widget to display.
//
// The function takes the following parameters:
//
//    - child (optional): GtkWidget, or NULL.
//
func (self *TreeExpander) SetChild(child Widgetter) {
	var _arg0 *C.GtkTreeExpander // out
	var _arg1 *C.GtkWidget       // out

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))
	if child != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	}

	C.gtk_tree_expander_set_child(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(child)
}

// SetListRow sets the tree list row that this expander should manage.
//
// The function takes the following parameters:
//
//    - listRow (optional): GtkTreeListRow, or NULL.
//
func (self *TreeExpander) SetListRow(listRow *TreeListRow) {
	var _arg0 *C.GtkTreeExpander // out
	var _arg1 *C.GtkTreeListRow  // out

	_arg0 = (*C.GtkTreeExpander)(unsafe.Pointer(self.Native()))
	if listRow != nil {
		_arg1 = (*C.GtkTreeListRow)(unsafe.Pointer(listRow.Native()))
	}

	C.gtk_tree_expander_set_list_row(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(listRow)
}
