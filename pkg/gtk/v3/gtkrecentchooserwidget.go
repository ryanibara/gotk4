// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeRecentChooserWidget returns the GType for the type RecentChooserWidget.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeRecentChooserWidget() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Gtk", "RecentChooserWidget").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalRecentChooserWidget)
	return gtype
}

// RecentChooserWidgetOverrider contains methods that are overridable.
type RecentChooserWidgetOverrider interface {
}

// RecentChooserWidget is a widget suitable for selecting recently used files.
// It is the main building block of a RecentChooserDialog. Most applications
// will only need to use the latter; you can use RecentChooserWidget as part of
// a larger window if you have special needs.
//
// Note that RecentChooserWidget does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserWidget struct {
	_ [0]func() // equal guard
	Box

	*coreglib.Object
	RecentChooser
}

var (
	_ coreglib.Objector = (*RecentChooserWidget)(nil)
	_ Containerer       = (*RecentChooserWidget)(nil)
)

func classInitRecentChooserWidgetter(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapRecentChooserWidget(obj *coreglib.Object) *RecentChooserWidget {
	return &RecentChooserWidget{
		Box: Box{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
					Object: obj,
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
		Object: obj,
		RecentChooser: RecentChooser{
			Object: obj,
		},
	}
}

func marshalRecentChooserWidget(p uintptr) (interface{}, error) {
	return wrapRecentChooserWidget(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewRecentChooserWidget creates a new RecentChooserWidget object. This is an
// embeddable widget used to access the recently used resources list.
//
// The function returns the following values:
//
//    - recentChooserWidget: new RecentChooserWidget.
//
func NewRecentChooserWidget() *RecentChooserWidget {
	_info := girepository.MustFind("Gtk", "RecentChooserWidget")
	_gret := _info.InvokeClassMethod("new_RecentChooserWidget", nil, nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}

// NewRecentChooserWidgetForManager creates a new RecentChooserWidget with a
// specified recent manager.
//
// This is useful if you have implemented your own recent manager, or if you
// have a customized instance of a RecentManager object.
//
// The function takes the following parameters:
//
//    - manager: RecentManager.
//
// The function returns the following values:
//
//    - recentChooserWidget: new RecentChooserWidget.
//
func NewRecentChooserWidgetForManager(manager *RecentManager) *RecentChooserWidget {
	var _args [1]girepository.Argument

	*(**C.void)(unsafe.Pointer(&_args[0])) = (*C.void)(unsafe.Pointer(coreglib.InternObject(manager).Native()))

	_info := girepository.MustFind("Gtk", "RecentChooserWidget")
	_gret := _info.InvokeClassMethod("new_RecentChooserWidget_for_manager", _args[:], nil)
	_cret := *(**C.void)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(manager)

	var _recentChooserWidget *RecentChooserWidget // out

	_recentChooserWidget = wrapRecentChooserWidget(coreglib.Take(unsafe.Pointer(_cret)))

	return _recentChooserWidget
}
