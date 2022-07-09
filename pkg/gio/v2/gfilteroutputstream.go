// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeFilterOutputStream returns the GType for the type FilterOutputStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFilterOutputStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "FilterOutputStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFilterOutputStream)
	return gtype
}

// FilterOutputStreamOverrider contains methods that are overridable.
type FilterOutputStreamOverrider interface {
}

// FilterOutputStream: base class for output stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterOutputStream struct {
	_ [0]func() // equal guard
	OutputStream
}

var (
	_ OutputStreamer = (*FilterOutputStream)(nil)
)

// FilterOutputStreamer describes types inherited from class FilterOutputStream.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FilterOutputStreamer interface {
	coreglib.Objector
	baseFilterOutputStream() *FilterOutputStream
}

var _ FilterOutputStreamer = (*FilterOutputStream)(nil)

func classInitFilterOutputStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFilterOutputStream(obj *coreglib.Object) *FilterOutputStream {
	return &FilterOutputStream{
		OutputStream: OutputStream{
			Object: obj,
		},
	}
}

func marshalFilterOutputStream(p uintptr) (interface{}, error) {
	return wrapFilterOutputStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (stream *FilterOutputStream) baseFilterOutputStream() *FilterOutputStream {
	return stream
}

// BaseFilterOutputStream returns the underlying base object.
func BaseFilterOutputStream(obj FilterOutputStreamer) *FilterOutputStream {
	return obj.baseFilterOutputStream()
}

// BaseStream gets the base stream for the filter stream.
//
// The function returns the following values:
//
//    - outputStream: Stream.
//
func (stream *FilterOutputStream) BaseStream() OutputStreamer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "FilterOutputStream").InvokeMethod("get_base_stream", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _outputStream OutputStreamer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.OutputStreamer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(OutputStreamer)
			return ok
		})
		rv, ok := casted.(OutputStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.OutputStreamer")
		}
		_outputStream = rv
	}

	return _outputStream
}

// CloseBaseStream returns whether the base stream will be closed when stream is
// closed.
//
// The function returns the following values:
//
//    - ok: TRUE if the base stream will be closed.
//
func (stream *FilterOutputStream) CloseBaseStream() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_gret := girepository.MustFind("Gio", "FilterOutputStream").InvokeMethod("get_close_base_stream", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetCloseBaseStream sets whether the base stream will be closed when stream is
// closed.
//
// The function takes the following parameters:
//
//    - closeBase: TRUE to close the base stream.
//
func (stream *FilterOutputStream) SetCloseBaseStream(closeBase bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	if closeBase {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gio", "FilterOutputStream").InvokeMethod("set_close_base_stream", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(closeBase)
}
