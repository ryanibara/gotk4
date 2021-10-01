// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// void _gotk4_gtk3_PageSetupDoneFunc(GtkPageSetup*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_error_get_type()), F: marshalPrintError},
		{T: externglib.Type(C.gtk_print_operation_action_get_type()), F: marshalPrintOperationAction},
		{T: externglib.Type(C.gtk_print_operation_result_get_type()), F: marshalPrintOperationResult},
		{T: externglib.Type(C.gtk_print_status_get_type()), F: marshalPrintStatus},
		{T: externglib.Type(C.gtk_print_operation_get_type()), F: marshalPrintOperationer},
	})
}

// PrintError: error codes that identify various errors that can occur while
// using the GTK+ printing support.
type PrintError int

const (
	// PrintErrorGeneral: unspecified error occurred.
	PrintErrorGeneral PrintError = iota
	// PrintErrorInternalError: internal error occurred.
	PrintErrorInternalError
	// PrintErrorNOMEM: memory allocation failed.
	PrintErrorNOMEM
	// PrintErrorInvalidFile: error occurred while loading a page setup or paper
	// size from a key file.
	PrintErrorInvalidFile
)

func marshalPrintError(p uintptr) (interface{}, error) {
	return PrintError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PrintError.
func (p PrintError) String() string {
	switch p {
	case PrintErrorGeneral:
		return "General"
	case PrintErrorInternalError:
		return "InternalError"
	case PrintErrorNOMEM:
		return "NOMEM"
	case PrintErrorInvalidFile:
		return "InvalidFile"
	default:
		return fmt.Sprintf("PrintError(%d)", p)
	}
}

// PrintErrorQuark registers an error quark for PrintOperation if necessary.
func PrintErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gtk_print_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)

	return _quark
}

// PrintOperationAction: action parameter to gtk_print_operation_run()
// determines what action the print operation should perform.
type PrintOperationAction int

const (
	// PrintOperationActionPrintDialog: show the print dialog.
	PrintOperationActionPrintDialog PrintOperationAction = iota
	// PrintOperationActionPrint: start to print without showing the print
	// dialog, based on the current print settings.
	PrintOperationActionPrint
	// PrintOperationActionPreview: show the print preview.
	PrintOperationActionPreview
	// PrintOperationActionExport: export to a file. This requires the
	// export-filename property to be set.
	PrintOperationActionExport
)

func marshalPrintOperationAction(p uintptr) (interface{}, error) {
	return PrintOperationAction(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PrintOperationAction.
func (p PrintOperationAction) String() string {
	switch p {
	case PrintOperationActionPrintDialog:
		return "PrintDialog"
	case PrintOperationActionPrint:
		return "Print"
	case PrintOperationActionPreview:
		return "Preview"
	case PrintOperationActionExport:
		return "Export"
	default:
		return fmt.Sprintf("PrintOperationAction(%d)", p)
	}
}

// PrintOperationResult: value of this type is returned by
// gtk_print_operation_run().
type PrintOperationResult int

const (
	// PrintOperationResultError: error has occurred.
	PrintOperationResultError PrintOperationResult = iota
	// PrintOperationResultApply: print settings should be stored.
	PrintOperationResultApply
	// PrintOperationResultCancel: print operation has been canceled, the print
	// settings should not be stored.
	PrintOperationResultCancel
	// PrintOperationResultInProgress: print operation is not complete yet. This
	// value will only be returned when running asynchronously.
	PrintOperationResultInProgress
)

func marshalPrintOperationResult(p uintptr) (interface{}, error) {
	return PrintOperationResult(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PrintOperationResult.
func (p PrintOperationResult) String() string {
	switch p {
	case PrintOperationResultError:
		return "Error"
	case PrintOperationResultApply:
		return "Apply"
	case PrintOperationResultCancel:
		return "Cancel"
	case PrintOperationResultInProgress:
		return "InProgress"
	default:
		return fmt.Sprintf("PrintOperationResult(%d)", p)
	}
}

// PrintStatus status gives a rough indication of the completion of a running
// print operation.
type PrintStatus int

const (
	// PrintStatusInitial: printing has not started yet; this status is set
	// initially, and while the print dialog is shown.
	PrintStatusInitial PrintStatus = iota
	// PrintStatusPreparing: this status is set while the begin-print signal is
	// emitted and during pagination.
	PrintStatusPreparing
	// PrintStatusGeneratingData: this status is set while the pages are being
	// rendered.
	PrintStatusGeneratingData
	// PrintStatusSendingData: print job is being sent off to the printer.
	PrintStatusSendingData
	// PrintStatusPending: print job has been sent to the printer, but is not
	// printed for some reason, e.g. the printer may be stopped.
	PrintStatusPending
	// PrintStatusPendingIssue: some problem has occurred during printing, e.g.
	// a paper jam.
	PrintStatusPendingIssue
	// PrintStatusPrinting: printer is processing the print job.
	PrintStatusPrinting
	// PrintStatusFinished: printing has been completed successfully.
	PrintStatusFinished
	// PrintStatusFinishedAborted: printing has been aborted.
	PrintStatusFinishedAborted
)

func marshalPrintStatus(p uintptr) (interface{}, error) {
	return PrintStatus(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for PrintStatus.
func (p PrintStatus) String() string {
	switch p {
	case PrintStatusInitial:
		return "Initial"
	case PrintStatusPreparing:
		return "Preparing"
	case PrintStatusGeneratingData:
		return "GeneratingData"
	case PrintStatusSendingData:
		return "SendingData"
	case PrintStatusPending:
		return "Pending"
	case PrintStatusPendingIssue:
		return "PendingIssue"
	case PrintStatusPrinting:
		return "Printing"
	case PrintStatusFinished:
		return "Finished"
	case PrintStatusFinishedAborted:
		return "FinishedAborted"
	default:
		return fmt.Sprintf("PrintStatus(%d)", p)
	}
}

// PageSetupDoneFunc: type of function that is passed to
// gtk_print_run_page_setup_dialog_async().
//
// This function will be called when the page setup dialog is dismissed, and
// also serves as destroy notify for data.
type PageSetupDoneFunc func(pageSetup *PageSetup)

//export _gotk4_gtk3_PageSetupDoneFunc
func _gotk4_gtk3_PageSetupDoneFunc(arg0 *C.GtkPageSetup, arg1 C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var pageSetup *PageSetup // out

	pageSetup = wrapPageSetup(externglib.Take(unsafe.Pointer(arg0)))

	fn := v.(PageSetupDoneFunc)
	fn(pageSetup)
}

// PrintRunPageSetupDialog runs a page setup dialog, letting the user modify the
// values from page_setup. If the user cancels the dialog, the returned
// PageSetup is identical to the passed in page_setup, otherwise it contains the
// modifications done in the dialog.
//
// Note that this function may use a recursive mainloop to show the page setup
// dialog. See gtk_print_run_page_setup_dialog_async() if this is a problem.
func PrintRunPageSetupDialog(parent *Window, pageSetup *PageSetup, settings *PrintSettings) *PageSetup {
	var _arg1 *C.GtkWindow        // out
	var _arg2 *C.GtkPageSetup     // out
	var _arg3 *C.GtkPrintSettings // out
	var _cret *C.GtkPageSetup     // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}
	if pageSetup != nil {
		_arg2 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))
	}
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_run_page_setup_dialog(_arg1, _arg2, _arg3)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(pageSetup)
	runtime.KeepAlive(settings)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _pageSetup
}

// PrintRunPageSetupDialogAsync runs a page setup dialog, letting the user
// modify the values from page_setup.
//
// In contrast to gtk_print_run_page_setup_dialog(), this function returns after
// showing the page setup dialog on platforms that support this, and calls
// done_cb from a signal handler for the ::response signal of the dialog.
func PrintRunPageSetupDialogAsync(parent *Window, pageSetup *PageSetup, settings *PrintSettings, doneCb PageSetupDoneFunc) {
	var _arg1 *C.GtkWindow           // out
	var _arg2 *C.GtkPageSetup        // out
	var _arg3 *C.GtkPrintSettings    // out
	var _arg4 C.GtkPageSetupDoneFunc // out
	var _arg5 C.gpointer

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}
	if pageSetup != nil {
		_arg2 = (*C.GtkPageSetup)(unsafe.Pointer(pageSetup.Native()))
	}
	_arg3 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg4 = (*[0]byte)(C._gotk4_gtk3_PageSetupDoneFunc)
	_arg5 = C.gpointer(gbox.AssignOnce(doneCb))

	C.gtk_print_run_page_setup_dialog_async(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(pageSetup)
	runtime.KeepAlive(settings)
	runtime.KeepAlive(doneCb)
}

// PrintOperationOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintOperationOverrider interface {
	BeginPrint(context *PrintContext)
	CustomWidgetApply(widget Widgetter)
	Done(result PrintOperationResult)
	DrawPage(context *PrintContext, pageNr int)
	EndPrint(context *PrintContext)
	Paginate(context *PrintContext) bool
	Preview(preview PrintOperationPreviewer, context *PrintContext, parent *Window) bool
	RequestPageSetup(context *PrintContext, pageNr int, setup *PageSetup)
	StatusChanged()
	UpdateCustomWidget(widget Widgetter, setup *PageSetup, settings *PrintSettings)
}

// PrintOperation is the high-level, portable printing API. It looks a bit
// different than other GTK+ dialogs such as the FileChooser, since some
// platforms don’t expose enough infrastructure to implement a good print
// dialog. On such platforms, GtkPrintOperation uses the native print dialog. On
// platforms which do not provide a native print dialog, GTK+ uses its own, see
// PrintUnixDialog.
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
type PrintOperation struct {
	*externglib.Object

	PrintOperationPreview
}

func wrapPrintOperation(obj *externglib.Object) *PrintOperation {
	return &PrintOperation{
		Object: obj,
		PrintOperationPreview: PrintOperationPreview{
			Object: obj,
		},
	}
}

func marshalPrintOperationer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintOperation(obj), nil
}

// NewPrintOperation creates a new PrintOperation.
func NewPrintOperation() *PrintOperation {
	var _cret *C.GtkPrintOperation // in

	_cret = C.gtk_print_operation_new()

	var _printOperation *PrintOperation // out

	_printOperation = wrapPrintOperation(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printOperation
}

// Cancel cancels a running print operation. This function may be called from a
// PrintOperation::begin-print, PrintOperation::paginate or
// PrintOperation::draw-page signal handler to stop the currently running print
// operation.
func (op *PrintOperation) Cancel() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	C.gtk_print_operation_cancel(_arg0)
	runtime.KeepAlive(op)
}

// DrawPageFinish: signalize that drawing of particular page is complete.
//
// It is called after completion of page drawing (e.g. drawing in another
// thread). If gtk_print_operation_set_defer_drawing() was called before, then
// this function has to be called by application. In another case it is called
// by the library itself.
func (op *PrintOperation) DrawPageFinish() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	C.gtk_print_operation_draw_page_finish(_arg0)
	runtime.KeepAlive(op)
}

// DefaultPageSetup returns the default page setup, see
// gtk_print_operation_set_default_page_setup().
func (op *PrintOperation) DefaultPageSetup() *PageSetup {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPageSetup      // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_default_page_setup(_arg0)
	runtime.KeepAlive(op)

	var _pageSetup *PageSetup // out

	_pageSetup = wrapPageSetup(externglib.Take(unsafe.Pointer(_cret)))

	return _pageSetup
}

// EmbedPageSetup gets the value of PrintOperation:embed-page-setup property.
func (op *PrintOperation) EmbedPageSetup() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_embed_page_setup(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Error: call this when the result of a print operation is
// GTK_PRINT_OPERATION_RESULT_ERROR, either as returned by
// gtk_print_operation_run(), or in the PrintOperation::done signal handler. The
// returned #GError will contain more details on what went wrong.
func (op *PrintOperation) Error() error {
	var _arg0 *C.GtkPrintOperation // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	C.gtk_print_operation_get_error(_arg0, &_cerr)
	runtime.KeepAlive(op)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// HasSelection gets the value of PrintOperation:has-selection property.
func (op *PrintOperation) HasSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_has_selection(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// NPagesToPrint returns the number of pages that will be printed.
//
// Note that this value is set during print preparation phase
// (GTK_PRINT_STATUS_PREPARING), so this function should never be called before
// the data generation phase (GTK_PRINT_STATUS_GENERATING_DATA). You can connect
// to the PrintOperation::status-changed signal and call
// gtk_print_operation_get_n_pages_to_print() when print status is
// GTK_PRINT_STATUS_GENERATING_DATA. This is typically used to track the
// progress of print operation.
func (op *PrintOperation) NPagesToPrint() int {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gint               // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_n_pages_to_print(_arg0)
	runtime.KeepAlive(op)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrintSettings returns the current print settings.
//
// Note that the return value is NULL until either
// gtk_print_operation_set_print_settings() or gtk_print_operation_run() have
// been called.
func (op *PrintOperation) PrintSettings() *PrintSettings {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.GtkPrintSettings  // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_print_settings(_arg0)
	runtime.KeepAlive(op)

	var _printSettings *PrintSettings // out

	_printSettings = wrapPrintSettings(externglib.Take(unsafe.Pointer(_cret)))

	return _printSettings
}

// Status returns the status of the print operation. Also see
// gtk_print_operation_get_status_string().
func (op *PrintOperation) Status() PrintStatus {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.GtkPrintStatus     // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_status(_arg0)
	runtime.KeepAlive(op)

	var _printStatus PrintStatus // out

	_printStatus = PrintStatus(_cret)

	return _printStatus
}

// StatusString returns a string representation of the status of the print
// operation. The string is translated and suitable for displaying the print
// status e.g. in a Statusbar.
//
// Use gtk_print_operation_get_status() to obtain a status value that is
// suitable for programmatic use.
func (op *PrintOperation) StatusString() string {
	var _arg0 *C.GtkPrintOperation // out
	var _cret *C.gchar             // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_status_string(_arg0)
	runtime.KeepAlive(op)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SupportSelection gets the value of PrintOperation:support-selection property.
func (op *PrintOperation) SupportSelection() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_get_support_selection(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsFinished: convenience function to find out if the print operation is
// finished, either successfully (GTK_PRINT_STATUS_FINISHED) or unsuccessfully
// (GTK_PRINT_STATUS_FINISHED_ABORTED).
//
// Note: when you enable print status tracking the print operation can be in a
// non-finished state even after done has been called, as the operation status
// then tracks the print job status on the printer.
func (op *PrintOperation) IsFinished() bool {
	var _arg0 *C.GtkPrintOperation // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	_cret = C.gtk_print_operation_is_finished(_arg0)
	runtime.KeepAlive(op)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Run runs the print operation, by first letting the user modify print settings
// in the print dialog, and then print the document.
//
// Normally that this function does not return until the rendering of all pages
// is complete. You can connect to the PrintOperation::status-changed signal on
// op to obtain some information about the progress of the print operation.
// Furthermore, it may use a recursive mainloop to show the print dialog.
//
// If you call gtk_print_operation_set_allow_async() or set the
// PrintOperation:allow-async property the operation will run asynchronously if
// this is supported on the platform. The PrintOperation::done signal will be
// emitted with the result of the operation when the it is done (i.e. when the
// dialog is canceled, or when the print succeeds or fails).
//
//    if (settings != NULL)
//      gtk_print_operation_set_print_settings (print, settings);
//
//    if (page_setup != NULL)
//      gtk_print_operation_set_default_page_setup (print, page_setup);
//
//    g_signal_connect (print, "begin-print",
//                      G_CALLBACK (begin_print), &data);
//    g_signal_connect (print, "draw-page",
//                      G_CALLBACK (draw_page), &data);
//
//    res = gtk_print_operation_run (print,
//                                   GTK_PRINT_OPERATION_ACTION_PRINT_DIALOG,
//                                   parent,
//                                   &error);
//
//    if (res == GTK_PRINT_OPERATION_RESULT_ERROR)
//     {
//       error_dialog = gtk_message_dialog_new (GTK_WINDOW (parent),
//      			                     GTK_DIALOG_DESTROY_WITH_PARENT,
//    					     GTK_MESSAGE_ERROR,
//    					     GTK_BUTTONS_CLOSE,
//    					     "Error printing file:\ns",
//    					     error->message);
//       g_signal_connect (error_dialog, "response",
//                         G_CALLBACK (gtk_widget_destroy), NULL);
//       gtk_widget_show (error_dialog);
//       g_error_free (error);
//     }
//    else if (res == GTK_PRINT_OPERATION_RESULT_APPLY)
//     {
//       if (settings != NULL)
//    g_object_unref (settings);
//       settings = g_object_ref (gtk_print_operation_get_print_settings (print));
//     }
//
// Note that gtk_print_operation_run() can only be called once on a given
// PrintOperation.
func (op *PrintOperation) Run(action PrintOperationAction, parent *Window) (PrintOperationResult, error) {
	var _arg0 *C.GtkPrintOperation      // out
	var _arg1 C.GtkPrintOperationAction // out
	var _arg2 *C.GtkWindow              // out
	var _cret C.GtkPrintOperationResult // in
	var _cerr *C.GError                 // in

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.GtkPrintOperationAction(action)
	if parent != nil {
		_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}

	_cret = C.gtk_print_operation_run(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(op)
	runtime.KeepAlive(action)
	runtime.KeepAlive(parent)

	var _printOperationResult PrintOperationResult // out
	var _goerr error                               // out

	_printOperationResult = PrintOperationResult(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _printOperationResult, _goerr
}

// SetAllowAsync sets whether the gtk_print_operation_run() may return before
// the print operation is completed. Note that some platforms may not allow
// asynchronous operation.
func (op *PrintOperation) SetAllowAsync(allowAsync bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if allowAsync {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_allow_async(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(allowAsync)
}

// SetCurrentPage sets the current page.
//
// If this is called before gtk_print_operation_run(), the user will be able to
// select to print only the current page.
//
// Note that this only makes sense for pre-paginated documents.
func (op *PrintOperation) SetCurrentPage(currentPage int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.gint(currentPage)

	C.gtk_print_operation_set_current_page(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(currentPage)
}

// SetCustomTabLabel sets the label for the tab holding custom widgets.
func (op *PrintOperation) SetCustomTabLabel(label string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_print_operation_set_custom_tab_label(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(label)
}

// SetDefaultPageSetup makes default_page_setup the default page setup for op.
//
// This page setup will be used by gtk_print_operation_run(), but it can be
// overridden on a per-page basis by connecting to the
// PrintOperation::request-page-setup signal.
func (op *PrintOperation) SetDefaultPageSetup(defaultPageSetup *PageSetup) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPageSetup      // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if defaultPageSetup != nil {
		_arg1 = (*C.GtkPageSetup)(unsafe.Pointer(defaultPageSetup.Native()))
	}

	C.gtk_print_operation_set_default_page_setup(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(defaultPageSetup)
}

// SetDeferDrawing sets up the PrintOperation to wait for calling of
// gtk_print_operation_draw_page_finish() from application. It can be used for
// drawing page in another thread.
//
// This function must be called in the callback of “draw-page” signal.
func (op *PrintOperation) SetDeferDrawing() {
	var _arg0 *C.GtkPrintOperation // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))

	C.gtk_print_operation_set_defer_drawing(_arg0)
	runtime.KeepAlive(op)
}

// SetEmbedPageSetup: embed page size combo box and orientation combo box into
// page setup page. Selected page setup is stored as default page setup in
// PrintOperation.
func (op *PrintOperation) SetEmbedPageSetup(embed bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if embed {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_embed_page_setup(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(embed)
}

// SetExportFilename sets up the PrintOperation to generate a file instead of
// showing the print dialog. The indended use of this function is for
// implementing “Export to PDF” actions. Currently, PDF is the only supported
// format.
//
// “Print to PDF” support is independent of this and is done by letting the user
// pick the “Print to PDF” item from the list of printers in the print dialog.
func (op *PrintOperation) SetExportFilename(filename string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_export_filename(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(filename)
}

// SetHasSelection sets whether there is a selection to print.
//
// Application has to set number of pages to which the selection will draw by
// gtk_print_operation_set_n_pages() in a callback of
// PrintOperation::begin-print.
func (op *PrintOperation) SetHasSelection(hasSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if hasSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_has_selection(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(hasSelection)
}

// SetJobName sets the name of the print job. The name is used to identify the
// job (e.g. in monitoring applications like eggcups).
//
// If you don’t set a job name, GTK+ picks a default one by numbering successive
// print jobs.
func (op *PrintOperation) SetJobName(jobName string) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(jobName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_operation_set_job_name(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(jobName)
}

// SetNPages sets the number of pages in the document.
//
// This must be set to a positive number before the rendering starts. It may be
// set in a PrintOperation::begin-print signal hander.
//
// Note that the page numbers passed to the PrintOperation::request-page-setup
// and PrintOperation::draw-page signals are 0-based, i.e. if the user chooses
// to print all pages, the last ::draw-page signal will be for page n_pages - 1.
func (op *PrintOperation) SetNPages(nPages int) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.gint(nPages)

	C.gtk_print_operation_set_n_pages(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(nPages)
}

// SetPrintSettings sets the print settings for op. This is typically used to
// re-establish print settings from a previous print operation, see
// gtk_print_operation_run().
func (op *PrintOperation) SetPrintSettings(printSettings *PrintSettings) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 *C.GtkPrintSettings  // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if printSettings != nil {
		_arg1 = (*C.GtkPrintSettings)(unsafe.Pointer(printSettings.Native()))
	}

	C.gtk_print_operation_set_print_settings(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(printSettings)
}

// SetShowProgress: if show_progress is TRUE, the print operation will show a
// progress dialog during the print operation.
func (op *PrintOperation) SetShowProgress(showProgress bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if showProgress {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_show_progress(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(showProgress)
}

// SetSupportSelection sets whether selection is supported by PrintOperation.
func (op *PrintOperation) SetSupportSelection(supportSelection bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if supportSelection {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_support_selection(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(supportSelection)
}

// SetTrackPrintStatus: if track_status is TRUE, the print operation will try to
// continue report on the status of the print job in the printer queues and
// printer. This can allow your application to show things like “out of paper”
// issues, and when the print job actually reaches the printer.
//
// This function is often implemented using some form of polling, so it should
// not be enabled unless needed.
func (op *PrintOperation) SetTrackPrintStatus(trackStatus bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if trackStatus {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_track_print_status(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(trackStatus)
}

// SetUnit sets up the transformation for the cairo context obtained from
// PrintContext in such a way that distances are measured in units of unit.
func (op *PrintOperation) SetUnit(unit Unit) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.GtkUnit            // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	_arg1 = C.GtkUnit(unit)

	C.gtk_print_operation_set_unit(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(unit)
}

// SetUseFullPage: if full_page is TRUE, the transformation for the cairo
// context obtained from PrintContext puts the origin at the top left corner of
// the page (which may not be the top left corner of the sheet, depending on
// page orientation and the number of pages per sheet). Otherwise, the origin is
// at the top left corner of the imageable area (i.e. inside the margins).
func (op *PrintOperation) SetUseFullPage(fullPage bool) {
	var _arg0 *C.GtkPrintOperation // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkPrintOperation)(unsafe.Pointer(op.Native()))
	if fullPage {
		_arg1 = C.TRUE
	}

	C.gtk_print_operation_set_use_full_page(_arg0, _arg1)
	runtime.KeepAlive(op)
	runtime.KeepAlive(fullPage)
}
