// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_gstring_get_type()), F: marshalString},
	})
}

// NewString creates a new #GString, initialized with the given string.
func NewString(init string) *String {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(init))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_string_new(arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *String) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// StringNewLen creates a new #GString with @len bytes of the @init buffer.
// Because a length is provided, @init need not be nul-terminated, and can
// contain embedded nul bytes.
//
// Since this function does not stop at nul bytes, it is the caller's
// responsibility to ensure that @init has at least @len addressable bytes.
func StringNewLen(init string, len int) *String {
	var arg1 *C.gchar
	var arg2 C.gssize

	arg1 = (*C.gchar)(C.CString(init))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	ret := C.g_string_new_len(arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *String) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// NewStringSized creates a new #GString, with enough space for @dfl_size bytes.
// This is useful if you are going to add a lot of text to the string and don't
// want it to be reallocated too often.
func NewStringSized(dflSize uint) *String {
	var arg1 C.gsize

	arg1 = C.gsize(dflSize)

	ret := C.g_string_sized_new(arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *String) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// String: the GString struct contains the public fields of a GString.
type String struct {
	native C.GString
}

// WrapString wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapString(ptr unsafe.Pointer) *String {
	if ptr == nil {
		return nil
	}

	return (*String)(ptr)
}

func marshalString(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapString(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *String) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Str gets the field inside the struct.
func (s *String) Str() string {
	var ret string
	ret = C.GoString(s.native.str)
	return ret
}

// Len gets the field inside the struct.
func (s *String) Len() uint {
	var ret uint
	ret = uint(s.native.len)
	return ret
}

// AllocatedLen gets the field inside the struct.
func (s *String) AllocatedLen() uint {
	var ret uint
	ret = uint(s.native.allocated_len)
	return ret
}

// Append adds a string onto the end of a #GString, expanding it if necessary.
func (s *String) Append(val string) *String {
	var arg0 *C.GString
	var arg1 *C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_string_append(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// AppendC adds a byte onto the end of a #GString, expanding it if necessary.
func (s *String) AppendC(c byte) *String {
	var arg0 *C.GString
	var arg1 C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gchar(c)

	ret := C.g_string_append_c(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// AppendLen appends @len bytes of @val to @string.
//
// If @len is positive, @val may contain embedded nuls and need not be
// nul-terminated. It is the caller's responsibility to ensure that @val has at
// least @len addressable bytes.
//
// If @len is negative, @val must be nul-terminated and @len is considered to
// request the entire string length. This makes g_string_append_len() equivalent
// to g_string_append().
func (s *String) AppendLen(val string, len int) *String {
	var arg0 *C.GString
	var arg1 *C.gchar
	var arg2 C.gssize

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	ret := C.g_string_append_len(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// AppendUnichar converts a Unicode character into UTF-8, and appends it to the
// string.
func (s *String) AppendUnichar(wc uint32) *String {
	var arg0 *C.GString
	var arg1 C.gunichar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gunichar(wc)

	ret := C.g_string_append_unichar(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// AppendURIEscaped appends @unescaped to @string, escaping any characters that
// are reserved in URIs using URI-style escape sequences.
func (s *String) AppendURIEscaped(unescaped string, reservedCharsAllowed string, allowUTF8 bool) *String {
	var arg0 *C.GString
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 C.gboolean

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(unescaped))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(reservedCharsAllowed))
	defer C.free(unsafe.Pointer(arg2))
	if allowUTF8 {
		arg3 = C.TRUE
	}

	ret := C.g_string_append_uri_escaped(arg0, arg1, arg2, arg3)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// ASCIIDown converts all uppercase ASCII letters to lowercase ASCII letters.
func (s *String) ASCIIDown() *String {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_ascii_down(arg0)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// ASCIIUp converts all lowercase ASCII letters to uppercase ASCII letters.
func (s *String) ASCIIUp() *String {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_ascii_up(arg0)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Assign copies the bytes from a string into a #GString, destroying any
// previous contents. It is rather like the standard strcpy() function, except
// that you do not have to worry about having enough space to copy the string.
func (s *String) Assign(rval string) *String {
	var arg0 *C.GString
	var arg1 *C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(rval))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_string_assign(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Down converts a #GString to lowercase.
func (s *String) Down() *String {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_down(arg0)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Equal compares two strings for equality, returning true if they are equal.
// For use with Table.
func (v *String) Equal(v2 *String) bool {
	var arg0 *C.GString
	var arg1 *C.GString

	arg0 = (*C.GString)(v.Native())
	arg1 = (*C.GString)(v2.Native())

	ret := C.g_string_equal(arg0, arg1)

	var ret0 bool

	ret0 = C.bool(ret) != 0

	return ret0
}

// Erase removes @len bytes from a #GString, starting at position @pos. The rest
// of the #GString is shifted down to fill the gap.
func (s *String) Erase(pos int, len int) *String {
	var arg0 *C.GString
	var arg1 C.gssize
	var arg2 C.gssize

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gssize(pos)
	arg2 = C.gssize(len)

	ret := C.g_string_erase(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Free frees the memory allocated for the #GString. If @free_segment is true it
// also frees the character data. If it's false, the caller gains ownership of
// the buffer and must free it after use with g_free().
func (s *String) Free(freeSegment bool) string {
	var arg0 *C.GString
	var arg1 C.gboolean

	arg0 = (*C.GString)(s.Native())
	if freeSegment {
		arg1 = C.TRUE
	}

	ret := C.g_string_free(arg0, arg1)

	var ret0 string

	ret0 = C.GoString(ret)
	C.free(unsafe.Pointer(ret))

	return ret0
}

// FreeToBytes transfers ownership of the contents of @string to a newly
// allocated #GBytes. The #GString structure itself is deallocated, and it is
// therefore invalid to use @string after invoking this function.
//
// Note that while #GString ensures that its buffer always has a trailing nul
// character (not reflected in its "len"), the returned #GBytes does not include
// this extra nul; i.e. it has length exactly equal to the "len" member.
func (s *String) FreeToBytes() *Bytes {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_free_to_bytes(arg0)

	var ret0 *Bytes

	{
		ret0 = WrapBytes(unsafe.Pointer(ret))
		runtime.SetFinalizer(ret0, func(v *Bytes) {
			C.free(unsafe.Pointer(v.Native()))
		})
	}

	return ret0
}

// Hash creates a hash code for @str; for use with Table.
func (s *String) Hash() uint {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_hash(arg0)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// Insert inserts a copy of a string into a #GString, expanding it if necessary.
func (s *String) Insert(pos int, val string) *String {
	var arg0 *C.GString
	var arg1 C.gssize
	var arg2 *C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gssize(pos)
	arg2 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_string_insert(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// InsertC inserts a byte into a #GString, expanding it if necessary.
func (s *String) InsertC(pos int, c byte) *String {
	var arg0 *C.GString
	var arg1 C.gssize
	var arg2 C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gssize(pos)
	arg2 = C.gchar(c)

	ret := C.g_string_insert_c(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// InsertLen inserts @len bytes of @val into @string at @pos.
//
// If @len is positive, @val may contain embedded nuls and need not be
// nul-terminated. It is the caller's responsibility to ensure that @val has at
// least @len addressable bytes.
//
// If @len is negative, @val must be nul-terminated and @len is considered to
// request the entire string length.
//
// If @pos is -1, bytes are inserted at the end of the string.
func (s *String) InsertLen(pos int, val string, len int) *String {
	var arg0 *C.GString
	var arg1 C.gssize
	var arg2 *C.gchar
	var arg3 C.gssize

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gssize(pos)
	arg2 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gssize(len)

	ret := C.g_string_insert_len(arg0, arg1, arg2, arg3)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// InsertUnichar converts a Unicode character into UTF-8, and insert it into the
// string at the given position.
func (s *String) InsertUnichar(pos int, wc uint32) *String {
	var arg0 *C.GString
	var arg1 C.gssize
	var arg2 C.gunichar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gssize(pos)
	arg2 = C.gunichar(wc)

	ret := C.g_string_insert_unichar(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Overwrite overwrites part of a string, lengthening it if necessary.
func (s *String) Overwrite(pos uint, val string) *String {
	var arg0 *C.GString
	var arg1 C.gsize
	var arg2 *C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gsize(pos)
	arg2 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg2))

	ret := C.g_string_overwrite(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// OverwriteLen overwrites part of a string, lengthening it if necessary. This
// function will work with embedded nuls.
func (s *String) OverwriteLen(pos uint, val string, len int) *String {
	var arg0 *C.GString
	var arg1 C.gsize
	var arg2 *C.gchar
	var arg3 C.gssize

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gsize(pos)
	arg2 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gssize(len)

	ret := C.g_string_overwrite_len(arg0, arg1, arg2, arg3)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Prepend adds a string on to the start of a #GString, expanding it if
// necessary.
func (s *String) Prepend(val string) *String {
	var arg0 *C.GString
	var arg1 *C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg1))

	ret := C.g_string_prepend(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// PrependC adds a byte onto the start of a #GString, expanding it if necessary.
func (s *String) PrependC(c byte) *String {
	var arg0 *C.GString
	var arg1 C.gchar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gchar(c)

	ret := C.g_string_prepend_c(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// PrependLen prepends @len bytes of @val to @string.
//
// If @len is positive, @val may contain embedded nuls and need not be
// nul-terminated. It is the caller's responsibility to ensure that @val has at
// least @len addressable bytes.
//
// If @len is negative, @val must be nul-terminated and @len is considered to
// request the entire string length. This makes g_string_prepend_len()
// equivalent to g_string_prepend().
func (s *String) PrependLen(val string, len int) *String {
	var arg0 *C.GString
	var arg1 *C.gchar
	var arg2 C.gssize

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(val))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(len)

	ret := C.g_string_prepend_len(arg0, arg1, arg2)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// PrependUnichar converts a Unicode character into UTF-8, and prepends it to
// the string.
func (s *String) PrependUnichar(wc uint32) *String {
	var arg0 *C.GString
	var arg1 C.gunichar

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gunichar(wc)

	ret := C.g_string_prepend_unichar(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Replace replaces the string @find with the string @replace in a #GString up
// to @limit times. If the number of instances of @find in the #GString is less
// than @limit, all instances are replaced. If the number of instances is `0`,
// all instances of @find are replaced.
func (s *String) Replace(find string, replace string, limit uint) uint {
	var arg0 *C.GString
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 C.guint

	arg0 = (*C.GString)(s.Native())
	arg1 = (*C.gchar)(C.CString(find))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(replace))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.guint(limit)

	ret := C.g_string_replace(arg0, arg1, arg2, arg3)

	var ret0 uint

	ret0 = uint(ret)

	return ret0
}

// SetSize sets the length of a #GString. If the length is less than the current
// length, the string will be truncated. If the length is greater than the
// current length, the contents of the newly added area are undefined. (However,
// as always, string->str[string->len] will be a nul byte.)
func (s *String) SetSize(len uint) *String {
	var arg0 *C.GString
	var arg1 C.gsize

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gsize(len)

	ret := C.g_string_set_size(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Truncate cuts off the end of the GString, leaving the first @len bytes.
func (s *String) Truncate(len uint) *String {
	var arg0 *C.GString
	var arg1 C.gsize

	arg0 = (*C.GString)(s.Native())
	arg1 = C.gsize(len)

	ret := C.g_string_truncate(arg0, arg1)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}

// Up converts a #GString to uppercase.
func (s *String) Up() *String {
	var arg0 *C.GString

	arg0 = (*C.GString)(s.Native())

	ret := C.g_string_up(arg0)

	var ret0 *String

	{
		ret0 = WrapString(unsafe.Pointer(ret))
	}

	return ret0
}
