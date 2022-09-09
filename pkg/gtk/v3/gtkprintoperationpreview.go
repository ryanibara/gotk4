// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectReady(gpointer, GtkPrintContext*, guintptr);
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize(gpointer, GtkPrintContext*, GtkPageSetup*, guintptr);
// gboolean _gotk4_gtk3_PrintOperationPreview_virtual_is_selected(void* fnptr, GtkPrintOperationPreview* arg0, gint arg1) {
//   return ((gboolean (*)(GtkPrintOperationPreview*, gint))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_PrintOperationPreview_virtual_end_preview(void* fnptr, GtkPrintOperationPreview* arg0) {
//   ((void (*)(GtkPrintOperationPreview*))(fnptr))(arg0);
// };
// void _gotk4_gtk3_PrintOperationPreview_virtual_got_page_size(void* fnptr, GtkPrintOperationPreview* arg0, GtkPrintContext* arg1, GtkPageSetup* arg2) {
//   ((void (*)(GtkPrintOperationPreview*, GtkPrintContext*, GtkPageSetup*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gtk3_PrintOperationPreview_virtual_ready(void* fnptr, GtkPrintOperationPreview* arg0, GtkPrintContext* arg1) {
//   ((void (*)(GtkPrintOperationPreview*, GtkPrintContext*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_PrintOperationPreview_virtual_render_page(void* fnptr, GtkPrintOperationPreview* arg0, gint arg1) {
//   ((void (*)(GtkPrintOperationPreview*, gint))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypePrintOperationPreview = coreglib.Type(C.gtk_print_operation_preview_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePrintOperationPreview, F: marshalPrintOperationPreview},
	})
}

//
// PrintOperationPreview wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PrintOperationPreview struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*PrintOperationPreview)(nil)
)

// PrintOperationPreviewer describes PrintOperationPreview's interface methods.
type PrintOperationPreviewer interface {
	coreglib.Objector

	// EndPreview ends a preview.
	EndPreview()
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	RenderPage(pageNr int)

	// Got-page-size signal is emitted once for each page that gets rendered to
	// the preview.
	ConnectGotPageSize(func(context *PrintContext, pageSetup *PageSetup)) coreglib.SignalHandle
	// Ready signal gets emitted once per preview operation, before the first
	// page is rendered.
	ConnectReady(func(context *PrintContext)) coreglib.SignalHandle
}

var _ PrintOperationPreviewer = (*PrintOperationPreview)(nil)

func wrapPrintOperationPreview(obj *coreglib.Object) *PrintOperationPreview {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	return wrapPrintOperationPreview(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectGotPageSize signal is emitted once for each page that gets rendered to
// the preview.
//
// A handler for this signal should update the context according to page_setup
// and set up a suitable cairo context, using
// gtk_print_context_set_cairo_context().
func (preview *PrintOperationPreview) ConnectGotPageSize(f func(context *PrintContext, pageSetup *PageSetup)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(preview, "got-page-size", false, unsafe.Pointer(C._gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize), f)
}

// ConnectReady signal gets emitted once per preview operation, before the first
// page is rendered.
//
// A handler for this signal can be used for setup tasks.
func (preview *PrintOperationPreview) ConnectReady(f func(context *PrintContext)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(preview, "ready", false, unsafe.Pointer(C._gotk4_gtk3_PrintOperationPreview_ConnectReady), f)
}

// EndPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (preview *PrintOperationPreview) EndPreview() {
	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))

	C.gtk_print_operation_preview_end_preview(_arg0)
	runtime.KeepAlive(preview)
}

// IsSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
//
// The function takes the following parameters:
//
//    - pageNr: page number.
//
// The function returns the following values:
//
//    - ok: TRUE if the page has been selected for printing.
//
func (preview *PrintOperationPreview) IsSelected(pageNr int) bool {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	_cret = C.gtk_print_operation_preview_is_selected(_arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenderPage renders a page to the preview, using the print context that was
// passed to the PrintOperation::preview handler together with preview.
//
// A custom iprint preview should use this function in its ::expose handler to
// render the currently selected page.
//
// Note that this function requires a suitable cairo context to be associated
// with the print context.
//
// The function takes the following parameters:
//
//    - pageNr: page to render.
//
func (preview *PrintOperationPreview) RenderPage(pageNr int) {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	C.gtk_print_operation_preview_render_page(_arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)
}

// endPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (preview *PrintOperationPreview) endPreview() {
	gclass := (*C.GtkPrintOperationPreviewIface)(coreglib.PeekParentClass(preview))
	fnarg := gclass.end_preview

	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))

	C._gotk4_gtk3_PrintOperationPreview_virtual_end_preview(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(preview)
}

// The function takes the following parameters:
//
//    - context
//    - pageSetup
//
func (preview *PrintOperationPreview) gotPageSize(context *PrintContext, pageSetup *PageSetup) {
	gclass := (*C.GtkPrintOperationPreviewIface)(coreglib.PeekParentClass(preview))
	fnarg := gclass.got_page_size

	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 *C.GtkPrintContext          // out
	var _arg2 *C.GtkPageSetup             // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = (*C.GtkPrintContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GtkPageSetup)(unsafe.Pointer(coreglib.InternObject(pageSetup).Native()))

	C._gotk4_gtk3_PrintOperationPreview_virtual_got_page_size(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(context)
	runtime.KeepAlive(pageSetup)
}

// isSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
//
// The function takes the following parameters:
//
//    - pageNr: page number.
//
// The function returns the following values:
//
//    - ok: TRUE if the page has been selected for printing.
//
func (preview *PrintOperationPreview) isSelected(pageNr int) bool {
	gclass := (*C.GtkPrintOperationPreviewIface)(coreglib.PeekParentClass(preview))
	fnarg := gclass.is_selected

	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	_cret = C._gotk4_gtk3_PrintOperationPreview_virtual_is_selected(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (preview *PrintOperationPreview) ready(context *PrintContext) {
	gclass := (*C.GtkPrintOperationPreviewIface)(coreglib.PeekParentClass(preview))
	fnarg := gclass.ready

	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 *C.GtkPrintContext          // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = (*C.GtkPrintContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	C._gotk4_gtk3_PrintOperationPreview_virtual_ready(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(context)
}

// renderPage renders a page to the preview, using the print context that was
// passed to the PrintOperation::preview handler together with preview.
//
// A custom iprint preview should use this function in its ::expose handler to
// render the currently selected page.
//
// Note that this function requires a suitable cairo context to be associated
// with the print context.
//
// The function takes the following parameters:
//
//    - pageNr: page to render.
//
func (preview *PrintOperationPreview) renderPage(pageNr int) {
	gclass := (*C.GtkPrintOperationPreviewIface)(coreglib.PeekParentClass(preview))
	fnarg := gclass.render_page

	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(coreglib.InternObject(preview).Native()))
	_arg1 = C.gint(pageNr)

	C._gotk4_gtk3_PrintOperationPreview_virtual_render_page(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)
}

// PrintOperationPreviewIface: instance of this type is always passed by
// reference.
type PrintOperationPreviewIface struct {
	*printOperationPreviewIface
}

// printOperationPreviewIface is the struct that's finalized.
type printOperationPreviewIface struct {
	native *C.GtkPrintOperationPreviewIface
}
