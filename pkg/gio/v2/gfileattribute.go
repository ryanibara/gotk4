// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeFileAttributeInfoList = coreglib.Type(C.g_file_attribute_info_list_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeFileAttributeInfoList, F: marshalFileAttributeInfoList},
	})
}

// FileAttributeInfo: information about a specific attribute.
//
// An instance of this type is always passed by reference.
type FileAttributeInfo struct {
	*fileAttributeInfo
}

// fileAttributeInfo is the struct that's finalized.
type fileAttributeInfo struct {
	native *C.GFileAttributeInfo
}

// Name: name of the attribute.
func (f *FileAttributeInfo) Name() string {
	valptr := &f.native.name
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}

// Type type of the attribute.
func (f *FileAttributeInfo) Type() FileAttributeType {
	valptr := &f.native._type
	var v FileAttributeType // out
	v = FileAttributeType(*valptr)
	return v
}

// Flags: set of AttributeInfoFlags.
func (f *FileAttributeInfo) Flags() FileAttributeInfoFlags {
	valptr := &f.native.flags
	var v FileAttributeInfoFlags // out
	v = FileAttributeInfoFlags(*valptr)
	return v
}

// FileAttributeInfoList acts as a lightweight registry for possible valid file
// attributes. The registry stores Key-Value pair formats as AttributeInfos.
//
// An instance of this type is always passed by reference.
type FileAttributeInfoList struct {
	*fileAttributeInfoList
}

// fileAttributeInfoList is the struct that's finalized.
type fileAttributeInfoList struct {
	native *C.GFileAttributeInfoList
}

func marshalFileAttributeInfoList(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &FileAttributeInfoList{&fileAttributeInfoList{(*C.GFileAttributeInfoList)(b)}}, nil
}

// NewFileAttributeInfoList constructs a struct FileAttributeInfoList.
func NewFileAttributeInfoList() *FileAttributeInfoList {
	var _cret *C.GFileAttributeInfoList // in

	_cret = C.g_file_attribute_info_list_new()

	var _fileAttributeInfoList *FileAttributeInfoList // out

	_fileAttributeInfoList = (*FileAttributeInfoList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fileAttributeInfoList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _fileAttributeInfoList
}

// Infos: array of AttributeInfos.
func (f *FileAttributeInfoList) Infos() *FileAttributeInfo {
	valptr := &f.native.infos
	var v *FileAttributeInfo // out
	v = (*FileAttributeInfo)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return v
}

// NInfos: number of values in the array.
func (f *FileAttributeInfoList) NInfos() int {
	valptr := &f.native.n_infos
	var v int // out
	v = int(*valptr)
	return v
}

// NInfos: number of values in the array.
func (f *FileAttributeInfoList) SetNInfos(nInfos int) {
	valptr := &f.native.n_infos
	*valptr = C.int(nInfos)
}

// Add adds a new attribute with name to the list, setting its type and flags.
//
// The function takes the following parameters:
//
//    - name of the attribute to add.
//    - typ for the attribute.
//    - flags for the attribute.
//
func (list *FileAttributeInfoList) Add(name string, typ FileAttributeType, flags FileAttributeInfoFlags) {
	var _arg0 *C.GFileAttributeInfoList // out
	var _arg1 *C.char                   // out
	var _arg2 C.GFileAttributeType      // out
	var _arg3 C.GFileAttributeInfoFlags // out

	_arg0 = (*C.GFileAttributeInfoList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GFileAttributeType(typ)
	_arg3 = C.GFileAttributeInfoFlags(flags)

	C.g_file_attribute_info_list_add(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(list)
	runtime.KeepAlive(name)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(flags)
}

// Dup makes a duplicate of a file attribute info list.
//
// The function returns the following values:
//
//    - fileAttributeInfoList: copy of the given list.
//
func (list *FileAttributeInfoList) Dup() *FileAttributeInfoList {
	var _arg0 *C.GFileAttributeInfoList // out
	var _cret *C.GFileAttributeInfoList // in

	_arg0 = (*C.GFileAttributeInfoList)(gextras.StructNative(unsafe.Pointer(list)))

	_cret = C.g_file_attribute_info_list_dup(_arg0)
	runtime.KeepAlive(list)

	var _fileAttributeInfoList *FileAttributeInfoList // out

	_fileAttributeInfoList = (*FileAttributeInfoList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_fileAttributeInfoList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _fileAttributeInfoList
}

// Lookup gets the file attribute with the name name from list.
//
// The function takes the following parameters:
//
//    - name of the attribute to look up.
//
// The function returns the following values:
//
//    - fileAttributeInfo for the name, or NULL if an attribute isn't found.
//
func (list *FileAttributeInfoList) Lookup(name string) *FileAttributeInfo {
	var _arg0 *C.GFileAttributeInfoList // out
	var _arg1 *C.char                   // out
	var _cret *C.GFileAttributeInfo     // in

	_arg0 = (*C.GFileAttributeInfoList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_info_list_lookup(_arg0, _arg1)
	runtime.KeepAlive(list)
	runtime.KeepAlive(name)

	var _fileAttributeInfo *FileAttributeInfo // out

	_fileAttributeInfo = (*FileAttributeInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _fileAttributeInfo
}
