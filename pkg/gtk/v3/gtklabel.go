// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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

	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// Angle gets the angle of rotation for the label. See
	// gtk_label_set_angle().
	Angle() float64
	// Attributes gets the attribute list that was set on the label using
	// gtk_label_set_attributes(), if any. This function does not reflect
	// attributes that come from the labels markup (see gtk_label_set_markup()).
	// If you want to get the effective attributes for the label, use
	// pango_layout_get_attribute (gtk_label_get_layout (label)).
	Attributes() pango.AttrList
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
	// SelectRegionLabel selects a range of characters in the label, if the
	// label is selectable. See gtk_label_set_selectable(). If the label is not
	// selectable, this function has no effect. If @start_offset or @end_offset
	// are -1, then the end of the label will be substituted.
	SelectRegionLabel(startOffset int, endOffset int)
	// SetAngleLabel sets the angle of rotation for the label. An angle of 90
	// reads from from bottom to top, an angle of 270, from top to bottom. The
	// angle setting for the label is ignored if the label is selectable,
	// wrapped, or ellipsized.
	SetAngleLabel(angle float64)
	// SetAttributesLabel sets a AttrList; the attributes in the list are
	// applied to the label text.
	//
	// The attributes set with this function will be applied and merged with any
	// other attributes previously effected by way of the Label:use-underline or
	// Label:use-markup properties. While it is not recommended to mix markup
	// strings with manually set attributes, if you must; know that the
	// attributes will be applied to the label after the markup string is
	// parsed.
	SetAttributesLabel(attrs pango.AttrList)
	// SetEllipsizeLabel sets the mode used to ellipsize (add an ellipsis:
	// "...") to the text if there is not enough space to render the entire
	// string.
	SetEllipsizeLabel(mode pango.EllipsizeMode)
	// SetJustifyLabel sets the alignment of the lines in the text of the label
	// relative to each other. GTK_JUSTIFY_LEFT is the default value when the
	// widget is first created with gtk_label_new(). If you instead want to set
	// the alignment of the label as a whole, use gtk_widget_set_halign()
	// instead. gtk_label_set_justify() has no effect on labels containing only
	// a single line.
	SetJustifyLabel(jtype Justification)
	// SetLabelLabel sets the text of the label. The label is interpreted as
	// including embedded underlines and/or Pango markup depending on the values
	// of the Label:use-underline and Label:use-markup properties.
	SetLabelLabel(str string)
	// SetLineWrapLabel toggles line wrapping within the Label widget. true
	// makes it break lines if text exceeds the widget’s size. false lets the
	// text get cut off by the edge of the widget if it exceeds the widget size.
	//
	// Note that setting line wrapping to true does not make the label wrap at
	// its parent container’s width, because GTK+ widgets conceptually can’t
	// make their requisition depend on the parent container’s size. For a label
	// that wraps at a specific position, set the label’s width using
	// gtk_widget_set_size_request().
	SetLineWrapLabel(wrap bool)
	// SetLineWrapModeLabel: if line wrapping is on (see
	// gtk_label_set_line_wrap()) this controls how the line wrapping is done.
	// The default is PANGO_WRAP_WORD which means wrap on word boundaries.
	SetLineWrapModeLabel(wrapMode pango.WrapMode)
	// SetLinesLabel sets the number of lines to which an ellipsized, wrapping
	// label should be limited. This has no effect if the label is not wrapping
	// or ellipsized. Set this to -1 if you don’t want to limit the number of
	// lines.
	SetLinesLabel(lines int)
	// SetMarkupLabel parses @str which is marked up with the [Pango text markup
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
	SetMarkupLabel(str string)
	// SetMarkupWithMnemonicLabel parses @str which is marked up with the [Pango
	// text markup language][PangoMarkupFormat], setting the label’s text and
	// attribute list based on the parse results. If characters in @str are
	// preceded by an underscore, they are underlined indicating that they
	// represent a keyboard accelerator called a mnemonic.
	//
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using gtk_label_set_mnemonic_widget().
	SetMarkupWithMnemonicLabel(str string)
	// SetMaxWidthCharsLabel sets the desired maximum width in characters of
	// @label to @n_chars.
	SetMaxWidthCharsLabel(nChars int)
	// SetMnemonicWidgetLabel: if the label has been set so that it has an
	// mnemonic key (using i.e. gtk_label_set_markup_with_mnemonic(),
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
	SetMnemonicWidgetLabel(widget Widget)
	// SetPatternLabel: the pattern of underlines you want under the existing
	// text within the Label widget. For example if the current text of the
	// label says “FooBarBaz” passing a pattern of “___ ___” will underline
	// “Foo” and “Baz” but not “Bar”.
	SetPatternLabel(pattern string)
	// SetSelectableLabel: selectable labels allow the user to select text from
	// the label, for copy-and-paste.
	SetSelectableLabel(setting bool)
	// SetSingleLineModeLabel sets whether the label is in single line mode.
	SetSingleLineModeLabel(singleLineMode bool)
	// SetTextLabel sets the text within the Label widget. It overwrites any
	// text that was there before.
	//
	// This function will clear any previously set mnemonic accelerators, and
	// set the Label:use-underline property to false as a side effect.
	//
	// This function will set the Label:use-markup property to false as a side
	// effect.
	//
	// See also: gtk_label_set_markup()
	SetTextLabel(str string)
	// SetTextWithMnemonicLabel sets the label’s text from the string @str. If
	// characters in @str are preceded by an underscore, they are underlined
	// indicating that they represent a keyboard accelerator called a mnemonic.
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using gtk_label_set_mnemonic_widget().
	SetTextWithMnemonicLabel(str string)
	// SetTrackVisitedLinksLabel sets whether the label should keep track of
	// clicked links (and use a different color for them).
	SetTrackVisitedLinksLabel(trackLinks bool)
	// SetUseMarkupLabel sets whether the text of the label contains markup in
	// [Pango’s text markup language][PangoMarkupFormat]. See
	// gtk_label_set_markup().
	SetUseMarkupLabel(setting bool)
	// SetUseUnderlineLabel: if true, an underline in the text indicates the
	// next character should be used for the mnemonic accelerator key.
	SetUseUnderlineLabel(setting bool)
	// SetWidthCharsLabel sets the desired width in characters of @label to
	// @n_chars.
	SetWidthCharsLabel(nChars int)
	// SetXalignLabel sets the Label:xalign property for @label.
	SetXalignLabel(xalign float32)
	// SetYalignLabel sets the Label:yalign property for @label.
	SetYalignLabel(yalign float32)
}

// label implements the Label class.
type label struct {
	Misc
}

// WrapLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapLabel(obj *externglib.Object) Label {
	return label{
		Misc: WrapMisc(obj),
	}
}

func marshalLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLabel(obj), nil
}

// NewLabel creates a new label with the given text inside it. You can pass nil
// to get an empty label widget.
func NewLabel(str string) Label {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_label_new(_arg1)

	var _label Label // out

	_label = WrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// NewLabelWithMnemonic creates a new Label, containing the text in @str.
//
// If characters in @str are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use '__' (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. The mnemonic key can be used to activate
// another widget, chosen automatically, or explicitly using
// gtk_label_set_mnemonic_widget().
//
// If gtk_label_set_mnemonic_widget() is not called, then the first activatable
// ancestor of the Label will be chosen as the mnemonic widget. For instance, if
// the label is inside a button or menu item, the button or menu item will
// automatically become the mnemonic widget and be activated by the mnemonic.
func NewLabelWithMnemonic(str string) Label {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_label_new_with_mnemonic(_arg1)

	var _label Label // out

	_label = WrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

func (l label) Angle() float64 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_angle(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (l label) Attributes() pango.AttrList {
	var _arg0 *C.GtkLabel      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_attributes(_arg0)

	var _attrList pango.AttrList // out

	_attrList = (pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)

	return _attrList
}

func (l label) CurrentURI() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_current_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (l label) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkLabel          // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

func (l label) Justify() Justification {
	var _arg0 *C.GtkLabel        // out
	var _cret C.GtkJustification // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_justify(_arg0)

	var _justification Justification // out

	_justification = Justification(_cret)

	return _justification
}

func (l label) Label() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (l label) Layout() pango.Layout {
	var _arg0 *C.GtkLabel    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_layout(_arg0)

	var _layout pango.Layout // out

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.Layout)

	return _layout
}

func (l label) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gint     // in
	var _arg2 *C.gint     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	C.gtk_label_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

func (l label) LineWrap() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_line_wrap(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) LineWrapMode() pango.WrapMode {
	var _arg0 *C.GtkLabel     // out
	var _cret C.PangoWrapMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_line_wrap_mode(_arg0)

	var _wrapMode pango.WrapMode // out

	_wrapMode = pango.WrapMode(_cret)

	return _wrapMode
}

func (l label) Lines() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_lines(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (l label) MaxWidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_max_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (l label) MnemonicKeyval() uint {
	var _arg0 *C.GtkLabel // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_mnemonic_keyval(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (l label) MnemonicWidget() Widget {
	var _arg0 *C.GtkLabel  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_mnemonic_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (l label) Selectable() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) SelectionBounds() (start int, end int, ok bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gint     // in
	var _arg2 *C.gint     // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_selection_bounds(_arg0, &_arg1, &_arg2)

	var _start int // out
	var _end int   // out
	var _ok bool   // out

	_start = int(_arg1)
	_end = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _start, _end, _ok
}

func (l label) SingleLineMode() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_single_line_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) Text() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (l label) TrackVisitedLinks() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_track_visited_links(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) UseMarkup() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_use_markup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) UseUnderline() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (l label) WidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (l label) Xalign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_xalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (l label) Yalign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))

	_cret = C.gtk_label_get_yalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (l label) SelectRegionLabel(startOffset int, endOffset int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out
	var _arg2 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)

	C.gtk_label_select_region(_arg0, _arg1, _arg2)
}

func (l label) SetAngleLabel(angle float64) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gdouble(angle)

	C.gtk_label_set_angle(_arg0, _arg1)
}

func (l label) SetAttributesLabel(attrs pango.AttrList) {
	var _arg0 *C.GtkLabel      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_label_set_attributes(_arg0, _arg1)
}

func (l label) SetEllipsizeLabel(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkLabel          // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_label_set_ellipsize(_arg0, _arg1)
}

func (l label) SetJustifyLabel(jtype Justification) {
	var _arg0 *C.GtkLabel        // out
	var _arg1 C.GtkJustification // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.GtkJustification(jtype)

	C.gtk_label_set_justify(_arg0, _arg1)
}

func (l label) SetLabelLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_label(_arg0, _arg1)
}

func (l label) SetLineWrapLabel(wrap bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_line_wrap(_arg0, _arg1)
}

func (l label) SetLineWrapModeLabel(wrapMode pango.WrapMode) {
	var _arg0 *C.GtkLabel     // out
	var _arg1 C.PangoWrapMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.PangoWrapMode(wrapMode)

	C.gtk_label_set_line_wrap_mode(_arg0, _arg1)
}

func (l label) SetLinesLabel(lines int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(lines)

	C.gtk_label_set_lines(_arg0, _arg1)
}

func (l label) SetMarkupLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup(_arg0, _arg1)
}

func (l label) SetMarkupWithMnemonicLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup_with_mnemonic(_arg0, _arg1)
}

func (l label) SetMaxWidthCharsLabel(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_label_set_max_width_chars(_arg0, _arg1)
}

func (l label) SetMnemonicWidgetLabel(widget Widget) {
	var _arg0 *C.GtkLabel  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_label_set_mnemonic_widget(_arg0, _arg1)
}

func (l label) SetPatternLabel(pattern string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_pattern(_arg0, _arg1)
}

func (l label) SetSelectableLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_selectable(_arg0, _arg1)
}

func (l label) SetSingleLineModeLabel(singleLineMode bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if singleLineMode {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_single_line_mode(_arg0, _arg1)
}

func (l label) SetTextLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text(_arg0, _arg1)
}

func (l label) SetTextWithMnemonicLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text_with_mnemonic(_arg0, _arg1)
}

func (l label) SetTrackVisitedLinksLabel(trackLinks bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if trackLinks {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_track_visited_links(_arg0, _arg1)
}

func (l label) SetUseMarkupLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_markup(_arg0, _arg1)
}

func (l label) SetUseUnderlineLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_underline(_arg0, _arg1)
}

func (l label) SetWidthCharsLabel(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_label_set_width_chars(_arg0, _arg1)
}

func (l label) SetXalignLabel(xalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_label_set_xalign(_arg0, _arg1)
}

func (l label) SetYalignLabel(yalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(l.Native()))
	_arg1 = C.gfloat(yalign)

	C.gtk_label_set_yalign(_arg0, _arg1)
}

func (l label) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(l))
}
