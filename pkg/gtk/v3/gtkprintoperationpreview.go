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
// extern gboolean _gotk4_gtk3_PrintOperationPreviewIface_is_selected(GtkPrintOperationPreview*, gint);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_end_preview(GtkPrintOperationPreview*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_got_page_size(GtkPrintOperationPreview*, GtkPrintContext*, GtkPageSetup*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_ready(GtkPrintOperationPreview*, GtkPrintContext*);
// extern void _gotk4_gtk3_PrintOperationPreviewIface_render_page(GtkPrintOperationPreview*, gint);
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize(gpointer, GtkPrintContext*, GtkPageSetup*, guintptr);
// extern void _gotk4_gtk3_PrintOperationPreview_ConnectReady(gpointer, GtkPrintContext*, guintptr);
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

// PrintOperationPreviewOverrider contains methods that are overridable.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()
	// The function takes the following parameters:
	//
	//    - context
	//    - pageSetup
	//
	GotPageSize(context *PrintContext, pageSetup *PageSetup)
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
	IsSelected(pageNr int) bool
	// The function takes the following parameters:
	//
	Ready(context *PrintContext)
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	//
	// A custom iprint preview should use this function in its ::expose handler
	// to render the currently selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	//
	// The function takes the following parameters:
	//
	//    - pageNr: page to render.
	//
	RenderPage(pageNr int)
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

func ifaceInitPrintOperationPreviewer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkPrintOperationPreviewIface)(unsafe.Pointer(gifacePtr))
	iface.end_preview = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_end_preview)
	iface.got_page_size = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_got_page_size)
	iface.is_selected = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_is_selected)
	iface.ready = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_ready)
	iface.render_page = (*[0]byte)(C._gotk4_gtk3_PrintOperationPreviewIface_render_page)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_end_preview
func _gotk4_gtk3_PrintOperationPreviewIface_end_preview(arg0 *C.GtkPrintOperationPreview) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	iface.EndPreview()
}

//export _gotk4_gtk3_PrintOperationPreviewIface_got_page_size
func _gotk4_gtk3_PrintOperationPreviewIface_got_page_size(arg0 *C.GtkPrintOperationPreview, arg1 *C.GtkPrintContext, arg2 *C.GtkPageSetup) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _context *PrintContext // out
	var _pageSetup *PageSetup  // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageSetup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg2)))

	iface.GotPageSize(_context, _pageSetup)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_is_selected
func _gotk4_gtk3_PrintOperationPreviewIface_is_selected(arg0 *C.GtkPrintOperationPreview, arg1 C.gint) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _pageNr int // out

	_pageNr = int(arg1)

	ok := iface.IsSelected(_pageNr)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_PrintOperationPreviewIface_ready
func _gotk4_gtk3_PrintOperationPreviewIface_ready(arg0 *C.GtkPrintOperationPreview, arg1 *C.GtkPrintContext) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _context *PrintContext // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))

	iface.Ready(_context)
}

//export _gotk4_gtk3_PrintOperationPreviewIface_render_page
func _gotk4_gtk3_PrintOperationPreviewIface_render_page(arg0 *C.GtkPrintOperationPreview, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(PrintOperationPreviewOverrider)

	var _pageNr int // out

	_pageNr = int(arg1)

	iface.RenderPage(_pageNr)
}

func wrapPrintOperationPreview(obj *coreglib.Object) *PrintOperationPreview {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	return wrapPrintOperationPreview(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize
func _gotk4_gtk3_PrintOperationPreview_ConnectGotPageSize(arg0 C.gpointer, arg1 *C.GtkPrintContext, arg2 *C.GtkPageSetup, arg3 C.guintptr) {
	var f func(context *PrintContext, pageSetup *PageSetup)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext, pageSetup *PageSetup))
	}

	var _context *PrintContext // out
	var _pageSetup *PageSetup  // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))
	_pageSetup = wrapPageSetup(coreglib.Take(unsafe.Pointer(arg2)))

	f(_context, _pageSetup)
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

//export _gotk4_gtk3_PrintOperationPreview_ConnectReady
func _gotk4_gtk3_PrintOperationPreview_ConnectReady(arg0 C.gpointer, arg1 *C.GtkPrintContext, arg2 C.guintptr) {
	var f func(context *PrintContext)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context *PrintContext))
	}

	var _context *PrintContext // out

	_context = wrapPrintContext(coreglib.Take(unsafe.Pointer(arg1)))

	f(_context)
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
