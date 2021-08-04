// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime/cgo"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_buildable_get_type()), F: marshalBuildabler},
	})
}

// BuildableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type BuildableOverrider interface {
	// AddChild adds a child to buildable. type is an optional string describing
	// how the child should be added.
	AddChild(builder *Builder, child *externglib.Object, typ string)
	// ConstructChild constructs a child of buildable with the name name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	ConstructChild(builder *Builder, name string) *externglib.Object
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the buildable.
	CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle)
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data *cgo.Handle)
	// CustomTagStart: this is called for each unknown element under <child>.
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (glib.MarkupParser, cgo.Handle, bool)
	// InternalChild: get the internal child called childname of the buildable
	// object.
	InternalChild(builder *Builder, childname string) *externglib.Object
	// Name gets the name of the buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the buildable.
	Name() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	ParserFinished(builder *Builder)
	// SetBuildableProperty sets the property name name to value on the
	// buildable object.
	SetBuildableProperty(builder *Builder, name string, value *externglib.Value)
	// SetName sets the name of the buildable object.
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
	*externglib.Object
}

// Buildabler describes Buildable's abstract methods.
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
	CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (glib.MarkupParser, cgo.Handle, bool)
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

func wrapBuildable(obj *externglib.Object) *Buildable {
	return &Buildable{
		Object: obj,
	}
}

func marshalBuildabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBuildable(obj), nil
}

// AddChild adds a child to buildable. type is an optional string describing how
// the child should be added.
func (buildable *Buildable) AddChild(builder *Builder, child *externglib.Object, typ string) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	if typ != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(typ)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	C.gtk_buildable_add_child(_arg0, _arg1, _arg2, _arg3)
}

// ConstructChild constructs a child of buildable with the name name.
//
// Builder calls this function if a “constructor” has been specified in the UI
// definition.
func (buildable *Buildable) ConstructChild(builder *Builder, name string) *externglib.Object {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_buildable_construct_child(_arg0, _arg1, _arg2)

	var _object *externglib.Object // out

	_object = externglib.AssumeOwnership(unsafe.Pointer(_cret))

	return _object
}

// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
// called once for each custom tag handled by the buildable.
func (buildable *Buildable) CustomFinished(builder *Builder, child *externglib.Object, tagname string, data cgo.Handle) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 C.gpointer      // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.gpointer)(unsafe.Pointer(data))

	C.gtk_buildable_custom_finished(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// CustomTagEnd: this is called at the end of each custom element handled by the
// buildable.
func (buildable *Buildable) CustomTagEnd(builder *Builder, child *externglib.Object, tagname string, data *cgo.Handle) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.gpointer     // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))
	if data != nil {
		_arg4 = (*C.gpointer)(unsafe.Pointer(data))
	}

	C.gtk_buildable_custom_tag_end(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// CustomTagStart: this is called for each unknown element under <child>.
func (buildable *Buildable) CustomTagStart(builder *Builder, child *externglib.Object, tagname string) (glib.MarkupParser, cgo.Handle, bool) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.GObject      // out
	var _arg3 *C.gchar        // out
	var _arg4 C.GMarkupParser // in
	var _arg5 C.gpointer      // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	if child != nil {
		_arg2 = (*C.GObject)(unsafe.Pointer(child.Native()))
	}
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(tagname)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gtk_buildable_custom_tag_start(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _parser glib.MarkupParser // out
	var _data cgo.Handle          // out
	var _ok bool                  // out

	_parser = *(*glib.MarkupParser)(gextras.NewStructNative(unsafe.Pointer((&_arg4))))
	_data = (cgo.Handle)(unsafe.Pointer(_arg5))
	if _cret != 0 {
		_ok = true
	}

	return _parser, _data, _ok
}

// InternalChild: get the internal child called childname of the buildable
// object.
func (buildable *Buildable) InternalChild(builder *Builder, childname string) *externglib.Object {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _cret *C.GObject      // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(childname)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_buildable_get_internal_child(_arg0, _arg1, _arg2)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Name gets the name of the buildable object.
//
// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
// used to construct the buildable.
func (buildable *Buildable) Name() string {
	var _arg0 *C.GtkBuildable // out
	var _cret *C.gchar        // in

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))

	_cret = C.gtk_buildable_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ParserFinished: called when the builder finishes the parsing of a [GtkBuilder
// UI definition][BUILDER-UI]. Note that this will be called once for each time
// gtk_builder_add_from_file() or gtk_builder_add_from_string() is called on a
// builder.
func (buildable *Buildable) ParserFinished(builder *Builder) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	C.gtk_buildable_parser_finished(_arg0, _arg1)
}

// SetBuildableProperty sets the property name name to value on the buildable
// object.
func (buildable *Buildable) SetBuildableProperty(builder *Builder, name string, value *externglib.Value) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.GtkBuilder   // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.GValue       // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gtk_buildable_set_buildable_property(_arg0, _arg1, _arg2, _arg3)
}

// SetName sets the name of the buildable object.
func (buildable *Buildable) SetName(name string) {
	var _arg0 *C.GtkBuildable // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(buildable.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_buildable_set_name(_arg0, _arg1)
}
