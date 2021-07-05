// Code generated by girgen. DO NOT EDIT.

package glib

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib.h>
import "C"

// BookmarkFileError: error codes returned by bookmark file parsing.
type BookmarkFileError int

const (
	// InvalidURI: URI was ill-formed
	BookmarkFileErrorInvalidURI BookmarkFileError = 0
	// InvalidValue: requested field was not found
	BookmarkFileErrorInvalidValue BookmarkFileError = 1
	// AppNotRegistered: requested application did not register a bookmark
	BookmarkFileErrorAppNotRegistered BookmarkFileError = 2
	// URINotFound: requested URI was not found
	BookmarkFileErrorURINotFound BookmarkFileError = 3
	// read: document was ill formed
	BookmarkFileErrorRead BookmarkFileError = 4
	// UnknownEncoding: the text being parsed was in an unknown encoding
	BookmarkFileErrorUnknownEncoding BookmarkFileError = 5
	// write: error occurred while writing
	BookmarkFileErrorWrite BookmarkFileError = 6
	// FileNotFound: requested file was not found
	BookmarkFileErrorFileNotFound BookmarkFileError = 7
)
