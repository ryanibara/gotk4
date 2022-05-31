// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_AssistantClass_apply(GtkAssistant*);
// extern void _gotk4_gtk3_AssistantClass_cancel(GtkAssistant*);
// extern void _gotk4_gtk3_AssistantClass_close(GtkAssistant*);
// extern void _gotk4_gtk3_AssistantClass_prepare(GtkAssistant*, GtkWidget*);
// extern void _gotk4_gtk3_Assistant_ConnectApply(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectCancel(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectEscape(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectPrepare(gpointer, GtkWidget*, guintptr);
import "C"

// glib.Type values for gtkassistant.go.
var (
	GTypeAssistantPageType = coreglib.Type(C.gtk_assistant_page_type_get_type())
	GTypeAssistant         = coreglib.Type(C.gtk_assistant_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAssistantPageType, F: marshalAssistantPageType},
		{T: GTypeAssistant, F: marshalAssistant},
	})
}

// AssistantPageType: enum for determining the page role inside the Assistant.
// It's used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType C.gint

const (
	// AssistantPageContent: page has regular contents. Both the Back and
	// forward buttons will be shown.
	AssistantPageContent AssistantPageType = iota
	// AssistantPageIntro: page contains an introduction to the assistant task.
	// Only the Forward button will be shown if there is a next page.
	AssistantPageIntro
	// AssistantPageConfirm: page lets the user confirm or deny the changes. The
	// Back and Apply buttons will be shown.
	AssistantPageConfirm
	// AssistantPageSummary: page informs the user of the changes done. Only the
	// Close button will be shown.
	AssistantPageSummary
	// AssistantPageProgress: used for tasks that take a long time to complete,
	// blocks the assistant until the page is marked as complete. Only the back
	// button will be shown.
	AssistantPageProgress
	// AssistantPageCustom: used for when other page types are not appropriate.
	// No buttons will be shown, and the application must add its own buttons
	// through gtk_assistant_add_action_widget().
	AssistantPageCustom
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AssistantPageType.
func (a AssistantPageType) String() string {
	switch a {
	case AssistantPageContent:
		return "Content"
	case AssistantPageIntro:
		return "Intro"
	case AssistantPageConfirm:
		return "Confirm"
	case AssistantPageSummary:
		return "Summary"
	case AssistantPageProgress:
		return "Progress"
	case AssistantPageCustom:
		return "Custom"
	default:
		return fmt.Sprintf("AssistantPageType(%d)", a)
	}
}

// AssistantPageFunc: function used by gtk_assistant_set_forward_page_func() to
// know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int32) (gint int32)

//export _gotk4_gtk3_AssistantPageFunc
func _gotk4_gtk3_AssistantPageFunc(arg1 C.gint, arg2 C.gpointer) (cret C.gint) {
	var fn AssistantPageFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AssistantPageFunc)
	}

	var _currentPage int32 // out

	_currentPage = int32(arg1)

	gint := fn(_currentPage)

	cret = C.gint(gint)

	return cret
}

// AssistantOverrider contains methods that are overridable.
type AssistantOverrider interface {
	Apply()
	Cancel()
	Close()
	// The function takes the following parameters:
	//
	Prepare(page Widgetter)
}

// Assistant is a widget used to represent a generally complex operation
// splitted in several steps, guiding the user through its pages and controlling
// the page flow to collect the necessary data.
//
// The design of GtkAssistant is that it controls what buttons to show and to
// make sensitive, based on what it knows about the page sequence and the
// [type][GtkAssistantPageType] of each page, in addition to state information
// like the page [completion][gtk-assistant-set-page-complete] and
// [committed][gtk-assistant-commit] status.
//
// If you have a case that doesn’t quite fit in Assistants way of handling
// buttons, you can use the K_ASSISTANT_PAGE_CUSTOM page type and handle buttons
// yourself.
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the Buildable interface exposes the
// action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*Assistant)(nil)
)

func classInitAssistanter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkAssistantClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkAssistantClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Apply() }); ok {
		pclass.apply = (*[0]byte)(C._gotk4_gtk3_AssistantClass_apply)
	}

	if _, ok := goval.(interface{ Cancel() }); ok {
		pclass.cancel = (*[0]byte)(C._gotk4_gtk3_AssistantClass_cancel)
	}

	if _, ok := goval.(interface{ Close() }); ok {
		pclass.close = (*[0]byte)(C._gotk4_gtk3_AssistantClass_close)
	}

	if _, ok := goval.(interface{ Prepare(page Widgetter) }); ok {
		pclass.prepare = (*[0]byte)(C._gotk4_gtk3_AssistantClass_prepare)
	}
}

//export _gotk4_gtk3_AssistantClass_apply
func _gotk4_gtk3_AssistantClass_apply(arg0 *C.GtkAssistant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Apply() })

	iface.Apply()
}

//export _gotk4_gtk3_AssistantClass_cancel
func _gotk4_gtk3_AssistantClass_cancel(arg0 *C.GtkAssistant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancel() })

	iface.Cancel()
}

//export _gotk4_gtk3_AssistantClass_close
func _gotk4_gtk3_AssistantClass_close(arg0 *C.GtkAssistant) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk3_AssistantClass_prepare
func _gotk4_gtk3_AssistantClass_prepare(arg0 *C.GtkAssistant, arg1 *C.GtkWidget) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Prepare(page Widgetter) })

	var _page Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_page = rv
	}

	iface.Prepare(_page)
}

func wrapAssistant(obj *coreglib.Object) *Assistant {
	return &Assistant{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	return wrapAssistant(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Assistant_ConnectApply
func _gotk4_gtk3_Assistant_ConnectApply(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectApply signal is emitted when the apply button is clicked.
//
// The default behavior of the Assistant is to switch to the page after the
// current page, unless the current page is the last one.
//
// A handler for the ::apply signal should carry out the actions for which the
// wizard has collected data. If the action takes a long time to complete, you
// might consider putting a page of type GTK_ASSISTANT_PAGE_PROGRESS after the
// confirmation page and handle this operation within the Assistant::prepare
// signal of the progress page.
func (assistant *Assistant) ConnectApply(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "apply", false, unsafe.Pointer(C._gotk4_gtk3_Assistant_ConnectApply), f)
}

//export _gotk4_gtk3_Assistant_ConnectCancel
func _gotk4_gtk3_Assistant_ConnectCancel(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectCancel signal is emitted when then the cancel button is clicked.
func (assistant *Assistant) ConnectCancel(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "cancel", false, unsafe.Pointer(C._gotk4_gtk3_Assistant_ConnectCancel), f)
}

//export _gotk4_gtk3_Assistant_ConnectClose
func _gotk4_gtk3_Assistant_ConnectClose(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectClose signal is emitted either when the close button of a summary page
// is clicked, or when the apply button in the last page in the flow (of type
// GTK_ASSISTANT_PAGE_CONFIRM) is clicked.
func (assistant *Assistant) ConnectClose(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "close", false, unsafe.Pointer(C._gotk4_gtk3_Assistant_ConnectClose), f)
}

//export _gotk4_gtk3_Assistant_ConnectEscape
func _gotk4_gtk3_Assistant_ConnectEscape(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

func (assistant *Assistant) ConnectEscape(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "escape", false, unsafe.Pointer(C._gotk4_gtk3_Assistant_ConnectEscape), f)
}

//export _gotk4_gtk3_Assistant_ConnectPrepare
func _gotk4_gtk3_Assistant_ConnectPrepare(arg0 C.gpointer, arg1 *C.GtkWidget, arg2 C.guintptr) {
	var f func(page Widgetter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(page Widgetter))
	}

	var _page Widgetter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_page = rv
	}

	f(_page)
}

// ConnectPrepare signal is emitted when a new page is set as the assistant's
// current page, before making the new page visible.
//
// A handler for this signal can do any preparations which are necessary before
// showing page.
func (assistant *Assistant) ConnectPrepare(f func(page Widgetter)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(assistant, "prepare", false, unsafe.Pointer(C._gotk4_gtk3_Assistant_ConnectPrepare), f)
}

// NewAssistant creates a new Assistant.
//
// The function returns the following values:
//
//    - assistant: newly created Assistant.
//
func NewAssistant() *Assistant {
	var _cret *C.void // in

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("new_Assistant", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _assistant *Assistant // out

	_assistant = wrapAssistant(coreglib.Take(unsafe.Pointer(_cret)))

	return _assistant
}

// AddActionWidget adds a widget to the action area of a Assistant.
//
// The function takes the following parameters:
//
//    - child: Widget.
//
func (assistant *Assistant) AddActionWidget(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("add_action_widget", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// AppendPage appends a page to the assistant.
//
// The function takes the following parameters:
//
//    - page: Widget.
//
// The function returns the following values:
//
//    - gint: index (starting at 0) of the inserted page.
//
func (assistant *Assistant) AppendPage(page Widgetter) int32 {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("append_page", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// Commit erases the visited page history so the back button is not shown on the
// current page, and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is hereafter
// deemed permanent and cannot be modified or undone. For example, showing a
// progress page to track a long-running, unreversible operation after the user
// has clicked apply on a confirmation page.
func (assistant *Assistant) Commit() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("commit", args[:], nil)

	runtime.KeepAlive(assistant)
}

// CurrentPage returns the page number of the current page.
//
// The function returns the following values:
//
//    - gint: index (starting from 0) of the current page in the assistant, or -1
//      if the assistant has no pages, or no current page.
//
func (assistant *Assistant) CurrentPage() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_current_page", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NPages returns the number of pages in the assistant.
//
// The function returns the following values:
//
//    - gint: number of pages in the assistant.
//
func (assistant *Assistant) NPages() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_n_pages", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number page_num.
//
// The function takes the following parameters:
//
//    - pageNum: index of a page in the assistant, or -1 to get the last page.
//
// The function returns the following values:
//
//    - widget (optional): child widget, or NULL if page_num is out of bounds.
//
func (assistant *Assistant) NthPage(pageNum int32) Widgetter {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.gint(pageNum)
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_nth_page", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

// PageComplete gets whether page is complete.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - ok: TRUE if page is complete.
//
func (assistant *Assistant) PageComplete(page Widgetter) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_complete", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHasPadding gets whether page has padding.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - ok: TRUE if page has padding.
//
func (assistant *Assistant) PageHasPadding(page Widgetter) bool {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_has_padding", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PageHeaderImage gets the header image for page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - pixbuf: header image for page, or NULL if there’s no header image for the
//      page.
//
func (assistant *Assistant) PageHeaderImage(page Widgetter) *gdkpixbuf.Pixbuf {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_header_image", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// PageSideImage gets the side image for page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - pixbuf: side image for page, or NULL if there’s no side image for the
//      page.
//
func (assistant *Assistant) PageSideImage(page Widgetter) *gdkpixbuf.Pixbuf {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_side_image", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	return _pixbuf
}

// PageTitle gets the title for page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - utf8: title for page.
//
func (assistant *Assistant) PageTitle(page Widgetter) string {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("get_page_title", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InsertPage inserts a page in the assistant at a given position.
//
// The function takes the following parameters:
//
//    - page: Widget.
//    - position: index (starting at 0) at which to insert the page, or -1 to
//      append the page to the assistant.
//
// The function returns the following values:
//
//    - gint: index (starting from 0) of the inserted page.
//
func (assistant *Assistant) InsertPage(page Widgetter, position int32) int32 {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 C.gint  // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = C.gint(position)
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("insert_page", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(position)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) NextPage() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("next_page", args[:], nil)

	runtime.KeepAlive(assistant)
}

// PrependPage prepends a page to the assistant.
//
// The function takes the following parameters:
//
//    - page: Widget.
//
// The function returns the following values:
//
//    - gint: index (starting at 0) of the inserted page.
//
func (assistant *Assistant) PrependPage(page Widgetter) int32 {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	_gret := girepository.MustFind("Gtk", "Assistant").InvokeMethod("prepend_page", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) PreviousPage() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("previous_page", args[:], nil)

	runtime.KeepAlive(assistant)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
//
// The function takes the following parameters:
//
//    - child: Widget.
//
func (assistant *Assistant) RemoveActionWidget(child Widgetter) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("remove_action_widget", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
}

// RemovePage removes the page_num’s page from assistant.
//
// The function takes the following parameters:
//
//    - pageNum: index of a page in the assistant, or -1 to remove the last page.
//
func (assistant *Assistant) RemovePage(pageNum int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.gint(pageNum)
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("remove_page", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetCurrentPage switches the page to page_num.
//
// Note that this will only be necessary in custom buttons, as the assistant
// flow can be set with gtk_assistant_set_forward_page_func().
//
// The function takes the following parameters:
//
//    - pageNum: index of the page to switch to, starting from 0. If negative,
//      the last page will be used. If greater than the number of pages in the
//      assistant, nothing will be done.
//
func (assistant *Assistant) SetCurrentPage(pageNum int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.gint(pageNum)
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_current_page", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetPageComplete sets whether page contents are complete.
//
// This will make assistant update the buttons state to be able to continue the
// task.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - complete completeness status of the page.
//
func (assistant *Assistant) SetPageComplete(page Widgetter, complete bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if complete {
		_arg2 = C.TRUE
	}
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_complete", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(complete)
}

// SetPageHasPadding sets whether the assistant is adding padding around the
// page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - hasPadding: whether this page has padding.
//
func (assistant *Assistant) SetPageHasPadding(page Widgetter, hasPadding bool) {
	var args [3]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if hasPadding {
		_arg2 = C.TRUE
	}
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_has_padding", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(hasPadding)
}

// SetPageHeaderImage sets a header image for page.
//
// Deprecated: Since GTK+ 3.2, a header is no longer shown; add your header
// decoration to the page content instead.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - pixbuf (optional): new header image page.
//
func (assistant *Assistant) SetPageHeaderImage(page Widgetter, pixbuf *gdkpixbuf.Pixbuf) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_header_image", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(pixbuf)
}

// SetPageSideImage sets a side image for page.
//
// This image used to be displayed in the side area of the assistant when page
// is the current page.
//
// Deprecated: Since GTK+ 3.2, sidebar images are not shown anymore.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - pixbuf (optional): new side image page.
//
func (assistant *Assistant) SetPageSideImage(page Widgetter, pixbuf *gdkpixbuf.Pixbuf) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		_arg2 = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_side_image", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(pixbuf)
}

// SetPageTitle sets a title for page.
//
// The title is displayed in the header area of the assistant when page is the
// current page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - title: new title for page.
//
func (assistant *Assistant) SetPageTitle(page Widgetter, title string) {
	var args [3]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _arg2 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))
	*(**Assistant)(unsafe.Pointer(&args[1])) = _arg1
	*(*Widgetter)(unsafe.Pointer(&args[2])) = _arg2

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("set_page_title", args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(title)
}

// UpdateButtonsState forces assistant to recompute the buttons state.
//
// GTK+ automatically takes care of this in most situations, e.g. when the user
// goes to a different page, or when the visibility or completeness of a page
// changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (assistant *Assistant) UpdateButtonsState() {
	var args [1]girepository.Argument
	var _arg0 *C.void // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**Assistant)(unsafe.Pointer(&args[0])) = _arg0

	girepository.MustFind("Gtk", "Assistant").InvokeMethod("update_buttons_state", args[:], nil)

	runtime.KeepAlive(assistant)
}
