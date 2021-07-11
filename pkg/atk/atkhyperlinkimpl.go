// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_hyperlink_impl_get_type()), F: marshalHyperlinkImpler},
	})
}

// HyperlinkImplOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type HyperlinkImplOverrider interface {
	// Hyperlink gets the hyperlink associated with this object.
	Hyperlink() *Hyperlink
}

// HyperlinkImpler describes HyperlinkImpl's methods.
type HyperlinkImpler interface {
	// Hyperlink gets the hyperlink associated with this object.
	Hyperlink() *Hyperlink
}

// HyperlinkImpl allows AtkObjects to refer to their associated AtkHyperlink
// instance, if one exists. AtkHyperlinkImpl differs from AtkHyperlink in that
// AtkHyperlinkImpl is an interface, whereas AtkHyperlink is a object type. The
// AtkHyperlinkImpl interface allows a client to query an AtkObject for the
// availability of an associated AtkHyperlink instance, and obtain that
// instance. It is thus particularly useful in cases where embedded content or
// inline content within a text object is present, since the embedding text
// object implements AtkHypertext and the inline/embedded objects are exposed as
// children which implement AtkHyperlinkImpl, in addition to their being
// obtainable via AtkHypertext:getLink followed by AtkHyperlink:getObject.
//
// The AtkHyperlinkImpl interface should be supported by objects exposed within
// the hierarchy as children of an AtkHypertext container which correspond to
// "links" or embedded content within the text. HTML anchors are not, for
// instance, normally exposed this way, but embedded images and components which
// appear inline in the content of a text object are. The AtkHyperlinkIface
// interface allows a means of determining which children are hyperlinks in this
// sense of the word, and for obtaining their corresponding AtkHyperlink object,
// from which the embedding range, URI, etc. can be obtained.
//
// To some extent this interface exists because, for historical reasons,
// AtkHyperlink was defined as an object type, not an interface. Thus, in order
// to interact with AtkObjects via AtkHyperlink semantics, a new interface was
// required.
type HyperlinkImpl struct {
	*externglib.Object
}

var (
	_ HyperlinkImpler = (*HyperlinkImpl)(nil)
	_ gextras.Nativer = (*HyperlinkImpl)(nil)
)

func wrapHyperlinkImpl(obj *externglib.Object) HyperlinkImpler {
	return &HyperlinkImpl{
		Object: obj,
	}
}

func marshalHyperlinkImpler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHyperlinkImpl(obj), nil
}

// Hyperlink gets the hyperlink associated with this object.
func (impl *HyperlinkImpl) Hyperlink() *Hyperlink {
	var _arg0 *C.AtkHyperlinkImpl // out
	var _cret *C.AtkHyperlink     // in

	_arg0 = (*C.AtkHyperlinkImpl)(unsafe.Pointer(impl.Native()))

	_cret = C.atk_hyperlink_impl_get_hyperlink(_arg0)

	var _hyperlink *Hyperlink // out

	_hyperlink = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*Hyperlink)

	return _hyperlink
}
