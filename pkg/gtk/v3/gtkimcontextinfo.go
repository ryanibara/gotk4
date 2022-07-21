// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"
)

// #include <stdlib.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// IMContextInfo: bookkeeping information about a loadable input method.
//
// An instance of this type is always passed by reference.
type IMContextInfo struct {
	*imContextInfo
}

// imContextInfo is the struct that's finalized.
type imContextInfo struct {
	native *C.GtkIMContextInfo
}

// ContextID: unique identification string of the input method.
func (i *IMContextInfo) ContextID() string {
	valptr := &i.native.context_id
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}

// ContextName: human-readable name of the input method.
func (i *IMContextInfo) ContextName() string {
	valptr := &i.native.context_name
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}

// Domain: translation domain to be used with dgettext().
func (i *IMContextInfo) Domain() string {
	valptr := &i.native.domain
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}

// DomainDirname: name of locale directory for use with bindtextdomain().
func (i *IMContextInfo) DomainDirname() string {
	valptr := &i.native.domain_dirname
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}

// DefaultLocales: colon-separated list of locales where this input method
// should be the default. The asterisk “*” sets the default for all locales.
func (i *IMContextInfo) DefaultLocales() string {
	valptr := &i.native.default_locales
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(*valptr)))
	return v
}
