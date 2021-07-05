// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_label_get_type()), F: marshalLabel},
	})
}

// Label: the `GtkLabel` widget displays a small amount of text.
//
// As the name implies, most labels are used to label another widget such as a
// [class@Button].
//
// !An example GtkLabel (label.png)
//
//
// CSS nodes
//
// “` label ├── [selection] ├── [link] ┊ ╰── [link] “`
//
// `GtkLabel` has a single CSS node with the name label. A wide variety of style
// classes may be applied to labels, such as .title, .subtitle, .dim-label, etc.
// In the `GtkShortcutsWindow`, labels are used with the .keycap style class.
//
// If the label has a selection, it gets a subnode with name selection.
//
// If the label has links, there is one subnode per link. These subnodes carry
// the link or visited state depending on whether they have been visited. In
// this case, label node also gets a .link style class.
//
//
// GtkLabel as GtkBuildable
//
// The GtkLabel implementation of the GtkBuildable interface supports a custom
// <attributes> element, which supports any number of <attribute> elements. The
// <attribute> element has attributes named “name“, “value“, “start“ and “end“
// and allows you to specify [struct@Pango.Attribute] values for this label.
//
// An example of a UI definition fragment specifying Pango attributes: “`xml
// <object class="GtkLabel"> <attributes> <attribute name="weight"
// value="PANGO_WEIGHT_BOLD"/> <attribute name="background" value="red"
// start="5" end="10"/> </attributes> </object> “`
//
// The start and end attributes specify the range of characters to which the
// Pango attribute applies. If start and end are not specified, the attribute is
// applied to the whole text. Note that specifying ranges does not make much
// sense with translatable attributes. Use markup embedded in the translatable
// content instead.
//
//
// Accessibility
//
// `GtkLabel` uses the K_ACCESSIBLE_ROLE_LABEL role.
//
//
// Mnemonics
//
// Labels may contain “mnemonics”. Mnemonics are underlined characters in the
// label, used for keyboard navigation. Mnemonics are created by providing a
// string with an underscore before the mnemonic character, such as `"_File"`,
// to the functions [ctor@Gtk.Label.new_with_mnemonic] or
// [method@Gtk.Label.set_text_with_mnemonic].
//
// Mnemonics automatically activate any activatable widget the label is inside,
// such as a [class@Gtk.Button]; if the label is not inside the mnemonic’s
// target widget, you have to tell the label about the target using
// [class@Gtk.Label.set_mnemonic_widget]. Here’s a simple example where the
// label is inside a button:
//
// “`c // Pressing Alt+H will activate this button GtkWidget *button =
// gtk_button_new (); GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
// gtk_button_set_child (GTK_BUTTON (button), label); “`
//
// There’s a convenience function to create buttons with a mnemonic label
// already inside:
//
// “`c // Pressing Alt+H will activate this button GtkWidget *button =
// gtk_button_new_with_mnemonic ("_Hello"); “`
//
// To create a mnemonic for a widget alongside the label, such as a
// [class@Gtk.Entry], you have to point the label at the entry with
// [method@Gtk.Label.set_mnemonic_widget]:
//
// “`c // Pressing Alt+H will focus the entry GtkWidget *entry = gtk_entry_new
// (); GtkWidget *label = gtk_label_new_with_mnemonic ("_Hello");
// gtk_label_set_mnemonic_widget (GTK_LABEL (label), entry); “`
//
// Markup (styled text)
//
// To make it easy to format text in a label (changing colors, fonts, etc.),
// label text can be provided in a simple markup format:
//
// Here’s how to create a label with a small font: “`c GtkWidget *label =
// gtk_label_new (NULL); gtk_label_set_markup (GTK_LABEL (label), "<small>Small
// text</small>"); “`
//
// (See the Pango manual for complete documentation] of available tags,
// [func@Pango.parse_markup])
//
// The markup passed to gtk_label_set_markup() must be valid; for example,
// literal <, > and & characters must be escaped as &lt;, &gt;, and &amp;. If
// you pass text obtained from the user, file, or a network to
// [method@Gtk.Label.set_markup], you’ll want to escape it with
// g_markup_escape_text() or g_markup_printf_escaped().
//
// Markup strings are just a convenient way to set the [struct@Pango.AttrList]
// on a label; [method@Gtk.Label.set_attributes] may be a simpler way to set
// attributes in some cases. Be careful though; [struct@Pango.AttrList] tends to
// cause internationalization problems, unless you’re applying attributes to the
// entire string (i.e. unless you set the range of each attribute to [0,
// G_MAXINT)). The reason is that specifying the start_index and end_index for a
// [struct@Pango.Attribute] requires knowledge of the exact string being
// displayed, so translations will cause problems.
//
//
// Selectable labels
//
// Labels can be made selectable with [method@Gtk.Label.set_selectable].
// Selectable labels allow the user to copy the label contents to the clipboard.
// Only labels that contain useful-to-copy information — such as error messages
// — should be made selectable.
//
//
// Text layout
//
// A label can contain any number of paragraphs, but will have performance
// problems if it contains more than a small number. Paragraphs are separated by
// newlines or other paragraph separators understood by Pango.
//
// Labels can automatically wrap text if you call [method@Gtk.Label.set_wrap].
//
// [method@Gtk.Label.set_justify] sets how the lines in a label align with one
// another. If you want to set how the label as a whole aligns in its available
// space, see the [property@Gtk.Widget:halign] and [property@Gtk.Widget:valign]
// properties.
//
// The [property@Gtk.Label:width-chars] and [property@Gtk.Label:max-width-chars]
// properties can be used to control the size allocation of ellipsized or
// wrapped labels. For ellipsizing labels, if either is specified (and less than
// the actual text size), it is used as the minimum width, and the actual text
// size is used as the natural width of the label. For wrapping labels,
// width-chars is used as the minimum width, if specified, and max-width-chars
// is used as the natural width. Even if max-width-chars specified, wrapping
// labels will be rewrapped to use all of the available width.
//
//
// Links
//
// GTK supports markup for clickable hyperlinks in addition to regular Pango
// markup. The markup for links is borrowed from HTML, using the `<a>` with
// “href“, “title“ and “class“ attributes. GTK renders links similar to the way
// they appear in web browsers, with colored, underlined text. The “title“
// attribute is displayed as a tooltip on the link. The “class“ attribute is
// used as style class on the CSS node for the link.
//
// An example looks like this:
//
// “`c const char *text = "Go to the" "<a href=\"http://www.gtk.org
// title=\"&lt;i&gt;Our&lt;/i&gt; website\">" "GTK website</a> for more...";
// GtkWidget *label = gtk_label_new (NULL); gtk_label_set_markup (GTK_LABEL
// (label), text); “`
//
// It is possible to implement custom handling for links and their tooltips with
// the [signal@Gtk.Label::activate-link] signal and the
// [method@Gtk.Label.get_current_uri] function.
type Label interface {
	Widget

	// AsAccessible casts the class to the Accessible interface.
	AsAccessible() Accessible
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable
	// AsConstraintTarget casts the class to the ConstraintTarget interface.
	AsConstraintTarget() ConstraintTarget

	// Attributes gets the labels attribute list.
	//
	// This is the [struct@Pango.AttrList] that was set on the label using
	// [method@Gtk.Label.set_attributes], if any. This function does not reflect
	// attributes that come from the labels markup (see
	// [method@Gtk.Label.set_markup]). If you want to get the effective
	// attributes for the label, use `pango_layout_get_attribute
	// (gtk_label_get_layout (self))`.
	Attributes() pango.AttrList
	// CurrentURI returns the URI for the currently active link in the label.
	//
	// The active link is the one under the mouse pointer or, in a selectable
	// label, the link in which the text cursor is currently positioned.
	//
	// This function is intended for use in a [signal@Gtk.Label::activate-link]
	// handler or for use in a [signal@Gtk.Widget::query-tooltip] handler.
	CurrentURI() string
	// Ellipsize returns the ellipsizing position of the label.
	//
	// See [method@Gtk.Label.set_ellipsize].
	Ellipsize() pango.EllipsizeMode
	// ExtraMenu gets the extra menu model of @label.
	//
	// See [method@Gtk.Label.set_extra_menu].
	ExtraMenu() gio.MenuModel
	// Justify returns the justification of the label.
	//
	// See [method@Gtk.Label.set_justify].
	Justify() Justification
	// Label fetches the text from a label.
	//
	// The returned text includes any embedded underlines indicating mnemonics
	// and Pango markup. (See [method@Gtk.Label.get_text]).
	Label() string
	// Layout gets the `PangoLayout` used to display the label.
	//
	// The layout is useful to e.g. convert text positions to pixel positions,
	// in combination with [method@Gtk.Label.get_layout_offsets]. The returned
	// layout is owned by the @label so need not be freed by the caller. The
	// @label is free to recreate its layout at any time, so it should be
	// considered read-only.
	Layout() pango.Layout
	// LayoutOffsets obtains the coordinates where the label will draw its
	// `PangoLayout`.
	//
	// The coordinates are useful to convert mouse events into coordinates
	// inside the [class@Pango.Layout], e.g. to take some action if some part of
	// the label is clicked. Remember when using the [class@Pango.Layout]
	// functions you need to convert to and from pixels using PANGO_PIXELS() or
	// [constant@Pango.SCALE].
	LayoutOffsets() (x int, y int)
	// Lines gets the number of lines to which an ellipsized, wrapping label
	// should be limited.
	//
	// See [method@Gtk.Label.set_lines].
	Lines() int
	// MaxWidthChars retrieves the desired maximum width of @label, in
	// characters.
	//
	// See [method@Gtk.Label.set_width_chars].
	MaxWidthChars() int
	// MnemonicKeyval: return the mnemonic accelerator.
	//
	// If the label has been set so that it has a mnemonic key this function
	// returns the keyval used for the mnemonic accelerator. If there is no
	// mnemonic set up it returns `GDK_KEY_VoidSymbol`.
	MnemonicKeyval() uint
	// MnemonicWidget retrieves the target of the mnemonic (keyboard shortcut)
	// of this label.
	//
	// See [method@Gtk.Label.set_mnemonic_widget].
	MnemonicWidget() Widget
	// Selectable returns whether the label is selectable.
	Selectable() bool
	// SelectionBounds gets the selected range of characters in the label.
	SelectionBounds() (start int, end int, ok bool)
	// SingleLineMode returns whether the label is in single line mode.
	SingleLineMode() bool
	// Text fetches the text from a label.
	//
	// The returned text is as it appears on screen. This does not include any
	// embedded underlines indicating mnemonics or Pango markup. (See
	// [method@Gtk.Label.get_label])
	Text() string
	// UseMarkup returns whether the label’s text is interpreted as Pango
	// markup.
	//
	// See [method@Gtk.Label.set_use_markup].
	UseMarkup() bool
	// UseUnderline returns whether an embedded underlines in the label indicate
	// mnemonics.
	//
	// See [method@Gtk.Label.set_use_underline].
	UseUnderline() bool
	// WidthChars retrieves the desired width of @label, in characters.
	//
	// See [method@Gtk.Label.set_width_chars].
	WidthChars() int
	// Wrap returns whether lines in the label are automatically wrapped.
	//
	// See [method@Gtk.Label.set_wrap].
	Wrap() bool
	// WrapMode returns line wrap mode used by the label.
	//
	// See [method@Gtk.Label.set_wrap_mode].
	WrapMode() pango.WrapMode
	// Xalign gets the `xalign` of the label.
	//
	// See the [property@Gtk.Label:xalign] property.
	Xalign() float32
	// Yalign gets the `yalign` of the label.
	//
	// See the [property@Gtk.Label:yalign] property.
	Yalign() float32
	// SelectRegionLabel selects a range of characters in the label, if the
	// label is selectable.
	//
	// See [method@Gtk.Label.set_selectable]. If the label is not selectable,
	// this function has no effect. If @start_offset or @end_offset are -1, then
	// the end of the label will be substituted.
	SelectRegionLabel(startOffset int, endOffset int)
	// SetAttributesLabel: apply attributes to the label text.
	//
	// The attributes set with this function will be applied and merged with any
	// other attributes previously effected by way of the
	// [property@Gtk.Label:use-underline] or [property@Gtk.Label:use-markup]
	// properties. While it is not recommended to mix markup strings with
	// manually set attributes, if you must; know that the attributes will be
	// applied to the label after the markup string is parsed.
	SetAttributesLabel(attrs pango.AttrList)
	// SetEllipsizeLabel sets the mode used to ellipsizei the text.
	//
	// The text will be ellipsized if there is not enough space to render the
	// entire string.
	SetEllipsizeLabel(mode pango.EllipsizeMode)
	// SetExtraMenuLabel sets a menu model to add when constructing the context
	// menu for @label.
	SetExtraMenuLabel(model gio.MenuModel)
	// SetJustifyLabel sets the alignment of the lines in the text of the label
	// relative to each other.
	//
	// GTK_JUSTIFY_LEFT is the default value when the widget is first created
	// with [ctor@Gtk.Label.new]. If you instead want to set the alignment of
	// the label as a whole, use [method@Gtk.Widget.set_halign] instead.
	// [method@Gtk.Label.set_justify] has no effect on labels containing only a
	// single line.
	SetJustifyLabel(jtype Justification)
	// SetLabelLabel sets the text of the label.
	//
	// The label is interpreted as including embedded underlines and/or Pango
	// markup depending on the values of the [property@Gtk.Label:use-underline]
	// and [property@Gtk.Label:use-markup] properties.
	SetLabelLabel(str string)
	// SetLinesLabel sets the number of lines to which an ellipsized, wrapping
	// label should be limited.
	//
	// This has no effect if the label is not wrapping or ellipsized. Set this
	// to -1 if you don’t want to limit the number of lines.
	SetLinesLabel(lines int)
	// SetMarkupLabel sets the labels text and attributes from markup.
	//
	// The string must be marked up with Pango markup (see
	// [func@Pango.parse_markup]).
	//
	// If the @str is external data, you may need to escape it with
	// g_markup_escape_text() or g_markup_printf_escaped():
	//
	// “`c GtkWidget *self = gtk_label_new (NULL); const char *str = "...";
	// const char *format = "<span style=\"italic\">\s</span>"; char *markup;
	//
	// markup = g_markup_printf_escaped (format, str); gtk_label_set_markup
	// (GTK_LABEL (self), markup); g_free (markup); “`
	//
	// This function will set the [property@Gtk.Label:use-markup] property to
	// true as a side effect.
	//
	// If you set the label contents using the [property@Gtk.Label:label]
	// property you should also ensure that you set the
	// [property@Gtk.Label:use-markup] property accordingly.
	//
	// See also: [method@Gtk.Label.set_text]
	SetMarkupLabel(str string)
	// SetMarkupWithMnemonicLabel sets the labels text, attributes and mnemonic
	// from markup.
	//
	// Parses @str which is marked up with Pango markup (see
	// [func@Pango.parse_markup]), setting the label’s text and attribute list
	// based on the parse results. If characters in @str are preceded by an
	// underscore, they are underlined indicating that they represent a keyboard
	// accelerator called a mnemonic.
	//
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using method@Gtk.Label.set_mnemonic_widget].
	SetMarkupWithMnemonicLabel(str string)
	// SetMaxWidthCharsLabel sets the desired maximum width in characters of
	// @label to @n_chars.
	SetMaxWidthCharsLabel(nChars int)
	// SetMnemonicWidgetLabel: associate the label with its mnemonic target.
	//
	// If the label has been set so that it has a mnemonic key (using i.e.
	// [method@Gtk.Label.set_markup_with_mnemonic],
	// [method@Gtk.Label.set_text_with_mnemonic],
	// [ctor@Gtk.Label.new_with_mnemonic] or the
	// [property@Gtk.Label:use_underline] property) the label can be associated
	// with a widget that is the target of the mnemonic. When the label is
	// inside a widget (like a [class@Gtk.Button] or a [class@Gtk.Notebook] tab)
	// it is automatically associated with the correct widget, but sometimes
	// (i.e. when the target is a [class@Gtk.Entry] next to the label) you need
	// to set it explicitly using this function.
	//
	// The target widget will be accelerated by emitting the
	// [signal@GtkWidget::mnemonic-activate] signal on it. The default handler
	// for this signal will activate the widget if there are no mnemonic
	// collisions and toggle focus between the colliding widgets otherwise.
	SetMnemonicWidgetLabel(widget Widget)
	// SetSelectableLabel makes text in the label selectable.
	//
	// Selectable labels allow the user to select text from the label, for
	// copy-and-paste.
	SetSelectableLabel(setting bool)
	// SetSingleLineModeLabel sets whether the label is in single line mode.
	SetSingleLineModeLabel(singleLineMode bool)
	// SetTextLabel sets the text within the `GtkLabel` widget.
	//
	// It overwrites any text that was there before.
	//
	// This function will clear any previously set mnemonic accelerators, and
	// set the [property@Gtk.Label:use-underline property] to false as a side
	// effect.
	//
	// This function will set the [property@Gtk.Label:use-markup] property to
	// false as a side effect.
	//
	// See also: [method@Gtk.Label.set_markup]
	SetTextLabel(str string)
	// SetTextWithMnemonicLabel sets the label’s text from the string @str.
	//
	// If characters in @str are preceded by an underscore, they are underlined
	// indicating that they represent a keyboard accelerator called a mnemonic.
	// The mnemonic key can be used to activate another widget, chosen
	// automatically, or explicitly using
	// [method@Gtk.Label.set_mnemonic_widget].
	SetTextWithMnemonicLabel(str string)
	// SetUseMarkupLabel sets whether the text of the label contains markup.
	//
	// See [method@Gtk.Label.set_markup].
	SetUseMarkupLabel(setting bool)
	// SetUseUnderlineLabel sets whether underlines in the text indicate
	// mnemonics.
	SetUseUnderlineLabel(setting bool)
	// SetWidthCharsLabel sets the desired width in characters of @label to
	// @n_chars.
	SetWidthCharsLabel(nChars int)
	// SetWrapLabel toggles line wrapping within the `GtkLabel` widget.
	//
	// true makes it break lines if text exceeds the widget’s size. false lets
	// the text get cut off by the edge of the widget if it exceeds the widget
	// size.
	//
	// Note that setting line wrapping to true does not make the label wrap at
	// its parent container’s width, because GTK widgets conceptually can’t make
	// their requisition depend on the parent container’s size. For a label that
	// wraps at a specific position, set the label’s width using
	// [method@Gtk.Widget.set_size_request].
	SetWrapLabel(wrap bool)
	// SetWrapModeLabel controls how line wrapping is done.
	//
	// This only affects the label if line wrapping is on. (See
	// [method@Gtk.Label.set_wrap]) The default is PANGO_WRAP_WORD which means
	// wrap on word boundaries.
	SetWrapModeLabel(wrapMode pango.WrapMode)
	// SetXalignLabel sets the `xalign` of the label.
	//
	// See the [property@Gtk.Label:xalign] property.
	SetXalignLabel(xalign float32)
	// SetYalignLabel sets the `yalign` of the label.
	//
	// See the [property@Gtk.Label:yalign] property.
	SetYalignLabel(yalign float32)
}

// label implements the Label class.
type label struct {
	Widget
}

// WrapLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapLabel(obj *externglib.Object) Label {
	return label{
		Widget: WrapWidget(obj),
	}
}

func marshalLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLabel(obj), nil
}

// NewLabel creates a new label with the given text inside it.
//
// You can pass nil to get an empty label widget.
func NewLabel(str string) Label {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_label_new(_arg1)

	var _label Label // out

	_label = WrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// NewLabelWithMnemonic creates a new `GtkLabel`, containing the text in @str.
//
// If characters in @str are preceded by an underscore, they are underlined. If
// you need a literal underscore character in a label, use '__' (two
// underscores). The first underlined character represents a keyboard
// accelerator called a mnemonic. The mnemonic key can be used to activate
// another widget, chosen automatically, or explicitly using
// [method@Gtk.Label.set_mnemonic_widget].
//
// If [method@Gtk.Label.set_mnemonic_widget] is not called, then the first
// activatable ancestor of the `GtkLabel` will be chosen as the mnemonic widget.
// For instance, if the label is inside a button or menu item, the button or
// menu item will automatically become the mnemonic widget and be activated by
// the mnemonic.
func NewLabelWithMnemonic(str string) Label {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_label_new_with_mnemonic(_arg1)

	var _label Label // out

	_label = WrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

func (s label) Attributes() pango.AttrList {
	var _arg0 *C.GtkLabel      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_attributes(_arg0)

	var _attrList pango.AttrList // out

	_attrList = (pango.AttrList)(unsafe.Pointer(_cret))
	C.pango_attr_list_ref(_cret)

	return _attrList
}

func (s label) CurrentURI() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_current_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s label) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkLabel          // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_ellipsize(_arg0)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

func (s label) ExtraMenu() gio.MenuModel {
	var _arg0 *C.GtkLabel   // out
	var _cret *C.GMenuModel // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_extra_menu(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.MenuModel)

	return _menuModel
}

func (s label) Justify() Justification {
	var _arg0 *C.GtkLabel        // out
	var _cret C.GtkJustification // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_justify(_arg0)

	var _justification Justification // out

	_justification = Justification(_cret)

	return _justification
}

func (s label) Label() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_label(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s label) Layout() pango.Layout {
	var _arg0 *C.GtkLabel    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_layout(_arg0)

	var _layout pango.Layout // out

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.Layout)

	return _layout
}

func (s label) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.int      // in
	var _arg2 *C.int      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	C.gtk_label_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

func (s label) Lines() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_lines(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s label) MaxWidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_max_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s label) MnemonicKeyval() uint {
	var _arg0 *C.GtkLabel // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_mnemonic_keyval(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s label) MnemonicWidget() Widget {
	var _arg0 *C.GtkLabel  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_mnemonic_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Widget)

	return _widget
}

func (s label) Selectable() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_selectable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s label) SelectionBounds() (start int, end int, ok bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.int      // in
	var _arg2 *C.int      // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

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

func (s label) SingleLineMode() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_single_line_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s label) Text() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.char     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s label) UseMarkup() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_use_markup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s label) UseUnderline() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_use_underline(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s label) WidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.int       // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (s label) Wrap() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_wrap(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s label) WrapMode() pango.WrapMode {
	var _arg0 *C.GtkLabel     // out
	var _cret C.PangoWrapMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_wrap_mode(_arg0)

	var _wrapMode pango.WrapMode // out

	_wrapMode = pango.WrapMode(_cret)

	return _wrapMode
}

func (s label) Xalign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.float     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_xalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (s label) Yalign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.float     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_label_get_yalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (s label) SelectRegionLabel(startOffset int, endOffset int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out
	var _arg2 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(startOffset)
	_arg2 = C.int(endOffset)

	C.gtk_label_select_region(_arg0, _arg1, _arg2)
}

func (s label) SetAttributesLabel(attrs pango.AttrList) {
	var _arg0 *C.GtkLabel      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs))

	C.gtk_label_set_attributes(_arg0, _arg1)
}

func (s label) SetEllipsizeLabel(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkLabel          // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_label_set_ellipsize(_arg0, _arg1)
}

func (s label) SetExtraMenuLabel(model gio.MenuModel) {
	var _arg0 *C.GtkLabel   // out
	var _arg1 *C.GMenuModel // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_label_set_extra_menu(_arg0, _arg1)
}

func (s label) SetJustifyLabel(jtype Justification) {
	var _arg0 *C.GtkLabel        // out
	var _arg1 C.GtkJustification // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkJustification(jtype)

	C.gtk_label_set_justify(_arg0, _arg1)
}

func (s label) SetLabelLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_label(_arg0, _arg1)
}

func (s label) SetLinesLabel(lines int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(lines)

	C.gtk_label_set_lines(_arg0, _arg1)
}

func (s label) SetMarkupLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup(_arg0, _arg1)
}

func (s label) SetMarkupWithMnemonicLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup_with_mnemonic(_arg0, _arg1)
}

func (s label) SetMaxWidthCharsLabel(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(nChars)

	C.gtk_label_set_max_width_chars(_arg0, _arg1)
}

func (s label) SetMnemonicWidgetLabel(widget Widget) {
	var _arg0 *C.GtkLabel  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_label_set_mnemonic_widget(_arg0, _arg1)
}

func (s label) SetSelectableLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_selectable(_arg0, _arg1)
}

func (s label) SetSingleLineModeLabel(singleLineMode bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	if singleLineMode {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_single_line_mode(_arg0, _arg1)
}

func (s label) SetTextLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text(_arg0, _arg1)
}

func (s label) SetTextWithMnemonicLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.char     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text_with_mnemonic(_arg0, _arg1)
}

func (s label) SetUseMarkupLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_markup(_arg0, _arg1)
}

func (s label) SetUseUnderlineLabel(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_underline(_arg0, _arg1)
}

func (s label) SetWidthCharsLabel(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(nChars)

	C.gtk_label_set_width_chars(_arg0, _arg1)
}

func (s label) SetWrapLabel(wrap bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_wrap(_arg0, _arg1)
}

func (s label) SetWrapModeLabel(wrapMode pango.WrapMode) {
	var _arg0 *C.GtkLabel     // out
	var _arg1 C.PangoWrapMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.PangoWrapMode(wrapMode)

	C.gtk_label_set_wrap_mode(_arg0, _arg1)
}

func (s label) SetXalignLabel(xalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(xalign)

	C.gtk_label_set_xalign(_arg0, _arg1)
}

func (s label) SetYalignLabel(yalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.float     // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(yalign)

	C.gtk_label_set_yalign(_arg0, _arg1)
}

func (l label) AsAccessible() Accessible {
	return WrapAccessible(gextras.InternObject(l))
}

func (l label) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(l))
}

func (l label) AsConstraintTarget() ConstraintTarget {
	return WrapConstraintTarget(gextras.InternObject(l))
}
