// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/box"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_toggle_button_get_type()), F: marshalToggleButton},
	})
}

// ToggleButton: a ToggleButton is a Button which will remain “pressed-in” when
// clicked. Clicking again will cause the toggle button to return to its normal
// state.
//
// A toggle button is created by calling either gtk_toggle_button_new() or
// gtk_toggle_button_new_with_label(). If using the former, it is advisable to
// pack a widget, (such as a Label and/or a Image), into the toggle button’s
// container. (See Button for more information).
//
// The state of a ToggleButton can be set specifically using
// gtk_toggle_button_set_active(), and retrieved using
// gtk_toggle_button_get_active().
//
// To simply switch the state of a toggle button, use
// gtk_toggle_button_toggled().
//
//
// CSS nodes
//
// GtkToggleButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .toggle style class.
//
// Creating two ToggleButton widgets.
//
//    static void output_state (GtkToggleButton *source, gpointer user_data) {
//      printf ("Active: d\n", gtk_toggle_button_get_active (source));
//    }
//
//    void make_toggles (void) {
//      GtkWidget *window, *toggle1, *toggle2;
//      GtkWidget *box;
//      const char *text;
//
//      window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//      box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 12);
//
//      text = "Hi, I’m a toggle button.";
//      toggle1 = gtk_toggle_button_new_with_label (text);
//
//      // Makes this toggle button invisible
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle1),
//                                  TRUE);
//
//      g_signal_connect (toggle1, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle1);
//
//      text = "Hi, I’m a toggle button.";
//      toggle2 = gtk_toggle_button_new_with_label (text);
//      gtk_toggle_button_set_mode (GTK_TOGGLE_BUTTON (toggle2),
//                                  FALSE);
//      g_signal_connect (toggle2, "toggled",
//                        G_CALLBACK (output_state),
//                        NULL);
//      gtk_container_add (GTK_CONTAINER (box), toggle2);
//
//      gtk_container_add (GTK_CONTAINER (window), box);
//      gtk_widget_show_all (window);
//    }
type ToggleButton interface {
	Button

	Active() bool

	Inconsistent() bool

	Mode() bool

	SetActiveToggleButton(isActive bool)

	SetInconsistentToggleButton(setting bool)

	SetModeToggleButton(drawIndicator bool)

	ToggledToggleButton()
}

// toggleButton implements the ToggleButton class.
type toggleButton struct {
	Button
}

// WrapToggleButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapToggleButton(obj *externglib.Object) ToggleButton {
	return toggleButton{
		Button: WrapButton(obj),
	}
}

func marshalToggleButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToggleButton(obj), nil
}

func NewToggleButton() ToggleButton {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_toggle_button_new()

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func NewToggleButtonWithLabel(label string) ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_label(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func NewToggleButtonWithMnemonic(label string) ToggleButton {
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_toggle_button_new_with_mnemonic(_arg1)

	var _toggleButton ToggleButton // out

	_toggleButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ToggleButton)

	return _toggleButton
}

func (t toggleButton) Active() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_toggle_button_get_active(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toggleButton) Inconsistent() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_toggle_button_get_inconsistent(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toggleButton) Mode() bool {
	var _arg0 *C.GtkToggleButton // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	_cret = C.gtk_toggle_button_get_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (t toggleButton) SetActiveToggleButton(isActive bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if isActive {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_active(_arg0, _arg1)
}

func (t toggleButton) SetInconsistentToggleButton(setting bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_inconsistent(_arg0, _arg1)
}

func (t toggleButton) SetModeToggleButton(drawIndicator bool) {
	var _arg0 *C.GtkToggleButton // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))
	if drawIndicator {
		_arg1 = C.TRUE
	}

	C.gtk_toggle_button_set_mode(_arg0, _arg1)
}

func (t toggleButton) ToggledToggleButton() {
	var _arg0 *C.GtkToggleButton // out

	_arg0 = (*C.GtkToggleButton)(unsafe.Pointer(t.Native()))

	C.gtk_toggle_button_toggled(_arg0)
}

func (b toggleButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b toggleButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b toggleButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b toggleButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b toggleButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b toggleButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b toggleButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b toggleButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b toggleButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b toggleButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a toggleButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a toggleButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a toggleButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a toggleButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a toggleButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b toggleButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b toggleButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b toggleButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b toggleButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b toggleButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b toggleButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b toggleButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b toggleButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b toggleButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b toggleButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a toggleButton) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a toggleButton) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a toggleButton) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a toggleButton) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a toggleButton) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a toggleButton) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}
