// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GTypeFileIcon returns the GType for the type FileIcon.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeFileIcon() coreglib.Type {
	gtype := coreglib.Type(C.g_file_icon_get_type())
	coreglib.RegisterGValueMarshaler(gtype, marshalFileIcon)
	return gtype
}

// FileIconOverrider contains methods that are overridable.
type FileIconOverrider interface {
}

// FileIcon specifies an icon by pointing to an image file to be used as icon.
type FileIcon struct {
	_ [0]func() // equal guard
	*coreglib.Object

	LoadableIcon
}

var (
	_ coreglib.Objector = (*FileIcon)(nil)
)

func classInitFileIconner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFileIcon(obj *coreglib.Object) *FileIcon {
	return &FileIcon{
		Object: obj,
		LoadableIcon: LoadableIcon{
			Icon: Icon{
				Object: obj,
			},
		},
	}
}

func marshalFileIcon(p uintptr) (interface{}, error) {
	return wrapFileIcon(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFileIcon creates a new icon for a file.
//
// The function takes the following parameters:
//
//    - file: #GFile.
//
// The function returns the following values:
//
//    - fileIcon for the given file, or NULL on error.
//
func NewFileIcon(file Filer) *FileIcon {
	var _arg1 *C.GFile // out
	var _cret *C.GIcon // in

	_arg1 = (*C.GFile)(unsafe.Pointer(coreglib.InternObject(file).Native()))

	_cret = C.g_file_icon_new(_arg1)
	runtime.KeepAlive(file)

	var _fileIcon *FileIcon // out

	_fileIcon = wrapFileIcon(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _fileIcon
}

// File gets the #GFile associated with the given icon.
//
// The function returns the following values:
//
//    - file: #GFile.
//
func (icon *FileIcon) File() *File {
	var _arg0 *C.GFileIcon // out
	var _cret *C.GFile     // in

	_arg0 = (*C.GFileIcon)(unsafe.Pointer(coreglib.InternObject(icon).Native()))

	_cret = C.g_file_icon_get_file(_arg0)
	runtime.KeepAlive(icon)

	var _file *File // out

	_file = wrapFile(coreglib.Take(unsafe.Pointer(_cret)))

	return _file
}
