// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// ApplicationCommandLineClass private data only.
//
// An instance of this type is always passed by reference.
type ApplicationCommandLineClass struct {
	*applicationCommandLineClass
}

// applicationCommandLineClass is the struct that's finalized.
type applicationCommandLineClass struct {
	native *C.GApplicationCommandLineClass
}