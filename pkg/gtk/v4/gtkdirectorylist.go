// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// GTypeDirectoryList returns the GType for the type DirectoryList.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDirectoryList() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "DirectoryList").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDirectoryList)
	return gtype
}

// DirectoryListOverrider contains methods that are overridable.
type DirectoryListOverrider interface {
}

// DirectoryList: GtkDirectoryList is a list model that wraps
// g_file_enumerate_children_async().
//
// It presents a GListModel and fills it asynchronously with the GFileInfos
// returned from that function.
//
// Enumeration will start automatically when a the gtk.DirectoryList:file
// property is set.
//
// While the GtkDirectoryList is being filled, the gtk.DirectoryList:loading
// property will be set to TRUE. You can listen to that property if you want to
// show information like a GtkSpinner or a "Loading..." text.
//
// If loading fails at any point, the gtk.DirectoryList:error property will be
// set to give more indication about the failure.
//
// The GFileInfos returned from a GtkDirectoryList have the "standard::file"
// attribute set to the GFile they refer to. This way you can get at the file
// that is referred to in the same way you would via
// g_file_enumerator_get_child(). This means you do not need access to the
// GtkDirectoryList, but can access the GFile directly from the GFileInfo when
// operating with a GtkListView or similar.
type DirectoryList struct {
	_ [0]func() // equal guard
	*coreglib.Object

	gio.ListModel
}

var (
	_ coreglib.Objector = (*DirectoryList)(nil)
)

func classInitDirectoryLister(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapDirectoryList(obj *coreglib.Object) *DirectoryList {
	return &DirectoryList{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalDirectoryList(p uintptr) (interface{}, error) {
	return wrapDirectoryList(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewDirectoryList creates a new GtkDirectoryList.
//
// The GtkDirectoryList is querying the given file with the given attributes.
//
// The function takes the following parameters:
//
//    - attributes (optional) to query with.
//    - file (optional) to query.
//
// The function returns the following values:
//
//    - directoryList: new GtkDirectoryList.
//
func NewDirectoryList(attributes string, file gio.Filer) *DirectoryList {
	var _args [2]girepository.Argument

	if attributes != "" {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(attributes)))
		defer C.free(unsafe.Pointer(_args[0]))
	}
	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("new_DirectoryList", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(attributes)
	runtime.KeepAlive(file)

	var _directoryList *DirectoryList // out

	_directoryList = wrapDirectoryList(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _directoryList
}

// Attributes gets the attributes queried on the children.
//
// The function returns the following values:
//
//    - utf8 (optional): queried attributes.
//
func (self *DirectoryList) Attributes() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("get_attributes", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Error gets the loading error, if any.
//
// If an error occurs during the loading process, the loading process will
// finish and this property allows querying the error that happened. This error
// will persist until a file is loaded again.
//
// An error being set does not mean that no files were loaded, and all
// successfully queried files will remain in the list.
//
// The function returns the following values:
//
//    - err (optional): loading error or NULL if loading finished successfully.
//
func (self *DirectoryList) Error() error {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("get_error", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _err error // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_err = gerror.Take(unsafe.Pointer(_cret))
	}

	return _err
}

// File gets the file whose children are currently enumerated.
//
// The function returns the following values:
//
//    - file (optional) whose children are enumerated.
//
func (self *DirectoryList) File() *gio.File {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("get_file", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _file *gio.File // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			obj := coreglib.Take(unsafe.Pointer(_cret))
			_file = &gio.File{
				Object: obj,
			}
		}
	}

	return _file
}

// IOPriority gets the IO priority set via gtk_directory_list_set_io_priority().
//
// The function returns the following values:
//
//    - gint: IO priority.
//
func (self *DirectoryList) IOPriority() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("get_io_priority", _args[:], nil)
	_cret = *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Monitored returns whether the directory list is monitoring the directory for
// changes.
//
// The function returns the following values:
//
//    - ok: TRUE if the directory is monitored.
//
func (self *DirectoryList) Monitored() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("get_monitored", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// IsLoading returns TRUE if the children enumeration is currently in progress.
//
// Files will be added to self from time to time while loading is going on. The
// order in which are added is undefined and may change in between runs.
//
// The function returns the following values:
//
//    - ok: TRUE if self is loading.
//
func (self *DirectoryList) IsLoading() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("is_loading", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// SetAttributes sets the attributes to be enumerated and starts the
// enumeration.
//
// If attributes is NULL, no attributes will be queried, but a list of
// GFileInfos will still be created.
//
// The function takes the following parameters:
//
//    - attributes (optional) to enumerate.
//
func (self *DirectoryList) SetAttributes(attributes string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if attributes != "" {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(attributes)))
		defer C.free(unsafe.Pointer(_args[1]))
	}

	girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("set_attributes", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(attributes)
}

// SetFile sets the file to be enumerated and starts the enumeration.
//
// If file is NULL, the result will be an empty list.
//
// The function takes the following parameters:
//
//    - file (optional) to be enumerated.
//
func (self *DirectoryList) SetFile(file gio.Filer) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if file != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(file).Native()))
	}

	girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("set_file", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(file)
}

// SetIOPriority sets the IO priority to use while loading directories.
//
// Setting the priority while self is loading will reprioritize the ongoing load
// as soon as possible.
//
// The default IO priority is G_PRIORITY_DEFAULT, which is higher than the GTK
// redraw priority. If you are loading a lot of directories in parallel,
// lowering it to something like G_PRIORITY_DEFAULT_IDLE may increase
// responsiveness.
//
// The function takes the following parameters:
//
//    - ioPriority: IO priority to use.
//
func (self *DirectoryList) SetIOPriority(ioPriority int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(ioPriority)

	girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("set_io_priority", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(ioPriority)
}

// SetMonitored sets whether the directory list will monitor the directory for
// changes. If monitoring is enabled, the ::items-changed signal will be emitted
// when the directory contents change.
//
// When monitoring is turned on after the initial creation of the directory
// list, the directory is reloaded to avoid missing files that appeared between
// the initial loading and when monitoring was turned on.
//
// The function takes the following parameters:
//
//    - monitored: TRUE to monitor the directory for changes.
//
func (self *DirectoryList) SetMonitored(monitored bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	if monitored {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "DirectoryList").InvokeMethod("set_monitored", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(monitored)
}
