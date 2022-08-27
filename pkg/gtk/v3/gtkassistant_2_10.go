// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void callbackDelete(gpointer);
// extern gint _gotk4_gtk3_AssistantPageFunc(gint, gpointer);
import "C"

// NewAssistant creates a new Assistant.
//
// The function returns the following values:
//
//    - assistant: newly created Assistant.
//
func NewAssistant() *Assistant {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_assistant_new()

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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_assistant_add_action_widget(_arg0, _arg1)
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
func (assistant *Assistant) AppendPage(page Widgetter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_append_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// CurrentPage returns the page number of the current page.
//
// The function returns the following values:
//
//    - gint: index (starting from 0) of the current page in the assistant, or -1
//      if the assistant has no pages, or no current page.
//
func (assistant *Assistant) CurrentPage() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_cret = C.gtk_assistant_get_current_page(_arg0)
	runtime.KeepAlive(assistant)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NPages returns the number of pages in the assistant.
//
// The function returns the following values:
//
//    - gint: number of pages in the assistant.
//
func (assistant *Assistant) NPages() int {
	var _arg0 *C.GtkAssistant // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	_cret = C.gtk_assistant_get_n_pages(_arg0)
	runtime.KeepAlive(assistant)

	var _gint int // out

	_gint = int(_cret)

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
func (assistant *Assistant) NthPage(pageNum int) Widgetter {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out
	var _cret *C.GtkWidget    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.gint(pageNum)

	_cret = C.gtk_assistant_get_nth_page(_arg0, _arg1)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_get_page_complete(_arg0, _arg1)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_get_page_header_image(_arg0, _arg1)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_get_page_side_image(_arg0, _arg1)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_get_page_title(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageType gets the page type of page.
//
// The function takes the following parameters:
//
//    - page of assistant.
//
// The function returns the following values:
//
//    - assistantPageType: page type of page.
//
func (assistant *Assistant) PageType(page Widgetter) AssistantPageType {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _cret C.GtkAssistantPageType // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_get_page_type(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _assistantPageType AssistantPageType // out

	_assistantPageType = AssistantPageType(_cret)

	return _assistantPageType
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
func (assistant *Assistant) InsertPage(page Widgetter, position int) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = C.gint(position)

	_cret = C.gtk_assistant_insert_page(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(position)

	var _gint int // out

	_gint = int(_cret)

	return _gint
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
func (assistant *Assistant) PrependPage(page Widgetter) int {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _cret C.gint          // in

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))

	_cret = C.gtk_assistant_prepend_page(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
//
// The function takes the following parameters:
//
//    - child: Widget.
//
func (assistant *Assistant) RemoveActionWidget(child Widgetter) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(child).Native()))

	C.gtk_assistant_remove_action_widget(_arg0, _arg1)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(child)
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
func (assistant *Assistant) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkAssistant // out
	var _arg1 C.gint          // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(_arg0, _arg1)
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
	var _arg0 *C.GtkAssistant        // out
	var _arg1 C.GtkAssistantPageFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	if pageFunc != nil {
		_arg1 = (*[0]byte)(C._gotk4_gtk3_AssistantPageFunc)
		_arg2 = C.gpointer(gbox.Assign(pageFunc))
		_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	C.gtk_assistant_set_forward_page_func(_arg0, _arg1, _arg2, _arg3)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 C.gboolean      // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if complete {
		_arg2 = C.TRUE
	}

	C.gtk_assistant_set_page_complete(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(complete)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	C.gtk_assistant_set_page_header_image(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	if pixbuf != nil {
		_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))
	}

	C.gtk_assistant_set_page_side_image(_arg0, _arg1, _arg2)
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
	var _arg0 *C.GtkAssistant // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.gchar        // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(title)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_assistant_set_page_title(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(title)
}

// SetPageType sets the page type for page.
//
// The page type determines the page behavior in the assistant.
//
// The function takes the following parameters:
//
//    - page of assistant.
//    - typ: new type for page.
//
func (assistant *Assistant) SetPageType(page Widgetter, typ AssistantPageType) {
	var _arg0 *C.GtkAssistant        // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 C.GtkAssistantPageType // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(page).Native()))
	_arg2 = C.GtkAssistantPageType(typ)

	C.gtk_assistant_set_page_type(_arg0, _arg1, _arg2)
	runtime.KeepAlive(assistant)
	runtime.KeepAlive(page)
	runtime.KeepAlive(typ)
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
	var _arg0 *C.GtkAssistant // out

	_arg0 = (*C.GtkAssistant)(unsafe.Pointer(coreglib.InternObject(assistant).Native()))

	C.gtk_assistant_update_buttons_state(_arg0)
	runtime.KeepAlive(assistant)
}
