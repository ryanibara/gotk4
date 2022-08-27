// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// NewExpander creates a new expander using label as the text of the label.
//
// The function takes the following parameters:
//
//    - label (optional): text of the label.
//
// The function returns the following values:
//
//    - expander: new Expander widget.
//
func NewExpander(label string) *Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_expander_new(_arg1)
	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(coreglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// NewExpanderWithMnemonic creates a new expander using label as the text of the
// label. If characters in label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use “__”
// (two underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. Pressing Alt and that key activates the
// button.
//
// The function takes the following parameters:
//
//    - label (optional): text of the label with an underscore in front of the
//      mnemonic character.
//
// The function returns the following values:
//
//    - expander: new Expander widget.
//
func NewExpanderWithMnemonic(label string) *Expander {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_expander_new_with_mnemonic(_arg1)
	runtime.KeepAlive(label)

	var _expander *Expander // out

	_expander = wrapExpander(coreglib.Take(unsafe.Pointer(_cret)))

	return _expander
}

// Expanded queries a Expander and returns its current state. Returns TRUE if
// the child widget is revealed.
//
// See gtk_expander_set_expanded().
//
// The function returns the following values:
//
//    - ok: current state of the expander.
//
func (expander *Expander) Expanded() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_expanded(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Label fetches the text from a label widget including any embedded underlines
// indicating mnemonics and Pango markup, as set by gtk_expander_set_label(). If
// the label text has not been set the return value will be NULL. This will be
// the case if you create an empty button with gtk_button_new() to use as a
// container.
//
// Note that this function behaved differently in versions prior to 2.14 and
// used to return the label text stripped of embedded underlines indicating
// mnemonics and Pango markup. This problem can be avoided by fetching the label
// text directly from the label widget.
//
// The function returns the following values:
//
//    - utf8 (optional): text of the label widget. This string is owned by the
//      widget and must not be modified or freed.
//
func (expander *Expander) Label() string {
	var _arg0 *C.GtkExpander // out
	var _cret *C.gchar       // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_label(_arg0)
	runtime.KeepAlive(expander)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
//
// The function returns the following values:
//
//    - widget (optional): label widget, or NULL if there is none.
//
func (expander *Expander) LabelWidget() Widgetter {
	var _arg0 *C.GtkExpander // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_label_widget(_arg0)
	runtime.KeepAlive(expander)

	var _widget Widgetter // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Widgetter)
				return ok
			})
			rv, ok := casted.(Widgetter)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
			}
			_widget = rv
		}
	}

	return _widget
}

// Spacing gets the value set by gtk_expander_set_spacing().
//
// Deprecated: Use margins on the child instead.
//
// The function returns the following values:
//
//    - gint: spacing between the expander and child.
//
func (expander *Expander) Spacing() int {
	var _arg0 *C.GtkExpander // out
	var _cret C.gint         // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_spacing(_arg0)
	runtime.KeepAlive(expander)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// UseMarkup returns whether the label’s text is interpreted as marked up with
// the [Pango text markup language][PangoMarkupFormat]. See
// gtk_expander_set_use_markup().
//
// The function returns the following values:
//
//    - ok: TRUE if the label’s text will be parsed for markup.
//
func (expander *Expander) UseMarkup() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_use_markup(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
//
// The function returns the following values:
//
//    - ok: TRUE if an embedded underline in the expander label indicates the
//      mnemonic accelerator keys.
//
func (expander *Expander) UseUnderline() bool {
	var _arg0 *C.GtkExpander // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))

	_cret = C.gtk_expander_get_use_underline(_arg0)
	runtime.KeepAlive(expander)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExpanded sets the state of the expander. Set to TRUE, if you want the
// child widget to be revealed, and FALSE if you want the child widget to be
// hidden.
//
// The function takes the following parameters:
//
//    - expanded: whether the child widget is revealed.
//
func (expander *Expander) SetExpanded(expanded bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if expanded {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_expanded(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(expanded)
}

// SetLabel sets the text of the label of the expander to label.
//
// This will also clear any previously set labels.
//
// The function takes the following parameters:
//
//    - label (optional): string.
//
func (expander *Expander) SetLabel(label string) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if label != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_expander_set_label(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(label)
}

// SetLabelWidget: set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
//
// The function takes the following parameters:
//
//    - labelWidget (optional): new label widget.
//
func (expander *Expander) SetLabelWidget(labelWidget Widgetter) {
	var _arg0 *C.GtkExpander // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if labelWidget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(coreglib.InternObject(labelWidget).Native()))
	}

	C.gtk_expander_set_label_widget(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(labelWidget)
}

// SetSpacing sets the spacing field of expander, which is the number of pixels
// to place between expander and the child.
//
// Deprecated: Use margins on the child instead.
//
// The function takes the following parameters:
//
//    - spacing: distance between the expander and child in pixels.
//
func (expander *Expander) SetSpacing(spacing int) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gint         // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	_arg1 = C.gint(spacing)

	C.gtk_expander_set_spacing(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(spacing)
}

// SetUseMarkup sets whether the text of the label contains markup in [Pango’s
// text markup language][PangoMarkupFormat]. See gtk_label_set_markup().
//
// The function takes the following parameters:
//
//    - useMarkup: TRUE if the label’s text should be parsed for markup.
//
func (expander *Expander) SetUseMarkup(useMarkup bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if useMarkup {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_markup(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(useMarkup)
}

// SetUseUnderline: if true, an underline in the text of the expander label
// indicates the next character should be used for the mnemonic accelerator key.
//
// The function takes the following parameters:
//
//    - useUnderline: TRUE if underlines in the text indicate mnemonics.
//
func (expander *Expander) SetUseUnderline(useUnderline bool) {
	var _arg0 *C.GtkExpander // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkExpander)(unsafe.Pointer(coreglib.InternObject(expander).Native()))
	if useUnderline {
		_arg1 = C.TRUE
	}

	C.gtk_expander_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(expander)
	runtime.KeepAlive(useUnderline)
}
