// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_target_flags_get_type()), F: marshalTargetFlags},
		{T: externglib.Type(C.gtk_target_entry_get_type()), F: marshalTargetEntry},
		{T: externglib.Type(C.gtk_target_list_get_type()), F: marshalTargetList},
	})
}

// TargetFlags enumeration is used to specify constraints on a TargetEntry.
type TargetFlags C.guint

const (
	// TargetSameApp: if this is set, the target will only be selected for drags
	// within a single application.
	TargetSameApp TargetFlags = 0b1
	// TargetSameWidget: if this is set, the target will only be selected for
	// drags within a single widget.
	TargetSameWidget TargetFlags = 0b10
	// TargetOtherApp: if this is set, the target will not be selected for drags
	// within a single application.
	TargetOtherApp TargetFlags = 0b100
	// TargetOtherWidget: if this is set, the target will not be selected for
	// drags withing a single widget.
	TargetOtherWidget TargetFlags = 0b1000
)

func marshalTargetFlags(p uintptr) (interface{}, error) {
	return TargetFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for TargetFlags.
func (t TargetFlags) String() string {
	if t == 0 {
		return "TargetFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(63)

	for t != 0 {
		next := t & (t - 1)
		bit := t - next

		switch bit {
		case TargetSameApp:
			builder.WriteString("SameApp|")
		case TargetSameWidget:
			builder.WriteString("SameWidget|")
		case TargetOtherApp:
			builder.WriteString("OtherApp|")
		case TargetOtherWidget:
			builder.WriteString("OtherWidget|")
		default:
			builder.WriteString(fmt.Sprintf("TargetFlags(0b%b)|", bit))
		}

		t = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if t contains other.
func (t TargetFlags) Has(other TargetFlags) bool {
	return (t & other) == other
}

// SelectionRemoveAll removes all handlers and unsets ownership of all
// selections for a widget. Called when widget is being destroyed. This function
// will not generally be called by applications.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
func SelectionRemoveAll(widget Widgetter) {
	var _arg1 *C.GtkWidget // out

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_selection_remove_all(_arg1)
	runtime.KeepAlive(widget)
}

// TargetTableNewFromList: this function creates an TargetEntry array that
// contains the same targets as the passed list. The returned table is newly
// allocated and should be freed using gtk_target_table_free() when no longer
// needed.
//
// The function takes the following parameters:
//
//    - list: TargetList.
//
func TargetTableNewFromList(list *TargetList) []TargetEntry {
	var _arg1 *C.GtkTargetList  // out
	var _cret *C.GtkTargetEntry // in
	var _arg2 C.gint            // in

	_arg1 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))

	_cret = C.gtk_target_table_new_from_list(_arg1, &_arg2)
	runtime.KeepAlive(list)

	var _targetEntrys []TargetEntry // out

	defer C.free(unsafe.Pointer(_cret))
	{
		src := unsafe.Slice(_cret, _arg2)
		_targetEntrys = make([]TargetEntry, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_targetEntrys[i] = *(*TargetEntry)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(&_targetEntrys[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.gtk_target_entry_free((*C.GtkTargetEntry)(intern.C))
				},
			)
		}
	}

	return _targetEntrys
}

// TargetEntry represents a single type of data than can be supplied for by a
// widget for a selection or for supplied or received during drag-and-drop.
//
// An instance of this type is always passed by reference.
type TargetEntry struct {
	*targetEntry
}

// targetEntry is the struct that's finalized.
type targetEntry struct {
	native *C.GtkTargetEntry
}

func marshalTargetEntry(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TargetEntry{&targetEntry{(*C.GtkTargetEntry)(b)}}, nil
}

// NewTargetEntry constructs a struct TargetEntry.
func NewTargetEntry(target string, flags uint, info uint) *TargetEntry {
	var _arg1 *C.gchar          // out
	var _arg2 C.guint           // out
	var _arg3 C.guint           // out
	var _cret *C.GtkTargetEntry // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(target)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(flags)
	_arg3 = C.guint(info)

	_cret = C.gtk_target_entry_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(target)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(info)

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetEntry)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_target_entry_free((*C.GtkTargetEntry)(intern.C))
		},
	)

	return _targetEntry
}

// Target: string representation of the target type.
func (t *TargetEntry) Target() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(t.native.target)))
	return v
}

// Flags for DND.
func (t *TargetEntry) Flags() uint {
	var v uint // out
	v = uint(t.native.flags)
	return v
}

// Info: application-assigned integer ID which will get passed as a parameter to
// e.g the Widget::selection-get signal. It allows the application to identify
// the target type without extensive string compares.
func (t *TargetEntry) Info() uint {
	var v uint // out
	v = uint(t.native.info)
	return v
}

// Copy makes a copy of a TargetEntry and its data.
func (data *TargetEntry) Copy() *TargetEntry {
	var _arg0 *C.GtkTargetEntry // out
	var _cret *C.GtkTargetEntry // in

	_arg0 = (*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer(data)))

	_cret = C.gtk_target_entry_copy(_arg0)
	runtime.KeepAlive(data)

	var _targetEntry *TargetEntry // out

	_targetEntry = (*TargetEntry)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetEntry)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_target_entry_free((*C.GtkTargetEntry)(intern.C))
		},
	)

	return _targetEntry
}

// TargetList is a reference counted list of TargetPair and should be treated as
// opaque.
//
// An instance of this type is always passed by reference.
type TargetList struct {
	*targetList
}

// targetList is the struct that's finalized.
type targetList struct {
	native *C.GtkTargetList
}

func marshalTargetList(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TargetList{&targetList{(*C.GtkTargetList)(b)}}, nil
}

// NewTargetList constructs a struct TargetList.
func NewTargetList(targets []TargetEntry) *TargetList {
	var _arg1 *C.GtkTargetEntry // out
	var _arg2 C.guint
	var _cret *C.GtkTargetList // in

	_arg2 = (C.guint)(len(targets))
	_arg1 = (*C.GtkTargetEntry)(C.calloc(C.size_t(len(targets)), C.size_t(C.sizeof_GtkTargetEntry)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GtkTargetEntry)(_arg1), len(targets))
		for i := range targets {
			out[i] = *(*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer((&targets[i]))))
		}
	}

	_cret = C.gtk_target_list_new(_arg1, _arg2)
	runtime.KeepAlive(targets)

	var _targetList *TargetList // out

	_targetList = (*TargetList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_targetList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_target_list_unref((*C.GtkTargetList)(intern.C))
		},
	)

	return _targetList
}

// AddImageTargets appends the image targets supported by SelectionData to the
// target list. All targets are added with the same info.
func (list *TargetList) AddImageTargets(info uint, writable bool) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out
	var _arg2 C.gboolean       // out

	_arg0 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = C.guint(info)
	if writable {
		_arg2 = C.TRUE
	}

	C.gtk_target_list_add_image_targets(_arg0, _arg1, _arg2)
	runtime.KeepAlive(list)
	runtime.KeepAlive(info)
	runtime.KeepAlive(writable)
}

// AddRichTextTargets appends the rich text targets registered with
// gtk_text_buffer_register_serialize_format() or
// gtk_text_buffer_register_deserialize_format() to the target list. All targets
// are added with the same info.
func (list *TargetList) AddRichTextTargets(info uint, deserializable bool, buffer *TextBuffer) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out
	var _arg2 C.gboolean       // out
	var _arg3 *C.GtkTextBuffer // out

	_arg0 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = C.guint(info)
	if deserializable {
		_arg2 = C.TRUE
	}
	_arg3 = (*C.GtkTextBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_target_list_add_rich_text_targets(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(list)
	runtime.KeepAlive(info)
	runtime.KeepAlive(deserializable)
	runtime.KeepAlive(buffer)
}

// AddTable prepends a table of TargetEntry to a target list.
func (list *TargetList) AddTable(targets []TargetEntry) {
	var _arg0 *C.GtkTargetList  // out
	var _arg1 *C.GtkTargetEntry // out
	var _arg2 C.guint

	_arg0 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg2 = (C.guint)(len(targets))
	_arg1 = (*C.GtkTargetEntry)(C.calloc(C.size_t(len(targets)), C.size_t(C.sizeof_GtkTargetEntry)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((*C.GtkTargetEntry)(_arg1), len(targets))
		for i := range targets {
			out[i] = *(*C.GtkTargetEntry)(gextras.StructNative(unsafe.Pointer((&targets[i]))))
		}
	}

	C.gtk_target_list_add_table(_arg0, _arg1, _arg2)
	runtime.KeepAlive(list)
	runtime.KeepAlive(targets)
}

// AddTextTargets appends the text targets supported by SelectionData to the
// target list. All targets are added with the same info.
func (list *TargetList) AddTextTargets(info uint) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = C.guint(info)

	C.gtk_target_list_add_text_targets(_arg0, _arg1)
	runtime.KeepAlive(list)
	runtime.KeepAlive(info)
}

// AddURITargets appends the URI targets supported by SelectionData to the
// target list. All targets are added with the same info.
func (list *TargetList) AddURITargets(info uint) {
	var _arg0 *C.GtkTargetList // out
	var _arg1 C.guint          // out

	_arg0 = (*C.GtkTargetList)(gextras.StructNative(unsafe.Pointer(list)))
	_arg1 = C.guint(info)

	C.gtk_target_list_add_uri_targets(_arg0, _arg1)
	runtime.KeepAlive(list)
	runtime.KeepAlive(info)
}

// TargetPair is used to represent the same information as a table of
// TargetEntry, but in an efficient form.
//
// An instance of this type is always passed by reference.
type TargetPair struct {
	*targetPair
}

// targetPair is the struct that's finalized.
type targetPair struct {
	native *C.GtkTargetPair
}
