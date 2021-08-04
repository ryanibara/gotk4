// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"strings"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_response_type_get_type()), F: marshalResponseType},
		{T: externglib.Type(C.gtk_dialog_flags_get_type()), F: marshalDialogFlags},
		{T: externglib.Type(C.gtk_dialog_get_type()), F: marshalDialogger},
	})
}

// ResponseType: predefined values for use as response ids in
// gtk_dialog_add_button().
//
// All predefined values are negative; GTK leaves values of 0 or greater for
// application-defined response ids.
type ResponseType int

const (
	// ResponseNone: returned if an action widget has no response id, or if the
	// dialog gets programmatically hidden or destroyed
	ResponseNone ResponseType = -1
	// ResponseReject: generic response id, not used by GTK dialogs
	ResponseReject ResponseType = -2
	// ResponseAccept: generic response id, not used by GTK dialogs
	ResponseAccept ResponseType = -3
	// ResponseDeleteEvent: returned if the dialog is deleted
	ResponseDeleteEvent ResponseType = -4
	// ResponseOK: returned by OK buttons in GTK dialogs
	ResponseOK ResponseType = -5
	// ResponseCancel: returned by Cancel buttons in GTK dialogs
	ResponseCancel ResponseType = -6
	// ResponseClose: returned by Close buttons in GTK dialogs
	ResponseClose ResponseType = -7
	// ResponseYes: returned by Yes buttons in GTK dialogs
	ResponseYes ResponseType = -8
	// ResponseNo: returned by No buttons in GTK dialogs
	ResponseNo ResponseType = -9
	// ResponseApply: returned by Apply buttons in GTK dialogs
	ResponseApply ResponseType = -10
	// ResponseHelp: returned by Help buttons in GTK dialogs
	ResponseHelp ResponseType = -11
)

func marshalResponseType(p uintptr) (interface{}, error) {
	return ResponseType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ResponseType.
func (r ResponseType) String() string {
	switch r {
	case ResponseNone:
		return "None"
	case ResponseReject:
		return "Reject"
	case ResponseAccept:
		return "Accept"
	case ResponseDeleteEvent:
		return "DeleteEvent"
	case ResponseOK:
		return "OK"
	case ResponseCancel:
		return "Cancel"
	case ResponseClose:
		return "Close"
	case ResponseYes:
		return "Yes"
	case ResponseNo:
		return "No"
	case ResponseApply:
		return "Apply"
	case ResponseHelp:
		return "Help"
	default:
		return fmt.Sprintf("ResponseType(%d)", r)
	}
}

// DialogFlags flags used to influence dialog construction.
type DialogFlags int

const (
	// DialogModal: make the constructed dialog modal
	DialogModal DialogFlags = 0b1
	// DialogDestroyWithParent: destroy the dialog when its parent is destroyed
	DialogDestroyWithParent DialogFlags = 0b10
	// DialogUseHeaderBar: create dialog with actions in header bar instead of
	// action area
	DialogUseHeaderBar DialogFlags = 0b100
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	return DialogFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the names in string for DialogFlags.
func (d DialogFlags) String() string {
	if d == 0 {
		return "DialogFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(54)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case DialogModal:
			builder.WriteString("Modal|")
		case DialogDestroyWithParent:
			builder.WriteString("DestroyWithParent|")
		case DialogUseHeaderBar:
			builder.WriteString("UseHeaderBar|")
		default:
			builder.WriteString(fmt.Sprintf("DialogFlags(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// DialogOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DialogOverrider interface {
	Close()
	// Response emits the ::response signal with the given response ID.
	//
	// Used to indicate that the user has responded to the dialog in some way.
	Response(responseId int)
}

// Dialog dialogs are a convenient way to prompt the user for a small amount of
// input.
//
// !An example GtkDialog (dialog.png)
//
// Typical uses are to display a message, ask a question, or anything else that
// does not require extensive effort on the user’s part.
//
// The main area of a GtkDialog is called the "content area", and is yours to
// populate with widgets such a GtkLabel or GtkEntry, to present your
// information, questions, or tasks to the user.
//
// In addition, dialogs allow you to add "action widgets". Most commonly, action
// widgets are buttons. Depending on the platform, action widgets may be
// presented in the header bar at the top of the window, or at the bottom of the
// window. To add action widgets, create your GtkDialog using
// gtk.Dialog.NewWithButtons, or use gtk.Dialog.AddButton(),
// gtk.Dialog.AddButtons(), or gtk.Dialog.AddActionWidget().
//
// GtkDialogs uses some heuristics to decide whether to add a close button to
// the window decorations. If any of the action buttons use the response ID
// GTK_RESPONSE_CLOSE or GTK_RESPONSE_CANCEL, the close button is omitted.
//
// Clicking a button that was added as an action widget will emit the
// gtk.Dialog::response signal with a response ID that you specified. GTK will
// never assign a meaning to positive response IDs; these are entirely
// user-defined. But for convenience, you can use the response IDs in the
// gtk.ResponseType enumeration (these all have values less than zero). If a
// dialog receives a delete event, the gtk.Dialog::response signal will be
// emitted with the GTK_RESPONSE_DELETE_EVENT response ID.
//
// Dialogs are created with a call to gtk.Dialog.New or
// gtk.Dialog.NewWithButtons. The latter is recommended; it allows you to set
// the dialog title, some convenient flags, and add buttons.
//
// A “modal” dialog (that is, one which freezes the rest of the application from
// user input), can be created by calling gtk.Window.SetModal() on the dialog.
// When using gtk.Dialog.NewWithButtons, you can also pass the GTK_DIALOG_MODAL
// flag to make a dialog modal.
//
// For the simple dialog in the following example, a gtk.MessageDialog would
// save some effort. But you’d need to create the dialog contents manually if
// you had more than a simple message in the dialog.
//
// An example for simple GtkDialog usage:
//
//    // Function to open a dialog box with a message
//    void
//    quick_message (GtkWindow *parent, char *message)
//    {
//     GtkWidget *dialog, *label, *content_area;
//     GtkDialogFlags flags;
//
//     // Create the widgets
//     flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//     dialog = gtk_dialog_new_with_buttons ("Message",
//                                           parent,
//                                           flags,
//                                           _("_OK"),
//                                           GTK_RESPONSE_NONE,
//                                           NULL);
//     content_area = gtk_dialog_get_content_area (GTK_DIALOG (dialog));
//     label = gtk_label_new (message);
//
//     // Ensure that the dialog box is destroyed when the user responds
//
//     g_signal_connect_swapped (dialog,
//                               "response",
//                               G_CALLBACK (gtk_window_destroy),
//                               dialog);
//
//     // Add the label, and show everything we’ve added
//
//     gtk_box_append (GTK_BOX (content_area), label);
//     gtk_widget_show (dialog);
//    }
//
//
//
// GtkDialog as GtkBuildable
//
// The GtkDialog implementation of the GtkBuildable interface exposes the
// content_area as an internal child with the name “content_area”.
//
// GtkDialog supports a custom <action-widgets> element, which can contain
// multiple <action-widget> elements. The “response” attribute specifies a
// numeric response, and the content of the element is the id of widget (which
// should be a child of the dialogs action_area). To mark a response as default,
// set the “default“ attribute of the <action-widget> element to true.
//
// GtkDialog supports adding action widgets by specifying “action“ as the “type“
// attribute of a <child> element. The widget will be added either to the action
// area or the headerbar of the dialog, depending on the “use-header-bar“
// property. The response id has to be associated with the action widget using
// the <action-widgets> element.
//
// An example of a Dialog UI definition fragment:
//
//    <object class="GtkDialog" id="dialog1">
//      <child type="action">
//        <object class="GtkButton" id="button_cancel"/>
//      </child>
//      <child type="action">
//        <object class="GtkButton" id="button_ok">
//        </object>
//      </child>
//      <action-widgets>
//        <action-widget response="cancel">button_cancel</action-widget>
//        <action-widget response="ok" default="true">button_ok</action-widget>
//      </action-widgets>
//    </object>
//
//
//
// Accessibility
//
// GtkDialog uses the GTK_ACCESSIBLE_ROLE_DIALOG role.
type Dialog struct {
	Window
}

func wrapDialog(obj *externglib.Object) *Dialog {
	return &Dialog{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Accessible: Accessible{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						ConstraintTarget: ConstraintTarget{
							Object: obj,
						},
						Object: obj,
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDialog(obj), nil
}

// NewDialog creates a new dialog box.
//
// Widgets should not be packed into the GtkWindow directly, but into the
// content_area and action_area, as described above.
func NewDialog() *Dialog {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_dialog_new()

	var _dialog *Dialog // out

	_dialog = wrapDialog(externglib.Take(unsafe.Pointer(_cret)))

	return _dialog
}

// AddActionWidget adds an activatable widget to the action area of a GtkDialog.
//
// GTK connects a signal handler that will emit the gtk.Dialog::response signal
// on the dialog when the widget is activated. The widget is appended to the end
// of the dialog’s action area.
//
// If you want to add a non-activatable widget, simply pack it into the
// action_area field of the GtkDialog struct.
func (dialog *Dialog) AddActionWidget(child Widgetter, responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(responseId)

	C.gtk_dialog_add_action_widget(_arg0, _arg1, _arg2)
}

// AddButton adds a button with the given text.
//
// GTK arranges things so that clicking the button will emit the
// gtk.Dialog::response signal with the given response_id. The button is
// appended to the end of the dialog’s action area. The button widget is
// returned, but usually you don’t need it.
func (dialog *Dialog) AddButton(buttonText string, responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(buttonText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(responseId)

	_cret = C.gtk_dialog_add_button(_arg0, _arg1, _arg2)

	var _widget Widgetter // out

	_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// ContentArea returns the content area of dialog.
func (dialog *Dialog) ContentArea() *Box {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_get_content_area(_arg0)

	var _box *Box // out

	_box = wrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// HeaderBar returns the header bar of dialog.
//
// Note that the headerbar is only used by the dialog if the
// gtk.Dialog:use-header-bar property is TRUE.
func (dialog *Dialog) HeaderBar() *HeaderBar {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))

	_cret = C.gtk_dialog_get_header_bar(_arg0)

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// ResponseForWidget gets the response id of a widget in the action area of a
// dialog.
func (dialog *Dialog) ResponseForWidget(widget Widgetter) int {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _cret C.int        // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_dialog_get_response_for_widget(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WidgetForResponse gets the widget button that uses the given response ID in
// the action area of a dialog.
func (dialog *Dialog) WidgetForResponse(responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.int(responseId)

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Response emits the ::response signal with the given response ID.
//
// Used to indicate that the user has responded to the dialog in some way.
func (dialog *Dialog) Response(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_response(_arg0, _arg1)
}

// SetDefaultResponse sets the default widget for the dialog based on the
// response ID.
//
// Pressing “Enter” normally activates the default widget.
func (dialog *Dialog) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_set_default_response(_arg0, _arg1)
}

// SetResponseSensitive: convenient way to sensitize/desensitize dialog buttons.
//
// Calls gtk_widget_set_sensitive (widget, setting) for each widget in the
// dialog’s action area with the given response_id.
func (dialog *Dialog) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(dialog.Native()))
	_arg1 = C.int(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_dialog_set_response_sensitive(_arg0, _arg1, _arg2)
}
