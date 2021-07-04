// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_assistant_page_type_get_type()), F: marshalAssistantPageType},
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistant},
	})
}

// AssistantPageType: an enum for determining the page role inside the
// Assistant. It's used to handle buttons sensitivity and visibility.
//
// Note that an assistant needs to end its page flow with a page of type
// GTK_ASSISTANT_PAGE_CONFIRM, GTK_ASSISTANT_PAGE_SUMMARY or
// GTK_ASSISTANT_PAGE_PROGRESS to be correct.
//
// The Cancel button will only be shown if the page isn’t “committed”. See
// gtk_assistant_commit() for details.
type AssistantPageType int

const (
	// content: the page has regular contents. Both the Back and forward buttons
	// will be shown.
	AssistantPageTypeContent AssistantPageType = 0
	// intro: the page contains an introduction to the assistant task. Only the
	// Forward button will be shown if there is a next page.
	AssistantPageTypeIntro AssistantPageType = 1
	// confirm: the page lets the user confirm or deny the changes. The Back and
	// Apply buttons will be shown.
	AssistantPageTypeConfirm AssistantPageType = 2
	// summary: the page informs the user of the changes done. Only the Close
	// button will be shown.
	AssistantPageTypeSummary AssistantPageType = 3
	// progress: used for tasks that take a long time to complete, blocks the
	// assistant until the page is marked as complete. Only the back button will
	// be shown.
	AssistantPageTypeProgress AssistantPageType = 4
	// custom: used for when other page types are not appropriate. No buttons
	// will be shown, and the application must add its own buttons through
	// gtk_assistant_add_action_widget().
	AssistantPageTypeCustom AssistantPageType = 5
)

func marshalAssistantPageType(p uintptr) (interface{}, error) {
	return AssistantPageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// AssistantPageFunc: a function used by gtk_assistant_set_forward_page_func()
// to know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int, gint int)

//export gotk4_AssistantPageFunc
func _AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) C.gint {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var currentPage int // out

	currentPage = int(arg0)

	fn := v.(AssistantPageFunc)
	gint := fn(currentPage)

	var cret C.gint // out

	cret = C.gint(gint)

	return cret
}

// Assistant: a Assistant is a widget used to represent a generally complex
// operation splitted in several steps, guiding the user through its pages and
// controlling the page flow to collect the necessary data.
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
// @action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant interface {
	Window

	AddActionWidgetAssistant(child Widget)

	AppendPageAssistant(page Widget) int

	CommitAssistant()

	CurrentPage() int

	NPages() int

	NthPage(pageNum int) Widget

	PageComplete(page Widget) bool

	PageHasPadding(page Widget) bool

	PageHeaderImage(page Widget) gdkpixbuf.Pixbuf

	PageSideImage(page Widget) gdkpixbuf.Pixbuf

	PageTitle(page Widget) string

	PageType(page Widget) AssistantPageType

	InsertPageAssistant(page Widget, position int) int

	NextPageAssistant()

	PrependPageAssistant(page Widget) int

	PreviousPageAssistant()

	RemoveActionWidgetAssistant(child Widget)

	RemovePageAssistant(pageNum int)

	SetCurrentPageAssistant(pageNum int)

	SetPageCompleteAssistant(page Widget, complete bool)

	SetPageHasPaddingAssistant(page Widget, hasPadding bool)

	SetPageHeaderImageAssistant(page Widget, pixbuf gdkpixbuf.Pixbuf)

	SetPageSideImageAssistant(page Widget, pixbuf gdkpixbuf.Pixbuf)

	SetPageTitleAssistant(page Widget, title string)

	SetPageTypeAssistant(page Widget, typ AssistantPageType)

	UpdateButtonsStateAssistant()
}

// assistant implements the Assistant class.
type assistant struct {
	Window
}

// WrapAssistant wraps a GObject to the right type. It is
// primarily used internally.
func WrapAssistant(obj *externglib.Object) Assistant {
	return assistant{
		Window: WrapWindow(obj),
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAssistant(obj), nil
}

func NewAssistant() Assistant {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_assistant_new()

	var _assistant Assistant // out

	_assistant = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Assistant)

	return _assistant
}

func (a assistant) AddActionWidgetAssistant(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
}

func (a assistant) AppendPageAssistant(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a assistant) CommitAssistant() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_commit(_arg0)
}

func (a assistant) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_assistant_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a assistant) NPages() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_assistant_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a assistant) NthPage(pageNum int) Widget {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_assistant_get_nth_page(_arg0, _arg1)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (a assistant) PageComplete(page Widget) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a assistant) PageHasPadding(page Widget) bool {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_has_padding(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a assistant) PageHeaderImage(page Widget) gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_header_image(_arg0, _arg1)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (a assistant) PageSideImage(page Widget) gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_side_image(_arg0, _arg1)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (a assistant) PageTitle(page Widget) string {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (a assistant) PageType(page Widget) AssistantPageType {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _cret C.GtkAssistantPageType // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_get_page_type(_arg0, _arg1)

	var _assistantPageType AssistantPageType // out

	_assistantPageType = AssistantPageType(_cret)

	return _assistantPageType
}

func (a assistant) InsertPageAssistant(page Widget, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = C.gint(position)

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a assistant) NextPageAssistant() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_next_page(_arg0)
}

func (a assistant) PrependPageAssistant(page Widget) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a assistant) PreviousPageAssistant() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_previous_page(_arg0)
}

func (a assistant) RemoveActionWidgetAssistant(child Widget) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
}

func (a assistant) RemovePageAssistant(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_remove_page(_arg0, _arg1)
}

func (a assistant) SetCurrentPageAssistant(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
}

func (a assistant) SetPageCompleteAssistant(page Widget, complete bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
}

func (a assistant) SetPageHasPaddingAssistant(page Widget, hasPadding bool) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if hasPadding {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_has_padding(_arg0, _arg1, _arg2)
}

func (a assistant) SetPageHeaderImageAssistant(page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_header_image(_arg0, _arg1, _arg2)
}

func (a assistant) SetPageSideImageAssistant(page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_side_image(_arg0, _arg1, _arg2)
}

func (a assistant) SetPageTitleAssistant(page Widget, title string) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
}

func (a assistant) SetPageTypeAssistant(page Widget, typ AssistantPageType) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkAssistantPageType // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	_arg2 = C.GtkAssistantPageType(typ)

	C.gtk_assistant_set_page_type(_arg0, _arg1, _arg2)
}

func (a assistant) UpdateButtonsStateAssistant() {
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
}

func (b assistant) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b assistant) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b assistant) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b assistant) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b assistant) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b assistant) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b assistant) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b assistant) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b assistant) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b assistant) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
