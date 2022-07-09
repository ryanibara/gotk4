// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// IOModulesScanAllInDirectory scans all the modules in the specified directory,
// ensuring that any extension point implemented by a module is registered.
//
// This may not actually load and initialize all the types in each module, some
// modules may be lazily loaded and initialized when an extension point it
// implements is used with e.g. g_io_extension_point_get_extensions() or
// g_io_extension_point_get_extension_by_name().
//
// If you need to guarantee that all types are loaded in all the modules, use
// g_io_modules_load_all_in_directory().
//
// The function takes the following parameters:
//
//    - dirname: pathname for a directory containing modules to scan.
//
func IOModulesScanAllInDirectory(dirname string) {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(dirname)))
	defer C.free(unsafe.Pointer(_args[0]))

	girepository.MustFind("Gio", "io_modules_scan_all_in_directory").Invoke(_args[:], nil)

	runtime.KeepAlive(dirname)
}
