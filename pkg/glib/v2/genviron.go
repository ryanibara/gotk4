// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"reflect"
	"unsafe"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib.h>
import "C"

// EnvironGetenv returns the value of the environment variable @variable in the
// provided list @envp.
func EnvironGetenv(envp []string, variable string) string {
	var arg1 **C.gchar
	var arg2 *C.gchar

	{
		var dst []*C.gchar
		ptr := C.malloc(unsafe.Sizeof((*struct{})(nil)) * (len(envp) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(envp)
		sliceHeader.Cap = len(envp)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(envp); i++ {
			src := envp[i]
			dst[i] = (*C.gchar)(C.CString(src))
			defer C.free(unsafe.Pointer(dst[i]))
		}

		arg1 = (**C.gchar)(unsafe.Pointer(ptr))
	}
	arg2 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_environ_getenv(arg1, arg2)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// EnvironSetenv sets the environment variable @variable in the provided list
// @envp to @value.
func EnvironSetenv(envp []string, variable string, value string, overwrite bool) []string {
	var arg1 **C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 C.gboolean

	{
		var dst []*C.gchar
		ptr := C.malloc(unsafe.Sizeof((*struct{})(nil)) * (len(envp) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(envp)
		sliceHeader.Cap = len(envp)

		for i := 0; i < len(envp); i++ {
			src := envp[i]
			dst[i] = (*C.gchar)(C.CString(src))
		}

		arg1 = (**C.gchar)(unsafe.Pointer(ptr))
	}
	arg2 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(arg3))
	if overwrite {
		arg4 = C.TRUE
	}

	ret := C.g_environ_setenv(arg1, arg2, arg3, arg4)

	var ret0 []string

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	return ret0
}

// EnvironUnsetenv removes the environment variable @variable from the provided
// environment @envp.
func EnvironUnsetenv(envp []string, variable string) []string {
	var arg1 **C.gchar
	var arg2 *C.gchar

	{
		var dst []*C.gchar
		ptr := C.malloc(unsafe.Sizeof((*struct{})(nil)) * (len(envp) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(envp)
		sliceHeader.Cap = len(envp)

		for i := 0; i < len(envp); i++ {
			src := envp[i]
			dst[i] = (*C.gchar)(C.CString(src))
		}

		arg1 = (**C.gchar)(unsafe.Pointer(ptr))
	}
	arg2 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_environ_unsetenv(arg1, arg2)

	var ret0 []string

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	return ret0
}

// GetEnviron gets the list of environment variables for the current process.
//
// The list is nil terminated and each item in the list is of the form
// 'NAME=VALUE'.
//
// This is equivalent to direct access to the 'environ' global variable, except
// portable.
//
// The return value is freshly allocated and it should be freed with
// g_strfreev() when it is no longer needed.
func GetEnviron() []string {
	ret := C.g_get_environ()

	var ret0 []string

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	return ret0
}

// Getenv returns the value of an environment variable.
//
// On UNIX, the name and value are byte strings which might or might not be in
// some consistent character set and encoding. On Windows, they are in UTF-8. On
// Windows, in case the environment variable's value contains references to
// other environment variables, they are expanded.
func Getenv(variable string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_getenv(arg1)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Listenv gets the names of all variables set in the environment.
//
// Programs that want to be portable to Windows should typically use this
// function and g_getenv() instead of using the environ array from the C library
// directly. On Windows, the strings in the environ array are in system codepage
// encoding, while in most of the typical use cases for environment variables in
// GLib-using programs you want the UTF-8 encoding that this function and
// g_getenv() provide.
func Listenv() []string {
	ret := C.g_listenv()

	var ret0 []string

	{
		var length uint
		for p := unsafe.Pointer(ret); *p != 0; p = unsafe.Pointer(uintptr(p) + 1) {
			length++
		}

		ret0 = make([]string, length)
		for i := 0; i < length; i++ {
			src := (*C.gchar)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + i))
			ret0[i] = C.GoString(src)
			C.free(unsafe.Pointer(src))
		}
	}

	return ret0
}

// Setenv sets an environment variable. On UNIX, both the variable's name and
// value can be arbitrary byte strings, except that the variable's name cannot
// contain '='. On Windows, they should be in UTF-8.
//
// Note that on some systems, when variables are overwritten, the memory used
// for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling in UNIX
// is not thread-safe, and your program may crash if one thread calls g_setenv()
// while another thread is calling getenv(). (And note that many functions, such
// as gettext(), call getenv() internally.) This function is only safe to use at
// the very start of your program, before creating any other threads (or
// creating objects that create worker threads of their own).
//
// If you need to set up the environment for a child process, you can use
// g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that array
// directly to execvpe(), g_spawn_async(), or the like.
func Setenv(variable string, value string, overwrite bool) bool {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 C.gboolean

	arg1 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(arg2))
	if overwrite {
		arg3 = C.TRUE
	}

	ret := C.g_setenv(arg1, arg2, arg3)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Unsetenv removes an environment variable from the environment.
//
// Note that on some systems, when variables are overwritten, the memory used
// for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling in UNIX
// is not thread-safe, and your program may crash if one thread calls
// g_unsetenv() while another thread is calling getenv(). (And note that many
// functions, such as gettext(), call getenv() internally.) This function is
// only safe to use at the very start of your program, before creating any other
// threads (or creating objects that create worker threads of their own).
//
// If you need to set up the environment for a child process, you can use
// g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that array
// directly to execvpe(), g_spawn_async(), or the like.
func Unsetenv(variable string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg1))

	C.g_unsetenv(arg1)
}
