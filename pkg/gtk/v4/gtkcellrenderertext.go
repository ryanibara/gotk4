// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_CellRendererText_ConnectEdited(gpointer, gchar*, gchar*, guintptr);
// extern void _gotk4_gtk4_CellRendererTextClass_edited(GtkCellRendererText*, char*, char*);
// void _gotk4_gtk4_CellRendererText_virtual_edited(void* fnptr, GtkCellRendererText* arg0, char* arg1, char* arg2) {
//   ((void (*)(GtkCellRendererText*, char*, char*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeCellRendererText = coreglib.Type(C.gtk_cell_renderer_text_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererText, F: marshalCellRendererText},
	})
}

// CellRendererTextOverrides contains methods that are overridable.
type CellRendererTextOverrides struct {
	// The function takes the following parameters:
	//
	//    - path
	//    - newText
	//
	Edited func(path, newText string)
}

func defaultCellRendererTextOverrides(v *CellRendererText) CellRendererTextOverrides {
	return CellRendererTextOverrides{
		Edited: v.edited,
	}
}

// CellRendererText renders text in a cell
//
// A CellRendererText renders a given text in its cell, using the font, color
// and style information provided by its properties. The text will be ellipsized
// if it is too long and the CellRendererText:ellipsize property allows it.
//
// If the CellRenderer:mode is GTK_CELL_RENDERER_MODE_EDITABLE, the
// CellRendererText allows to edit its text using an entry.
type CellRendererText struct {
	_ [0]func() // equal guard
	CellRenderer
}

var (
	_ CellRendererer = (*CellRendererText)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CellRendererText, *CellRendererTextClass, CellRendererTextOverrides](
		GTypeCellRendererText,
		initCellRendererTextClass,
		wrapCellRendererText,
		defaultCellRendererTextOverrides,
	)
}

func initCellRendererTextClass(gclass unsafe.Pointer, overrides CellRendererTextOverrides, classInitFunc func(*CellRendererTextClass)) {
	pclass := (*C.GtkCellRendererTextClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeCellRendererText))))

	if overrides.Edited != nil {
		pclass.edited = (*[0]byte)(C._gotk4_gtk4_CellRendererTextClass_edited)
	}

	if classInitFunc != nil {
		class := (*CellRendererTextClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellRendererText(obj *coreglib.Object) *CellRendererText {
	return &CellRendererText{
		CellRenderer: CellRenderer{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalCellRendererText(p uintptr) (interface{}, error) {
	return wrapCellRendererText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEdited: this signal is emitted after renderer has been edited.
//
// It is the responsibility of the application to update the model and store
// new_text at the position indicated by path.
func (renderer *CellRendererText) ConnectEdited(f func(path, newText string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(renderer, "edited", false, unsafe.Pointer(C._gotk4_gtk4_CellRendererText_ConnectEdited), f)
}

// NewCellRendererText creates a new CellRendererText. Adjust how text is drawn
// using object properties. Object properties can be set globally (with
// g_object_set()). Also, with TreeViewColumn, you can bind a property to a
// value in a TreeModel. For example, you can bind the “text” property on the
// cell renderer to a string value in the model, thus rendering a different
// string in each row of the TreeView.
//
// The function returns the following values:
//
//    - cellRendererText: new cell renderer.
//
func NewCellRendererText() *CellRendererText {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_text_new()

	var _cellRendererText *CellRendererText // out

	_cellRendererText = wrapCellRendererText(coreglib.Take(unsafe.Pointer(_cret)))

	return _cellRendererText
}

// SetFixedHeightFromFont sets the height of a renderer to explicitly be
// determined by the “font” and “y_pad” property set on it. Further changes in
// these properties do not affect the height, so they must be accompanied by a
// subsequent call to this function. Using this function is inflexible, and
// should really only be used if calculating the size of a cell is too slow (ie,
// a massive number of cells displayed). If number_of_rows is -1, then the fixed
// height is unset, and the height is determined by the properties again.
//
// The function takes the following parameters:
//
//    - numberOfRows: number of rows of text each cell renderer is allocated, or
//      -1.
//
func (renderer *CellRendererText) SetFixedHeightFromFont(numberOfRows int) {
	var _arg0 *C.GtkCellRendererText // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(coreglib.InternObject(renderer).Native()))
	_arg1 = C.int(numberOfRows)

	C.gtk_cell_renderer_text_set_fixed_height_from_font(_arg0, _arg1)
	runtime.KeepAlive(renderer)
	runtime.KeepAlive(numberOfRows)
}

// The function takes the following parameters:
//
//    - path
//    - newText
//
func (cellRendererText *CellRendererText) edited(path, newText string) {
	gclass := (*C.GtkCellRendererTextClass)(coreglib.PeekParentClass(cellRendererText))
	fnarg := gclass.edited

	var _arg0 *C.GtkCellRendererText // out
	var _arg1 *C.char                // out
	var _arg2 *C.char                // out

	_arg0 = (*C.GtkCellRendererText)(unsafe.Pointer(coreglib.InternObject(cellRendererText).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(newText)))
	defer C.free(unsafe.Pointer(_arg2))

	C._gotk4_gtk4_CellRendererText_virtual_edited(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(cellRendererText)
	runtime.KeepAlive(path)
	runtime.KeepAlive(newText)
}

// CellRendererTextClass: instance of this type is always passed by reference.
type CellRendererTextClass struct {
	*cellRendererTextClass
}

// cellRendererTextClass is the struct that's finalized.
type cellRendererTextClass struct {
	native *C.GtkCellRendererTextClass
}

func (c *CellRendererTextClass) ParentClass() *CellRendererClass {
	valptr := &c.native.parent_class
	var _v *CellRendererClass // out
	_v = (*CellRendererClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
