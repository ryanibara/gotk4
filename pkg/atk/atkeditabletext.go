// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
// extern void _gotk4_atk1_EditableTextIface_copy_text(void*, gint, gint);
// extern void _gotk4_atk1_EditableTextIface_cut_text(void*, gint, gint);
// extern void _gotk4_atk1_EditableTextIface_delete_text(void*, gint, gint);
// extern void _gotk4_atk1_EditableTextIface_insert_text(void*, void*, gint, void*);
// extern void _gotk4_atk1_EditableTextIface_paste_text(void*, gint);
// extern void _gotk4_atk1_EditableTextIface_set_text_contents(void*, void*);
import "C"

// GTypeEditableText returns the GType for the type EditableText.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeEditableText() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Atk", "EditableText").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalEditableText)
	return gtype
}

// EditableTextOverrider contains methods that are overridable.
type EditableTextOverrider interface {
	// CopyText: copy text from start_pos up to, but not including end_pos to
	// the clipboard.
	//
	// The function takes the following parameters:
	//
	//    - startPos: start position.
	//    - endPos: end position.
	//
	CopyText(startPos, endPos int32)
	// CutText: copy text from start_pos up to, but not including end_pos to the
	// clipboard and then delete from the widget.
	//
	// The function takes the following parameters:
	//
	//    - startPos: start position.
	//    - endPos: end position.
	//
	CutText(startPos, endPos int32)
	// DeleteText: delete text start_pos up to, but not including end_pos.
	//
	// The function takes the following parameters:
	//
	//    - startPos: start position.
	//    - endPos: end position.
	//
	DeleteText(startPos, endPos int32)
	// InsertText: insert text at a given position.
	//
	// The function takes the following parameters:
	//
	//    - str: text to insert.
	//    - length of text to insert, in bytes.
	//    - position: caller initializes this to the position at which to insert
	//      the text. After the call it points at the position after the newly
	//      inserted text.
	//
	InsertText(str string, length int32, position *int32)
	// PasteText: paste text from clipboard to specified position.
	//
	// The function takes the following parameters:
	//
	//    - position to paste.
	//
	PasteText(position int32)
	// SetTextContents: set text contents of text.
	//
	// The function takes the following parameters:
	//
	//    - str: string to set for text contents of text.
	//
	SetTextContents(str string)
}

// EditableText should be implemented by UI components which contain text which
// the user can edit, via the Object corresponding to that component (see
// Object).
//
// EditableText is a subclass of Text, and as such, an object which implements
// EditableText is by definition an Text implementor as well.
//
// See also: Text.
//
// EditableText wraps an interface. This means the user can get the
// underlying type by calling Cast().
type EditableText struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*EditableText)(nil)
)

// EditableTexter describes EditableText's interface methods.
type EditableTexter interface {
	coreglib.Objector

	// CopyText: copy text from start_pos up to, but not including end_pos to
	// the clipboard.
	CopyText(startPos, endPos int32)
	// CutText: copy text from start_pos up to, but not including end_pos to the
	// clipboard and then delete from the widget.
	CutText(startPos, endPos int32)
	// DeleteText: delete text start_pos up to, but not including end_pos.
	DeleteText(startPos, endPos int32)
	// InsertText: insert text at a given position.
	InsertText(str string, length int32, position *int32)
	// PasteText: paste text from clipboard to specified position.
	PasteText(position int32)
	// SetTextContents: set text contents of text.
	SetTextContents(str string)
}

var _ EditableTexter = (*EditableText)(nil)

func ifaceInitEditableTexter(gifacePtr, data C.gpointer) {
	iface := girepository.MustFind("Atk", "EditableTextIface")
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("copy_text"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_copy_text)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("cut_text"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_cut_text)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("delete_text"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_delete_text)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("insert_text"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_insert_text)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("paste_text"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_paste_text)
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(gifacePtr), iface.StructFieldOffset("set_text_contents"))) = unsafe.Pointer(C._gotk4_atk1_EditableTextIface_set_text_contents)
}

//export _gotk4_atk1_EditableTextIface_copy_text
func _gotk4_atk1_EditableTextIface_copy_text(arg0 *C.void, arg1 C.gint, arg2 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _startPos int32 // out
	var _endPos int32   // out

	_startPos = int32(arg1)
	_endPos = int32(arg2)

	iface.CopyText(_startPos, _endPos)
}

//export _gotk4_atk1_EditableTextIface_cut_text
func _gotk4_atk1_EditableTextIface_cut_text(arg0 *C.void, arg1 C.gint, arg2 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _startPos int32 // out
	var _endPos int32   // out

	_startPos = int32(arg1)
	_endPos = int32(arg2)

	iface.CutText(_startPos, _endPos)
}

//export _gotk4_atk1_EditableTextIface_delete_text
func _gotk4_atk1_EditableTextIface_delete_text(arg0 *C.void, arg1 C.gint, arg2 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _startPos int32 // out
	var _endPos int32   // out

	_startPos = int32(arg1)
	_endPos = int32(arg2)

	iface.DeleteText(_startPos, _endPos)
}

//export _gotk4_atk1_EditableTextIface_insert_text
func _gotk4_atk1_EditableTextIface_insert_text(arg0 *C.void, arg1 *C.void, arg2 C.gint, arg3 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _str string      // out
	var _length int32    // out
	var _position *int32 // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	_length = int32(arg2)
	_position = (*int32)(unsafe.Pointer(arg3))

	iface.InsertText(_str, _length, _position)
}

//export _gotk4_atk1_EditableTextIface_paste_text
func _gotk4_atk1_EditableTextIface_paste_text(arg0 *C.void, arg1 C.gint) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _position int32 // out

	_position = int32(arg1)

	iface.PasteText(_position)
}

//export _gotk4_atk1_EditableTextIface_set_text_contents
func _gotk4_atk1_EditableTextIface_set_text_contents(arg0 *C.void, arg1 *C.void) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(EditableTextOverrider)

	var _str string // out

	_str = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetTextContents(_str)
}

func wrapEditableText(obj *coreglib.Object) *EditableText {
	return &EditableText{
		Object: obj,
	}
}

func marshalEditableText(p uintptr) (interface{}, error) {
	return wrapEditableText(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CopyText: copy text from start_pos up to, but not including end_pos to the
// clipboard.
//
// The function takes the following parameters:
//
//    - startPos: start position.
//    - endPos: end position.
//
func (text *EditableText) CopyText(startPos, endPos int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("copy_text", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// CutText: copy text from start_pos up to, but not including end_pos to the
// clipboard and then delete from the widget.
//
// The function takes the following parameters:
//
//    - startPos: start position.
//    - endPos: end position.
//
func (text *EditableText) CutText(startPos, endPos int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("cut_text", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// DeleteText: delete text start_pos up to, but not including end_pos.
//
// The function takes the following parameters:
//
//    - startPos: start position.
//    - endPos: end position.
//
func (text *EditableText) DeleteText(startPos, endPos int32) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(startPos)
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(endPos)

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("delete_text", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// InsertText: insert text at a given position.
//
// The function takes the following parameters:
//
//    - str: text to insert.
//    - length of text to insert, in bytes.
//    - position: caller initializes this to the position at which to insert the
//      text. After the call it points at the position after the newly inserted
//      text.
//
func (text *EditableText) InsertText(str string, length int32, position *int32) {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.gint)(unsafe.Pointer(&_args[2])) = C.gint(length)
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(position))

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("insert_text", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(str)
	runtime.KeepAlive(length)
	runtime.KeepAlive(position)
}

// PasteText: paste text from clipboard to specified position.
//
// The function takes the following parameters:
//
//    - position to paste.
//
func (text *EditableText) PasteText(position int32) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(position)

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("paste_text", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(position)
}

// SetTextContents: set text contents of text.
//
// The function takes the following parameters:
//
//    - str: string to set for text contents of text.
//
func (text *EditableText) SetTextContents(str string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(text).Native()))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_args[1]))

	_info := girepository.MustFind("Atk", "EditableText")
	_info.InvokeIfaceMethod("set_text_contents", _args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(str)
}
