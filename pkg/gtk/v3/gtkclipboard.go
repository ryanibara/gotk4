// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern void _gotk4_gtk3_ClipboardImageReceivedFunc(void*, void*, gpointer);
// extern void _gotk4_gtk3_ClipboardTextReceivedFunc(void*, void*, gpointer);
// extern void _gotk4_gtk3_ClipboardURIReceivedFunc(void*, void**, gpointer);
// extern void _gotk4_gtk3_Clipboard_ConnectOwnerChange(gpointer, void*, guintptr);
import "C"

// glib.Type values for gtkclipboard.go.
var GTypeClipboard = coreglib.Type(C.gtk_clipboard_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeClipboard, F: marshalClipboard},
	})
}

// ClipboardImageReceivedFunc: function to be called when the results of
// gtk_clipboard_request_image() are received, or when the request fails.
type ClipboardImageReceivedFunc func(clipboard *Clipboard, pixbuf *gdkpixbuf.Pixbuf)

//export _gotk4_gtk3_ClipboardImageReceivedFunc
func _gotk4_gtk3_ClipboardImageReceivedFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
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

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	{
		obj := coreglib.Take(unsafe.Pointer(arg2))
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
func _gotk4_gtk3_ClipboardReceivedFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
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

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn(_clipboard, _selectionData)
}

// ClipboardTextReceivedFunc: function to be called when the results of
// gtk_clipboard_request_text() are received, or when the request fails.
type ClipboardTextReceivedFunc func(clipboard *Clipboard, text string)

//export _gotk4_gtk3_ClipboardTextReceivedFunc
func _gotk4_gtk3_ClipboardTextReceivedFunc(arg1 *C.void, arg2 *C.void, arg3 C.gpointer) {
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

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_text = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	}

	fn(_clipboard, _text)
}

// ClipboardURIReceivedFunc: function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard *Clipboard, uris []string)

//export _gotk4_gtk3_ClipboardURIReceivedFunc
func _gotk4_gtk3_ClipboardURIReceivedFunc(arg1 *C.void, arg2 **C.void, arg3 C.gpointer) {
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

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))
	{
		var i int
		var z *C.void
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
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Clipboard)(nil)
)

func wrapClipboard(obj *coreglib.Object) *Clipboard {
	return &Clipboard{
		Object: obj,
	}
}

func marshalClipboard(p uintptr) (interface{}, error) {
	return wrapClipboard(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_Clipboard_ConnectOwnerChange
func _gotk4_gtk3_Clipboard_ConnectOwnerChange(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) {
	var f func(event *gdk.EventOwnerChange)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
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
func (clipboard *Clipboard) ConnectOwnerChange(f func(event *gdk.EventOwnerChange)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(clipboard, "owner-change", false, unsafe.Pointer(C._gotk4_gtk3_Clipboard_ConnectOwnerChange), f)
}

// Clear clears the contents of the clipboard. Generally this should only be
// called between the time you call gtk_clipboard_set_with_owner() or
// gtk_clipboard_set_with_data(), and when the clear_func you supplied is
// called. Otherwise, the clipboard may be owned by someone else.
func (clipboard *Clipboard) Clear() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("clear", _args[:], nil)

	runtime.KeepAlive(clipboard)
}

// Display gets the Display associated with clipboard.
//
// The function returns the following values:
//
//    - display associated with clipboard.
//
func (clipboard *Clipboard) Display() *gdk.Display {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("get_display", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _display *gdk.Display // out

	{
		obj := coreglib.Take(unsafe.Pointer(_cret))
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
func (clipboard *Clipboard) Owner() *coreglib.Object {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("get_owner", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _object *coreglib.Object // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_object = coreglib.Take(unsafe.Pointer(_cret))
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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_ClipboardImageReceivedFunc)
	_args[2] = C.gpointer(gbox.AssignOnce(callback))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("request_image", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_ClipboardTextReceivedFunc)
	_args[2] = C.gpointer(gbox.AssignOnce(callback))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("request_text", _args[:], nil)

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
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_ClipboardURIReceivedFunc)
	_args[2] = C.gpointer(gbox.AssignOnce(callback))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("request_uris", _args[:], nil)

	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(callback)
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(pixbuf).Native()))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("set_image", _args[:], nil)

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
func (clipboard *Clipboard) SetText(text string, len int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(len)

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("set_text", _args[:], nil)

	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(text)
	runtime.KeepAlive(len)
}

// Store stores the current clipboard data somewhere so that it will stay around
// after the application has quit.
func (clipboard *Clipboard) Store() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	girepository.MustFind("Gtk", "Clipboard").InvokeMethod("store", _args[:], nil)

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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_for_image", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_for_text", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_for_uris", _args[:], nil)
	_cret = *(***C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _utf8s []string // out

	if *(***C.void)(unsafe.Pointer(&_cret)) != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.void
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_is_image_available", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(buffer).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_is_rich_text_available", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(buffer)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_is_text_available", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(clipboard).Native()))

	_gret := girepository.MustFind("Gtk", "Clipboard").InvokeMethod("wait_is_uris_available", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_gret := girepository.MustFind("Gtk", "get_default").Invoke(_args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(display)

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(_cret)))

	return _clipboard
}
