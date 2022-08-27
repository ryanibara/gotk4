// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void _gotk4_gio2_Resolver_ConnectReload(gpointer, guintptr);
// extern void _gotk4_gio2_ResolverClass_reload(GResolver*);
// extern void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
// extern gchar* _gotk4_gio2_ResolverClass_lookup_by_address_finish(GResolver*, GAsyncResult*, GError**);
// extern gchar* _gotk4_gio2_ResolverClass_lookup_by_address(GResolver*, GInetAddress*, GCancellable*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_service_finish(GResolver*, GAsyncResult*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_records_finish(GResolver*, GAsyncResult*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_records(GResolver*, gchar*, GResolverRecordType, GCancellable*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_by_name_with_flags_finish(GResolver*, GAsyncResult*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_by_name_with_flags(GResolver*, gchar*, GResolverNameLookupFlags, GCancellable*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_by_name_finish(GResolver*, GAsyncResult*, GError**);
// extern GList* _gotk4_gio2_ResolverClass_lookup_by_name(GResolver*, gchar*, GCancellable*, GError**);
// void _gotk4_gio2_Resolver_virtual_lookup_service_async(void* fnptr, GResolver* arg0, gchar* arg1, GCancellable* arg2, GAsyncReadyCallback arg3, gpointer arg4) {
//   ((void (*)(GResolver*, gchar*, GCancellable*, GAsyncReadyCallback, gpointer))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// void _gotk4_gio2_Resolver_virtual_reload(void* fnptr, GResolver* arg0) {
//   ((void (*)(GResolver*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeResolver = coreglib.Type(C.g_resolver_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeResolver, F: marshalResolver},
	})
}

// ResolverOverrides contains methods that are overridable.
type ResolverOverrides struct {
	// LookupByAddress: synchronously reverse-resolves address to determine its
	// associated hostname.
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//    - address to reverse-resolve.
	//
	// The function returns the following values:
	//
	//    - utf8: hostname (either ASCII-only, or in ASCII-encoded form), or NULL
	//      on error.
	//
	LookupByAddress func(ctx context.Context, address *InetAddress) (string, error)
	// LookupByAddressFinish retrieves the result of a previous call to
	// g_resolver_lookup_by_address_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - result passed to your ReadyCallback.
	//
	// The function returns the following values:
	//
	//    - utf8: hostname (either ASCII-only, or in ASCII-encoded form), or NULL
	//      on error.
	//
	LookupByAddressFinish func(result AsyncResulter) (string, error)
	// LookupByName: synchronously resolves hostname to determine its associated
	// IP address(es). hostname may be an ASCII-only or UTF-8 hostname, or the
	// textual form of an IP address (in which case this just becomes a wrapper
	// around g_inet_address_new_from_string()).
	//
	// On success, g_resolver_lookup_by_name() will return a non-empty #GList of
	// Address, sorted in order of preference and guaranteed to not contain
	// duplicates. That is, if using the result to connect to hostname, you
	// should attempt to connect to the first address first, then the second if
	// the first fails, etc. If you are using the result to listen on a socket,
	// it is appropriate to add each result using e.g.
	// g_socket_listener_add_address().
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error and NULL will be returned.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	//
	// If you are planning to connect to a socket on the resolved IP address, it
	// may be easier to create a Address and use its Connectable interface.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//    - hostname to look up.
	//
	// The function returns the following values:
	//
	//    - list: non-empty #GList of Address, or NULL on error. You must unref
	//      each of the addresses and free the list when you are done with it.
	//      (You can use g_resolver_free_addresses() to do this.).
	//
	LookupByName func(ctx context.Context, hostname string) ([]*InetAddress, error)
	// LookupByNameFinish retrieves the result of a call to
	// g_resolver_lookup_by_name_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - result passed to your ReadyCallback.
	//
	// The function returns the following values:
	//
	//    - list Address, or NULL on error. See g_resolver_lookup_by_name() for
	//      more details.
	//
	LookupByNameFinish func(result AsyncResulter) ([]*InetAddress, error)
	// LookupByNameWithFlags: this differs from g_resolver_lookup_by_name() in
	// that you can modify the lookup behavior with flags. For example this can
	// be used to limit results with RESOLVER_NAME_LOOKUP_FLAGS_IPV4_ONLY.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//    - hostname to look up.
	//    - flags: extra NameLookupFlags for the lookup.
	//
	// The function returns the following values:
	//
	//    - list: non-empty #GList of Address, or NULL on error. You must unref
	//      each of the addresses and free the list when you are done with it.
	//      (You can use g_resolver_free_addresses() to do this.).
	//
	LookupByNameWithFlags func(ctx context.Context, hostname string, flags ResolverNameLookupFlags) ([]*InetAddress, error)
	// LookupByNameWithFlagsFinish retrieves the result of a call to
	// g_resolver_lookup_by_name_with_flags_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - result passed to your ReadyCallback.
	//
	// The function returns the following values:
	//
	//    - list Address, or NULL on error. See g_resolver_lookup_by_name() for
	//      more details.
	//
	LookupByNameWithFlagsFinish func(result AsyncResulter) ([]*InetAddress, error)
	// LookupRecords: synchronously performs a DNS record lookup for the given
	// rrname and returns a list of records as #GVariant tuples. See RecordType
	// for information on what the records contain for each record_type.
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error and NULL will be returned.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional) or NULL.
	//    - rrname: DNS name to look up the record for.
	//    - recordType: type of DNS record to look up.
	//
	// The function returns the following values:
	//
	//    - list: non-empty #GList of #GVariant, or NULL on error. You must free
	//      each of the records and the list when you are done with it. (You can
	//      use g_list_free_full() with g_variant_unref() to do this.).
	//
	LookupRecords func(ctx context.Context, rrname string, recordType ResolverRecordType) ([]*glib.Variant, error)
	// LookupRecordsFinish retrieves the result of a previous call to
	// g_resolver_lookup_records_async(). Returns a non-empty list of records as
	// #GVariant tuples. See RecordType for information on what the records
	// contain.
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - result passed to your ReadyCallback.
	//
	// The function returns the following values:
	//
	//    - list: non-empty #GList of #GVariant, or NULL on error. You must free
	//      each of the records and the list when you are done with it. (You can
	//      use g_list_free_full() with g_variant_unref() to do this.).
	//
	LookupRecordsFinish func(result AsyncResulter) ([]*glib.Variant, error)
	// LookupServiceFinish retrieves the result of a previous call to
	// g_resolver_lookup_service_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	//
	// The function takes the following parameters:
	//
	//    - result passed to your ReadyCallback.
	//
	// The function returns the following values:
	//
	//    - list: non-empty #GList of Target, or NULL on error. See
	//      g_resolver_lookup_service() for more details.
	//
	LookupServiceFinish func(result AsyncResulter) ([]*SrvTarget, error)
	Reload              func()
}

func defaultResolverOverrides(v *Resolver) ResolverOverrides {
	return ResolverOverrides{
		LookupByAddress:             v.lookupByAddress,
		LookupByAddressFinish:       v.lookupByAddressFinish,
		LookupByName:                v.lookupByName,
		LookupByNameFinish:          v.lookupByNameFinish,
		LookupByNameWithFlags:       v.lookupByNameWithFlags,
		LookupByNameWithFlagsFinish: v.lookupByNameWithFlagsFinish,
		LookupRecords:               v.lookupRecords,
		LookupRecordsFinish:         v.lookupRecordsFinish,
		LookupServiceFinish:         v.lookupServiceFinish,
		Reload:                      v.reload,
	}
}

// Resolver provides cancellable synchronous and asynchronous DNS resolution,
// for hostnames (g_resolver_lookup_by_address(), g_resolver_lookup_by_name()
// and their async variants) and SRV (service) records
// (g_resolver_lookup_service()).
//
// Address and Service provide wrappers around #GResolver functionality that
// also implement Connectable, making it easy to connect to a remote
// host/service.
type Resolver struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Resolver)(nil)
)

// Resolverer describes types inherited from class Resolver.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type Resolverer interface {
	coreglib.Objector
	baseResolver() *Resolver
}

var _ Resolverer = (*Resolver)(nil)

func init() {
	coreglib.RegisterClassInfo[*Resolver, *ResolverClass, ResolverOverrides](
		GTypeResolver,
		initResolverClass,
		wrapResolver,
		defaultResolverOverrides,
	)
}

func initResolverClass(gclass unsafe.Pointer, overrides ResolverOverrides, classInitFunc func(*ResolverClass)) {
	pclass := (*C.GResolverClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeResolver))))

	if overrides.LookupByAddress != nil {
		pclass.lookup_by_address = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_address)
	}

	if overrides.LookupByAddressFinish != nil {
		pclass.lookup_by_address_finish = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_address_finish)
	}

	if overrides.LookupByName != nil {
		pclass.lookup_by_name = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_name)
	}

	if overrides.LookupByNameFinish != nil {
		pclass.lookup_by_name_finish = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_name_finish)
	}

	if overrides.LookupByNameWithFlags != nil {
		pclass.lookup_by_name_with_flags = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_name_with_flags)
	}

	if overrides.LookupByNameWithFlagsFinish != nil {
		pclass.lookup_by_name_with_flags_finish = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_by_name_with_flags_finish)
	}

	if overrides.LookupRecords != nil {
		pclass.lookup_records = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_records)
	}

	if overrides.LookupRecordsFinish != nil {
		pclass.lookup_records_finish = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_records_finish)
	}

	if overrides.LookupServiceFinish != nil {
		pclass.lookup_service_finish = (*[0]byte)(C._gotk4_gio2_ResolverClass_lookup_service_finish)
	}

	if overrides.Reload != nil {
		pclass.reload = (*[0]byte)(C._gotk4_gio2_ResolverClass_reload)
	}

	if classInitFunc != nil {
		class := (*ResolverClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapResolver(obj *coreglib.Object) *Resolver {
	return &Resolver{
		Object: obj,
	}
}

func marshalResolver(p uintptr) (interface{}, error) {
	return wrapResolver(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (resolver *Resolver) baseResolver() *Resolver {
	return resolver
}

// BaseResolver returns the underlying base object.
func BaseResolver(obj Resolverer) *Resolver {
	return obj.baseResolver()
}

// ConnectReload is emitted when the resolver notices that the system resolver
// configuration has changed.
func (resolver *Resolver) ConnectReload(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(resolver, "reload", false, unsafe.Pointer(C._gotk4_gio2_Resolver_ConnectReload), f)
}

// The function takes the following parameters:
//
//    - ctx (optional)
//    - rrname
//    - callback (optional)
//
func (resolver *Resolver) lookupServiceAsync(ctx context.Context, rrname string, callback AsyncReadyCallback) {
	gclass := (*C.GResolverClass)(coreglib.PeekParentClass(resolver))
	fnarg := gclass.lookup_service_async

	var _arg0 *C.GResolver          // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(rrname)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C._gotk4_gio2_Resolver_virtual_lookup_service_async(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(rrname)
	runtime.KeepAlive(callback)
}

func (resolver *Resolver) reload() {
	gclass := (*C.GResolverClass)(coreglib.PeekParentClass(resolver))
	fnarg := gclass.reload

	var _arg0 *C.GResolver // out

	_arg0 = (*C.GResolver)(unsafe.Pointer(coreglib.InternObject(resolver).Native()))

	C._gotk4_gio2_Resolver_virtual_reload(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(resolver)
}

// ResolverClass: instance of this type is always passed by reference.
type ResolverClass struct {
	*resolverClass
}

// resolverClass is the struct that's finalized.
type resolverClass struct {
	native *C.GResolverClass
}
