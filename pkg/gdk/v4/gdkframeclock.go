// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_frame_clock_phase_get_type()), F: marshalFrameClockPhase},
		{T: externglib.Type(C.gdk_frame_clock_get_type()), F: marshalFrameClocker},
	})
}

// FrameClockPhase: used to represent the different paint clock phases that can
// be requested.
//
// The elements of the enumeration correspond to the signals of GdkFrameClock.
type FrameClockPhase C.guint

const (
	// FrameClockPhaseNone: no phase.
	FrameClockPhaseNone FrameClockPhase = 0b0
	// FrameClockPhaseFlushEvents corresponds to GdkFrameClock::flush-events.
	// Should not be handled by applications.
	FrameClockPhaseFlushEvents FrameClockPhase = 0b1
	// FrameClockPhaseBeforePaint corresponds to GdkFrameClock::before-paint.
	// Should not be handled by applications.
	FrameClockPhaseBeforePaint FrameClockPhase = 0b10
	// FrameClockPhaseUpdate corresponds to GdkFrameClock::update.
	FrameClockPhaseUpdate FrameClockPhase = 0b100
	// FrameClockPhaseLayout corresponds to GdkFrameClock::layout. Should not be
	// handled by applicatiosn.
	FrameClockPhaseLayout FrameClockPhase = 0b1000
	// FrameClockPhasePaint corresponds to GdkFrameClock::paint.
	FrameClockPhasePaint FrameClockPhase = 0b10000
	// FrameClockPhaseResumeEvents corresponds to GdkFrameClock::resume-events.
	// Should not be handled by applications.
	FrameClockPhaseResumeEvents FrameClockPhase = 0b100000
	// FrameClockPhaseAfterPaint corresponds to GdkFrameClock::after-paint.
	// Should not be handled by applications.
	FrameClockPhaseAfterPaint FrameClockPhase = 0b1000000
)

func marshalFrameClockPhase(p uintptr) (interface{}, error) {
	return FrameClockPhase(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for FrameClockPhase.
func (f FrameClockPhase) String() string {
	if f == 0 {
		return "FrameClockPhase(0)"
	}

	var builder strings.Builder
	builder.Grow(192)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FrameClockPhaseNone:
			builder.WriteString("None|")
		case FrameClockPhaseFlushEvents:
			builder.WriteString("FlushEvents|")
		case FrameClockPhaseBeforePaint:
			builder.WriteString("BeforePaint|")
		case FrameClockPhaseUpdate:
			builder.WriteString("Update|")
		case FrameClockPhaseLayout:
			builder.WriteString("Layout|")
		case FrameClockPhasePaint:
			builder.WriteString("Paint|")
		case FrameClockPhaseResumeEvents:
			builder.WriteString("ResumeEvents|")
		case FrameClockPhaseAfterPaint:
			builder.WriteString("AfterPaint|")
		default:
			builder.WriteString(fmt.Sprintf("FrameClockPhase(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FrameClockPhase) Has(other FrameClockPhase) bool {
	return (f & other) == other
}

// FrameClock: GdkFrameClock tells the application when to update and repaint a
// surface.
//
// This may be synced to the vertical refresh rate of the monitor, for example.
// Even when the frame clock uses a simple timer rather than a hardware-based
// vertical sync, the frame clock helps because it ensures everything paints at
// the same time (reducing the total number of frames).
//
// The frame clock can also automatically stop painting when it knows the frames
// will not be visible, or scale back animation framerates.
//
// GdkFrameClock is designed to be compatible with an OpenGL-based
// implementation or with mozRequestAnimationFrame in Firefox, for example.
//
// A frame clock is idle until someone requests a frame with
// gdk.FrameClock.RequestPhase(). At some later point that makes sense for the
// synchronization being implemented, the clock will process a frame and emit
// signals for each phase that has been requested. (See the signals of the
// GdkFrameClock class for documentation of the phases.
// GDK_FRAME_CLOCK_PHASE_UPDATE and the gdkframeclock::update signal are most
// interesting for application writers, and are used to update the animations,
// using the frame time given by gdk.FrameClock.GetFrameTime.
//
// The frame time is reported in microseconds and generally in the same
// timescale as g_get_monotonic_time(), however, it is not the same as
// g_get_monotonic_time(). The frame time does not advance during the time a
// frame is being painted, and outside of a frame, an attempt is made so that
// all calls to gdk.FrameClock.GetFrameTime() that are called at a “similar”
// time get the same value. This means that if different animations are timed by
// looking at the difference in time between an initial value from
// gdk.FrameClock.GetFrameTime() and the value inside the gdkframeclock::update
// signal of the clock, they will stay exactly synchronized.
type FrameClock struct {
	*externglib.Object
}

var (
	_ externglib.Objector = (*FrameClock)(nil)
)

// FrameClocker describes types inherited from class FrameClock.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type FrameClocker interface {
	externglib.Objector
	baseFrameClock() *FrameClock
}

var _ FrameClocker = (*FrameClock)(nil)

func wrapFrameClock(obj *externglib.Object) *FrameClock {
	return &FrameClock{
		Object: obj,
	}
}

func marshalFrameClocker(p uintptr) (interface{}, error) {
	return wrapFrameClock(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (frameClock *FrameClock) baseFrameClock() *FrameClock {
	return frameClock
}

// BaseFrameClock returns the underlying base object.
func BaseFrameClock(obj FrameClocker) *FrameClock {
	return obj.baseFrameClock()
}

// ConnectAfterPaint: this signal ends processing of the frame.
//
// Applications should generally not handle this signal.
func (frameClock *FrameClock) ConnectAfterPaint(f func()) externglib.SignalHandle {
	return frameClock.Connect("after-paint", f)
}

// ConnectBeforePaint begins processing of the frame.
//
// Applications should generally not handle this signal.
func (frameClock *FrameClock) ConnectBeforePaint(f func()) externglib.SignalHandle {
	return frameClock.Connect("before-paint", f)
}

// ConnectFlushEvents: used to flush pending motion events that are being
// batched up and compressed together.
//
// Applications should not handle this signal.
func (frameClock *FrameClock) ConnectFlushEvents(f func()) externglib.SignalHandle {
	return frameClock.Connect("flush-events", f)
}

// ConnectLayout: emitted as the second step of toolkit and application
// processing of the frame.
//
// Any work to update sizes and positions of application elements should be
// performed. GTK normally handles this internally.
func (frameClock *FrameClock) ConnectLayout(f func()) externglib.SignalHandle {
	return frameClock.Connect("layout", f)
}

// ConnectPaint: emitted as the third step of toolkit and application processing
// of the frame.
//
// The frame is repainted. GDK normally handles this internally and emits
// gdk.Surface::render signals which are turned into gtk.Widget::snapshot
// signals by GTK.
func (frameClock *FrameClock) ConnectPaint(f func()) externglib.SignalHandle {
	return frameClock.Connect("paint", f)
}

// ConnectResumeEvents: emitted after processing of the frame is finished.
//
// This signal is handled internally by GTK to resume normal event processing.
// Applications should not handle this signal.
func (frameClock *FrameClock) ConnectResumeEvents(f func()) externglib.SignalHandle {
	return frameClock.Connect("resume-events", f)
}

// ConnectUpdate: emitted as the first step of toolkit and application
// processing of the frame.
//
// Animations should be updated using gdk.FrameClock.GetFrameTime().
// Applications can connect directly to this signal, or use
// gtk.Widget.AddTickCallback() as a more convenient interface.
func (frameClock *FrameClock) ConnectUpdate(f func()) externglib.SignalHandle {
	return frameClock.Connect("update", f)
}

// BeginUpdating starts updates for an animation.
//
// Until a matching call to gdk.FrameClock.EndUpdating() is made, the frame
// clock will continually request a new frame with the
// GDK_FRAME_CLOCK_PHASE_UPDATE phase. This function may be called multiple
// times and frames will be requested until gdk_frame_clock_end_updating() is
// called the same number of times.
func (frameClock *FrameClock) BeginUpdating() {
	var _arg0 *C.GdkFrameClock // out

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	C.gdk_frame_clock_begin_updating(_arg0)
	runtime.KeepAlive(frameClock)
}

// EndUpdating stops updates for an animation.
//
// See the documentation for gdk.FrameClock.BeginUpdating().
func (frameClock *FrameClock) EndUpdating() {
	var _arg0 *C.GdkFrameClock // out

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	C.gdk_frame_clock_end_updating(_arg0)
	runtime.KeepAlive(frameClock)
}

// CurrentTimings gets the frame timings for the current frame.
//
// The function returns the following values:
//
//    - frameTimings (optional): GdkFrameTimings for the frame currently being
//      processed, or even no frame is being processed, for the previous frame.
//      Before any frames have been processed, returns NULL.
//
func (frameClock *FrameClock) CurrentTimings() *FrameTimings {
	var _arg0 *C.GdkFrameClock   // out
	var _cret *C.GdkFrameTimings // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	_cret = C.gdk_frame_clock_get_current_timings(_arg0)
	runtime.KeepAlive(frameClock)

	var _frameTimings *FrameTimings // out

	if _cret != nil {
		_frameTimings = (*FrameTimings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gdk_frame_timings_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_frameTimings)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_frame_timings_unref((*C.GdkFrameTimings)(intern.C))
			},
		)
	}

	return _frameTimings
}

// FPS calculates the current frames-per-second, based on the frame timings of
// frame_clock.
//
// The function returns the following values:
//
//    - gdouble: current fps, as a double.
//
func (frameClock *FrameClock) FPS() float64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.double         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	_cret = C.gdk_frame_clock_get_fps(_arg0)
	runtime.KeepAlive(frameClock)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// FrameCounter: GdkFrameClock maintains a 64-bit counter that increments for
// each frame drawn.
//
// The function returns the following values:
//
//    - gint64: inside frame processing, the value of the frame counter for the
//      current frame. Outside of frame processing, the frame counter for the
//      last frame.
//
func (frameClock *FrameClock) FrameCounter() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	_cret = C.gdk_frame_clock_get_frame_counter(_arg0)
	runtime.KeepAlive(frameClock)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// FrameTime gets the time that should currently be used for animations.
//
// Inside the processing of a frame, it’s the time used to compute the animation
// position of everything in a frame. Outside of a frame, it's the time of the
// conceptual “previous frame,” which may be either the actual previous frame
// time, or if that’s too old, an updated time.
//
// The function returns the following values:
//
//    - gint64: timestamp in microseconds, in the timescale of of
//      g_get_monotonic_time().
//
func (frameClock *FrameClock) FrameTime() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	_cret = C.gdk_frame_clock_get_frame_time(_arg0)
	runtime.KeepAlive(frameClock)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// HistoryStart returns the frame counter for the oldest frame available in
// history.
//
// GdkFrameClock internally keeps a history of GdkFrameTimings objects for
// recent frames that can be retrieved with gdk.FrameClock.GetTimings(). The set
// of stored frames is the set from the counter values given by
// gdk.FrameClock.GetHistoryStart() and gdk.FrameClock.GetFrameCounter(),
// inclusive.
//
// The function returns the following values:
//
//    - gint64: frame counter value for the oldest frame that is available in the
//      internal frame history of the GdkFrameClock.
//
func (frameClock *FrameClock) HistoryStart() int64 {
	var _arg0 *C.GdkFrameClock // out
	var _cret C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	_cret = C.gdk_frame_clock_get_history_start(_arg0)
	runtime.KeepAlive(frameClock)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// RefreshInfo predicts a presentation time, based on history.
//
// Using the frame history stored in the frame clock, finds the last known
// presentation time and refresh interval, and assuming that presentation times
// are separated by the refresh interval, predicts a presentation time that is a
// multiple of the refresh interval after the last presentation time, and later
// than base_time.
//
// The function takes the following parameters:
//
//    - baseTime: base time for determining a presentaton time.
//
// The function returns the following values:
//
//    - refreshIntervalReturn (optional): location to store the determined
//      refresh interval, or NULL. A default refresh interval of 1/60th of a
//      second will be stored if no history is present.
//    - presentationTimeReturn: location to store the next candidate presentation
//      time after the given base time. 0 will be will be stored if no history is
//      present.
//
func (frameClock *FrameClock) RefreshInfo(baseTime int64) (refreshIntervalReturn int64, presentationTimeReturn int64) {
	var _arg0 *C.GdkFrameClock // out
	var _arg1 C.gint64         // out
	var _arg2 C.gint64         // in
	var _arg3 C.gint64         // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))
	_arg1 = C.gint64(baseTime)

	C.gdk_frame_clock_get_refresh_info(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(frameClock)
	runtime.KeepAlive(baseTime)

	var _refreshIntervalReturn int64  // out
	var _presentationTimeReturn int64 // out

	_refreshIntervalReturn = int64(_arg2)
	_presentationTimeReturn = int64(_arg3)

	return _refreshIntervalReturn, _presentationTimeReturn
}

// Timings retrieves a GdkFrameTimings object holding timing information for the
// current frame or a recent frame.
//
// The GdkFrameTimings object may not yet be complete: see
// gdk.FrameTimings.GetComplete().
//
// The function takes the following parameters:
//
//    - frameCounter: frame counter value identifying the frame to be received.
//
// The function returns the following values:
//
//    - frameTimings (optional): GdkFrameTimings object for the specified frame,
//      or NULL if it is not available. See gdk.FrameClock.GetHistoryStart().
//
func (frameClock *FrameClock) Timings(frameCounter int64) *FrameTimings {
	var _arg0 *C.GdkFrameClock   // out
	var _arg1 C.gint64           // out
	var _cret *C.GdkFrameTimings // in

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))
	_arg1 = C.gint64(frameCounter)

	_cret = C.gdk_frame_clock_get_timings(_arg0, _arg1)
	runtime.KeepAlive(frameClock)
	runtime.KeepAlive(frameCounter)

	var _frameTimings *FrameTimings // out

	if _cret != nil {
		_frameTimings = (*FrameTimings)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.gdk_frame_timings_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_frameTimings)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_frame_timings_unref((*C.GdkFrameTimings)(intern.C))
			},
		)
	}

	return _frameTimings
}

// RequestPhase asks the frame clock to run a particular phase.
//
// The signal corresponding the requested phase will be emitted the next time
// the frame clock processes. Multiple calls to gdk_frame_clock_request_phase()
// will be combined together and only one frame processed. If you are displaying
// animated content and want to continually request the
// GDK_FRAME_CLOCK_PHASE_UPDATE phase for a period of time, you should use
// gdk.FrameClock.BeginUpdating() instead, since this allows GTK to adjust
// system parameters to get maximally smooth animations.
//
// The function takes the following parameters:
//
//    - phase that is requested.
//
func (frameClock *FrameClock) RequestPhase(phase FrameClockPhase) {
	var _arg0 *C.GdkFrameClock     // out
	var _arg1 C.GdkFrameClockPhase // out

	_arg0 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))
	_arg1 = C.GdkFrameClockPhase(phase)

	C.gdk_frame_clock_request_phase(_arg0, _arg1)
	runtime.KeepAlive(frameClock)
	runtime.KeepAlive(phase)
}
