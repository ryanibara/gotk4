// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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
	var _arg1 *C.char // out

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(dirname)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_io_modules_scan_all_in_directory(_arg1)
	runtime.KeepAlive(dirname)
}
