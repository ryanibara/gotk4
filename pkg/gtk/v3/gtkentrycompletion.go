// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gtk3_EntryCompletionClass_cursor_on_match(void*, void*, void*);
// extern gboolean _gotk4_gtk3_EntryCompletionClass_insert_prefix(void*, void*);
// extern gboolean _gotk4_gtk3_EntryCompletionClass_match_selected(void*, void*, void*);
// extern gboolean _gotk4_gtk3_EntryCompletionMatchFunc(void*, void*, void*, gpointer);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch(gpointer, void*, void*, guintptr);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectInsertPrefix(gpointer, void*, guintptr);
// extern gboolean _gotk4_gtk3_EntryCompletion_ConnectMatchSelected(gpointer, void*, void*, guintptr);
// extern void _gotk4_gtk3_EntryCompletionClass_action_activated(void*, gint);
// extern void _gotk4_gtk3_EntryCompletionClass_no_matches(void*);
// extern void _gotk4_gtk3_EntryCompletion_ConnectActionActivated(gpointer, gint, guintptr);
// extern void _gotk4_gtk3_EntryCompletion_ConnectNoMatches(gpointer, guintptr);
// extern void callbackDelete(gpointer);
import "C"

// glib.Type values for gtkentrycompletion.go.
var GTypeEntryCompletion = coreglib.Type(C.gtk_entry_completion_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeEntryCompletion, F: marshalEntryCompletion},
	})
}

// EntryCompletionMatchFunc: function which decides whether the row indicated by
// iter matches a given key, and should be displayed as a possible completion
// for key. Note that key is normalized and case-folded (see g_utf8_normalize()
// and g_utf8_casefold()). If this is not appropriate, match functions have
// access to the unmodified key via gtk_entry_get_text (GTK_ENTRY
// (gtk_entry_completion_get_entry ())).
type EntryCompletionMatchFunc func(completion *EntryCompletion, key string, iter *TreeIter) (ok bool)

//export _gotk4_gtk3_EntryCompletionMatchFunc
func _gotk4_gtk3_EntryCompletionMatchFunc(arg1 *C.void, arg2 *C.void, arg3 *C.void, arg4 C.gpointer) (cret C.gboolean) {
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

	_completion = wrapEntryCompletion(coreglib.Take(unsafe.Pointer(arg1)))
	_key = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_iter = (*TreeIter)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	ok := fn(_completion, _key, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// EntryCompletionOverrider contains methods that are overridable.
type EntryCompletionOverrider interface {
	// The function takes the following parameters:
	//
	ActionActivated(index_ int32)
	// The function takes the following parameters:
	//
	//    - model
	//    - iter
	//
	// The function returns the following values:
	//
	CursorOnMatch(model TreeModeller, iter *TreeIter) bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	InsertPrefix(prefix string) bool
	// The function takes the following parameters:
	//
	//    - model
	//    - iter
	//
	// The function returns the following values:
	//
	MatchSelected(model TreeModeller, iter *TreeIter) bool
	NoMatches()
}

// EntryCompletion is an auxiliary object to be used in conjunction with Entry
// to provide the completion functionality. It implements the CellLayout
// interface, to allow the user to add extra cells to the TreeView with
// completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, EntryCompletion checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see gtk_entry_completion_set_text_column()), but
// this can be overridden with a custom match function (see
// gtk_entry_completion_set_match_func()).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// EntryCompletion::match-selected signal and updating the entry in the signal
// handler. Note that you should return TRUE from the signal handler to suppress
// the default behaviour.
//
// To add completion functionality to an entry, use gtk_entry_set_completion().
//
// In addition to regular completion matches, which will be inserted into the
// entry when they are selected, EntryCompletion also allows to display
// “actions” in the popup window. Their appearance is similar to menuitems, to
// differentiate them clearly from completion strings. When an action is
// selected, the EntryCompletion::action-activated signal is emitted.
//
// GtkEntryCompletion uses a TreeModelFilter model to represent the subset of
// the entire model that is currently matching. While the GtkEntryCompletion
// signals EntryCompletion::match-selected and EntryCompletion::cursor-on-match
// take the original model and an iter pointing to that model as arguments,
// other callbacks and signals (such as CellLayoutDataFuncs or
// CellArea::apply-attributes) will generally take the filter model as argument.
// As long as you are only calling gtk_tree_model_get(), this will make no
// difference to you. If for some reason, you need the original model, use
// gtk_tree_model_filter_get_model(). Don’t forget to use
// gtk_tree_model_filter_convert_iter_to_child_iter() to obtain a matching iter.
type EntryCompletion struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
	CellLayout
}

var (
	_ coreglib.Objector = (*EntryCompletion)(nil)
)

func classInitEntryCompletioner(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkEntryCompletionClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkEntryCompletionClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ ActionActivated(index_ int32) }); ok {
		pclass.action_activated = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_action_activated)
	}

	if _, ok := goval.(interface {
		CursorOnMatch(model TreeModeller, iter *TreeIter) bool
	}); ok {
		pclass.cursor_on_match = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_cursor_on_match)
	}

	if _, ok := goval.(interface{ InsertPrefix(prefix string) bool }); ok {
		pclass.insert_prefix = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_insert_prefix)
	}

	if _, ok := goval.(interface {
		MatchSelected(model TreeModeller, iter *TreeIter) bool
	}); ok {
		pclass.match_selected = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_match_selected)
	}

	if _, ok := goval.(interface{ NoMatches() }); ok {
		pclass.no_matches = (*[0]byte)(C._gotk4_gtk3_EntryCompletionClass_no_matches)
	}
}

//export _gotk4_gtk3_EntryCompletionClass_action_activated
func _gotk4_gtk3_EntryCompletionClass_action_activated(arg0 *C.void, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ ActionActivated(index_ int32) })

	var _index_ int32 // out

	_index_ = int32(arg1)

	iface.ActionActivated(_index_)
}

//export _gotk4_gtk3_EntryCompletionClass_cursor_on_match
func _gotk4_gtk3_EntryCompletionClass_cursor_on_match(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		CursorOnMatch(model TreeModeller, iter *TreeIter) bool
	})

	var _model TreeModeller // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

	ok := iface.CursorOnMatch(_model, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_EntryCompletionClass_insert_prefix
func _gotk4_gtk3_EntryCompletionClass_insert_prefix(arg0 *C.void, arg1 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ InsertPrefix(prefix string) bool })

	var _prefix string // out

	_prefix = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := iface.InsertPrefix(_prefix)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_EntryCompletionClass_match_selected
func _gotk4_gtk3_EntryCompletionClass_match_selected(arg0 *C.void, arg1 *C.void, arg2 *C.void) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface {
		MatchSelected(model TreeModeller, iter *TreeIter) bool
	})

	var _model TreeModeller // out
	var _iter *TreeIter     // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gtk.TreeModeller is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

	ok := iface.MatchSelected(_model, _iter)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_EntryCompletionClass_no_matches
func _gotk4_gtk3_EntryCompletionClass_no_matches(arg0 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ NoMatches() })

	iface.NoMatches()
}

func wrapEntryCompletion(obj *coreglib.Object) *EntryCompletion {
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
	return wrapEntryCompletion(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_gtk3_EntryCompletion_ConnectActionActivated
func _gotk4_gtk3_EntryCompletion_ConnectActionActivated(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(index int32)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(index int32))
	}

	var _index int32 // out

	_index = int32(arg1)

	f(_index)
}

// ConnectActionActivated gets emitted when an action is activated.
func (completion *EntryCompletion) ConnectActionActivated(f func(index int32)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "action-activated", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectActionActivated), f)
}

//export _gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch
func _gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(model TreeModeller, iter *TreeIter) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
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

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

// ConnectCursorOnMatch gets emitted when a match from the cursor is on a match
// of the list. The default behaviour is to replace the contents of the entry
// with the contents of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk_entry_completion_set_model().
func (completion *EntryCompletion) ConnectCursorOnMatch(f func(model TreeModeller, iter *TreeIter) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "cursor-on-match", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectCursorOnMatch), f)
}

//export _gotk4_gtk3_EntryCompletion_ConnectInsertPrefix
func _gotk4_gtk3_EntryCompletion_ConnectInsertPrefix(arg0 C.gpointer, arg1 *C.void, arg2 C.guintptr) (cret C.gboolean) {
	var f func(prefix string) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
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

// ConnectInsertPrefix gets emitted when the inline autocompletion is triggered.
// The default behaviour is to make the entry display the whole prefix and
// select the newly inserted part.
//
// Applications may connect to this signal in order to insert only a smaller
// part of the prefix into the entry - e.g. the entry used in the FileChooser
// inserts only the part of the prefix up to the next '/'.
func (completion *EntryCompletion) ConnectInsertPrefix(f func(prefix string) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "insert-prefix", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectInsertPrefix), f)
}

//export _gotk4_gtk3_EntryCompletion_ConnectMatchSelected
func _gotk4_gtk3_EntryCompletion_ConnectMatchSelected(arg0 C.gpointer, arg1 *C.void, arg2 *C.void, arg3 C.guintptr) (cret C.gboolean) {
	var f func(model TreeModeller, iter *TreeIter) (ok bool)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
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

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
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

// ConnectMatchSelected gets emitted when a match from the list is selected. The
// default behaviour is to replace the contents of the entry with the contents
// of the text column in the row pointed to by iter.
//
// Note that model is the model that was passed to
// gtk_entry_completion_set_model().
func (completion *EntryCompletion) ConnectMatchSelected(f func(model TreeModeller, iter *TreeIter) (ok bool)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "match-selected", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectMatchSelected), f)
}

//export _gotk4_gtk3_EntryCompletion_ConnectNoMatches
func _gotk4_gtk3_EntryCompletion_ConnectNoMatches(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

// ConnectNoMatches gets emitted when the filter model has zero number of rows
// in completion_complete method. (In other words when GtkEntryCompletion is out
// of suggestions).
func (completion *EntryCompletion) ConnectNoMatches(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(completion, "no-matches", false, unsafe.Pointer(C._gotk4_gtk3_EntryCompletion_ConnectNoMatches), f)
}

// NewEntryCompletion creates a new EntryCompletion object.
//
// The function returns the following values:
//
//    - entryCompletion: newly created EntryCompletion object.
//
func NewEntryCompletion() *EntryCompletion {
	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("new_EntryCompletion", nil, nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// NewEntryCompletionWithArea creates a new EntryCompletion object using the
// specified area to layout cells in the underlying TreeViewColumn for the
// drop-down menu.
//
// The function takes the following parameters:
//
//    - area used to layout cells.
//
// The function returns the following values:
//
//    - entryCompletion: newly created EntryCompletion object.
//
func NewEntryCompletionWithArea(area CellAreaer) *EntryCompletion {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(area).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("new_EntryCompletion_with_area", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(area)

	var _entryCompletion *EntryCompletion // out

	_entryCompletion = wrapEntryCompletion(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _entryCompletion
}

// Complete requests a completion operation, or in other words a refiltering of
// the current list with completions, using the current key. The completion list
// view will be updated accordingly.
func (completion *EntryCompletion) Complete() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("complete", _args[:], nil)

	runtime.KeepAlive(completion)
}

// ComputePrefix computes the common prefix that is shared by all rows in
// completion that start with key. If no row matches key, NULL will be returned.
// Note that a text column must have been set for this function to work, see
// gtk_entry_completion_set_text_column() for details.
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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("compute_prefix", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)
	runtime.KeepAlive(key)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}

// DeleteAction deletes the action at index_ from completion’s action list.
//
// Note that index_ is a relative position and the position of an action may
// have changed since it was inserted.
//
// The function takes the following parameters:
//
//    - index_: index of the item to delete.
//
func (completion *EntryCompletion) DeleteAction(index_ int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(index_)

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("delete_action", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(index_)
}

// CompletionPrefix: get the original text entered by the user that triggered
// the completion or NULL if there’s no completion ongoing.
//
// The function returns the following values:
//
//    - utf8: prefix for the current completion.
//
func (completion *EntryCompletion) CompletionPrefix() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_completion_prefix", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Entry gets the entry completion has been attached to.
//
// The function returns the following values:
//
//    - widget: entry completion has been attached to.
//
func (completion *EntryCompletion) Entry() Widgetter {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_entry", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _widget Widgetter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_inline_completion", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_inline_selection", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
func (completion *EntryCompletion) MinimumKeyLength() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_minimum_key_length", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// Model returns the model the EntryCompletion is using as data source. Returns
// NULL if the model is unset.
//
// The function returns the following values:
//
//    - treeModel (optional) or NULL if none is currently being used.
//
func (completion *EntryCompletion) Model() *TreeModel {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_model", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _treeModel *TreeModel // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_treeModel = wrapTreeModel(coreglib.Take(unsafe.Pointer(_cret)))
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_popup_completion", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_popup_set_width", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_popup_single_match", _args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
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
func (completion *EntryCompletion) TextColumn() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	_gret := girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("get_text_column", _args[:], nil)
	_cret = *(*C.gint)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(completion)

	var _gint int32 // out

	_gint = int32(*(*C.gint)(unsafe.Pointer(&_cret)))

	return _gint
}

// InsertActionMarkup inserts an action in completion’s action item list at
// position index_ with markup markup.
//
// The function takes the following parameters:
//
//    - index_: index of the item to insert.
//    - markup of the item to insert.
//
func (completion *EntryCompletion) InsertActionMarkup(index_ int32, markup string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(index_)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(markup)))
	defer C.free(unsafe.Pointer(_args[2]))

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("insert_action_markup", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(markup)
}

// InsertActionText inserts an action in completion’s action item list at
// position index_ with text text. If you want the action item to have markup,
// use gtk_entry_completion_insert_action_markup().
//
// Note that index_ is a relative position in the list of actions and the
// position of an action can change when deleting a different action.
//
// The function takes the following parameters:
//
//    - index_: index of the item to insert.
//    - text of the item to insert.
//
func (completion *EntryCompletion) InsertActionText(index_ int32, text string) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(index_)
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[2]))

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("insert_action_text", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(index_)
	runtime.KeepAlive(text)
}

// InsertPrefix requests a prefix insertion.
func (completion *EntryCompletion) InsertPrefix() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("insert_prefix", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if inlineCompletion {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_inline_completion", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if inlineSelection {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_inline_selection", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(inlineSelection)
}

// SetMatchFunc sets the match function for completion to be func. The match
// function is used to determine if a row should or should not be in the
// completion list.
//
// The function takes the following parameters:
//
//    - fn to use.
//
func (completion *EntryCompletion) SetMatchFunc(fn EntryCompletionMatchFunc) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gpointer)(unsafe.Pointer(&_args[1])) = (*[0]byte)(C._gotk4_gtk3_EntryCompletionMatchFunc)
	_args[2] = C.gpointer(gbox.Assign(fn))
	_args[3] = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_match_func", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(fn)
}

// SetMinimumKeyLength requires the length of the search key for completion to
// be at least length. This is useful for long lists, where completing using a
// small key takes a lot of time and will come up with meaningless results
// anyway (ie, a too large dataset).
//
// The function takes the following parameters:
//
//    - length: minimum length of the key in order to start completing.
//
func (completion *EntryCompletion) SetMinimumKeyLength(length int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(length)

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_minimum_key_length", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(length)
}

// SetModel sets the model for a EntryCompletion. If completion already has a
// model set, it will remove it before setting the new model. If model is NULL,
// then it will unset the model.
//
// The function takes the following parameters:
//
//    - model (optional): TreeModel.
//
func (completion *EntryCompletion) SetModel(model TreeModeller) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if model != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_model", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if popupCompletion {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_popup_completion", _args[:], nil)

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
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if popupSetWidth {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_popup_set_width", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSetWidth)
}

// SetPopupSingleMatch sets whether the completion popup window will appear even
// if there is only a single match. You may want to set this to FALSE if you are
// using [inline completion][GtkEntryCompletion--inline-completion].
//
// The function takes the following parameters:
//
//    - popupSingleMatch: TRUE if the popup should appear even for a single
//      match.
//
func (completion *EntryCompletion) SetPopupSingleMatch(popupSingleMatch bool) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	if popupSingleMatch {
		*(*C.gboolean)(unsafe.Pointer(&_args[1])) = C.TRUE
	}

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_popup_single_match", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(popupSingleMatch)
}

// SetTextColumn: convenience function for setting up the most used case of this
// code: a completion list with just strings. This function will set up
// completion to have a list displaying all (and just) strings in the completion
// list, and to get those strings from column in the model of completion.
//
// This functions creates and adds a CellRendererText for the selected column.
// If you need to set the text column, but don't want the cell renderer, use
// g_object_set() to set the EntryCompletion:text-column property directly.
//
// The function takes the following parameters:
//
//    - column in the model of completion to get strings from.
//
func (completion *EntryCompletion) SetTextColumn(column int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(completion).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(column)

	girepository.MustFind("Gtk", "EntryCompletion").InvokeMethod("set_text_column", _args[:], nil)

	runtime.KeepAlive(completion)
	runtime.KeepAlive(column)
}
