// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

//export _gotk4_gdk4_ContentProviderClass_attach_clipboard
func _gotk4_gdk4_ContentProviderClass_attach_clipboard(arg0 *C.GdkContentProvider, arg1 *C.GdkClipboard) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.AttachClipboard == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.AttachClipboard, got none")
	}

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.AttachClipboard(_clipboard)
}

//export _gotk4_gdk4_ContentProviderClass_content_changed
func _gotk4_gdk4_ContentProviderClass_content_changed(arg0 *C.GdkContentProvider) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.ContentChanged == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.ContentChanged, got none")
	}

	overrides.ContentChanged()
}

//export _gotk4_gdk4_ContentProviderClass_detach_clipboard
func _gotk4_gdk4_ContentProviderClass_detach_clipboard(arg0 *C.GdkContentProvider, arg1 *C.GdkClipboard) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.DetachClipboard == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.DetachClipboard, got none")
	}

	var _clipboard *Clipboard // out

	_clipboard = wrapClipboard(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.DetachClipboard(_clipboard)
}

//export _gotk4_gdk4_ContentProviderClass_get_value
func _gotk4_gdk4_ContentProviderClass_get_value(arg0 *C.GdkContentProvider, arg1 *C.GValue, _cerr **C.GError) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.Value == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.Value, got none")
	}

	var _value *coreglib.Value // out

	_value = coreglib.ValueFromNative(unsafe.Pointer(arg1))

	_goerr := overrides.Value(_value)

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gdk4_ContentProviderClass_ref_formats
func _gotk4_gdk4_ContentProviderClass_ref_formats(arg0 *C.GdkContentProvider) (cret *C.GdkContentFormats) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.RefFormats == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.RefFormats, got none")
	}

	contentFormats := overrides.RefFormats()

	var _ *ContentFormats

	cret = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(contentFormats)))

	return cret
}

//export _gotk4_gdk4_ContentProviderClass_ref_storable_formats
func _gotk4_gdk4_ContentProviderClass_ref_storable_formats(arg0 *C.GdkContentProvider) (cret *C.GdkContentFormats) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.RefStorableFormats == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.RefStorableFormats, got none")
	}

	contentFormats := overrides.RefStorableFormats()

	var _ *ContentFormats

	cret = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(contentFormats)))

	return cret
}

//export _gotk4_gdk4_ContentProviderClass_write_mime_type_finish
func _gotk4_gdk4_ContentProviderClass_write_mime_type_finish(arg0 *C.GdkContentProvider, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ContentProviderOverrides](instance0)
	if overrides.WriteMIMETypeFinish == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ContentProviderOverrides.WriteMIMETypeFinish, got none")
	}

	var _result gio.AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.AsyncResulter)
			return ok
		})
		rv, ok := casted.(gio.AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_result = rv
	}

	_goerr := overrides.WriteMIMETypeFinish(_result)

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}