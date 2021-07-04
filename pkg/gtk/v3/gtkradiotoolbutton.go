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
		{T: externglib.Type(C.gtk_radio_tool_button_get_type()), F: marshalRadioToolButton},
	})
}

// RadioToolButton: a RadioToolButton is a ToolItem that contains a radio
// button, that is, a button that is part of a group of toggle buttons where
// only one button can be active at a time.
//
// Use gtk_radio_tool_button_new() to create a new GtkRadioToolButton. Use
// gtk_radio_tool_button_new_from_widget() to create a new GtkRadioToolButton
// that is part of the same group as an existing GtkRadioToolButton.
//
//
// CSS nodes
//
// GtkRadioToolButton has a single CSS node with name toolbutton.
type RadioToolButton interface {
	ToggleToolButton
}

// radioToolButton implements the RadioToolButton class.
type radioToolButton struct {
	ToggleToolButton
}

// WrapRadioToolButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioToolButton(obj *externglib.Object) RadioToolButton {
	return radioToolButton{
		ToggleToolButton: WrapToggleToolButton(obj),
	}
}

func marshalRadioToolButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioToolButton(obj), nil
}

func NewRadioToolButtonFromWidget(group RadioToolButton) RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _cret *C.GtkToolItem        // in

	_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))

	_cret = C.gtk_radio_tool_button_new_from_widget(_arg1)

	var _radioToolButton RadioToolButton // out

	_radioToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(RadioToolButton)

	return _radioToolButton
}

func NewRadioToolButtonWithStockFromWidget(group RadioToolButton, stockId string) RadioToolButton {
	var _arg1 *C.GtkRadioToolButton // out
	var _arg2 *C.gchar              // out
	var _cret *C.GtkToolItem        // in

	_arg1 = (*C.GtkRadioToolButton)(unsafe.Pointer(group.Native()))
	_arg2 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_tool_button_new_with_stock_from_widget(_arg1, _arg2)

	var _radioToolButton RadioToolButton // out

	_radioToolButton = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(RadioToolButton)

	return _radioToolButton
}

func (b radioToolButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioToolButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioToolButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioToolButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioToolButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioToolButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b radioToolButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b radioToolButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioToolButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioToolButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a radioToolButton) DoSetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).DoSetRelatedAction(action)
}

func (a radioToolButton) RelatedAction() Action {
	return WrapActivatable(gextras.InternObject(a)).RelatedAction()
}

func (a radioToolButton) UseActionAppearance() bool {
	return WrapActivatable(gextras.InternObject(a)).UseActionAppearance()
}

func (a radioToolButton) SetRelatedAction(action Action) {
	WrapActivatable(gextras.InternObject(a)).SetRelatedAction(action)
}

func (a radioToolButton) SetUseActionAppearance(useAppearance bool) {
	WrapActivatable(gextras.InternObject(a)).SetUseActionAppearance(useAppearance)
}

func (a radioToolButton) SyncActionProperties(action Action) {
	WrapActivatable(gextras.InternObject(a)).SyncActionProperties(action)
}

func (a radioToolButton) ActionName() string {
	return WrapActionable(gextras.InternObject(a)).ActionName()
}

func (a radioToolButton) ActionTargetValue() *glib.Variant {
	return WrapActionable(gextras.InternObject(a)).ActionTargetValue()
}

func (a radioToolButton) SetActionName(actionName string) {
	WrapActionable(gextras.InternObject(a)).SetActionName(actionName)
}

func (a radioToolButton) SetActionTargetValue(targetValue *glib.Variant) {
	WrapActionable(gextras.InternObject(a)).SetActionTargetValue(targetValue)
}

func (a radioToolButton) SetDetailedActionName(detailedActionName string) {
	WrapActionable(gextras.InternObject(a)).SetDetailedActionName(detailedActionName)
}

func (b radioToolButton) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioToolButton) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioToolButton) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioToolButton) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data *interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioToolButton) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioToolButton) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b radioToolButton) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b radioToolButton) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioToolButton) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioToolButton) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}
