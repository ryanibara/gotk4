// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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

// Native returns the underlying C source pointer.
func (s *ScannerConfig) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}
