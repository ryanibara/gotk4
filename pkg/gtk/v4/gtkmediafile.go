// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_media_file_get_type()), F: marshalMediaFile},
	})
}

// MediaFile: `GtkMediaFile` implements `GtkMediaStream` for files.
//
// This provides a simple way to play back video files with GTK.
//
// GTK provides a GIO extension point for `GtkMediaFile` implementations to
// allow for external implementations using various media frameworks.
//
// GTK itself includes implementations using GStreamer and ffmpeg.
type MediaFile interface {
	MediaStream

	ClearMediaFile()

	InputStream() gio.InputStream

	SetFilenameMediaFile(filename string)

	SetInputStreamMediaFile(stream gio.InputStream)

	SetResourceMediaFile(resourcePath string)
}

// mediaFile implements the MediaFile class.
type mediaFile struct {
	MediaStream
}

// WrapMediaFile wraps a GObject to the right type. It is
// primarily used internally.
func WrapMediaFile(obj *externglib.Object) MediaFile {
	return mediaFile{
		MediaStream: WrapMediaStream(obj),
	}
}

func marshalMediaFile(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaFile(obj), nil
}

func NewMediaFile() MediaFile {
	var _cret *C.GtkMediaStream // in

	_cret = C.gtk_media_file_new()

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MediaFile)

	return _mediaFile
}

func NewMediaFileForFilename(filename string) MediaFile {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_filename(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MediaFile)

	return _mediaFile
}

func NewMediaFileForInputStream(stream gio.InputStream) MediaFile {
	var _arg1 *C.GInputStream   // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.gtk_media_file_new_for_input_stream(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MediaFile)

	return _mediaFile
}

func NewMediaFileForResource(resourcePath string) MediaFile {
	var _arg1 *C.char           // out
	var _cret *C.GtkMediaStream // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_resource(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MediaFile)

	return _mediaFile
}

func (s mediaFile) ClearMediaFile() {
	var _arg0 *C.GtkMediaFile // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))

	C.gtk_media_file_clear(_arg0)
}

func (s mediaFile) InputStream() gio.InputStream {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GInputStream // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_file_get_input_stream(_arg0)

	var _inputStream gio.InputStream // out

	_inputStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.InputStream)

	return _inputStream
}

func (s mediaFile) SetFilenameMediaFile(filename string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_filename(_arg0, _arg1)
}

func (s mediaFile) SetInputStreamMediaFile(stream gio.InputStream) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GInputStream // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))

	C.gtk_media_file_set_input_stream(_arg0, _arg1)
}

func (s mediaFile) SetResourceMediaFile(resourcePath string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_resource(_arg0, _arg1)
}
