// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_operation_get_type()), F: marshalPrintOperation},
	})
}

// PrintRunPageSetupDialogAsync runs a page setup dialog, letting the user
// modify the values from @page_setup.
//
// In contrast to gtk_print_run_page_setup_dialog(), this function returns after
// showing the page setup dialog on platforms that support this, and calls
// @done_cb from a signal handler for the ::response signal of the dialog.
func PrintRunPageSetupDialogAsync(parent Window, pageSetup PageSetup, settings PrintSettings, doneCb PageSetupDoneFunc) {
	var _arg1 *C.GtkWindow
	var _arg2 *C.GtkPageSetup
	var _arg3 *C.GtkPrintSettings
	var _arg4 C.GtkPageSetupDoneFunc
	var _arg5 C.gpointer

	_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	_arg2 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg4 = (*[0]byte)(C.gotk4_PageSetupDoneFunc)
	_arg5 = C.gpointer(box.Assign(doneCb))

	C.gtk_print_run_page_setup_dialog_async(_arg1, _arg2, _arg3, _arg4, _arg5)
}

// PrintOperation: gtkPrintOperation is the high-level, portable printing API.
// It looks a bit different than other GTK+ dialogs such as the FileChooser,
// since some platforms don’t expose enough infrastructure to implement a good
// print dialog. On such platforms, GtkPrintOperation uses the native print
// dialog. On platforms which do not provide a native print dialog, GTK+ uses
// its own, see PrintUnixDialog.
//
// The typical way to use the high-level printing API is to create a
// GtkPrintOperation object with gtk_print_operation_new() when the user selects
// to print. Then you set some properties on it, e.g. the page size, any
// PrintSettings from previous print operations, the number of pages, the
// current page, etc.
//
// Then you start the print operation by calling gtk_print_operation_run(). It
// will then show a dialog, let the user select a printer and options. When the
// user finished the dialog various signals will be emitted on the
// PrintOperation, the main one being PrintOperation::draw-page, which you are
// supposed to catch and render the page on the provided PrintContext using
// Cairo.
//
// The high-level printing API
//
//    static GtkPrintSettings *settings = NULL;
//
//    static void
//    do_print (void)
//    {
//      GtkPrintOperation *print;
//      GtkPrintOperationResult res;
//
//      print = gtk_print_operation_new ();
//
//      if (settings != NULL)
//        gtk_print_operation_set_print_settings (print, settings);
//
//      g_signal_connect (print, "begin_print", G_CALLBACK (begin_print), NULL);
//      g_signal_connect (print, "draw_page", G_CALLBACK (draw_page), NULL);
//
//      res = gtk_print_operation_run (print, GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG,
//                                     GTK_WINDOW (main_window), NULL);
//
//      if (res == GTK_PRINT_OPERATION_RESULT_APPLY)
//        {
//          if (settings != NULL)
//            g_object_unref (settings);
//          settings = g_object_ref (gtk_print_operation_get_print_settings (print));
//        }
//
//      g_object_unref (print);
//    }
//
// By default GtkPrintOperation uses an external application to do print
// preview. To implement a custom print preview, an application must connect to
// the preview signal. The functions gtk_print_operation_preview_render_page(),
// gtk_print_operation_preview_end_preview() and
// gtk_print_operation_preview_is_selected() are useful when implementing a
// print preview.
type PrintOperation interface {
	gextras.Objector
	PrintOperationPreview

	// Cancel cancels a running print operation. This function may be called
	// from a PrintOperation::begin-print, PrintOperation::paginate or
	// PrintOperation::draw-page signal handler to stop the currently running
	// print operation.
	Cancel()
	// DrawPageFinish: signalize that drawing of particular page is complete.
	//
	// It is called after completion of page drawing (e.g. drawing in another
	// thread). If gtk_print_operation_set_defer_drawing() was called before,
	// then this function has to be called by application. In another case it is
	// called by the library itself.
	DrawPageFinish()
	// EmbedPageSetup gets the value of PrintOperation:embed-page-setup
	// property.
	EmbedPageSetup() bool
	// Error: call this when the result of a print operation is
	// GTK_PRINT_OPERATION_RESULT_ERROR, either as returned by
	// gtk_print_operation_run(), or in the PrintOperation::done signal handler.
	// The returned #GError will contain more details on what went wrong.
	Error() error
	// HasSelection gets the value of PrintOperation:has-selection property.
	HasSelection() bool
	// NPagesToPrint returns the number of pages that will be printed.
	//
	// Note that this value is set during print preparation phase
	// (GTK_PRINT_STATUS_PREPARING), so this function should never be called
	// before the data generation phase (GTK_PRINT_STATUS_GENERATING_DATA). You
	// can connect to the PrintOperation::status-changed signal and call
	// gtk_print_operation_get_n_pages_to_print() when print status is
	// GTK_PRINT_STATUS_GENERATING_DATA. This is typically used to track the
	// progress of print operation.
	NPagesToPrint() int
	// StatusString returns a string representation of the status of the print
	// operation. The string is translated and suitable for displaying the print
	// status e.g. in a Statusbar.
	//
	// Use gtk_print_operation_get_status() to obtain a status value that is
	// suitable for programmatic use.
	StatusString() string
	// SupportSelection gets the value of PrintOperation:support-selection
	// property.
	SupportSelection() bool
	// IsFinished: a convenience function to find out if the print operation is
	// finished, either successfully (GTK_PRINT_STATUS_FINISHED) or
	// unsuccessfully (GTK_PRINT_STATUS_FINISHED_ABORTED).
	//
	// Note: when you enable print status tracking the print operation can be in
	// a non-finished state even after done has been called, as the operation
	// status then tracks the print job status on the printer.
	IsFinished() bool
	// SetAllowAsync sets whether the gtk_print_operation_run() may return
	// before the print operation is completed. Note that some platforms may not
	// allow asynchronous operation.
	SetAllowAsync(allowAsync bool)
	// SetCurrentPage sets the current page.
	//
	// If this is called before gtk_print_operation_run(), the user will be able
	// to select to print only the current page.
	//
	// Note that this only makes sense for pre-paginated documents.
	SetCurrentPage(currentPage int)
	// SetCustomTabLabel sets the label for the tab holding custom widgets.
	SetCustomTabLabel(label string)
	// SetDefaultPageSetup makes @default_page_setup the default page setup for
	// @op.
	//
	// This page setup will be used by gtk_print_operation_run(), but it can be
	// overridden on a per-page basis by connecting to the
	// PrintOperation::request-page-setup signal.
	SetDefaultPageSetup(defaultPageSetup PageSetup)
	// SetDeferDrawing sets up the PrintOperation to wait for calling of
	// gtk_print_operation_draw_page_finish() from application. It can be used
	// for drawing page in another thread.
	//
	// This function must be called in the callback of “draw-page” signal.
	SetDeferDrawing()
	// SetEmbedPageSetup: embed page size combo box and orientation combo box
	// into page setup page. Selected page setup is stored as default page setup
	// in PrintOperation.
	SetEmbedPageSetup(embed bool)
	// SetExportFilename sets up the PrintOperation to generate a file instead
	// of showing the print dialog. The indended use of this function is for
	// implementing “Export to PDF” actions. Currently, PDF is the only
	// supported format.
	//
	// “Print to PDF” support is independent of this and is done by letting the
	// user pick the “Print to PDF” item from the list of printers in the print
	// dialog.
	SetExportFilename(filename *string)
	// SetHasSelection sets whether there is a selection to print.
	//
	// Application has to set number of pages to which the selection will draw
	// by gtk_print_operation_set_n_pages() in a callback of
	// PrintOperation::begin-print.
	SetHasSelection(hasSelection bool)
	// SetJobName sets the name of the print job. The name is used to identify
	// the job (e.g. in monitoring applications like eggcups).
	//
	// If you don’t set a job name, GTK+ picks a default one by numbering
	// successive print jobs.
	SetJobName(jobName string)
	// SetNPages sets the number of pages in the document.
	//
	// This must be set to a positive number before the rendering starts. It may
	// be set in a PrintOperation::begin-print signal hander.
	//
	// Note that the page numbers passed to the
	// PrintOperation::request-page-setup and PrintOperation::draw-page signals
	// are 0-based, i.e. if the user chooses to print all pages, the last
	// ::draw-page signal will be for page @n_pages - 1.
	SetNPages(nPages int)
	// SetPrintSettings sets the print settings for @op. This is typically used
	// to re-establish print settings from a previous print operation, see
	// gtk_print_operation_run().
	SetPrintSettings(printSettings PrintSettings)
	// SetShowProgress: if @show_progress is true, the print operation will show
	// a progress dialog during the print operation.
	SetShowProgress(showProgress bool)
	// SetSupportSelection sets whether selection is supported by
	// PrintOperation.
	SetSupportSelection(supportSelection bool)
	// SetTrackPrintStatus: if track_status is true, the print operation will
	// try to continue report on the status of the print job in the printer
	// queues and printer. This can allow your application to show things like
	// “out of paper” issues, and when the print job actually reaches the
	// printer.
	//
	// This function is often implemented using some form of polling, so it
	// should not be enabled unless needed.
	SetTrackPrintStatus(trackStatus bool)
	// SetUnit sets up the transformation for the cairo context obtained from
	// PrintContext in such a way that distances are measured in units of @unit.
	SetUnit(unit Unit)
	// SetUseFullPage: if @full_page is true, the transformation for the cairo
	// context obtained from PrintContext puts the origin at the top left corner
	// of the page (which may not be the top left corner of the sheet, depending
	// on page orientation and the number of pages per sheet). Otherwise, the
	// origin is at the top left corner of the imageable area (i.e. inside the
	// margins).
	SetUseFullPage(fullPage bool)
}

// printOperation implements the PrintOperation interface.
type printOperation struct {
	gextras.Objector
	PrintOperationPreview
}

var _ PrintOperation = (*printOperation)(nil)

// WrapPrintOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapPrintOperation(obj *externglib.Object) PrintOperation {
	return PrintOperation{
		Objector:              obj,
		PrintOperationPreview: WrapPrintOperationPreview(obj),
	}
}

func marshalPrintOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintOperation(obj), nil
}

// Cancel cancels a running print operation. This function may be called
// from a PrintOperation::begin-print, PrintOperation::paginate or
// PrintOperation::draw-page signal handler to stop the currently running
// print operation.
func (o printOperation) Cancel() {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_cancel(_arg0)
}

// DrawPageFinish: signalize that drawing of particular page is complete.
//
// It is called after completion of page drawing (e.g. drawing in another
// thread). If gtk_print_operation_set_defer_drawing() was called before,
// then this function has to be called by application. In another case it is
// called by the library itself.
func (o printOperation) DrawPageFinish() {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_draw_page_finish(_arg0)
}

// EmbedPageSetup gets the value of PrintOperation:embed-page-setup
// property.
func (o printOperation) EmbedPageSetup() bool {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean

	_cret = C.gtk_print_operation_get_embed_page_setup(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Error: call this when the result of a print operation is
// GTK_PRINT_OPERATION_RESULT_ERROR, either as returned by
// gtk_print_operation_run(), or in the PrintOperation::done signal handler.
// The returned #GError will contain more details on what went wrong.
func (o printOperation) Error() error {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cerr *C.GError

	C.gtk_print_operation_get_error(_arg0, _cerr)

	var _goerr error

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// HasSelection gets the value of PrintOperation:has-selection property.
func (o printOperation) HasSelection() bool {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean

	_cret = C.gtk_print_operation_get_has_selection(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// NPagesToPrint returns the number of pages that will be printed.
//
// Note that this value is set during print preparation phase
// (GTK_PRINT_STATUS_PREPARING), so this function should never be called
// before the data generation phase (GTK_PRINT_STATUS_GENERATING_DATA). You
// can connect to the PrintOperation::status-changed signal and call
// gtk_print_operation_get_n_pages_to_print() when print status is
// GTK_PRINT_STATUS_GENERATING_DATA. This is typically used to track the
// progress of print operation.
func (o printOperation) NPagesToPrint() int {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gint

	_cret = C.gtk_print_operation_get_n_pages_to_print(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// StatusString returns a string representation of the status of the print
// operation. The string is translated and suitable for displaying the print
// status e.g. in a Statusbar.
//
// Use gtk_print_operation_get_status() to obtain a status value that is
// suitable for programmatic use.
func (o printOperation) StatusString() string {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret *C.gchar

	_cret = C.gtk_print_operation_get_status_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SupportSelection gets the value of PrintOperation:support-selection
// property.
func (o printOperation) SupportSelection() bool {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean

	_cret = C.gtk_print_operation_get_support_selection(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// IsFinished: a convenience function to find out if the print operation is
// finished, either successfully (GTK_PRINT_STATUS_FINISHED) or
// unsuccessfully (GTK_PRINT_STATUS_FINISHED_ABORTED).
//
// Note: when you enable print status tracking the print operation can be in
// a non-finished state even after done has been called, as the operation
// status then tracks the print job status on the printer.
func (o printOperation) IsFinished() bool {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	var _cret C.gboolean

	_cret = C.gtk_print_operation_is_finished(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// SetAllowAsync sets whether the gtk_print_operation_run() may return
// before the print operation is completed. Note that some platforms may not
// allow asynchronous operation.
func (o printOperation) SetAllowAsync(allowAsync bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if allowAsync {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_allow_async(_arg0, _arg1)
}

// SetCurrentPage sets the current page.
//
// If this is called before gtk_print_operation_run(), the user will be able
// to select to print only the current page.
//
// Note that this only makes sense for pre-paginated documents.
func (o printOperation) SetCurrentPage(currentPage int) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gint

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.gint(currentPage)

	C.gtk_print_operation_set_current_page(_arg0, _arg1)
}

// SetCustomTabLabel sets the label for the tab holding custom widgets.
func (o printOperation) SetCustomTabLabel(label string) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 *C.gchar

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_custom_tab_label(_arg0, _arg1)
}

// SetDefaultPageSetup makes @default_page_setup the default page setup for
// @op.
//
// This page setup will be used by gtk_print_operation_run(), but it can be
// overridden on a per-page basis by connecting to the
// PrintOperation::request-page-setup signal.
func (o printOperation) SetDefaultPageSetup(defaultPageSetup PageSetup) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 *C.GtkPageSetup

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(defaultPageSetup.Native()))

	C.gtk_print_operation_set_default_page_setup(_arg0, _arg1)
}

// SetDeferDrawing sets up the PrintOperation to wait for calling of
// gtk_print_operation_draw_page_finish() from application. It can be used
// for drawing page in another thread.
//
// This function must be called in the callback of “draw-page” signal.
func (o printOperation) SetDeferDrawing() {
	var _arg0 *C.GtkPrintOperation

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))

	C.gtk_print_operation_set_defer_drawing(_arg0)
}

// SetEmbedPageSetup: embed page size combo box and orientation combo box
// into page setup page. Selected page setup is stored as default page setup
// in PrintOperation.
func (o printOperation) SetEmbedPageSetup(embed bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if embed {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_embed_page_setup(_arg0, _arg1)
}

// SetExportFilename sets up the PrintOperation to generate a file instead
// of showing the print dialog. The indended use of this function is for
// implementing “Export to PDF” actions. Currently, PDF is the only
// supported format.
//
// “Print to PDF” support is independent of this and is done by letting the
// user pick the “Print to PDF” item from the list of printers in the print
// dialog.
func (o printOperation) SetExportFilename(filename *string) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 *C.gchar

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_export_filename(_arg0, _arg1)
}

// SetHasSelection sets whether there is a selection to print.
//
// Application has to set number of pages to which the selection will draw
// by gtk_print_operation_set_n_pages() in a callback of
// PrintOperation::begin-print.
func (o printOperation) SetHasSelection(hasSelection bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if hasSelection {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_has_selection(_arg0, _arg1)
}

// SetJobName sets the name of the print job. The name is used to identify
// the job (e.g. in monitoring applications like eggcups).
//
// If you don’t set a job name, GTK+ picks a default one by numbering
// successive print jobs.
func (o printOperation) SetJobName(jobName string) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 *C.gchar

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.gchar)(C.CString(jobName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_job_name(_arg0, _arg1)
}

// SetNPages sets the number of pages in the document.
//
// This must be set to a positive number before the rendering starts. It may
// be set in a PrintOperation::begin-print signal hander.
//
// Note that the page numbers passed to the
// PrintOperation::request-page-setup and PrintOperation::draw-page signals
// are 0-based, i.e. if the user chooses to print all pages, the last
// ::draw-page signal will be for page @n_pages - 1.
func (o printOperation) SetNPages(nPages int) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gint

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.gint(nPages)

	C.gtk_print_operation_set_n_pages(_arg0, _arg1)
}

// SetPrintSettings sets the print settings for @op. This is typically used
// to re-establish print settings from a previous print operation, see
// gtk_print_operation_run().
func (o printOperation) SetPrintSettings(printSettings PrintSettings) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 *C.GtkPrintSettings

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(printSettings.Native()))

	C.gtk_print_operation_set_print_settings(_arg0, _arg1)
}

// SetShowProgress: if @show_progress is true, the print operation will show
// a progress dialog during the print operation.
func (o printOperation) SetShowProgress(showProgress bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if showProgress {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_show_progress(_arg0, _arg1)
}

// SetSupportSelection sets whether selection is supported by
// PrintOperation.
func (o printOperation) SetSupportSelection(supportSelection bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if supportSelection {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_support_selection(_arg0, _arg1)
}

// SetTrackPrintStatus: if track_status is true, the print operation will
// try to continue report on the status of the print job in the printer
// queues and printer. This can allow your application to show things like
// “out of paper” issues, and when the print job actually reaches the
// printer.
//
// This function is often implemented using some form of polling, so it
// should not be enabled unless needed.
func (o printOperation) SetTrackPrintStatus(trackStatus bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if trackStatus {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_track_print_status(_arg0, _arg1)
}

// SetUnit sets up the transformation for the cairo context obtained from
// PrintContext in such a way that distances are measured in units of @unit.
func (o printOperation) SetUnit(unit Unit) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.GtkUnit

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (C.GtkUnit)(unit)

	C.gtk_print_operation_set_unit(_arg0, _arg1)
}

// SetUseFullPage: if @full_page is true, the transformation for the cairo
// context obtained from PrintContext puts the origin at the top left corner
// of the page (which may not be the top left corner of the sheet, depending
// on page orientation and the number of pages per sheet). Otherwise, the
// origin is at the top left corner of the imageable area (i.e. inside the
// margins).
func (o printOperation) SetUseFullPage(fullPage bool) {
	var _arg0 *C.GtkPrintOperation
	var _arg1 C.gboolean

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(o.Native()))
	if fullPage {
		_arg1 = C.gboolean(1)
	}

	C.gtk_print_operation_set_use_full_page(_arg0, _arg1)
}
