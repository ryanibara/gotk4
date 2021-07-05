// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
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
		{T: externglib.Type(C.atk_document_get_type()), F: marshalDocument},
	})
}

// Document: the AtkDocument interface should be supported by any object whose
// content is a representation or view of a document. The AtkDocument interface
// should appear on the toplevel container for the document content; however
// AtkDocument instances may be nested (i.e. an AtkDocument may be a descendant
// of another AtkDocument) in those cases where one document contains "embedded
// content" which can reasonably be considered a document in its own right.
type Document interface {
	gextras.Objector

	// AttributeValue retrieves the value of the given @attribute_name inside
	// @document.
	AttributeValue(attributeName string) string
	// CurrentPageNumber retrieves the current page number inside @document.
	CurrentPageNumber() int
	// Document gets a gpointer that points to an instance of the DOM. It is up
	// to the caller to check atk_document_get_type to determine how to cast
	// this pointer.
	//
	// Deprecated: since version .
	Document() interface{}
	// DocumentType gets a string indicating the document type.
	//
	// Deprecated: since version .
	DocumentType() string
	// Locale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale
	// of the content of this document instance. Individual text substrings or
	// images within this document may have a different locale, see
	// atk_text_get_attributes and atk_image_get_image_locale.
	//
	// Deprecated: since version 2.7.90.
	Locale() string
	// PageCount retrieves the total number of pages inside @document.
	PageCount() int
	// SetAttributeValue sets the value for the given @attribute_name inside
	// @document.
	SetAttributeValue(attributeName string, attributeValue string) bool
}

// document implements the Document interface.
type document struct {
	gextras.Objector
}

var _ Document = (*document)(nil)

// WrapDocument wraps a GObject to a type that implements
// interface Document. It is primarily used internally.
func WrapDocument(obj *externglib.Object) Document {
	return document{
		Objector: obj,
	}
}

func marshalDocument(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDocument(obj), nil
}

func (d document) AttributeValue(attributeName string) string {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(attributeName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_document_get_attribute_value(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d document) CurrentPageNumber() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))

	_cret = C.atk_document_get_current_page_number(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d document) Document() interface{} {
	var _arg0 *C.AtkDocument // out
	var _cret C.gpointer     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))

	_cret = C.atk_document_get_document(_arg0)

	var _gpointer interface{} // out

	_gpointer = box.Get(uintptr(_cret))

	return _gpointer
}

func (d document) DocumentType() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))

	_cret = C.atk_document_get_document_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d document) Locale() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))

	_cret = C.atk_document_get_locale(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d document) PageCount() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))

	_cret = C.atk_document_get_page_count(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (d document) SetAttributeValue(attributeName string, attributeValue string) bool {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.gchar)(C.CString(attributeName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(attributeValue))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.atk_document_set_attribute_value(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}