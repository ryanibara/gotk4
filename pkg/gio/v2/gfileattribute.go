// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_attribute_info_list_get_type()), F: marshalFileAttributeInfoList},
	})
}

// FileAttributeInfo: information about a specific attribute.
type FileAttributeInfo struct {
	native C.GFileAttributeInfo
}

// WrapFileAttributeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeInfo(ptr unsafe.Pointer) *FileAttributeInfo {
	if ptr == nil {
		return nil
	}

	return (*FileAttributeInfo)(ptr)
}

func marshalFileAttributeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFileAttributeInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FileAttributeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// Name gets the field inside the struct.
func (f *FileAttributeInfo) Name() string {
	var ret string
	ret = C.GoString(f.native.name)
	return ret
}

// Type gets the field inside the struct.
func (f *FileAttributeInfo) Type() FileAttributeType {
	var ret FileAttributeType
	ret = FileAttributeType(f.native._type)
	return ret
}

// Flags gets the field inside the struct.
func (f *FileAttributeInfo) Flags() FileAttributeInfoFlags {
	var ret FileAttributeInfoFlags
	ret = FileAttributeInfoFlags(f.native.flags)
	return ret
}

// FileAttributeInfoList acts as a lightweight registry for possible valid file
// attributes. The registry stores Key-Value pair formats as AttributeInfos.
type FileAttributeInfoList struct {
	native C.GFileAttributeInfoList
}

// WrapFileAttributeInfoList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeInfoList(ptr unsafe.Pointer) *FileAttributeInfoList {
	if ptr == nil {
		return nil
	}

	return (*FileAttributeInfoList)(ptr)
}

func marshalFileAttributeInfoList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFileAttributeInfoList(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FileAttributeInfoList) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// NewFileAttributeInfoList constructs a struct FileAttributeInfoList.
func NewFileAttributeInfoList() *FileAttributeInfoList {
	ret := C.g_file_attribute_info_list_new()

	var ret0 *FileAttributeInfoList

	{
		ret0 = WrapFileAttributeInfoList(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *FileAttributeInfoList) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Infos gets the field inside the struct.
func (f *FileAttributeInfoList) Infos() *FileAttributeInfo {
	var ret *FileAttributeInfo
	{
		ret = WrapFileAttributeInfo(unsafe.Pointer(f.native.infos))
	}
	return ret
}

// NInfos gets the field inside the struct.
func (f *FileAttributeInfoList) NInfos() int {
	var ret int
	ret = int(f.native.n_infos)
	return ret
}

// Add adds a new attribute with @name to the @list, setting its @type and
// @flags.
func (l *FileAttributeInfoList) Add(name string, typ FileAttributeType, flags FileAttributeInfoFlags) {
	var arg0 *C.GFileAttributeInfoList
	var arg1 *C.char
	var arg2 C.GFileAttributeType
	var arg3 C.GFileAttributeInfoFlags

	arg0 = (*C.GFileAttributeInfoList)(l.Native())
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GFileAttributeType)(typ)
	arg3 = (C.GFileAttributeInfoFlags)(flags)

	C.g_file_attribute_info_list_add(arg0, arg1, arg2, arg3)
}

// Dup makes a duplicate of a file attribute info list.
func (l *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	var arg0 *C.GFileAttributeInfoList

	arg0 = (*C.GFileAttributeInfoList)(l.Native())

	ret := C.g_file_attribute_info_list_dup(arg0)

	var ret0 *FileAttributeInfoList

	{
		ret0 = WrapFileAttributeInfoList(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *FileAttributeInfoList) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Lookup gets the file attribute with the name @name from @list.
func (l *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	var arg0 *C.GFileAttributeInfoList
	var arg1 *C.char

	arg0 = (*C.GFileAttributeInfoList)(l.Native())
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_file_attribute_info_list_lookup(arg0, arg1)

	var ret0 *FileAttributeInfo

	{
		ret0 = WrapFileAttributeInfo(unsafe.Pointer(ret))
	}

	return ret0
}

// Ref references a file attribute info list.
func (l *FileAttributeInfoList) Ref() *FileAttributeInfoList {
	var arg0 *C.GFileAttributeInfoList

	arg0 = (*C.GFileAttributeInfoList)(l.Native())

	ret := C.g_file_attribute_info_list_ref(arg0)

	var ret0 *FileAttributeInfoList

	{
		ret0 = WrapFileAttributeInfoList(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *FileAttributeInfoList) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Unref removes a reference from the given @list. If the reference count falls
// to zero, the @list is deleted.
func (l *FileAttributeInfoList) Unref() {
	var arg0 *C.GFileAttributeInfoList

	arg0 = (*C.GFileAttributeInfoList)(l.Native())

	C.g_file_attribute_info_list_unref(arg0)
}
