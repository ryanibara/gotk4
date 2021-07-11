// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_print_operation_preview_get_type()), F: marshalPrintOperationPreviewer},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()

	GotPageSize(context PrintContexter, pageSetup PageSetupper)
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool

	Ready(context PrintContexter)
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with @preview.
	//
	// A custom iprint preview should use this function in its ::expose handler
	// to render the currently selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	RenderPage(pageNr int)
}

// PrintOperationPreviewer describes PrintOperationPreview's methods.
type PrintOperationPreviewer interface {
	// EndPreview ends a preview.
	EndPreview()
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with @preview.
	RenderPage(pageNr int)
}

type PrintOperationPreview struct {
	*externglib.Object
}

var (
	_ PrintOperationPreviewer = (*PrintOperationPreview)(nil)
	_ gextras.Nativer         = (*PrintOperationPreview)(nil)
)

func wrapPrintOperationPreview(obj *externglib.Object) PrintOperationPreviewer {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreviewer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintOperationPreview(obj), nil
}

// EndPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (preview *PrintOperationPreview) EndPreview() {
	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))

	C.gtk_print_operation_preview_end_preview(_arg0)
}

// IsSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
func (preview *PrintOperationPreview) IsSelected(pageNr int) bool {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))
	_arg1 = C.gint(pageNr)

	_cret = C.gtk_print_operation_preview_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenderPage renders a page to the preview, using the print context that was
// passed to the PrintOperation::preview handler together with @preview.
//
// A custom iprint preview should use this function in its ::expose handler to
// render the currently selected page.
//
// Note that this function requires a suitable cairo context to be associated
// with the print context.
func (preview *PrintOperationPreview) RenderPage(pageNr int) {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))
	_arg1 = C.gint(pageNr)

	C.gtk_print_operation_preview_render_page(_arg0, _arg1)
}
