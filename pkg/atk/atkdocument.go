// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
// extern gchar* _gotk4_atk1_DocumentIface_get_document_type(AtkDocument*);
// extern gint _gotk4_atk1_DocumentIface_get_current_page_number(AtkDocument*);
// extern gint _gotk4_atk1_DocumentIface_get_page_count(AtkDocument*);
// extern gpointer _gotk4_atk1_DocumentIface_get_document(AtkDocument*);
// extern void _gotk4_atk1_Document_ConnectLoadComplete(gpointer, guintptr);
// extern void _gotk4_atk1_Document_ConnectLoadStopped(gpointer, guintptr);
// extern void _gotk4_atk1_Document_ConnectPageChanged(gpointer, gint, guintptr);
// extern void _gotk4_atk1_Document_ConnectReload(gpointer, guintptr);
import "C"

// glib.Type values for atkdocument.go.
var GTypeDocument = externglib.Type(C.atk_document_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeDocument, F: marshalDocument},
	})
}

// Document interface should be supported by any object whose content is a
// representation or view of a document. The AtkDocument interface should appear
// on the toplevel container for the document content; however AtkDocument
// instances may be nested (i.e. an AtkDocument may be a descendant of another
// AtkDocument) in those cases where one document contains "embedded content"
// which can reasonably be considered a document in its own right.
//
// Document wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Document struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Document)(nil)
)

// Documenter describes Document's interface methods.
type Documenter interface {
	externglib.Objector

	// AttributeValue retrieves the value of the given attribute_name inside
	// document.
	AttributeValue(attributeName string) string
	// CurrentPageNumber retrieves the current page number inside document.
	CurrentPageNumber() int
	// Document gets a gpointer that points to an instance of the DOM.
	Document() cgo.Handle
	// DocumentType gets a string indicating the document type.
	DocumentType() string
	// Locale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale
	// of the content of this document instance.
	Locale() string
	// PageCount retrieves the total number of pages inside document.
	PageCount() int
	// SetAttributeValue sets the value for the given attribute_name inside
	// document.
	SetAttributeValue(attributeName, attributeValue string) bool

	// Load-complete: 'load-complete' signal is emitted when a pending load of a
	// static document has completed.
	ConnectLoadComplete(func()) externglib.SignalHandle
	// Load-stopped: 'load-stopped' signal is emitted when a pending load of
	// document contents is cancelled, paused, or otherwise interrupted by the
	// user or application logic.
	ConnectLoadStopped(func()) externglib.SignalHandle
	// Page-changed: 'page-changed' signal is emitted when the current page of a
	// document changes, e.g.
	ConnectPageChanged(func(pageNumber int)) externglib.SignalHandle
	// Reload: 'reload' signal is emitted when the contents of a document is
	// refreshed from its source.
	ConnectReload(func()) externglib.SignalHandle
}

var _ Documenter = (*Document)(nil)

func wrapDocument(obj *externglib.Object) *Document {
	return &Document{
		Object: obj,
	}
}

func marshalDocument(p uintptr) (interface{}, error) {
	return wrapDocument(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

//export _gotk4_atk1_Document_ConnectLoadComplete
func _gotk4_atk1_Document_ConnectLoadComplete(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectLoadComplete: 'load-complete' signal is emitted when a pending load of
// a static document has completed. This signal is to be expected by ATK clients
// if and when AtkDocument implementors expose ATK_STATE_BUSY. If the state of
// an AtkObject which implements AtkDocument does not include ATK_STATE_BUSY, it
// should be safe for clients to assume that the AtkDocument's static contents
// are fully loaded into the container. (Dynamic document contents should be
// exposed via other signals.).
func (document *Document) ConnectLoadComplete(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(document, "load-complete", false, unsafe.Pointer(C._gotk4_atk1_Document_ConnectLoadComplete), f)
}

//export _gotk4_atk1_Document_ConnectLoadStopped
func _gotk4_atk1_Document_ConnectLoadStopped(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectLoadStopped: 'load-stopped' signal is emitted when a pending load of
// document contents is cancelled, paused, or otherwise interrupted by the user
// or application logic. It should not however be emitted while waiting for a
// resource (for instance while blocking on a file or network read) unless a
// user-significant timeout has occurred.
func (document *Document) ConnectLoadStopped(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(document, "load-stopped", false, unsafe.Pointer(C._gotk4_atk1_Document_ConnectLoadStopped), f)
}

//export _gotk4_atk1_Document_ConnectPageChanged
func _gotk4_atk1_Document_ConnectPageChanged(arg0 C.gpointer, arg1 C.gint, arg2 C.guintptr) {
	var f func(pageNumber int)
	{
		closure := externglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(pageNumber int))
	}

	var _pageNumber int // out

	_pageNumber = int(arg1)

	f(_pageNumber)
}

// ConnectPageChanged: 'page-changed' signal is emitted when the current page of
// a document changes, e.g. pressing page up/down in a document viewer.
func (document *Document) ConnectPageChanged(f func(pageNumber int)) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(document, "page-changed", false, unsafe.Pointer(C._gotk4_atk1_Document_ConnectPageChanged), f)
}

//export _gotk4_atk1_Document_ConnectReload
func _gotk4_atk1_Document_ConnectReload(arg0 C.gpointer, arg1 C.guintptr) {
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

// ConnectReload: 'reload' signal is emitted when the contents of a document is
// refreshed from its source. Once 'reload' has been emitted, a matching
// 'load-complete' or 'load-stopped' signal should follow, which clients may
// await before interrogating ATK for the latest document content.
func (document *Document) ConnectReload(f func()) externglib.SignalHandle {
	return externglib.ConnectGeneratedClosure(document, "reload", false, unsafe.Pointer(C._gotk4_atk1_Document_ConnectReload), f)
}

// AttributeValue retrieves the value of the given attribute_name inside
// document.
//
// The function takes the following parameters:
//
//    - attributeName: character string representing the name of the attribute
//      whose value is being queried.
//
// The function returns the following values:
//
//    - utf8 (optional): string value associated with the named attribute for
//      this document, or NULL if a value for attribute_name has not been
//      specified for this document.
//
func (document *Document) AttributeValue(attributeName string) string {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attributeName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.atk_document_get_attribute_value(_arg0, _arg1)
	runtime.KeepAlive(document)
	runtime.KeepAlive(attributeName)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CurrentPageNumber retrieves the current page number inside document.
//
// The function returns the following values:
//
//    - gint: current page number inside document, or -1 if not implemented, not
//      know by the implementor, or irrelevant.
//
func (document *Document) CurrentPageNumber() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.atk_document_get_current_page_number(_arg0)
	runtime.KeepAlive(document)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Document gets a gpointer that points to an instance of the DOM. It is up to
// the caller to check atk_document_get_type to determine how to cast this
// pointer.
//
// Deprecated: Since 2.12. document is already a representation of the document.
// Use it directly, or one of its children, as an instance of the DOM.
//
// The function returns the following values:
//
//    - gpointer (optional) that points to an instance of the DOM.
//
func (document *Document) Document() cgo.Handle {
	var _arg0 *C.AtkDocument // out
	var _cret C.gpointer     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.atk_document_get_document(_arg0)
	runtime.KeepAlive(document)

	var _gpointer cgo.Handle // out

	_gpointer = (cgo.Handle)(unsafe.Pointer(_cret))

	return _gpointer
}

// DocumentType gets a string indicating the document type.
//
// Deprecated: Since 2.12. Please use atk_document_get_attributes() to ask for
// the document type if it applies.
//
// The function returns the following values:
//
//    - utf8: string indicating the document type.
//
func (document *Document) DocumentType() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.atk_document_get_document_type(_arg0)
	runtime.KeepAlive(document)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Locale gets a UTF-8 string indicating the POSIX-style LC_MESSAGES locale of
// the content of this document instance. Individual text substrings or images
// within this document may have a different locale, see atk_text_get_attributes
// and atk_image_get_image_locale.
//
// Deprecated: Please use atk_object_get_object_locale() instead.
//
// The function returns the following values:
//
//    - utf8: UTF-8 string indicating the POSIX-style LC_MESSAGES locale of the
//      document content as a whole, or NULL if the document content does not
//      specify a locale.
//
func (document *Document) Locale() string {
	var _arg0 *C.AtkDocument // out
	var _cret *C.gchar       // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.atk_document_get_locale(_arg0)
	runtime.KeepAlive(document)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageCount retrieves the total number of pages inside document.
//
// The function returns the following values:
//
//    - gint: total page count of document, or -1 if not implemented, not know by
//      the implementor or irrelevant.
//
func (document *Document) PageCount() int {
	var _arg0 *C.AtkDocument // out
	var _cret C.gint         // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))

	_cret = C.atk_document_get_page_count(_arg0)
	runtime.KeepAlive(document)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SetAttributeValue sets the value for the given attribute_name inside
// document.
//
// The function takes the following parameters:
//
//    - attributeName: character string representing the name of the attribute
//      whose value is being set.
//    - attributeValue: string value to be associated with attribute_name.
//
// The function returns the following values:
//
//    - ok: TRUE if attribute_value is successfully associated with
//      attribute_name for this document, and FALSE if if the document does not
//      allow the attribute to be modified.
//
func (document *Document) SetAttributeValue(attributeName, attributeValue string) bool {
	var _arg0 *C.AtkDocument // out
	var _arg1 *C.gchar       // out
	var _arg2 *C.gchar       // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkDocument)(unsafe.Pointer(externglib.InternObject(document).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(attributeName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(attributeValue)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.atk_document_set_attribute_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(document)
	runtime.KeepAlive(attributeName)
	runtime.KeepAlive(attributeValue)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
