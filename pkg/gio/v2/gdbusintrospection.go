// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeDBusAnnotationInfo returns the GType for the type DBusAnnotationInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusAnnotationInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusAnnotationInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusAnnotationInfo)
	return gtype
}

// GTypeDBusArgInfo returns the GType for the type DBusArgInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusArgInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusArgInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusArgInfo)
	return gtype
}

// GTypeDBusInterfaceInfo returns the GType for the type DBusInterfaceInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusInterfaceInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusInterfaceInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusInterfaceInfo)
	return gtype
}

// GTypeDBusMethodInfo returns the GType for the type DBusMethodInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusMethodInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusMethodInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusMethodInfo)
	return gtype
}

// GTypeDBusNodeInfo returns the GType for the type DBusNodeInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusNodeInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusNodeInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusNodeInfo)
	return gtype
}

// GTypeDBusPropertyInfo returns the GType for the type DBusPropertyInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusPropertyInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusPropertyInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusPropertyInfo)
	return gtype
}

// GTypeDBusSignalInfo returns the GType for the type DBusSignalInfo.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeDBusSignalInfo() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gio", "DBusSignalInfo").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalDBusSignalInfo)
	return gtype
}

// DBusAnnotationInfo: information about an annotation.
//
// An instance of this type is always passed by reference.
type DBusAnnotationInfo struct {
	*dBusAnnotationInfo
}

// dBusAnnotationInfo is the struct that's finalized.
type dBusAnnotationInfo struct {
	native unsafe.Pointer
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusAnnotationInfo{&dBusAnnotationInfo{(unsafe.Pointer)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusAnnotationInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusAnnotationInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Key: name of the annotation, e.g. "org.freedesktop.DBus.Deprecated".
func (d *DBusAnnotationInfo) Key() string {
	offset := girepository.MustFind("Gio", "DBusAnnotationInfo").StructFieldOffset("key")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Value: value of the annotation.
func (d *DBusAnnotationInfo) Value() string {
	offset := girepository.MustFind("Gio", "DBusAnnotationInfo").StructFieldOffset("value")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusAnnotationInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusAnnotationInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusAnnotationInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusAnnotationInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}

// DBusAnnotationInfoLookup looks up the value of an annotation.
//
// The cost of this function is O(n) in number of annotations.
//
// The function takes the following parameters:
//
//    - annotations (optional): NULL-terminated array of annotations or NULL.
//    - name of the annotation to look up.
//
// The function returns the following values:
//
//    - utf8 (optional): value or NULL if not found. Do not free, it is owned by
//      annotations.
//
func DBusAnnotationInfoLookup(annotations []*DBusAnnotationInfo, name string) string {
	var _args [2]girepository.Argument

	if annotations != nil {
		{
			*(***C.void)(unsafe.Pointer(&_args[0])) = (**C.void)(C.calloc(C.size_t((len(annotations) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
			defer C.free(unsafe.Pointer(*(***C.void)(unsafe.Pointer(&_args[0]))))
			{
				out := unsafe.Slice(_args[0], len(annotations)+1)
				var zero *C.void
				out[len(annotations)] = zero
				for i := range annotations {
					*(**C.void)(unsafe.Pointer(&out[i])) = (*C.void)(gextras.StructNative(unsafe.Pointer(annotations[i])))
				}
			}
		}
	}
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "lookup")
	_gret := _info.InvokeFunction(_args[:], nil)
	_cret := *(**C.gchar)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(annotations)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if *(**C.gchar)(unsafe.Pointer(&_cret)) != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_cret)))))
	}

	return _utf8
}

// DBusArgInfo: information about an argument for a method or a signal.
//
// An instance of this type is always passed by reference.
type DBusArgInfo struct {
	*dBusArgInfo
}

// dBusArgInfo is the struct that's finalized.
type dBusArgInfo struct {
	native unsafe.Pointer
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusArgInfo{&dBusArgInfo{(unsafe.Pointer)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusArgInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusArgInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Name of the argument, e.g. unix_user_id.
func (d *DBusArgInfo) Name() string {
	offset := girepository.MustFind("Gio", "DBusArgInfo").StructFieldOffset("name")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Signature d-Bus signature of the argument (a single complete type).
func (d *DBusArgInfo) Signature() string {
	offset := girepository.MustFind("Gio", "DBusArgInfo").StructFieldOffset("signature")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusArgInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusArgInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusArgInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusArgInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}

// DBusInterfaceInfo: information about a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusInterfaceInfo struct {
	*dBusInterfaceInfo
}

// dBusInterfaceInfo is the struct that's finalized.
type dBusInterfaceInfo struct {
	native unsafe.Pointer
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusInterfaceInfo{&dBusInterfaceInfo{(unsafe.Pointer)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusInterfaceInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Name: name of the D-Bus interface, e.g. "org.freedesktop.DBus.Properties".
func (d *DBusInterfaceInfo) Name() string {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("name")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Methods: pointer to a NULL-terminated array of pointers to BusMethodInfo
// structures or NULL if there are no methods.
func (d *DBusInterfaceInfo) Methods() []*DBusMethodInfo {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("methods")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusMethodInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusMethodInfo, i)
		for i := range src {
			v[i] = (*DBusMethodInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_method_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Signals: pointer to a NULL-terminated array of pointers to BusSignalInfo
// structures or NULL if there are no signals.
func (d *DBusInterfaceInfo) Signals() []*DBusSignalInfo {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("signals")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusSignalInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusSignalInfo, i)
		for i := range src {
			v[i] = (*DBusSignalInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_signal_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Properties: pointer to a NULL-terminated array of pointers to BusPropertyInfo
// structures or NULL if there are no properties.
func (d *DBusInterfaceInfo) Properties() []*DBusPropertyInfo {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("properties")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusPropertyInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusPropertyInfo, i)
		for i := range src {
			v[i] = (*DBusPropertyInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_property_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusInterfaceInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusInterfaceInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusInterfaceInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with info, the existing cache is used and its
// use count is increased.
//
// Note that info cannot be modified until g_dbus_interface_info_cache_release()
// is called.
func (info *DBusInterfaceInfo) CacheBuild() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))

	_info := girepository.MustFind("Gio", "DBusInterfaceInfo")
	_info.InvokeRecordMethod("cache_build", _args[:], nil)

	runtime.KeepAlive(info)
}

// CacheRelease decrements the usage count for the cache for info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (info *DBusInterfaceInfo) CacheRelease() {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))

	_info := girepository.MustFind("Gio", "DBusInterfaceInfo")
	_info.InvokeRecordMethod("cache_release", _args[:], nil)

	runtime.KeepAlive(info)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on info.
//
// The function takes the following parameters:
//
//    - name d-Bus method name (typically in CamelCase).
//
// The function returns the following values:
//
//    - dBusMethodInfo (optional) or NULL if not found. Do not free, it is owned
//      by info.
//
func (info *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "DBusInterfaceInfo")
	_gret := _info.InvokeRecordMethod("lookup_method", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusMethodInfo *DBusMethodInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusMethodInfo = (*DBusMethodInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
		C.g_dbus_method_info_ref(*(**C.void)(unsafe.Pointer(&_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusMethodInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on info.
//
// The function takes the following parameters:
//
//    - name d-Bus property name (typically in CamelCase).
//
// The function returns the following values:
//
//    - dBusPropertyInfo (optional) or NULL if not found. Do not free, it is
//      owned by info.
//
func (info *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "DBusInterfaceInfo")
	_gret := _info.InvokeRecordMethod("lookup_property", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusPropertyInfo = (*DBusPropertyInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
		C.g_dbus_property_info_ref(*(**C.void)(unsafe.Pointer(&_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusPropertyInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on info.
//
// The function takes the following parameters:
//
//    - name d-Bus signal name (typically in CamelCase).
//
// The function returns the following values:
//
//    - dBusSignalInfo (optional) or NULL if not found. Do not free, it is owned
//      by info.
//
func (info *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "DBusInterfaceInfo")
	_gret := _info.InvokeRecordMethod("lookup_signal", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusSignalInfo *DBusSignalInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusSignalInfo = (*DBusSignalInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
		C.g_dbus_signal_info_ref(*(**C.void)(unsafe.Pointer(&_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusSignalInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _dBusSignalInfo
}

// DBusMethodInfo: information about a method on an D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusMethodInfo struct {
	*dBusMethodInfo
}

// dBusMethodInfo is the struct that's finalized.
type dBusMethodInfo struct {
	native unsafe.Pointer
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusMethodInfo{&dBusMethodInfo{(unsafe.Pointer)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusMethodInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Name: name of the D-Bus method, e.g. RequestName.
func (d *DBusMethodInfo) Name() string {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("name")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// InArgs: pointer to a NULL-terminated array of pointers to BusArgInfo
// structures or NULL if there are no in arguments.
func (d *DBusMethodInfo) InArgs() []*DBusArgInfo {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("in_args")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_arg_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// OutArgs: pointer to a NULL-terminated array of pointers to BusArgInfo
// structures or NULL if there are no out arguments.
func (d *DBusMethodInfo) OutArgs() []*DBusArgInfo {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("out_args")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_arg_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusMethodInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusMethodInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusMethodInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
//
// An instance of this type is always passed by reference.
type DBusNodeInfo struct {
	*dBusNodeInfo
}

// dBusNodeInfo is the struct that's finalized.
type dBusNodeInfo struct {
	native unsafe.Pointer
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusNodeInfo{&dBusNodeInfo{(unsafe.Pointer)(b)}}, nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (*DBusNodeInfo, error) {
	var _args [1]girepository.Argument

	*(**C.gchar)(unsafe.Pointer(&_args[0])) = (*C.gchar)(unsafe.Pointer(C.CString(xmlData)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[0]))))

	_info := girepository.MustFind("Gio", "DBusNodeInfo")
	_gret := _info.InvokeRecordMethod("new_for_xml", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(xmlData)

	var _dBusNodeInfo *DBusNodeInfo // out
	var _goerr error                // out

	_dBusNodeInfo = (*DBusNodeInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dBusNodeInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	if *(**C.void)(unsafe.Pointer(&_cerr)) != nil {
		_goerr = gerror.Take(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cerr))))
	}

	return _dBusNodeInfo, _goerr
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusNodeInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Path: path of the node or NULL if omitted. Note that this may be a relative
// path. See the D-Bus specification for more details.
func (d *DBusNodeInfo) Path() string {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("path")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Interfaces: pointer to a NULL-terminated array of pointers to
// BusInterfaceInfo structures or NULL if there are no interfaces.
func (d *DBusNodeInfo) Interfaces() []*DBusInterfaceInfo {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("interfaces")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusInterfaceInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusInterfaceInfo, i)
		for i := range src {
			v[i] = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_interface_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Nodes: pointer to a NULL-terminated array of pointers to BusNodeInfo
// structures or NULL if there are no nodes.
func (d *DBusNodeInfo) Nodes() []*DBusNodeInfo {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("nodes")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusNodeInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusNodeInfo, i)
		for i := range src {
			v[i] = (*DBusNodeInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_node_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusNodeInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusNodeInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusNodeInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
//
// The function takes the following parameters:
//
//    - name d-Bus interface name.
//
// The function returns the following values:
//
//    - dBusInterfaceInfo (optional) or NULL if not found. Do not free, it is
//      owned by info.
//
func (info *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var _args [2]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(gextras.StructNative(unsafe.Pointer(info)))
	*(**C.gchar)(unsafe.Pointer(&_args[1])) = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&_args[1]))))

	_info := girepository.MustFind("Gio", "DBusNodeInfo")
	_gret := _info.InvokeRecordMethod("lookup_interface", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	if *(**C.void)(unsafe.Pointer(&_cret)) != nil {
		_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&_cret)))))
		C.g_dbus_interface_info_ref(*(**C.void)(unsafe.Pointer(&_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _dBusInterfaceInfo
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusPropertyInfo struct {
	*dBusPropertyInfo
}

// dBusPropertyInfo is the struct that's finalized.
type dBusPropertyInfo struct {
	native unsafe.Pointer
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusPropertyInfo{&dBusPropertyInfo{(unsafe.Pointer)(b)}}, nil
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusSignalInfo struct {
	*dBusSignalInfo
}

// dBusSignalInfo is the struct that's finalized.
type dBusSignalInfo struct {
	native unsafe.Pointer
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusSignalInfo{&dBusSignalInfo{(unsafe.Pointer)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusSignalInfo) RefCount() int32 {
	offset := girepository.MustFind("Gio", "DBusSignalInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v int32 // out
	v = int32(*(*C.gint)(unsafe.Pointer(&*valptr)))
	return v
}

// Name: name of the D-Bus signal, e.g. "NameOwnerChanged".
func (d *DBusSignalInfo) Name() string {
	offset := girepository.MustFind("Gio", "DBusSignalInfo").StructFieldOffset("name")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*(**C.gchar)(unsafe.Pointer(&*valptr)))))
	return v
}

// Args: pointer to a NULL-terminated array of pointers to BusArgInfo structures
// or NULL if there are no arguments.
func (d *DBusSignalInfo) Args() []*DBusArgInfo {
	offset := girepository.MustFind("Gio", "DBusSignalInfo").StructFieldOffset("args")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_arg_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusSignalInfo) Annotations() []*DBusAnnotationInfo {
	offset := girepository.MustFind("Gio", "DBusSignalInfo").StructFieldOffset("annotations")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.void
		for p := *(***C.void)(unsafe.Pointer(&*valptr)); *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(*(***C.void)(unsafe.Pointer(&*valptr)), i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(*(**C.void)(unsafe.Pointer(&src[i])))))
			C.g_dbus_annotation_info_ref(*(**C.void)(unsafe.Pointer(&src[i])))
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.free(intern.C)
				},
			)
		}
	}
	return v
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusSignalInfo) SetRefCount(refCount int32) {
	offset := girepository.MustFind("Gio", "DBusSignalInfo").StructFieldOffset("ref_count")
	valptr := (*uintptr)(unsafe.Add(d.native, offset))
	*(*C.gint)(unsafe.Pointer(&*valptr)) = C.gint(refCount)
}
