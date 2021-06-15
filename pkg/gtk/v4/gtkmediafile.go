// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
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
	gdk.Paintable

	// Clear resets the media file to be empty.
	Clear()
	// File returns the file that @self is currently playing from.
	//
	// When @self is not playing or not playing from a file, nil is returned.
	File() gio.File
	// InputStream returns the stream that @self is currently playing from.
	//
	// When @self is not playing or not playing from a stream, nil is returned.
	InputStream() gio.InputStream
	// SetFile sets the `GtkMediaFile` to play the given file.
	//
	// If any file is still playing, stop playing it.
	SetFile(file gio.File)
	// SetFilename sets the `GtkMediaFile to play the given file.
	//
	// This is a utility function that converts the given @filename to a `GFile`
	// and calls [method@Gtk.MediaFile.set_file].
	SetFilename(filename string)
	// SetInputStream sets the `GtkMediaFile` to play the given stream.
	//
	// If anything is still playing, stop playing it.
	//
	// Full control about the @stream is assumed for the duration of playback.
	// The stream will not be closed.
	SetInputStream(stream gio.InputStream)
	// SetResource sets the `GtkMediaFile to play the given resource.
	//
	// This is a utility function that converts the given @resource_path to a
	// `GFile` and calls [method@Gtk.MediaFile.set_file].
	SetResource(resourcePath string)
}

// mediaFile implements the MediaFile class.
type mediaFile struct {
	MediaStream
	gdk.Paintable
}

var _ MediaFile = (*mediaFile)(nil)

// WrapMediaFile wraps a GObject to the right type. It is
// primarily used internally.
func WrapMediaFile(obj *externglib.Object) MediaFile {
	return mediaFile{
		MediaStream:   WrapMediaStream(obj),
		gdk.Paintable: gdk.WrapPaintable(obj),
	}
}

func marshalMediaFile(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMediaFile(obj), nil
}

// NewMediaFile constructs a class MediaFile.
func NewMediaFile() MediaFile {
	var _cret C.GtkMediaFile // in

	_cret = C.gtk_media_file_new()

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MediaFile)

	return _mediaFile
}

// NewMediaFileForFile constructs a class MediaFile.
func NewMediaFileForFile(file gio.File) MediaFile {
	var _arg1 *C.GFile       // out
	var _cret C.GtkMediaFile // in

	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	_cret = C.gtk_media_file_new_for_file(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MediaFile)

	return _mediaFile
}

// NewMediaFileForFilename constructs a class MediaFile.
func NewMediaFileForFilename(filename string) MediaFile {
	var _arg1 *C.char        // out
	var _cret C.GtkMediaFile // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_filename(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MediaFile)

	return _mediaFile
}

// NewMediaFileForInputStream constructs a class MediaFile.
func NewMediaFileForInputStream(stream gio.InputStream) MediaFile {
	var _arg1 *C.GInputStream // out
	var _cret C.GtkMediaFile  // in

	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.gtk_media_file_new_for_input_stream(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MediaFile)

	return _mediaFile
}

// NewMediaFileForResource constructs a class MediaFile.
func NewMediaFileForResource(resourcePath string) MediaFile {
	var _arg1 *C.char        // out
	var _cret C.GtkMediaFile // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_media_file_new_for_resource(_arg1)

	var _mediaFile MediaFile // out

	_mediaFile = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MediaFile)

	return _mediaFile
}

// Clear resets the media file to be empty.
func (s mediaFile) Clear() {
	var _arg0 *C.GtkMediaFile // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))

	C.gtk_media_file_clear(_arg0)
}

// File returns the file that @self is currently playing from.
//
// When @self is not playing or not playing from a file, nil is returned.
func (s mediaFile) File() gio.File {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GFile        // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_file_get_file(_arg0)

	var _file gio.File // out

	_file = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.File)

	return _file
}

// InputStream returns the stream that @self is currently playing from.
//
// When @self is not playing or not playing from a stream, nil is returned.
func (s mediaFile) InputStream() gio.InputStream {
	var _arg0 *C.GtkMediaFile // out
	var _cret *C.GInputStream // in

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_media_file_get_input_stream(_arg0)

	var _inputStream gio.InputStream // out

	_inputStream = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.InputStream)

	return _inputStream
}

// SetFile sets the `GtkMediaFile` to play the given file.
//
// If any file is still playing, stop playing it.
func (s mediaFile) SetFile(file gio.File) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GFile        // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_media_file_set_file(_arg0, _arg1)
}

// SetFilename sets the `GtkMediaFile to play the given file.
//
// This is a utility function that converts the given @filename to a `GFile`
// and calls [method@Gtk.MediaFile.set_file].
func (s mediaFile) SetFilename(filename string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_filename(_arg0, _arg1)
}

// SetInputStream sets the `GtkMediaFile` to play the given stream.
//
// If anything is still playing, stop playing it.
//
// Full control about the @stream is assumed for the duration of playback.
// The stream will not be closed.
func (s mediaFile) SetInputStream(stream gio.InputStream) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.GInputStream // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GInputStream)(unsafe.Pointer(stream.Native()))

	C.gtk_media_file_set_input_stream(_arg0, _arg1)
}

// SetResource sets the `GtkMediaFile to play the given resource.
//
// This is a utility function that converts the given @resource_path to a
// `GFile` and calls [method@Gtk.MediaFile.set_file].
func (s mediaFile) SetResource(resourcePath string) {
	var _arg0 *C.GtkMediaFile // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GtkMediaFile)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_media_file_set_resource(_arg0, _arg1)
}
