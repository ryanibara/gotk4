// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_completion_get_type()), F: marshalEntryCompletion},
	})
}

// EntryCompletionMatchFunc: a function which decides whether the row indicated
// by @iter matches a given @key, and should be displayed as a possible
// completion for @key.
//
// Note that @key is normalized and case-folded (see g_utf8_normalize() and
// g_utf8_casefold()). If this is not appropriate, match functions have access
// to the unmodified key via `gtk_editable_get_text (GTK_EDITABLE
// (gtk_entry_completion_get_entry ()))`.
type EntryCompletionMatchFunc func(completion EntryCompletion, key string, iter *TreeIter, ok bool)

//export gotk4_EntryCompletionMatchFunc
func _EntryCompletionMatchFunc(arg0 *C.GtkEntryCompletion, arg1 *C.char, arg2 *C.GtkTreeIter, arg3 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var completion EntryCompletion // out
	var key string                 // out
	var iter *TreeIter             // out

	completion = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(EntryCompletion)
	key = C.GoString(arg1)
	iter = (*TreeIter)(unsafe.Pointer(arg2))

	fn := v.(EntryCompletionMatchFunc)
	ok := fn(completion, key, iter)

	var cret C.gboolean // out

	if ok {
		cret = C.TRUE
	}

	return cret
}

// EntryCompletion: `GtkEntryCompletion` is an auxiliary object to provide
// completion functionality for `GtkEntry`.
//
// It implements the [iface@Gtk.CellLayout] interface, to allow the user to add
// extra cells to the `GtkTreeView` with completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, `GtkEntryCompletion` checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see [method@Gtk.EntryCompletion.set_text_column]),
// but this can be overridden with a custom match function (see
// [method@Gtk.EntryCompletion.set_match_func]).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// [signal@Gtk.EntryCompletion::match-selected] signal and updating the entry in
// the signal handler. Note that you should return true from the signal handler
// to suppress the default behaviour.
//
// To add completion functionality to an entry, use
// [method@Gtk.Entry.set_completion].
//
// `GtkEntryCompletion` uses a [class@Gtk.TreeModelFilter] model to represent
// the subset of the entire model that is currently matching. While the
// `GtkEntryCompletion` signals [signal@Gtk.EntryCompletion::match-selected] and
// [signal@Gtk.EntryCompletion::cursor-on-match] take the original model and an
// iter pointing to that model as arguments, other callbacks and signals (such
// as `GtkCellLayoutDataFunc` or [signal@Gtk.CellArea::apply-attributes)] will
// generally take the filter model as argument. As long as you are only calling
// [method@Gtk.TreeModel.get], this will make no difference to you. If for some
// reason, you need the original model, use
// [method@Gtk.TreeModelFilter.get_model]. Don’t forget to use
// [method@Gtk.TreeModelFilter.convert_iter_to_child_iter] to obtain a matching
// iter.
type EntryCompletion interface {
	Buildable
	CellLayout

	CompleteEntryCompletion()

	ComputePrefixEntryCompletion(key string) string

	CompletionPrefix() string

	Entry() Widget

	InlineCompletion() bool

	InlineSelection() bool

	MinimumKeyLength() int

	Model() TreeModel

	PopupCompletion() bool

	PopupSetWidth() bool

	PopupSingleMatch() bool

	TextColumn() int

	InsertPrefixEntryCompletion()

	SetInlineCompletionEntryCompletion(inlineCompletion bool)

	SetInlineSelectionEntryCompletion(inlineSelection bool)

	SetMinimumKeyLengthEntryCompletion(length int)

	SetModelEntryCompletion(model TreeModel)

	SetPopupCompletionEntryCompletion(popupCompletion bool)

	SetPopupSetWidthEntryCompletion(popupSetWidth bool)

	SetPopupSingleMatchEntryCompletion(popupSingleMatch bool)

	SetTextColumnEntryCompletion(column int)
}

// entryCompletion implements the EntryCompletion class.
type entryCompletion struct {
	gextras.Objector
}

// WrapEntryCompletion wraps a GObject to the right type. It is
// primarily used internally.
func WrapEntryCompletion(obj *externglib.Object) EntryCompletion {
	return entryCompletion{
		Objector: obj,
	}
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntryCompletion(obj), nil
}

func NewEntryCompletion() EntryCompletion {
	var _cret *C.GtkEntryCompletion // in

	_cret = C.gtk_entry_completion_new()

	var _entryCompletion EntryCompletion // out

	_entryCompletion = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EntryCompletion)

	return _entryCompletion
}

func NewEntryCompletionWithArea(area CellArea) EntryCompletion {
	var _arg1 *C.GtkCellArea        // out
	var _cret *C.GtkEntryCompletion // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_entry_completion_new_with_area(_arg1)

	var _entryCompletion EntryCompletion // out

	_entryCompletion = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(EntryCompletion)

	return _entryCompletion
}

func (c entryCompletion) CompleteEntryCompletion() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_complete(_arg0)
}

func (c entryCompletion) ComputePrefixEntryCompletion(key string) string {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_entry_completion_compute_prefix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (c entryCompletion) CompletionPrefix() string {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_completion_prefix(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c entryCompletion) Entry() Widget {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_entry(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (c entryCompletion) InlineCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c entryCompletion) InlineSelection() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c entryCompletion) MinimumKeyLength() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_minimum_key_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c entryCompletion) Model() TreeModel {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

func (c entryCompletion) PopupCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c entryCompletion) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c entryCompletion) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c entryCompletion) TextColumn() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c entryCompletion) InsertPrefixEntryCompletion() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_insert_prefix(_arg0)
}

func (c entryCompletion) SetInlineCompletionEntryCompletion(inlineCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_completion(_arg0, _arg1)
}

func (c entryCompletion) SetInlineSelectionEntryCompletion(inlineSelection bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineSelection {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_selection(_arg0, _arg1)
}

func (c entryCompletion) SetMinimumKeyLengthEntryCompletion(length int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(length)

	C.gtk_entry_completion_set_minimum_key_length(_arg0, _arg1)
}

func (c entryCompletion) SetModelEntryCompletion(model TreeModel) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_entry_completion_set_model(_arg0, _arg1)
}

func (c entryCompletion) SetPopupCompletionEntryCompletion(popupCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_completion(_arg0, _arg1)
}

func (c entryCompletion) SetPopupSetWidthEntryCompletion(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
}

func (c entryCompletion) SetPopupSingleMatchEntryCompletion(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
}

func (c entryCompletion) SetTextColumnEntryCompletion(column int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(column)

	C.gtk_entry_completion_set_text_column(_arg0, _arg1)
}

func (b entryCompletion) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}

func (c entryCompletion) AddAttribute(cell CellRenderer, attribute string, column int) {
	WrapCellLayout(gextras.InternObject(c)).AddAttribute(cell, attribute, column)
}

func (c entryCompletion) Clear() {
	WrapCellLayout(gextras.InternObject(c)).Clear()
}

func (c entryCompletion) ClearAttributes(cell CellRenderer) {
	WrapCellLayout(gextras.InternObject(c)).ClearAttributes(cell)
}

func (c entryCompletion) Area() CellArea {
	return WrapCellLayout(gextras.InternObject(c)).Area()
}

func (c entryCompletion) PackEnd(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackEnd(cell, expand)
}

func (c entryCompletion) PackStart(cell CellRenderer, expand bool) {
	WrapCellLayout(gextras.InternObject(c)).PackStart(cell, expand)
}

func (c entryCompletion) Reorder(cell CellRenderer, position int) {
	WrapCellLayout(gextras.InternObject(c)).Reorder(cell, position)
}
