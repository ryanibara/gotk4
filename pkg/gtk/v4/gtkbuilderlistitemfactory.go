// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_builder_list_item_factory_get_type()), F: marshalBuilderListItemFactory},
	})
}

// BuilderListItemFactory: `GtkBuilderListItemFactory` is a `GtkListItemFactory`
// that creates widgets by instantiating `GtkBuilder` UI templates.
//
// The templates must be extending `GtkListItem`, and typically use
// `GtkExpression`s to obtain data from the items in the model.
//
// Example: “`xml <interface> <template class="GtkListItem"> <property
// name="child"> <object class="GtkLabel"> <property name="xalign">0</property>
// <binding name="label"> <lookup name="name" type="SettingsKey"> <lookup
// name="item">GtkListItem</lookup> </lookup> </binding> </object> </property>
// </template> </interface> “`
type BuilderListItemFactory interface {
	ListItemFactory

	Resource() string

	Scope() BuilderScope
}

// builderListItemFactory implements the BuilderListItemFactory class.
type builderListItemFactory struct {
	ListItemFactory
}

// WrapBuilderListItemFactory wraps a GObject to the right type. It is
// primarily used internally.
func WrapBuilderListItemFactory(obj *externglib.Object) BuilderListItemFactory {
	return builderListItemFactory{
		ListItemFactory: WrapListItemFactory(obj),
	}
}

func marshalBuilderListItemFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuilderListItemFactory(obj), nil
}

func NewBuilderListItemFactoryFromResource(scope BuilderScope, resourcePath string) BuilderListItemFactory {
	var _arg1 *C.GtkBuilderScope    // out
	var _arg2 *C.char               // out
	var _cret *C.GtkListItemFactory // in

	_arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))
	_arg2 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_builder_list_item_factory_new_from_resource(_arg1, _arg2)

	var _builderListItemFactory BuilderListItemFactory // out

	_builderListItemFactory = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(BuilderListItemFactory)

	return _builderListItemFactory
}

func (s builderListItemFactory) Resource() string {
	var _arg0 *C.GtkBuilderListItemFactory // out
	var _cret *C.char                      // in

	_arg0 = (*C.GtkBuilderListItemFactory)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_builder_list_item_factory_get_resource(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s builderListItemFactory) Scope() BuilderScope {
	var _arg0 *C.GtkBuilderListItemFactory // out
	var _cret *C.GtkBuilderScope           // in

	_arg0 = (*C.GtkBuilderListItemFactory)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_builder_list_item_factory_get_scope(_arg0)

	var _builderScope BuilderScope // out

	_builderScope = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(BuilderScope)

	return _builderScope
}
