// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// BuildAttributeListForCopy prepares the file attribute query string for
// copying to file.
//
// This function prepares an attribute query string to be passed to
// g_file_query_info() to get a list of attributes normally copied with the file
// (see g_file_copy_attributes() for the detailed description). This function is
// used by the implementation of g_file_copy_attributes() and is useful when one
// needs to query and set the attributes in two stages (e.g., for recursive move
// of a directory).
//
// The function takes the following parameters:
//
//    - ctx (optional): optional #GCancellable object, NULL to ignore.
//    - flags: set of CopyFlags.
//
// The function returns the following values:
//
//    - utf8: attribute query string for g_file_query_info(), or NULL if an error
//      occurs.
//
func (file *File) BuildAttributeListForCopy(ctx context.Context, flags FileCopyFlags) (string, error) {
	var _arg0 *C.GFile         // out
	var _arg2 *C.GCancellable  // out
	var _arg1 C.GFileCopyFlags // out
	var _cret *C.char          // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.GFileCopyFlags(flags)

	_cret = C.g_file_build_attribute_list_for_copy(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(file)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(flags)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}
