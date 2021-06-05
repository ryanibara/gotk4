// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_directory_list_get_type()), F: marshalDirectoryList},
	})
}

// DirectoryList is a list model that wraps g_file_enumerate_children_async().
// It presents a Model and fills it asynchronously with the Infos returned from
// that function.
//
// Enumeration will start automatically when a the DirectoryList:file property
// is set.
//
// While the DirectoryList is being filled, the DirectoryList:loading property
// will be set to true. You can listen to that property if you want to show
// information like a Spinner or a "Loading..." text.
//
// If loading fails at any point, the DirectoryList:error property will be set
// to give more indication about the failure.
//
// The Infos returned from a DirectoryList have the "standard::file" attribute
// set to the #GFile they refer to. This way you can get at the file that is
// referred to in the same way you would via g_file_enumerator_get_child(). This
// means you do not need access to the DirectoryList but can access the #GFile
// directly from the Info when operating with a ListView or similar.
type DirectoryList interface {
	gextras.Objector
	gio.ListModel

	// Attributes gets the attributes queried on the children.
	Attributes() string
	// Error gets the loading error, if any.
	//
	// If an error occurs during the loading process, the loading process will
	// finish and this property allows querying the error that happened. This
	// error will persist until a file is loaded again.
	//
	// An error being set does not mean that no files were loaded, and all
	// successfully queried files will remain in the list.
	Error() error
	// File gets the file whose children are currently enumerated.
	File() gio.File
	// IOPriority gets the IO priority set via
	// gtk_directory_list_set_io_priority().
	IOPriority() int
	// Monitored returns whether the directory list is monitoring the directory
	// for changes.
	Monitored() bool
	// IsLoading returns true if the children enumeration is currently in
	// progress.
	//
	// Files will be added to @self from time to time while loading is going on.
	// The order in which are added is undefined and may change in between runs.
	IsLoading() bool
	// SetAttributes sets the @attributes to be enumerated and starts the
	// enumeration.
	//
	// If @attributes is nil, no attributes will be queried, but a list of Infos
	// will still be created.
	SetAttributes(attributes string)
	// SetFile sets the @file to be enumerated and starts the enumeration.
	//
	// If @file is nil, the result will be an empty list.
	SetFile(file gio.File)
	// SetIOPriority sets the IO priority to use while loading directories.
	//
	// Setting the priority while @self is loading will reprioritize the ongoing
	// load as soon as possible.
	//
	// The default IO priority is G_PRIORITY_DEFAULT, which is higher than the
	// GTK redraw priority. If you are loading a lot of directories in parallel,
	// lowering it to something like G_PRIORITY_DEFAULT_IDLE may increase
	// responsiveness.
	SetIOPriority(ioPriority int)
	// SetMonitored sets whether the directory list will monitor the directory
	// for changes. If monitoring is enabled, the Model::items-changed signal
	// will be emitted when the directory contents change.
	//
	// When monitoring is turned on after the initial creation of the directory
	// list, the directory is reloaded to avoid missing files that appeared
	// between the initial loading and when monitoring was turned on.
	SetMonitored(monitored bool)
}

// directoryList implements the DirectoryList interface.
type directoryList struct {
	gextras.Objector
	gio.ListModel
}

var _ DirectoryList = (*directoryList)(nil)

// WrapDirectoryList wraps a GObject to the right type. It is
// primarily used internally.
func WrapDirectoryList(obj *externglib.Object) DirectoryList {
	return DirectoryList{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalDirectoryList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDirectoryList(obj), nil
}

// NewDirectoryList constructs a class DirectoryList.
func NewDirectoryList(attributes string, file gio.File) DirectoryList {
	var arg1 *C.char
	var arg2 *C.GFile

	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GFile)(unsafe.Pointer(file.Native()))

	var cret C.GtkDirectoryList
	var ret1 DirectoryList

	cret = C.gtk_directory_list_new(attributes, file)

	ret1 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DirectoryList)

	return ret1
}

// Attributes gets the attributes queried on the children.
func (s directoryList) Attributes() string {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret *C.char
	var ret1 string

	cret = C.gtk_directory_list_get_attributes(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Error gets the loading error, if any.
//
// If an error occurs during the loading process, the loading process will
// finish and this property allows querying the error that happened. This
// error will persist until a file is loaded again.
//
// An error being set does not mean that no files were loaded, and all
// successfully queried files will remain in the list.
func (s directoryList) Error() error {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret *C.GError
	var ret1 error

	cret = C.gtk_directory_list_get_error(arg0)

	ret1 = gerror.Take(unsafe.Pointer(cret))

	return ret1
}

// File gets the file whose children are currently enumerated.
func (s directoryList) File() gio.File {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret *C.GFile
	var ret1 gio.File

	cret = C.gtk_directory_list_get_file(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gio.File)

	return ret1
}

// IOPriority gets the IO priority set via
// gtk_directory_list_set_io_priority().
func (s directoryList) IOPriority() int {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret C.int
	var ret1 int

	cret = C.gtk_directory_list_get_io_priority(arg0)

	ret1 = C.int(cret)

	return ret1
}

// Monitored returns whether the directory list is monitoring the directory
// for changes.
func (s directoryList) Monitored() bool {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_directory_list_get_monitored(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// IsLoading returns true if the children enumeration is currently in
// progress.
//
// Files will be added to @self from time to time while loading is going on.
// The order in which are added is undefined and may change in between runs.
func (s directoryList) IsLoading() bool {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_directory_list_is_loading(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SetAttributes sets the @attributes to be enumerated and starts the
// enumeration.
//
// If @attributes is nil, no attributes will be queried, but a list of Infos
// will still be created.
func (s directoryList) SetAttributes(attributes string) {
	var arg0 *C.GtkDirectoryList
	var arg1 *C.char

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_directory_list_set_attributes(arg0, attributes)
}

// SetFile sets the @file to be enumerated and starts the enumeration.
//
// If @file is nil, the result will be an empty list.
func (s directoryList) SetFile(file gio.File) {
	var arg0 *C.GtkDirectoryList
	var arg1 *C.GFile

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_directory_list_set_file(arg0, file)
}

// SetIOPriority sets the IO priority to use while loading directories.
//
// Setting the priority while @self is loading will reprioritize the ongoing
// load as soon as possible.
//
// The default IO priority is G_PRIORITY_DEFAULT, which is higher than the
// GTK redraw priority. If you are loading a lot of directories in parallel,
// lowering it to something like G_PRIORITY_DEFAULT_IDLE may increase
// responsiveness.
func (s directoryList) SetIOPriority(ioPriority int) {
	var arg0 *C.GtkDirectoryList
	var arg1 C.int

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = C.int(ioPriority)

	C.gtk_directory_list_set_io_priority(arg0, ioPriority)
}

// SetMonitored sets whether the directory list will monitor the directory
// for changes. If monitoring is enabled, the Model::items-changed signal
// will be emitted when the directory contents change.
//
// When monitoring is turned on after the initial creation of the directory
// list, the directory is reloaded to avoid missing files that appeared
// between the initial loading and when monitoring was turned on.
func (s directoryList) SetMonitored(monitored bool) {
	var arg0 *C.GtkDirectoryList
	var arg1 C.gboolean

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	if monitored {
		arg1 = C.gboolean(1)
	}

	C.gtk_directory_list_set_monitored(arg0, monitored)
}
