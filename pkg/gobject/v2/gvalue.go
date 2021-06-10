// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_value_get_type()), F: marshalValue},
	})
}

// Value: an opaque structure used to hold different types of values. The data
// within the structure has protected scope: it is accessible only to functions
// within a ValueTable structure, or implementations of the g_value_*() API.
// That is, code portions which implement new fundamental types. #GValue users
// cannot make any assumptions about how data is stored within the 2 element
// @data union, and the @g_type member should only be accessed through the
// G_VALUE_TYPE() macro.
type Value struct {
	native C.GValue
}

// WrapValue wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapValue(ptr unsafe.Pointer) *Value {
	if ptr == nil {
		return nil
	}

	return (*Value)(ptr)
}

func marshalValue(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapValue(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (v *Value) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// Copy copies the value of @src_value into @dest_value.
func (s *Value) Copy(destValue **externglib.Value) {
	var _arg0 *C.GValue
	var _arg1 *C.GValue

	_arg0 = (*C.GValue)(s.GValue)
	_arg1 = (*C.GValue)(destValue.GValue)

	C.g_value_copy(_arg0, _arg1)
}

// DupBoxed: get the contents of a G_TYPE_BOXED derived #GValue. Upon getting,
// the boxed value is duplicated and needs to be later freed with
// g_boxed_free(), e.g. like: g_boxed_free (G_VALUE_TYPE (@value),
// return_value);
func (v *Value) DupBoxed() interface{} {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_dup_boxed(_arg0)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// DupObject: get the contents of a G_TYPE_OBJECT derived #GValue, increasing
// its reference count. If the contents of the #GValue are nil, then nil will be
// returned.
func (v *Value) DupObject() gextras.Objector {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_dup_object(_arg0)

	var _object gextras.Objector

	_object = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(gextras.Objector)

	return _object
}

// DupString: get a copy the contents of a G_TYPE_STRING #GValue.
func (v *Value) DupString() string {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret *C.gchar

	_cret = C.g_value_dup_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FitsPointer determines if @value will fit inside the size of a pointer value.
// This is an internal function introduced mainly for C marshallers.
func (v *Value) FitsPointer() bool {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gboolean

	_cret = C.g_value_fits_pointer(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Boolean: get the contents of a G_TYPE_BOOLEAN #GValue.
func (v *Value) Boolean() bool {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gboolean

	_cret = C.g_value_get_boolean(_arg0)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Boxed: get the contents of a G_TYPE_BOXED derived #GValue.
func (v *Value) Boxed() interface{} {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_get_boxed(_arg0)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// Char: do not use this function; it is broken on platforms where the char type
// is unsigned, such as ARM and PowerPC. See g_value_get_schar().
//
// Get the contents of a G_TYPE_CHAR #GValue.
func (v *Value) Char() byte {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gchar

	_cret = C.g_value_get_char(_arg0)

	var _gchar byte

	_gchar = (byte)(_cret)

	return _gchar
}

// Double: get the contents of a G_TYPE_DOUBLE #GValue.
func (v *Value) Double() float64 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gdouble

	_cret = C.g_value_get_double(_arg0)

	var _gdouble float64

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Enum: get the contents of a G_TYPE_ENUM #GValue.
func (v *Value) Enum() int {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gint

	_cret = C.g_value_get_enum(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Flags: get the contents of a G_TYPE_FLAGS #GValue.
func (v *Value) Flags() uint {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.guint

	_cret = C.g_value_get_flags(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// Float: get the contents of a G_TYPE_FLOAT #GValue.
func (v *Value) Float() float32 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gfloat

	_cret = C.g_value_get_float(_arg0)

	var _gfloat float32

	_gfloat = (float32)(_cret)

	return _gfloat
}

// GType: get the contents of a G_TYPE_GTYPE #GValue.
func (v *Value) GType() externglib.Type {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.GType

	_cret = C.g_value_get_gtype(_arg0)

	var _gType externglib.Type

	_gType = externglib.Type(_cret)

	return _gType
}

// Int: get the contents of a G_TYPE_INT #GValue.
func (v *Value) Int() int {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gint

	_cret = C.g_value_get_int(_arg0)

	var _gint int

	_gint = (int)(_cret)

	return _gint
}

// Int64: get the contents of a G_TYPE_INT64 #GValue.
func (v *Value) Int64() int64 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gint64

	_cret = C.g_value_get_int64(_arg0)

	var _gint64 int64

	_gint64 = (int64)(_cret)

	return _gint64
}

// Long: get the contents of a G_TYPE_LONG #GValue.
func (v *Value) Long() int32 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.glong

	_cret = C.g_value_get_long(_arg0)

	var _glong int32

	_glong = (int32)(_cret)

	return _glong
}

// Object: get the contents of a G_TYPE_OBJECT derived #GValue.
func (v *Value) Object() gextras.Objector {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_get_object(_arg0)

	var _object gextras.Objector

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gextras.Objector)

	return _object
}

// Pointer: get the contents of a pointer #GValue.
func (v *Value) Pointer() interface{} {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_get_pointer(_arg0)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// Schar: get the contents of a G_TYPE_CHAR #GValue.
func (v *Value) Schar() int8 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gint8

	_cret = C.g_value_get_schar(_arg0)

	var _gint8 int8

	_gint8 = (int8)(_cret)

	return _gint8
}

// String: get the contents of a G_TYPE_STRING #GValue.
func (v *Value) String() string {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret *C.gchar

	_cret = C.g_value_get_string(_arg0)

	var _utf8 string

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Uchar: get the contents of a G_TYPE_UCHAR #GValue.
func (v *Value) Uchar() byte {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.guchar

	_cret = C.g_value_get_uchar(_arg0)

	var _guint8 byte

	_guint8 = (byte)(_cret)

	return _guint8
}

// Uint: get the contents of a G_TYPE_UINT #GValue.
func (v *Value) Uint() uint {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.guint

	_cret = C.g_value_get_uint(_arg0)

	var _guint uint

	_guint = (uint)(_cret)

	return _guint
}

// Uint64: get the contents of a G_TYPE_UINT64 #GValue.
func (v *Value) Uint64() uint64 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.guint64

	_cret = C.g_value_get_uint64(_arg0)

	var _guint64 uint64

	_guint64 = (uint64)(_cret)

	return _guint64
}

// Ulong: get the contents of a G_TYPE_ULONG #GValue.
func (v *Value) Ulong() uint32 {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gulong

	_cret = C.g_value_get_ulong(_arg0)

	var _gulong uint32

	_gulong = (uint32)(_cret)

	return _gulong
}

// Init initializes @value with the default value of @type.
func (v *Value) Init(gType externglib.Type) **externglib.Value {
	var _arg0 *C.GValue
	var _arg1 C.GType

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.GType(gType)

	var _cret *C.GValue

	_cret = C.g_value_init(_arg0, _arg1)

	var _ret **externglib.Value

	_ret = externglib.ValueFromNative(unsafe.Pointer(_cret))

	return _ret
}

// InitFromInstance initializes and sets @value from an instantiatable type via
// the value_table's collect_value() function.
//
// Note: The @value will be initialised with the exact type of @instance. If you
// wish to set the @value's type to a different GType (such as a parent class
// GType), you need to manually call g_value_init() and g_value_set_instance().
func (v *Value) InitFromInstance(instance TypeInstance) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (C.gpointer)(unsafe.Pointer(instance.Native()))

	C.g_value_init_from_instance(_arg0, _arg1)
}

// PeekPointer returns the value contents as pointer. This function asserts that
// g_value_fits_pointer() returned true for the passed in value. This is an
// internal function introduced mainly for C marshallers.
func (v *Value) PeekPointer() interface{} {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret C.gpointer

	_cret = C.g_value_peek_pointer(_arg0)

	var _gpointer interface{}

	_gpointer = (interface{})(_cret)

	return _gpointer
}

// Reset clears the current value in @value and resets it to the default value
// (as if the value had just been initialized).
func (v *Value) Reset() **externglib.Value {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	var _cret *C.GValue

	_cret = C.g_value_reset(_arg0)

	var _ret **externglib.Value

	_ret = externglib.ValueFromNative(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *externglib.Value) {
		C.g_value_unset((*C.GValue)(v.GValue))
	})

	return _ret
}

// SetBoolean: set the contents of a G_TYPE_BOOLEAN #GValue to @v_boolean.
func (v *Value) SetBoolean(vBoolean bool) {
	var _arg0 *C.GValue
	var _arg1 C.gboolean

	_arg0 = (*C.GValue)(v.GValue)
	if vBoolean {
		_arg1 = C.gboolean(1)
	}

	C.g_value_set_boolean(_arg0, _arg1)
}

// SetBoxed: set the contents of a G_TYPE_BOXED derived #GValue to @v_boxed.
func (v *Value) SetBoxed(vBoxed interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vBoxed)

	C.g_value_set_boxed(_arg0, _arg1)
}

// SetBoxedTakeOwnership: this is an internal function introduced mainly for C
// marshallers.
func (v *Value) SetBoxedTakeOwnership(vBoxed interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vBoxed)

	C.g_value_set_boxed_take_ownership(_arg0, _arg1)
}

// SetChar: set the contents of a G_TYPE_CHAR #GValue to @v_char.
func (v *Value) SetChar(vChar byte) {
	var _arg0 *C.GValue
	var _arg1 C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gchar(vChar)

	C.g_value_set_char(_arg0, _arg1)
}

// SetDouble: set the contents of a G_TYPE_DOUBLE #GValue to @v_double.
func (v *Value) SetDouble(vDouble float64) {
	var _arg0 *C.GValue
	var _arg1 C.gdouble

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gdouble(vDouble)

	C.g_value_set_double(_arg0, _arg1)
}

// SetEnum: set the contents of a G_TYPE_ENUM #GValue to @v_enum.
func (v *Value) SetEnum(vEnum int) {
	var _arg0 *C.GValue
	var _arg1 C.gint

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gint(vEnum)

	C.g_value_set_enum(_arg0, _arg1)
}

// SetFlags: set the contents of a G_TYPE_FLAGS #GValue to @v_flags.
func (v *Value) SetFlags(vFlags uint) {
	var _arg0 *C.GValue
	var _arg1 C.guint

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.guint(vFlags)

	C.g_value_set_flags(_arg0, _arg1)
}

// SetFloat: set the contents of a G_TYPE_FLOAT #GValue to @v_float.
func (v *Value) SetFloat(vFloat float32) {
	var _arg0 *C.GValue
	var _arg1 C.gfloat

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gfloat(vFloat)

	C.g_value_set_float(_arg0, _arg1)
}

// SetGType: set the contents of a G_TYPE_GTYPE #GValue to @v_gtype.
func (v *Value) SetGType(vGtype externglib.Type) {
	var _arg0 *C.GValue
	var _arg1 C.GType

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.GType(vGtype)

	C.g_value_set_gtype(_arg0, _arg1)
}

// SetInstance sets @value from an instantiatable type via the value_table's
// collect_value() function.
func (v *Value) SetInstance(instance interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(instance)

	C.g_value_set_instance(_arg0, _arg1)
}

// SetInt: set the contents of a G_TYPE_INT #GValue to @v_int.
func (v *Value) SetInt(vInt int) {
	var _arg0 *C.GValue
	var _arg1 C.gint

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gint(vInt)

	C.g_value_set_int(_arg0, _arg1)
}

// SetInt64: set the contents of a G_TYPE_INT64 #GValue to @v_int64.
func (v *Value) SetInt64(vInt64 int64) {
	var _arg0 *C.GValue
	var _arg1 C.gint64

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gint64(vInt64)

	C.g_value_set_int64(_arg0, _arg1)
}

// SetInternedString: set the contents of a G_TYPE_STRING #GValue to @v_string.
// The string is assumed to be static and interned (canonical, for example from
// g_intern_string()), and is thus not duplicated when setting the #GValue.
func (v *Value) SetInternedString(vString string) {
	var _arg0 *C.GValue
	var _arg1 *C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_value_set_interned_string(_arg0, _arg1)
}

// SetLong: set the contents of a G_TYPE_LONG #GValue to @v_long.
func (v *Value) SetLong(vLong int32) {
	var _arg0 *C.GValue
	var _arg1 C.glong

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.glong(vLong)

	C.g_value_set_long(_arg0, _arg1)
}

// SetObject: set the contents of a G_TYPE_OBJECT derived #GValue to @v_object.
//
// g_value_set_object() increases the reference count of @v_object (the #GValue
// holds a reference to @v_object). If you do not wish to increase the reference
// count of the object (i.e. you wish to pass your current reference to the
// #GValue because you no longer need it), use g_value_take_object() instead.
//
// It is important that your #GValue holds a reference to @v_object (either its
// own, or one it has taken) to ensure that the object won't be destroyed while
// the #GValue still exists).
func (v *Value) SetObject(vObject gextras.Objector) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GObject)(unsafe.Pointer(vObject.Native()))

	C.g_value_set_object(_arg0, _arg1)
}

// SetObjectTakeOwnership: this is an internal function introduced mainly for C
// marshallers.
func (v *Value) SetObjectTakeOwnership(vObject interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vObject)

	C.g_value_set_object_take_ownership(_arg0, _arg1)
}

// SetParam: set the contents of a G_TYPE_PARAM #GValue to @param.
func (v *Value) SetParam(param ParamSpec) {
	var _arg0 *C.GValue
	var _arg1 *C.GParamSpec

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GParamSpec)(unsafe.Pointer(param.Native()))

	C.g_value_set_param(_arg0, _arg1)
}

// SetParamTakeOwnership: this is an internal function introduced mainly for C
// marshallers.
func (v *Value) SetParamTakeOwnership(param ParamSpec) {
	var _arg0 *C.GValue
	var _arg1 *C.GParamSpec

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GParamSpec)(unsafe.Pointer(param.Native()))

	C.g_value_set_param_take_ownership(_arg0, _arg1)
}

// SetPointer: set the contents of a pointer #GValue to @v_pointer.
func (v *Value) SetPointer(vPointer interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vPointer)

	C.g_value_set_pointer(_arg0, _arg1)
}

// SetSchar: set the contents of a G_TYPE_CHAR #GValue to @v_char.
func (v *Value) SetSchar(vChar int8) {
	var _arg0 *C.GValue
	var _arg1 C.gint8

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gint8(vChar)

	C.g_value_set_schar(_arg0, _arg1)
}

// SetStaticBoxed: set the contents of a G_TYPE_BOXED derived #GValue to
// @v_boxed. The boxed value is assumed to be static, and is thus not duplicated
// when setting the #GValue.
func (v *Value) SetStaticBoxed(vBoxed interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vBoxed)

	C.g_value_set_static_boxed(_arg0, _arg1)
}

// SetStaticString: set the contents of a G_TYPE_STRING #GValue to @v_string.
// The string is assumed to be static, and is thus not duplicated when setting
// the #GValue.
//
// If the the string is a canonical string, using g_value_set_interned_string()
// is more appropriate.
func (v *Value) SetStaticString(vString string) {
	var _arg0 *C.GValue
	var _arg1 *C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_value_set_static_string(_arg0, _arg1)
}

// SetString: set the contents of a G_TYPE_STRING #GValue to @v_string.
func (v *Value) SetString(vString string) {
	var _arg0 *C.GValue
	var _arg1 *C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_value_set_string(_arg0, _arg1)
}

// SetStringTakeOwnership: this is an internal function introduced mainly for C
// marshallers.
func (v *Value) SetStringTakeOwnership(vString string) {
	var _arg0 *C.GValue
	var _arg1 *C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_value_set_string_take_ownership(_arg0, _arg1)
}

// SetUchar: set the contents of a G_TYPE_UCHAR #GValue to @v_uchar.
func (v *Value) SetUchar(vUchar byte) {
	var _arg0 *C.GValue
	var _arg1 C.guchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.guchar(vUchar)

	C.g_value_set_uchar(_arg0, _arg1)
}

// SetUint: set the contents of a G_TYPE_UINT #GValue to @v_uint.
func (v *Value) SetUint(vUint uint) {
	var _arg0 *C.GValue
	var _arg1 C.guint

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.guint(vUint)

	C.g_value_set_uint(_arg0, _arg1)
}

// SetUint64: set the contents of a G_TYPE_UINT64 #GValue to @v_uint64.
func (v *Value) SetUint64(vUint64 uint64) {
	var _arg0 *C.GValue
	var _arg1 C.guint64

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.guint64(vUint64)

	C.g_value_set_uint64(_arg0, _arg1)
}

// SetUlong: set the contents of a G_TYPE_ULONG #GValue to @v_ulong.
func (v *Value) SetUlong(vUlong uint32) {
	var _arg0 *C.GValue
	var _arg1 C.gulong

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gulong(vUlong)

	C.g_value_set_ulong(_arg0, _arg1)
}

// SetVariant: set the contents of a variant #GValue to @variant. If the variant
// is floating, it is consumed.
func (v *Value) SetVariant(variant *glib.Variant) {
	var _arg0 *C.GValue
	var _arg1 *C.GVariant

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	C.g_value_set_variant(_arg0, _arg1)
}

// TakeBoxed sets the contents of a G_TYPE_BOXED derived #GValue to @v_boxed and
// takes over the ownership of the caller’s reference to @v_boxed; the caller
// doesn’t have to unref it any more.
func (v *Value) TakeBoxed(vBoxed interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vBoxed)

	C.g_value_take_boxed(_arg0, _arg1)
}

// TakeObject sets the contents of a G_TYPE_OBJECT derived #GValue to @v_object
// and takes over the ownership of the caller’s reference to @v_object; the
// caller doesn’t have to unref it any more (i.e. the reference count of the
// object is not increased).
//
// If you want the #GValue to hold its own reference to @v_object, use
// g_value_set_object() instead.
func (v *Value) TakeObject(vObject interface{}) {
	var _arg0 *C.GValue
	var _arg1 C.gpointer

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = C.gpointer(vObject)

	C.g_value_take_object(_arg0, _arg1)
}

// TakeParam sets the contents of a G_TYPE_PARAM #GValue to @param and takes
// over the ownership of the caller’s reference to @param; the caller doesn’t
// have to unref it any more.
func (v *Value) TakeParam(param ParamSpec) {
	var _arg0 *C.GValue
	var _arg1 *C.GParamSpec

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GParamSpec)(unsafe.Pointer(param.Native()))

	C.g_value_take_param(_arg0, _arg1)
}

// TakeString sets the contents of a G_TYPE_STRING #GValue to @v_string.
func (v *Value) TakeString(vString string) {
	var _arg0 *C.GValue
	var _arg1 *C.gchar

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_value_take_string(_arg0, _arg1)
}

// TakeVariant: set the contents of a variant #GValue to @variant, and takes
// over the ownership of the caller's reference to @variant; the caller doesn't
// have to unref it any more (i.e. the reference count of the variant is not
// increased).
//
// If @variant was floating then its floating reference is converted to a hard
// reference.
//
// If you want the #GValue to hold its own reference to @variant, use
// g_value_set_variant() instead.
//
// This is an internal function introduced mainly for C marshallers.
func (v *Value) TakeVariant(variant *glib.Variant) {
	var _arg0 *C.GValue
	var _arg1 *C.GVariant

	_arg0 = (*C.GValue)(v.GValue)
	_arg1 = (*C.GVariant)(unsafe.Pointer(variant.Native()))

	C.g_value_take_variant(_arg0, _arg1)
}

// Transform tries to cast the contents of @src_value into a type appropriate to
// store in @dest_value, e.g. to transform a G_TYPE_INT value into a
// G_TYPE_FLOAT value. Performing transformations between value types might
// incur precision lossage. Especially transformations into strings might reveal
// seemingly arbitrary results and shouldn't be relied upon for production code
// (such as rcfile value or object property serialization).
func (s *Value) Transform(destValue **externglib.Value) bool {
	var _arg0 *C.GValue
	var _arg1 *C.GValue

	_arg0 = (*C.GValue)(s.GValue)
	_arg1 = (*C.GValue)(destValue.GValue)

	var _cret C.gboolean

	_cret = C.g_value_transform(_arg0, _arg1)

	var _ok bool

	if _cret {
		_ok = true
	}

	return _ok
}

// Unset clears the current value in @value (if any) and "unsets" the type, this
// releases all resources associated with this GValue. An unset value is the
// same as an uninitialized (zero-filled) #GValue structure.
func (v *Value) Unset() {
	var _arg0 *C.GValue

	_arg0 = (*C.GValue)(v.GValue)

	C.g_value_unset(_arg0)
}
