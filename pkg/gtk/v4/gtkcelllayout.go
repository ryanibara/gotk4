// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// void _gotk4_gtk4_CellLayoutDataFunc(GtkCellLayout*, GtkCellRenderer*, GtkTreeModel*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_layout_get_type()), F: marshalCellLayouter},
	})
}

// CellLayoutDataFunc: function which should set the value of cell_layout’s cell
// renderer(s) as appropriate.
type CellLayoutDataFunc func(cellLayout CellLayouter, cell CellRendererer, treeModel TreeModeller, iter *TreeIter)

//export _gotk4_gtk4_CellLayoutDataFunc
func _gotk4_gtk4_CellLayoutDataFunc(arg0 *C.GtkCellLayout, arg1 *C.GtkCellRenderer, arg2 *C.GtkTreeModel, arg3 *C.GtkTreeIter, arg4 C.gpointer) {
	v := gbox.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var cellLayout CellLayouter // out
	var cell CellRendererer     // out
	var treeModel TreeModeller  // out
	var iter *TreeIter          // out

	{
		objptr := unsafe.Pointer(arg0)
		if objptr == nil {
			panic("object of type gtk.CellLayouter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(CellLayouter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.CellLayouter")
		}
		cellLayout = rv
	}
	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.CellRendererer is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(CellRendererer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.CellRendererer")
		}
		cell = rv
	}
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(TreeModeller)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.TreeModeller")
		}
		treeModel = rv
	}
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	fn := v.(CellLayoutDataFunc)
	fn(cellLayout, cell, treeModel, iter)
}

// CellLayoutOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type CellLayoutOverrider interface {
	// AddAttribute adds an attribute mapping to the list in cell_layout.
	//
	// The column is the column of the model to get a value from, and the
	// attribute is the parameter on cell to be set from the value. So for
	// example if column 2 of the model contains strings, you could have the
	// “text” attribute of a CellRendererText get its values from column 2.
	AddAttribute(cell CellRendererer, attribute string, column int)
	// Clear unsets all the mappings on all renderers on cell_layout and removes
	// all renderers from cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be cell_layout if called
	// on a CellArea or might be NULL if no CellArea is used by cell_layout.
	Area() CellAreaer
	// Cells returns the cell renderers which have been added to cell_layout.
	Cells() []CellRendererer
	// PackEnd adds the cell to the end of cell_layout. If expand is FALSE, then
	// the cell is allocated no more space than it needs. Any unused space is
	// divided evenly between cells for which expand is TRUE.
	//
	// Note that reusing the same cell renderer is not supported.
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the cell into the beginning of cell_layout. If expand is
	// FALSE, then the cell is allocated no more space than it needs. Any unused
	// space is divided evenly between cells for which expand is TRUE.
	//
	// Note that reusing the same cell renderer is not supported.
	PackStart(cell CellRendererer, expand bool)
	// Reorder re-inserts cell at position.
	//
	// Note that cell has already to be packed into cell_layout for this to
	// function properly.
	Reorder(cell CellRendererer, position int)
	// SetCellDataFunc sets the CellLayoutDataFunc to use for cell_layout.
	//
	// This function is used instead of the standard attributes mapping for
	// setting the column value, and should set the value of cell_layout’s cell
	// renderer(s) as appropriate.
	//
	// func may be NULL to remove a previously set function.
	SetCellDataFunc(cell CellRendererer, fn CellLayoutDataFunc)
}

// CellLayout: interface for packing cells
//
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
// out cells (all GtkCellLayouts in GTK use a GtkCellArea) [cell
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
type CellLayout struct {
	*externglib.Object
}

// CellLayouter describes CellLayout's abstract methods.
type CellLayouter interface {
	externglib.Objector

	// AddAttribute adds an attribute mapping to the list in cell_layout.
	AddAttribute(cell CellRendererer, attribute string, column int)
	// Clear unsets all the mappings on all renderers on cell_layout and removes
	// all renderers from cell_layout.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_cell_layout_set_attributes().
	ClearAttributes(cell CellRendererer)
	// Area returns the underlying CellArea which might be cell_layout if called
	// on a CellArea or might be NULL if no CellArea is used by cell_layout.
	Area() CellAreaer
	// Cells returns the cell renderers which have been added to cell_layout.
	Cells() []CellRendererer
	// PackEnd adds the cell to the end of cell_layout.
	PackEnd(cell CellRendererer, expand bool)
	// PackStart packs the cell into the beginning of cell_layout.
	PackStart(cell CellRendererer, expand bool)
	// Reorder re-inserts cell at position.
	Reorder(cell CellRendererer, position int)
	// SetCellDataFunc sets the CellLayoutDataFunc to use for cell_layout.
	SetCellDataFunc(cell CellRendererer, fn CellLayoutDataFunc)
}

var _ CellLayouter = (*CellLayout)(nil)

func wrapCellLayout(obj *externglib.Object) *CellLayout {
	return &CellLayout{
		Object: obj,
	}
}

func marshalCellLayouter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellLayout(obj), nil
}

// AddAttribute adds an attribute mapping to the list in cell_layout.
//
// The column is the column of the model to get a value from, and the attribute
// is the parameter on cell to be set from the value. So for example if column 2
// of the model contains strings, you could have the “text” attribute of a
// CellRendererText get its values from column 2.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - attribute on the renderer.
//    - column position on the model to get the attribute from.
//
func (cellLayout *CellLayout) AddAttribute(cell CellRendererer, attribute string, column int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out
	var _arg3 C.int              // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(attribute)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(column)

	C.gtk_cell_layout_add_attribute(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(attribute)
	runtime.KeepAlive(column)
}

// Clear unsets all the mappings on all renderers on cell_layout and removes all
// renderers from cell_layout.
func (cellLayout *CellLayout) Clear() {
	var _arg0 *C.GtkCellLayout // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))

	C.gtk_cell_layout_clear(_arg0)
	runtime.KeepAlive(cellLayout)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_cell_layout_set_attributes().
//
// The function takes the following parameters:
//
//    - cell to clear the attribute mapping on.
//
func (cellLayout *CellLayout) ClearAttributes(cell CellRendererer) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_layout_clear_attributes(_arg0, _arg1)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
}

// Area returns the underlying CellArea which might be cell_layout if called on
// a CellArea or might be NULL if no CellArea is used by cell_layout.
func (cellLayout *CellLayout) Area() CellAreaer {
	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GtkCellArea   // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))

	_cret = C.gtk_cell_layout_get_area(_arg0)
	runtime.KeepAlive(cellLayout)

	var _cellArea CellAreaer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(CellAreaer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.CellAreaer")
			}
			_cellArea = rv
		}
	}

	return _cellArea
}

// Cells returns the cell renderers which have been added to cell_layout.
func (cellLayout *CellLayout) Cells() []CellRendererer {
	var _arg0 *C.GtkCellLayout // out
	var _cret *C.GList         // in

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))

	_cret = C.gtk_cell_layout_get_cells(_arg0)
	runtime.KeepAlive(cellLayout)

	var _list []CellRendererer // out

	_list = make([]CellRendererer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkCellRenderer)(v)
		var dst CellRendererer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.CellRendererer is nil")
			}

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(CellRendererer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.CellRendererer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// PackEnd adds the cell to the end of cell_layout. If expand is FALSE, then the
// cell is allocated no more space than it needs. Any unused space is divided
// evenly between cells for which expand is TRUE.
//
// Note that reusing the same cell renderer is not supported.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - expand: TRUE if cell is to be given extra space allocated to
//    cell_layout.
//
func (cellLayout *CellLayout) PackEnd(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_end(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(expand)
}

// PackStart packs the cell into the beginning of cell_layout. If expand is
// FALSE, then the cell is allocated no more space than it needs. Any unused
// space is divided evenly between cells for which expand is TRUE.
//
// Note that reusing the same cell renderer is not supported.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - expand: TRUE if cell is to be given extra space allocated to
//    cell_layout.
//
func (cellLayout *CellLayout) PackStart(cell CellRendererer, expand bool) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.gboolean         // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_cell_layout_pack_start(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(expand)
}

// Reorder re-inserts cell at position.
//
// Note that cell has already to be packed into cell_layout for this to function
// properly.
//
// The function takes the following parameters:
//
//    - cell to reorder.
//    - position: new position to insert cell at.
//
func (cellLayout *CellLayout) Reorder(cell CellRendererer, position int) {
	var _arg0 *C.GtkCellLayout   // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.int              // out

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	_arg2 = C.int(position)

	C.gtk_cell_layout_reorder(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(position)
}

// SetCellDataFunc sets the CellLayoutDataFunc to use for cell_layout.
//
// This function is used instead of the standard attributes mapping for setting
// the column value, and should set the value of cell_layout’s cell renderer(s)
// as appropriate.
//
// func may be NULL to remove a previously set function.
//
// The function takes the following parameters:
//
//    - cell: CellRenderer.
//    - fn to use, or NULL.
//
func (cellLayout *CellLayout) SetCellDataFunc(cell CellRendererer, fn CellLayoutDataFunc) {
	var _arg0 *C.GtkCellLayout        // out
	var _arg1 *C.GtkCellRenderer      // out
	var _arg2 C.GtkCellLayoutDataFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify

	_arg0 = (*C.GtkCellLayout)(unsafe.Pointer(cellLayout.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if fn != nil {
		_arg2 = (*[0]byte)(C._gotk4_gtk4_CellLayoutDataFunc)
		_arg3 = C.gpointer(gbox.Assign(fn))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_cell_layout_set_cell_data_func(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(cellLayout)
	runtime.KeepAlive(cell)
	runtime.KeepAlive(fn)
}
