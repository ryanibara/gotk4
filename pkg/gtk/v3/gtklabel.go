// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_label_get_type()), F: marshalLabeller},
	})
}

// LabelOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LabelOverrider interface {
	ActivateLink(uri string) bool
	CopyClipboard()
	MoveCursor(step MovementStep, count int, extendSelection bool)
	PopulatePopup(menu *Menu)
}

// Label widget displays a small amount of text. As the name implies, most
// labels are used to label another widget such as a Button, a MenuItem, or a
// ComboBox.
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
type Label struct {
	Misc
}

func wrapLabel(obj *externglib.Object) *Label {
	return &Label{
		Misc: Misc{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalLabeller(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLabel(obj), nil
}

// NewLabel creates a new label with the given text inside it. You can pass NULL
// to get an empty label widget.
func NewLabel(str string) *Label {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if str != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_label_new(_arg1)
	runtime.KeepAlive(str)

	var _label *Label // out

	_label = wrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// NewLabelWithMnemonic creates a new Label, containing the text in str.
//
// If characters in str are preceded by an underscore, they are underlined. If
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
func NewLabelWithMnemonic(str string) *Label {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if str != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_label_new_with_mnemonic(_arg1)
	runtime.KeepAlive(str)

	var _label *Label // out

	_label = wrapLabel(externglib.Take(unsafe.Pointer(_cret)))

	return _label
}

// Angle gets the angle of rotation for the label. See gtk_label_set_angle().
func (label *Label) Angle() float64 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_angle(_arg0)
	runtime.KeepAlive(label)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Attributes gets the attribute list that was set on the label using
// gtk_label_set_attributes(), if any. This function does not reflect attributes
// that come from the labels markup (see gtk_label_set_markup()). If you want to
// get the effective attributes for the label, use pango_layout_get_attribute
// (gtk_label_get_layout (label)).
func (label *Label) Attributes() *pango.AttrList {
	var _arg0 *C.GtkLabel      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_attributes(_arg0)
	runtime.KeepAlive(label)

	var _attrList *pango.AttrList // out

	if _cret != nil {
		_attrList = (*pango.AttrList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.pango_attr_list_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_attrList)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.pango_attr_list_unref((*C.PangoAttrList)(intern.C))
			},
		)
	}

	return _attrList
}

// CurrentURI returns the URI for the currently active link in the label. The
// active link is the one under the mouse pointer or, in a selectable label, the
// link in which the text cursor is currently positioned.
//
// This function is intended for use in a Label::activate-link handler or for
// use in a Widget::query-tooltip handler.
func (label *Label) CurrentURI() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_current_uri(_arg0)
	runtime.KeepAlive(label)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Ellipsize returns the ellipsizing position of the label. See
// gtk_label_set_ellipsize().
func (label *Label) Ellipsize() pango.EllipsizeMode {
	var _arg0 *C.GtkLabel          // out
	var _cret C.PangoEllipsizeMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_ellipsize(_arg0)
	runtime.KeepAlive(label)

	var _ellipsizeMode pango.EllipsizeMode // out

	_ellipsizeMode = pango.EllipsizeMode(_cret)

	return _ellipsizeMode
}

// Justify returns the justification of the label. See gtk_label_set_justify().
func (label *Label) Justify() Justification {
	var _arg0 *C.GtkLabel        // out
	var _cret C.GtkJustification // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_justify(_arg0)
	runtime.KeepAlive(label)

	var _justification Justification // out

	_justification = Justification(_cret)

	return _justification
}

// Label fetches the text from a label widget including any embedded underlines
// indicating mnemonics and Pango markup. (See gtk_label_get_text()).
func (label *Label) Label() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_label(_arg0)
	runtime.KeepAlive(label)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Layout gets the Layout used to display the label. The layout is useful to
// e.g. convert text positions to pixel positions, in combination with
// gtk_label_get_layout_offsets(). The returned layout is owned by the label so
// need not be freed by the caller. The label is free to recreate its layout at
// any time, so it should be considered read-only.
func (label *Label) Layout() *pango.Layout {
	var _arg0 *C.GtkLabel    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_layout(_arg0)
	runtime.KeepAlive(label)

	var _layout *pango.Layout // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_layout = &pango.Layout{
			Object: obj,
		}
	}

	return _layout
}

// LayoutOffsets obtains the coordinates where the label will draw the Layout
// representing the text in the label; useful to convert mouse events into
// coordinates inside the Layout, e.g. to take some action if some part of the
// label is clicked. Of course you will need to create a EventBox to receive the
// events, and pack the label inside it, since labels are windowless (they
// return FALSE from gtk_widget_get_has_window()). Remember when using the
// Layout functions you need to convert to and from pixels using PANGO_PIXELS()
// or NGO_SCALE.
func (label *Label) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	C.gtk_label_get_layout_offsets(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(label)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

// LineWrap returns whether lines in the label are automatically wrapped. See
// gtk_label_set_line_wrap().
func (label *Label) LineWrap() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_line_wrap(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LineWrapMode returns line wrap mode used by the label. See
// gtk_label_set_line_wrap_mode().
func (label *Label) LineWrapMode() pango.WrapMode {
	var _arg0 *C.GtkLabel     // out
	var _cret C.PangoWrapMode // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_line_wrap_mode(_arg0)
	runtime.KeepAlive(label)

	var _wrapMode pango.WrapMode // out

	_wrapMode = pango.WrapMode(_cret)

	return _wrapMode
}

// Lines gets the number of lines to which an ellipsized, wrapping label should
// be limited. See gtk_label_set_lines().
func (label *Label) Lines() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_lines(_arg0)
	runtime.KeepAlive(label)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MaxWidthChars retrieves the desired maximum width of label, in characters.
// See gtk_label_set_width_chars().
func (label *Label) MaxWidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_max_width_chars(_arg0)
	runtime.KeepAlive(label)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// MnemonicKeyval: if the label has been set so that it has an mnemonic key this
// function returns the keyval used for the mnemonic accelerator. If there is no
// mnemonic set up it returns K_KEY_VoidSymbol.
func (label *Label) MnemonicKeyval() uint {
	var _arg0 *C.GtkLabel // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_mnemonic_keyval(_arg0)
	runtime.KeepAlive(label)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MnemonicWidget retrieves the target of the mnemonic (keyboard shortcut) of
// this label. See gtk_label_set_mnemonic_widget().
func (label *Label) MnemonicWidget() Widgetter {
	var _arg0 *C.GtkLabel  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_mnemonic_widget(_arg0)
	runtime.KeepAlive(label)

	var _widget Widgetter // out

	if _cret != nil {
		_widget = (externglib.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)
	}

	return _widget
}

// Selectable gets the value set by gtk_label_set_selectable().
func (label *Label) Selectable() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_selectable(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SelectionBounds gets the selected range of characters in the label, returning
// TRUE if there’s a selection.
func (label *Label) SelectionBounds() (start int, end int, ok bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_selection_bounds(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(label)

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

// SingleLineMode returns whether the label is in single line mode.
func (label *Label) SingleLineMode() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_single_line_mode(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Text fetches the text from a label widget, as displayed on the screen. This
// does not include any embedded underlines indicating mnemonics or Pango
// markup. (See gtk_label_get_label())
func (label *Label) Text() string {
	var _arg0 *C.GtkLabel // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_text(_arg0)
	runtime.KeepAlive(label)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TrackVisitedLinks returns whether the label is currently keeping track of
// clicked links.
func (label *Label) TrackVisitedLinks() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_track_visited_links(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseMarkup returns whether the label’s text is interpreted as marked up with
// the [Pango text markup language][PangoMarkupFormat]. See
// gtk_label_set_use_markup ().
func (label *Label) UseMarkup() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_use_markup(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UseUnderline returns whether an embedded underline in the label indicates a
// mnemonic. See gtk_label_set_use_underline().
func (label *Label) UseUnderline() bool {
	var _arg0 *C.GtkLabel // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_use_underline(_arg0)
	runtime.KeepAlive(label)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WidthChars retrieves the desired width of label, in characters. See
// gtk_label_set_width_chars().
func (label *Label) WidthChars() int {
	var _arg0 *C.GtkLabel // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_width_chars(_arg0)
	runtime.KeepAlive(label)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// XAlign gets the Label:xalign property for label.
func (label *Label) XAlign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_xalign(_arg0)
	runtime.KeepAlive(label)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// YAlign gets the Label:yalign property for label.
func (label *Label) YAlign() float32 {
	var _arg0 *C.GtkLabel // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))

	_cret = C.gtk_label_get_yalign(_arg0)
	runtime.KeepAlive(label)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// SelectRegion selects a range of characters in the label, if the label is
// selectable. See gtk_label_set_selectable(). If the label is not selectable,
// this function has no effect. If start_offset or end_offset are -1, then the
// end of the label will be substituted.
func (label *Label) SelectRegion(startOffset int, endOffset int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out
	var _arg2 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gint(startOffset)
	_arg2 = C.gint(endOffset)

	C.gtk_label_select_region(_arg0, _arg1, _arg2)
	runtime.KeepAlive(label)
	runtime.KeepAlive(startOffset)
	runtime.KeepAlive(endOffset)
}

// SetAngle sets the angle of rotation for the label. An angle of 90 reads from
// from bottom to top, an angle of 270, from top to bottom. The angle setting
// for the label is ignored if the label is selectable, wrapped, or ellipsized.
func (label *Label) SetAngle(angle float64) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gdouble(angle)

	C.gtk_label_set_angle(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(angle)
}

// SetAttributes sets a AttrList; the attributes in the list are applied to the
// label text.
//
// The attributes set with this function will be applied and merged with any
// other attributes previously effected by way of the Label:use-underline or
// Label:use-markup properties. While it is not recommended to mix markup
// strings with manually set attributes, if you must; know that the attributes
// will be applied to the label after the markup string is parsed.
func (label *Label) SetAttributes(attrs *pango.AttrList) {
	var _arg0 *C.GtkLabel      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if attrs != nil {
		_arg1 = (*C.PangoAttrList)(gextras.StructNative(unsafe.Pointer(attrs)))
	}

	C.gtk_label_set_attributes(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(attrs)
}

// SetEllipsize sets the mode used to ellipsize (add an ellipsis: "...") to the
// text if there is not enough space to render the entire string.
func (label *Label) SetEllipsize(mode pango.EllipsizeMode) {
	var _arg0 *C.GtkLabel          // out
	var _arg1 C.PangoEllipsizeMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.PangoEllipsizeMode(mode)

	C.gtk_label_set_ellipsize(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(mode)
}

// SetJustify sets the alignment of the lines in the text of the label relative
// to each other. GTK_JUSTIFY_LEFT is the default value when the widget is first
// created with gtk_label_new(). If you instead want to set the alignment of the
// label as a whole, use gtk_widget_set_halign() instead.
// gtk_label_set_justify() has no effect on labels containing only a single
// line.
func (label *Label) SetJustify(jtype Justification) {
	var _arg0 *C.GtkLabel        // out
	var _arg1 C.GtkJustification // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.GtkJustification(jtype)

	C.gtk_label_set_justify(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(jtype)
}

// SetLabel sets the text of the label. The label is interpreted as including
// embedded underlines and/or Pango markup depending on the values of the
// Label:use-underline and Label:use-markup properties.
func (label *Label) SetLabel(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_label(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(str)
}

// SetLineWrap toggles line wrapping within the Label widget. TRUE makes it
// break lines if text exceeds the widget’s size. FALSE lets the text get cut
// off by the edge of the widget if it exceeds the widget size.
//
// Note that setting line wrapping to TRUE does not make the label wrap at its
// parent container’s width, because GTK+ widgets conceptually can’t make their
// requisition depend on the parent container’s size. For a label that wraps at
// a specific position, set the label’s width using
// gtk_widget_set_size_request().
func (label *Label) SetLineWrap(wrap bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if wrap {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_line_wrap(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(wrap)
}

// SetLineWrapMode: if line wrapping is on (see gtk_label_set_line_wrap()) this
// controls how the line wrapping is done. The default is PANGO_WRAP_WORD which
// means wrap on word boundaries.
func (label *Label) SetLineWrapMode(wrapMode pango.WrapMode) {
	var _arg0 *C.GtkLabel     // out
	var _arg1 C.PangoWrapMode // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.PangoWrapMode(wrapMode)

	C.gtk_label_set_line_wrap_mode(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(wrapMode)
}

// SetLines sets the number of lines to which an ellipsized, wrapping label
// should be limited. This has no effect if the label is not wrapping or
// ellipsized. Set this to -1 if you don’t want to limit the number of lines.
func (label *Label) SetLines(lines int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gint(lines)

	C.gtk_label_set_lines(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(lines)
}

// SetMarkup parses str which is marked up with the [Pango text markup
// language][PangoMarkupFormat], setting the label’s text and attribute list
// based on the parse results.
//
// If the str is external data, you may need to escape it with
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
// This function will set the Label:use-markup property to TRUE as a side
// effect.
//
// If you set the label contents using the Label:label property you should also
// ensure that you set the Label:use-markup property accordingly.
//
// See also: gtk_label_set_text()
func (label *Label) SetMarkup(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(str)
}

// SetMarkupWithMnemonic parses str which is marked up with the [Pango text
// markup language][PangoMarkupFormat], setting the label’s text and attribute
// list based on the parse results. If characters in str are preceded by an
// underscore, they are underlined indicating that they represent a keyboard
// accelerator called a mnemonic.
//
// The mnemonic key can be used to activate another widget, chosen
// automatically, or explicitly using gtk_label_set_mnemonic_widget().
func (label *Label) SetMarkupWithMnemonic(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_markup_with_mnemonic(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(str)
}

// SetMaxWidthChars sets the desired maximum width in characters of label to
// n_chars.
func (label *Label) SetMaxWidthChars(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_label_set_max_width_chars(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(nChars)
}

// SetMnemonicWidget: if the label has been set so that it has an mnemonic key
// (using i.e. gtk_label_set_markup_with_mnemonic(),
// gtk_label_set_text_with_mnemonic(), gtk_label_new_with_mnemonic() or the
// “use_underline” property) the label can be associated with a widget that is
// the target of the mnemonic. When the label is inside a widget (like a Button
// or a Notebook tab) it is automatically associated with the correct widget,
// but sometimes (i.e. when the target is a Entry next to the label) you need to
// set it explicitly using this function.
//
// The target widget will be accelerated by emitting the
// GtkWidget::mnemonic-activate signal on it. The default handler for this
// signal will activate the widget if there are no mnemonic collisions and
// toggle focus between the colliding widgets otherwise.
func (label *Label) SetMnemonicWidget(widget Widgetter) {
	var _arg0 *C.GtkLabel  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if widget != nil {
		_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	}

	C.gtk_label_set_mnemonic_widget(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(widget)
}

// SetPattern: pattern of underlines you want under the existing text within the
// Label widget. For example if the current text of the label says “FooBarBaz”
// passing a pattern of “___ ___” will underline “Foo” and “Baz” but not “Bar”.
func (label *Label) SetPattern(pattern string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pattern)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_pattern(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(pattern)
}

// SetSelectable: selectable labels allow the user to select text from the
// label, for copy-and-paste.
func (label *Label) SetSelectable(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_selectable(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(setting)
}

// SetSingleLineMode sets whether the label is in single line mode.
func (label *Label) SetSingleLineMode(singleLineMode bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if singleLineMode {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_single_line_mode(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(singleLineMode)
}

// SetText sets the text within the Label widget. It overwrites any text that
// was there before.
//
// This function will clear any previously set mnemonic accelerators, and set
// the Label:use-underline property to FALSE as a side effect.
//
// This function will set the Label:use-markup property to FALSE as a side
// effect.
//
// See also: gtk_label_set_markup()
func (label *Label) SetText(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(str)
}

// SetTextWithMnemonic sets the label’s text from the string str. If characters
// in str are preceded by an underscore, they are underlined indicating that
// they represent a keyboard accelerator called a mnemonic. The mnemonic key can
// be used to activate another widget, chosen automatically, or explicitly using
// gtk_label_set_mnemonic_widget().
func (label *Label) SetTextWithMnemonic(str string) {
	var _arg0 *C.GtkLabel // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_label_set_text_with_mnemonic(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(str)
}

// SetTrackVisitedLinks sets whether the label should keep track of clicked
// links (and use a different color for them).
func (label *Label) SetTrackVisitedLinks(trackLinks bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if trackLinks {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_track_visited_links(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(trackLinks)
}

// SetUseMarkup sets whether the text of the label contains markup in [Pango’s
// text markup language][PangoMarkupFormat]. See gtk_label_set_markup().
func (label *Label) SetUseMarkup(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_markup(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(setting)
}

// SetUseUnderline: if true, an underline in the text indicates the next
// character should be used for the mnemonic accelerator key.
func (label *Label) SetUseUnderline(setting bool) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_label_set_use_underline(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(setting)
}

// SetWidthChars sets the desired width in characters of label to n_chars.
func (label *Label) SetWidthChars(nChars int) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_label_set_width_chars(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(nChars)
}

// SetXAlign sets the Label:xalign property for label.
func (label *Label) SetXAlign(xalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_label_set_xalign(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(xalign)
}

// SetYAlign sets the Label:yalign property for label.
func (label *Label) SetYAlign(yalign float32) {
	var _arg0 *C.GtkLabel // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkLabel)(unsafe.Pointer(label.Native()))
	_arg1 = C.gfloat(yalign)

	C.gtk_label_set_yalign(_arg0, _arg1)
	runtime.KeepAlive(label)
	runtime.KeepAlive(yalign)
}
