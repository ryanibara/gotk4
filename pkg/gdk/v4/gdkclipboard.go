// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk4_Clipboard_ConnectChanged(gpointer, guintptr);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

// glib.Type values for gdkclipboard.go.
var GTypeClipboard = externglib.Type(C.gdk_clipboard_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeClipboard, F: marshalClipboard},
	})
}

// Clipboard: GdkClipboard object represents data shared between applications or
// inside an application.
//
// To get a GdkClipboard object, use gdk.Display.GetClipboard() or
// gdk.Display.GetPrimaryClipboard(). You can find out about the data that is
// currently available in a clipboard using gdk.Clipboard.GetFormats().
//
// To make text or image data available in a clipboard, use
// gdk.Clipboard.SetText() or gdk.Clipboard.SetTexture(). For other data, you
// can use gdk.Clipboard.SetContent(), which takes a gdk.ContentProvider object.
//
// To read textual or image data from a clipboard, use
// gdk.Clipboard.ReadTextAsync() or gdk.Clipboard.ReadTextureAsync(). For other
// data, use gdk.Clipboard.ReadAsync(), which provides a GInputStream object.
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

//export _gotk4_gdk4_Clipboard_ConnectChanged
func _gotk4_gdk4_Clipboard_ConnectChanged(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectChanged is emitted when the clipboard changes ownership.
func (clipboard *Clipboard) ConnectChanged(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(clipboard, "changed", false, unsafe.Pointer(C._gotk4_gdk4_Clipboard_ConnectChanged), f)
}

// Content returns the GdkContentProvider currently set on clipboard.
//
// If the clipboard is empty or its contents are not owned by the current
// process, NULL will be returned.
//
// The function returns the following values:
//
//    - contentProvider (optional): content of a clipboard or NULL if the
//      clipboard does not maintain any content.
//
func (clipboard *Clipboard) Content() *ContentProvider {
	var _arg0 *C.GdkClipboard       // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gdk_clipboard_get_content(_arg0)
	runtime.KeepAlive(clipboard)

	var _contentProvider *ContentProvider // out

	if _cret != nil {
		_contentProvider = wrapContentProvider(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _contentProvider
}

// Display gets the GdkDisplay that the clipboard was created for.
//
// The function returns the following values:
//
//    - display: GdkDisplay.
//
func (clipboard *Clipboard) Display() *Display {
	var _arg0 *C.GdkClipboard // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gdk_clipboard_get_display(_arg0)
	runtime.KeepAlive(clipboard)

	var _display *Display // out

	_display = wrapDisplay(externglib.Take(unsafe.Pointer(_cret)))

	return _display
}

// Formats gets the formats that the clipboard can provide its current contents
// in.
//
// The function returns the following values:
//
//    - contentFormats formats of the clipboard.
//
func (clipboard *Clipboard) Formats() *ContentFormats {
	var _arg0 *C.GdkClipboard      // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gdk_clipboard_get_formats(_arg0)
	runtime.KeepAlive(clipboard)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.gdk_content_formats_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// IsLocal returns if the clipboard is local.
//
// A clipboard is considered local if it was last claimed by the running
// application.
//
// Note that gdk.Clipboard.GetContent() may return NULL even on a local
// clipboard. In this case the clipboard is empty.
//
// The function returns the following values:
//
//    - ok: TRUE if the clipboard is local.
//
func (clipboard *Clipboard) IsLocal() bool {
	var _arg0 *C.GdkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))

	_cret = C.gdk_clipboard_is_local(_arg0)
	runtime.KeepAlive(clipboard)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReadAsync: asynchronously requests an input stream to read the clipboard's
// contents from.
//
// When the operation is finished callback will be called. You must then call
// gdk.Clipboard.ReadFinish() to get the result of the operation.
//
// The clipboard will choose the most suitable mime type from the given list to
// fulfill the request, preferring the ones listed first.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - mimeTypes: NULL-terminated array of mime types to choose from.
//    - ioPriority: i/O priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (clipboard *Clipboard) ReadAsync(ctx context.Context, mimeTypes []string, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkClipboard       // out
	var _arg3 *C.GCancellable       // out
	var _arg1 **C.char              // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	{
		_arg1 = (**C.char)(C.calloc(C.size_t((len(mimeTypes) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(mimeTypes)+1)
			var zero *C.char
			out[len(mimeTypes)] = zero
			for i := range mimeTypes {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(mimeTypes[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_clipboard_read_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(mimeTypes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadFinish finishes an asynchronous clipboard read.
//
// See gdk.Clipboard.ReadAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - outMimeType (optional): pointer to store the chosen mime type in or NULL.
//    - inputStream (optional): GInputStream or NULL on error.
//
func (clipboard *Clipboard) ReadFinish(result gio.AsyncResulter) (string, gio.InputStreamer, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.char         // in
	var _cret *C.GInputStream // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.gdk_clipboard_read_finish(_arg0, _arg1, &_arg2, &_cerr)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(result)

	var _outMimeType string            // out
	var _inputStream gio.InputStreamer // out
	var _goerr error                   // out

	if _arg2 != nil {
		_outMimeType = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	}
	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.InputStreamer)
				return ok
			})
			rv, ok := casted.(gio.InputStreamer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.InputStreamer")
			}
			_inputStream = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _outMimeType, _inputStream, _goerr
}

// ReadTextAsync: asynchronously request the clipboard contents converted to a
// string.
//
// When the operation is finished callback will be called. You must then call
// gdk.Clipboard.ReadTextFinish() to get the result.
//
// This is a simple wrapper around gdk.Clipboard.ReadValueAsync(). Use that
// function or gdk.Clipboard.ReadAsync() directly if you need more control over
// the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (clipboard *Clipboard) ReadTextAsync(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkClipboard       // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_clipboard_read_text_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ReadTextFinish finishes an asynchronous clipboard read.
//
// See gdk.Clipboard.ReadTextAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - utf8 (optional): new string or NULL on error.
//
func (clipboard *Clipboard) ReadTextFinish(result gio.AsyncResulter) (string, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.char         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.gdk_clipboard_read_text_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(result)

	var _utf8 string // out
	var _goerr error // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// ReadTextureAsync: asynchronously request the clipboard contents converted to
// a GdkPixbuf.
//
// When the operation is finished callback will be called. You must then call
// gdk.Clipboard.ReadTextureFinish() to get the result.
//
// This is a simple wrapper around gdk.Clipboard.ReadValueAsync(). Use that
// function or gdk.Clipboard.ReadAsync directly if you need more control over
// the operation.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - callback (optional) to call when the request is satisfied.
//
func (clipboard *Clipboard) ReadTextureAsync(ctx context.Context, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkClipboard       // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_clipboard_read_texture_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ReadTextureFinish finishes an asynchronous clipboard read.
//
// See gdk.Clipboard.ReadTextureAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - texture (optional): new GdkTexture or NULL on error.
//
func (clipboard *Clipboard) ReadTextureFinish(result gio.AsyncResulter) (Texturer, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GdkTexture   // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.gdk_clipboard_read_texture_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(result)

	var _texture Texturer // out
	var _goerr error      // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Texturer)
				return ok
			})
			rv, ok := casted.(Texturer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Texturer")
			}
			_texture = rv
		}
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _texture, _goerr
}

// ReadValueAsync: asynchronously request the clipboard contents converted to
// the given type.
//
// When the operation is finished callback will be called. You must then call
// gdk.Clipboard.ReadValueFinish() to get the resulting GValue.
//
// For local clipboard contents that are available in the given GType, the value
// will be copied directly. Otherwise, GDK will try to use
// content_deserialize_async to convert the clipboard's data.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - typ: GType to read.
//    - ioPriority: i/O priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (clipboard *Clipboard) ReadValueAsync(ctx context.Context, typ externglib.Type, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkClipboard       // out
	var _arg3 *C.GCancellable       // out
	var _arg1 C.GType               // out
	var _arg2 C.int                 // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GType(typ)
	_arg2 = C.int(ioPriority)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_clipboard_read_value_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// ReadValueFinish finishes an asynchronous clipboard read.
//
// See gdk.Clipboard.ReadValueAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
// The function returns the following values:
//
//    - value: GValue containing the result.
//
func (clipboard *Clipboard) ReadValueFinish(result gio.AsyncResulter) (*externglib.Value, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GValue       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	_cret = C.gdk_clipboard_read_value_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(result)

	var _value *externglib.Value // out
	var _goerr error             // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}

// SetContent sets a new content provider on clipboard.
//
// The clipboard will claim the GdkDisplay's resources and advertise these new
// contents to other applications.
//
// In the rare case of a failure, this function will return FALSE. The clipboard
// will then continue reporting its old contents and ignore provider.
//
// If the contents are read by either an external application or the clipboard's
// read functions, clipboard will select the best format to transfer the
// contents and then request that format from provider.
//
// The function takes the following parameters:
//
//    - provider (optional): new contents of clipboard or NULL to clear the
//      clipboard.
//
// The function returns the following values:
//
//    - ok: TRUE if setting the clipboard succeeded.
//
func (clipboard *Clipboard) SetContent(provider *ContentProvider) bool {
	var _arg0 *C.GdkClipboard       // out
	var _arg1 *C.GdkContentProvider // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	if provider != nil {
		_arg1 = (*C.GdkContentProvider)(unsafe.Pointer(externglib.InternObject(provider).Native()))
	}

	_cret = C.gdk_clipboard_set_content(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(provider)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetText puts the given text into the clipboard.
//
// The function takes the following parameters:
//
//    - text: text to put into the clipboard.
//
func (clipboard *Clipboard) SetText(text string) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_clipboard_set_text(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(text)
}

// SetTexture puts the given texture into the clipboard.
//
// The function takes the following parameters:
//
//    - texture: GdkTexture to put into the clipboard.
//
func (clipboard *Clipboard) SetTexture(texture Texturer) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GdkTexture   // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GdkTexture)(unsafe.Pointer(externglib.InternObject(texture).Native()))

	C.gdk_clipboard_set_texture(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(texture)
}

// Set sets the clipboard to contain the given value.
//
// The function takes the following parameters:
//
//    - value: GValue to set.
//
func (clipboard *Clipboard) Set(value *externglib.Value) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GValue       // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gdk_clipboard_set_value(_arg0, _arg1)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(value)
}

// StoreAsync: asynchronously instructs the clipboard to store its contents
// remotely.
//
// If the clipboard is not local, this function does nothing but report success.
//
// The callback must call gdk.Clipboard.StoreFinish().
//
// The purpose of this call is to preserve clipboard contents beyond the
// lifetime of an application, so this function is typically called on exit.
// Depending on the platform, the functionality may not be available unless a
// "clipboard manager" is running.
//
// This function is called automatically when a gtk.Application is shut down, so
// you likely don't need to call it.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional GCancellable object, NULL to ignore.
//    - ioPriority: i/O priority of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (clipboard *Clipboard) StoreAsync(ctx context.Context, ioPriority int, callback gio.AsyncReadyCallback) {
	var _arg0 *C.GdkClipboard       // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.int(ioPriority)
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.gdk_clipboard_store_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// StoreFinish finishes an asynchronous clipboard store.
//
// See gdk.Clipboard.StoreAsync().
//
// The function takes the following parameters:
//
//    - result: GAsyncResult.
//
func (clipboard *Clipboard) StoreFinish(result gio.AsyncResulter) error {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(externglib.InternObject(clipboard).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(externglib.InternObject(result).Native()))

	C.gdk_clipboard_store_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(clipboard)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
