// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern char* _gotk4_gio2_FileIOStreamClass_get_etag(void*);
// extern gboolean _gotk4_gio2_FileIOStreamClass_can_seek(void*);
// extern gboolean _gotk4_gio2_FileIOStreamClass_can_truncate(void*);
// extern void _gotk4_gio2_AsyncReadyCallback(void*, void*, gpointer);
// extern void* _gotk4_gio2_FileIOStreamClass_query_info(void*, char*, void*, GError**);
// extern void* _gotk4_gio2_FileIOStreamClass_query_info_finish(void*, void*, GError**);
import "C"

// GTypeFileIOStream returns the GType for the type FileIOStream.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFileIOStream() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "FileIOStream").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalFileIOStream)
	return gtype
}

// FileIOStreamOverrider contains methods that are overridable.
type FileIOStreamOverrider interface {
	// The function returns the following values:
	//
	CanSeek() bool
	// The function returns the following values:
	//
	CanTruncate() bool
	// ETag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	//
	// The function returns the following values:
	//
	//    - utf8 (optional): entity tag for the stream.
	//
	ETag() string
	// QueryInfo queries a file io stream for the given attributes. This
	// function blocks while querying the stream. For the asynchronous version
	// of this function, see g_file_io_stream_query_info_async(). While the
	// stream is blocked, the stream will set the pending flag internally, and
	// any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with error being set to G_IO_ERROR_NOT_SUPPORTED). I
	// all cases of failure, NULL will be returned.
	//
	// If cancellable is not NULL, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and NULL will
	// be returned.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): optional #GCancellable object, NULL to ignore.
	//    - attributes: file attribute query string.
	//
	// The function returns the following values:
	//
	//    - fileInfo for the stream, or NULL on error.
	//
	QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_io_stream_query_info_async().
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - fileInfo for the finished query.
	//
	QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
}

// FileIOStream provides io streams that both read and write to the same file
// handle.
//
// GFileIOStream implements #GSeekable, which allows the io stream to jump to
// arbitrary positions in the file and to truncate the file, provided the
// filesystem of the file supports these operations.
//
// To find the position of a file io stream, use g_seekable_tell().
//
// To find out if a file io stream supports seeking, use g_seekable_can_seek().
// To position a file io stream, use g_seekable_seek(). To find out if a file io
// stream supports truncating, use g_seekable_can_truncate(). To truncate a file
// io stream, use g_seekable_truncate().
//
// The default implementation of all the IOStream operations and the
// implementation of #GSeekable just call into the same operations on the output
// stream.
type FileIOStream struct {
	_ [0]func() // equal guard
	IOStream

	*coreglib.Object
	Seekable
}

var (
	_ IOStreamer        = (*FileIOStream)(nil)
	_ coreglib.Objector = (*FileIOStream)(nil)
)

func classInitFileIOStreamer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := girepository.MustFind("Gio", "FileIOStreamClass")

	if _, ok := goval.(interface{ CanSeek() bool }); ok {
		o := pclass.StructFieldOffset("can_seek")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FileIOStreamClass_can_seek)
	}

	if _, ok := goval.(interface{ CanTruncate() bool }); ok {
		o := pclass.StructFieldOffset("can_truncate")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FileIOStreamClass_can_truncate)
	}

	if _, ok := goval.(interface{ ETag() string }); ok {
		o := pclass.StructFieldOffset("get_etag")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FileIOStreamClass_get_etag)
	}

	if _, ok := goval.(interface {
		QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	}); ok {
		o := pclass.StructFieldOffset("query_info")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FileIOStreamClass_query_info)
	}

	if _, ok := goval.(interface {
		QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	}); ok {
		o := pclass.StructFieldOffset("query_info_finish")
		*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gclassPtr), o)) = unsafe.Pointer(C._gotk4_gio2_FileIOStreamClass_query_info_finish)
	}
}

//export _gotk4_gio2_FileIOStreamClass_can_seek
func _gotk4_gio2_FileIOStreamClass_can_seek(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CanSeek() bool })

	ok := iface.CanSeek()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_FileIOStreamClass_can_truncate
func _gotk4_gio2_FileIOStreamClass_can_truncate(arg0 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ CanTruncate() bool })

	ok := iface.CanTruncate()

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gio2_FileIOStreamClass_get_etag
func _gotk4_gio2_FileIOStreamClass_get_etag(arg0 *C.void) (cret *C.char) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ETag() string })

	utf8 := iface.ETag()

	if utf8 != "" {
		cret = (*C.char)(unsafe.Pointer(C.CString(utf8)))
	}

	return cret
}

//export _gotk4_gio2_FileIOStreamClass_query_info
func _gotk4_gio2_FileIOStreamClass_query_info(arg0 *C.void, arg1 *C.char, arg2 *C.void, _cerr **C.GError) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		QueryInfo(ctx context.Context, attributes string) (*FileInfo, error)
	})

	var _cancellable context.Context // out
	var _attributes string           // out

	if arg2 != nil {
		_cancellable = gcancel.NewCancellableContext(unsafe.Pointer(arg2))
	}
	_attributes = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	fileInfo, _goerr := iface.QueryInfo(_cancellable, _attributes)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gio2_FileIOStreamClass_query_info_finish
func _gotk4_gio2_FileIOStreamClass_query_info_finish(arg0 *C.void, arg1 *C.void, _cerr **C.GError) (cret *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		QueryInfoFinish(result AsyncResulter) (*FileInfo, error)
	})

	var _result AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	fileInfo, _goerr := iface.QueryInfoFinish(_result)

	cret = (*C.void)(unsafe.Pointer(coreglib.InternObject(fileInfo).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(fileInfo).Native()))
	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapFileIOStream(obj *coreglib.Object) *FileIOStream {
	return &FileIOStream{
		IOStream: IOStream{
			Object: obj,
		},
		Object: obj,
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalFileIOStream(p uintptr) (interface{}, error) {
	return wrapFileIOStream(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ETag gets the entity tag for the file when it has been written. This must be
// called after the stream has been written and closed, as the etag can change
// while writing.
//
// The function returns the following values:
//
//    - utf8 (optional): entity tag for the stream.
//
func (stream *FileIOStream) ETag() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))

	_info := girepository.MustFind("Gio", "FileIOStream")
	_gret := _info.InvokeClassMethod("get_etag", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)

	var _utf8 string // out

	if *(**C.char)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
		defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret))))
	}

	return _utf8
}

// QueryInfo queries a file io stream for the given attributes. This function
// blocks while querying the stream. For the asynchronous version of this
// function, see g_file_io_stream_query_info_async(). While the stream is
// blocked, the stream will set the pending flag internally, and any other
// operations on the stream will fail with G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with error being set to
// G_IO_ERROR_CLOSED), the stream has pending operations (with error being set
// to G_IO_ERROR_PENDING), or if querying info is not supported for the stream's
// interface (with error being set to G_IO_ERROR_NOT_SUPPORTED). I all cases of
// failure, NULL will be returned.
//
// If cancellable is not NULL, then the operation can be cancelled by triggering
// the cancellable object from another thread. If the operation was cancelled,
// the error G_IO_ERROR_CANCELLED will be set, and NULL will be returned.
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - attributes: file attribute query string.
//
// The function returns the following values:
//
//    - fileInfo for the stream, or NULL on error.
//
func (stream *FileIOStream) QueryInfo(ctx context.Context, attributes string) (*FileInfo, error) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[2] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "FileIOStream")
	_gret := _info.InvokeClassMethod("query_info", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _fileInfo, _goerr
}

// QueryInfoAsync: asynchronously queries the stream for a Info. When completed,
// callback will be called with a Result which can be used to finish the
// operation with g_file_io_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_io_stream_query_info().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - attributes: file attribute query string.
//    - ioPriority: [I/O priority][gio-GIOScheduler] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (stream *FileIOStream) QueryInfoAsync(ctx context.Context, attributes string, ioPriority int32, callback AsyncReadyCallback) {
	var _args [6]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_args[3] = (*C.void)(unsafe.Pointer(cancellable.Native()))
	}
	*(**C.char)(unsafe.Pointer(&_args[1])) = (*C.char)(unsafe.Pointer(C.CString(attributes)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[1]))))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(ioPriority)
	if callback != nil {
		*(*C.gpointer)(unsafe.Pointer(&_args[4])) = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_args[5] = C.gpointer(gbox.AssignOnce(callback))
	}

	_info := girepository.MustFind("Gio", "FileIOStream")
	_info.InvokeClassMethod("query_info_async", _args[:], nil)

	runtime.KeepAlive(stream)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(attributes)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// QueryInfoFinish finalizes the asynchronous query started by
// g_file_io_stream_query_info_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - fileInfo for the finished query.
//
func (stream *FileIOStream) QueryInfoFinish(result AsyncResulter) (*FileInfo, error) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(stream).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	_info := girepository.MustFind("Gio", "FileIOStream")
	_gret := _info.InvokeClassMethod("query_info_finish", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stream)
	runtime.KeepAlive(result)

	var _fileInfo *FileInfo // out
	var _goerr error        // out

	_fileInfo = wrapFileInfo(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _fileInfo, _goerr
}
