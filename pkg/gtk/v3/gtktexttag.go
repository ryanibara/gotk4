// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_TextTagClass_event(GtkTextTag*, GObject*, GdkEvent*, GtkTextIter*);
// extern gboolean _gotk4_gtk3_TextTag_ConnectEvent(gpointer, GObject, GdkEvent, GtkTextIter*, guintptr);
import "C"

// glib.Type values for gtktexttag.go.
var GTypeTextTag = externglib.Type(C.gtk_text_tag_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeTextTag, F: marshalTextTag},
	})
}

// TextTagOverrider contains methods that are overridable.
type TextTagOverrider interface {
	externglib.Objector
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
	Event(eventObject *externglib.Object, event *gdk.Event, iter *TextIter) bool
}

// WrapTextTagOverrider wraps the TextTagOverrider
// interface implementation to access the instance methods.
func WrapTextTagOverrider(obj TextTagOverrider) *TextTag {
	return wrapTextTag(externglib.BaseObject(obj))
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
	*externglib.Object
}

var (
	_ externglib.Objector = (*TextTag)(nil)
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
		Event(eventObject *externglib.Object, event *gdk.Event, iter *TextIter) bool
	}); ok {
		pclass.event = (*[0]byte)(C._gotk4_gtk3_TextTagClass_event)
	}
}

//export _gotk4_gtk3_TextTagClass_event
func _gotk4_gtk3_TextTagClass_event(arg0 *C.GtkTextTag, arg1 *C.GObject, arg2 *C.GdkEvent, arg3 *C.GtkTextIter) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		Event(eventObject *externglib.Object, event *gdk.Event, iter *TextIter) bool
	})

	var _eventObject *externglib.Object // out
	var _event *gdk.Event               // out
	var _iter *TextIter                 // out

	_eventObject = externglib.Take(unsafe.Pointer(arg1))
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

func wrapTextTag(obj *externglib.Object) *TextTag {
	return &TextTag{
		Object: obj,
	}
}

func marshalTextTag(p uintptr) (interface{}, error) {
	return wrapTextTag(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_TextTag_ConnectEvent
func _gotk4_gtk3_TextTag_ConnectEvent(arg0 C.gpointer, arg1 C.GObject, arg2 C.GdkEvent, arg3 *C.GtkTextIter, arg4 C.guintptr) (cret C.gboolean) {
	var f func(object *externglib.Object, event *gdk.Event, iter *TextIter) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg4))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object *externglib.Object, event *gdk.Event, iter *TextIter) (ok bool))
	}

	var _object *externglib.Object // out
	var _event *gdk.Event          // out
	var _iter *TextIter            // out

	_object = externglib.Take(unsafe.Pointer(&arg1))
	{
		v := (*gdk.Event)(gextras.NewStructNative(unsafe.Pointer((&arg2))))
		v = gdk.CopyEventer(v)
		_event = v
	}
	_iter = (*TextIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	ok := f(_object, _event, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectEvent signal is emitted when an event occurs on a region of the buffer
// marked with this tag.
func (tag *TextTag) ConnectEvent(f func(object *externglib.Object, event *gdk.Event, iter *TextIter) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(tag, "event", false, unsafe.Pointer(C._gotk4_gtk3_TextTag_ConnectEvent), f)
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
	var _arg1 *C.gchar      // out
	var _cret *C.GtkTextTag // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_text_tag_new(_arg1)
	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	_textTag = wrapTextTag(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

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
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))
	if sizeChanged {
		_arg1 = C.TRUE
	}

	C.gtk_text_tag_changed(_arg0, _arg1)
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
func (tag *TextTag) Event(eventObject *externglib.Object, event *gdk.Event, iter *TextIter) bool {
	var _arg0 *C.GtkTextTag  // out
	var _arg1 *C.GObject     // out
	var _arg2 *C.GdkEvent    // out
	var _arg3 *C.GtkTextIter // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(eventObject.Native()))
	_arg2 = (*C.GdkEvent)(gextras.StructNative(unsafe.Pointer(event)))
	_arg3 = (*C.GtkTextIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.gtk_text_tag_event(_arg0, _arg1, _arg2, _arg3)
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
func (tag *TextTag) Priority() int {
	var _arg0 *C.GtkTextTag // out
	var _cret C.gint        // in

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))

	_cret = C.gtk_text_tag_get_priority(_arg0)
	runtime.KeepAlive(tag)

	var _gint int // out

	_gint = int(_cret)

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
func (tag *TextTag) SetPriority(priority int) {
	var _arg0 *C.GtkTextTag // out
	var _arg1 C.gint        // out

	_arg0 = (*C.GtkTextTag)(unsafe.Pointer(externglib.InternObject(tag).Native()))
	_arg1 = C.gint(priority)

	C.gtk_text_tag_set_priority(_arg0, _arg1)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(priority)
}
