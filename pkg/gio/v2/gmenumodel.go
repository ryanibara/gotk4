// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// void _gotk4_gio2_MenuModel_virtual_get_item_attributes(void* fnptr, GMenuModel* arg0, gint arg1, GHashTable** arg2) {
//   ((void (*)(GMenuModel*, gint, GHashTable**))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_MenuModel_virtual_get_item_links(void* fnptr, GMenuModel* arg0, gint arg1, GHashTable** arg2) {
//   ((void (*)(GMenuModel*, gint, GHashTable**))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// itemAttributes gets all the attributes associated with the item in the menu
// model.
//
// The function takes the following parameters:
//
//    - itemIndex to query.
//
// The function returns the following values:
//
//    - attributes attributes on the item.
//
func (model *MenuModel) itemAttributes(itemIndex int) map[string]*glib.Variant {
	gclass := (*C.GMenuModelClass)(coreglib.PeekParentClass(model))
	fnarg := gclass.get_item_attributes

	var _arg0 *C.GMenuModel // out
	var _arg1 C.gint        // out
	var _arg2 *C.GHashTable // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	_arg1 = C.gint(itemIndex)

	C._gotk4_gio2_MenuModel_virtual_get_item_attributes(unsafe.Pointer(fnarg), _arg0, _arg1, &_arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)

	var _attributes map[string]*glib.Variant // out

	_attributes = make(map[string]*glib.Variant, gextras.HashTableSize(unsafe.Pointer(_arg2)))
	gextras.MoveHashTable(unsafe.Pointer(_arg2), true, func(k, v unsafe.Pointer) {
		ksrc := *(**C.gchar)(k)
		vsrc := *(**C.GVariant)(v)
		var kdst string        // out
		var vdst *glib.Variant // out
		kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
		vdst = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(vsrc)))
		C.g_variant_ref(vsrc)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(vdst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
		_attributes[kdst] = vdst
	})

	return _attributes
}

// itemLinks gets all the links associated with the item in the menu model.
//
// The function takes the following parameters:
//
//    - itemIndex to query.
//
// The function returns the following values:
//
//    - links links from the item.
//
func (model *MenuModel) itemLinks(itemIndex int) map[string]MenuModeller {
	gclass := (*C.GMenuModelClass)(coreglib.PeekParentClass(model))
	fnarg := gclass.get_item_links

	var _arg0 *C.GMenuModel // out
	var _arg1 C.gint        // out
	var _arg2 *C.GHashTable // in

	_arg0 = (*C.GMenuModel)(unsafe.Pointer(coreglib.InternObject(model).Native()))
	_arg1 = C.gint(itemIndex)

	C._gotk4_gio2_MenuModel_virtual_get_item_links(unsafe.Pointer(fnarg), _arg0, _arg1, &_arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(itemIndex)

	var _links map[string]MenuModeller // out

	_links = make(map[string]MenuModeller, gextras.HashTableSize(unsafe.Pointer(_arg2)))
	gextras.MoveHashTable(unsafe.Pointer(_arg2), true, func(k, v unsafe.Pointer) {
		ksrc := *(**C.gchar)(k)
		vsrc := *(**C.GMenuModel)(v)
		var kdst string       // out
		var vdst MenuModeller // out
		kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
		{
			objptr := unsafe.Pointer(vsrc)
			if objptr == nil {
				panic("object of type gio.MenuModeller is nil")
			}

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(MenuModeller)
				return ok
			})
			rv, ok := casted.(MenuModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
			}
			vdst = rv
		}
		_links[kdst] = vdst
	})

	return _links
}

// MenuAttributeIterClass: instance of this type is always passed by reference.
type MenuAttributeIterClass struct {
	*menuAttributeIterClass
}

// menuAttributeIterClass is the struct that's finalized.
type menuAttributeIterClass struct {
	native *C.GMenuAttributeIterClass
}

// MenuLinkIterClass: instance of this type is always passed by reference.
type MenuLinkIterClass struct {
	*menuLinkIterClass
}

// menuLinkIterClass is the struct that's finalized.
type menuLinkIterClass struct {
	native *C.GMenuLinkIterClass
}

// MenuModelClass: instance of this type is always passed by reference.
type MenuModelClass struct {
	*menuModelClass
}

// menuModelClass is the struct that's finalized.
type menuModelClass struct {
	native *C.GMenuModelClass
}
