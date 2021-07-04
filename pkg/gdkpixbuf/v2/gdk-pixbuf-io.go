// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_format_get_type()), F: marshalPixbufFormat},
	})
}

// PixbufFormatFlags flags which allow a module to specify further details about
// the supported operations.
type PixbufFormatFlags int

const (
	// PixbufFormatFlagsWritable: the module can write out images in the format.
	PixbufFormatFlagsWritable PixbufFormatFlags = 0b1
	// PixbufFormatFlagsScalable: the image format is scalable
	PixbufFormatFlagsScalable PixbufFormatFlags = 0b10
	// PixbufFormatFlagsThreadsafe: the module is threadsafe. gdk-pixbuf ignores
	// modules that are not marked as threadsafe. (Since 2.28).
	PixbufFormatFlagsThreadsafe PixbufFormatFlags = 0b100
)

// PixbufModulePreparedFunc defines the type of the function that gets called
// once the initial setup of @pixbuf is done.
//
// PixbufLoader uses a function of this type to emit the "<link
// linkend="GdkPixbufLoader-area-prepared">area_prepared</link>" signal.
type PixbufModulePreparedFunc func(pixbuf Pixbuf, anim PixbufAnimation)

//export gotk4_PixbufModulePreparedFunc
func _PixbufModulePreparedFunc(arg0 *C.GdkPixbuf, arg1 *C.GdkPixbufAnimation, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var pixbuf Pixbuf        // out
	var anim PixbufAnimation // out

	pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Pixbuf)
	anim = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(PixbufAnimation)

	fn := v.(PixbufModulePreparedFunc)
	fn(pixbuf, anim)
}

// PixbufModuleSizeFunc defines the type of the function that gets called once
// the size of the loaded image is known.
//
// The function is expected to set @width and @height to the desired size to
// which the image should be scaled. If a module has no efficient way to achieve
// the desired scaling during the loading of the image, it may either ignore the
// size request, or only approximate it - gdk-pixbuf will then perform the
// required scaling on the completely loaded image.
//
// If the function sets @width or @height to zero, the module should interpret
// this as a hint that it will be closed soon and shouldn't allocate further
// resources. This convention is used to implement gdk_pixbuf_get_file_info()
// efficiently.
type PixbufModuleSizeFunc func(width *int, height *int)

//export gotk4_PixbufModuleSizeFunc
func _PixbufModuleSizeFunc(arg0 *C.gint, arg1 *C.gint, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var width *int  // out
	var height *int // out

	width = (*int)(unsafe.Pointer(arg0))
	height = (*int)(unsafe.Pointer(arg1))

	fn := v.(PixbufModuleSizeFunc)
	fn(width, height)
}

// PixbufModuleUpdatedFunc defines the type of the function that gets called
// every time a region of @pixbuf is updated.
//
// PixbufLoader uses a function of this type to emit the "<link
// linkend="GdkPixbufLoader-area-updated">area_updated</link>" signal.
type PixbufModuleUpdatedFunc func(pixbuf Pixbuf, x int, y int, width int, height int)

//export gotk4_PixbufModuleUpdatedFunc
func _PixbufModuleUpdatedFunc(arg0 *C.GdkPixbuf, arg1 C.int, arg2 C.int, arg3 C.int, arg4 C.int, arg5 C.gpointer) {
	v := box.Get(uintptr(arg5))
	if v == nil {
		panic(`callback not found`)
	}

	var pixbuf Pixbuf // out
	var x int         // out
	var y int         // out
	var width int     // out
	var height int    // out

	pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Pixbuf)
	x = int(arg1)
	y = int(arg2)
	width = int(arg3)
	height = int(arg4)

	fn := v.(PixbufModuleUpdatedFunc)
	fn(pixbuf, x, y, width, height)
}

// PixbufFormat: a `GdkPixbufFormat` contains information about the image format
// accepted by a module.
//
// Only modules should access the fields directly, applications should use the
// `gdk_pixbuf_format_*` family of functions.
type PixbufFormat C.GdkPixbufFormat

// WrapPixbufFormat wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPixbufFormat(ptr unsafe.Pointer) *PixbufFormat {
	return (*PixbufFormat)(ptr)
}

func marshalPixbufFormat(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*PixbufFormat)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PixbufFormat) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// Copy disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) Copy() *PixbufFormat {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.GdkPixbufFormat // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_copy(_arg0)

	var _pixbufFormat *PixbufFormat // out

	_pixbufFormat = (*PixbufFormat)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_pixbufFormat, func(v **PixbufFormat) {
		C.free(unsafe.Pointer(v))
	})

	return _pixbufFormat
}

// Free disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) Free() {
	var _arg0 *C.GdkPixbufFormat // out

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	C.gdk_pixbuf_format_free(_arg0)
}

// Description disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) Description() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_get_description(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Extensions disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) Extensions() []string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret **C.gchar

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_get_extensions(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// License disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) License() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_get_license(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// MIMETypes disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) MIMETypes() []string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret **C.gchar

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_get_mime_types(_arg0)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// Name disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) Name() string {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// IsDisabled disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) IsDisabled() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_is_disabled(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSaveOptionSupported disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) IsSaveOptionSupported(optionKey string) bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(optionKey))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_pixbuf_format_is_save_option_supported(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsScalable disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) IsScalable() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_is_scalable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsWritable disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) IsWritable() bool {
	var _arg0 *C.GdkPixbufFormat // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))

	_cret = C.gdk_pixbuf_format_is_writable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDisabled disables or enables an image format.
//
// If a format is disabled, GdkPixbuf won't use the image loader for this format
// to load images.
//
// Applications can use this to avoid using image loaders with an inappropriate
// license, see gdk_pixbuf_format_get_license().
func (f *PixbufFormat) SetDisabled(disabled bool) {
	var _arg0 *C.GdkPixbufFormat // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GdkPixbufFormat)(unsafe.Pointer(f.Native()))
	if disabled {
		_arg1 = C.TRUE
	}

	C.gdk_pixbuf_format_set_disabled(_arg0, _arg1)
}

// PixbufModule: a `GdkPixbufModule` contains the necessary functions to load
// and save images in a certain file format.
//
// If `GdkPixbuf` has been compiled with `GModule` support, it can be extended
// by modules which can load (and perhaps also save) new image and animation
// formats.
//
//
// Implementing modules
//
// The `GdkPixbuf` interfaces needed for implementing modules are contained in
// `gdk-pixbuf-io.h` (and `gdk-pixbuf-animation.h` if the module supports
// animations). They are not covered by the same stability guarantees as the
// regular GdkPixbuf API. To underline this fact, they are protected by the
// `GDK_PIXBUF_ENABLE_BACKEND` pre-processor symbol.
//
// Each loadable module must contain a `GdkPixbufModuleFillVtableFunc` function
// named `fill_vtable`, which will get called when the module is loaded and must
// set the function pointers of the `GdkPixbufModule`.
//
// In order to make format-checking work before actually loading the modules
// (which may require calling `dlopen` to load image libraries), modules export
// their signatures (and other information) via the `fill_info` function. An
// external utility, `gdk-pixbuf-query-loaders`, uses this to create a text file
// containing a list of all available loaders and their signatures. This file is
// then read at runtime by `GdkPixbuf` to obtain the list of available loaders
// and their signatures.
//
// Modules may only implement a subset of the functionality available via
// `GdkPixbufModule`. If a particular functionality is not implemented, the
// `fill_vtable` function will simply not set the corresponding function
// pointers of the `GdkPixbufModule` structure. If a module supports incremental
// loading (i.e. provides `begin_load`, `stop_load` and `load_increment`), it
// doesn't have to implement `load`, since `GdkPixbuf` can supply a generic
// `load` implementation wrapping the incremental loading.
//
//
// Installing modules
//
// Installing a module is a two-step process:
//
//    - copy the module file(s) to the loader directory (normally
//      `$libdir/gdk-pixbuf-2.0/$version/loaders`, unless overridden by the
//      environment variable `GDK_PIXBUF_MODULEDIR`)
//    - call `gdk-pixbuf-query-loaders` to update the module file (normally
//      `$libdir/gdk-pixbuf-2.0/$version/loaders.cache`, unless overridden
//      by the environment variable `GDK_PIXBUF_MODULE_FILE`)
type PixbufModule C.GdkPixbufModule

// WrapPixbufModule wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPixbufModule(ptr unsafe.Pointer) *PixbufModule {
	return (*PixbufModule)(ptr)
}

// Native returns the underlying C source pointer.
func (p *PixbufModule) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}

// PixbufModulePattern: the signature prefix for a module.
//
// The signature of a module is a set of prefixes. Prefixes are encoded as pairs
// of ordinary strings, where the second string, called the mask, if not `NULL`,
// must be of the same length as the first one and may contain ' ', '!', 'x',
// 'z', and 'n' to indicate bytes that must be matched, not matched,
// "don't-care"-bytes, zeros and non-zeros, respectively.
//
// Each prefix has an associated integer that describes the relevance of the
// prefix, with 0 meaning a mismatch and 100 a "perfect match".
//
// Starting with gdk-pixbuf 2.8, the first byte of the mask may be '*',
// indicating an unanchored pattern that matches not only at the beginning, but
// also in the middle. Versions prior to 2.8 will interpret the '*' like an 'x'.
//
// The signature of a module is stored as an array of `GdkPixbufModulePatterns`.
// The array is terminated by a pattern where the `prefix` is `NULL`.
//
// “`c GdkPixbufModulePattern *signature[] = { { "abcdx", " !x z", 100 }, {
// "bla", NULL, 90 }, { NULL, NULL, 0 } }; “`
//
// In the example above, the signature matches e.g. "auud\0" with relevance 100,
// and "blau" with relevance 90.
type PixbufModulePattern C.GdkPixbufModulePattern

// WrapPixbufModulePattern wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPixbufModulePattern(ptr unsafe.Pointer) *PixbufModulePattern {
	return (*PixbufModulePattern)(ptr)
}

// Native returns the underlying C source pointer.
func (p *PixbufModulePattern) Native() unsafe.Pointer {
	return unsafe.Pointer(p)
}
