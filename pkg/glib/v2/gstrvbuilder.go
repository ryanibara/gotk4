// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// StrvBuilder is a method of easily building dynamically sized NULL-terminated
// string arrays.
//
// The following example shows how to build a two element array:
//
//    g_autoptr(GStrvBuilder) builder = g_strv_builder_new ();
//    g_strv_builder_add (builder, "hello");
//    g_strv_builder_add (builder, "world");
//    g_auto(GStrv) array = g_strv_builder_end (builder);
type StrvBuilder struct {
	native C.GStrvBuilder
}

// WrapStrvBuilder wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStrvBuilder(ptr unsafe.Pointer) *StrvBuilder {
	if ptr == nil {
		return nil
	}

	return (*StrvBuilder)(ptr)
}

func marshalStrvBuilder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStrvBuilder(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StrvBuilder) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Add: add a string to the end of the array.
//
// Since 2.68
func (b *StrvBuilder) Add(value string) {
	var arg0 *C.GStrvBuilder
	var arg1 *C.char

	arg0 = (*C.GStrvBuilder)(b.Native())
	arg1 = (*C.gchar)(C.CString(value))
	defer C.free(unsafe.Pointer(arg1))

	C.g_strv_builder_add(arg0, arg1)
}

// End ends the builder process and returns the constructed NULL-terminated
// string array. The returned value should be freed with g_strfreev() when no
// longer needed.
func (b *StrvBuilder) End() []string {
	var arg0 *C.GStrvBuilder

	arg0 = (*C.GStrvBuilder)(b.Native())

	ret := C.g_strv_builder_end(arg0)

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

// Ref: atomically increments the reference count of @builder by one. This
// function is thread-safe and may be called from any thread.
func (b *StrvBuilder) Ref() *StrvBuilder {
	var arg0 *C.GStrvBuilder

	arg0 = (*C.GStrvBuilder)(b.Native())

	ret := C.g_strv_builder_ref(arg0)

	var ret0 *StrvBuilder

	{
		ret0 = WrapStrvBuilder(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *StrvBuilder) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Unref decreases the reference count on @builder.
//
// In the event that there are no more references, releases all memory
// associated with the Builder.
func (b *StrvBuilder) Unref() {
	var arg0 *C.GStrvBuilder

	arg0 = (*C.GStrvBuilder)(b.Native())

	C.g_strv_builder_unref(arg0)
}
