// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

// glib.Type values for gmarkup.go.
var GTypeMarkupParseContext = coreglib.Type(C.g_markup_parse_context_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeMarkupParseContext, F: marshalMarkupParseContext},
	})
}

// MarkupError: error codes returned by markup parsing.
type MarkupError C.gint

const (
	// MarkupErrorBadUTF8: text being parsed was not valid UTF-8.
	MarkupErrorBadUTF8 MarkupError = iota
	// MarkupErrorEmpty: document contained nothing, or only whitespace.
	MarkupErrorEmpty
	// MarkupErrorParse: document was ill-formed.
	MarkupErrorParse
	// MarkupErrorUnknownElement: error should be set by Parser functions;
	// element wasn't known.
	MarkupErrorUnknownElement
	// MarkupErrorUnknownAttribute: error should be set by Parser functions;
	// attribute wasn't known.
	MarkupErrorUnknownAttribute
	// MarkupErrorInvalidContent: error should be set by Parser functions;
	// content was invalid.
	MarkupErrorInvalidContent
	// MarkupErrorMissingAttribute: error should be set by Parser functions; a
	// required attribute was missing.
	MarkupErrorMissingAttribute
)

// String returns the name in string for MarkupError.
func (m MarkupError) String() string {
	switch m {
	case MarkupErrorBadUTF8:
		return "BadUTF8"
	case MarkupErrorEmpty:
		return "Empty"
	case MarkupErrorParse:
		return "Parse"
	case MarkupErrorUnknownElement:
		return "UnknownElement"
	case MarkupErrorUnknownAttribute:
		return "UnknownAttribute"
	case MarkupErrorInvalidContent:
		return "InvalidContent"
	case MarkupErrorMissingAttribute:
		return "MissingAttribute"
	default:
		return fmt.Sprintf("MarkupError(%d)", m)
	}
}

// MarkupCollectType: mixed enumerated type and flags field. You must specify
// one type (string, strdup, boolean, tristate). Additionally, you may
// optionally bitwise OR the type with the flag G_MARKUP_COLLECT_OPTIONAL.
//
// It is likely that this enum will be extended in the future to support other
// types.
type MarkupCollectType C.guint

const (
	// MarkupCollectInvalid: used to terminate the list of attributes to
	// collect.
	MarkupCollectInvalid MarkupCollectType = 0b0
	// MarkupCollectString: collect the string pointer directly from the
	// attribute_values[] array. Expects a parameter of type (const char **). If
	// G_MARKUP_COLLECT_OPTIONAL is specified and the attribute isn't present
	// then the pointer will be set to NULL.
	MarkupCollectString MarkupCollectType = 0b1
	// MarkupCollectStrdup as with G_MARKUP_COLLECT_STRING, but expects a
	// parameter of type (char **) and g_strdup()s the returned pointer. The
	// pointer must be freed with g_free().
	MarkupCollectStrdup MarkupCollectType = 0b10
	// MarkupCollectBoolean expects a parameter of type (gboolean *) and parses
	// the attribute value as a boolean. Sets FALSE if the attribute isn't
	// present. Valid boolean values consist of (case-insensitive) "false", "f",
	// "no", "n", "0" and "true", "t", "yes", "y", "1".
	MarkupCollectBoolean MarkupCollectType = 0b11
	// MarkupCollectTristate as with G_MARKUP_COLLECT_BOOLEAN, but in the case
	// of a missing attribute a value is set that compares equal to neither
	// FALSE nor TRUE G_MARKUP_COLLECT_OPTIONAL is implied.
	MarkupCollectTristate MarkupCollectType = 0b100
	// MarkupCollectOptional: can be bitwise ORed with the other fields. If
	// present, allows the attribute not to appear. A default value is set
	// depending on what value type is used.
	MarkupCollectOptional MarkupCollectType = 0b10000000000000000
)

// String returns the names in string for MarkupCollectType.
func (m MarkupCollectType) String() string {
	if m == 0 {
		return "MarkupCollectType(0)"
	}

	var builder strings.Builder
	builder.Grow(125)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case MarkupCollectInvalid:
			builder.WriteString("Invalid|")
		case MarkupCollectString:
			builder.WriteString("String|")
		case MarkupCollectStrdup:
			builder.WriteString("Strdup|")
		case MarkupCollectBoolean:
			builder.WriteString("Boolean|")
		case MarkupCollectTristate:
			builder.WriteString("Tristate|")
		case MarkupCollectOptional:
			builder.WriteString("Optional|")
		default:
			builder.WriteString(fmt.Sprintf("MarkupCollectType(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m MarkupCollectType) Has(other MarkupCollectType) bool {
	return (m & other) == other
}

// MarkupParseFlags flags that affect the behaviour of the parser.
type MarkupParseFlags C.guint

const (
	// MarkupDoNotUseThisUnsupportedFlag: flag you should not use.
	MarkupDoNotUseThisUnsupportedFlag MarkupParseFlags = 0b1
	// MarkupTreatCdataAsText: when this flag is set, CDATA marked sections are
	// not passed literally to the passthrough function of the parser. Instead,
	// the content of the section (without the <![CDATA[ and ]]>) is passed to
	// the text function. This flag was added in GLib 2.12.
	MarkupTreatCdataAsText MarkupParseFlags = 0b10
	// MarkupPrefixErrorPosition: normally errors caught by GMarkup itself have
	// line/column information prefixed to them to let the caller know the
	// location of the error. When this flag is set the location information is
	// also prefixed to errors generated by the Parser implementation functions.
	MarkupPrefixErrorPosition MarkupParseFlags = 0b100
	// MarkupIgnoreQualified: ignore (don't report) qualified attributes and
	// tags, along with their contents. A qualified attribute or tag is one that
	// contains ':' in its name (ie: is in another namespace). Since: 2.40.
	MarkupIgnoreQualified MarkupParseFlags = 0b1000
)

// String returns the names in string for MarkupParseFlags.
func (m MarkupParseFlags) String() string {
	if m == 0 {
		return "MarkupParseFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(104)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case MarkupDoNotUseThisUnsupportedFlag:
			builder.WriteString("DoNotUseThisUnsupportedFlag|")
		case MarkupTreatCdataAsText:
			builder.WriteString("TreatCdataAsText|")
		case MarkupPrefixErrorPosition:
			builder.WriteString("PrefixErrorPosition|")
		case MarkupIgnoreQualified:
			builder.WriteString("IgnoreQualified|")
		default:
			builder.WriteString(fmt.Sprintf("MarkupParseFlags(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m MarkupParseFlags) Has(other MarkupParseFlags) bool {
	return (m & other) == other
}

// MarkupEscapeText escapes text so that the markup parser will parse it
// verbatim. Less than, greater than, ampersand, etc. are replaced with the
// corresponding entities. This function would typically be used when writing
// out a file to be parsed with the markup parser.
//
// Note that this function doesn't protect whitespace and line endings from
// being processed according to the XML rules for normalization of line endings
// and attribute values.
//
// Note also that this function will produce character references in the range
// of &#x1; ... &#x1f; for all control sequences except for tabstop, newline and
// carriage return. The character references in this range are not valid XML
// 1.0, but they are valid XML 1.1 and will be accepted by the GMarkup parser.
//
// The function takes the following parameters:
//
//    - text: some valid UTF-8 text.
//    - length of text in bytes, or -1 if the text is nul-terminated.
//
// The function returns the following values:
//
//    - utf8: newly allocated string with the escaped text.
//
func MarkupEscapeText(text string, length int) string {
	var _arg1 *C.gchar // out
	var _arg2 C.gssize // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.g_markup_escape_text(_arg1, _arg2)
	runtime.KeepAlive(text)
	runtime.KeepAlive(length)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MarkupParseContext: parse context is used to parse a stream of bytes that you
// expect to contain marked-up text.
//
// See g_markup_parse_context_new(), Parser, and so on for more details.
//
// An instance of this type is always passed by reference.
type MarkupParseContext struct {
	*markupParseContext
}

// markupParseContext is the struct that's finalized.
type markupParseContext struct {
	native *C.GMarkupParseContext
}

func marshalMarkupParseContext(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &MarkupParseContext{&markupParseContext{(*C.GMarkupParseContext)(b)}}, nil
}

// EndParse signals to the ParseContext that all data has been fed into the
// parse context with g_markup_parse_context_parse().
//
// This function reports an error if the document isn't complete, for example if
// elements are still open.
func (context *MarkupParseContext) EndParse() error {
	var _arg0 *C.GMarkupParseContext // out
	var _cerr *C.GError              // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	C.g_markup_parse_context_end_parse(_arg0, &_cerr)
	runtime.KeepAlive(context)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Element retrieves the name of the currently open element.
//
// If called from the start_element or end_element handlers this will give the
// element_name as passed to those functions. For the parent elements, see
// g_markup_parse_context_get_element_stack().
//
// The function returns the following values:
//
//    - utf8: name of the currently open element, or NULL.
//
func (context *MarkupParseContext) Element() string {
	var _arg0 *C.GMarkupParseContext // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	_cret = C.g_markup_parse_context_get_element(_arg0)
	runtime.KeepAlive(context)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Position retrieves the current line number and the number of the character on
// that line. Intended for use in error messages; there are no strict semantics
// for what constitutes the "current" line number other than "the best number we
// could come up with for error messages.".
//
// The function returns the following values:
//
//    - lineNumber (optional): return location for a line number, or NULL.
//    - charNumber (optional): return location for a char-on-line number, or
//      NULL.
//
func (context *MarkupParseContext) Position() (lineNumber int32, charNumber int32) {
	var _arg0 *C.GMarkupParseContext // out
	var _arg1 C.gint                 // in
	var _arg2 C.gint                 // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))

	C.g_markup_parse_context_get_position(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(context)

	var _lineNumber int32 // out
	var _charNumber int32 // out

	_lineNumber = int32(_arg1)
	_charNumber = int32(_arg2)

	return _lineNumber, _charNumber
}

// Parse: feed some data to the ParseContext.
//
// The data need not be valid UTF-8; an error will be signaled if it's invalid.
// The data need not be an entire document; you can feed a document into the
// parser incrementally, via multiple calls to this function. Typically, as you
// receive data from a network connection or file, you feed each received chunk
// of data into this function, aborting the process if an error occurs. Once an
// error is reported, no further data may be fed to the ParseContext; all errors
// are fatal.
//
// The function takes the following parameters:
//
//    - text: chunk of text to parse.
//    - textLen: length of text in bytes.
//
func (context *MarkupParseContext) Parse(text string, textLen int) error {
	var _arg0 *C.GMarkupParseContext // out
	var _arg1 *C.gchar               // out
	var _arg2 C.gssize               // out
	var _cerr *C.GError              // in

	_arg0 = (*C.GMarkupParseContext)(gextras.StructNative(unsafe.Pointer(context)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(textLen)

	C.g_markup_parse_context_parse(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(context)
	runtime.KeepAlive(text)
	runtime.KeepAlive(textLen)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// MarkupParser: any of the fields in Parser can be NULL, in which case they
// will be ignored. Except for the error function, any of these callbacks can
// set an error; in particular the G_MARKUP_ERROR_UNKNOWN_ELEMENT,
// G_MARKUP_ERROR_UNKNOWN_ATTRIBUTE, and G_MARKUP_ERROR_INVALID_CONTENT errors
// are intended to be set from these callbacks. If you set an error from a
// callback, g_markup_parse_context_parse() will report that error back to its
// caller.
//
// An instance of this type is always passed by reference.
type MarkupParser struct {
	*markupParser
}

// markupParser is the struct that's finalized.
type markupParser struct {
	native *C.GMarkupParser
}
