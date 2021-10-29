// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern void callbackDelete(gpointer);
// gboolean _gotk4_gtk4_EntryCompletionMatchFunc(GtkEntryCompletion*, char*, GtkTreeIter*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_completion_get_type()), F: marshalEntryCompletioner},
	})
}

// EntryCompletionMatchFunc: function which decides whether the row indicated by
// iter matches a given key, and should be displayed as a possible completion
// for key.
//
// Note that key is normalized and case-folded (see g_utf8_normalize() and
// g_utf8_casefold()). If this is not appropriate, match functions have access
// to the unmodified key via gtk_editable_get_text (GTK_EDITABLE
// (gtk_entry_completion_get_entry ())).
type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter) (ok bool)

//export _gotk4_gtk4_EntryCompletionMatchFunc
func _gotk4_gtk4_EntryCompletionMatchFunc(arg0 *C.GtkEntryCompletion, arg1 *C.char, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gboolean) {
	v := gbox.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var completion *EntryCompletion // out
	var key string                  // out
	var iter *TreeIter              // out

	completion = wrapEntryCompletion(externglib.Take(unsafe.Pointer(arg0)))
	key = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	fn := v.(EntryCompletionMatchFunc)
	ok := fn(completion, key, iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// EntryCompletion: GtkEntryCompletion is an auxiliary object to provide
// completion functionality for GtkEntry.
//
// It implements the gtk.CellLayout interface, to allow the user to add extra
// cells to the GtkTreeView with completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, GtkEntryCompletion checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see gtk.EntryCompletion.SetTextColumn()), but this
// can be overridden with a custom match function (see
// gtk.EntryCompletion.SetMatchFunc()).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// gtk.EntryCompletion::match-selected signal and updating the entry in the
// signal handler. Note that you should return TRUE from the signal handler to
// suppress the default behaviour.
//
// To add completion functionality to an entry, use gtk.Entry.SetCompletion().
//
// GtkEntryCompletion uses a gtk.TreeModelFilter model to represent the subset
// of the entire model that is currently matching. While the GtkEntryCompletion
// signals gtk.EntryCompletion::match-selected and
// gtk.EntryCompletion::cursor-on-match take the original model and an iter
// pointing to that model as arguments, other callbacks and signals (such as
// GtkCellLayoutDataFunc or gtk.CellArea::apply-attributes) will generally take
// the filter model as argument. As long as you are only calling
// gtk.TreeModel.Get(), this will make no difference to you. If for some reason,
// you need the original model, use gtk.TreeModelFilter.GetModel(). Don’t forget
// to use gtk.TreeModelFilter.ConvertIterToChildIter() to obtain a matching
// iter.
type EntryCompletion struct {
	*externglib.Object

	Buildable
	CellLayout
}

var (
	_ externglib.Objector = (*EntryCompletion)(nil)
)

func wrapEntryCompletion(obj *externglib.Object) *EntryCompletion {
	return &EntryCompletion{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
		CellLayout: CellLayout{
			Object: obj,
		},
	}
}

func marshalEntryCompletioner(p uintptr) (interface{}, error) {
	return wrapEntryCompletion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewEntryCompletion creates a new GtkEntryCompletion object.
func NewEntryCompletion() *EntryCompletion {
	var _cret *C.GtkEntryCompletion // in

	_cret = C.gtk_entry_completion_new()

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// NewEntryCompletionWithArea creates a new GtkEntryCompletion object using the
// specified area.
//
// The GtkCellArea is used to layout cells in the underlying GtkTreeViewColumn
// for the drop-down menu.
//
// The function takes the following parameters:
//
//    - area used to layout cells.
//
func NewEntryCompletionWithArea(area CellAreaer) *EntryCompletion {
	var _arg1 *C.GtkCellArea        // out
	var _cret *C.GtkEntryCompletion // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_entry_completion_new_with_area(_arg1)
	runtime.KeepAlive(area)

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// Complete requests a completion operation, or in other words a refiltering of
// the current list with completions, using the current key.
//
// The completion list view will be updated accordingly.
func (completion *EntryCompletion) Complete() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_completion_complete(_arg0)
	runtime.KeepAlive(completion)
}

// ComputePrefix computes the common prefix that is shared by all rows in
// completion that start with key.
//
// If no row matches key, NULL will be returned. Note that a text column must
// have been set for this function to work, see
// gtk.EntryCompletion.SetTextColumn() for details.
//
// The function takes the following parameters:
//
//    - key: text to complete for.
//
func (completion *EntryCompletion) ComputePrefix(key string) string {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_entry_completion_compute_prefix(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(key)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// CompletionPrefix: get the original text entered by the user that triggered
// the completion or NULL if there’s no completion ongoing.
func (completion *EntryCompletion) CompletionPrefix() string {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_completion_prefix(_arg0)
	runtime.KeepAlive(completion)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Entry gets the entry completion has been attached to.
func (completion *EntryCompletion) Entry() Widgetter {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_entry(_arg0)
	runtime.KeepAlive(completion)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		rv, ok := (externglib.CastObject(object)).(Widgetter)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// InlineCompletion returns whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (completion *EntryCompletion) InlineCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_inline_completion(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InlineSelection returns TRUE if inline-selection mode is turned on.
func (completion *EntryCompletion) InlineSelection() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_inline_selection(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MinimumKeyLength returns the minimum key length as set for completion.
func (completion *EntryCompletion) MinimumKeyLength() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_minimum_key_length(_arg0)
	runtime.KeepAlive(completion)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the model the GtkEntryCompletion is using as data source.
//
// Returns NULL if the model is unset.
func (completion *EntryCompletion) Model() TreeModeller {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_model(_arg0)
	runtime.KeepAlive(completion)

	var _treeModel TreeModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(TreeModeller)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.TreeModeller")
			}
			_treeModel = rv
		}
	}

	return _treeModel
}

// PopupCompletion returns whether the completions should be presented in a
// popup window.
func (completion *EntryCompletion) PopupCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_completion(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSetWidth returns whether the completion popup window will be resized to
// the width of the entry.
func (completion *EntryCompletion) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSingleMatch returns whether the completion popup window will appear even
// if there is only a single match.
func (completion *EntryCompletion) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)
	runtime.KeepAlive(completion)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextColumn returns the column in the model of completion to get strings from.
func (completion *EntryCompletion) TextColumn() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	_cret = C.gtk_entry_completion_get_text_column(_arg0)
	runtime.KeepAlive(completion)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertPrefix requests a prefix insertion.
func (completion *EntryCompletion) InsertPrefix() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_completion_insert_prefix(_arg0)
	runtime.KeepAlive(completion)
}

// SetInlineCompletion sets whether the common prefix of the possible
// completions should be automatically inserted in the entry.
//
// The function takes the following parameters:
//
//    - inlineCompletion: TRUE to do inline completion.
//
func (completion *EntryCompletion) SetInlineCompletion(inlineCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if inlineCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_completion(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(inlineCompletion)
}

// SetInlineSelection sets whether it is possible to cycle through the possible
// completions inside the entry.
//
// The function takes the following parameters:
//
//    - inlineSelection: TRUE to do inline selection.
//
func (completion *EntryCompletion) SetInlineSelection(inlineSelection bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if inlineSelection {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_selection(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(inlineSelection)
}

// SetMatchFunc sets the match function for completion to be func.
//
// The match function is used to determine if a row should or should not be in
// the completion list.
//
// The function takes the following parameters:
//
//    - fn: GtkEntryCompletionMatchFunc to use.
//
func (completion *EntryCompletion) SetMatchFunc(fn EntryCompletionMatchFunc) {
	var _arg0 *C.GtkEntryCompletion         // out
	var _arg1 C.GtkEntryCompletionMatchFunc // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_EntryCompletionMatchFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gtk_entry_completion_set_match_func(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(fn)
}

// SetMinimumKeyLength requires the length of the search key for completion to
// be at least length.
//
// This is useful for long lists, where completing using a small key takes a lot
// of time and will come up with meaningless results anyway (ie, a too large
// dataset).
//
// The function takes the following parameters:
//
//    - length: minimum length of the key in order to start completing.
//
func (completion *EntryCompletion) SetMinimumKeyLength(length int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.int(length)

	C.gtk_entry_completion_set_minimum_key_length(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(length)
}

// SetModel sets the model for a GtkEntryCompletion.
//
// If completion already has a model set, it will remove it before setting the
// new model. If model is NULL, then it will unset the model.
//
// The function takes the following parameters:
//
//    - model: GtkTreeModel.
//
func (completion *EntryCompletion) SetModel(model TreeModeller) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if model != nil {
		_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_entry_completion_set_model(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(model)
}

// SetPopupCompletion sets whether the completions should be presented in a
// popup window.
//
// The function takes the following parameters:
//
//    - popupCompletion: TRUE to do popup completion.
//
func (completion *EntryCompletion) SetPopupCompletion(popupCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_completion(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupCompletion)
}

// SetPopupSetWidth sets whether the completion popup window will be resized to
// be the same width as the entry.
//
// The function takes the following parameters:
//
//    - popupSetWidth: TRUE to make the width of the popup the same as the
//    entry.
//
func (completion *EntryCompletion) SetPopupSetWidth(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSetWidth)
}

// SetPopupSingleMatch sets whether the completion popup window will appear even
// if there is only a single match.
//
// You may want to set this to FALSE if you are using
// gtk.EntryCompletion:inline-completion.
//
// The function takes the following parameters:
//
//    - popupSingleMatch: TRUE if the popup should appear even for a single
//    match.
//
func (completion *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSingleMatch)
}

// SetTextColumn: convenience function for setting up the most used case of this
// code: a completion list with just strings.
//
// This function will set up completion to have a list displaying all (and just)
// strings in the completion list, and to get those strings from column in the
// model of completion.
//
// This functions creates and adds a GtkCellRendererText for the selected
// column. If you need to set the text column, but don't want the cell renderer,
// use g_object_set() to set the gtk.EntryCompletion:text-column property
// directly.
//
// The function takes the following parameters:
//
//    - column in the model of completion to get strings from.
//
func (completion *EntryCompletion) SetTextColumn(column int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))
	_arg1 = C.int(column)

	C.gtk_entry_completion_set_text_column(_arg0, _arg1)
	runtime.KeepAlive(completion)
	runtime.KeepAlive(column)
}

// ConnectCursorOnMatch: emitted when a match from the cursor is on a match of
// the list.
//
// The default behaviour is to replace the contents of the entry with the
// contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk.EntryCompletion.SetModel().
func (completion *EntryCompletion) ConnectCursorOnMatch(f func(model TreeModeller, iter *TreeIter) bool) externglib.SignalHandle {
	return completion.Connect("cursor-on-match", f)
}

// ConnectInsertPrefix: emitted when the inline autocompletion is triggered.
//
// The default behaviour is to make the entry display the whole prefix and
// select the newly inserted part.
//
// Applications may connect to this signal in order to insert only a smaller
// part of the prefix into the entry - e.g. the entry used in the FileChooser
// inserts only the part of the prefix up to the next '/'.
func (completion *EntryCompletion) ConnectInsertPrefix(f func(prefix string) bool) externglib.SignalHandle {
	return completion.Connect("insert-prefix", f)
}

// ConnectMatchSelected: emitted when a match from the list is selected.
//
// The default behaviour is to replace the contents of the entry with the
// contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk.EntryCompletion.SetModel().
func (completion *EntryCompletion) ConnectMatchSelected(f func(model TreeModeller, iter *TreeIter) bool) externglib.SignalHandle {
	return completion.Connect("match-selected", f)
}

// ConnectNoMatches: emitted when the filter model has zero number of rows in
// completion_complete method.
//
// In other words when GtkEntryCompletion is out of suggestions.
func (completion *EntryCompletion) ConnectNoMatches(f func()) externglib.SignalHandle {
	return completion.Connect("no-matches", f)
}
