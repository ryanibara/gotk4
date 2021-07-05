// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
//
// void gotk4_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_enumerator_get_type()), F: marshalFileEnumerator},
	})
}

// FileEnumerator allows you to operate on a set of #GFiles, returning a Info
// structure for each file enumerated (e.g. g_file_enumerate_children() will
// return a Enumerator for each of the children within a directory).
//
// To get the next file's information from a Enumerator, use
// g_file_enumerator_next_file() or its asynchronous version,
// g_file_enumerator_next_files_async(). Note that the asynchronous version will
// return a list of Infos, whereas the synchronous will only return the next
// file in the enumerator.
//
// The ordering of returned files is unspecified for non-Unix platforms; for
// more information, see g_dir_read_name(). On Unix, when operating on local
// files, returned files will be sorted by inode number. Effectively you can
// assume that the ordering of returned files will be stable between successive
// calls (and applications) assuming the directory is unchanged.
//
// If your application needs a specific ordering, such as by name or
// modification time, you will have to implement that in your application code.
//
// To close a Enumerator, use g_file_enumerator_close(), or its asynchronous
// version, g_file_enumerator_close_async(). Once a Enumerator is closed, no
// further actions may be performed on it, and it should be freed with
// g_object_unref().
type FileEnumerator interface {
	gextras.Objector

	// CloseFileEnumerator releases all resources used by this enumerator,
	// making the enumerator return G_IO_ERROR_CLOSED on all calls.
	//
	// This will be automatically called when the last reference is dropped, but
	// you might want to call this function to make sure resources are released
	// as early as possible.
	CloseFileEnumerator(cancellable Cancellable) error
	// CloseAsyncFileEnumerator: asynchronously closes the file enumerator.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned in
	// g_file_enumerator_close_finish().
	CloseAsyncFileEnumerator(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// CloseFinishFileEnumerator finishes closing a file enumerator, started
	// from g_file_enumerator_close_async().
	//
	// If the file enumerator was already closed when
	// g_file_enumerator_close_async() was called, then this function will
	// report G_IO_ERROR_CLOSED in @error, and return false. If the file
	// enumerator had pending operation when the close operation was started,
	// then this function will report G_IO_ERROR_PENDING, and return false. If
	// @cancellable was not nil, then the operation may have been cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and false will
	// be returned.
	CloseFinishFileEnumerator(result AsyncResult) error
	// Child: return a new #GFile which refers to the file named by @info in the
	// source directory of @enumerator. This function is primarily intended to
	// be used inside loops with g_file_enumerator_next_file().
	//
	// This is a convenience method that's equivalent to:
	//
	//    gchar *name = g_file_info_get_name (info);
	//    GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
	//                                     name);
	Child(info FileInfo) File
	// Container: get the #GFile container which is being enumerated.
	Container() File
	// HasPendingFileEnumerator checks if the file enumerator has pending
	// operations.
	HasPendingFileEnumerator() bool
	// IsClosedFileEnumerator checks if the file enumerator has been closed.
	IsClosedFileEnumerator() bool
	// IterateFileEnumerator: this is a version of g_file_enumerator_next_file()
	// that's easier to use correctly from C programs. With
	// g_file_enumerator_next_file(), the gboolean return value signifies "end
	// of iteration or error", which requires allocation of a temporary #GError.
	//
	// In contrast, with this function, a false return from
	// g_file_enumerator_iterate() *always* means "error". End of iteration is
	// signaled by @out_info or @out_child being nil.
	//
	// Another crucial difference is that the references for @out_info and
	// @out_child are owned by @direnum (they are cached as hidden properties).
	// You must not unref them in your own code. This makes memory management
	// significantly easier for C code in combination with loops.
	//
	// Finally, this function optionally allows retrieving a #GFile as well.
	//
	// You must specify at least one of @out_info or @out_child.
	//
	// The code pattern for correctly using g_file_enumerator_iterate() from C
	// is:
	//
	//    direnum = g_file_enumerate_children (file, ...);
	//    while (TRUE)
	//      {
	//        GFileInfo *info;
	//        if (!g_file_enumerator_iterate (direnum, &info, NULL, cancellable, error))
	//          goto out;
	//        if (!info)
	//          break;
	//        ... do stuff with "info"; do not unref it! ...
	//      }
	//
	//    out:
	//      g_object_unref (direnum); // Note: frees the last @info
	IterateFileEnumerator(cancellable Cancellable) (FileInfo, File, error)
	// NextFileFileEnumerator returns information for the next file in the
	// enumerated object. Will block until the information is available. The
	// Info returned from this function will contain attributes that match the
	// attribute string that was passed when the Enumerator was created.
	//
	// See the documentation of Enumerator for information about the order of
	// returned files.
	//
	// On error, returns nil and sets @error to the error. If the enumerator is
	// at the end, nil will be returned and @error will be unset.
	NextFileFileEnumerator(cancellable Cancellable) (FileInfo, error)
	// NextFilesAsyncFileEnumerator: request information for a number of files
	// from the enumerator asynchronously. When all i/o for the operation is
	// finished the @callback will be called with the requested information.
	//
	// See the documentation of Enumerator for information about the order of
	// returned files.
	//
	// The callback can be called with less than @num_files files in case of
	// error or at the end of the enumerator. In case of a partial error the
	// callback will be called with any succeeding items and no error, and on
	// the next request the error will be reported. If a request is cancelled
	// the callback will be called with G_IO_ERROR_CANCELLED.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	NextFilesAsyncFileEnumerator(numFiles int, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
	// SetPendingFileEnumerator sets the file enumerator as having pending
	// operations.
	SetPendingFileEnumerator(pending bool)
}

// fileEnumerator implements the FileEnumerator class.
type fileEnumerator struct {
	gextras.Objector
}

// WrapFileEnumerator wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileEnumerator(obj *externglib.Object) FileEnumerator {
	return fileEnumerator{
		Objector: obj,
	}
}

func marshalFileEnumerator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileEnumerator(obj), nil
}

func (e fileEnumerator) CloseFileEnumerator(cancellable Cancellable) error {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 *C.GCancellable    // out
	var _cerr **C.GError         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_file_enumerator_close(_arg0, _arg1, &_cerr)

	var _goerr error // out

	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _goerr
}

func (e fileEnumerator) CloseAsyncFileEnumerator(ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileEnumerator    // out
	var _arg1 C.int                 // out
	var _arg2 *C.GCancellable       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = C.int(ioPriority)
	_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg3 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg4 = C.gpointer(box.Assign(callback))

	C.g_file_enumerator_close_async(_arg0, _arg1, _arg2, _arg3, _arg4)
}

func (e fileEnumerator) CloseFinishFileEnumerator(result AsyncResult) error {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 *C.GAsyncResult    // out
	var _cerr **C.GError         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_file_enumerator_close_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _goerr
}

func (e fileEnumerator) Child(info FileInfo) File {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 *C.GFileInfo       // out
	var _cret *C.GFile           // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GFileInfo)(unsafe.Pointer(info.Native()))

	_cret = C.g_file_enumerator_get_child(_arg0, _arg1)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

func (e fileEnumerator) Container() File {
	var _arg0 *C.GFileEnumerator // out
	var _cret *C.GFile           // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	_cret = C.g_file_enumerator_get_container(_arg0)

	var _file File // out

	_file = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(File)

	return _file
}

func (e fileEnumerator) HasPendingFileEnumerator() bool {
	var _arg0 *C.GFileEnumerator // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	_cret = C.g_file_enumerator_has_pending(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e fileEnumerator) IsClosedFileEnumerator() bool {
	var _arg0 *C.GFileEnumerator // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	_cret = C.g_file_enumerator_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d fileEnumerator) IterateFileEnumerator(cancellable Cancellable) (FileInfo, File, error) {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 **C.GFileInfo      // in
	var _arg2 **C.GFile          // in
	var _arg3 *C.GCancellable    // out
	var _cerr **C.GError         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(d.Native()))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	C.g_file_enumerator_iterate(_arg0, &_arg1, &_arg2, _arg3, &_cerr)

	var _outInfo FileInfo // out
	var _outChild File    // out
	var _goerr error      // out

	{
		var refTmpIn *C.GFileInfo
		var refTmpOut fileInfo

		refTmpIn = *_arg1

		refTmpOut = gextras.CastObject(externglib.Take(unsafe.Pointer(refTmpIn))).(fileInfo)

		_outInfo = refTmpOut
	}
	{
		var refTmpIn *C.GFile
		var refTmpOut file

		refTmpIn = *_arg2

		refTmpOut = gextras.CastObject(externglib.Take(unsafe.Pointer(refTmpIn))).(file)

		_outChild = refTmpOut
	}
	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _outInfo, _outChild, _goerr
}

func (e fileEnumerator) NextFileFileEnumerator(cancellable Cancellable) (FileInfo, error) {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 *C.GCancellable    // out
	var _cret *C.GFileInfo       // in
	var _cerr **C.GError         // in

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_file_enumerator_next_file(_arg0, _arg1, &_cerr)

	var _fileInfo FileInfo // out
	var _goerr error       // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(FileInfo)
	{
		var refTmpIn *C.GError
		var refTmpOut error

		refTmpIn = *_cerr

		refTmpOut = gerror.Take(unsafe.Pointer(refTmpIn))

		_goerr = refTmpOut
	}

	return _fileInfo, _goerr
}

func (e fileEnumerator) NextFilesAsyncFileEnumerator(numFiles int, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileEnumerator    // out
	var _arg1 C.int                 // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	_arg1 = C.int(numFiles)
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_file_enumerator_next_files_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}

func (e fileEnumerator) SetPendingFileEnumerator(pending bool) {
	var _arg0 *C.GFileEnumerator // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	if pending {
		_arg1 = C.TRUE
	}

	C.g_file_enumerator_set_pending(_arg0, _arg1)
}
