// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_ClipboardImageReceivedFunc(GtkClipboard*, GdkPixbuf*, gpointer);
// extern void _gotk4_gtk3_ClipboardTextReceivedFunc(GtkClipboard*, gchar*, gpointer);
// extern void _gotk4_gtk3_ClipboardURIReceivedFunc(GtkClipboard*, gchar**, gpointer);
// extern void _gotk4_gtk3_Clipboard_ConnectOwnerChange(gpointer, GdkEventOwnerChange*, guintptr);
import "C"

// glib.Type values for gtkclipboard.go.
var GTypeClipboard = externglib.Type(C.gtk_clipboard_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeClipboard, F: marshalClipboard},
	})
}

// ClipboardImageReceivedFunc: function to be called when the results of
// gtk_clipboard_request_image() are received, or when the request fails.
type ClipboardImageReceivedFunc func(clipboard *Clipboard, pixbuf *gdkpixbuf.Pixbuf)

//export _gotk4_gtk3_ClipboardImageReceivedFunc
func _gotk4_gtk3_ClipboardImageReceivedFunc(arg1 *C.GtkClipboard, arg2 *C.GdkPixbuf, arg3 C.gpointer) {
	var fn ClipboardImageReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardImageReceivedFunc)
	}

	var _clipboard *Clipboard     // out
	var _pixbuf *gdkpixbuf.Pixbuf // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg1)))
	{
		obj := externglib.Take(unsafe.Pointer(arg2))
		_pixbuf = &gdkpixbuf.Pixbuf{
			Object: obj,
			LoadableIcon: gio.LoadableIcon{
				Icon: gio.Icon{
					Object: obj,
				},
			},
		}
	}

	fn(_clipboard, _pixbuf)
}

// ClipboardReceivedFunc: function to be called when the results of
// gtk_clipboard_request_contents() are received, or when the request fails.
type ClipboardReceivedFunc func(clipboard *Clipboard, selectionData *SelectionData)

//export _gotk4_gtk3_ClipboardReceivedFunc
func _gotk4_gtk3_ClipboardReceivedFunc(arg1 *C.GtkClipboard, arg2 *C.GtkSelectionData, arg3 C.gpointer) {
	var fn ClipboardReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardReceivedFunc)
	}

	var _clipboard *Clipboard         // out
	var _selectionData *SelectionData // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn(_clipboard, _selectionData)
}

// ClipboardTextReceivedFunc: function to be called when the results of
// gtk_clipboard_request_text() are received, or when the request fails.
type ClipboardTextReceivedFunc func(clipboard *Clipboard, text string)

//export _gotk4_gtk3_ClipboardTextReceivedFunc
func _gotk4_gtk3_ClipboardTextReceivedFunc(arg1 *C.GtkClipboard, arg2 *C.gchar, arg3 C.gpointer) {
	var fn ClipboardTextReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardTextReceivedFunc)
	}

	var _clipboard *Clipboard // out
	var _text string          // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	}

	fn(_clipboard, _text)
}

// ClipboardURIReceivedFunc: function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard *Clipboard, uris []string)

//export _gotk4_gtk3_ClipboardURIReceivedFunc
func _gotk4_gtk3_ClipboardURIReceivedFunc(arg1 *C.GtkClipboard, arg2 **C.gchar, arg3 C.gpointer) {
	var fn ClipboardURIReceivedFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ClipboardURIReceivedFunc)
	}

	var _clipboard *Clipboard // out
	var _uris []string        // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(arg1)))
	{
		var i int
		var z *C.gchar
		for p := arg2; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(arg2, i)
		_uris = make([]string, i)
		for i := range src {
			_uris[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	fn(_clipboard, _uris)
}

// Clipboard object represents a clipboard of data shared between different
// processes or between different widgets in the same process. Each clipboard is
// identified by a name encoded as a Atom. (Conversion to and from strings can
// be done with gdk_atom_intern() and gdk_atom_name().) The default clipboard
// corresponds to the “CLIPBOARD” atom; another commonly used clipboard is the
// “PRIMARY” clipboard, which, in X, traditionally contains the currently
// selected text.
//
// To support having a number of different formats on the clipboard at the same
// time, the clipboard mechanism allows providing callbacks instead of the
// actual data. When you set the contents of the clipboard, you can either
// supply the data directly (via functions like gtk_clipboard_set_text()), or
// you can supply a callback to be called at a later time when the data is
// needed (via gtk_clipboard_set_with_data() or gtk_clipboard_set_with_owner().)
// Providing a callback also avoids having to make copies of the data when it is
// not needed.
//
// gtk_clipboard_set_with_data() and gtk_clipboard_set_with_owner() are quite
// similar; the choice between the two depends mostly on which is more
// convenient in a particular situation. The former is most useful when you want
// to have a blob of data with callbacks to convert it into the various data
// types that you advertise. When the clear_func you provided is called, you
// simply free the data blob. The latter is more useful when the contents of
// clipboard reflect the internal state of a #GObject (As an example, for the
// PRIMARY clipboard, when an entry widget provides the clipboard’s contents the
// contents are simply the text within the selected region.) If the contents
// change, the entry widget can call gtk_clipboard_set_with_owner() to update
// the timestamp for clipboard ownership, without having to worry about
// clear_func being called.
//
// Requesting the data from the clipboard is essentially asynchronous. If the
// contents of the clipboard are provided within the same process, then a direct
// function call will be made to retrieve the data, but if they are provided by
// another process, then the data needs to be retrieved from the other process,
// which may take some time. To avoid blocking the user interface, the call to
// request the selection, gtk_clipboard_request_contents() takes a callback that
// will be called when the contents are received (or when the request fails.) If
// you don’t want to deal with providing a separate callback, you can also use
// gtk_clipboard_wait_for_contents(). What this does is run the GLib main loop
// recursively waiting for the contents. This can simplify the code flow, but
// you still have to be aware that other callbacks in your program can be called
// while this recursive mainloop is running.
//
// Along with the functions to get the clipboard contents as an arbitrary data
// chunk, there are also functions to retrieve it as text,
// gtk_clipboard_request_text() and gtk_clipboard_wait_for_text(). These
// functions take care of determining which formats are advertised by the
// clipboard provider, asking for the clipboard in the best available format and
// converting the results into the UTF-8 encoding. (The standard form for
// representing strings in GTK+.).
type Clipboard struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Clipboard)(nil)
)

func wrapClipboard(obj *externglib.Object) *Clipboard {
	return &Clipboard{
		Object: obj,
	}
}

func marshalClipboard(p uintptr) (interface{}, error) {
	return wrapClipboard(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Clipboard_ConnectOwnerChange
func _gotk4_gtk3_Clipboard_ConnectOwnerChange(arg0 C.gpointer, arg1 *C.GdkEventOwnerChange, arg2 C.guintptr) {
	var f func(event *gdk.EventOwnerChange)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(event *gdk.EventOwnerChange))
	}

	var _event *gdk.EventOwnerChange // out

	_event = (*gdk.EventOwnerChange)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_event)
}

// ConnectOwnerChange signal is emitted when GTK+ receives an event that
// indicates that the ownership of the selection associated with clipboard has
// changed.
func (clipboard *Clipboard) ConnectOwnerChange(f func(event *gdk.EventOwnerChange)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(clipboard, "owner-change", false, unsafe.Pointer(C._gotk4_gtk3_Clipboard_ConnectOwnerChange), f)
}

// Clear clears the contents of the clipboard. Generally this should only be
// called between the time you call gtk_clipboard_set_with_owner() or
// gtk_clipboard_set_with_data(), and when the clear_func you supplied is
// called. Otherwise, the clipboard may be owned by someone else.
func (clipboard *Clipboard) Clear() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	C.gtk_clipboard_clear(_arg0)
	runtime.KeepAlive(clipboard)
}

// Display gets the Display associated with clipboard.
//
// The function returns the following values:
//
//    - display associated with clipboard.
//
func (clipboard *Clipboard) Display() *gdk.Display {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_get_display(_arg0)
	runtime.KeepAlive(clipboard)

	var _display *gdk.Display // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_display = &gdk.Display{
			Object: obj,
		}
	}

	return _display
}

// Owner: if the clipboard contents callbacks were set with
// gtk_clipboard_set_with_owner(), and the gtk_clipboard_set_with_data() or
// gtk_clipboard_clear() has not subsequently called, returns the owner set by
// gtk_clipboard_set_with_owner().
//
// The function returns the following values:
//
//    - object (optional): owner of the clipboard, if any; otherwise NULL.
//
func (clipboard *Clipboard) Owner() *externglib.Object {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_get_owner(_arg0)
	runtime.KeepAlive(clipboard)

	var _object *externglib.Object // out

	if _cret != nil {
		_object = externglib.Take(unsafe.Pointer(_cret))
	}

	return _object
}

// RequestImage requests the contents of the clipboard as image. When the image
// is later received, it will be converted to a Pixbuf, and callback will be
// called.
//
// The pixbuf parameter to callback will contain the resulting Pixbuf if the
// request succeeded, or NULL if it failed. This could happen for various
// reasons, in particular if the clipboard was empty or if the contents of the
// clipboard could not be converted into an image.
//
// The function takes the following parameters:
//
//    - callback: function to call when the image is received, or the retrieval
//      fails. (It will always be called one way or the other.).
//
func (clipboard *Clipboard) RequestImage(callback ClipboardImageReceivedFunc) {
	var _arg0 *C.GtkClipboard                 // out
	var _arg1 C.GtkClipboardImageReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardImageReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_image(_arg0, _arg1, _arg2)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(callback)
}

// RequestText requests the contents of the clipboard as text. When the text is
// later received, it will be converted to UTF-8 if necessary, and callback will
// be called.
//
// The text parameter to callback will contain the resulting text if the request
// succeeded, or NULL if it failed. This could happen for various reasons, in
// particular if the clipboard was empty or if the contents of the clipboard
// could not be converted into text form.
//
// The function takes the following parameters:
//
//    - callback: function to call when the text is received, or the retrieval
//      fails. (It will always be called one way or the other.).
//
func (clipboard *Clipboard) RequestText(callback ClipboardTextReceivedFunc) {
	var _arg0 *C.GtkClipboard                // out
	var _arg1 C.GtkClipboardTextReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardTextReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(callback)
}

// RequestURIs requests the contents of the clipboard as URIs. When the URIs are
// later received callback will be called.
//
// The uris parameter to callback will contain the resulting array of URIs if
// the request succeeded, or NULL if it failed. This could happen for various
// reasons, in particular if the clipboard was empty or if the contents of the
// clipboard could not be converted into URI form.
//
// The function takes the following parameters:
//
//    - callback: function to call when the URIs are received, or the retrieval
//      fails. (It will always be called one way or the other.).
//
func (clipboard *Clipboard) RequestURIs(callback ClipboardURIReceivedFunc) {
	var _arg0 *C.GtkClipboard               // out
	var _arg1 C.GtkClipboardURIReceivedFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk3_ClipboardURIReceivedFunc)
	_arg2 = C.gpointer(gbox.AssignOnce(callback))

	C.gtk_clipboard_request_uris(_arg0, _arg1, _arg2)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(callback)
}

// SetCanStore hints that the clipboard data should be stored somewhere when the
// application exits or when gtk_clipboard_store () is called.
//
// This value is reset when the clipboard owner changes. Where the clipboard
// data is stored is platform dependent, see gdk_display_store_clipboard () for
// more information.
//
// The function takes the following parameters:
//
//    - targets (optional): array containing information about which forms should
//      be stored or NULL to indicate that all forms should be stored.
//
func (clipboard *Clipboard) SetCanStore(targets []TargetEntry) {
	var _arg0 *C.GtkClipboard   // out
	var _arg1 *C.GtkTargetEntry // out
	var _arg2 C.gint

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg2 = (C.gint)(len(targets))
	_arg1 = (*C.GtkTargetEntry)(C.calloc(C.size_t(len(targets)), C.size_t(C.sizeof_GtkTargetEntry)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GtkTargetEntry)(_arg1), len(targets))
		for i := range targets {
			out[i] = *(*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer((&targets[i]))))
		}
	}

	C.gtk_clipboard_set_can_store(_arg0, _arg1, _arg2)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(targets)
}

// SetImage sets the contents of the clipboard to the given Pixbuf. GTK+ will
// take responsibility for responding for requests for the image, and for
// converting the image into the requested format.
//
// The function takes the following parameters:
//
//    - pixbuf: Pixbuf.
//
func (clipboard *Clipboard) SetImage(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.GdkPixbuf    // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(externglib.InternObject(pixbuf).Native()))

	C.gtk_clipboard_set_image(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(pixbuf)
}

// SetText sets the contents of the clipboard to the given UTF-8 string. GTK+
// will make a copy of the text and take responsibility for responding for
// requests for the text, and for converting the text into the requested format.
//
// The function takes the following parameters:
//
//    - text: UTF-8 string.
//    - len: length of text, in bytes, or -1, in which case the length will be
//      determined with strlen().
//
func (clipboard *Clipboard) SetText(text string, len int) {
	var _arg0 *C.GtkClipboard // out
	var _arg1 *C.gchar        // out
	var _arg2 C.gint          // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(len)

	C.gtk_clipboard_set_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(text)
	runtime.KeepAlive(len)
}

// Store stores the current clipboard data somewhere so that it will stay around
// after the application has quit.
func (clipboard *Clipboard) Store() {
	var _arg0 *C.GtkClipboard // out

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	C.gtk_clipboard_store(_arg0)
	runtime.KeepAlive(clipboard)
}

// WaitForImage requests the contents of the clipboard as image and converts the
// result to a Pixbuf. This function waits for the data to be received using the
// main loop, so events, timeouts, etc, may be dispatched during the wait.
//
// The function returns the following values:
//
//    - pixbuf (optional): newly-allocated Pixbuf object which must be disposed
//      with g_object_unref(), or NULL if retrieving the selection data failed.
//      (This could happen for various reasons, in particular if the clipboard
//      was empty or if the contents of the clipboard could not be converted into
//      an image.).
//
func (clipboard *Clipboard) WaitForImage() *gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.GdkPixbuf    // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_for_image(_arg0)
	runtime.KeepAlive(clipboard)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// WaitForText requests the contents of the clipboard as text and converts the
// result to UTF-8 if necessary. This function waits for the data to be received
// using the main loop, so events, timeouts, etc, may be dispatched during the
// wait.
//
// The function returns the following values:
//
//    - utf8 (optional): newly-allocated UTF-8 string which must be freed with
//      g_free(), or NULL if retrieving the selection data failed. (This could
//      happen for various reasons, in particular if the clipboard was empty or
//      if the contents of the clipboard could not be converted into text form.).
//
func (clipboard *Clipboard) WaitForText() string {
	var _arg0 *C.GtkClipboard // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_for_text(_arg0)
	runtime.KeepAlive(clipboard)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// WaitForURIs requests the contents of the clipboard as URIs. This function
// waits for the data to be received using the main loop, so events, timeouts,
// etc, may be dispatched during the wait.
//
// The function returns the following values:
//
//    - utf8s (optional): a newly-allocated NULL-terminated array of strings
//      which must be freed with g_strfreev(), or NULL if retrieving the
//      selection data failed. (This could happen for various reasons, in
//      particular if the clipboard was empty or if the contents of the clipboard
//      could not be converted into URI form.).
//
func (clipboard *Clipboard) WaitForURIs() []string {
	var _arg0 *C.GtkClipboard // out
	var _cret **C.gchar       // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_for_uris(_arg0)
	runtime.KeepAlive(clipboard)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// WaitIsImageAvailable: test to see if there is an image available to be pasted
// This is done by requesting the TARGETS atom and checking if it contains any
// of the supported image targets. This function waits for the data to be
// received using the main loop, so events, timeouts, etc, may be dispatched
// during the wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_image()
// since it doesn’t need to retrieve the actual image data.
//
// The function returns the following values:
//
//    - ok: TRUE is there is an image available, FALSE otherwise.
//
func (clipboard *Clipboard) WaitIsImageAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_is_image_available(_arg0)
	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsRichTextAvailable: test to see if there is rich text available to be
// pasted This is done by requesting the TARGETS atom and checking if it
// contains any of the supported rich text targets. This function waits for the
// data to be received using the main loop, so events, timeouts, etc, may be
// dispatched during the wait.
//
// This function is a little faster than calling
// gtk_clipboard_wait_for_rich_text() since it doesn’t need to retrieve the
// actual text.
//
// The function takes the following parameters:
//
//    - buffer: TextBuffer.
//
// The function returns the following values:
//
//    - ok: TRUE is there is rich text available, FALSE otherwise.
//
func (clipboard *Clipboard) WaitIsRichTextAvailable(buffer *TextBuffer) bool {
	var _arg0 *C.GtkClipboard  // out
	var _arg1 *C.GtkTextBuffer // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GtkTextBuffer)(unsafe.Pointer(externglib.InternObject(buffer).Native()))

	_cret = C.gtk_clipboard_wait_is_rich_text_available(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsTextAvailable: test to see if there is text available to be pasted This
// is done by requesting the TARGETS atom and checking if it contains any of the
// supported text targets. This function waits for the data to be received using
// the main loop, so events, timeouts, etc, may be dispatched during the wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_text()
// since it doesn’t need to retrieve the actual text.
//
// The function returns the following values:
//
//    - ok: TRUE is there is text available, FALSE otherwise.
//
func (clipboard *Clipboard) WaitIsTextAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_is_text_available(_arg0)
	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WaitIsURIsAvailable: test to see if there is a list of URIs available to be
// pasted This is done by requesting the TARGETS atom and checking if it
// contains the URI targets. This function waits for the data to be received
// using the main loop, so events, timeouts, etc, may be dispatched during the
// wait.
//
// This function is a little faster than calling gtk_clipboard_wait_for_uris()
// since it doesn’t need to retrieve the actual URI data.
//
// The function returns the following values:
//
//    - ok: TRUE is there is an URI list available, FALSE otherwise.
//
func (clipboard *Clipboard) WaitIsURIsAvailable() bool {
	var _arg0 *C.GtkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gtk_clipboard_wait_is_uris_available(_arg0)
	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ClipboardGetDefault returns the default clipboard object for use with
// cut/copy/paste menu items and keyboard shortcuts.
//
// The function takes the following parameters:
//
//    - display for which the clipboard is to be retrieved.
//
// The function returns the following values:
//
//    - clipboard: default clipboard object.
//
func ClipboardGetDefault(display *gdk.Display) *Clipboard {
	var _arg1 *C.GdkDisplay   // out
	var _cret *C.GtkClipboard // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(externglib.InternObject(display).Native()))

	_cret = C.gtk_clipboard_get_default(_arg1)
	runtime.KeepAlive(display)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(externglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}
