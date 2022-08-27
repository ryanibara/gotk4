// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// FILE_ATTRIBUTE_STANDARD_IS_VOLATILE: key in the "standard" namespace for
// checking if a file is volatile. This is meant for opaque, non-POSIX-like
// backends to indicate that the URI is not persistent. Applications should look
// at FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET for the persistent URI.
//
// Corresponding AttributeType is G_FILE_ATTRIBUTE_TYPE_BOOLEAN.
const FILE_ATTRIBUTE_STANDARD_IS_VOLATILE = "standard::is-volatile"
