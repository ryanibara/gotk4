// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk4_CellAreaContextClass_allocate(void*, int, int);
// extern void _gotk4_gtk4_CellAreaContextClass_get_preferred_height_for_width(void*, int, void*, void*);
// extern void _gotk4_gtk4_CellAreaContextClass_get_preferred_width_for_height(void*, int, void*, void*);
// extern void _gotk4_gtk4_CellAreaContextClass_reset(void*);
import "C"

// glib.Type values for gtkcellareacontext.go.
var GTypeCellAreaContext = coreglib.Type(C.gtk_cell_area_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeCellAreaContext, F: marshalCellAreaContext},
	})
}

// CellAreaContextOverrider contains methods that are overridable.
type CellAreaContextOverrider interface {
	// Allocate allocates a width and/or a height for all rows which are to be
	// rendered with context.
	//
	// Usually allocation is performed only horizontally or sometimes vertically
	// since a group of rows are usually rendered side by side vertically or
	// horizontally and share either the same width or the same height.
	// Sometimes they are allocated in both horizontal and vertical orientations
	// producing a homogeneous effect of the rows. This is generally the case
	// for TreeView when TreeView:fixed-height-mode is enabled.
	//
	// The function takes the following parameters:
	//
	//    - width: allocated width for all TreeModel rows rendered with context,
	//      or -1.
	//    - height: allocated height for all TreeModel rows rendered with
	//      context, or -1.
	//
	Allocate(width, height int32)
	// PreferredHeightForWidth gets the accumulative preferred height for width
	// for all rows which have been requested for the same said width with this
	// context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	//
	// The function takes the following parameters:
	//
	//    - width: proposed width for allocation.
	//
	// The function returns the following values:
	//
	//    - minimumHeight (optional): location to store the minimum height, or
	//      NULL.
	//    - naturalHeight (optional): location to store the natural height, or
	//      NULL.
	//
	PreferredHeightForWidth(width int32) (minimumHeight, naturalHeight int32)
	// PreferredWidthForHeight gets the accumulative preferred width for height
	// for all rows which have been requested for the same said height with this
	// context.
	//
	// After gtk_cell_area_context_reset() is called and/or before ever
	// requesting the size of a CellArea, the returned values are -1.
	//
	// The function takes the following parameters:
	//
	//    - height: proposed height for allocation.
	//
	// The function returns the following values:
	//
	//    - minimumWidth (optional): location to store the minimum width, or
	//      NULL.
	//    - naturalWidth (optional): location to store the natural width, or
	//      NULL.
	//
	PreferredWidthForHeight(height int32) (minimumWidth, naturalWidth int32)
	// Reset resets any previously cached request and allocation data.
	//
	// When underlying TreeModel data changes its important to reset the context
	// if the content size is allowed to shrink. If the content size is only
	// allowed to grow (this is usually an option for views rendering large data
	// stores as a measure of optimization), then only the row that changed or
	// was inserted needs to be (re)requested with
	// gtk_cell_area_get_preferred_width().
	//
	// When the new overall size of the context requires that the allocated size
	// changes (or whenever this allocation changes at all), the variable row
	// sizes need to be re-requested for every row.
	//
	// For instance, if the rows are displayed all with the same width from top
	// to bottom then a change in the allocated width necessitates a
	// recalculation of all the displayed row heights using
	// gtk_cell_area_get_preferred_height_for_width().
	Reset()
}

// CellAreaContext stores geometrical information for a series of rows in a
// GtkCellArea
//
// The CellAreaContext object is created by a given CellArea implementation via
// its CellAreaClass.create_context() virtual method and is used to store cell
// sizes and alignments for a series of TreeModel rows that are requested and
// rendered in the same context.
//
// CellLayout widgets can create any number of contexts in which to request and
// render groups of data rows. However, it’s important that the same context
// which was used to request sizes for a given TreeModel row also be used for
// the same row when calling other CellArea APIs such as gtk_cell_area_render()
// and gtk_cell_area_event().
type CellAreaContext struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*CellAreaContext)(nil)
)

func classInitCellAreaContexter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellAreaContextClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCellAreaContextClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Allocate(width, height int32) }); ok {
		pclass.allocate = (*[0]byte)(C._gotk4_gtk4_CellAreaContextClass_allocate)
	}

	if _, ok := goval.(interface {
		PreferredHeightForWidth(width int32) (minimumHeight, naturalHeight int32)
	}); ok {
		pclass.get_preferred_height_for_width = (*[0]byte)(C._gotk4_gtk4_CellAreaContextClass_get_preferred_height_for_width)
	}

	if _, ok := goval.(interface {
		PreferredWidthForHeight(height int32) (minimumWidth, naturalWidth int32)
	}); ok {
		pclass.get_preferred_width_for_height = (*[0]byte)(C._gotk4_gtk4_CellAreaContextClass_get_preferred_width_for_height)
	}

	if _, ok := goval.(interface{ Reset() }); ok {
		pclass.reset = (*[0]byte)(C._gotk4_gtk4_CellAreaContextClass_reset)
	}
}

//export _gotk4_gtk4_CellAreaContextClass_allocate
func _gotk4_gtk4_CellAreaContextClass_allocate(arg0 *C.void, arg1 C.int, arg2 C.int) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Allocate(width, height int32) })

	var _width int32  // out
	var _height int32 // out

	_width = int32(arg1)
	_height = int32(arg2)

	iface.Allocate(_width, _height)
}

//export _gotk4_gtk4_CellAreaContextClass_get_preferred_height_for_width
func _gotk4_gtk4_CellAreaContextClass_get_preferred_height_for_width(arg0 *C.void, arg1 C.int, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		PreferredHeightForWidth(width int32) (minimumHeight, naturalHeight int32)
	})

	var _width int32 // out

	_width = int32(arg1)

	minimumHeight, naturalHeight := iface.PreferredHeightForWidth(_width)

	*arg2 = (*C.void)(unsafe.Pointer(minimumHeight))
	*arg3 = (*C.void)(unsafe.Pointer(naturalHeight))
}

//export _gotk4_gtk4_CellAreaContextClass_get_preferred_width_for_height
func _gotk4_gtk4_CellAreaContextClass_get_preferred_width_for_height(arg0 *C.void, arg1 C.int, arg2 *C.void, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		PreferredWidthForHeight(height int32) (minimumWidth, naturalWidth int32)
	})

	var _height int32 // out

	_height = int32(arg1)

	minimumWidth, naturalWidth := iface.PreferredWidthForHeight(_height)

	*arg2 = (*C.void)(unsafe.Pointer(minimumWidth))
	*arg3 = (*C.void)(unsafe.Pointer(naturalWidth))
}

//export _gotk4_gtk4_CellAreaContextClass_reset
func _gotk4_gtk4_CellAreaContextClass_reset(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Reset() })

	iface.Reset()
}

func wrapCellAreaContext(obj *coreglib.Object) *CellAreaContext {
	return &CellAreaContext{
		Object: obj,
	}
}

func marshalCellAreaContext(p uintptr) (interface{}, error) {
	return wrapCellAreaContext(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Allocate allocates a width and/or a height for all rows which are to be
// rendered with context.
//
// Usually allocation is performed only horizontally or sometimes vertically
// since a group of rows are usually rendered side by side vertically or
// horizontally and share either the same width or the same height. Sometimes
// they are allocated in both horizontal and vertical orientations producing a
// homogeneous effect of the rows. This is generally the case for TreeView when
// TreeView:fixed-height-mode is enabled.
//
// The function takes the following parameters:
//
//    - width: allocated width for all TreeModel rows rendered with context, or
//      -1.
//    - height: allocated height for all TreeModel rows rendered with context, or
//      -1.
//
func (context *CellAreaContext) Allocate(width, height int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(height)

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("allocate", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Allocation fetches the current allocation size for context.
//
// If the context was not allocated in width or height, or if the context was
// recently reset with gtk_cell_area_context_reset(), the returned value will be
// -1.
//
// The function returns the following values:
//
//    - width (optional): location to store the allocated width, or NULL.
//    - height (optional): location to store the allocated height, or NULL.
//
func (context *CellAreaContext) Allocation() (width, height int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_allocation", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _width int32  // out
	var _height int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_width = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_height = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _width, _height
}

// Area fetches the CellArea this context was created by.
//
// This is generally unneeded by layouting widgets; however, it is important for
// the context implementation itself to fetch information about the area it is
// being used for.
//
// For instance at CellAreaContextClass.allocate() time it’s important to know
// details about any cell spacing that the CellArea is configured with in order
// to compute a proper allocation.
//
// The function returns the following values:
//
//    - cellArea this context was created by.
//
func (context *CellAreaContext) Area() CellAreaer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_gret := girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_area", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(context)

	var _cellArea CellAreaer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.CellAreaer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(CellAreaer)
			return ok
		})
		rv, ok := casted.(CellAreaer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.CellAreaer")
		}
		_cellArea = rv
	}

	return _cellArea
}

// PreferredHeight gets the accumulative preferred height for all rows which
// have been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are 0.
//
// The function returns the following values:
//
//    - minimumHeight (optional): location to store the minimum height, or NULL.
//    - naturalHeight (optional): location to store the natural height, or NULL.
//
func (context *CellAreaContext) PreferredHeight() (minimumHeight, naturalHeight int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_preferred_height", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _minimumHeight int32 // out
	var _naturalHeight int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_minimumHeight = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_naturalHeight = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _minimumHeight, _naturalHeight
}

// PreferredHeightForWidth gets the accumulative preferred height for width for
// all rows which have been requested for the same said width with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are -1.
//
// The function takes the following parameters:
//
//    - width: proposed width for allocation.
//
// The function returns the following values:
//
//    - minimumHeight (optional): location to store the minimum height, or NULL.
//    - naturalHeight (optional): location to store the natural height, or NULL.
//
func (context *CellAreaContext) PreferredHeightForWidth(width int32) (minimumHeight, naturalHeight int32) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(width)

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_preferred_height_for_width", _args[:], _outs[:])

	runtime.KeepAlive(context)
	runtime.KeepAlive(width)

	var _minimumHeight int32 // out
	var _naturalHeight int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_minimumHeight = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_naturalHeight = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _minimumHeight, _naturalHeight
}

// PreferredWidth gets the accumulative preferred width for all rows which have
// been requested with this context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are 0.
//
// The function returns the following values:
//
//    - minimumWidth (optional): location to store the minimum width, or NULL.
//    - naturalWidth (optional): location to store the natural width, or NULL.
//
func (context *CellAreaContext) PreferredWidth() (minimumWidth, naturalWidth int32) {
	var _args [1]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_preferred_width", _args[:], _outs[:])

	runtime.KeepAlive(context)

	var _minimumWidth int32 // out
	var _naturalWidth int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_minimumWidth = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_naturalWidth = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _minimumWidth, _naturalWidth
}

// PreferredWidthForHeight gets the accumulative preferred width for height for
// all rows which have been requested for the same said height with this
// context.
//
// After gtk_cell_area_context_reset() is called and/or before ever requesting
// the size of a CellArea, the returned values are -1.
//
// The function takes the following parameters:
//
//    - height: proposed height for allocation.
//
// The function returns the following values:
//
//    - minimumWidth (optional): location to store the minimum width, or NULL.
//    - naturalWidth (optional): location to store the natural width, or NULL.
//
func (context *CellAreaContext) PreferredWidthForHeight(height int32) (minimumWidth, naturalWidth int32) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(height)

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("get_preferred_width_for_height", _args[:], _outs[:])

	runtime.KeepAlive(context)
	runtime.KeepAlive(height)

	var _minimumWidth int32 // out
	var _naturalWidth int32 // out

	if *(**C.void)(unsafe.Pointer(&_outs[0])) != nil {
		_minimumWidth = *(*int32)(unsafe.Pointer(_outs[0]))
	}
	if *(**C.void)(unsafe.Pointer(&_outs[1])) != nil {
		_naturalWidth = *(*int32)(unsafe.Pointer(_outs[1]))
	}

	return _minimumWidth, _naturalWidth
}

// PushPreferredHeight causes the minimum and/or natural height to grow if the
// new proposed sizes exceed the current minimum and natural height.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested height
// over a series of gtk_cell_area_get_preferred_height() requests.
//
// The function takes the following parameters:
//
//    - minimumHeight: proposed new minimum height for context.
//    - naturalHeight: proposed new natural height for context.
//
func (context *CellAreaContext) PushPreferredHeight(minimumHeight, naturalHeight int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(minimumHeight)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(naturalHeight)

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("push_preferred_height", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumHeight)
	runtime.KeepAlive(naturalHeight)
}

// PushPreferredWidth causes the minimum and/or natural width to grow if the new
// proposed sizes exceed the current minimum and natural width.
//
// This is used by CellAreaContext implementations during the request process
// over a series of TreeModel rows to progressively push the requested width
// over a series of gtk_cell_area_get_preferred_width() requests.
//
// The function takes the following parameters:
//
//    - minimumWidth: proposed new minimum width for context.
//    - naturalWidth: proposed new natural width for context.
//
func (context *CellAreaContext) PushPreferredWidth(minimumWidth, naturalWidth int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(minimumWidth)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(naturalWidth)

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("push_preferred_width", _args[:], nil)

	runtime.KeepAlive(context)
	runtime.KeepAlive(minimumWidth)
	runtime.KeepAlive(naturalWidth)
}

// Reset resets any previously cached request and allocation data.
//
// When underlying TreeModel data changes its important to reset the context if
// the content size is allowed to shrink. If the content size is only allowed to
// grow (this is usually an option for views rendering large data stores as a
// measure of optimization), then only the row that changed or was inserted
// needs to be (re)requested with gtk_cell_area_get_preferred_width().
//
// When the new overall size of the context requires that the allocated size
// changes (or whenever this allocation changes at all), the variable row sizes
// need to be re-requested for every row.
//
// For instance, if the rows are displayed all with the same width from top to
// bottom then a change in the allocated width necessitates a recalculation of
// all the displayed row heights using
// gtk_cell_area_get_preferred_height_for_width().
func (context *CellAreaContext) Reset() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	girepository.MustFind("Gtk", "CellAreaContext").InvokeMethod("reset", _args[:], nil)

	runtime.KeepAlive(context)
}
