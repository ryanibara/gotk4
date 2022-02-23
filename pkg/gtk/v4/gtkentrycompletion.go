// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
// extern gboolean _gotk4_gtk4_EntryCompletionMatchFunc(GtkEntryCompletion*, char*, GtkTreeIter*, gpointer);
// extern gboolean _gotk4_gtk4_EntryCompletion_ConnectCursorOnMatch(gpointer, GtkTreeModel*, GtkTreeIter*, guintptr);
// extern gboolean _gotk4_gtk4_EntryCompletion_ConnectInsertPrefix(gpointer, gchar*, guintptr);
// extern gboolean _gotk4_gtk4_EntryCompletion_ConnectMatchSelected(gpointer, GtkTreeModel*, GtkTreeIter*, guintptr);
// extern void _gotk4_gtk4_EntryCompletion_ConnectNoMatches(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkentrycompletion.go.
var GTypeEntryCompletion = externglib.Type(C.gtk_entry_completion_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeEntryCompletion, F: marshalEntryCompletion},
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
func _gotk4_gtk4_EntryCompletionMatchFunc(arg1 *C.GtkEntryCompletion, arg2 *C.char, arg3 *C.GtkTreeIter, arg4 C.gpointer) (cret C.gboolean) {
	var fn EntryCompletionMatchFunc
	{
		v := gbox.Get(uintptr(arg4))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(EntryCompletionMatchFunc)
	}

	var _completion *EntryCompletion // out
	var _key string                  // out
	var _iter *TreeIter              // out

	_completion = wrapEntryCompletion(externglib.Take(unsafe.Pointer(arg1)))
	_key = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	ok := fn(_completion, _key, _iter)

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
	_ [0]func() // equal guard
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

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	return wrapEntryCompletion(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk4_EntryCompletion_ConnectCursorOnMatch
func _gotk4_gtk4_EntryCompletion_ConnectCursorOnMatch(arg0 C.gpointer, arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 C.guintptr) (cret C.gboolean) {
	var f func(model TreeModeller, iter *TreeIter) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(model TreeModeller, iter *TreeIter) (ok bool))
	}

	var _model TreeModeller // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := f(_model, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectCursorOnMatch: emitted when a match from the cursor is on a match of
// the list.
//
// The default behaviour is to replace the contents of the entry with the
// contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk.EntryCompletion.SetModel().
func (completion *EntryCompletion) ConnectCursorOnMatch(f func(model TreeModeller, iter *TreeIter) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(completion, "cursor-on-match", false, unsafe.Pointer(C._gotk4_gtk4_EntryCompletion_ConnectCursorOnMatch), f)
}

//export _gotk4_gtk4_EntryCompletion_ConnectInsertPrefix
func _gotk4_gtk4_EntryCompletion_ConnectInsertPrefix(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) (cret C.gboolean) {
	var f func(prefix string) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(prefix string) (ok bool))
	}

	var _prefix string // out

	_prefix = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := f(_prefix)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectInsertPrefix: emitted when the inline autocompletion is triggered.
//
// The default behaviour is to make the entry display the whole prefix and
// select the newly inserted part.
//
// Applications may connect to this signal in order to insert only a smaller
// part of the prefix into the entry - e.g. the entry used in the FileChooser
// inserts only the part of the prefix up to the next '/'.
func (completion *EntryCompletion) ConnectInsertPrefix(f func(prefix string) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(completion, "insert-prefix", false, unsafe.Pointer(C._gotk4_gtk4_EntryCompletion_ConnectInsertPrefix), f)
}

//export _gotk4_gtk4_EntryCompletion_ConnectMatchSelected
func _gotk4_gtk4_EntryCompletion_ConnectMatchSelected(arg0 C.gpointer, arg1 *C.GtkTreeModel, arg2 *C.GtkTreeIter, arg3 C.guintptr) (cret C.gboolean) {
	var f func(model TreeModeller, iter *TreeIter) (ok bool)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(model TreeModeller, iter *TreeIter) (ok bool))
	}

	var _model TreeModeller // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TreeModeller)
			return ok
		})
		rv, ok := casted.(TreeModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
		}
		_model = rv
	}
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := f(_model, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// ConnectMatchSelected: emitted when a match from the list is selected.
//
// The default behaviour is to replace the contents of the entry with the
// contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk.EntryCompletion.SetModel().
func (completion *EntryCompletion) ConnectMatchSelected(f func(model TreeModeller, iter *TreeIter) (ok bool)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(completion, "match-selected", false, unsafe.Pointer(C._gotk4_gtk4_EntryCompletion_ConnectMatchSelected), f)
}

//export _gotk4_gtk4_EntryCompletion_ConnectNoMatches
func _gotk4_gtk4_EntryCompletion_ConnectNoMatches(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectNoMatches: emitted when the filter model has zero number of rows in
// completion_complete method.
//
// In other words when GtkEntryCompletion is out of suggestions.
func (completion *EntryCompletion) ConnectNoMatches(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(completion, "no-matches", false, unsafe.Pointer(C._gotk4_gtk4_EntryCompletion_ConnectNoMatches), f)
}

// NewEntryCompletion creates a new GtkEntryCompletion object.
//
// The function returns the following values:
//
//    - entryCompletion: newly created GtkEntryCompletion object.
//
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
// The function returns the following values:
//
//    - entryCompletion: newly created GtkEntryCompletion object.
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
// The function returns the following values:
//
//    - utf8 (optional): common prefix all rows starting with key or NULL if no
//      row matches key.
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
//
// The function returns the following values:
//
//    - utf8 (optional): prefix for the current completion.
//
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
//
// The function returns the following values:
//
//    - widget: entry completion has been attached to.
//
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
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		_widget = rv
	}

	return _widget
}

// InlineCompletion returns whether the common prefix of the possible
// completions should be automatically inserted in the entry.
//
// The function returns the following values:
//
//    - ok: TRUE if inline completion is turned on.
//
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
//
// The function returns the following values:
//
//    - ok: TRUE if inline-selection mode is on.
//
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
//
// The function returns the following values:
//
//    - gint: currently used minimum key length.
//
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
//
// The function returns the following values:
//
//    - treeModel (optional) or NULL if none is currently being used.
//
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
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(TreeModeller)
				return ok
			})
			rv, ok := casted.(TreeModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.TreeModeller")
			}
			_treeModel = rv
		}
	}

	return _treeModel
}

// PopupCompletion returns whether the completions should be presented in a
// popup window.
//
// The function returns the following values:
//
//    - ok: TRUE if popup completion is turned on.
//
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
//
// The function returns the following values:
//
//    - ok: TRUE if the popup window will be resized to the width of the entry.
//
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
//
// The function returns the following values:
//
//    - ok: TRUE if the popup window will appear regardless of the number of
//      matches.
//
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
//
// The function returns the following values:
//
//    - gint: column containing the strings.
//
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
//    - model (optional): GtkTreeModel.
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
//    - popupSetWidth: TRUE to make the width of the popup the same as the entry.
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
//      match.
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
