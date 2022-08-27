// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// gboolean _gotk4_gio2_File_virtual_delete_file_finish(void* fnptr, GFile* arg0, GAsyncResult* arg1, GError** arg2) {
//   return ((gboolean (*)(GFile*, GAsyncResult*, GError**))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_File_virtual_delete_file_async(void* fnptr, GFile* arg0, int arg1, GCancellable* arg2, GAsyncReadyCallback arg3, gpointer arg4) {
//   ((void (*)(GFile*, int, GCancellable*, GAsyncReadyCallback, gpointer))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
import "C"

// DeleteAsync: asynchronously delete a file. If the file is a directory, it
// will only be deleted if it is empty. This has the same semantics as
// g_unlink().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (file *File) DeleteAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback) {
	var _arg0 *C.GFile              // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
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

	C.g_file_delete_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(file)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// DeleteFinish finishes deleting a file started with g_file_delete_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (file *File) DeleteFinish(result AsyncResulter) error {
	var _arg0 *C.GFile        // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C.g_file_delete_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(file)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// deleteFileAsync: asynchronously delete a file. If the file is a directory, it
// will only be deleted if it is empty. This has the same semantics as
// g_unlink().
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - ioPriority: [I/O priority][io-priority] of the request.
//    - callback (optional) to call when the request is satisfied.
//
func (file *File) deleteFileAsync(ctx context.Context, ioPriority int, callback AsyncReadyCallback) {
	gclass := (*C.GFileIface)(coreglib.PeekParentClass(file))
	fnarg := gclass.delete_file_async

	var _arg0 *C.GFile              // out
	var _arg2 *C.GCancellable       // out
	var _arg1 C.int                 // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
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

	C._gotk4_gio2_File_virtual_delete_file_async(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(file)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(ioPriority)
	runtime.KeepAlive(callback)
}

// deleteFileFinish finishes deleting a file started with g_file_delete_async().
//
// The function takes the following parameters:
//
//    - result: Result.
//
func (file *File) deleteFileFinish(result AsyncResulter) error {
	gclass := (*C.GFileIface)(coreglib.PeekParentClass(file))
	fnarg := gclass.delete_file_finish

	var _arg0 *C.GFile        // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(coreglib.InternObject(result).Native()))

	C._gotk4_gio2_File_virtual_delete_file_finish(unsafe.Pointer(fnarg), _arg0, _arg1, &_cerr)
	runtime.KeepAlive(file)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
