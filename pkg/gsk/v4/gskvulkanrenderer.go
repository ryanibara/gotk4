// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_vulkan_renderer_get_type()), F: marshalVulkanRenderer},
	})
}

// VulkanRenderer: GSK renderer that is using Vulkan.
type VulkanRenderer interface {
	Renderer
}

// vulkanRenderer implements the VulkanRenderer class.
type vulkanRenderer struct {
	Renderer
}

// WrapVulkanRenderer wraps a GObject to the right type. It is
// primarily used internally.
func WrapVulkanRenderer(obj *externglib.Object) VulkanRenderer {
	return vulkanRenderer{
		Renderer: WrapRenderer(obj),
	}
}

func marshalVulkanRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapVulkanRenderer(obj), nil
}

func NewVulkanRenderer() VulkanRenderer {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_vulkan_renderer_new()

	var _vulkanRenderer VulkanRenderer // out

	_vulkanRenderer = WrapVulkanRenderer(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _vulkanRenderer
}
