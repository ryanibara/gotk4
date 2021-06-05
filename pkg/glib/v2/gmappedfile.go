// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_mapped_file_get_type()), F: marshalMappedFile},
	})
}

// MappedFile: the File represents a file mapping created with
// g_mapped_file_new(). It has only private members and should not be accessed
// directly.
type MappedFile struct {
	native C.GMappedFile
}

// WrapMappedFile wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapMappedFile(ptr unsafe.Pointer) *MappedFile {
	if ptr == nil {
		return nil
	}

	return (*MappedFile)(ptr)
}

func marshalMappedFile(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapMappedFile(unsafe.Pointer(b)), nil
}

// NewMappedFile constructs a struct MappedFile.
func NewMappedFile(filename string, writable bool) (mappedFile *MappedFile, err error) {
	var arg1 *C.gchar
	var arg2 C.gboolean

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	if writable {
		arg2 = C.gboolean(1)
	}

	var errout *C.GError
	var goerr error
	var cret *C.GMappedFile
	var ret2 *MappedFile

	cret = C.g_mapped_file_new(filename, writable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = WrapMappedFile(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret2, func(v *MappedFile) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goerr, ret2
}

// NewMappedFileFromFd constructs a struct MappedFile.
func NewMappedFileFromFd(fd int, writable bool) (mappedFile *MappedFile, err error) {
	var arg1 C.gint
	var arg2 C.gboolean

	arg1 = C.gint(fd)
	if writable {
		arg2 = C.gboolean(1)
	}

	var errout *C.GError
	var goerr error
	var cret *C.GMappedFile
	var ret2 *MappedFile

	cret = C.g_mapped_file_new_from_fd(fd, writable, &errout)

	goerr = gerror.Take(unsafe.Pointer(errout))
	ret2 = WrapMappedFile(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret2, func(v *MappedFile) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goerr, ret2
}

// Native returns the underlying C source pointer.
func (m *MappedFile) Native() unsafe.Pointer {
	return unsafe.Pointer(&m.native)
}

// Free: this call existed before File had refcounting and is currently exactly
// the same as g_mapped_file_unref().
func (f *MappedFile) Free() {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	C.g_mapped_file_free(arg0)
}

// Bytes creates a new #GBytes which references the data mapped from @file. The
// mapped contents of the file must not be modified after creating this bytes
// object, because a #GBytes should be immutable.
func (f *MappedFile) Bytes() *Bytes {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var cret *C.GBytes
	var ret1 *Bytes

	cret = C.g_mapped_file_get_bytes(arg0)

	ret1 = WrapBytes(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *Bytes) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Contents returns the contents of a File.
//
// Note that the contents may not be zero-terminated, even if the File is backed
// by a text file.
//
// If the file is empty then nil is returned.
func (f *MappedFile) Contents() string {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.g_mapped_file_get_contents(arg0)

	ret1 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return ret1
}

// Length returns the length of the contents of a File.
func (f *MappedFile) Length() uint {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var cret C.gsize
	var ret1 uint

	cret = C.g_mapped_file_get_length(arg0)

	ret1 = C.gsize(cret)

	return ret1
}

// Ref increments the reference count of @file by one. It is safe to call this
// function from any thread.
func (f *MappedFile) Ref() *MappedFile {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	var cret *C.GMappedFile
	var ret1 *MappedFile

	cret = C.g_mapped_file_ref(arg0)

	ret1 = WrapMappedFile(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret1, func(v *MappedFile) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1
}

// Unref decrements the reference count of @file by one. If the reference count
// drops to 0, unmaps the buffer of @file and frees it.
//
// It is safe to call this function from any thread.
//
// Since 2.22
func (f *MappedFile) Unref() {
	var arg0 *C.GMappedFile

	arg0 = (*C.GMappedFile)(unsafe.Pointer(f.Native()))

	C.g_mapped_file_unref(arg0)
}
