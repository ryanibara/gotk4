// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_bookmark_list_get_type()), F: marshalBookmarkList},
	})
}

// BookmarkList: `GtkBookmarkList` is a list model that wraps `GBookmarkFile`.
//
// It presents a `GListModel` and fills it asynchronously with the `GFileInfo`s
// returned from that function.
//
// The `GFileInfo`s in the list have some attributes in the recent namespace
// added: `recent::private` (boolean) and `recent:applications` (stringv).
type BookmarkList interface {
	gextras.Objector
	gio.ListModel

	// Attributes gets the attributes queried on the children.
	Attributes() string
	// Filename returns the filename of the bookmark file that this list is
	// loading.
	Filename() string
	// IOPriority gets the IO priority to use while loading file.
	IOPriority() int
	// IsLoading returns true if the files are currently being loaded.
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
	// SetIOPriority sets the IO priority to use while loading files.
	//
	// The default IO priority is G_PRIORITY_DEFAULT.
	SetIOPriority(ioPriority int)
}

// bookmarkList implements the BookmarkList interface.
type bookmarkList struct {
	gextras.Objector
	gio.ListModel
}

var _ BookmarkList = (*bookmarkList)(nil)

// WrapBookmarkList wraps a GObject to the right type. It is
// primarily used internally.
func WrapBookmarkList(obj *externglib.Object) BookmarkList {
	return BookmarkList{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalBookmarkList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBookmarkList(obj), nil
}

// NewBookmarkList constructs a class BookmarkList.
func NewBookmarkList(filename string, attributes string) BookmarkList {
	var arg1 *C.char
	var arg2 *C.char

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.gtk_bookmark_list_new(arg1, arg2)

	var ret0 BookmarkList

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(BookmarkList)

	return ret0
}

// Attributes gets the attributes queried on the children.
func (s bookmarkList) Attributes() string {
	var arg0 *C.GtkBookmarkList

	arg0 = (*C.GtkBookmarkList)(s.Native())

	ret := C.gtk_bookmark_list_get_attributes(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Filename returns the filename of the bookmark file that this list is
// loading.
func (s bookmarkList) Filename() string {
	var arg0 *C.GtkBookmarkList

	arg0 = (*C.GtkBookmarkList)(s.Native())

	ret := C.gtk_bookmark_list_get_filename(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// IOPriority gets the IO priority to use while loading file.
func (s bookmarkList) IOPriority() int {
	var arg0 *C.GtkBookmarkList

	arg0 = (*C.GtkBookmarkList)(s.Native())

	ret := C.gtk_bookmark_list_get_io_priority(arg0)

	var ret0 int

	ret0 = int(ret)

	return ret0
}

// IsLoading returns true if the files are currently being loaded.
//
// Files will be added to @self from time to time while loading is going on.
// The order in which are added is undefined and may change in between runs.
func (s bookmarkList) IsLoading() bool {
	var arg0 *C.GtkBookmarkList

	arg0 = (*C.GtkBookmarkList)(s.Native())

	ret := C.gtk_bookmark_list_is_loading(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// SetAttributes sets the @attributes to be enumerated and starts the
// enumeration.
//
// If @attributes is nil, no attributes will be queried, but a list of Infos
// will still be created.
func (s bookmarkList) SetAttributes(attributes string) {
	var arg0 *C.GtkBookmarkList
	var arg1 *C.char

	arg0 = (*C.GtkBookmarkList)(s.Native())
	arg1 = (*C.gchar)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_bookmark_list_set_attributes(arg0, arg1)
}

// SetIOPriority sets the IO priority to use while loading files.
//
// The default IO priority is G_PRIORITY_DEFAULT.
func (s bookmarkList) SetIOPriority(ioPriority int) {
	var arg0 *C.GtkBookmarkList
	var arg1 C.int

	arg0 = (*C.GtkBookmarkList)(s.Native())
	arg1 = C.int(ioPriority)

	C.gtk_bookmark_list_set_io_priority(arg0, arg1)
}
