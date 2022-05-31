// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_TextTagClass_event(GtkTextTag*, GObject*, GdkEvent*, GtkTextIter*);
import "C"

// glib.Type values for gtktexttag.go.
var GTypeTextTag = coreglib.Type(C.gtk_text_tag_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeTextTag, F: marshalTextTag},
	})
}

// TextTagOverrider contains methods that are overridable.
type TextTagOverrider interface {
	// Event emits the “event” signal on the TextTag.
	//
	// The function takes the following parameters:
	//
	//    - eventObject: object that received the event, such as a widget.
	//    - event: event.
	//    - iter: location where the event was received.
	//
	// The function returns the following values:
	//
	//    - ok: result of signal emission (whether the event was handled).
	//
	Event(eventObject *coreglib.Object, event *gdk.Event, iter *TextIter) bool
}

// TextTag: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
// Tags should be in the TextTagTable for a given TextBuffer before using them
// with that buffer.
//
// gtk_text_buffer_create_tag() is the best way to create tags. See “gtk3-demo”
// for numerous examples.
//
// For each property of TextTag, there is a “set” property, e.g. “font-set”
// corresponds to “font”. These “set” properties reflect whether a property has
// been set or not. They are maintained by GTK+ and you should not set them
// independently.
type TextTag struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TextTag)(nil)
)

func classInitTextTagger(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkTextTagClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkTextTagClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface {
		Event(eventObject *coreglib.Object, event *gdk.Event, iter *TextIter) bool
	}); ok {
		pclass.event = (*[0]byte)(C._gotk4_gtk3_TextTagClass_event)
	}
}

//export _gotk4_gtk3_TextTagClass_event
func _gotk4_gtk3_TextTagClass_event(arg0 *C.GtkTextTag, arg1 *C.GObject, arg2 *C.GdkEvent, arg3 *C.GtkTextIter) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Event(eventObject *coreglib.Object, event *gdk.Event, iter *TextIter) bool
	})

	var _eventObject *coreglib.Object // out
	var _event *gdk.Event             // out
	var _iter *TextIter               // out

	_eventObject = coreglib.Take(unsafe.Pointer(arg1))
	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer(arg2)))
		v = gdk.CopyEventer(v)
		_event = v
	}
	_iter = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	ok := iface.Event(_eventObject, _event, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTextTag(obj *coreglib.Object) *TextTag {
	return &TextTag{
		Object: obj,
	}
}

func marshalTextTag(p uintptr) (interface{}, error) {
	return wrapTextTag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewTextTag creates a TextTag. Configure the tag using object arguments, i.e.
// using g_object_set().
//
// The function takes the following parameters:
//
//    - name (optional): tag name, or NULL.
//
// The function returns the following values:
//
//    - textTag: new TextTag.
//
func NewTextTag(name string) *TextTag {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret *C.void // in

	if name != "" {
		_arg0 = (*C.void)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg0))
	}
	*(*string)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TextTag").InvokeMethod("new_TextTag", args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	_textTag = wrapTextTag(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTag
}

// Changed emits the TextTagTable::tag-changed signal on the TextTagTable where
// the tag is included.
//
// The signal is already emitted when setting a TextTag property. This function
// is useful for a TextTag subclass.
//
// The function takes the following parameters:
//
//    - sizeChanged: whether the change affects the TextView layout.
//
func (tag *TextTag) Changed(sizeChanged bool) {
	var args [2]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 C.gboolean // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	if sizeChanged {
		_arg1 = C.TRUE
	}
	*(**TextTag)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "TextTag").InvokeMethod("changed", args[:], nil)

	runtime.KeepAlive(tag)
	runtime.KeepAlive(sizeChanged)
}

// Event emits the “event” signal on the TextTag.
//
// The function takes the following parameters:
//
//    - eventObject: object that received the event, such as a widget.
//    - event: event.
//    - iter: location where the event was received.
//
// The function returns the following values:
//
//    - ok: result of signal emission (whether the event was handled).
//
func (tag *TextTag) Event(eventObject *coreglib.Object, event *gdk.Event, iter *TextIter) bool {
	var args [4]girepository.Argument
	var _arg0 *C.void    // out
	var _arg1 *C.void    // out
	var _arg2 *C.void    // out
	var _arg3 *C.void    // out
	var _cret C.gboolean // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(eventObject.Native()))
	_arg2 = (*C.void)(gextras.StructNative(unsafe.Pointer(event)))
	_arg3 = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**TextTag)(unsafe.Pointer(&args[1])) = _arg1
	*(**coreglib.Object)(unsafe.Pointer(&args[2])) = _arg2
	*(**gdk.Event)(unsafe.Pointer(&args[3])) = _arg3

	_gret := girepository.MustFind("Gtk", "TextTag").InvokeMethod("event", args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tag)
	runtime.KeepAlive(eventObject)
	runtime.KeepAlive(event)
	runtime.KeepAlive(iter)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Priority: get the tag priority.
//
// The function returns the following values:
//
//    - gint tag’s priority.
//
func (tag *TextTag) Priority() int32 {
	var args [1]girepository.Argument
	var _arg0 *C.void // out
	var _cret C.gint  // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	*(**TextTag)(unsafe.Pointer(&args[0])) = _arg0

	_gret := girepository.MustFind("Gtk", "TextTag").InvokeMethod("get_priority", args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(tag)

	var _gint int32 // out

	_gint = int32(_cret)

	return _gint
}

// SetPriority sets the priority of a TextTag. Valid priorities start at 0 and
// go to one less than gtk_text_tag_table_get_size(). Each tag in a table has a
// unique priority; setting the priority of one tag shifts the priorities of all
// the other tags in the table to maintain a unique priority for each tag.
// Higher priority tags “win” if two tags both set the same text attribute. When
// adding a tag to a tag table, it will be assigned the highest priority in the
// table by default; so normally the precedence of a set of tags is the order in
// which they were added to the table, or created with
// gtk_text_buffer_create_tag(), which adds the tag to the buffer’s table
// automatically.
//
// The function takes the following parameters:
//
//    - priority: new priority.
//
func (tag *TextTag) SetPriority(priority int32) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 C.gint  // out

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(tag).Native()))
	_arg1 = C.gint(priority)
	*(**TextTag)(unsafe.Pointer(&args[1])) = _arg1

	girepository.MustFind("Gtk", "TextTag").InvokeMethod("set_priority", args[:], nil)

	runtime.KeepAlive(tag)
	runtime.KeepAlive(priority)
}
