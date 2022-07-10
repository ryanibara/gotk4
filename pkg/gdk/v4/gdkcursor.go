// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeCursor returns the GType for the type Cursor.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeCursor() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gdk", "Cursor").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalCursor)
	return gtype
}

// Cursor: GdkCursor is used to create and destroy cursors.
//
// Cursors are immutable objects, so once you created them, there is no way to
// modify them later. You should create a new cursor when you want to change
// something about it.
//
// Cursors by themselves are not very interesting: they must be bound to a
// window for users to see them. This is done with gdk.Surface.SetCursor() or
// gdk.Surface.SetDeviceCursor(). Applications will typically use higher-level
// GTK functions such as gtk.Widget.SetCursor()` instead.
//
// Cursors are not bound to a given gdk.Display, so they can be shared. However,
// the appearance of cursors may vary when used on different platforms.
//
//
// Named and texture cursors
//
// There are multiple ways to create cursors. The platform's own cursors can be
// created with gdk.Cursor.NewFromName. That function lists the commonly
// available names that are shared with the CSS specification. Other names may
// be available, depending on the platform in use. On some platforms, what
// images are used for named cursors may be influenced by the cursor theme.
//
// Another option to create a cursor is to use gdk.Cursor.NewFromTexture and
// provide an image to use for the cursor.
//
// To ease work with unsupported cursors, a fallback cursor can be provided. If
// a gdk.Surface cannot use a cursor because of the reasons mentioned above, it
// will try the fallback cursor. Fallback cursors can themselves have fallback
// cursors again, so it is possible to provide a chain of progressively easier
// to support cursors. If none of the provided cursors can be supported, the
// default cursor will be the ultimate fallback.
type Cursor struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Cursor)(nil)
)

func wrapCursor(obj *coreglib.Object) *Cursor {
	return &Cursor{
		Object: obj,
	}
}

func marshalCursor(p uintptr) (interface{}, error) {
	return wrapCursor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewCursorFromName creates a new cursor by looking up name in the current
// cursor theme.
//
// A recommended set of cursor names that will work across different platforms
// can be found in the CSS specification:
//
// | | | | | | --- | --- | ---- | --- | | "none" | ! (default_cursor.png)
// "default" | ! (help_cursor.png) "help" | ! (pointer_cursor.png) "pointer" | |
// ! (context_menu_cursor.png) "context-menu" | ! (progress_cursor.png)
// "progress" | ! (wait_cursor.png) "wait" | ! (cell_cursor.png) "cell" | | !
// (crosshair_cursor.png) "crosshair" | ! (text_cursor.png) "text" | !
// (vertical_text_cursor.png) "vertical-text" | ! (alias_cursor.png) "alias" | |
// ! (copy_cursor.png) "copy" | ! (no_drop_cursor.png) "no-drop" | !
// (move_cursor.png) "move" | ! (not_allowed_cursor.png) "not-allowed" | | !
// (grab_cursor.png) "grab" | ! (grabbing_cursor.png) "grabbing" | !
// (all_scroll_cursor.png) "all-scroll" | ! (col_resize_cursor.png) "col-resize"
// | | ! (row_resize_cursor.png) "row-resize" | ! (n_resize_cursor.png)
// "n-resize" | ! (e_resize_cursor.png) "e-resize" | ! (s_resize_cursor.png)
// "s-resize" | | ! (w_resize_cursor.png) "w-resize" | ! (ne_resize_cursor.png)
// "ne-resize" | ! (nw_resize_cursor.png) "nw-resize" | ! (sw_resize_cursor.png)
// "sw-resize" | | ! (se_resize_cursor.png) "se-resize" | !
// (ew_resize_cursor.png) "ew-resize" | ! (ns_resize_cursor.png) "ns-resize" | !
// (nesw_resize_cursor.png) "nesw-resize" | | ! (nwse_resize_cursor.png)
// "nwse-resize" | ! (zoom_in_cursor.png) "zoom-in" | ! (zoom_out_cursor.png)
// "zoom-out" | |.
//
// The function takes the following parameters:
//
//    - name of the cursor.
//    - fallback (optional): NULL or the GdkCursor to fall back to when this one
//      cannot be supported.
//
// The function returns the following values:
//
//    - cursor (optional): new GdkCursor, or NULL if there is no cursor with the
//      given name.
//
func NewCursorFromName(name string, fallback *Cursor) *Cursor {
	var _args [2]girepository.Argument

	*(**C.char)(unsafe.Pointer(&_args[0])) = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_args[0]))))
	if fallback != nil {
		*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fallback).Native()))
	}

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("new_Cursor_from_name", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(name)
	runtime.KeepAlive(fallback)

	var _cursor *Cursor // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_cursor = wrapCursor(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _cursor
}

// NewCursorFromTexture creates a new cursor from a GdkTexture.
//
// The function takes the following parameters:
//
//    - texture providing the pixel data.
//    - hotspotX: horizontal offset of the “hotspot” of the cursor.
//    - hotspotY: vertical offset of the “hotspot” of the cursor.
//    - fallback (optional): NULL or the GdkCursor to fall back to when this one
//      cannot be supported.
//
// The function returns the following values:
//
//    - cursor: new GdkCursor.
//
func NewCursorFromTexture(texture Texturer, hotspotX, hotspotY int32, fallback *Cursor) *Cursor {
	var _args [4]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(texture).Native()))
	*(*C.int)(unsafe.Pointer(&_args[1])) = C.int(hotspotX)
	*(*C.int)(unsafe.Pointer(&_args[2])) = C.int(hotspotY)
	if fallback != nil {
		*(**C.void)(unsafe.Pointer(&_args[3])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(fallback).Native()))
	}

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("new_Cursor_from_texture", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(texture)
	runtime.KeepAlive(hotspotX)
	runtime.KeepAlive(hotspotY)
	runtime.KeepAlive(fallback)

	var _cursor *Cursor // out

	_cursor = wrapCursor(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _cursor
}

// Fallback returns the fallback for this cursor.
//
// The fallback will be used if this cursor is not available on a given
// GdkDisplay. For named cursors, this can happen when using nonstandard names
// or when using an incomplete cursor theme. For textured cursors, this can
// happen when the texture is too large or when the GdkDisplay it is used on
// does not support textured cursors.
//
// The function returns the following values:
//
//    - ret (optional): fallback of the cursor or NULL to use the default cursor
//      as fallback.
//
func (cursor *Cursor) Fallback() *Cursor {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("get_fallback", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cursor)

	var _ret *Cursor // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_ret = wrapCursor(coreglib.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	}

	return _ret
}

// HotspotX returns the horizontal offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function will
// only return the hotspot position for cursors created with
// gdk.Cursor.NewFromTexture.
//
// The function returns the following values:
//
//    - gint: horizontal offset of the hotspot or 0 for named cursors.
//
func (cursor *Cursor) HotspotX() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("get_hotspot_x", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cursor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// HotspotY returns the vertical offset of the hotspot.
//
// The hotspot indicates the pixel that will be directly above the cursor.
//
// Note that named cursors may have a nonzero hotspot, but this function will
// only return the hotspot position for cursors created with
// gdk.Cursor.NewFromTexture.
//
// The function returns the following values:
//
//    - gint: vertical offset of the hotspot or 0 for named cursors.
//
func (cursor *Cursor) HotspotY() int32 {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("get_hotspot_y", _args[:], nil)
	_cret := *(*C.int)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cursor)

	var _gint int32 // out

	_gint = int32(*(*C.int)(unsafe.Pointer(&_cret)))

	return _gint
}

// Name returns the name of the cursor.
//
// If the cursor is not a named cursor, NULL will be returned.
//
// The function returns the following values:
//
//    - utf8 (optional): name of the cursor or NULL if it is not a named cursor.
//
func (cursor *Cursor) Name() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("get_name", _args[:], nil)
	_cret := *(**C.char)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cursor)

	var _utf8 string // out

	if *(**C.char)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.char)(unsafe.Pointer(&_cret)))))
	}

	return _utf8
}

// Texture returns the texture for the cursor.
//
// If the cursor is a named cursor, NULL will be returned.
//
// The function returns the following values:
//
//    - texture (optional) for cursor or NULL if it is a named cursor.
//
func (cursor *Cursor) Texture() Texturer {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(cursor).Native()))

	_info := girepository.MustFind("Gdk", "Cursor")
	_gret := _info.InvokeClassMethod("get_texture", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(cursor)

	var _texture Texturer // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		{
			objptr := unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Texturer)
				return ok
			})
			rv, ok := casted.(Texturer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Texturer")
			}
			_texture = rv
		}
	}

	return _texture
}
