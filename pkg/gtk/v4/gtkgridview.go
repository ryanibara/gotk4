// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_GridView_ConnectActivate(gpointer, guint, guintptr);
import "C"

// glib.Type values for gtkgridview.go.
var GTypeGridView = externglib.Type(C.gtk_grid_view_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeGridView, F: marshalGridView},
	})
}

// GridViewOverrider contains methods that are overridable.
type GridViewOverrider interface {
}

// GridView: GtkGridView presents a large dynamic grid of items.
//
// GtkGridView uses its factory to generate one child widget for each visible
// item and shows them in a grid. The orientation of the grid view determines if
// the grid reflows vertically or horizontally.
//
// GtkGridView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// gtk.GridView:enable-rubberband.
//
// To learn more about the list widget framework, see the overview
// (section-list-widget.html).
//
// CSS nodes
//
//    gridview
//    ├── child
//    │
//    ├── child
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkGridView uses a single CSS node with name gridview. Each child uses a
// single CSS node with name child. For rubberband selection, a subnode with
// name rubberband is used.
//
//
// Accessibility
//
// GtkGridView uses the GTK_ACCESSIBLE_ROLE_GRID role, and the items use the
// GTK_ACCESSIBLE_ROLE_GRID_CELL role.
type GridView struct {
	_ [0]func() // equal guard
	ListBase
}

var (
	_ ListBaser = (*GridView)(nil)
)

func classInitGridViewer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGridView(obj *externglib.Object) *GridView {
	return &GridView{
		ListBase: ListBase{
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
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
			Scrollable: Scrollable{
				Object: obj,
			},
		},
	}
}

func marshalGridView(p uintptr) (interface{}, error) {
	return wrapGridView(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_GridView_ConnectActivate
func _gotk4_gtk4_GridView_ConnectActivate(arg0 C.gpointer, arg1 C.guint, arg2 C.guintptr) {
	var f func(position uint)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(position uint))
	}

	var _position uint // out

	_position = uint(arg1)

	f(_position)
}

// ConnectActivate: emitted when a cell has been activated by the user, usually
// via activating the GtkGridView|list.activate-item action.
//
// This allows for a convenient way to handle activation in a gridview. See
// gtk.ListItem:activatable for details on how to use this signal.
func (self *GridView) ConnectActivate(f func(position uint)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(self, "activate", false, unsafe.Pointer(C._gotk4_gtk4_GridView_ConnectActivate), f)
}

// NewGridView creates a new GtkGridView that uses the given factory for mapping
// items to widgets.
//
// The function takes ownership of the arguments, so you can write code like
//
//    grid_view = gtk_grid_view_new (create_model (),
//      gtk_builder_list_item_factory_new_from_resource ("/resource.ui"));.
//
// The function takes the following parameters:
//
//    - model (optional) to use, or NULL.
//    - factory (optional) to populate items with, or NULL.
//
// The function returns the following values:
//
//    - gridView: new GtkGridView using the given model and factory.
//
func NewGridView(model SelectionModeller, factory *ListItemFactory) *GridView {
	var _arg1 *C.GtkSelectionModel  // out
	var _arg2 *C.GtkListItemFactory // out
	var _cret *C.GtkWidget          // in

	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(model).Native()))
	}
	if factory != nil {
		_arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(externglib.InternObject(factory).Native()))
		C.g_object_ref(C.gpointer(externglib.InternObject(factory).Native()))
	}

	_cret = C.gtk_grid_view_new(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(factory)

	var _gridView *GridView // out

	_gridView = wrapGridView(externglib.Take(unsafe.Pointer(_cret)))

	return _gridView
}

// EnableRubberband returns whether rows can be selected by dragging with the
// mouse.
//
// The function returns the following values:
//
//    - ok: TRUE if rubberband selection is enabled.
//
func (self *GridView) EnableRubberband() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_enable_rubberband(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Factory gets the factory that's currently used to populate list items.
//
// The function returns the following values:
//
//    - listItemFactory (optional): factory in use.
//
func (self *GridView) Factory() *ListItemFactory {
	var _arg0 *C.GtkGridView        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_factory(_arg0)
	runtime.KeepAlive(self)

	var _listItemFactory *ListItemFactory // out

	if _cret != nil {
		_listItemFactory = wrapListItemFactory(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _listItemFactory
}

// MaxColumns gets the maximum number of columns that the grid will use.
//
// The function returns the following values:
//
//    - guint: maximum number of columns.
//
func (self *GridView) MaxColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_max_columns(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MinColumns gets the minimum number of columns that the grid will use.
//
// The function returns the following values:
//
//    - guint: minimum number of columns.
//
func (self *GridView) MinColumns() uint {
	var _arg0 *C.GtkGridView // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_min_columns(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Model gets the model that's currently used to read the items displayed.
//
// The function returns the following values:
//
//    - selectionModel (optional): model in use.
//
func (self *GridView) Model() SelectionModeller {
	var _arg0 *C.GtkGridView       // out
	var _cret *C.GtkSelectionModel // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_model(_arg0)
	runtime.KeepAlive(self)

	var _selectionModel SelectionModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(SelectionModeller)
				return ok
			})
			rv, ok := casted.(SelectionModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.SelectionModeller")
			}
			_selectionModel = rv
		}
	}

	return _selectionModel
}

// SingleClickActivate returns whether items will be activated on single click
// and selected on hover.
//
// The function returns the following values:
//
//    - ok: TRUE if items are activated on single click.
//
func (self *GridView) SingleClickActivate() bool {
	var _arg0 *C.GtkGridView // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.gtk_grid_view_get_single_click_activate(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetEnableRubberband sets whether selections can be changed by dragging with
// the mouse.
//
// The function takes the following parameters:
//
//    - enableRubberband: TRUE to enable rubberband selection.
//
func (self *GridView) SetEnableRubberband(enableRubberband bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if enableRubberband {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_enable_rubberband(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(enableRubberband)
}

// SetFactory sets the GtkListItemFactory to use for populating list items.
//
// The function takes the following parameters:
//
//    - factory (optional) to use or NULL for none.
//
func (self *GridView) SetFactory(factory *ListItemFactory) {
	var _arg0 *C.GtkGridView        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if factory != nil {
		_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(externglib.InternObject(factory).Native()))
	}

	C.gtk_grid_view_set_factory(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(factory)
}

// SetMaxColumns sets the maximum number of columns to use.
//
// This number must be at least 1.
//
// If max_columns is smaller than the minimum set via
// gtk.GridView.SetMinColumns(), that value is used instead.
//
// The function takes the following parameters:
//
//    - maxColumns: maximum number of columns.
//
func (self *GridView) SetMaxColumns(maxColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(maxColumns)

	C.gtk_grid_view_set_max_columns(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(maxColumns)
}

// SetMinColumns sets the minimum number of columns to use.
//
// This number must be at least 1.
//
// If min_columns is smaller than the minimum set via
// gtk.GridView.SetMaxColumns(), that value is ignored.
//
// The function takes the following parameters:
//
//    - minColumns: minimum number of columns.
//
func (self *GridView) SetMinColumns(minColumns uint) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	_arg1 = C.guint(minColumns)

	C.gtk_grid_view_set_min_columns(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(minColumns)
}

// SetModel sets the imodel to use.
//
// This must be a gtk.SelectionModel.
//
// The function takes the following parameters:
//
//    - model (optional) to use or NULL for none.
//
func (self *GridView) SetModel(model SelectionModeller) {
	var _arg0 *C.GtkGridView       // out
	var _arg1 *C.GtkSelectionModel // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if model != nil {
		_arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(externglib.InternObject(model).Native()))
	}

	C.gtk_grid_view_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSingleClickActivate sets whether items should be activated on single click
// and selected on hover.
//
// The function takes the following parameters:
//
//    - singleClickActivate: TRUE to activate items on single click.
//
func (self *GridView) SetSingleClickActivate(singleClickActivate bool) {
	var _arg0 *C.GtkGridView // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkGridView)(unsafe.Pointer(externglib.InternObject(self).Native()))
	if singleClickActivate {
		_arg1 = C.TRUE
	}

	C.gtk_grid_view_set_single_click_activate(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(singleClickActivate)
}
