// Code generated by girgen. DO NOT EDIT.

package glib

// #include <stdlib.h>
// #include <glib.h>
import "C"

// ASCII_DTOSTR_BUF_SIZE: good size for a buffer to be passed into
// g_ascii_dtostr(). It is guaranteed to be enough for all output of that
// function on systems with 64bit IEEE-compatible doubles.
//
// The typical usage would be something like:
//
//	char buf[G_ASCII_DTOSTR_BUF_SIZE];
//
//	fprintf (out, "value=s\n", g_ascii_dtostr (buf, sizeof (buf), value));.
const ASCII_DTOSTR_BUF_SIZE = 39

// STR_DELIMITERS: standard delimiters, used in g_strdelimit().
const STR_DELIMITERS = "_-|> <."
