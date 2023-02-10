package glib

// #include <glib.h>
// #include <glib-object.h>
// #include "glib.go.h"
//
// extern void _gotk4_gobject_init_class(gpointer, gpointer);
// extern void _gotk4_gobject_init_instance(GTypeInstance*, gpointer);
// extern void _gotk4_gobject_dispose(GObject*);
// extern void _gotk4_gobject_get_property(GObject*, guint, GValue*, GParamSpec*);
// extern void _gotk4_gobject_set_property(GObject*, guint, GValue*, GParamSpec*);
import "C"

import (
	"log"
	"math"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"unicode"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

func init() {
	RegisterClassInfo[
		*Object,
		*struct{},
		ObjectOverrides,
	](
		TypeObject,
		initObjectClass,
		func(o *Object) *Object { return o },
		func(o *Object) ObjectOverrides {
			// This is actually quite funny. Instead of having our own thing
			// that the caller can choose to call here, we just don't let the
			// caller choose anything. Instead, we will always do our thing in
			// the C callback.
			return ObjectOverrides{
				Init:    func() {},
				Dispose: func() {},
			}
		},
	)
}

// ObjectOverrides contains optional virtual methods hookable onto the Object
// type.
type ObjectOverrides struct {
	// Init is called during init_instance, which is actually when a
	// GTypeInstance is instantiated, not GObject.
	Init func()
	// Dispose is called during GObject's dispose.
	Dispose func()
}

func initObjectClass(gclass unsafe.Pointer, overrides ObjectOverrides, initFunc func(*struct{})) {
	// Install our GObject methods.
	gobjectClass := (*C.GObjectClass)(gclass)
	gobjectClass.dispose = (*[0]byte)(C._gotk4_gobject_dispose)
	gobjectClass.get_property = (*[0]byte)(C._gotk4_gobject_get_property)
	gobjectClass.set_property = (*[0]byte)(C._gotk4_gobject_set_property)
}

// RegisteredSubclass is a type that described a registered Go subclass type.
type RegisteredSubclass[T any] registeredSubclass

func rtypeElem[T any]() reflect.Type {
	var zero T
	rtype := reflect.TypeOf(zero)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	return rtype
}

func rtypeElemV(v any) reflect.Type {
	// Fast path?
	rtype := reflect.TypeOf(v)
	if rtype.Kind() == reflect.Ptr {
		rtype = rtype.Elem()
	}
	return rtype
}

type registerOpts struct {
	typeOpts   map[*classTypeInfo]registerTypeOpts // nil == currentType
	paramSpecs []*ParamSpec
	gtypeName  string
}

// registerTypeOpts contains register options specific to a single subclass
// type.
type registerTypeOpts struct {
	override  func(Objector) any
	classInit func(any)
}

func (o *registerOpts) wantTypeOpts() {
	if o.typeOpts == nil {
		o.typeOpts = make(map[*classTypeInfo]registerTypeOpts)
	}
}

// RegisterOptsFunc is a function type that modifies the behavior of a
// RegisterSubclass call.
type RegisterOptsFunc func(*registerOpts)

// WithParamSpecs adds additional ParamSpecs into a type for its properties.
func WithParamSpecs(paramSpecs []*ParamSpec) RegisterOptsFunc {
	return func(opts *registerOpts) { opts.paramSpecs = append(opts.paramSpecs, paramSpecs...) }
}

// WithGTypeName overrides the autogenerated GLib type name, which is the
// namespace concatenated with the Go type name.
func WithGTypeName(gtypeName string) RegisterOptsFunc {
	return func(opts *registerOpts) { opts.gtypeName = gtypeName }
}

// WithClassInit adds a custom ClassInit function. Use this method to add a
// function that, for example, takes in a WidgetClass.
func WithClassInit[ClassT any](f func(ClassT)) RegisterOptsFunc {
	return func(opts *registerOpts) {
		opts.wantTypeOpts()
		rtype := rtypeElem[ClassT]()

		ptype, ok := classTypeClassTypes[rtype]
		if !ok {
			log.Panicf("gotk4: cannot add ClassInit func for unknown Class type %s", rtype)
		}

		topt := opts.typeOpts[ptype]
		topt.classInit = func(v any) {
			class, _ := v.(ClassT)
			f(class)
		}
		opts.typeOpts[ptype] = topt
	}
}

// WithOverrides adds one XOverrides instance into the subclass. The user can
// use this to override any particular virtual method of any of its parent
// classes.
func WithOverrides[T Objector, OverridesT any](f func(T) OverridesT) RegisterOptsFunc {
	return func(opts *registerOpts) {
		opts.wantTypeOpts()
		rtype := rtypeElem[OverridesT]()

		ptype, ok := classOverridesTypes[rtype]
		if !ok {
			log.Panicf("gotk4: cannot add Overrides func for unknown Overrides type %s", rtype)
			// ptype = nil
		}

		topt := opts.typeOpts[ptype]
		topt.override = func(obj Objector) any {
			v, _ := obj.(T)
			return f(v)
		}
		opts.typeOpts[ptype] = topt
	}
}

// classInitParamType gets the Class type of the classInit function.
func classInitParamType(funcValue any) reflect.Type {
	// We know this is a func(*T).
	rtype := reflect.TypeOf(funcValue)
	firstParamType := rtype.In(0)
	return firstParamType.Elem()
}

// RegisterSubclass is RegisterSubclassWithConstructor, but a zero-value
// instance of T is automatically created.
func RegisterSubclass[T Objector](opts ...RegisterOptsFunc) *RegisteredSubclass[T] {
	rtype := rtypeElem[T]()
	return RegisterSubclassWithConstructor(func() T {
		return reflect.New(rtype).Interface().(T)
	}, opts...)
}

// RegisterSubclassWithConstructor registers a new type T that is a subclass of
// its parent type, which is the first field that must be embedded.
//
// ctor has to be idempotent (i.e. can be called multiple times w/o side
// effects).
func RegisterSubclassWithConstructor[T Objector](ctor func() T, opts ...RegisterOptsFunc) *RegisteredSubclass[T] {
	rtype := rtypeElem[T]()

	subclass, ok := knownTypes[rtype]
	if ok {
		log.Panicln(
			"gotk4: type", rtype,
			"is already registered as a subclass of", subclass.parentType.GoInstanceType)
	}

	subclass = registerSubclass(rtype, func() Objector { return ctor() }, opts)
	knownTypes[rtype] = subclass
	knownGTypes[subclass.gType] = subclass

	return castRegisteredSubclass[T](subclass)
}

func castRegisteredSubclass[T any](src *registeredSubclass) *RegisteredSubclass[T] {
	return (*RegisteredSubclass[T])(unsafe.Pointer(src))
}

// New creates an instance of the subclass object with no properties.
func (r *RegisteredSubclass[T]) New() T {
	return r.NewWithProperties(nil)
}

// NewWithProperties creates an instance of the subclass object with the given
// properties.
func (r *RegisteredSubclass[T]) NewWithProperties(properties map[string]any) T {
	var names_ **C.gchar
	var values_ *C.GValue

	if len(properties) > 0 {
		names := make([]*C.char, 0, len(properties))
		values := make([]C.GValue, 0, len(properties))

		for name, value := range properties {
			cname := (*C.char)(C.CString(name))
			defer C.free(unsafe.Pointer(cname))

			gvalue := NewValue(value)
			defer runtime.KeepAlive(gvalue)

			names = append(names, cname)
			values = append(values, *gvalue.gvalue)
		}

		names_ = &names[0]
		values_ = &values[0]
	}

	cval := C.g_object_new_with_properties(
		C.GType(r.gType),
		C.guint(len(properties)),
		names_, values_,
	)

	// So we created an instance. We have the instance in private. We can just
	// dig it up and avoid the whole type-casting mess.
	private := privateFromInstance(unsafe.Pointer(cval))
	return private.instance().(T)
}

// Type returns the GType of the registered Go subclass.
func (r *RegisteredSubclass[T]) Type() Type {
	return r.gType
}

type registeredSubclass struct {
	goType      reflect.Type // NOT POINTER
	gType       Type
	parentType  *classTypeInfo
	typeClass   unsafe.Pointer
	constructor func() Objector
	properties  map[*C.GParamSpec]subclassProperty

	opts registerOpts
}

type subclassProperty struct {
	name     string
	fieldIdx []int
}

func (p *subclassProperty) instance(goval any) propertyInstance {
	prop := reflect.Indirect(reflect.ValueOf(goval)).FieldByIndex(p.fieldIdx)
	return prop.Addr().Interface().(propertyInstance)
}

var (
	knownTypes  = map[reflect.Type]*registeredSubclass{}
	knownGTypes = map[Type]*registeredSubclass{}
)

func subclassFromData(data C.gpointer) *registeredSubclass {
	return gbox.Get(uintptr(data)).(*registeredSubclass)
}

func registerSubclass(rtype reflect.Type, ctor func() Objector, optsFuncs []RegisterOptsFunc) *registeredSubclass {
	subclass := &registeredSubclass{
		goType:      rtype,
		parentType:  extractParentType(rtype),
		constructor: ctor,
	}

	for _, fn := range optsFuncs {
		fn(&subclass.opts)
	}

	// We want len(opts.paramSpecs) == len(properties). Ideally, ParamSpecs
	// should be optional and we should automatically detect the type, however
	// that is not the case.
	for _, spec := range subclass.opts.paramSpecs {
		// Permanently take a reference, since we're globalling this forever
		// anyway.
		C.g_param_spec_ref(spec.intern)
	}

	// Scan for Property fields.
	subclass.properties = make(map[*C.GParamSpec]subclassProperty, len(subclass.opts.paramSpecs))
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		if !strings.HasSuffix(field.Type.PkgPath(), "/core/glib") {
			continue
		}
		// Hack!
		if !strings.HasPrefix(field.Type.Name(), "Property[") {
			continue
		}

		name := field.Tag.Get("glib")
		if name == "" {
			log.Panicf("field %s of type %s has no glib tag", field.Name, rtype)
		}

		if !reflect.PtrTo(field.Type).Implements(rtypeProperty) {
			log.Panicf("BUG: *Property[T] of %s does not implement propertyInstance", reflect.PtrTo(field.Type))
		}

		// Validate the name.
		var gspec *C.GParamSpec
		for _, spec := range subclass.opts.paramSpecs {
			if spec.Name() == name {
				gspec = spec.intern
				break
			}
		}

		if gspec == nil {
			// Try and generate our own.
			nilProp := reflect.NewAt(field.Type, nil).Interface().(propertyInstance)
			ptyp := rtypeToParamType(nilProp.rtype())

			gspec = gParamSpecBlank(name, ptyp)
			paramSpecTypeInit(unsafe.Pointer(gspec), ptyp)

			// Append. This means WE CANNOT RELY ON THE INDEX!
			subclass.opts.paramSpecs = append(subclass.opts.paramSpecs, &ParamSpec{
				&paramSpec{gspec},
			})
		}

		subclass.properties[gspec] = subclassProperty{
			name:     name,
			fieldIdx: field.Index,
		}
	}

	var typeQuery C.GTypeQuery
	C.g_type_query(C.GType(subclass.parentType.GType), &typeQuery)
	if typeQuery._type == 0 {
		log.Panicln("GType", subclass.parentType.GType, "is generated but is unknown")
	}

	var typeInfo C.GTypeInfo

	// Why are these ushort anyway?
	typeInfo.class_size = C.gushort(typeQuery.class_size)
	typeInfo.class_data = C.gconstpointer(gbox.Assign(subclass))
	typeInfo.class_init = C.GClassInitFunc(C._gotk4_gobject_init_class)

	typeInfo.instance_size = C.gushort(typeQuery.instance_size)
	typeInfo.instance_init = C.GInstanceInitFunc(C._gotk4_gobject_init_instance)

	gtypeName := subclass.opts.gtypeName
	if gtypeName == "" {
		gtypeName = transformGTypeName(rtype)
	}

	ctypeName := (*C.gchar)(C.CString(gtypeName))
	defer C.free(unsafe.Pointer(ctypeName))

	gtype := C.g_type_register_static(
		C.GType(subclass.parentType.GType),
		ctypeName,
		&typeInfo,
		C.GTypeFlags(0))
	subclass.gType = Type(gtype)

	return subclass
}

// TODO: autogenerate this:
//
// type enumerator[T any] interface {
// 	Enumerate() []T
// }

// G_TYPE_PARAM_CHAR
// G_TYPE_PARAM_UCHAR
// G_TYPE_PARAM_BOOLEAN
// G_TYPE_PARAM_INT
// G_TYPE_PARAM_UINT
// G_TYPE_PARAM_LONG
// G_TYPE_PARAM_ULONG
// G_TYPE_PARAM_INT64
// G_TYPE_PARAM_UINT64
// G_TYPE_PARAM_UNICHAR
// G_TYPE_PARAM_ENUM
// G_TYPE_PARAM_FLAGS
// G_TYPE_PARAM_FLOAT
// G_TYPE_PARAM_DOUBLE
// G_TYPE_PARAM_STRING
// G_TYPE_PARAM_PARAM
// G_TYPE_PARAM_BOXED
// G_TYPE_PARAM_POINTER
// G_TYPE_PARAM_OBJECT
// G_TYPE_PARAM_OVERRIDE
// G_TYPE_PARAM_GTYPE
// G_TYPE_PARAM_VARIANT
func rtypeToParamType(rtype reflect.Type) Type {
	if rtype == rtypePVariant {
		return Type(C.G_TYPE_PARAM_VARIANT)
	}
	if rtype.Implements(rtypeObjector) {
		return Type(C.G_TYPE_PARAM_OBJECT)
	}
	switch rtype.Kind() {
	case reflect.String:
		return Type(C.G_TYPE_PARAM_STRING)
	case reflect.Int8:
		return Type(C.G_TYPE_PARAM_CHAR)
	case reflect.Int16, reflect.Int32:
		return Type(C.G_TYPE_PARAM_INT)
	case reflect.Int64, reflect.Int:
		return Type(C.G_TYPE_PARAM_INT64)
	case reflect.Uint8:
		return Type(C.G_TYPE_PARAM_UCHAR)
	case reflect.Uint16, reflect.Uint32:
		return Type(C.G_TYPE_PARAM_UINT)
	case reflect.Uint64, reflect.Uint:
		return Type(C.G_TYPE_PARAM_UINT64)
	case reflect.Float32:
		return Type(C.G_TYPE_PARAM_FLOAT)
	case reflect.Float64:
		return Type(C.G_TYPE_PARAM_DOUBLE)
	case reflect.Bool:
		return Type(C.G_TYPE_PARAM_BOOLEAN)
	case reflect.Ptr:
		// TODO: gbox
	}
	log.Panicf("gotk4: unsupported type %T cannot be converted to GTypeParam", rtype)
	return 0
}

// TODO: remove all this shit and just use the Go functions. Please.
func paramSpecTypeInit(spec unsafe.Pointer, ptype Type) {
	switch C.GType(ptype) {
	case C.G_TYPE_PARAM_INT:
		ispec := (*C.GParamSpecInt)(spec)
		ispec.minimum = C.gint(-math.MaxInt32)
		ispec.maximum = C.gint(+math.MaxInt32)
	case C.G_TYPE_PARAM_INT64:
		ispec := (*C.GParamSpecInt64)(spec)
		ispec.minimum = C.gint64(-math.MaxInt64)
		ispec.maximum = C.gint64(+math.MaxInt64)
	case C.G_TYPE_PARAM_UINT:
		uspec := (*C.GParamSpecUInt)(spec)
		uspec.minimum = C.guint(0)
		uspec.maximum = C.guint(+math.MaxUint32)
	case C.G_TYPE_PARAM_UINT64:
		uspec := (*C.GParamSpecUInt64)(spec)
		uspec.minimum = C.guint64(0)
		uspec.maximum = C.guint64(+math.MaxUint64)
	case C.G_TYPE_PARAM_FLOAT:
		fspec := (*C.GParamSpecFloat)(spec)
		fspec.minimum = C.gfloat(-math.MaxFloat32)
		fspec.maximum = C.gfloat(+math.MaxFloat32)
	case C.G_TYPE_PARAM_DOUBLE:
		fspec := (*C.GParamSpecDouble)(spec)
		fspec.minimum = C.gdouble(-math.MaxFloat64)
		fspec.maximum = C.gdouble(+math.MaxFloat64)
	}
}

// gParamSpecBlank is used for generating a ParamSpec from just the struct
// field.
func gParamSpecBlank(name string, gtype Type) *C.GParamSpec {
	cname := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(cname))

	emptyStr := (*C.gchar)(C.CString(""))
	defer C.free(unsafe.Pointer(emptyStr))

	paramSpec := (*C.GParamSpec)(C.g_param_spec_internal(
		C.GType(gtype),
		cname, emptyStr, emptyStr,
		C.GParamFlags(ParamReadWrite),
	))
	return C.g_param_spec_ref_sink(paramSpec)
}

func transformGTypeName(rtype reflect.Type) string {
	goType := rtype.String()

	first, second, ok := strings.Cut(goType, ".")
	if ok {
		return capitalizeFirst(first) + capitalizeFirst(second)
	}

	return capitalizeFirst(goType)
}

func capitalizeFirst(str string) string {
	if str == "" {
		return str
	}
	runes := []rune(str)
	return string(unicode.ToUpper(runes[0])) + string(runes[1:])
}

func extractParentType(rtype reflect.Type) *classTypeInfo {
	if rtype.Kind() != reflect.Struct {
		log.Panicln("given type is not a struct or a *struct")
	}

	field := rtype.Field(0)
	if !field.Anonymous {
		log.Panicln("first field (parent) must be embedded")
	}

	typeInfo, ok := classInstanceTypes[field.Type]
	if !ok {
		// TODO: allow inheriting from a Go type.
		log.Panicln("unknown type", field.Type)
	}

	return typeInfo
}

func (r *registeredSubclass) setParent(instance any, parent unsafe.Pointer) {
	// transfer: none! See GInstanceInitFunc's docs.
	gobject := Take(unsafe.Pointer(parent))

	// We want to set the first field of our new Go class instance, which is the
	// parent, to be the initialized parent. There's a pretty nifty way of doing
	// this: we can hijack our base methods and use that.

	// gobjectVal might have type *gtk.Widget, for example.
	gobjectVal := r.parentType.WrapClass(gobject)

	rval := reflect.ValueOf(instance).Elem()

	parentVal := rval.Field(0)
	if !parentVal.IsZero() {
		log.Panicf("cannot construct subclass %s instance with non-nil parent", r.goType)
	}

	// v.Widget = *(gobjectVal.(*gtk.Widget))
	parentVal.Set(reflect.ValueOf(gobjectVal).Elem())
}

// Overrides gets the XOverrides instance associated with the given object's
// parent. The instance will have already been modified by the subclass if it
// did.
func Overrides[OverridesT any](obj Objector) OverridesT {
	rtype := rtypeElemV(obj)

	if subclass, ok := knownTypes[rtype]; ok {
		topt, ok := subclass.opts.typeOpts[subclass.parentType]
		if ok && topt.override != nil {
			overrides, ok := topt.override(obj).(OverridesT)
			if ok {
				return overrides
			}

			var z OverridesT
			log.Panicf(
				"gotk4: Overrides: object type %s's parent overrides type is %s, cannot assert to %T",
				rtype, subclass.parentType.GoOverridesType, z)
		}

		var z OverridesT
		return z
	}

	if typeInfo, ok := classInstanceTypes[rtype]; ok {
		overrides, ok := typeInfo.Overrides(obj).(OverridesT)
		if ok {
			return overrides
		}

		var z OverridesT
		log.Panicf(
			"gotk4: Overrides: object type %s's overrides type is %s, cannot assert to %T",
			rtype, typeInfo.GoOverridesType, z)
	}

	obase := BaseObject(obj)
	return OverridesFromObj[OverridesT](obase)
}

// OverridesFromObj is like Overrides, except it specifically takes the base
// object. The GType is asked for instead of using the underlying Go type.
func OverridesFromObj[OverridesT any](obj *Object) OverridesT {
	gtype := obj.TypeFromInstance()

	if subclass, ok := knownGTypes[gtype]; ok {
		topt, ok := subclass.opts.typeOpts[subclass.parentType]
		if ok && topt.override != nil {
			overrides, ok := topt.override(obj).(OverridesT)
			if ok {
				return overrides
			}

			var z OverridesT
			log.Panicf(
				"gotk4: OverridesFromObj: object type %s's parent overrides type is %s, cannot assert to %T",
				gtype, subclass.parentType.GoOverridesType, z)
		}

		var z OverridesT
		return z
	}

	if typeInfo, ok := classGTypes[gtype]; ok {
		overrides, ok := typeInfo.Overrides(obj).(OverridesT)
		if ok {
			return overrides
		}

		var z OverridesT
		log.Panicf(
			"gotk4: OverridesFromObj: object type %s's overrides type is %s, cannot assert to %T",
			gtype, typeInfo.GoOverridesType, z)
	}

	var z OverridesT
	return z
}

// ParentOverrides is similar to Overrides, except the object's parent type is
// used instead of the object's type.
func ParentOverrides[OverridesT any](obj Objector) OverridesT {
	rtype := rtypeElemV(obj)

	subclass, ok := knownTypes[rtype]
	if !ok {
		log.Panicf(
			"gotk4: cannot get ParentOverrides of non-subclass type %s (unsupported)",
			rtype)
	}

	typeInfo, ok := classInstanceTypes[subclass.parentType.GoInstanceType]
	if ok {
		overrides, ok := typeInfo.Overrides(obj).(OverridesT)
		if ok {
			return overrides
		}
		var z OverridesT
		log.Panicf(
			"gotk4: object type %s's overrides type is %s, cannot assert to %T",
			rtype, typeInfo.GoOverridesType, z)
	}

	var z OverridesT
	return z
}

// PeekParentClass returns the C.TClass type using peek_parent, which is free of
// any subclassing changes.
func PeekParentClass(obj Objector) unsafe.Pointer {
	base := BaseObject(obj)
	parentGType := base.TypeFromInstance().Parent()
	return unsafe.Pointer(C.g_type_class_peek(C.GType(parentGType)))
}

type classTypeInfo struct {
	GType           Type
	GoInstanceType  reflect.Type
	GoClassType     reflect.Type
	GoOverridesType reflect.Type
	// InitClass will be called to initialize a class using the given Go value.
	// It should type assert goValue and sets the available functions into the
	// gclass (*GObjectTypeClass) type parameter.
	//
	// goValue might not point to a valid value.
	//
	// initFunc will be of type func(*TClass).
	InitClass func(gclass unsafe.Pointer, overrides, initFunc any)
	// WrapClass wraps the given Object to the right Ojector of GoType.
	WrapClass func(obj *Object) Objector
	// Overrides creates an overrides value from the given object.
	Overrides func(obj Objector) any
}

var (
	classInstanceTypes  = make(map[reflect.Type]*classTypeInfo, 1024)
	classTypeClassTypes = make(map[reflect.Type]*classTypeInfo, 1024)
	classOverridesTypes = make(map[reflect.Type]*classTypeInfo, 1024)
	classGTypes         = make(map[Type]*classTypeInfo, 1024)
)

// RegisterClassInfo registers the given class type info. This function is NOT
// thread-safe; only call it in init().
func RegisterClassInfo[InstanceT Objector, ClassT, OverridesT any](
	gtype Type,
	initClassFunc func(gclass unsafe.Pointer, overrides OverridesT, initFunc func(ClassT)),
	wrapClassFunc func(*Object) InstanceT,
	overridesFunc func(InstanceT) OverridesT,
) {
	typeInfo := &classTypeInfo{
		GType:           gtype,
		GoInstanceType:  rtypeElem[InstanceT](),
		GoClassType:     rtypeElem[ClassT](),
		GoOverridesType: rtypeElem[OverridesT](),
		InitClass: func(gclass unsafe.Pointer, overridesV, initFuncV any) {
			overrides, _ := overridesV.(OverridesT)
			initFunc, _ := initFuncV.(func(ClassT))
			initClassFunc(gclass, overrides, initFunc)
		},
		WrapClass: func(obj *Object) Objector {
			return wrapClassFunc(obj)
		},
		Overrides: func(obj Objector) any {
			return overridesFunc(obj.(InstanceT))
		},
	}

	classGTypes[typeInfo.GType] = typeInfo
	classInstanceTypes[typeInfo.GoInstanceType] = typeInfo
	classTypeClassTypes[typeInfo.GoClassType] = typeInfo
	classOverridesTypes[typeInfo.GoOverridesType] = typeInfo
}

// registeredGClass binds gclass (which is a *GTypeClassInfo describing XClass
// structs) to data (which is the ID that gets us a *registeredSubclass).
// It exists for init_instance.
//
// gpointer (*GTypeClassInfo) -> *privateGoInstance
var registeredPrivateInstances sync.Map

// privateGoInstance maps the two private fields.
type privateGoInstance struct {
	subclassID C.gpointer // constant
	instanceID C.gpointer // used for finalizing
}

func privateFromInstance(obj unsafe.Pointer) *privateGoInstance {
	gtype := typeFromObject(obj)

	private := C.g_type_instance_get_private((*C.GTypeInstance)(obj), C.GType(gtype))
	if private == nil {
		log.Panicf("cannot get private from unknown object %s (%p)", Type(gtype), obj)
	}

	p := (*privateGoInstance)(unsafe.Pointer(private))
	return p
}

func (p *privateGoInstance) subclass() *registeredSubclass {
	return subclassFromData(p.subclassID)
}

func (p *privateGoInstance) instance() any {
	return gbox.Get(uintptr(p.instanceID))
}

//export _gotk4_gobject_init_class
func _gotk4_gobject_init_class(gclass, data C.gpointer) {
	subclass := subclassFromData(data)
	subclass.typeClass = unsafe.Pointer(gclass)

	_, dup := registeredPrivateInstances.LoadOrStore(gclass, &privateGoInstance{subclassID: data})
	if dup {
		log.Panicf("init_class called on the same gclass %s (%p) twice", subclass.goType, gclass)
	}

	// I believe this function is only called once. The internal macros actually
	// just set a global variable right inside this callback, so I'd assume
	// that's the case. It's weird, but still.
	//
	// Also, privateOffset is originally in int, which is 32-bit. Why? Why do
	// they do this? Why not guintptr? What???
	//
	// You know what? Nevermind. Why do they deprecate the function that I'm
	// supposed to use, and then write in the comment that the generated macro
	// that's supposed to solve everything in an ideal world (???) just
	// literally calls that deprecated function? Nevermind that, why is the
	// alternative A FUCKING PRIVATE FUNCTION?! COME ON WHAT THE FUCK?!
	//
	// Anyway, here's the deprecated one. We have Cgo flags that silence
	// deprecation notices because GLib is a fucking clown.
	C.g_type_class_add_private(gclass, C.gsize(unsafe.Sizeof(privateGoInstance{})))

	for parentType, topt := range subclass.opts.typeOpts {
		if parentType == nil {
			parentType = subclass.parentType
		}

		var overrides any
		if topt.override != nil {
			overrides = topt.override(nil)
		}

		parentType.InitClass(unsafe.Pointer(gclass), overrides, topt.classInit)
	}

	// Install properties, if any.
	if len(subclass.opts.paramSpecs) > 0 {
		gParamSpecs := make([]*C.GParamSpec, len(subclass.opts.paramSpecs)+1)
		for i, spec := range subclass.opts.paramSpecs {
			// [0] is reserved by GLib.
			gParamSpecs[i+1] = spec.intern
		}

		C.g_object_class_install_properties(
			(*C.GObjectClass)(gclass),
			C.guint(len(gParamSpecs)),
			&gParamSpecs[0],
		)
	}
}

//export _gotk4_gobject_init_instance
func _gotk4_gobject_init_instance(obj *C.GTypeInstance, gclass C.gpointer) {
	// Reminder: obj of type *GTypeInstance IS a regular *GObject if we're
	// initializing a class! We can consider it as such.

	// Grab our registeredSubclass ID. We don't actually need the gclass past
	// this point.
	privateV, ok := registeredPrivateInstances.Load(gclass)
	if !ok {
		log.Panicf(
			"init_instance called on unregistered gclass %s (%p)",
			typeFromObject(unsafe.Pointer(obj)), gclass)
	}

	private := privateV.(*privateGoInstance)
	subclass := private.subclass()

	// Allocate and construct a new instance.
	instance := subclass.constructor()
	private.instanceID = C.gpointer(gbox.Assign(instance))

	// Bind our new Go class' parent field.
	subclass.setParent(instance, unsafe.Pointer(obj))

	// Copy our fully initialized private instance values to GLib's allocated
	// object private one.
	*privateFromInstance(unsafe.Pointer(obj)) = *private

	// Initialize its properties.
	if len(subclass.properties) > 0 {
		for gspec, propProto := range subclass.properties {
			prop := propProto.instance(instance)
			prop.init(instance, gspec)
		}
	}

	// Call its initializer.
	if overrides := Overrides[ObjectOverrides](instance); overrides.Init != nil {
		overrides.Init()
	}
}

//export _gotk4_gobject_dispose
func _gotk4_gobject_dispose(obj *C.GObject) {
	private := privateFromInstance(unsafe.Pointer(obj))
	instance := private.instance().(Objector)

	// Call its disposal function.
	if overrides := Overrides[ObjectOverrides](instance); overrides.Dispose != nil {
		overrides.Dispose()
	}

	// Unbind our instance from the global store.
	// TODO: should we do this in Finalize or Dispose?
	gbox.Delete(uintptr(private.instanceID))
}

func getAndValidateProperty(obj *C.GObject, spec *C.GParamSpec) propertyInstance {
	private := privateFromInstance(unsafe.Pointer(obj))

	instance := private.instance()
	subclass := private.subclass()

	propertyType, ok := subclass.properties[spec]
	if !ok {
		log.Panicf(
			"gotk4: requested property %s for object %s is unknown",
			C.GoString(spec.name), typeFromObject(unsafe.Pointer(obj)))
	}

	return propertyType.instance(instance)
}

//export _gotk4_gobject_get_property
func _gotk4_gobject_get_property(obj *C.GObject, propID C.guint, value *C.GValue, spec *C.GParamSpec) {
	if propID == 0 {
		return // reserved prop? dunno
	}

	property := getAndValidateProperty(obj, spec)
	property.get(ValueFromNative(unsafe.Pointer(value)))
}

//export _gotk4_gobject_set_property
func _gotk4_gobject_set_property(obj *C.GObject, propID C.guint, value *C.GValue, spec *C.GParamSpec) {
	if propID == 0 {
		return // reserved prop? dunno
	}

	property := getAndValidateProperty(obj, spec)
	property.set(ValueFromNative(unsafe.Pointer(value)))
}

// Property describes a Go GObject property. It is used as an alternative to
// manually written property getter/setters.
type Property[T any] struct {
	parent Objector
	name   string
	value  T

	gspec *C.GParamSpec
}

type propertyInstance interface {
	rtype() reflect.Type
	init(obj Objector, spec *C.GParamSpec)
	set(*Value)
	get(*Value)
}

var _ propertyInstance = (*Property[any])(nil)

var (
	rtypeProperty = reflect.TypeOf((*propertyInstance)(nil)).Elem()
	rtypeObjector = reflect.TypeOf((*Objector)(nil)).Elem()
	rtypePVariant = reflect.TypeOf((*Variant)(nil))
)

func (p *Property[T]) rtype() reflect.Type {
	var z T
	return reflect.TypeOf(z)
}

func (p *Property[T]) init(obj Objector, spec *C.GParamSpec) {
	p.parent = obj
	p.gspec = spec
	p.name = C.GoString(spec.name)
}

func (p *Property[T]) gtype() Type {
	return Type(p.gspec.value_type)
}

func (p *Property[T]) set(value *Value) {
	govalue := value.GoValueAsType(p.gtype())
	// This will panic if the type is not convertible.
	p.value = reflect.ValueOf(govalue).Convert(p.rtype()).Interface().(T)
}

func (p *Property[T]) get(value *Value) {
	value.InitGoValue(p.value)
}

// Set sets the new property.
func (p *Property[T]) Set(v T) {
	base := BaseObject(p.parent)
	base.SetObjectProperty(p.name, v)
}

// Get gets the value of the property.
func (p *Property[T]) Get() T {
	base := BaseObject(p.parent)
	pval := base.ObjectProperty(p.name)
	if rtype := reflect.TypeOf(pval); rtype != p.rtype() {
		return reflect.ValueOf(pval).Convert(p.rtype()).Interface().(T)
	}
	return pval.(T)
}

// Notify calls f everytime the property changes.
func (p *Property[T]) Notify(f func(T)) SignalHandle {
	base := BaseObject(p.parent)
	return base.NotifyProperty(p.name, func() { f(p.Get()) })
}
