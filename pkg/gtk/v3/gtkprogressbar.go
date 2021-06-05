// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_progress_bar_get_type()), F: marshalProgressBar},
	})
}

// ProgressBar: the ProgressBar is typically used to display the progress of a
// long running operation. It provides a visual clue that processing is
// underway. The GtkProgressBar can be used in two different modes: percentage
// mode and activity mode.
//
// When an application can determine how much work needs to take place (e.g.
// read a fixed number of bytes from a file) and can monitor its progress, it
// can use the GtkProgressBar in percentage mode and the user sees a growing bar
// indicating the percentage of the work that has been completed. In this mode,
// the application is required to call gtk_progress_bar_set_fraction()
// periodically to update the progress bar.
//
// When an application has no accurate way of knowing the amount of work to do,
// it can use the ProgressBar in activity mode, which shows activity by a block
// moving back and forth within the progress area. In this mode, the application
// is required to call gtk_progress_bar_pulse() periodically to update the
// progress bar.
//
// There is quite a bit of flexibility provided to control the appearance of the
// ProgressBar. Functions are provided to control the orientation of the bar,
// optional text can be displayed along with the bar, and the step size used in
// activity mode can be set.
//
// CSS nodes
//
//    progressbar[.osd]
//    ├── [text]
//    ╰── trough[.empty][.full]
//        ╰── progress[.pulse]
//
// GtkProgressBar has a main CSS node with name progressbar and subnodes with
// names text and trough, of which the latter has a subnode named progress. The
// text subnode is only present if text is shown. The progress subnode has the
// style class .pulse when in activity mode. It gets the style classes .left,
// .right, .top or .bottom added when the progress 'touches' the corresponding
// end of the GtkProgressBar. The .osd class on the progressbar node is for use
// in overlays like the one Epiphany has for page loading progress.
type ProgressBar interface {
	Widget
	Buildable
	Orientable

	// Ellipsize returns the ellipsizing position of the progress bar. See
	// gtk_progress_bar_set_ellipsize().
	Ellipsize() pango.EllipsizeMode
	// Fraction returns the current fraction of the task that’s been completed.
	Fraction() float64
	// Inverted gets the value set by gtk_progress_bar_set_inverted().
	Inverted() bool
	// PulseStep retrieves the pulse step set with
	// gtk_progress_bar_set_pulse_step().
	PulseStep() float64
	// ShowText gets the value of the ProgressBar:show-text property. See
	// gtk_progress_bar_set_show_text().
	ShowText() bool
	// Text retrieves the text that is displayed with the progress bar, if any,
	// otherwise nil. The return value is a reference to the text, not a copy of
	// it, so will become invalid if you change the text in the progress bar.
	Text() string
	// Pulse indicates that some progress has been made, but you don’t know how
	// much. Causes the progress bar to enter “activity mode,” where a block
	// bounces back and forth. Each call to gtk_progress_bar_pulse() causes the
	// block to move by a little bit (the amount of movement per pulse is
	// determined by gtk_progress_bar_set_pulse_step()).
	Pulse()
	// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") the
	// text if there is not enough space to render the entire string.
	SetEllipsize(mode pango.EllipsizeMode)
	// SetFraction causes the progress bar to “fill in” the given fraction of
	// the bar. The fraction should be between 0.0 and 1.0, inclusive.
	SetFraction(fraction float64)
	// SetInverted progress bars normally grow from top to bottom or left to
	// right. Inverted progress bars grow in the opposite direction.
	SetInverted(inverted bool)
	// SetPulseStep sets the fraction of total progress bar length to move the
	// bouncing block for each call to gtk_progress_bar_pulse().
	SetPulseStep(fraction float64)
	// SetShowText sets whether the progress bar will show text next to the bar.
	// The shown text is either the value of the ProgressBar:text property or,
	// if that is nil, the ProgressBar:fraction value, as a percentage.
	//
	// To make a progress bar that is styled and sized suitably for containing
	// text (even if the actual text is blank), set ProgressBar:show-text to
	// true and ProgressBar:text to the empty string (not nil).
	SetShowText(showText bool)
	// SetText causes the given @text to appear next to the progress bar.
	//
	// If @text is nil and ProgressBar:show-text is true, the current value of
	// ProgressBar:fraction will be displayed as a percentage.
	//
	// If @text is non-nil and ProgressBar:show-text is true, the text will be
	// displayed. In this case, it will not display the progress percentage. If
	// @text is the empty string, the progress bar will still be styled and
	// sized suitably for containing text, as long as ProgressBar:show-text is
	// true.
	SetText(text string)
}

// progressBar implements the ProgressBar interface.
type progressBar struct {
	Widget
	Buildable
	Orientable
}

var _ ProgressBar = (*progressBar)(nil)

// WrapProgressBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapProgressBar(obj *externglib.Object) ProgressBar {
	return ProgressBar{
		Widget:     WrapWidget(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalProgressBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProgressBar(obj), nil
}

// NewProgressBar constructs a class ProgressBar.
func NewProgressBar() ProgressBar {
	var cret C.GtkProgressBar
	var ret1 ProgressBar

	cret = C.gtk_progress_bar_new()

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ProgressBar)

	return ret1
}

// Ellipsize returns the ellipsizing position of the progress bar. See
// gtk_progress_bar_set_ellipsize().
func (p progressBar) Ellipsize() pango.EllipsizeMode {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret C.PangoEllipsizeMode
	var ret1 pango.EllipsizeMode

	cret = C.gtk_progress_bar_get_ellipsize(arg0)

	ret1 = pango.EllipsizeMode(cret)

	return ret1
}

// Fraction returns the current fraction of the task that’s been completed.
func (p progressBar) Fraction() float64 {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_progress_bar_get_fraction(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// Inverted gets the value set by gtk_progress_bar_set_inverted().
func (p progressBar) Inverted() bool {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_progress_bar_get_inverted(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// PulseStep retrieves the pulse step set with
// gtk_progress_bar_set_pulse_step().
func (p progressBar) PulseStep() float64 {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_progress_bar_get_pulse_step(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// ShowText gets the value of the ProgressBar:show-text property. See
// gtk_progress_bar_set_show_text().
func (p progressBar) ShowText() bool {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_progress_bar_get_show_text(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Text retrieves the text that is displayed with the progress bar, if any,
// otherwise nil. The return value is a reference to the text, not a copy of
// it, so will become invalid if you change the text in the progress bar.
func (p progressBar) Text() string {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_progress_bar_get_text(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Pulse indicates that some progress has been made, but you don’t know how
// much. Causes the progress bar to enter “activity mode,” where a block
// bounces back and forth. Each call to gtk_progress_bar_pulse() causes the
// block to move by a little bit (the amount of movement per pulse is
// determined by gtk_progress_bar_set_pulse_step()).
func (p progressBar) Pulse() {
	var arg0 *C.GtkProgressBar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))

	C.gtk_progress_bar_pulse(arg0)
}

// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") the
// text if there is not enough space to render the entire string.
func (p progressBar) SetEllipsize(mode pango.EllipsizeMode) {
	var arg0 *C.GtkProgressBar
	var arg1 C.PangoEllipsizeMode

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	arg1 = (C.PangoEllipsizeMode)(mode)

	C.gtk_progress_bar_set_ellipsize(arg0, mode)
}

// SetFraction causes the progress bar to “fill in” the given fraction of
// the bar. The fraction should be between 0.0 and 1.0, inclusive.
func (p progressBar) SetFraction(fraction float64) {
	var arg0 *C.GtkProgressBar
	var arg1 C.gdouble

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_fraction(arg0, fraction)
}

// SetInverted progress bars normally grow from top to bottom or left to
// right. Inverted progress bars grow in the opposite direction.
func (p progressBar) SetInverted(inverted bool) {
	var arg0 *C.GtkProgressBar
	var arg1 C.gboolean

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if inverted {
		arg1 = C.gboolean(1)
	}

	C.gtk_progress_bar_set_inverted(arg0, inverted)
}

// SetPulseStep sets the fraction of total progress bar length to move the
// bouncing block for each call to gtk_progress_bar_pulse().
func (p progressBar) SetPulseStep(fraction float64) {
	var arg0 *C.GtkProgressBar
	var arg1 C.gdouble

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	arg1 = C.gdouble(fraction)

	C.gtk_progress_bar_set_pulse_step(arg0, fraction)
}

// SetShowText sets whether the progress bar will show text next to the bar.
// The shown text is either the value of the ProgressBar:text property or,
// if that is nil, the ProgressBar:fraction value, as a percentage.
//
// To make a progress bar that is styled and sized suitably for containing
// text (even if the actual text is blank), set ProgressBar:show-text to
// true and ProgressBar:text to the empty string (not nil).
func (p progressBar) SetShowText(showText bool) {
	var arg0 *C.GtkProgressBar
	var arg1 C.gboolean

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	if showText {
		arg1 = C.gboolean(1)
	}

	C.gtk_progress_bar_set_show_text(arg0, showText)
}

// SetText causes the given @text to appear next to the progress bar.
//
// If @text is nil and ProgressBar:show-text is true, the current value of
// ProgressBar:fraction will be displayed as a percentage.
//
// If @text is non-nil and ProgressBar:show-text is true, the text will be
// displayed. In this case, it will not display the progress percentage. If
// @text is the empty string, the progress bar will still be styled and
// sized suitably for containing text, as long as ProgressBar:show-text is
// true.
func (p progressBar) SetText(text string) {
	var arg0 *C.GtkProgressBar
	var arg1 *C.gchar

	arg0 = (*C.GtkProgressBar)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_progress_bar_set_text(arg0, text)
}
