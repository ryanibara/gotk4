// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeIconFactory returns the GType for the type IconFactory.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeIconFactory() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "IconFactory").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalIconFactory)
	return gtype
}

// IconFactoryOverrider contains methods that are overridable.
type IconFactoryOverrider interface {
}

// IconFactory: icon factory manages a collection of IconSet; a IconSet manages
// a set of variants of a particular icon (i.e. a IconSet contains variants for
// different sizes and widget states). Icons in an icon factory are named by a
// stock ID, which is a simple string identifying the icon. Each Style has a
// list of IconFactory derived from the current theme; those icon factories are
// consulted first when searching for an icon. If the theme doesn’t set a
// particular icon, GTK+ looks for the icon in a list of default icon factories,
// maintained by gtk_icon_factory_add_default() and
// gtk_icon_factory_remove_default(). Applications with icons should add a
// default icon factory with their icons, which will allow themes to override
// the icons for the application.
//
// To display an icon, always use gtk_style_lookup_icon_set() on the widget that
// will display the icon, or the convenience function gtk_widget_render_icon().
// These functions take the theme into account when looking up the icon to use
// for a given stock ID.
//
//
// GtkIconFactory as GtkBuildable
//
// GtkIconFactory supports a custom <sources> element, which can contain
// multiple <source> elements. The following attributes are allowed:
//
// - stock-id
//
//    The stock id of the source, a string. This attribute is
//    mandatory
//
// - filename
//
//    The filename of the source, a string.  This attribute is
//    optional
//
// - icon-name
//
//    The icon name for the source, a string.  This attribute is
//    optional.
//
// - size
//
//    Size of the icon, a IconSize enum value.  This attribute is
//    optional.
//
// - direction
//
//    Direction of the source, a TextDirection enum value.  This
//    attribute is optional.
//
// - state
//
//    State of the source, a StateType enum value.  This
//    attribute is optional.
//
// A IconFactory UI definition fragment. ##
//
//    <object class="GtkIconFactory" id="iconfactory1">
//      <sources>
//        <source stock-id="apple-red" filename="apple-red.png"/>
//      </sources>
//    </object>
//    <object class="GtkWindow" id="window1">
//      <child>
//        <object class="GtkButton" id="apple_button">
//          <property name="label">apple-red</property>
//          <property name="use-stock">True</property>
//        </object>
//      </child>
//    </object>.
type IconFactory struct {
	_ [0]func() // equal guard
	*coreglib.Object

	Buildable
}

var (
	_ coreglib.Objector = (*IconFactory)(nil)
)

func classInitIconFactorier(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapIconFactory(obj *coreglib.Object) *IconFactory {
	return &IconFactory{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalIconFactory(p uintptr) (interface{}, error) {
	return wrapIconFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewIconFactory creates a new IconFactory. An icon factory manages a
// collection of IconSets; a IconSet manages a set of variants of a particular
// icon (i.e. a IconSet contains variants for different sizes and widget
// states). Icons in an icon factory are named by a stock ID, which is a simple
// string identifying the icon. Each Style has a list of IconFactorys derived
// from the current theme; those icon factories are consulted first when
// searching for an icon. If the theme doesn’t set a particular icon, GTK+ looks
// for the icon in a list of default icon factories, maintained by
// gtk_icon_factory_add_default() and gtk_icon_factory_remove_default().
// Applications with icons should add a default icon factory with their icons,
// which will allow themes to override the icons for the application.
//
// Deprecated: Use IconTheme instead.
//
// The function returns the following values:
//
//    - iconFactory: new IconFactory.
//
func NewIconFactory() *IconFactory {
	_info := girepository.MustFind("Gtk", "IconFactory")
	_gret := _info.InvokeClassMethod("new_IconFactory", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _iconFactory *IconFactory // out

	_iconFactory = wrapIconFactory(coreglib.AssumeOwnership(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))

	return _iconFactory
}

// Add adds the given icon_set to the icon factory, under the name stock_id.
// stock_id should be namespaced for your application, e.g.
// “myapp-whatever-icon”. Normally applications create a IconFactory, then add
// it to the list of default factories with gtk_icon_factory_add_default(). Then
// they pass the stock_id to widgets such as Image to display the icon. Themes
// can provide an icon with the same name (such as "myapp-whatever-icon") to
// override your application’s default icons. If an icon already existed in
// factory for stock_id, it is unreferenced and replaced with the new icon_set.
//
// Deprecated: Use IconTheme instead.
//
// The function takes the following parameters:
//
//    - stockId: icon name.
//    - iconSet: icon set.
//
func (factory *IconFactory) Add(stockId string, iconSet *IconSet) {
	var _args [3]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))
	*(**C.void)(unsafe.Pointer(&_args[2])) = (*C.void)(gextras.StructNative(unsafe.Pointer(iconSet)))

	_info := girepository.MustFind("Gtk", "IconFactory")
	_info.InvokeClassMethod("add", _args[:], nil)

	runtime.KeepAlive(factory)
	runtime.KeepAlive(stockId)
	runtime.KeepAlive(iconSet)
}

// AddDefault adds an icon factory to the list of icon factories searched by
// gtk_style_lookup_icon_set(). This means that, for example,
// gtk_image_new_from_stock() will be able to find icons in factory. There will
// normally be an icon factory added for each library or application that comes
// with icons. The default icon factories can be overridden by themes.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) AddDefault() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_info := girepository.MustFind("Gtk", "IconFactory")
	_info.InvokeClassMethod("add_default", _args[:], nil)

	runtime.KeepAlive(factory)
}

// Lookup looks up stock_id in the icon factory, returning an icon set if found,
// otherwise NULL. For display to the user, you should use
// gtk_style_lookup_icon_set() on the Style for the widget that will display the
// icon, instead of using this function directly, so that themes are taken into
// account.
//
// Deprecated: Use IconTheme instead.
//
// The function takes the following parameters:
//
//    - stockId: icon name.
//
// The function returns the following values:
//
//    - iconSet: icon set of stock_id.
//
func (factory *IconFactory) Lookup(stockId string) *IconSet {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gtk", "IconFactory")
	_gret := _info.InvokeClassMethod("lookup", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(factory)
	runtime.KeepAlive(stockId)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	C.gtk_icon_set_ref(*(**C.void)(unsafe.Pointer(&_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_iconSet)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _iconSet
}

// RemoveDefault removes an icon factory from the list of default icon
// factories. Not normally used; you might use it for a library that can be
// unloaded or shut down.
//
// Deprecated: Use IconTheme instead.
func (factory *IconFactory) RemoveDefault() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_info := girepository.MustFind("Gtk", "IconFactory")
	_info.InvokeClassMethod("remove_default", _args[:], nil)

	runtime.KeepAlive(factory)
}

// IconFactoryLookupDefault looks for an icon in the list of default icon
// factories. For display to the user, you should use
// gtk_style_lookup_icon_set() on the Style for the widget that will display the
// icon, instead of using this function directly, so that themes are taken into
// account.
//
// Deprecated: Use IconTheme instead.
//
// The function takes the following parameters:
//
//    - stockId: icon name.
//
// The function returns the following values:
//
//    - iconSet or NULL.
//
func IconFactoryLookupDefault(stockId string) *IconSet {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gtk", "lookup_default")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(stockId)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	C.gtk_icon_set_ref(*(**C.void)(unsafe.Pointer(&_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_iconSet)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _iconSet
}
