// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// extern gboolean _gotk4_gio2_AsyncInitableIface_init_finish(GAsyncInitable*, GAsyncResult*, GError**);
import "C"

// glib.Type values for gasyncinitable.go.
var GTypeAsyncInitable = coreglib.Type(C.g_async_initable_get_type())

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		{T: GTypeAsyncInitable, F: marshalAsyncInitable},
	})
}

// AsyncInitableOverrider contains methods that are overridable.
type AsyncInitableOverrider interface {
	// InitFinish finishes asynchronous initialization and returns the result.
	// See g_async_initable_init_async().
	//
	// The function takes the following parameters:
	//
	//    - res: Result.
	//
	InitFinish(res AsyncResulter) error
}

// AsyncInitable: this is the asynchronous version of #GInitable; it behaves the
// same in all ways except that initialization is asynchronous. For more details
// see the descriptions on #GInitable.
//
// A class may implement both the #GInitable and Initable interfaces.
//
// Users of objects implementing this are not intended to use the interface
// method directly; instead it will be used automatically in various ways. For C
// applications you generally just call g_async_initable_new_async() directly,
// or indirectly via a foo_thing_new_async() wrapper. This will call
// g_async_initable_init_async() under the cover, calling back with NULL and a
// set GError on failure.
//
// A typical implementation might look something like this:
//
//    enum {
//       NOT_INITIALIZED,
//       INITIALIZING,
//       INITIALIZED
//    };
//
//    static void
//    _foo_ready_cb (Foo *self)
//    {
//      GList *l;
//
//      self->priv->state = INITIALIZED;
//
//      for (l = self->priv->init_results; l != NULL; l = l->next)
//        {
//          GTask *task = l->data;
//
//          if (self->priv->success)
//            g_task_return_boolean (task, TRUE);
//          else
//            g_task_return_new_error (task, ...);
//          g_object_unref (task);
//        }
//
//      g_list_free (self->priv->init_results);
//      self->priv->init_results = NULL;
//    }
//
//    static void
//    foo_init_async (GAsyncInitable       *initable,
//                    int                   io_priority,
//                    GCancellable         *cancellable,
//                    GAsyncReadyCallback   callback,
//                    gpointer              user_data)
//    {
//      Foo *self = FOO (initable);
//      GTask *task;
//
//      task = g_task_new (initable, cancellable, callback, user_data);
//      g_task_set_name (task, G_STRFUNC);
//
//      switch (self->priv->state)
//        {
//          case NOT_INITIALIZED:
//            _foo_get_ready (self);
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            self->priv->state = INITIALIZING;
//            break;
//          case INITIALIZING:
//            self->priv->init_results = g_list_append (self->priv->init_results,
//                                                      task);
//            break;
//          case INITIALIZED:
//            if (!self->priv->success)
//              g_task_return_new_error (task, ...);
//            else
//              g_task_return_boolean (task, TRUE);
//            g_object_unref (task);
//            break;
//        }
//    }
//
//    static gboolean
//    foo_init_finish (GAsyncInitable       *initable,
//                     GAsyncResult         *result,
//                     GError              **error)
//    {
//      g_return_val_if_fail (g_task_is_valid (result, initable), FALSE);
//
//      return g_task_propagate_boolean (G_TASK (result), error);
//    }
//
//    static void
//    foo_async_initable_iface_init (gpointer g_iface,
//                                   gpointer data)
//    {
//      GAsyncInitableIface *iface = g_iface;
//
//      iface->init_async = foo_init_async;
//      iface->init_finish = foo_init_finish;
//    }.
//
// AsyncInitable wraps an interface. This means the user can get the
// underlying type by calling Cast().
type AsyncInitable struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*AsyncInitable)(nil)
)

// AsyncInitabler describes AsyncInitable's interface methods.
type AsyncInitabler interface {
	coreglib.Objector

	// InitFinish finishes asynchronous initialization and returns the result.
	InitFinish(res AsyncResulter) error
	// NewFinish finishes the async construction for the various
	// g_async_initable_new calls, returning the created object or NULL on
	// error.
	NewFinish(res AsyncResulter) (*coreglib.Object, error)
}

var _ AsyncInitabler = (*AsyncInitable)(nil)

func ifaceInitAsyncInitabler(gifacePtr, data C.gpointer) {
	iface := (*C.GAsyncInitableIface)(unsafe.Pointer(gifacePtr))
	iface.init_finish = (*[0]byte)(C._gotk4_gio2_AsyncInitableIface_init_finish)
}

//export _gotk4_gio2_AsyncInitableIface_init_finish
func _gotk4_gio2_AsyncInitableIface_init_finish(arg0 *C.GAsyncInitable, arg1 *C.GAsyncResult, _cerr **C.GError) (cret C.gboolean) {
	goval := coreglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(AsyncInitableOverrider)

	var _res AsyncResulter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gio.AsyncResulter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AsyncResulter)
			return ok
		})
		rv, ok := casted.(AsyncResulter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AsyncResulter")
		}
		_res = rv
	}

	_goerr := iface.InitFinish(_res)

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.void)(gerror.New(_goerr))
	}

	return cret
}

func wrapAsyncInitable(obj *coreglib.Object) *AsyncInitable {
	return &AsyncInitable{
		Object: obj,
	}
}

func marshalAsyncInitable(p uintptr) (interface{}, error) {
	return wrapAsyncInitable(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// InitFinish finishes asynchronous initialization and returns the result. See
// g_async_initable_init_async().
//
// The function takes the following parameters:
//
//    - res: Result.
//
func (initable *AsyncInitable) InitFinish(res AsyncResulter) error {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))
	*(**AsyncInitable)(unsafe.Pointer(&args[1])) = _arg1

	runtime.KeepAlive(initable)
	runtime.KeepAlive(res)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// NewFinish finishes the async construction for the various
// g_async_initable_new calls, returning the created object or NULL on error.
//
// The function takes the following parameters:
//
//    - res from the callback.
//
// The function returns the following values:
//
//    - object: newly created #GObject, or NULL on error. Free with
//      g_object_unref().
//
func (initable *AsyncInitable) NewFinish(res AsyncResulter) (*coreglib.Object, error) {
	var args [2]girepository.Argument
	var _arg0 *C.void // out
	var _arg1 *C.void // out
	var _cret *C.void // in
	var _cerr *C.void // in

	_arg0 = (*C.void)(unsafe.Pointer(coreglib.InternObject(initable).Native()))
	_arg1 = (*C.void)(unsafe.Pointer(coreglib.InternObject(res).Native()))
	*(**AsyncInitable)(unsafe.Pointer(&args[1])) = _arg1

	_cret = *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(initable)
	runtime.KeepAlive(res)

	var _object *coreglib.Object // out
	var _goerr error             // out

	_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _object, _goerr
}
