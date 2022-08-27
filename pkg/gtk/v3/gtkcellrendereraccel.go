// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellRendererAccel_ConnectAccelEdited(gpointer, gchar*, guint, GdkModifierType, guint, guintptr);
// extern void _gotk4_gtk3_CellRendererAccel_ConnectAccelCleared(gpointer, gchar*, guintptr);
// extern void _gotk4_gtk3_CellRendererAccelClass_accel_edited(GtkCellRendererAccel*, gchar*, guint, GdkModifierType, guint);
// extern void _gotk4_gtk3_CellRendererAccelClass_accel_cleared(GtkCellRendererAccel*, gchar*);
// void _gotk4_gtk3_CellRendererAccel_virtual_accel_cleared(void* fnptr, GtkCellRendererAccel* arg0, gchar* arg1) {
//   ((void (*)(GtkCellRendererAccel*, gchar*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gtk3_CellRendererAccel_virtual_accel_edited(void* fnptr, GtkCellRendererAccel* arg0, gchar* arg1, guint arg2, GdkModifierType arg3, guint arg4) {
//   ((void (*)(GtkCellRendererAccel*, gchar*, guint, GdkModifierType, guint))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
import "C"

// GType values.
var (
	GTypeCellRendererAccelMode = coreglib.Type(C.gtk_cell_renderer_accel_mode_get_type())
	GTypeCellRendererAccel     = coreglib.Type(C.gtk_cell_renderer_accel_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeCellRendererAccelMode, F: marshalCellRendererAccelMode},
		coreglib.TypeMarshaler{T: GTypeCellRendererAccel, F: marshalCellRendererAccel},
	})
}

// CellRendererAccelMode determines if the edited accelerators are GTK+
// accelerators. If they are, consumed modifiers are suppressed, only
// accelerators accepted by GTK+ are allowed, and the accelerators are rendered
// in the same way as they are in menus.
type CellRendererAccelMode C.gint

const (
	// CellRendererAccelModeGTK: GTK+ accelerators mode.
	CellRendererAccelModeGTK CellRendererAccelMode = iota
	// CellRendererAccelModeOther: other accelerator mode.
	CellRendererAccelModeOther
)

func marshalCellRendererAccelMode(p uintptr) (interface{}, error) {
	return CellRendererAccelMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for CellRendererAccelMode.
func (c CellRendererAccelMode) String() string {
	switch c {
	case CellRendererAccelModeGTK:
		return "GTK"
	case CellRendererAccelModeOther:
		return "Other"
	default:
		return fmt.Sprintf("CellRendererAccelMode(%d)", c)
	}
}

// CellRendererAccelOverrides contains methods that are overridable.
type CellRendererAccelOverrides struct {
	// The function takes the following parameters:
	//
	AccelCleared func(pathString string)
	// The function takes the following parameters:
	//
	//    - pathString
	//    - accelKey
	//    - accelMods
	//    - hardwareKeycode
	//
	AccelEdited func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)
}

func defaultCellRendererAccelOverrides(v *CellRendererAccel) CellRendererAccelOverrides {
	return CellRendererAccelOverrides{
		AccelCleared: v.accelCleared,
		AccelEdited:  v.accelEdited,
	}
}

// CellRendererAccel displays a keyboard accelerator (i.e. a key combination
// like Control + a). If the cell renderer is editable, the accelerator can be
// changed by simply typing the new combination.
//
// The CellRendererAccel cell renderer was added in GTK+ 2.10.
type CellRendererAccel struct {
	_ [0]func() // equal guard
	CellRendererText
}

var (
	_ CellRendererer = (*CellRendererAccel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*CellRendererAccel, *CellRendererAccelClass, CellRendererAccelOverrides](
		GTypeCellRendererAccel,
		initCellRendererAccelClass,
		wrapCellRendererAccel,
		defaultCellRendererAccelOverrides,
	)
}

func initCellRendererAccelClass(gclass unsafe.Pointer, overrides CellRendererAccelOverrides, classInitFunc func(*CellRendererAccelClass)) {
	pclass := (*C.GtkCellRendererAccelClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeCellRendererAccel))))

	if overrides.AccelCleared != nil {
		pclass.accel_cleared = (*[0]byte)(C._gotk4_gtk3_CellRendererAccelClass_accel_cleared)
	}

	if overrides.AccelEdited != nil {
		pclass.accel_edited = (*[0]byte)(C._gotk4_gtk3_CellRendererAccelClass_accel_edited)
	}

	if classInitFunc != nil {
		class := (*CellRendererAccelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapCellRendererAccel(obj *coreglib.Object) *CellRendererAccel {
	return &CellRendererAccel{
		CellRendererText: CellRendererText{
			CellRenderer: CellRenderer{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalCellRendererAccel(p uintptr) (interface{}, error) {
	return wrapCellRendererAccel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAccelCleared gets emitted when the user has removed the accelerator.
func (v *CellRendererAccel) ConnectAccelCleared(f func(pathString string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "accel-cleared", false, unsafe.Pointer(C._gotk4_gtk3_CellRendererAccel_ConnectAccelCleared), f)
}

// ConnectAccelEdited gets emitted when the user has selected a new accelerator.
func (v *CellRendererAccel) ConnectAccelEdited(f func(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "accel-edited", false, unsafe.Pointer(C._gotk4_gtk3_CellRendererAccel_ConnectAccelEdited), f)
}

// The function takes the following parameters:
//
func (accel *CellRendererAccel) accelCleared(pathString string) {
	gclass := (*C.GtkCellRendererAccelClass)(coreglib.PeekParentClass(accel))
	fnarg := gclass.accel_cleared

	var _arg0 *C.GtkCellRendererAccel // out
	var _arg1 *C.gchar                // out

	_arg0 = (*C.GtkCellRendererAccel)(unsafe.Pointer(coreglib.InternObject(accel).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pathString)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_gtk3_CellRendererAccel_virtual_accel_cleared(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(accel)
	runtime.KeepAlive(pathString)
}

// The function takes the following parameters:
//
//    - pathString
//    - accelKey
//    - accelMods
//    - hardwareKeycode
//
func (accel *CellRendererAccel) accelEdited(pathString string, accelKey uint, accelMods gdk.ModifierType, hardwareKeycode uint) {
	gclass := (*C.GtkCellRendererAccelClass)(coreglib.PeekParentClass(accel))
	fnarg := gclass.accel_edited

	var _arg0 *C.GtkCellRendererAccel // out
	var _arg1 *C.gchar                // out
	var _arg2 C.guint                 // out
	var _arg3 C.GdkModifierType       // out
	var _arg4 C.guint                 // out

	_arg0 = (*C.GtkCellRendererAccel)(unsafe.Pointer(coreglib.InternObject(accel).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pathString)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint(accelKey)
	_arg3 = C.GdkModifierType(accelMods)
	_arg4 = C.guint(hardwareKeycode)

	C._gotk4_gtk3_CellRendererAccel_virtual_accel_edited(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(accel)
	runtime.KeepAlive(pathString)
	runtime.KeepAlive(accelKey)
	runtime.KeepAlive(accelMods)
	runtime.KeepAlive(hardwareKeycode)
}

// CellRendererAccelClass: instance of this type is always passed by reference.
type CellRendererAccelClass struct {
	*cellRendererAccelClass
}

// cellRendererAccelClass is the struct that's finalized.
type cellRendererAccelClass struct {
	native *C.GtkCellRendererAccelClass
}

func (c *CellRendererAccelClass) ParentClass() *CellRendererTextClass {
	valptr := &c.native.parent_class
	var _v *CellRendererTextClass // out
	_v = (*CellRendererTextClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
