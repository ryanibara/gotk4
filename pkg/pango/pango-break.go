// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// DefaultBreak: this is the default break algorithm.
//
// It applies Unicode rules without language-specific tailoring, therefore the
// analyis argument is unused and can be NULL.
//
// See pango_tailor_break() for language-specific breaks.
//
// The function takes the following parameters:
//
//    - text to break. Must be valid UTF-8.
//    - length of text in bytes (may be -1 if text is nul-terminated).
//    - analysis (optional) for the text.
//    - attrs: logical attributes to fill in.
//    - attrsLen: size of the array passed as attrs.
//
func DefaultBreak(text string, length int32, analysis *Analysis, attrs *LogAttr, attrsLen int32) {
	var _args [5]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[0]))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(length)
	if analysis != nil {
		*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(analysis)))
	}
	*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(gextras.StructNative(unsafe.Pointer(attrs)))
	*(*C.int)(unsafe.Pointer(&_args[4])) = C.int(attrsLen)

	girepository.MustFind("Pango", "default_break").Invoke(_args[:], nil)

	runtime.KeepAlive(text)
	runtime.KeepAlive(length)
	runtime.KeepAlive(analysis)
	runtime.KeepAlive(attrs)
	runtime.KeepAlive(attrsLen)
}

// FindParagraphBoundary locates a paragraph boundary in text.
//
// A boundary is caused by delimiter characters, such as a newline, carriage
// return, carriage return-newline pair, or Unicode paragraph separator
// character. The index of the run of delimiters is returned in
// paragraph_delimiter_index. The index of the start of the paragrap (index
// after all delimiters) is stored in next_paragraph_start.
//
// If no delimiters are found, both paragraph_delimiter_index and
// next_paragraph_start are filled with the length of text (an index one off the
// end).
//
// The function takes the following parameters:
//
//    - text: UTF-8 text.
//    - length of text in bytes, or -1 if nul-terminated.
//
// The function returns the following values:
//
//    - paragraphDelimiterIndex: return location for index of delimiter.
//    - nextParagraphStart: return location for start of next paragraph.
//
func FindParagraphBoundary(text string, length int32) (paragraphDelimiterIndex, nextParagraphStart int32) {
	var _args [2]girepository.Argument
	var _outs [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_args[0]))
	*(*C.gint)(unsafe.Pointer(&_args[1])) = C.gint(length)

	girepository.MustFind("Pango", "find_paragraph_boundary").Invoke(_args[:], _outs[:])

	runtime.KeepAlive(text)
	runtime.KeepAlive(length)

	var _paragraphDelimiterIndex int32 // out
	var _nextParagraphStart int32      // out

	_paragraphDelimiterIndex = *(*int32)(unsafe.Pointer(_outs[0]))
	_nextParagraphStart = *(*int32)(unsafe.Pointer(_outs[1]))

	return _paragraphDelimiterIndex, _nextParagraphStart
}

// LogAttr: PangoLogAttr structure stores information about the attributes of a
// single character.
//
// An instance of this type is always passed by reference.
type LogAttr struct {
	*logAttr
}

// logAttr is the struct that's finalized.
type logAttr struct {
	native unsafe.Pointer
}
