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
// #include <glib-object.h>
// extern gint _gotk4_gtk3_AssistantPageFunc(gint, gpointer);
// extern void _gotk4_gtk3_AssistantClass_apply(void*);
// extern void _gotk4_gtk3_AssistantClass_cancel(void*);
// extern void _gotk4_gtk3_AssistantClass_close(void*);
// extern void _gotk4_gtk3_AssistantClass_prepare(void*, void*);
// extern void _gotk4_gtk3_Assistant_ConnectApply(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectCancel(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectClose(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectEscape(gpointer, guintptr);
// extern void _gotk4_gtk3_Assistant_ConnectPrepare(gpointer, void*, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// GTypeAssistantPageType returns the GType for the type AssistantPageType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAssistantPageType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "AssistantPageType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAssistantPageType)
	return gtype
}

// GTypeAssistant returns the GType for the type Assistant.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeAssistant() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "Assistant").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalAssistant)
	return gtype
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
	pclass := girepository.MustFind("Gtk", "AssistantClass")

	if _, ok := goval.(interface{ Apply() }); ok {
		o := pclass.StructFieldOffset("apply")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_AssistantClass_apply)
	}

	if _, ok := goval.(interface{ Cancel() }); ok {
		o := pclass.StructFieldOffset("cancel")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_AssistantClass_cancel)
	}

	if _, ok := goval.(interface{ Close() }); ok {
		o := pclass.StructFieldOffset("close")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_AssistantClass_close)
	}

	if _, ok := goval.(interface{ Prepare(page Widgetter) }); ok {
		o := pclass.StructFieldOffset("prepare")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gtk3_AssistantClass_prepare)
	}
}

//export _gotk4_gtk3_AssistantClass_apply
func _gotk4_gtk3_AssistantClass_apply(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Apply() })

	iface.Apply()
}

//export _gotk4_gtk3_AssistantClass_cancel
func _gotk4_gtk3_AssistantClass_cancel(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Cancel() })

	iface.Cancel()
}

//export _gotk4_gtk3_AssistantClass_close
func _gotk4_gtk3_AssistantClass_close(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk3_AssistantClass_prepare
func _gotk4_gtk3_AssistantClass_prepare(arg0 *C.void, arg1 *C.void) {
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
func _gotk4_gtk3_Assistant_ConnectPrepare(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
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
	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("new_Assistant", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("add_action_widget", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("append_page", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("commit", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_current_page", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// NPages returns the number of pages in the assistant.
//
// The function returns the following values:
//
//    - gint: number of pages in the assistant.
//
func (assistant *Assistant) NPages() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_n_pages", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(pageNum)

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_nth_page", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)

	var _widget Widgetter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_page_complete", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_page_has_padding", _args[:], nil)
	_cret := *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_page_header_image", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_page_side_image", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("get_page_title", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(position)

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("insert_page", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(position)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next page.
//
// This function is for use when creating pages of the K_ASSISTANT_PAGE_CUSTOM
// type.
func (assistant *Assistant) NextPage() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("next_page", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_gret := _info.InvokeClassMethod("prepend_page", _args[:], nil)
	_cret := *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("previous_page", _args[:], nil)

	runtime.KeepAlive(assistant)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
//
// The function takes the following parameters:
//
//    - child: Widget.
//
func (assistant *Assistant) RemoveActionWidget(child Widgetter) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("remove_action_widget", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(pageNum)

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("remove_page", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(pageNum)

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_current_page", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageNum)
}

// SetForwardPageFunc sets the page forwarding function to be page_func.
//
// This function will be used to determine what will be the next page when the
// user presses the forward button. Setting page_func to NULL will make the
// assistant to use the default forward function, which just goes to the next
// visible page.
//
// The function takes the following parameters:
//
//    - pageFunc (optional) or NULL to use the default one.
//
func (assistant *Assistant) SetForwardPageFunc(pageFunc AssistantPageFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	if pageFunc != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_AssistantPageFunc)
		_args[2] = C.gpointer(gbox.Assign(pageFunc))
		_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_forward_page_func", _args[:], nil)

	runtime.KeepAlive(assistant)
	runtime.KeepAlive(pageFunc)
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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if complete {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_page_complete", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if hasPadding {
		*(*C.gboolean)(unsafe.Pointer(&_args[2])) = C.TRUE
	}

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_page_has_padding", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_page_header_image", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_page_side_image", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_args[2]))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("set_page_title", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_info := girepository.MustFind("Gtk", "Assistant")
	_info.InvokeClassMethod("update_buttons_state", _args[:], nil)

	runtime.KeepAlive(assistant)
}
