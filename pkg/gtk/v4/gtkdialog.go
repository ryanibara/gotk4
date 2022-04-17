// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void _gotk4_gtk4_DialogClass_close(GtkDialog*);
// extern void _gotk4_gtk4_DialogClass_response(GtkDialog*, int);
// extern void _gotk4_gtk4_Dialog_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk4_Dialog_ConnectResponse(gpointer, gint, guintptr);
import "C"

// glib.Type values for gtkdialog.go.
var (
	GTypeResponseType = externglib.Type(C.gtk_response_type_get_type())
	GTypeDialogFlags  = externglib.Type(C.gtk_dialog_flags_get_type())
	GTypeDialog       = externglib.Type(C.gtk_dialog_get_type())
)

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeResponseType, F: marshalResponseType},
		{T: GTypeDialogFlags, F: marshalDialogFlags},
		{T: GTypeDialog, F: marshalDialog},
	})
}

// ResponseType: predefined values for use as response ids in
// gtk_dialog_add_button().
//
// All predefined values are negative; GTK leaves values of 0 or greater for
// application-defined response ids.
type ResponseType C.gint

const (
	// ResponseNone: returned if an action widget has no response id, or if the
	// dialog gets programmatically hidden or destroyed.
	ResponseNone ResponseType = -1
	// ResponseReject: generic response id, not used by GTK dialogs.
	ResponseReject ResponseType = -2
	// ResponseAccept: generic response id, not used by GTK dialogs.
	ResponseAccept ResponseType = -3
	// ResponseDeleteEvent: returned if the dialog is deleted.
	ResponseDeleteEvent ResponseType = -4
	// ResponseOK: returned by OK buttons in GTK dialogs.
	ResponseOK ResponseType = -5
	// ResponseCancel: returned by Cancel buttons in GTK dialogs.
	ResponseCancel ResponseType = -6
	// ResponseClose: returned by Close buttons in GTK dialogs.
	ResponseClose ResponseType = -7
	// ResponseYes: returned by Yes buttons in GTK dialogs.
	ResponseYes ResponseType = -8
	// ResponseNo: returned by No buttons in GTK dialogs.
	ResponseNo ResponseType = -9
	// ResponseApply: returned by Apply buttons in GTK dialogs.
	ResponseApply ResponseType = -10
	// ResponseHelp: returned by Help buttons in GTK dialogs.
	ResponseHelp ResponseType = -11
)

func marshalResponseType(p uintptr) (interface{}, error) {
	return ResponseType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
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
type DialogFlags C.guint

const (
	// DialogModal: make the constructed dialog modal.
	DialogModal DialogFlags = 0b1
	// DialogDestroyWithParent: destroy the dialog when its parent is destroyed.
	DialogDestroyWithParent DialogFlags = 0b10
	// DialogUseHeaderBar: create dialog with actions in header bar instead of
	// action area.
	DialogUseHeaderBar DialogFlags = 0b100
)

func marshalDialogFlags(p uintptr) (interface{}, error) {
	return DialogFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
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

// Has returns true if d contains other.
func (d DialogFlags) Has(other DialogFlags) bool {
	return (d & other) == other
}

// DialogOverrider contains methods that are overridable.
type DialogOverrider interface {
	externglib.Objector
	Close()
	// Response emits the ::response signal with the given response ID.
	//
	// Used to indicate that the user has responded to the dialog in some way.
	//
	// The function takes the following parameters:
	//
	//    - responseId: response ID.
	//
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
	_ [0]func() // equal guard
	Window
}

var (
	_ Widgetter           = (*Dialog)(nil)
	_ externglib.Objector = (*Dialog)(nil)
)

func classInitDialogger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkDialogClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkDialogClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Close() }); ok {
		pclass.close = (*[0]byte)(C._gotk4_gtk4_DialogClass_close)
	}

	if _, ok := goval.(interface{ Response(responseId int) }); ok {
		pclass.response = (*[0]byte)(C._gotk4_gtk4_DialogClass_response)
	}
}

//export _gotk4_gtk4_DialogClass_close
func _gotk4_gtk4_DialogClass_close(arg0 *C.GtkDialog) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk4_DialogClass_response
func _gotk4_gtk4_DialogClass_response(arg0 *C.GtkDialog, arg1 C.int) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Response(responseId int) })

	var _responseId int // out

	_responseId = int(arg1)

	iface.Response(_responseId)
}

func wrapDialog(obj *externglib.Object) *Dialog {
	return &Dialog{
		Window: Window{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Root: Root{
				NativeSurface: NativeSurface{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						Accessible: Accessible{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						ConstraintTarget: ConstraintTarget{
							Object: obj,
						},
					},
				},
			},
			ShortcutManager: ShortcutManager{
				Object: obj,
			},
		},
	}
}

func marshalDialog(p uintptr) (interface{}, error) {
	return wrapDialog(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_Dialog_ConnectClose
func _gotk4_gtk4_Dialog_ConnectClose(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClose is emitted when the user uses a keybinding to close the dialog.
//
// This is a keybinding signal (class.SignalAction.html).
//
// The default binding for this signal is the Escape key.
func (dialog *Dialog) ConnectClose(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(dialog, "close", false, unsafe.Pointer(C._gotk4_gtk4_Dialog_ConnectClose), f)
}

//export _gotk4_gtk4_Dialog_ConnectResponse
func _gotk4_gtk4_Dialog_ConnectResponse(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(responseId int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(responseId int))
	}

	var _responseId int // out

	_responseId = int(arg1)

	f(_responseId)
}

// ConnectResponse is emitted when an action widget is clicked.
//
// The signal is also emitted when the dialog receives a delete event, and when
// gtk.Dialog.Response() is called. On a delete event, the response ID is
// GTK_RESPONSE_DELETE_EVENT. Otherwise, it depends on which action widget was
// clicked.
func (dialog *Dialog) ConnectResponse(f func(responseId int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(dialog, "response", false, unsafe.Pointer(C._gotk4_gtk4_Dialog_ConnectResponse), f)
}

// NewDialog creates a new dialog box.
//
// Widgets should not be packed into the GtkWindow directly, but into the
// content_area and action_area, as described above.
//
// The function returns the following values:
//
//    - dialog: new dialog as a GtkWidget.
//
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
//
// The function takes the following parameters:
//
//    - child: activatable widget.
//    - responseId: response ID for child.
//
func (dialog *Dialog) AddActionWidget(child Widgetter, responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(child).Native()))
	_arg2 = C.int(responseId)

	C.gtk_dialog_add_action_widget(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(child)
	runtime.KeepAlive(responseId)
}

// AddButton adds a button with the given text.
//
// GTK arranges things so that clicking the button will emit the
// gtk.Dialog::response signal with the given response_id. The button is
// appended to the end of the dialog’s action area. The button widget is
// returned, but usually you don’t need it.
//
// The function takes the following parameters:
//
//    - buttonText: text of button.
//    - responseId: response ID for the button.
//
// The function returns the following values:
//
//    - widget: GtkButton widget that was added.
//
func (dialog *Dialog) AddButton(buttonText string, responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.char      // out
	var _arg2 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(buttonText)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(responseId)

	_cret = C.gtk_dialog_add_button(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(buttonText)
	runtime.KeepAlive(responseId)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// ContentArea returns the content area of dialog.
//
// The function returns the following values:
//
//    - box: content area Box.
//
func (dialog *Dialog) ContentArea() *Box {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))

	_cret = C.gtk_dialog_get_content_area(_arg0)
	runtime.KeepAlive(dialog)

	var _box *Box // out

	_box = wrapBox(externglib.Take(unsafe.Pointer(_cret)))

	return _box
}

// HeaderBar returns the header bar of dialog.
//
// Note that the headerbar is only used by the dialog if the
// gtk.Dialog:use-header-bar property is TRUE.
//
// The function returns the following values:
//
//    - headerBar: header bar.
//
func (dialog *Dialog) HeaderBar() *HeaderBar {
	var _arg0 *C.GtkDialog // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))

	_cret = C.gtk_dialog_get_header_bar(_arg0)
	runtime.KeepAlive(dialog)

	var _headerBar *HeaderBar // out

	_headerBar = wrapHeaderBar(externglib.Take(unsafe.Pointer(_cret)))

	return _headerBar
}

// ResponseForWidget gets the response id of a widget in the action area of a
// dialog.
//
// The function takes the following parameters:
//
//    - widget in the action area of dialog.
//
// The function returns the following values:
//
//    - gint: response id of widget, or GTK_RESPONSE_NONE if widget doesn’t have
//      a response id set.
//
func (dialog *Dialog) ResponseForWidget(widget Widgetter) int {
	var _arg0 *C.GtkDialog // out
	var _arg1 *C.GtkWidget // out
	var _cret C.int        // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(externglib.InternObject(widget).Native()))

	_cret = C.gtk_dialog_get_response_for_widget(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(widget)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// WidgetForResponse gets the widget button that uses the given response ID in
// the action area of a dialog.
//
// The function takes the following parameters:
//
//    - responseId: response ID used by the dialog widget.
//
// The function returns the following values:
//
//    - widget (optional) button that uses the given response_id, or NULL.
//
func (dialog *Dialog) WidgetForResponse(responseId int) Widgetter {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = C.int(responseId)

	_cret = C.gtk_dialog_get_widget_for_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Response emits the ::response signal with the given response ID.
//
// Used to indicate that the user has responded to the dialog in some way.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (dialog *Dialog) Response(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
}

// SetDefaultResponse sets the default widget for the dialog based on the
// response ID.
//
// Pressing “Enter” normally activates the default widget.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//
func (dialog *Dialog) SetDefaultResponse(responseId int) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = C.int(responseId)

	C.gtk_dialog_set_default_response(_arg0, _arg1)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
}

// SetResponseSensitive: convenient way to sensitize/desensitize dialog buttons.
//
// Calls gtk_widget_set_sensitive (widget, setting) for each widget in the
// dialog’s action area with the given response_id.
//
// The function takes the following parameters:
//
//    - responseId: response ID.
//    - setting: TRUE for sensitive.
//
func (dialog *Dialog) SetResponseSensitive(responseId int, setting bool) {
	var _arg0 *C.GtkDialog // out
	var _arg1 C.int        // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GtkDialog)(unsafe.Pointer(externglib.InternObject(dialog).Native()))
	_arg1 = C.int(responseId)
	if setting {
		_arg2 = C.TRUE
	}

	C.gtk_dialog_set_response_sensitive(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dialog)
	runtime.KeepAlive(responseId)
	runtime.KeepAlive(setting)
}
