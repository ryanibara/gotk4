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
		{T: externglib.Type(C.gtk_label_get_type()), F: marshalLabel},
	})
}

// Label: the Label widget displays a small amount of text. As the name implies,
// most labels are used to label another widget such as a Button, a MenuItem, or
// a ComboBox.
//
// CSS nodes
//
//    const gchar *text =
//    "Go to the"
//    "<a href=\"http://www.gtk.org title=\"&lt;i&gt;Our&lt;/i&gt; website\">"
//    "GTK+ website</a> for more...";
//    GtkWidget *label = gtk_label_new (NULL);
//    gtk_label_set_markup (GTK_LABEL (label), text);
//
// It is possible to implement custom handling for links and their tooltips with
// the Label::activate-link signal and the gtk_label_get_current_uri() function.
type Label interface {
	Misc
	Buildable

	// Angle gets the angle of rotation for the label. See
	// gtk_label_set_angle().
	Angle() float64
	// Attributes gets the attribute list that was set on the label using
	// gtk_label_set_attributes(), if any. This function does not reflect
	// attributes that come from the labels markup (see gtk_label_set_markup()).
	// If you want to get the effective attributes for the label, use
	// pango_layout_get_attribute (gtk_label_get_layout (label)).
	Attributes() *pango.AttrList
	// CurrentURI returns the URI for the currently active link in the label.
	// The active link is the one under the mouse pointer or, in a selectable
	// label, the link in which the text cursor is currently positioned.
	//
	// This function is intended for use in a Label::activate-link handler or
	// for use in a Widget::query-tooltip handler.
	CurrentURI() string
	// Ellipsize returns the ellipsizing position of the label. See
	// gtk_label_set_ellipsize().
	Ellipsize() pango.EllipsizeMode
	// Justify returns the justification of the label. See
	// gtk_label_set_justify().
	Justify() Justification
	// Label fetches the text from a label widget including any embedded
	// underlines indicating mnemonics and Pango markup. (See
	// gtk_label_get_text()).
	Label() string
	// Layout gets the Layout used to display the label. The layout is useful to
	// e.g. convert text positions to pixel positions, in combination with
	// gtk_label_get_layout_offsets(). The returned layout is owned by the
	// @label so need not be freed by the caller. The @label is free to recreate
	// its layout at any time, so it should be considered read-only.
	Layout() pango.Layout
	// LayoutOffsets obtains the coordinates where the label will draw the
	// Layout representing the text in the label; useful to convert mouse events
	// into coordinates inside the Layout, e.g. to take some action if some part
	// of the label is clicked. Of course you will need to create a EventBox to
	// receive the events, and pack the label inside it, since labels are
	// windowless (they return false from gtk_widget_get_has_window()). Remember
	// when using the Layout functions you need to convert to and from pixels
	// using PANGO_PIXELS() or NGO_SCALE.
	LayoutOffsets() (x int, y int)
	// LineWrap returns whether lines in the label are automatically wrapped.
	// See gtk_label_set_line_wrap().
	LineWrap() bool
	// LineWrapMode returns line wrap mode used by the label. See
	// gtk_label_set_line_wrap_mode().
	LineWrapMode() pango.WrapMode
	// Lines gets the number of lines to which an ellipsized, wrapping label
	// should be limited. See gtk_label_set_lines().
	Lines() int
	// MaxWidthChars retrieves the desired maximum width of @label, in
	// characters. See gtk_label_set_width_chars().
	MaxWidthChars() int
	// MnemonicKeyval: if the label has been set so that it has an mnemonic key
	// this function returns the keyval used for the mnemonic accelerator. If
	// there is no mnemonic set up it returns K_KEY_VoidSymbol.
	MnemonicKeyval() uint
	// MnemonicWidget retrieves the target of the mnemonic (keyboard shortcut)
	// of this label. See gtk_label_set_mnemonic_widget().
	MnemonicWidget() Widget
	// Selectable gets the value set by gtk_label_set_selectable().
	Selectable() bool
	// SelectionBounds gets the selected range of characters in the label,
	// returning true if there’s a selection.
	SelectionBounds() (start int, end int, ok bool)
	// SingleLineMode returns whether the label is in single line mode.
	SingleLineMode() bool
	// Text fetches the text from a label widget, as displayed on the screen.
	// This does not include any embedded underlines indicating mnemonics or
	// Pango markup. (See gtk_label_get_label())
	Text() string
	// TrackVisitedLinks returns whether the label is currently keeping track of
	// clicked links.
	TrackVisitedLinks() bool
	// UseMarkup returns whether the label’s text is interpreted as marked up
	// with the [Pango text markup language][PangoMarkupFormat]. See
	// gtk_label_set_use_markup ().
	UseMarkup() bool
	// UseUnderline returns whether an embedded underline in the label indicates
	// a mnemonic. See gtk_label_set_use_underline().
	UseUnderline() bool
	// WidthChars retrieves the desired width of @label, in characters. See
	// gtk_label_set_width_chars().
	WidthChars() int
	// Xalign gets the Label:xalign property for @label.
	Xalign() float32
	// Yalign gets the Label:yalign property for @label.
	Yalign() float32
	// SelectRegion selects a range of characters in the label, if the label is
	// selectable. See gtk_label_set_selectable(). If the label is not
	// selectable, this function has no effect. If @start_offset or @end_offset
	// are -1, then the end of the label will be substituted.
	SelectRegion(startOffset int, endOffset int)
	// SetAngle sets the angle of rotation for the label. An angle of 90 reads
	// from from bottom to top, an angle of 270, from top to bottom. The angle
	// setting for the label is ignored if the label is selectable, wrapped, or
	// ellipsized.
	SetAngle(angle float64)
	// SetAttributes sets a AttrList; the attributes in the list are applied to
	// the label text.
	//
	// The attributes set with this function will be applied and merged with any
	// other attributes previously effected by way of the Label:use-underline or
	// Label:use-markup properties. While it is not recommended to mix markup
	// strings with manually set attributes, if you must; know that the
	// attributes will be applied to the label after the markup string is
	// parsed.
	SetAttributes(attrs *pango.AttrList)
	// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") to
	// the text if there is not enough space to render the entire string.
	SetEllipsize(mode pango.EllipsizeMode)
	// SetJustify sets the alignment of the lines in the text of the label
	// relative to each other. GTK_JUSTIFY_LEFT is the default value when the
	// widget is first created with gtk_label_new(). If you instead want to set
	// the alignment of the label as a whole, use gtk_widget_set_halign()
	// instead. gtk_label_set_justify() has no effect on labels containing only
	// a single line.
	SetJustify(jtype Justification)
	// SetLabel sets the text of the label. The label is interpreted as
	// including embedded underlines and/or Pango markup depending on the values
	// of the Label:use-underline and Label:use-markup properties.
	SetLabel(str string)
	// SetLineWrap toggles line wrapping within the Label widget. true makes it
	// break lines if text exceeds the widget’s size. false lets the text get
	// cut off by the edge of the widget if it exceeds the widget size.
	//
	// Note that setting line wrapping to true does not make the label wrap at
	// its parent container’s width, because GTK+ widgets conceptually can’t
	// make their requisition depend on the parent container’s size. For a label
	// that wraps at a specific position, set the label’s width using
	// gtk_widget_set_size_request().
	SetLineWrap(wrap bool)
	// SetLineWrapMode: if line wrapping is on (see gtk_label_set_line_wrap())
	// this controls how the line wrapping is done. The default is
	// PANGO_WRAP_WORD which means wrap on word boundaries.
	SetLineWrapMode(wrapMode pango.WrapMode)
	// SetLines sets the number of lines to which an ellipsized, wrapping label
	// should be limited. This has no effect if the label is not wrapping or
	// ellipsized. Set this to -1 if you don’t want to limit the number of
	// lines.
	SetLines(lines int)
	// SetMarkup parses @str which is marked up with the [Pango text markup
	// language][PangoMarkupFormat], setting the label’s text and attribute list
	// based on the parse results.
	//
	// If the @str is external data, you may need to escape it with
	// g_markup_escape_text() or g_markup_printf_escaped():
	//
	//    GtkWidget *label = gtk_label_new (NULL);
	//    const char *str = "some text";
	//    const char *format = "<span style=\"italic\">\s</span>";
	//    char *markup;
	//
	//    markup = g_markup_printf_escaped (format, str);
	//    gtk_label_set_markup (GTK_LABEL (label), markup);
	//    g_free (markup);
	//
	// This function will set the Label:use-markup property to true as a side
	// effect.
	//
	// If you set the label contents using the Label:label property you should
	// also ensure that you set the Label:use-markup property accordingly.
	//
	// See also: gtk_label_set_text()
	SetMarkup(str string)
	// SetMarkupWithMnemonic parses @str which is marked up with the [Pango text
	// markup language][PangoMarkupFormat], setting the label’s text and
	// attribute list based on the parse results. If characters in @str are
	// preceded by an underscore, they are underlined indicating that they
	// represent a keyboard accelerator called a mnemonic.
	//
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using gtk_label_set_mnemonic_widget().
	SetMarkupWithMnemonic(str string)
	// SetMaxWidthChars sets the desired maximum width in characters of @label
	// to @n_chars.
	SetMaxWidthChars(nChars int)
	// SetMnemonicWidget: if the label has been set so that it has an mnemonic
	// key (using i.e. gtk_label_set_markup_with_mnemonic(),
	// gtk_label_set_text_with_mnemonic(), gtk_label_new_with_mnemonic() or the
	// “use_underline” property) the label can be associated with a widget that
	// is the target of the mnemonic. When the label is inside a widget (like a
	// Button or a Notebook tab) it is automatically associated with the correct
	// widget, but sometimes (i.e. when the target is a Entry next to the label)
	// you need to set it explicitly using this function.
	//
	// The target widget will be accelerated by emitting the
	// GtkWidget::mnemonic-activate signal on it. The default handler for this
	// signal will activate the widget if there are no mnemonic collisions and
	// toggle focus between the colliding widgets otherwise.
	SetMnemonicWidget(widget Widget)
	// SetPattern: the pattern of underlines you want under the existing text
	// within the Label widget. For example if the current text of the label
	// says “FooBarBaz” passing a pattern of “___ ___” will underline “Foo” and
	// “Baz” but not “Bar”.
	SetPattern(pattern string)
	// SetSelectable: selectable labels allow the user to select text from the
	// label, for copy-and-paste.
	SetSelectable(setting bool)
	// SetSingleLineMode sets whether the label is in single line mode.
	SetSingleLineMode(singleLineMode bool)
	// SetText sets the text within the Label widget. It overwrites any text
	// that was there before.
	//
	// This function will clear any previously set mnemonic accelerators, and
	// set the Label:use-underline property to false as a side effect.
	//
	// This function will set the Label:use-markup property to false as a side
	// effect.
	//
	// See also: gtk_label_set_markup()
	SetText(str string)
	// SetTextWithMnemonic sets the label’s text from the string @str. If
	// characters in @str are preceded by an underscore, they are underlined
	// indicating that they represent a keyboard accelerator called a mnemonic.
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using gtk_label_set_mnemonic_widget().
	SetTextWithMnemonic(str string)
	// SetTrackVisitedLinks sets whether the label should keep track of clicked
	// links (and use a different color for them).
	SetTrackVisitedLinks(trackLinks bool)
	// SetUseMarkup sets whether the text of the label contains markup in
	// [Pango’s text markup language][PangoMarkupFormat]. See
	// gtk_label_set_markup().
	SetUseMarkup(setting bool)
	// SetUseUnderline: if true, an underline in the text indicates the next
	// character should be used for the mnemonic accelerator key.
	SetUseUnderline(setting bool)
	// SetWidthChars sets the desired width in characters of @label to @n_chars.
	SetWidthChars(nChars int)
	// SetXalign sets the Label:xalign property for @label.
	SetXalign(xalign float32)
	// SetYalign sets the Label:yalign property for @label.
	SetYalign(yalign float32)
}

// label implements the Label interface.
type label struct {
	Misc
	Buildable
}

var _ Label = (*label)(nil)

// WrapLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapLabel(obj *externglib.Object) Label {
	return Label{
		Misc:      WrapMisc(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLabel(obj), nil
}

// NewLabel constructs a class Label.
func NewLabel(str string) Label {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkLabel
	var ret1 Label

	cret = C.gtk_label_new(str)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Label)

	return ret1
}

// NewLabelWithMnemonic constructs a class Label.
func NewLabelWithMnemonic(str string) Label {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkLabel
	var ret1 Label

	cret = C.gtk_label_new_with_mnemonic(str)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Label)

	return ret1
}

// Angle gets the angle of rotation for the label. See
// gtk_label_set_angle().
func (l label) Angle() float64 {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gdouble
	var ret1 float64

	cret = C.gtk_label_get_angle(arg0)

	ret1 = C.gdouble(cret)

	return ret1
}

// Attributes gets the attribute list that was set on the label using
// gtk_label_set_attributes(), if any. This function does not reflect
// attributes that come from the labels markup (see gtk_label_set_markup()).
// If you want to get the effective attributes for the label, use
// pango_layout_get_attribute (gtk_label_get_layout (label)).
func (l label) Attributes() *pango.AttrList {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.PangoAttrList
	var ret1 *pango.AttrList

	cret = C.gtk_label_get_attributes(arg0)

	ret1 = pango.WrapAttrList(unsafe.Pointer(cret))

	return ret1
}

// CurrentURI returns the URI for the currently active link in the label.
// The active link is the one under the mouse pointer or, in a selectable
// label, the link in which the text cursor is currently positioned.
//
// This function is intended for use in a Label::activate-link handler or
// for use in a Widget::query-tooltip handler.
func (l label) CurrentURI() string {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_label_get_current_uri(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Ellipsize returns the ellipsizing position of the label. See
// gtk_label_set_ellipsize().
func (l label) Ellipsize() pango.EllipsizeMode {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.PangoEllipsizeMode
	var ret1 pango.EllipsizeMode

	cret = C.gtk_label_get_ellipsize(arg0)

	ret1 = pango.EllipsizeMode(cret)

	return ret1
}

// Justify returns the justification of the label. See
// gtk_label_set_justify().
func (l label) Justify() Justification {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.GtkJustification
	var ret1 Justification

	cret = C.gtk_label_get_justify(arg0)

	ret1 = Justification(cret)

	return ret1
}

// Label fetches the text from a label widget including any embedded
// underlines indicating mnemonics and Pango markup. (See
// gtk_label_get_text()).
func (l label) Label() string {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_label_get_label(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// Layout gets the Layout used to display the label. The layout is useful to
// e.g. convert text positions to pixel positions, in combination with
// gtk_label_get_layout_offsets(). The returned layout is owned by the
// @label so need not be freed by the caller. The @label is free to recreate
// its layout at any time, so it should be considered read-only.
func (l label) Layout() pango.Layout {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.PangoLayout
	var ret1 pango.Layout

	cret = C.gtk_label_get_layout(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(pango.Layout)

	return ret1
}

// LayoutOffsets obtains the coordinates where the label will draw the
// Layout representing the text in the label; useful to convert mouse events
// into coordinates inside the Layout, e.g. to take some action if some part
// of the label is clicked. Of course you will need to create a EventBox to
// receive the events, and pack the label inside it, since labels are
// windowless (they return false from gtk_widget_get_has_window()). Remember
// when using the Layout functions you need to convert to and from pixels
// using PANGO_PIXELS() or NGO_SCALE.
func (l label) LayoutOffsets() (x int, y int) {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var arg1 C.gint
	var ret1 int
	var arg2 C.gint
	var ret2 int

	C.gtk_label_get_layout_offsets(arg0, &arg1, &arg2)

	*ret1 = C.gint(arg1)
	*ret2 = C.gint(arg2)

	return ret1, ret2
}

// LineWrap returns whether lines in the label are automatically wrapped.
// See gtk_label_set_line_wrap().
func (l label) LineWrap() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_line_wrap(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// LineWrapMode returns line wrap mode used by the label. See
// gtk_label_set_line_wrap_mode().
func (l label) LineWrapMode() pango.WrapMode {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.PangoWrapMode
	var ret1 pango.WrapMode

	cret = C.gtk_label_get_line_wrap_mode(arg0)

	ret1 = pango.WrapMode(cret)

	return ret1
}

// Lines gets the number of lines to which an ellipsized, wrapping label
// should be limited. See gtk_label_set_lines().
func (l label) Lines() int {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_label_get_lines(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// MaxWidthChars retrieves the desired maximum width of @label, in
// characters. See gtk_label_set_width_chars().
func (l label) MaxWidthChars() int {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_label_get_max_width_chars(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// MnemonicKeyval: if the label has been set so that it has an mnemonic key
// this function returns the keyval used for the mnemonic accelerator. If
// there is no mnemonic set up it returns K_KEY_VoidSymbol.
func (l label) MnemonicKeyval() uint {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.guint
	var ret1 uint

	cret = C.gtk_label_get_mnemonic_keyval(arg0)

	ret1 = C.guint(cret)

	return ret1
}

// MnemonicWidget retrieves the target of the mnemonic (keyboard shortcut)
// of this label. See gtk_label_set_mnemonic_widget().
func (l label) MnemonicWidget() Widget {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.GtkWidget
	var ret1 Widget

	cret = C.gtk_label_get_mnemonic_widget(arg0)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return ret1
}

// Selectable gets the value set by gtk_label_set_selectable().
func (l label) Selectable() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_selectable(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// SelectionBounds gets the selected range of characters in the label,
// returning true if there’s a selection.
func (l label) SelectionBounds() (start int, end int, ok bool) {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var arg1 C.gint
	var ret1 int
	var arg2 C.gint
	var ret2 int
	var cret C.gboolean
	var ret3 bool

	cret = C.gtk_label_get_selection_bounds(arg0, &arg1, &arg2)

	*ret1 = C.gint(arg1)
	*ret2 = C.gint(arg2)
	ret3 = C.bool(cret) != C.false

	return ret1, ret2, ret3
}

// SingleLineMode returns whether the label is in single line mode.
func (l label) SingleLineMode() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_single_line_mode(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// Text fetches the text from a label widget, as displayed on the screen.
// This does not include any embedded underlines indicating mnemonics or
// Pango markup. (See gtk_label_get_label())
func (l label) Text() string {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret *C.gchar
	var ret1 string

	cret = C.gtk_label_get_text(arg0)

	ret1 = C.GoString(cret)

	return ret1
}

// TrackVisitedLinks returns whether the label is currently keeping track of
// clicked links.
func (l label) TrackVisitedLinks() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_track_visited_links(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// UseMarkup returns whether the label’s text is interpreted as marked up
// with the [Pango text markup language][PangoMarkupFormat]. See
// gtk_label_set_use_markup ().
func (l label) UseMarkup() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_use_markup(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// UseUnderline returns whether an embedded underline in the label indicates
// a mnemonic. See gtk_label_set_use_underline().
func (l label) UseUnderline() bool {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ret1 bool

	cret = C.gtk_label_get_use_underline(arg0)

	ret1 = C.bool(cret) != C.false

	return ret1
}

// WidthChars retrieves the desired width of @label, in characters. See
// gtk_label_set_width_chars().
func (l label) WidthChars() int {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gint
	var ret1 int

	cret = C.gtk_label_get_width_chars(arg0)

	ret1 = C.gint(cret)

	return ret1
}

// Xalign gets the Label:xalign property for @label.
func (l label) Xalign() float32 {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gfloat
	var ret1 float32

	cret = C.gtk_label_get_xalign(arg0)

	ret1 = C.gfloat(cret)

	return ret1
}

// Yalign gets the Label:yalign property for @label.
func (l label) Yalign() float32 {
	var arg0 *C.GtkLabel

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	var cret C.gfloat
	var ret1 float32

	cret = C.gtk_label_get_yalign(arg0)

	ret1 = C.gfloat(cret)

	return ret1
}

// SelectRegion selects a range of characters in the label, if the label is
// selectable. See gtk_label_set_selectable(). If the label is not
// selectable, this function has no effect. If @start_offset or @end_offset
// are -1, then the end of the label will be substituted.
func (l label) SelectRegion(startOffset int, endOffset int) {
	var arg0 *C.GtkLabel
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(startOffset)
	arg2 = C.gint(endOffset)

	C.gtk_label_select_region(arg0, startOffset, endOffset)
}

// SetAngle sets the angle of rotation for the label. An angle of 90 reads
// from from bottom to top, an angle of 270, from top to bottom. The angle
// setting for the label is ignored if the label is selectable, wrapped, or
// ellipsized.
func (l label) SetAngle(angle float64) {
	var arg0 *C.GtkLabel
	var arg1 C.gdouble

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gdouble(angle)

	C.gtk_label_set_angle(arg0, angle)
}

// SetAttributes sets a AttrList; the attributes in the list are applied to
// the label text.
//
// The attributes set with this function will be applied and merged with any
// other attributes previously effected by way of the Label:use-underline or
// Label:use-markup properties. While it is not recommended to mix markup
// strings with manually set attributes, if you must; know that the
// attributes will be applied to the label after the markup string is
// parsed.
func (l label) SetAttributes(attrs *pango.AttrList) {
	var arg0 *C.GtkLabel
	var arg1 *C.PangoAttrList

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs.Native()))

	C.gtk_label_set_attributes(arg0, attrs)
}

// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") to
// the text if there is not enough space to render the entire string.
func (l label) SetEllipsize(mode pango.EllipsizeMode) {
	var arg0 *C.GtkLabel
	var arg1 C.PangoEllipsizeMode

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (C.PangoEllipsizeMode)(mode)

	C.gtk_label_set_ellipsize(arg0, mode)
}

// SetJustify sets the alignment of the lines in the text of the label
// relative to each other. GTK_JUSTIFY_LEFT is the default value when the
// widget is first created with gtk_label_new(). If you instead want to set
// the alignment of the label as a whole, use gtk_widget_set_halign()
// instead. gtk_label_set_justify() has no effect on labels containing only
// a single line.
func (l label) SetJustify(jtype Justification) {
	var arg0 *C.GtkLabel
	var arg1 C.GtkJustification

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (C.GtkJustification)(jtype)

	C.gtk_label_set_justify(arg0, jtype)
}

// SetLabel sets the text of the label. The label is interpreted as
// including embedded underlines and/or Pango markup depending on the values
// of the Label:use-underline and Label:use-markup properties.
func (l label) SetLabel(str string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_label(arg0, str)
}

// SetLineWrap toggles line wrapping within the Label widget. true makes it
// break lines if text exceeds the widget’s size. false lets the text get
// cut off by the edge of the widget if it exceeds the widget size.
//
// Note that setting line wrapping to true does not make the label wrap at
// its parent container’s width, because GTK+ widgets conceptually can’t
// make their requisition depend on the parent container’s size. For a label
// that wraps at a specific position, set the label’s width using
// gtk_widget_set_size_request().
func (l label) SetLineWrap(wrap bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if wrap {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_line_wrap(arg0, wrap)
}

// SetLineWrapMode: if line wrapping is on (see gtk_label_set_line_wrap())
// this controls how the line wrapping is done. The default is
// PANGO_WRAP_WORD which means wrap on word boundaries.
func (l label) SetLineWrapMode(wrapMode pango.WrapMode) {
	var arg0 *C.GtkLabel
	var arg1 C.PangoWrapMode

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (C.PangoWrapMode)(wrapMode)

	C.gtk_label_set_line_wrap_mode(arg0, wrapMode)
}

// SetLines sets the number of lines to which an ellipsized, wrapping label
// should be limited. This has no effect if the label is not wrapping or
// ellipsized. Set this to -1 if you don’t want to limit the number of
// lines.
func (l label) SetLines(lines int) {
	var arg0 *C.GtkLabel
	var arg1 C.gint

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(lines)

	C.gtk_label_set_lines(arg0, lines)
}

// SetMarkup parses @str which is marked up with the [Pango text markup
// language][PangoMarkupFormat], setting the label’s text and attribute list
// based on the parse results.
//
// If the @str is external data, you may need to escape it with
// g_markup_escape_text() or g_markup_printf_escaped():
//
//    GtkWidget *label = gtk_label_new (NULL);
//    const char *str = "some text";
//    const char *format = "<span style=\"italic\">\s</span>";
//    char *markup;
//
//    markup = g_markup_printf_escaped (format, str);
//    gtk_label_set_markup (GTK_LABEL (label), markup);
//    g_free (markup);
//
// This function will set the Label:use-markup property to true as a side
// effect.
//
// If you set the label contents using the Label:label property you should
// also ensure that you set the Label:use-markup property accordingly.
//
// See also: gtk_label_set_text()
func (l label) SetMarkup(str string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_markup(arg0, str)
}

// SetMarkupWithMnemonic parses @str which is marked up with the [Pango text
// markup language][PangoMarkupFormat], setting the label’s text and
// attribute list based on the parse results. If characters in @str are
// preceded by an underscore, they are underlined indicating that they
// represent a keyboard accelerator called a mnemonic.
//
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using gtk_label_set_mnemonic_widget().
func (l label) SetMarkupWithMnemonic(str string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_markup_with_mnemonic(arg0, str)
}

// SetMaxWidthChars sets the desired maximum width in characters of @label
// to @n_chars.
func (l label) SetMaxWidthChars(nChars int) {
	var arg0 *C.GtkLabel
	var arg1 C.gint

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(nChars)

	C.gtk_label_set_max_width_chars(arg0, nChars)
}

// SetMnemonicWidget: if the label has been set so that it has an mnemonic
// key (using i.e. gtk_label_set_markup_with_mnemonic(),
// gtk_label_set_text_with_mnemonic(), gtk_label_new_with_mnemonic() or the
// “use_underline” property) the label can be associated with a widget that
// is the target of the mnemonic. When the label is inside a widget (like a
// Button or a Notebook tab) it is automatically associated with the correct
// widget, but sometimes (i.e. when the target is a Entry next to the label)
// you need to set it explicitly using this function.
//
// The target widget will be accelerated by emitting the
// GtkWidget::mnemonic-activate signal on it. The default handler for this
// signal will activate the widget if there are no mnemonic collisions and
// toggle focus between the colliding widgets otherwise.
func (l label) SetMnemonicWidget(widget Widget) {
	var arg0 *C.GtkLabel
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_label_set_mnemonic_widget(arg0, widget)
}

// SetPattern: the pattern of underlines you want under the existing text
// within the Label widget. For example if the current text of the label
// says “FooBarBaz” passing a pattern of “___ ___” will underline “Foo” and
// “Baz” but not “Bar”.
func (l label) SetPattern(pattern string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_pattern(arg0, pattern)
}

// SetSelectable: selectable labels allow the user to select text from the
// label, for copy-and-paste.
func (l label) SetSelectable(setting bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_selectable(arg0, setting)
}

// SetSingleLineMode sets whether the label is in single line mode.
func (l label) SetSingleLineMode(singleLineMode bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if singleLineMode {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_single_line_mode(arg0, singleLineMode)
}

// SetText sets the text within the Label widget. It overwrites any text
// that was there before.
//
// This function will clear any previously set mnemonic accelerators, and
// set the Label:use-underline property to false as a side effect.
//
// This function will set the Label:use-markup property to false as a side
// effect.
//
// See also: gtk_label_set_markup()
func (l label) SetText(str string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_text(arg0, str)
}

// SetTextWithMnemonic sets the label’s text from the string @str. If
// characters in @str are preceded by an underscore, they are underlined
// indicating that they represent a keyboard accelerator called a mnemonic.
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using gtk_label_set_mnemonic_widget().
func (l label) SetTextWithMnemonic(str string) {
	var arg0 *C.GtkLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_label_set_text_with_mnemonic(arg0, str)
}

// SetTrackVisitedLinks sets whether the label should keep track of clicked
// links (and use a different color for them).
func (l label) SetTrackVisitedLinks(trackLinks bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if trackLinks {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_track_visited_links(arg0, trackLinks)
}

// SetUseMarkup sets whether the text of the label contains markup in
// [Pango’s text markup language][PangoMarkupFormat]. See
// gtk_label_set_markup().
func (l label) SetUseMarkup(setting bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_use_markup(arg0, setting)
}

// SetUseUnderline: if true, an underline in the text indicates the next
// character should be used for the mnemonic accelerator key.
func (l label) SetUseUnderline(setting bool) {
	var arg0 *C.GtkLabel
	var arg1 C.gboolean

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_label_set_use_underline(arg0, setting)
}

// SetWidthChars sets the desired width in characters of @label to @n_chars.
func (l label) SetWidthChars(nChars int) {
	var arg0 *C.GtkLabel
	var arg1 C.gint

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gint(nChars)

	C.gtk_label_set_width_chars(arg0, nChars)
}

// SetXalign sets the Label:xalign property for @label.
func (l label) SetXalign(xalign float32) {
	var arg0 *C.GtkLabel
	var arg1 C.gfloat

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gfloat(xalign)

	C.gtk_label_set_xalign(arg0, xalign)
}

// SetYalign sets the Label:yalign property for @label.
func (l label) SetYalign(yalign float32) {
	var arg0 *C.GtkLabel
	var arg1 C.gfloat

	arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	arg1 = C.gfloat(yalign)

	C.gtk_label_set_yalign(arg0, yalign)
}

type LabelPrivate struct {
	native C.GtkLabelPrivate
}

// WrapLabelPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLabelPrivate(ptr unsafe.Pointer) *LabelPrivate {
	if ptr == nil {
		return nil
	}

	return (*LabelPrivate)(ptr)
}

func marshalLabelPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapLabelPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *LabelPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

type LabelSelectionInfo struct {
	native C.GtkLabelSelectionInfo
}

// WrapLabelSelectionInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLabelSelectionInfo(ptr unsafe.Pointer) *LabelSelectionInfo {
	if ptr == nil {
		return nil
	}

	return (*LabelSelectionInfo)(ptr)
}

func marshalLabelSelectionInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapLabelSelectionInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *LabelSelectionInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}
