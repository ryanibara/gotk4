// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
import "C"

// glib.Type values for gtkbuilderlistitemfactory.go.
var GTypeBuilderListItemFactory = coreglib.Type(C.gtk_builder_list_item_factory_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeBuilderListItemFactory, F: marshalBuilderListItemFactory},
	})
}

// BuilderListItemFactory: GtkBuilderListItemFactory is a GtkListItemFactory
// that creates widgets by instantiating GtkBuilder UI templates.
//
// The templates must be extending GtkListItem, and typically use GtkExpressions
// to obtain data from the items in the model.
//
// Example:
//
//    <interface>
//      <template class="GtkListItem">
//        <property name="child">
//          <object class="GtkLabel">
//            <property name="xalign">0</property>
//            <binding name="label">
//              <lookup name="name" type="SettingsKey">
//                <lookup name="item">GtkListItem</lookup>
//              </lookup>
//            </binding>
//          </object>
//        </property>
//      </template>
//    </interface>.
type BuilderListItemFactory struct {
	_ [0]func() // equal guard
	ListItemFactory
}

var (
	_ coreglib.Objector = (*BuilderListItemFactory)(nil)
)

func wrapBuilderListItemFactory(obj *coreglib.Object) *BuilderListItemFactory {
	return &BuilderListItemFactory{
		ListItemFactory: ListItemFactory{
			Object: obj,
		},
	}
}

func marshalBuilderListItemFactory(p uintptr) (interface{}, error) {
	return wrapBuilderListItemFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewBuilderListItemFactoryFromBytes creates a new GtkBuilderListItemFactory
// that instantiates widgets using bytes as the data to pass to GtkBuilder.
//
// The function takes the following parameters:
//
//    - scope (optional) to use when instantiating.
//    - bytes: GBytes containing the ui file to instantiate.
//
// The function returns the following values:
//
//    - builderListItemFactory: new GtkBuilderListItemFactory.
//
func NewBuilderListItemFactoryFromBytes(scope BuilderScoper, bytes *glib.Bytes) *BuilderListItemFactory {
	var _args [2]girepository.Argument

	if scope != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scope).Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(gextras.StructNative(unsafe.Pointer(bytes)))

	_gret := girepository.MustFind("Gtk", "BuilderListItemFactory").InvokeMethod("new_BuilderListItemFactory_from_bytes", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scope)
	runtime.KeepAlive(bytes)

	var _builderListItemFactory *BuilderListItemFactory // out

	_builderListItemFactory = wrapBuilderListItemFactory(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderListItemFactory
}

// NewBuilderListItemFactoryFromResource creates a new GtkBuilderListItemFactory
// that instantiates widgets using data read from the given resource_path to
// pass to GtkBuilder.
//
// The function takes the following parameters:
//
//    - scope (optional) to use when instantiating.
//    - resourcePath: valid path to a resource that contains the data.
//
// The function returns the following values:
//
//    - builderListItemFactory: new GtkBuilderListItemFactory.
//
func NewBuilderListItemFactoryFromResource(scope BuilderScoper, resourcePath string) *BuilderListItemFactory {
	var _args [2]girepository.Argument

	if scope != nil {
		*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(scope).Native()))
	}
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_args[1]))

	_gret := girepository.MustFind("Gtk", "BuilderListItemFactory").InvokeMethod("new_BuilderListItemFactory_from_resource", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(scope)
	runtime.KeepAlive(resourcePath)

	var _builderListItemFactory *BuilderListItemFactory // out

	_builderListItemFactory = wrapBuilderListItemFactory(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builderListItemFactory
}

// Bytes gets the data used as the GtkBuilder UI template for constructing
// listitems.
//
// The function returns the following values:
//
//    - bytes: GtkBuilder data.
//
func (self *BuilderListItemFactory) Bytes() *glib.Bytes {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "BuilderListItemFactory").InvokeMethod("get_bytes", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	C.g_bytes_ref(_cret)
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _bytes
}

// Resource: if the data references a resource, gets the path of that resource.
//
// The function returns the following values:
//
//    - utf8 (optional): path to the resource or NULL if none.
//
func (self *BuilderListItemFactory) Resource() string {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "BuilderListItemFactory").InvokeMethod("get_resource", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _utf8 string // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Scope gets the scope used when constructing listitems.
//
// The function returns the following values:
//
//    - builderScope (optional): scope used when constructing listitems.
//
func (self *BuilderListItemFactory) Scope() *BuilderScope {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(self).Native()))

	_gret := girepository.MustFind("Gtk", "BuilderListItemFactory").InvokeMethod("get_scope", _args[:], nil)
	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(self)

	var _builderScope *BuilderScope // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_builderScope = wrapBuilderScope(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _builderScope
}
