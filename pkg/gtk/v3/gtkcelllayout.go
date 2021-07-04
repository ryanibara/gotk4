// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_cell_layout_get_type()), F: marshalCellLayout},
	})
}

// CellLayoutDataFunc: a function which should set the value of @cell_layout’s
// cell renderer(s) as appropriate.
type CellLayoutDataFunc func(cellLayout CellLayout, cell CellRenderer, treeModel TreeModel, iter *TreeIter)

//export gotk4_CellLayoutDataFunc
func _CellLayoutDataFunc(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var cellLayout CellLayout // out
	var cell CellRenderer     // out
	var treeModel TreeModel   // out
	var iter *TreeIter        // out

	cellLayout = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(CellLayout)
	cell = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(CellRenderer)
	treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(arg2))).(TreeModel)
	iter = (*TreeIter)(unsafe.Pointer(arg3))

	fn := v.(CellLayoutDataFunc)
	fn(cellLayout, cell, treeModel, iter)
}

// CellLayout is an interface to be implemented by all objects which want to
// provide a TreeViewColumn like API for packing cells, setting attributes and
// data funcs.
//
// One of the notable features provided by implementations of GtkCellLayout are
// attributes. Attributes let you set the properties in flexible ways. They can
// just be set to constant values like regular properties. But they can also be
// mapped to a column of the underlying tree model with
// gtk_cell_layout_set_attributes(), which means that the value of the attribute
// can change from cell to cell as they are rendered by the cell renderer.
// Finally, it is possible to specify a function with
// gtk_cell_layout_set_cell_data_func() that is called to determine the value of
// the attribute for each cell that is rendered.
//
//
// GtkCellLayouts as GtkBuildable
//
// Implementations of GtkCellLayout which also implement the GtkBuildable
// interface (CellView, IconView, ComboBox, EntryCompletion, TreeViewColumn)
// accept GtkCellRenderer objects as <child> elements in UI definitions. They
// support a custom <attributes> element for their children, which can contain
// multiple <attribute> elements. Each <attribute> element has a name attribute
// which specifies a property of the cell renderer; the content of the element
// is the attribute value.
//
// This is an example of a UI definition fragment specifying attributes:
//
//    <object class="GtkCellView">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <attributes>
//          <attribute name="text">0</attribute>
//        </attributes>
//      </child>"
//    </object>
//
// Furthermore for implementations of GtkCellLayout that use a CellArea to lay
// out cells (all GtkCellLayouts in GTK+ use a GtkCellArea) [cell
// properties][cell-properties] can also be defined in the format by specifying
// the custom <cell-packing> attribute which can contain multiple <property>
// elements defined in the normal way.
//
// Here is a UI definition fragment specifying cell properties:
//
//    <object class="GtkTreeViewColumn">
//      <child>
//        <object class="GtkCellRendererText"/>
//        <cell-packing>
//          <property name="align">True</property>
//          <property name="expand">False</property>
//        </cell-packing>
//      </child>"
//    </object>
//
//
// Subclassing GtkCellLayout implementations
//
// When subclassing a widget that implements CellLayout like IconView or
// ComboBox, there are some considerations related to the fact that these
// widgets internally use a CellArea. The cell area is exposed as a
// construct-only property by these widgets. This means that it is possible to
// e.g. do
//
//    static void
//    my_combo_box_init (MyComboBox *b)
//    {
//      GtkCellRenderer *cell;
//
//      cell = gtk_cell_renderer_pixbuf_new ();
//      // The following call causes the default cell area for combo boxes,
//      // a GtkCellAreaBox, to be instantiated
//      gtk_cell_layout_pack_start (GTK_CELL_LAYOUT (b), cell, FALSE);
//      ...
//    }
//
//    GtkWidget *
//    my_combo_box_new (GtkCellArea *area)
//    {
//      // This call is going to cause a warning about area being ignored
//      return g_object_new (MY_TYPE_COMBO_BOX, "cell-area", area, NULL);
//    }
//
// If supporting alternative cell areas with your derived widget is not
// important, then this does not have to concern you. If you want to support
// alternative cell areas, you can do so by moving the problematic calls out of
// init() and into a constructor() for your class.
type CellLayout interface {
	gextras.Objector

	// AddAttribute sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	AddAttribute(cell CellRenderer, attribute string, column int)
	// Clear sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	Clear()
	// ClearAttributes sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	ClearAttributes(cell CellRenderer)
	// Area sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	Area() CellArea
	// PackEnd sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	PackEnd(cell CellRenderer, expand bool)
	// PackStart sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	PackStart(cell CellRenderer, expand bool)
	// Reorder sets the CellLayoutDataFunc to use for @cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of @cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// @func may be nil to remove a previously set function.
	Reorder(cell CellRenderer, position int)
}

// cellLayout implements the CellLayout interface.
type cellLayout struct {
	gextras.Objector
}

var _ CellLayout = (*cellLayout)(nil)

// WrapCellLayout wraps a GObject to a type that implements
// interface CellLayout. It is primarily used internally.
func WrapCellLayout(obj *externglib.Object) CellLayout {
	return cellLayout{
		Objector: obj,
	}
}

func marshalCellLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellLayout(obj), nil
}

func (c cellLayout) AddAttribute(cell CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.gchar           // out
	var _arg3 C.gint             // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(column)

	C.gtk_cell_layout_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

func (c cellLayout) Clear() {
	var _arg0 *C.GtkCellLayout // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	C.gtk_cell_layout_clear(_arg0)
}

func (c cellLayout) ClearAttributes(cell CellRenderer) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_layout_clear_attributes(_arg0, _arg1)
}

func (c cellLayout) Area() CellArea {
	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GtkCellArea   // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_cell_layout_get_area(_arg0)

	var _cellArea CellArea // out

	_cellArea = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(CellArea)

	return _cellArea
}

func (c cellLayout) PackEnd(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_end(_arg0, _arg1, _arg2)
}

func (c cellLayout) PackStart(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_start(_arg0, _arg1, _arg2)
}

func (c cellLayout) Reorder(cell CellRenderer, position int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gint             // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = C.gint(position)

	C.gtk_cell_layout_reorder(_arg0, _arg1, _arg2)
}
