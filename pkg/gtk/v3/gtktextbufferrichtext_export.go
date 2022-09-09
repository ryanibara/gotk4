// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

//export _gotk4_gtk3_TextBufferDeserializeFunc
func _gotk4_gtk3_TextBufferDeserializeFunc(arg1 *C.GtkTextBuffer, arg2 *C.GtkTextBuffer, arg3 *C.GtkTextIter, arg4 *C.guint8, arg5 C.gsize, arg6 C.gboolean, arg7 C.gpointer, _cerr **C.GError) (cret C.gboolean) {
	var fn TextBufferDeserializeFunc
	{
		v := gbox.Get(uintptr(arg7))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TextBufferDeserializeFunc)
	}

	var _registerBuffer *TextBuffer // out
	var _contentBuffer *TextBuffer  // out
	var _iter *TextIter             // out
	var _data []byte                // out
	var _createTags bool            // out

	_registerBuffer = wrapTextBuffer(coreglib.Take(unsafe.Pointer(arg1)))
	_contentBuffer = wrapTextBuffer(coreglib.Take(unsafe.Pointer(arg2)))
	_iter = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	_data = make([]byte, arg5)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(arg4)), arg5))
	if arg6 != 0 {
		_createTags = true
	}

	_goerr := fn(_registerBuffer, _contentBuffer, _iter, _data, _createTags)

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gtk3_TextBufferSerializeFunc
func _gotk4_gtk3_TextBufferSerializeFunc(arg1 *C.GtkTextBuffer, arg2 *C.GtkTextBuffer, arg3 *C.GtkTextIter, arg4 *C.GtkTextIter, arg5 *C.gsize, arg6 C.gpointer) (cret *C.guint8) {
	var fn TextBufferSerializeFunc
	{
		v := gbox.Get(uintptr(arg6))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TextBufferSerializeFunc)
	}

	var _registerBuffer *TextBuffer // out
	var _contentBuffer *TextBuffer  // out
	var _start *TextIter            // out
	var _end *TextIter              // out

	_registerBuffer = wrapTextBuffer(coreglib.Take(unsafe.Pointer(arg1)))
	_contentBuffer = wrapTextBuffer(coreglib.Take(unsafe.Pointer(arg2)))
	_start = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	_end = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	length, guint8 := fn(_registerBuffer, _contentBuffer, _start, _end)

	var _ uint
	var _ *byte

	*arg5 = C.gsize(length)
	if guint8 != nil {
		cret = (*C.guint8)(unsafe.Pointer(guint8))
	}

	return cret
}
