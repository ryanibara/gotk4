// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_progress_bar_get_type()), F: marshalProgressBar},
	})
}

// ProgressBar: `GtkProgressBar` is typically used to display the progress of a
// long running operation.
//
// It provides a visual clue that processing is underway. `GtkProgressBar` can
// be used in two different modes: percentage mode and activity mode.
//
// !An example GtkProgressBar (progressbar.png)
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the `GtkProgressBar` in percentage mode and the user sees a growing
// bar indicating the percentage of the work that has been completed. In this
// mode, the application is required to call
// [method@Gtk.ProgressBar.set_fraction] periodically to update the progress
// bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the `GtkProgressBar` in activity mode, which shows activity by a
// block moving back and forth within the progress area. In this mode, the
// application is required to call [method@Gtk.ProgressBar.pulse] periodically
// to update the progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// `GtkProgressBar`. Functions are provided to control the orientation of the
// bar, optional text can be displayed along with the bar, and the step size
// used in activity mode can be set.
//
//
// CSS nodes
//
// “` progressbar[.osd] ├── [text] ╰── trough[.empty][.full] ╰──
// progress[.pulse] “`
//
// `GtkProgressBar` has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
//
//
// Accessibility
//
// `GtkProgressBar` uses the K_ACCESSIBLE_ROLE_PROGRESS_BAR role.
type ProgressBar interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget
	// AsOrientable casts the class to the Orientable interface.
	AsOrientable() Orientable

	// Ellipsize returns the ellipsizing position of the progress bar.
	//
	// See [method@Gtk.ProgressBar.set_ellipsize].
	Ellipsize() pango.EllipsizeMode
	// Fraction returns the current fraction of the task that’s been completed.
	Fraction() float64
	// Inverted returns whether the progress bar is inverted.
	Inverted() bool
	// PulseStep retrieves the pulse step.
	//
	// See [method@Gtk.ProgressBar.set_pulse_step].
	PulseStep() float64
	// ShowText returns whether the `GtkProgressBar` shows text.
	//
	// See [method@Gtk.ProgressBar.set_show_text].
	ShowText() bool
	// Text retrieves the text that is displayed with the progress bar.
	//
	// The return value is a reference to the text, not a copy of it, so will
	// become invalid if you change the text in the progress bar.
	Text() string
	// PulseProgressBar indicates that some progress has been made, but you
	// don’t know how much.
	//
	// Causes the progress bar to enter “activity mode,” where a block bounces
	// back and forth. Each call to [method@Gtk.ProgressBar.pulse] causes the
	// block to move by a little bit (the amount of movement per pulse is
	// determined by [method@Gtk.ProgressBar.set_pulse_step]).
	PulseProgressBar()
	// SetEllipsizeProgressBar sets the mode used to ellipsize the text.
	//
	// The text is ellipsized if there is not enough space to render the entire
	// string.
	SetEllipsizeProgressBar(mode pango.EllipsizeMode)
	// SetFractionProgressBar causes the progress bar to “fill in” the given
	// fraction of the bar.
	//
	// The fraction should be between 0.0 and 1.0, inclusive.
	SetFractionProgressBar(fraction float64)
	// SetInvertedProgressBar sets whether the progress bar is inverted.
	//
	// Progress bars normally grow from top to bottom or left to right. Inverted
	// progress bars grow in the opposite direction.
	SetInvertedProgressBar(inverted bool)
	// SetPulseStepProgressBar sets the fraction of total progress bar length to
	// move the bouncing block.
	//
	// The bouncing block is moved when [method@Gtk.ProgressBar.pulse] is
	// called.
	SetPulseStepProgressBar(fraction float64)
	// SetShowTextProgressBar sets whether the progress bar will show text next
	// to the bar.
	//
	// The shown text is either the value of the [property@Gtk.ProgressBar:text]
	// property or, if that is nil, the [property@Gtk.ProgressBar:fraction]
	// value, as a percentage.
	//
	// To make a progress bar that is styled and sized suitably for containing
	// text (even if the actual text is blank), set
	// [property@Gtk.ProgressBar:show-text] to true and
	// [property@Gtk.ProgressBar:text] to the empty string (not nil).
	SetShowTextProgressBar(showText bool)
	// SetTextProgressBar causes the given @text to appear next to the progress
	// bar.
	//
	// If @text is nil and [property@Gtk.ProgressBar:show-text] is true, the
	// current value of [property@Gtk.ProgressBar:fraction] will be displayed as
	// a percentage.
	//
	// If @text is non-nil and [property@Gtk.ProgressBar:show-text] is true, the
	// text will be displayed. In this case, it will not display the progress
	// percentage. If @text is the empty string, the progress bar will still be
	// styled and sized suitably for containing text, as long as
	// [property@Gtk.ProgressBar:show-text] is true.
	SetTextProgressBar(text string)
}

// progressBar implements the ProgressBar class.
type progressBar struct {
	Widget
}

// WrapProgressBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapProgressBar(obj *externglib.Object) ProgressBar {
	return progressBar{
		Widget: WrapWidget(obj),
	}
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProgressBar(obj), nil
}

// NewProgressBar creates a new `GtkProgressBar`.
func NewProgressBar() ProgressBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_progress_bar_new()

	var _progressBar ProgressBar // out

	_progressBar = WrapProgressBar(externglib.Take(unsafe.Pointer(_cret)))

	return _progressBar
}

func (p progressBar) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkProgressBar    // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

func (p progressBar) Fraction() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.double          // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (p progressBar) Inverted() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p progressBar) PulseStep() float64 {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.double          // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (p progressBar) ShowText() bool {
	var _arg0 *C.GtkProgressBar // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_show_text(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p progressBar) Text() string {
	var _arg0 *C.GtkProgressBar // out
	var _cret *C.char           // in

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	_cret = C.gtk_progress_bar_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (p progressBar) PulseProgressBar() {
	var _arg0 *C.GtkProgressBar // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	C.gtk_progress_bar_pulse(_arg0)
}

func (p progressBar) SetEllipsizeProgressBar(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkProgressBar    // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_progress_bar_set_ellipsize(_arg0, _arg1)
}

func (p progressBar) SetFractionProgressBar(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.double(fraction)

	C.gtk_progress_bar_set_fraction(_arg0, _arg1)
}

func (p progressBar) SetInvertedProgressBar(inverted bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_inverted(_arg0, _arg1)
}

func (p progressBar) SetPulseStepProgressBar(fraction float64) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.double          // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = C.double(fraction)

	C.gtk_progress_bar_set_pulse_step(_arg0, _arg1)
}

func (p progressBar) SetShowTextProgressBar(showText bool) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if showText {
		_arg1 = C.TRUE
	}

	C.gtk_progress_bar_set_show_text(_arg0, _arg1)
}

func (p progressBar) SetTextProgressBar(text string) {
	var _arg0 *C.GtkProgressBar // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_progress_bar_set_text(_arg0, _arg1)
}

func (p progressBar) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(p))
}

func (p progressBar) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(p))
}

func (p progressBar) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(p))
}

func (p progressBar) AsOrientable() Orientable {
	return WrapOrientable(gextras.InternObject(p))
}
