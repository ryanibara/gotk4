// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_memory_format_get_type()), F: marshalMemoryFormat},
		{T: externglib.Type(C.gdk_memory_texture_get_type()), F: marshalMemoryTexturer},
	})
}

// MemoryFormat: GdkMemoryFormat describes a format that bytes can have in
// memory.
//
// It describes formats by listing the contents of the memory passed to it. So
// GDK_MEMORY_A8R8G8B8 will be 1 byte (8 bits) of alpha, followed by a byte each
// of red, green and blue. It is not endian-dependent, so CAIRO_FORMAT_ARGB32 is
// represented by different GdkMemoryFormats on architectures with different
// endiannesses.
//
// Its naming is modelled after VkFormat (see
// https://www.khronos.org/registry/vulkan/specs/1.0/html/vkspec.htmlFormat for
// details).
type MemoryFormat C.gint

const (
	// MemoryB8G8R8A8Premultiplied: 4 bytes; for blue, green, red, alpha. The
	// color values are premultiplied with the alpha value.
	MemoryB8G8R8A8Premultiplied MemoryFormat = iota
	// MemoryA8R8G8B8Premultiplied: 4 bytes; for alpha, red, green, blue. The
	// color values are premultiplied with the alpha value.
	MemoryA8R8G8B8Premultiplied
	// MemoryR8G8B8A8Premultiplied: 4 bytes; for red, green, blue, alpha The
	// color values are premultiplied with the alpha value.
	MemoryR8G8B8A8Premultiplied
	// MemoryB8G8R8A8: 4 bytes; for blue, green, red, alpha.
	MemoryB8G8R8A8
	// MemoryA8R8G8B8: 4 bytes; for alpha, red, green, blue.
	MemoryA8R8G8B8
	// MemoryR8G8B8A8: 4 bytes; for red, green, blue, alpha.
	MemoryR8G8B8A8
	// MemoryA8B8G8R8: 4 bytes; for alpha, blue, green, red.
	MemoryA8B8G8R8
	// MemoryR8G8B8: 3 bytes; for red, green, blue. The data is opaque.
	MemoryR8G8B8
	// MemoryB8G8R8: 3 bytes; for blue, green, red. The data is opaque.
	MemoryB8G8R8
	// MemoryNFormats: number of formats. This value will change as more formats
	// get added, so do not rely on its concrete integer.
	MemoryNFormats
)

func marshalMemoryFormat(p uintptr) (interface{}, error) {
	return MemoryFormat(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for MemoryFormat.
func (m MemoryFormat) String() string {
	switch m {
	case MemoryB8G8R8A8Premultiplied:
		return "B8G8R8A8Premultiplied"
	case MemoryA8R8G8B8Premultiplied:
		return "A8R8G8B8Premultiplied"
	case MemoryR8G8B8A8Premultiplied:
		return "R8G8B8A8Premultiplied"
	case MemoryB8G8R8A8:
		return "B8G8R8A8"
	case MemoryA8R8G8B8:
		return "A8R8G8B8"
	case MemoryR8G8B8A8:
		return "R8G8B8A8"
	case MemoryA8B8G8R8:
		return "A8B8G8R8"
	case MemoryR8G8B8:
		return "R8G8B8"
	case MemoryB8G8R8:
		return "B8G8R8"
	case MemoryNFormats:
		return "NFormats"
	default:
		return fmt.Sprintf("MemoryFormat(%d)", m)
	}
}

// MemoryTexture: GdkTexture representing image data in memory.
type MemoryTexture struct {
	_ [0]func() // equal guard
	Texture
}

var (
	_ Texturer = (*MemoryTexture)(nil)
)

func wrapMemoryTexture(obj *externglib.Object) *MemoryTexture {
	return &MemoryTexture{
		Texture: Texture{
			Object: obj,
			Paintable: Paintable{
				Object: obj,
			},
		},
	}
}

func marshalMemoryTexturer(p uintptr) (interface{}, error) {
	return wrapMemoryTexture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewMemoryTexture creates a new texture for a blob of image data.
//
// The GBytes must contain stride x height pixels in the given format.
//
// The function takes the following parameters:
//
//    - width of the texture.
//    - height of the texture.
//    - format of the data.
//    - bytes: GBytes containing the pixel data.
//    - stride for the data.
//
// The function returns the following values:
//
//    - memoryTexture: newly-created GdkTexture.
//
func NewMemoryTexture(width, height int, format MemoryFormat, bytes *glib.Bytes, stride uint) *MemoryTexture {
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _arg3 C.GdkMemoryFormat // out
	var _arg4 *C.GBytes         // out
	var _arg5 C.gsize           // out
	var _cret *C.GdkTexture     // in

	_arg1 = C.int(width)
	_arg2 = C.int(height)
	_arg3 = C.GdkMemoryFormat(format)
	_arg4 = (*C.GBytes)(gextras.StructNative(unsafe.Pointer(bytes)))
	_arg5 = C.gsize(stride)

	_cret = C.gdk_memory_texture_new(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
	runtime.KeepAlive(format)
	runtime.KeepAlive(bytes)
	runtime.KeepAlive(stride)

	var _memoryTexture *MemoryTexture // out

	_memoryTexture = wrapMemoryTexture(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _memoryTexture
}
