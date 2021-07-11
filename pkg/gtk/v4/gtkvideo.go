// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_video_get_type()), F: marshalVideoer},
	})
}

// Videoer describes Video's methods.
type Videoer interface {
	// Autoplay returns true if videos have been set to loop.
	Autoplay() bool
	// File gets the file played by @self or nil if not playing back a file.
	File() *gio.File
	// Loop returns true if videos have been set to loop.
	Loop() bool
	// MediaStream gets the media stream managed by @self or nil if none.
	MediaStream() *MediaStream
	// SetAutoplay sets whether @self automatically starts playback when it
	// becomes visible or when a new file gets loaded.
	SetAutoplay(autoplay bool)
	// SetFile makes @self play the given @file.
	SetFile(file gio.Filer)
	// SetFilename makes @self play the given @filename.
	SetFilename(filename string)
	// SetLoop sets whether new files loaded by @self should be set to loop.
	SetLoop(loop bool)
	// SetMediaStream sets the media stream to be played back.
	SetMediaStream(stream MediaStreamer)
	// SetResource makes @self play the resource at the given @resource_path.
	SetResource(resourcePath string)
}

// Video: `GtkVideo` is a widget to show a `GtkMediaStream` with media controls.
//
// !An example GtkVideo (video.png)
//
// The controls are available separately as [class@Gtk.MediaControls]. If you
// just want to display a video without controls, you can treat it like any
// other paintable and for example put it into a [class@Gtk.Picture].
//
// `GtkVideo` aims to cover use cases such as previews, embedded animations,
// etc. It supports autoplay, looping, and simple media controls. It does not
// have support for video overlays, multichannel audio, device selection, or
// input. If you are writing a full-fledged video player, you may want to use
// the [class@Gdk.Paintable] API and a media framework such as Gstreamer
// directly.
type Video struct {
	Widget
}

var (
	_ Videoer         = (*Video)(nil)
	_ gextras.Nativer = (*Video)(nil)
)

func wrapVideo(obj *externglib.Object) Videoer {
	return &Video{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalVideoer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVideo(obj), nil
}

// NewVideo creates a new empty `GtkVideo`.
func NewVideo() *Video {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_video_new()

	var _video *Video // out

	_video = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Video)

	return _video
}

// NewVideoForFile creates a `GtkVideo` to play back the given @file.
func NewVideoForFile(file gio.Filer) *Video {
	var _arg1 *C.GFile     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))

	_cret = C.gtk_video_new_for_file(_arg1)

	var _video *Video // out

	_video = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Video)

	return _video
}

// NewVideoForFilename creates a `GtkVideo` to play back the given @filename.
//
// This is a utility function that calls [ctor@Gtk.Video.new_for_file], See that
// function for details.
func NewVideoForFilename(filename string) *Video {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_video_new_for_filename(_arg1)

	var _video *Video // out

	_video = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Video)

	return _video
}

// NewVideoForMediaStream creates a `GtkVideo` to play back the given @stream.
func NewVideoForMediaStream(stream MediaStreamer) *Video {
	var _arg1 *C.GtkMediaStream // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))

	_cret = C.gtk_video_new_for_media_stream(_arg1)

	var _video *Video // out

	_video = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Video)

	return _video
}

// NewVideoForResource creates a `GtkVideo` to play back the resource at the
// given @resource_path.
//
// This is a utility function that calls [ctor@Gtk.Video.new_for_file].
func NewVideoForResource(resourcePath string) *Video {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_video_new_for_resource(_arg1)

	var _video *Video // out

	_video = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*Video)

	return _video
}

// Autoplay returns true if videos have been set to loop.
func (self *Video) Autoplay() bool {
	var _arg0 *C.GtkVideo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_video_get_autoplay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// File gets the file played by @self or nil if not playing back a file.
func (self *Video) File() *gio.File {
	var _arg0 *C.GtkVideo // out
	var _cret *C.GFile    // in

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_video_get_file(_arg0)

	var _file *gio.File // out

	_file = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.File)

	return _file
}

// Loop returns true if videos have been set to loop.
func (self *Video) Loop() bool {
	var _arg0 *C.GtkVideo // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_video_get_loop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MediaStream gets the media stream managed by @self or nil if none.
func (self *Video) MediaStream() *MediaStream {
	var _arg0 *C.GtkVideo       // out
	var _cret *C.GtkMediaStream // in

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_video_get_media_stream(_arg0)

	var _mediaStream *MediaStream // out

	_mediaStream = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*MediaStream)

	return _mediaStream
}

// SetAutoplay sets whether @self automatically starts playback when it becomes
// visible or when a new file gets loaded.
func (self *Video) SetAutoplay(autoplay bool) {
	var _arg0 *C.GtkVideo // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	if autoplay {
		_arg1 = C.TRUE
	}

	C.gtk_video_set_autoplay(_arg0, _arg1)
}

// SetFile makes @self play the given @file.
func (self *Video) SetFile(file gio.Filer) {
	var _arg0 *C.GtkVideo // out
	var _arg1 *C.GFile    // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GFile)(unsafe.Pointer((file).(gextras.Nativer).Native()))

	C.gtk_video_set_file(_arg0, _arg1)
}

// SetFilename makes @self play the given @filename.
//
// This is a utility function that calls gtk_video_set_file(),
func (self *Video) SetFilename(filename string) {
	var _arg0 *C.GtkVideo // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_video_set_filename(_arg0, _arg1)
}

// SetLoop sets whether new files loaded by @self should be set to loop.
func (self *Video) SetLoop(loop bool) {
	var _arg0 *C.GtkVideo // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	if loop {
		_arg1 = C.TRUE
	}

	C.gtk_video_set_loop(_arg0, _arg1)
}

// SetMediaStream sets the media stream to be played back.
//
// @self will take full control of managing the media stream. If you want to
// manage a media stream yourself, consider using a [class@Gtk.Picture] for
// display.
//
// If you want to display a file, consider using [method@Gtk.Video.set_file]
// instead.
func (self *Video) SetMediaStream(stream MediaStreamer) {
	var _arg0 *C.GtkVideo       // out
	var _arg1 *C.GtkMediaStream // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GtkMediaStream)(unsafe.Pointer((stream).(gextras.Nativer).Native()))

	C.gtk_video_set_media_stream(_arg0, _arg1)
}

// SetResource makes @self play the resource at the given @resource_path.
//
// This is a utility function that calls [method@Gtk.Video.set_file].
func (self *Video) SetResource(resourcePath string) {
	var _arg0 *C.GtkVideo // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkVideo)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_video_set_resource(_arg0, _arg1)
}
