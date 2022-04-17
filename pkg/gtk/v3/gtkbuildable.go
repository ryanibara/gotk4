// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern GObject* _gotk4_gtk3_BuildableIface_construct_child(GtkBuildable*, GtkBuilder*, gchar*);
// extern GObject* _gotk4_gtk3_BuildableIface_get_internal_child(GtkBuildable*, GtkBuilder*, gchar*);
// extern gboolean _gotk4_gtk3_BuildableIface_custom_tag_start(GtkBuildable*, GtkBuilder*, GObject*, gchar*, GMarkupParser*, gpointer*);
// extern gchar* _gotk4_gtk3_BuildableIface_get_name(GtkBuildable*);
// extern void _gotk4_gtk3_BuildableIface_add_child(GtkBuildable*, GtkBuilder*, GObject*, gchar*);
// extern void _gotk4_gtk3_BuildableIface_custom_finished(GtkBuildable*, GtkBuilder*, GObject*, gchar*, gpointer);
// extern void _gotk4_gtk3_BuildableIface_custom_tag_end(GtkBuildable*, GtkBuilder*, GObject*, gchar*, gpointer*);
// extern void _gotk4_gtk3_BuildableIface_parser_finished(GtkBuildable*, GtkBuilder*);
// extern void _gotk4_gtk3_BuildableIface_set_buildable_property(GtkBuildable*, GtkBuilder*, gchar*, GValue*);
// extern void _gotk4_gtk3_BuildableIface_set_name(GtkBuildable*, gchar*);
import "C"

// glib.Type values for gtkbuildable.go.
var GTypeBuildable = externglib.Type(C.gtk_buildable_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeBuildable, F: marshalBuildable},
	})
}

// BuildableOverrider contains methods that are overridable.
type BuildableOverrider interface {
	externglib.Objector
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child to add.
	//    - typ (optional): kind of child or NULL.
	//
	AddChild(builder *Builder, child *externglib.Object, typ string)
	// ConstructChild constructs a child of buildable with the name name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - name of child to construct.
	//
	// The function returns the following values:
	//
	//    - object: constructed child.
	//
	ConstructChild(builder *Builder, name string) *externglib.Object
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of the tag.
	//    - data (optional): user data created in custom_tag_start.
	//
	CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//    - data (optional): user data that will be passed in to parser
	//      functions.
	//
	CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data *cgo.Handle)
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// The function takes the following parameters:
	//
	//    - builder used to construct this object.
	//    - child (optional) object or NULL for non-child tags.
	//    - tagname: name of tag.
	//
	// The function returns the following values:
	//
	//    - parser to fill in.
	//    - data (optional): return location for user data that will be passed in
	//      to parser functions.
	//    - ok: TRUE if a object has a custom implementation, FALSE if it
	//      doesn't.
	//
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (*glib.MarkupParser, cgo.Handle, bool)
	// InternalChild: get the internal child called childname of the buildable
	// object.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - childname: name of child.
	//
	// The function returns the following values:
	//
	//    - object: internal child of the buildable object.
	//
	InternalChild(builder *Builder, childname string) *externglib.Object
	// Name gets the name of the buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the buildable.
	//
	// The function returns the following values:
	//
	//    - utf8: name set with gtk_buildable_set_name().
	//
	Name() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//
	ParserFinished(builder *Builder)
	// SetBuildableProperty sets the property name name to value on the
	// buildable object.
	//
	// The function takes the following parameters:
	//
	//    - builder: Builder.
	//    - name of property.
	//    - value of property.
	//
	SetBuildableProperty(builder *Builder, name string, value *externglib.Value)
	// SetName sets the name of the buildable object.
	//
	// The function takes the following parameters:
	//
	//    - name to set.
	//
	SetName(name string)
}

// Buildable allows objects to extend and customize their deserialization from
// [GtkBuilder UI descriptions][BUILDER-UI]. The interface includes methods for
// setting names and properties of objects, parsing custom tags and constructing
// child objects.
//
// The GtkBuildable interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK+. The main user of this interface
// is Builder. There should be very little need for applications to call any of
// these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// Builder format or run any extra routines at deserialization time.
type Buildable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Buildable)(nil)
)

// Buildabler describes Buildable's interface methods.
type Buildabler interface {
	externglib.Objector

	// AddChild adds a child to buildable.
	AddChild(builder *Builder, child *externglib.Object, typ string)
	// ConstructChild constructs a child of buildable with the name name.
	ConstructChild(builder *Builder, name string) *externglib.Object
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the buildable.
	CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data *cgo.Handle)
	// CustomTagStart: this is called for each unknown element under <child>.
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (*glib.MarkupParser, cgo.Handle, bool)
	// InternalChild: get the internal child called childname of the buildable
	// object.
	InternalChild(builder *Builder, childname string) *externglib.Object
	// Name gets the name of the buildable object.
	Name() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI].
	ParserFinished(builder *Builder)
	// SetBuildableProperty sets the property name name to value on the
	// buildable object.
	SetBuildableProperty(builder *Builder, name string, value *externglib.Value)
	// SetName sets the name of the buildable object.
	SetName(name string)
}

var _ Buildabler = (*Buildable)(nil)

func ifaceInitBuildabler(gifacePtr, data C.gpointer) {
	iface := (*C.GtkBuildableIface)(unsafe.Pointer(gifacePtr))
	iface.add_child = (*[0]byte)(C._gotk4_gtk3_BuildableIface_add_child)
	iface.construct_child = (*[0]byte)(C._gotk4_gtk3_BuildableIface_construct_child)
	iface.custom_finished = (*[0]byte)(C._gotk4_gtk3_BuildableIface_custom_finished)
	iface.custom_tag_end = (*[0]byte)(C._gotk4_gtk3_BuildableIface_custom_tag_end)
	iface.custom_tag_start = (*[0]byte)(C._gotk4_gtk3_BuildableIface_custom_tag_start)
	iface.get_internal_child = (*[0]byte)(C._gotk4_gtk3_BuildableIface_get_internal_child)
	iface.get_name = (*[0]byte)(C._gotk4_gtk3_BuildableIface_get_name)
	iface.parser_finished = (*[0]byte)(C._gotk4_gtk3_BuildableIface_parser_finished)
	iface.set_buildable_property = (*[0]byte)(C._gotk4_gtk3_BuildableIface_set_buildable_property)
	iface.set_name = (*[0]byte)(C._gotk4_gtk3_BuildableIface_set_name)
}

//export _gotk4_gtk3_BuildableIface_add_child
func _gotk4_gtk3_BuildableIface_add_child(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _typ string               // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_child = externglib.Take(unsafe.Pointer(arg2))
	if arg3 != nil {
		_typ = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	}

	iface.AddChild(_builder, _child, _typ)
}

//export _gotk4_gtk3_BuildableIface_construct_child
func _gotk4_gtk3_BuildableIface_construct_child(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.gchar) (cret *C.GObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _name string      // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.ConstructChild(_builder, _name)

	cret = (*C.GObject)(unsafe.Pointer(object.Native()))
	C.g_object_ref(C.gpointer(object.Native()))

	return cret
}

//export _gotk4_gtk3_BuildableIface_custom_finished
func _gotk4_gtk3_BuildableIface_custom_finished(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.gchar, arg4 C.gpointer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out
	var _data cgo.Handle          // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	_data = (cgo.Handle)(unsafe.Pointer(arg4))

	iface.CustomFinished(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk3_BuildableIface_custom_tag_end
func _gotk4_gtk3_BuildableIface_custom_tag_end(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.gchar, arg4 *C.gpointer) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out
	var _data *cgo.Handle         // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))
	if arg4 != nil {
		_data = (*cgo.Handle)(unsafe.Pointer(arg4))
	}

	iface.CustomTagEnd(_builder, _child, _tagname, _data)
}

//export _gotk4_gtk3_BuildableIface_custom_tag_start
func _gotk4_gtk3_BuildableIface_custom_tag_start(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.GObject, arg3 *C.gchar, arg4 *C.GMarkupParser, arg5 *C.gpointer) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder         // out
	var _child *externglib.Object // out
	var _tagname string           // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_child = externglib.Take(unsafe.Pointer(arg2))
	}
	_tagname = C.GoString((*C.gchar)(unsafe.Pointer(arg3)))

	parser, data, ok := iface.CustomTagStart(_builder, _child, _tagname)

	*arg4 = *(*C.GMarkupParser)(gextras.StructNative(unsafe.Pointer(parser)))
	*arg5 = (C.gpointer)(unsafe.Pointer(data))
	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_BuildableIface_get_internal_child
func _gotk4_gtk3_BuildableIface_get_internal_child(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.gchar) (cret *C.GObject) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out
	var _childname string // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_childname = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	object := iface.InternalChild(_builder, _childname)

	cret = (*C.GObject)(unsafe.Pointer(object.Native()))

	return cret
}

//export _gotk4_gtk3_BuildableIface_get_name
func _gotk4_gtk3_BuildableIface_get_name(arg0 *C.GtkBuildable) (cret *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	utf8 := iface.Name()

	cret = (*C.gchar)(unsafe.Pointer(C.CString(utf8)))
	defer C.free(unsafe.Pointer(cret))

	return cret
}

//export _gotk4_gtk3_BuildableIface_parser_finished
func _gotk4_gtk3_BuildableIface_parser_finished(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))

	iface.ParserFinished(_builder)
}

//export _gotk4_gtk3_BuildableIface_set_buildable_property
func _gotk4_gtk3_BuildableIface_set_buildable_property(arg0 *C.GtkBuildable, arg1 *C.GtkBuilder, arg2 *C.gchar, arg3 *C.GValue) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _builder *Builder        // out
	var _name string             // out
	var _value *externglib.Value // out

	_builder = wrapBuilder(externglib.Take(unsafe.Pointer(arg1)))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))
	_value = externglib.ValueFromNative(unsafe.Pointer(arg3))

	iface.SetBuildableProperty(_builder, _name, _value)
}

//export _gotk4_gtk3_BuildableIface_set_name
func _gotk4_gtk3_BuildableIface_set_name(arg0 *C.GtkBuildable, arg1 *C.gchar) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(BuildableOverrider)

	var _name string // out

	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	iface.SetName(_name)
}

func wrapBuildable(obj *externglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	return wrapBuildable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AddChild adds a child to buildable. type is an optional string describing how
// the child should be added.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - child to add.
//    - typ (optional): kind of child or NULL.
//
func (buildable *Buildable) AddChild(builder *Builder, child *externglib.Object, typ string) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	if typ != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.gtk_buildable_add_child(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(typ)
}

// ConstructChild constructs a child of buildable with the name name.
//
// Builder calls this function if a “constructor” has been specified in the UI
// definition.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - name of child to construct.
//
// The function returns the following values:
//
//    - object: constructed child.
//
func (buildable *Buildable) ConstructChild(builder *Builder, name string) *externglib.Object {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_buildable_construct_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)

	var _object *externglib.Object // out

	_object = externglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
// called once for each custom tag handled by the buildable.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of the tag.
//    - data (optional): user data created in custom_tag_start.
//
func (buildable *Buildable) CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 C.gpointer      // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.gpointer)(unsafe.Pointer(data))

	C.gtk_buildable_custom_finished(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)
	runtime.KeepAlive(data)
}

// CustomTagEnd: this is called at the end of each custom element handled by the
// buildable.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of tag.
//    - data (optional): user data that will be passed in to parser functions.
//
func (buildable *Buildable) CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data *cgo.Handle) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.gpointer     // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))
	if data != nil {
		_arg4 = (*C.gpointer)(unsafe.Pointer(data))
	}

	C.gtk_buildable_custom_tag_end(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)
	runtime.KeepAlive(data)
}

// CustomTagStart: this is called for each unknown element under <child>.
//
// The function takes the following parameters:
//
//    - builder used to construct this object.
//    - child (optional) object or NULL for non-child tags.
//    - tagname: name of tag.
//
// The function returns the following values:
//
//    - parser to fill in.
//    - data (optional): return location for user data that will be passed in to
//      parser functions.
//    - ok: TRUE if a object has a custom implementation, FALSE if it doesn't.
//
func (buildable *Buildable) CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (*glib.MarkupParser, cgo.Handle, bool) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 C.GMarkupParser // in
	var _arg5 C.gpointer      // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_buildable_custom_tag_start(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(child)
	runtime.KeepAlive(tagname)

	var _parser *glib.MarkupParser // out
	var _data cgo.Handle           // out
	var _ok bool                   // out

	_parser = (*glib.MarkupParser)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))
	_data = (cgo.Handle)(unsafe.Pointer(_arg5))
	if _cret != 0 {
		_ok = true
	}

	return _parser, _data, _ok
}

// InternalChild: get the internal child called childname of the buildable
// object.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - childname: name of child.
//
// The function returns the following values:
//
//    - object: internal child of the buildable object.
//
func (buildable *Buildable) InternalChild(builder *Builder, childname string) *externglib.Object {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(childname)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_buildable_get_internal_child(_arg0, _arg1, _arg2)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(childname)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Name gets the name of the buildable object.
//
// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
// used to construct the buildable.
//
// The function returns the following values:
//
//    - utf8: name set with gtk_buildable_set_name().
//
func (buildable *Buildable) Name() string {
	var _arg0 *C.GtkBuildable // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))

	_cret = C.gtk_buildable_get_name(_arg0)
	runtime.KeepAlive(buildable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ParserFinished: called when the builder finishes the parsing of a [GtkBuilder
// UI definition][BUILDER-UI]. Note that this will be called once for each time
// gtk_builder_add_from_file() or gtk_builder_add_from_string() is called on a
// builder.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//
func (buildable *Buildable) ParserFinished(builder *Builder) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))

	C.gtk_buildable_parser_finished(_arg0, _arg1)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
}

// SetBuildableProperty sets the property name name to value on the buildable
// object.
//
// The function takes the following parameters:
//
//    - builder: Builder.
//    - name of property.
//    - value of property.
//
func (buildable *Buildable) SetBuildableProperty(builder *Builder, name string, value *externglib.Value) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(externglib.InternObject(builder).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_buildable_set_buildable_property(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(builder)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// SetName sets the name of the buildable object.
//
// The function takes the following parameters:
//
//    - name to set.
//
func (buildable *Buildable) SetName(name string) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(externglib.InternObject(buildable).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_buildable_set_name(_arg0, _arg1)
	runtime.KeepAlive(buildable)
	runtime.KeepAlive(name)
}
