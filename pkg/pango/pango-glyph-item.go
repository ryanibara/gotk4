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
import "C"

// glib.Type values for pango-glyph-item.go.
var (
	GTypeGlyphItem     = coreglib.Type(C.pango_glyph_item_get_type())
	GTypeGlyphItemIter = coreglib.Type(C.pango_glyph_item_iter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeGlyphItem, F: marshalGlyphItem},
		{T: GTypeGlyphItemIter, F: marshalGlyphItemIter},
	})
}

// GlyphItem: PangoGlyphItem is a pair of a PangoItem and the glyphs resulting
// from shaping the items text.
//
// As an example of the usage of PangoGlyphItem, the results of shaping text
// with PangoLayout is a list of PangoLayoutLine, each of which contains a list
// of PangoGlyphItem.
//
// An instance of this type is always passed by reference.
type GlyphItem struct {
	*glyphItem
}

// glyphItem is the struct that's finalized.
type glyphItem struct {
	native unsafe.Pointer
}

func marshalGlyphItem(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &GlyphItem{&glyphItem{(unsafe.Pointer)(b)}}, nil
}

// Item: corresponding PangoItem.
func (g *GlyphItem) Item() *Item {
	offset := girepository.MustFind("Pango", "GlyphItem").StructFieldOffset("item")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v *Item // out
	v = (*Item)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return v
}

// Glyphs: corresponding PangoGlyphString.
func (g *GlyphItem) Glyphs() *GlyphString {
	offset := girepository.MustFind("Pango", "GlyphItem").StructFieldOffset("glyphs")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v *GlyphString // out
	v = (*GlyphString)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return v
}

// ApplyAttrs splits a shaped item (PangoGlyphItem) into multiple items based on
// an attribute list.
//
// The idea is that if you have attributes that don't affect shaping, such as
// color or underline, to avoid affecting shaping, you filter them out
// (pango.AttrList.Filter()), apply the shaping process and then reapply them to
// the result using this function.
//
// All attributes that start or end inside a cluster are applied to that
// cluster; for instance, if half of a cluster is underlined and the other-half
// strikethrough, then the cluster will end up with both underline and
// strikethrough attributes. In these cases, it may happen that
// item->extra_attrs for some of the result items can have multiple attributes
// of the same type.
//
// This function takes ownership of glyph_item; it will be reused as one of the
// elements in the list.
//
// The function takes the following parameters:
//
//    - text that list applies to.
//    - list: PangoAttrList.
//
// The function returns the following values:
//
//    - sList: a list of glyph items resulting from splitting glyph_item. Free
//      the elements using pango.GlyphItem.Free(), the list using g_slist_free().
//
func (glyphItem *GlyphItem) ApplyAttrs(text string, list *AttrList) []*GlyphItem {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(list)))

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(text)
	runtime.KeepAlive(list)

	var _sList []*GlyphItem // out

	_sList = make([]*GlyphItem, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.void)(v)
		var dst *GlyphItem // out
		dst = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_glyph_item_free((*C.PangoGlyphItem)(intern.C))
			},
		)
		_sList = append(_sList, dst)
	})

	return _sList
}

// Copy: make a deep copy of an existing PangoGlyphItem structure.
//
// The function returns the following values:
//
//    - glyphItem (optional): newly allocated PangoGlyphItem, which should be
//      freed with pango_glyph_item_free(), or NULL if orig was NULL.
//
func (orig *GlyphItem) Copy() *GlyphItem {
	var _args [1]girepository.Argument

	if orig != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(orig)))
	}

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(orig)

	var _glyphItem *GlyphItem // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_glyphItem = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_glyphItem)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_glyph_item_free((*C.PangoGlyphItem)(intern.C))
			},
		)
	}

	return _glyphItem
}

// Split modifies orig to cover only the text after split_index, and returns a
// new item that covers the text before split_index that used to be in orig.
//
// You can think of split_index as the length of the returned item. split_index
// may not be 0, and it may not be greater than or equal to the length of orig
// (that is, there must be at least one byte assigned to each item, you can't
// create a zero-length item).
//
// This function is similar in function to pango_item_split() (and uses it
// internally.).
//
// The function takes the following parameters:
//
//    - text to which positions in orig apply.
//    - splitIndex: byte index of position to split item, relative to the start
//      of the item.
//
// The function returns the following values:
//
//    - glyphItem: newly allocated item representing text before split_index,
//      which should be freed with pango_glyph_item_free().
//
func (orig *GlyphItem) Split(text string, splitIndex int32) *GlyphItem {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(orig)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[1]))
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(splitIndex)

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(orig)
	runtime.KeepAlive(text)
	runtime.KeepAlive(splitIndex)

	var _glyphItem *GlyphItem // out

	_glyphItem = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_glyphItem)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.pango_glyph_item_free((*C.PangoGlyphItem)(intern.C))
		},
	)

	return _glyphItem
}

// GlyphItemIter: PangoGlyphItemIter is an iterator over the clusters in a
// PangoGlyphItem.
//
// The *forward direction* of the iterator is the logical direction of text.
// That is, with increasing start_index and start_char values. If glyph_item is
// right-to-left (that is, if glyph_item->item->analysis.level is odd), then
// start_glyph decreases as the iterator moves forward. Moreover, in
// right-to-left cases, start_glyph is greater than end_glyph.
//
// An iterator should be initialized using either
// pango_glyph_item_iter_init_start() or pango_glyph_item_iter_init_end(), for
// forward and backward iteration respectively, and walked over using any
// desired mixture of pango_glyph_item_iter_next_cluster() and
// pango_glyph_item_iter_prev_cluster().
//
// A common idiom for doing a forward iteration over the clusters is:
//
//    PangoGlyphItemIter cluster_iter;
//    gboolean have_cluster;
//
//    for (have_cluster = pango_glyph_item_iter_init_start (&cluster_iter,
//                                                          glyph_item, text);
//         have_cluster;
//         have_cluster = pango_glyph_item_iter_next_cluster (&cluster_iter))
//    {
//      ...
//    }
//
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset to get to the text of glyph_item. The start_index
// and end_index values can directly index into text. The start_glyph,
// end_glyph, start_char, and end_char values however are zero-based for the
// glyph_item. For each cluster, the item pointed at by the start variables is
// included in the cluster while the one pointed at by end variables is not.
//
// None of the members of a PangoGlyphItemIter should be modified manually.
//
// An instance of this type is always passed by reference.
type GlyphItemIter struct {
	*glyphItemIter
}

// glyphItemIter is the struct that's finalized.
type glyphItemIter struct {
	native unsafe.Pointer
}

func marshalGlyphItemIter(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &GlyphItemIter{&glyphItemIter{(unsafe.Pointer)(b)}}, nil
}

func (g *GlyphItemIter) GlyphItem() *GlyphItem {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("glyph_item")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v *GlyphItem // out
	v = (*GlyphItem)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return v
}

func (g *GlyphItemIter) Text() string {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("text")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(valptr)))
	return v
}

func (g *GlyphItemIter) StartGlyph() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("start_glyph")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

func (g *GlyphItemIter) StartIndex() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("start_index")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

func (g *GlyphItemIter) StartChar() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("start_char")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

func (g *GlyphItemIter) EndGlyph() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("end_glyph")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

func (g *GlyphItemIter) EndIndex() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("end_index")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

func (g *GlyphItemIter) EndChar() int32 {
	offset := girepository.MustFind("Pango", "GlyphItemIter").StructFieldOffset("end_char")
	valptr := unsafe.Add(unsafe.Pointer(g), offset)
	var v int32 // out
	v = int32(*(*C.int)(unsafe.Pointer(&valptr)))
	return v
}

// Copy: make a shallow copy of an existing PangoGlyphItemIter structure.
//
// The function returns the following values:
//
//    - glyphItemIter (optional): newly allocated PangoGlyphItemIter, which
//      should be freed with pango_glyph_item_iter_free(), or NULL if orig was
//      NULL.
//
func (orig *GlyphItemIter) Copy() *GlyphItemIter {
	var _args [1]girepository.Argument

	if orig != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(orig)))
	}

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(orig)

	var _glyphItemIter *GlyphItemIter // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_glyphItemIter = (*GlyphItemIter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_glyphItemIter)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_glyph_item_iter_free((*C.PangoGlyphItemIter)(intern.C))
			},
		)
	}

	return _glyphItemIter
}

// InitEnd initializes a PangoGlyphItemIter structure to point to the last
// cluster in a glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
//
// The function takes the following parameters:
//
//    - glyphItem: glyph item to iterate over.
//    - text corresponding to the glyph item.
//
// The function returns the following values:
//
//    - ok: FALSE if there are no clusters in the glyph item.
//
func (iter *GlyphItemIter) InitEnd(glyphItem *GlyphItem, text string) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)
	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(text)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// InitStart initializes a PangoGlyphItemIter structure to point to the first
// cluster in a glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
//
// The function takes the following parameters:
//
//    - glyphItem: glyph item to iterate over.
//    - text corresponding to the glyph item.
//
// The function returns the following values:
//
//    - ok: FALSE if there are no clusters in the glyph item.
//
func (iter *GlyphItemIter) InitStart(glyphItem *GlyphItem, text string) bool {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(glyphItem)))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[2]))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)
	runtime.KeepAlive(glyphItem)
	runtime.KeepAlive(text)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// NextCluster advances the iterator to the next cluster in the glyph item.
//
// See PangoGlyphItemIter for details of cluster orders.
//
// The function returns the following values:
//
//    - ok: TRUE if the iterator was advanced, FALSE if we were already on the
//      last cluster.
//
func (iter *GlyphItemIter) NextCluster() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}

// PrevCluster moves the iterator to the preceding cluster in the glyph item.
// See PangoGlyphItemIter for details of cluster orders.
//
// The function returns the following values:
//
//    - ok: TRUE if the iterator was moved, FALSE if we were already on the first
//      cluster.
//
func (iter *GlyphItemIter) PrevCluster() bool {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(iter)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
