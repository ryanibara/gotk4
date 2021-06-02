// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <stdbool.h>
// #include <glib.h>
import "C"

// ErrorType: the possible errors, used in the @v_error field of Value, when the
// token is a G_TOKEN_ERROR.
type ErrorType int

const (
	// ErrorTypeUnknown: unknown error
	ErrorTypeUnknown ErrorType = 0
	// ErrorTypeUnexpEOF: unexpected end of file
	ErrorTypeUnexpEOF ErrorType = 1
	// ErrorTypeUnexpEOFInString: unterminated string constant
	ErrorTypeUnexpEOFInString ErrorType = 2
	// ErrorTypeUnexpEOFInComment: unterminated comment
	ErrorTypeUnexpEOFInComment ErrorType = 3
	// ErrorTypeNonDigitInConst: non-digit character in a number
	ErrorTypeNonDigitInConst ErrorType = 4
	// ErrorTypeDigitRadix: digit beyond radix in a number
	ErrorTypeDigitRadix ErrorType = 5
	// ErrorTypeFloatRadix: non-decimal floating point number
	ErrorTypeFloatRadix ErrorType = 6
	// ErrorTypeFloatMalformed: malformed floating point number
	ErrorTypeFloatMalformed ErrorType = 7
)

// TokenType: the possible types of token returned from each
// g_scanner_get_next_token() call.
type TokenType int

const (
	// TokenTypeEOF: the end of the file
	TokenTypeEOF TokenType = 0
	// TokenTypeLeftParen: a '(' character
	TokenTypeLeftParen TokenType = 40
	// TokenTypeRightParen: a ')' character
	TokenTypeRightParen TokenType = 41
	// TokenTypeLeftCurly: a '{' character
	TokenTypeLeftCurly TokenType = 123
	// TokenTypeRightCurly: a '}' character
	TokenTypeRightCurly TokenType = 125
	// TokenTypeLeftBrace: a '[' character
	TokenTypeLeftBrace TokenType = 91
	// TokenTypeRightBrace: a ']' character
	TokenTypeRightBrace TokenType = 93
	// TokenTypeEqualSign: a '=' character
	TokenTypeEqualSign TokenType = 61
	// TokenTypeComma: a ',' character
	TokenTypeComma TokenType = 44
	// TokenTypeNone: not a token
	TokenTypeNone TokenType = 256
	// TokenTypeError: an error occurred
	TokenTypeError TokenType = 257
	// TokenTypeChar: a character
	TokenTypeChar TokenType = 258
	// TokenTypeBinary: a binary integer
	TokenTypeBinary TokenType = 259
	// TokenTypeOctal: an octal integer
	TokenTypeOctal TokenType = 260
	// TokenTypeInt: an integer
	TokenTypeInt TokenType = 261
	// TokenTypeHex: a hex integer
	TokenTypeHex TokenType = 262
	// TokenTypeFloat: a floating point number
	TokenTypeFloat TokenType = 263
	// TokenTypeString: a string
	TokenTypeString TokenType = 264
	// TokenTypeSymbol: a symbol
	TokenTypeSymbol TokenType = 265
	// TokenTypeIdentifier: an identifier
	TokenTypeIdentifier TokenType = 266
	// TokenTypeIdentifierNull: a null identifier
	TokenTypeIdentifierNull TokenType = 267
	// TokenTypeCommentSingle: one line comment
	TokenTypeCommentSingle TokenType = 268
	// TokenTypeCommentMulti: multi line comment
	TokenTypeCommentMulti TokenType = 269
)

// Scanner: the data structure representing a lexical scanner.
//
// You should set @input_name after creating the scanner, since it is used by
// the default message handler when displaying warnings and errors. If you are
// scanning a file, the filename would be a good choice.
//
// The @user_data and @max_parse_errors fields are not used. If you need to
// associate extra data with the scanner you can place them here.
//
// If you want to use your own message handler you can set the @msg_handler
// field. The type of the message handler function is declared by MsgFunc.
type Scanner struct {
	native C.GScanner
}

// WrapScanner wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScanner(ptr unsafe.Pointer) *Scanner {
	if ptr == nil {
		return nil
	}

	return (*Scanner)(ptr)
}

func marshalScanner(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScanner(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *Scanner) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// UserData gets the field inside the struct.
func (s *Scanner) UserData() interface{} {
	var ret interface{}
	ret = box.Get(uintptr(s.native.user_data))
	return ret
}

// MaxParseErrors gets the field inside the struct.
func (s *Scanner) MaxParseErrors() uint {
	var ret uint
	ret = uint(s.native.max_parse_errors)
	return ret
}

// ParseErrors gets the field inside the struct.
func (s *Scanner) ParseErrors() uint {
	var ret uint
	ret = uint(s.native.parse_errors)
	return ret
}

// InputName gets the field inside the struct.
func (s *Scanner) InputName() string {
	var ret string
	ret = C.GoString(s.native.input_name)
	return ret
}

// Qdata gets the field inside the struct.
func (s *Scanner) Qdata() *Data {
	var ret *Data
	{
		ret = WrapData(unsafe.Pointer(s.native.qdata))
	}
	return ret
}

// Config gets the field inside the struct.
func (s *Scanner) Config() *ScannerConfig {
	var ret *ScannerConfig
	{
		ret = WrapScannerConfig(unsafe.Pointer(s.native.config))
	}
	return ret
}

// Token gets the field inside the struct.
func (s *Scanner) Token() TokenType {
	var ret TokenType
	ret = TokenType(s.native.token)
	return ret
}

// Line gets the field inside the struct.
func (s *Scanner) Line() uint {
	var ret uint
	ret = uint(s.native.line)
	return ret
}

// Position gets the field inside the struct.
func (s *Scanner) Position() uint {
	var ret uint
	ret = uint(s.native.position)
	return ret
}

// NextLine gets the field inside the struct.
func (s *Scanner) NextLine() uint {
	var ret uint
	ret = uint(s.native.next_line)
	return ret
}

// NextPosition gets the field inside the struct.
func (s *Scanner) NextPosition() uint {
	var ret uint
	ret = uint(s.native.next_position)
	return ret
}

// CurLine returns the current line in the input stream (counting from 1). This
// is the line of the last token parsed via g_scanner_get_next_token().
func (s *Scanner) CurLine() uint {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_cur_line(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// CurPosition returns the current position in the current line (counting from
// 0). This is the position of the last token parsed via
// g_scanner_get_next_token().
func (s *Scanner) CurPosition() uint {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_cur_position(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// CurToken gets the current token type. This is simply the @token field in the
// #GScanner structure.
func (s *Scanner) CurToken() TokenType {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_cur_token(arg0)

	var ret0 TokenType

	ret0 = TokenType(ret)

	return ret0
}

// Destroy frees all memory used by the #GScanner.
func (s *Scanner) Destroy() {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	C.g_scanner_destroy(arg0)
}

// EOF returns true if the scanner has reached the end of the file or text
// buffer.
func (s *Scanner) EOF() bool {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_eof(arg0)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// NextToken parses the next token just like g_scanner_peek_next_token() and
// also removes it from the input stream. The token data is placed in the
// @token, @value, @line, and @position fields of the #GScanner structure.
func (s *Scanner) NextToken() TokenType {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_get_next_token(arg0)

	var ret0 TokenType

	ret0 = TokenType(ret)

	return ret0
}

// InputFile prepares to scan a file.
func (s *Scanner) InputFile(inputFd int) {
	var arg0 *C.GScanner
	var arg1 C.gint

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.gint(inputFd)

	C.g_scanner_input_file(arg0, arg1)
}

// InputText prepares to scan a text buffer.
func (s *Scanner) InputText(text string, textLen uint) {
	var arg0 *C.GScanner
	var arg1 *C.gchar
	var arg2 C.guint

	arg0 = (*C.GScanner)(s.Native())
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.guint(textLen)

	C.g_scanner_input_text(arg0, arg1, arg2)
}

// LookupSymbol looks up a symbol in the current scope and return its value. If
// the symbol is not bound in the current scope, nil is returned.
func (s *Scanner) LookupSymbol(symbol string) interface{} {
	var arg0 *C.GScanner
	var arg1 *C.gchar

	arg0 = (*C.GScanner)(s.Native())
	arg1 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_scanner_lookup_symbol(arg0, arg1)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

// PeekNextToken parses the next token, without removing it from the input
// stream. The token data is placed in the @next_token, @next_value, @next_line,
// and @next_position fields of the #GScanner structure.
//
// Note that, while the token is not removed from the input stream (i.e. the
// next call to g_scanner_get_next_token() will return the same token), it will
// not be reevaluated. This can lead to surprising results when changing scope
// or the scanner configuration after peeking the next token. Getting the next
// token after switching the scope or configuration will return whatever was
// peeked before, regardless of any symbols that may have been added or removed
// in the new scope.
func (s *Scanner) PeekNextToken() TokenType {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	ret := C.g_scanner_peek_next_token(arg0)

	var ret0 TokenType

	ret0 = TokenType(ret)

	return ret0
}

// ScopeAddSymbol adds a symbol to the given scope.
func (s *Scanner) ScopeAddSymbol(scopeID uint, symbol string, value interface{}) {
	var arg0 *C.GScanner
	var arg1 C.guint
	var arg2 *C.gchar
	var arg3 C.gpointer

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.guint(scopeID)
	arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gpointer(box.Assign(value))

	C.g_scanner_scope_add_symbol(arg0, arg1, arg2, arg3)
}

// ScopeForeachSymbol calls the given function for each of the symbol/value
// pairs in the given scope of the #GScanner. The function is passed the symbol
// and value of each pair, and the given @user_data parameter.
func (s *Scanner) ScopeForeachSymbol(scopeID uint, fn HFunc) {
	var arg0 *C.GScanner
	var arg1 C.guint
	var arg2 C.GHFunc
	var arg3 C.gpointer

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.guint(scopeID)
	arg2 = (*[0]byte)(C.gotk4_HFunc)
	arg3 = C.gpointer(box.Assign(fn))

	C.g_scanner_scope_foreach_symbol(arg0, arg1, arg2, arg3)
}

// ScopeLookupSymbol looks up a symbol in a scope and return its value. If the
// symbol is not bound in the scope, nil is returned.
func (s *Scanner) ScopeLookupSymbol(scopeID uint, symbol string) interface{} {
	var arg0 *C.GScanner
	var arg1 C.guint
	var arg2 *C.gchar

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.guint(scopeID)
	arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_scanner_scope_lookup_symbol(arg0, arg1, arg2)

	var ret0 interface{}

	ret0 = box.Get(uintptr(ret)).(interface{})

	return ret0
}

// ScopeRemoveSymbol removes a symbol from a scope.
func (s *Scanner) ScopeRemoveSymbol(scopeID uint, symbol string) {
	var arg0 *C.GScanner
	var arg1 C.guint
	var arg2 *C.gchar

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.guint(scopeID)
	arg2 = (*C.gchar)(C.CString(symbol))
	defer C.free(unsafe.Pointer(arg2))

	C.g_scanner_scope_remove_symbol(arg0, arg1, arg2)
}

// SetScope sets the current scope.
func (s *Scanner) SetScope(scopeID uint) uint {
	var arg0 *C.GScanner
	var arg1 C.guint

	arg0 = (*C.GScanner)(s.Native())
	arg1 = C.guint(scopeID)

	ret := C.g_scanner_set_scope(arg0, arg1)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// SyncFileOffset rewinds the filedescriptor to the current buffer position and
// blows the file read ahead buffer. This is useful for third party uses of the
// scanners filedescriptor, which hooks onto the current scanning position.
func (s *Scanner) SyncFileOffset() {
	var arg0 *C.GScanner

	arg0 = (*C.GScanner)(s.Native())

	C.g_scanner_sync_file_offset(arg0)
}

// UnexpToken outputs a message through the scanner's msg_handler, resulting
// from an unexpected token in the input stream. Note that you should not call
// g_scanner_peek_next_token() followed by g_scanner_unexp_token() without an
// intermediate call to g_scanner_get_next_token(), as g_scanner_unexp_token()
// evaluates the scanner's current token (not the peeked token) to construct
// part of the message.
func (s *Scanner) UnexpToken(expectedToken TokenType, identifierSpec string, symbolSpec string, symbolName string, message string, isError int) {
	var arg0 *C.GScanner
	var arg1 C.GTokenType
	var arg2 *C.gchar
	var arg3 *C.gchar
	var arg4 *C.gchar
	var arg5 *C.gchar
	var arg6 C.gint

	arg0 = (*C.GScanner)(s.Native())
	arg1 = (C.GTokenType)(expectedToken)
	arg2 = (*C.gchar)(C.CString(identifierSpec))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(symbolSpec))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(symbolName))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.gchar)(C.CString(message))
	defer C.free(unsafe.Pointer(arg5))
	arg6 = C.gint(isError)

	C.g_scanner_unexp_token(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ScannerConfig specifies the #GScanner parser configuration. Most settings can
// be changed during the parsing phase and will affect the lexical parsing of
// the next unpeeked token.
type ScannerConfig struct {
	native C.GScannerConfig
}

// WrapScannerConfig wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScannerConfig(ptr unsafe.Pointer) *ScannerConfig {
	if ptr == nil {
		return nil
	}

	return (*ScannerConfig)(ptr)
}

func marshalScannerConfig(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScannerConfig(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *ScannerConfig) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// CsetSkipCharacters gets the field inside the struct.
func (s *ScannerConfig) CsetSkipCharacters() string {
	var ret string
	ret = C.GoString(s.native.cset_skip_characters)
	return ret
}

// CsetIdentifierFirst gets the field inside the struct.
func (s *ScannerConfig) CsetIdentifierFirst() string {
	var ret string
	ret = C.GoString(s.native.cset_identifier_first)
	return ret
}

// CsetIdentifierNth gets the field inside the struct.
func (s *ScannerConfig) CsetIdentifierNth() string {
	var ret string
	ret = C.GoString(s.native.cset_identifier_nth)
	return ret
}

// CpairCommentSingle gets the field inside the struct.
func (s *ScannerConfig) CpairCommentSingle() string {
	var ret string
	ret = C.GoString(s.native.cpair_comment_single)
	return ret
}
