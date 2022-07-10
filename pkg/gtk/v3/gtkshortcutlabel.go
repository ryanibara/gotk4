// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeShortcutLabel returns the GType for the type ShortcutLabel.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeShortcutLabel() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "ShortcutLabel").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalShortcutLabel)
	return gtype
}

// ShortcutLabel is a widget that represents a single keyboard shortcut or
// gesture in the user interface.
type ShortcutLabel struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Containerer       = (*ShortcutLabel)(nil)
	_ coreglib.Objector = (*ShortcutLabel)(nil)
)

func wrapShortcutLabel(obj *coreglib.Object) *ShortcutLabel {
	return &ShortcutLabel{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutLabel(p uintptr) (interface{}, error) {
	return wrapShortcutLabel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewShortcutLabel creates a new ShortcutLabel with accelerator set.
//
// The function takes the following parameters:
//
//    - accelerator: initial accelerator.
//
// The function returns the following values:
//
//    - shortcutLabel: newly-allocated ShortcutLabel.
//
func NewShortcutLabel(accelerator string) *ShortcutLabel {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "ShortcutLabel")
	_gret := _info.InvokeClassMethod("new_ShortcutLabel", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(accelerator)

	var _shortcutLabel *ShortcutLabel // out

	_shortcutLabel = wrapShortcutLabel(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _shortcutLabel
}

// Accelerator retrieves the current accelerator of self.
//
// The function returns the following values:
//
//    - utf8 (optional): current accelerator.
//
func (self *ShortcutLabel) Accelerator() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "ShortcutLabel")
	_gret := _info.InvokeClassMethod("get_accelerator", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.gchar)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	}

	return _utf8
}

// DisabledText retrieves the text that is displayed when no accelerator is set.
//
// The function returns the following values:
//
//    - utf8 (optional): current text displayed when no accelerator is set.
//
func (self *ShortcutLabel) DisabledText() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_info := girepository.MustFind("Gtk", "ShortcutLabel")
	_gret := _info.InvokeClassMethod("get_disabled_text", _args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.gchar)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	}

	return _utf8
}

// SetAccelerator sets the accelerator to be displayed by self.
//
// The function takes the following parameters:
//
//    - accelerator: new accelerator.
//
func (self *ShortcutLabel) SetAccelerator(accelerator string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(accelerator)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ShortcutLabel")
	_info.InvokeClassMethod("set_accelerator", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(accelerator)
}

// SetDisabledText sets the text to be displayed by self when no accelerator is
// set.
//
// The function takes the following parameters:
//
//    - disabledText: text to be displayed when no accelerator is set.
//
func (self *ShortcutLabel) SetDisabledText(disabledText string) {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(disabledText)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "ShortcutLabel")
	_info.InvokeClassMethod("set_disabled_text", _args[:], nil)

	runtime.KeepAlive(self)
	runtime.KeepAlive(disabledText)
}
