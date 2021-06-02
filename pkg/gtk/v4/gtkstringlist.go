// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"reflect"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_string_list_get_type()), F: marshalStringList},
		{T: externglib.Type(C.gtk_string_object_get_type()), F: marshalStringObject},
	})
}

// StringList: `GtkStringList` is a list model that wraps an array of strings.
//
// The objects in the model have a "string" property.
//
// `GtkStringList` is well-suited for any place where you would typically use a
// `char*[]`, but need a list model.
//
//
// GtkStringList as GtkBuildable
//
// The `GtkStringList` implementation of the `GtkBuildable` interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying a `GtkStringList`
//
// “`xml <object class="GtkStringList"> <items> <item
// translatable="yes">Factory</item> <item translatable="yes">Home</item> <item
// translatable="yes">Subway</item> </items> </object> “`
type StringList interface {
	gextras.Objector
	gio.ListModel
	Buildable

	// Append appends @string to @self.
	//
	// The @string will be copied. See [method@Gtk.StringList.take] for a way to
	// avoid that.
	Append(string string)
	// String gets the string that is at @position in @self.
	//
	// If @self does not contain @position items, nil is returned.
	//
	// This function returns the const char *. To get the object wrapping it,
	// use g_list_model_get_item().
	String(position uint) string
	// Remove removes the string at @position from @self.
	//
	// @position must be smaller than the current length of the list.
	Remove(position uint)
	// Splice changes @self by removing @n_removals strings and adding
	// @additions to it.
	//
	// This function is more efficient than [method@Gtk.StringList.append] and
	// [method@Gtk.StringList.remove], because it only emits the ::items-changed
	// signal once for the change.
	//
	// This function copies the strings in @additions.
	//
	// The parameters @position and @n_removals must be correct (ie: @position +
	// @n_removals must be less than or equal to the length of the list at the
	// time this function is called).
	Splice(position uint, nRemovals uint, additions []string)
	// Take adds @string to self at the end, and takes ownership of it.
	//
	// This variant of [method@Gtk.StringList.append] is convenient for
	// formatting strings:
	//
	// “`c gtk_string_list_take (self, g_strdup_print ("d dollars", lots)); “`
	Take(string string)
}

// stringList implements the StringList interface.
type stringList struct {
	gextras.Objector
	gio.ListModel
	Buildable
}

var _ StringList = (*stringList)(nil)

// WrapStringList wraps a GObject to the right type. It is
// primarily used internally.
func WrapStringList(obj *externglib.Object) StringList {
	return StringList{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalStringList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStringList(obj), nil
}

// NewStringList constructs a class StringList.
func NewStringList(strings []string) StringList {
	var arg1 **C.char

	{
		var dst []*C.char
		ptr := C.malloc(unsafe.Sizeof((*struct{})(nil)) * (len(strings) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(strings)
		sliceHeader.Cap = len(strings)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(strings); i++ {
			src := strings[i]
			dst[i] = (*C.gchar)(C.CString(src))
			defer C.free(unsafe.Pointer(dst[i]))
		}

		arg1 = (**C.char)(unsafe.Pointer(ptr))
	}

	ret := C.gtk_string_list_new(arg1)

	var ret0 StringList

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StringList)

	return ret0
}

// Append appends @string to @self.
//
// The @string will be copied. See [method@Gtk.StringList.take] for a way to
// avoid that.
func (s stringList) Append(string string) {
	var arg0 *C.GtkStringList
	var arg1 *C.char

	arg0 = (*C.GtkStringList)(s.Native())
	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_string_list_append(arg0, arg1)
}

// String gets the string that is at @position in @self.
//
// If @self does not contain @position items, nil is returned.
//
// This function returns the const char *. To get the object wrapping it,
// use g_list_model_get_item().
func (s stringList) String(position uint) string {
	var arg0 *C.GtkStringList
	var arg1 C.guint

	arg0 = (*C.GtkStringList)(s.Native())
	arg1 = C.guint(position)

	ret := C.gtk_string_list_get_string(arg0, arg1)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}

// Remove removes the string at @position from @self.
//
// @position must be smaller than the current length of the list.
func (s stringList) Remove(position uint) {
	var arg0 *C.GtkStringList
	var arg1 C.guint

	arg0 = (*C.GtkStringList)(s.Native())
	arg1 = C.guint(position)

	C.gtk_string_list_remove(arg0, arg1)
}

// Splice changes @self by removing @n_removals strings and adding
// @additions to it.
//
// This function is more efficient than [method@Gtk.StringList.append] and
// [method@Gtk.StringList.remove], because it only emits the ::items-changed
// signal once for the change.
//
// This function copies the strings in @additions.
//
// The parameters @position and @n_removals must be correct (ie: @position +
// @n_removals must be less than or equal to the length of the list at the
// time this function is called).
func (s stringList) Splice(position uint, nRemovals uint, additions []string) {
	var arg0 *C.GtkStringList
	var arg1 C.guint
	var arg2 C.guint
	var arg3 **C.char

	arg0 = (*C.GtkStringList)(s.Native())
	arg1 = C.guint(position)
	arg2 = C.guint(nRemovals)
	{
		var dst []*C.char
		ptr := C.malloc(unsafe.Sizeof((*struct{})(nil)) * (len(additions) + 1))
		sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&dst))
		sliceHeader.Data = uintptr(unsafe.Pointer(ptr))
		sliceHeader.Len = len(additions)
		sliceHeader.Cap = len(additions)
		defer C.free(unsafe.Pointer(ptr))

		for i := 0; i < len(additions); i++ {
			src := additions[i]
			dst[i] = (*C.gchar)(C.CString(src))
			defer C.free(unsafe.Pointer(dst[i]))
		}

		arg3 = (**C.char)(unsafe.Pointer(ptr))
	}

	C.gtk_string_list_splice(arg0, arg1, arg2, arg3)
}

// Take adds @string to self at the end, and takes ownership of it.
//
// This variant of [method@Gtk.StringList.append] is convenient for
// formatting strings:
//
// “`c gtk_string_list_take (self, g_strdup_print ("d dollars", lots)); “`
func (s stringList) Take(string string) {
	var arg0 *C.GtkStringList
	var arg1 *C.char

	arg0 = (*C.GtkStringList)(s.Native())
	arg1 = (*C.gchar)(C.CString(string))

	C.gtk_string_list_take(arg0, arg1)
}

// StringObject: `GtkStringObject` is the type of items in a `GtkStringList`.
//
// A `GtkStringObject` is a wrapper around a `const char*`; it has a
// [property@Gtk.StringObject:string] property.
type StringObject interface {
	gextras.Objector

	// String returns the string contained in a `GtkStringObject`.
	String() string
}

// stringObject implements the StringObject interface.
type stringObject struct {
	gextras.Objector
}

var _ StringObject = (*stringObject)(nil)

// WrapStringObject wraps a GObject to the right type. It is
// primarily used internally.
func WrapStringObject(obj *externglib.Object) StringObject {
	return StringObject{
		Objector: obj,
	}
}

func marshalStringObject(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStringObject(obj), nil
}

// NewStringObject constructs a class StringObject.
func NewStringObject(string string) StringObject {
	var arg1 *C.char

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.gtk_string_object_new(arg1)

	var ret0 StringObject

	ret0 = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(ret.Native()))).(StringObject)

	return ret0
}

// String returns the string contained in a `GtkStringObject`.
func (s stringObject) String() string {
	var arg0 *C.GtkStringObject

	arg0 = (*C.GtkStringObject)(s.Native())

	ret := C.gtk_string_object_get_string(arg0)

	var ret0 string

	ret0 = C.GoString(ret)

	return ret0
}
