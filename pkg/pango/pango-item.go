// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeItem returns the GType for the type Item.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeItem() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "Item").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalItem)
	return gtype
}

// ANALYSIS_FLAG_CENTERED_BASELINE: whether the segment should be shifted to
// center around the baseline.
//
// This is mainly used in vertical writing directions.
const ANALYSIS_FLAG_CENTERED_BASELINE = 1

// ANALYSIS_FLAG_IS_ELLIPSIS: whether this run holds ellipsized text.
const ANALYSIS_FLAG_IS_ELLIPSIS = 2

// ANALYSIS_FLAG_NEED_HYPHEN: whether to add a hyphen at the end of the run
// during shaping.
const ANALYSIS_FLAG_NEED_HYPHEN = 4

// Analysis: PangoAnalysis structure stores information about the properties of
// a segment of text.
//
// An instance of this type is always passed by reference.
type Analysis struct {
	*analysis
}

// analysis is the struct that's finalized.
type analysis struct {
	native unsafe.Pointer
}

// ShapeEngine: unused.
func (a *Analysis) ShapeEngine() unsafe.Pointer {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("shape_engine")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v unsafe.Pointer // out
	v = (unsafe.Pointer)(unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&*valptr))))
	return v
}

// LangEngine: unused.
func (a *Analysis) LangEngine() unsafe.Pointer {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("lang_engine")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v unsafe.Pointer // out
	v = (unsafe.Pointer)(unsafe.Pointer(*(*C.gpointer)(unsafe.Pointer(&*valptr))))
	return v
}

// Font: font for this segment.
func (a *Analysis) Font() Fonter {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("font")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v Fonter // out
	{
		objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&*valptr)))
		if objptr == nil {
			panic("object of type pango.Fonter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Fonter)
			return ok
		})
		rv, ok := casted.(Fonter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching pango.Fonter")
		}
		v = rv
	}
	return v
}

// Level: bidirectional level for this segment.
func (a *Analysis) Level() byte {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("level")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v byte // out
	v = byte(*(*C.guint8)(unsafe.Pointer(&*valptr)))
	return v
}

// Gravity: glyph orientation for this segment (A PangoGravity).
func (a *Analysis) Gravity() byte {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("gravity")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v byte // out
	v = byte(*(*C.guint8)(unsafe.Pointer(&*valptr)))
	return v
}

// Flags: boolean flags for this segment (Since: 1.16).
func (a *Analysis) Flags() byte {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("flags")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v byte // out
	v = byte(*(*C.guint8)(unsafe.Pointer(&*valptr)))
	return v
}

// Script: detected script for this segment (A PangoScript) (Since: 1.18).
func (a *Analysis) Script() byte {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("script")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v byte // out
	v = byte(*(*C.guint8)(unsafe.Pointer(&*valptr)))
	return v
}

// Language: detected language for this segment.
func (a *Analysis) Language() *Language {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("language")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	var v *Language // out
	v = (*Language)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&*valptr)))))
	return v
}

// Level: bidirectional level for this segment.
func (a *Analysis) SetLevel(level byte) {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("level")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	*(*C.guint8)(unsafe.Pointer(&*valptr)) = C.guint8(level)
}

// Gravity: glyph orientation for this segment (A PangoGravity).
func (a *Analysis) SetGravity(gravity byte) {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("gravity")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	*(*C.guint8)(unsafe.Pointer(&*valptr)) = C.guint8(gravity)
}

// Flags: boolean flags for this segment (Since: 1.16).
func (a *Analysis) SetFlags(flags byte) {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("flags")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	*(*C.guint8)(unsafe.Pointer(&*valptr)) = C.guint8(flags)
}

// Script: detected script for this segment (A PangoScript) (Since: 1.18).
func (a *Analysis) SetScript(script byte) {
	offset := girepository.MustFind("Pango", "Analysis").StructFieldOffset("script")
	valptr := (*uintptr)(unsafe.Add(a.native, offset))
	*(*C.guint8)(unsafe.Pointer(&*valptr)) = C.guint8(script)
}

// Item: PangoItem structure stores information about a segment of text.
//
// You typically obtain PangoItems by itemizing a piece of text with itemize.
//
// An instance of this type is always passed by reference.
type Item struct {
	*item
}

// item is the struct that's finalized.
type item struct {
	native unsafe.Pointer
}

func marshalItem(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Item{&item{(unsafe.Pointer)(b)}}, nil
}

// NewItem constructs a struct Item.
func NewItem() *Item {
	_info := girepository.MustFind("Pango", "Item")
	_gret := _info.InvokeRecordMethod("new", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _item *Item // out

	_item = (*Item)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_item)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("Pango", "Item").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _item
}

// ApplyAttrs: add attributes to a PangoItem.
//
// The idea is that you have attributes that don't affect itemization, such as
// font features, so you filter them out using pango.AttrList.Filter(), itemize
// your text, then reapply the attributes to the resulting items using this
// function.
//
// The iter should be positioned before the range of the item, and will be
// advanced past it. This function is meant to be called in a loop over the
// items resulting from itemization, while passing the iter to each call.
//
// The function takes the following parameters:
//
//    - iter: PangoAttrIterator.
//
func (item *Item) ApplyAttrs(iter *AttrIterator) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(item)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_info := girepository.MustFind("Pango", "Item")
	_info.InvokeRecordMethod("apply_attrs", _args[:], nil)

	runtime.KeepAlive(item)
	runtime.KeepAlive(iter)
}

// Copy an existing PangoItem structure.
//
// The function returns the following values:
//
//    - ret (optional): newly allocated PangoItem, which should be freed with
//      pango.Item.Free(), or NULL if item was NULL.
//
func (item *Item) Copy() *Item {
	var _args [1]girepository.Argument

	if item != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(item)))
	}

	_info := girepository.MustFind("Pango", "Item")
	_gret := _info.InvokeRecordMethod("copy", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(item)

	var _ret *Item // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_ret = (*Item)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_ret)),
			func(intern *struct{ C unsafe.Pointer }) {
				{
					var args [1]girepository.Argument
					*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
					girepository.MustFind("Pango", "Item").InvokeRecordMethod("free", args[:], nil)
				}
			},
		)
	}

	return _ret
}

// Split modifies orig to cover only the text after split_index, and returns a
// new item that covers the text before split_index that used to be in orig.
//
// You can think of split_index as the length of the returned item. split_index
// may not be 0, and it may not be greater than or equal to the length of orig
// (that is, there must be at least one byte assigned to each item, you can't
// create a zero-length item). split_offset is the length of the first item in
// chars, and must be provided because the text used to generate the item isn't
// available, so pango_item_split() can't count the char length of the split
// items itself.
//
// The function takes the following parameters:
//
//    - splitIndex: byte index of position to split item, relative to the start
//      of the item.
//    - splitOffset: number of chars between start of orig and split_index.
//
// The function returns the following values:
//
//    - item: new item representing text before split_index, which should be
//      freed with pango.Item.Free().
//
func (orig *Item) Split(splitIndex int32, splitOffset int32) *Item {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(orig)))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(splitIndex)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(splitOffset)

	_info := girepository.MustFind("Pango", "Item")
	_gret := _info.InvokeRecordMethod("split", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(orig)
	runtime.KeepAlive(splitIndex)
	runtime.KeepAlive(splitOffset)

	var _item *Item // out

	_item = (*Item)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_item)),
		func(intern *struct{ C unsafe.Pointer }) {
			{
				var args [1]girepository.Argument
				*(*unsafe.Pointer)(unsafe.Pointer(&args[0])) = unsafe.Pointer(intern.C)
				girepository.MustFind("Pango", "Item").InvokeRecordMethod("free", args[:], nil)
			}
		},
	)

	return _item
}
