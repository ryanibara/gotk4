// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_hyperlink_state_flags_get_type()), F: marshalHyperlinkStateFlags},
		{T: externglib.Type(C.atk_hyperlink_get_type()), F: marshalHyperlinker},
	})
}

// HyperlinkStateFlags describes the type of link
type HyperlinkStateFlags int

const (
	// HyperlinkStateFlagsInline: link is inline
	HyperlinkStateFlagsInline HyperlinkStateFlags = 0b1
)

func marshalHyperlinkStateFlags(p uintptr) (interface{}, error) {
	return HyperlinkStateFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// HyperlinkOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HyperlinkOverrider interface {
	// EndIndex gets the index with the hypertext document at which this link
	// ends.
	EndIndex() int
	// NAnchors gets the number of anchors associated with this hyperlink.
	NAnchors() int
	// GetObject returns the item associated with this hyperlinks nth anchor.
	// For instance, the returned Object will implement Text if @link_ is a text
	// hyperlink, Image if @link_ is an image hyperlink etc.
	//
	// Multiple anchors are primarily used by client-side image maps.
	GetObject(i int) *ObjectClass
	// StartIndex gets the index with the hypertext document at which this link
	// begins.
	StartIndex() int
	// URI: get a the URI associated with the anchor specified by @i of @link_.
	//
	// Multiple anchors are primarily used by client-side image maps.
	URI(i int) string
	// IsSelectedLink determines whether this AtkHyperlink is selected
	//
	// Deprecated: Please use ATK_STATE_FOCUSABLE for all links, and
	// ATK_STATE_FOCUSED for focused links.
	IsSelectedLink() bool
	// IsValid: since the document that a link is associated with may have
	// changed this method returns true if the link is still valid (with respect
	// to the document it references) and false otherwise.
	IsValid() bool
	LinkActivated()
	LinkState() uint
}

// Hyperlinker describes Hyperlink's methods.
type Hyperlinker interface {
	// EndIndex gets the index with the hypertext document at which this link
	// ends.
	EndIndex() int
	// NAnchors gets the number of anchors associated with this hyperlink.
	NAnchors() int
	// GetObject returns the item associated with this hyperlinks nth anchor.
	GetObject(i int) *ObjectClass
	// StartIndex gets the index with the hypertext document at which this link
	// begins.
	StartIndex() int
	// URI: get a the URI associated with the anchor specified by @i of @link_.
	URI(i int) string
	// IsInline indicates whether the link currently displays some or all of its
	// content inline.
	IsInline() bool
	// IsSelectedLink determines whether this AtkHyperlink is selected
	// Deprecated: Please use ATK_STATE_FOCUSABLE for all links, and
	// ATK_STATE_FOCUSED for focused links.
	IsSelectedLink() bool
	// IsValid: since the document that a link is associated with may have
	// changed this method returns true if the link is still valid (with respect
	// to the document it references) and false otherwise.
	IsValid() bool
}

// Hyperlink: ATK object which encapsulates a link or set of links (for instance
// in the case of client-side image maps) in a hypertext document. It may
// implement the AtkAction interface. AtkHyperlink may also be used to refer to
// inline embedded content, since it allows specification of a start and end
// offset within the host AtkHypertext object.
type Hyperlink struct {
	*externglib.Object

	Action
}

var (
	_ Hyperlinker     = (*Hyperlink)(nil)
	_ gextras.Nativer = (*Hyperlink)(nil)
)

func wrapHyperlink(obj *externglib.Object) *Hyperlink {
	return &Hyperlink{
		Object: obj,
		Action: Action{
			Object: obj,
		},
	}
}

func marshalHyperlinker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHyperlink(obj), nil
}

// EndIndex gets the index with the hypertext document at which this link ends.
func (link_ *Hyperlink) EndIndex() int {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_get_end_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NAnchors gets the number of anchors associated with this hyperlink.
func (link_ *Hyperlink) NAnchors() int {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_get_n_anchors(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GetObject returns the item associated with this hyperlinks nth anchor. For
// instance, the returned Object will implement Text if @link_ is a text
// hyperlink, Image if @link_ is an image hyperlink etc.
//
// Multiple anchors are primarily used by client-side image maps.
func (link_ *Hyperlink) GetObject(i int) *ObjectClass {
	var _arg0 *C.AtkHyperlink // out
	var _arg1 C.gint          // out
	var _cret *C.AtkObject    // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_hyperlink_get_object(_arg0, _arg1)

	var _object *ObjectClass // out

	_object = wrapObject(externglib.Take(unsafe.Pointer(_cret)))

	return _object
}

// StartIndex gets the index with the hypertext document at which this link
// begins.
func (link_ *Hyperlink) StartIndex() int {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gint          // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_get_start_index(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// URI: get a the URI associated with the anchor specified by @i of @link_.
//
// Multiple anchors are primarily used by client-side image maps.
func (link_ *Hyperlink) URI(i int) string {
	var _arg0 *C.AtkHyperlink // out
	var _arg1 C.gint          // out
	var _cret *C.gchar        // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))
	_arg1 = C.gint(i)

	_cret = C.atk_hyperlink_get_uri(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// IsInline indicates whether the link currently displays some or all of its
// content inline. Ordinary HTML links will usually return false, but an inline
// &lt;src&gt; HTML element will return true.
func (link_ *Hyperlink) IsInline() bool {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_is_inline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSelectedLink determines whether this AtkHyperlink is selected
//
// Deprecated: Please use ATK_STATE_FOCUSABLE for all links, and
// ATK_STATE_FOCUSED for focused links.
func (link_ *Hyperlink) IsSelectedLink() bool {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_is_selected_link(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsValid: since the document that a link is associated with may have changed
// this method returns true if the link is still valid (with respect to the
// document it references) and false otherwise.
func (link_ *Hyperlink) IsValid() bool {
	var _arg0 *C.AtkHyperlink // out
	var _cret C.gboolean      // in

	_arg0 = (*C.AtkHyperlink)(unsafe.Pointer(link_.Native()))

	_cret = C.atk_hyperlink_is_valid(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
