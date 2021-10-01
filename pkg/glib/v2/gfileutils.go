// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// FileError values corresponding to errno codes returned from file operations
// on UNIX. Unlike errno codes, GFileError values are available on all systems,
// even Windows. The exact meaning of each code depends on what sort of file
// operation you were performing; the UNIX documentation gives more details. The
// following error code descriptions come from the GNU C Library manual, and are
// under the copyright of that manual.
//
// It's not very portable to make detailed assumptions about exactly which
// errors will be returned from a given operation. Some errors don't occur on
// some systems, etc., sometimes there are subtle differences in when a system
// will report a given error, etc.
type FileError int

const (
	// FileErrorExist: operation not permitted; only the owner of the file (or
	// other resource) or processes with special privileges can perform the
	// operation.
	FileErrorExist FileError = iota
	// FileErrorIsdir: file is a directory; you cannot open a directory for
	// writing, or create or remove hard links to it.
	FileErrorIsdir
	// FileErrorAcces: permission denied; the file permissions do not allow the
	// attempted operation.
	FileErrorAcces
	// FileErrorNametoolong: filename too long.
	FileErrorNametoolong
	// FileErrorNoent: no such file or directory. This is a "file doesn't exist"
	// error for ordinary files that are referenced in contexts where they are
	// expected to already exist.
	FileErrorNoent
	// FileErrorNotdir: file that isn't a directory was specified when a
	// directory is required.
	FileErrorNotdir
	// FileErrorNxio: no such device or address. The system tried to use the
	// device represented by a file you specified, and it couldn't find the
	// device. This can mean that the device file was installed incorrectly, or
	// that the physical device is missing or not correctly attached to the
	// computer.
	FileErrorNxio
	// FileErrorNodev: underlying file system of the specified file does not
	// support memory mapping.
	FileErrorNodev
	// FileErrorRofs: directory containing the new link can't be modified
	// because it's on a read-only file system.
	FileErrorRofs
	// FileErrorTxtbsy: text file busy.
	FileErrorTxtbsy
	// FileErrorFault: you passed in a pointer to bad memory. (GLib won't
	// reliably return this, don't pass in pointers to bad memory.).
	FileErrorFault
	// FileErrorLoop: too many levels of symbolic links were encountered in
	// looking up a file name. This often indicates a cycle of symbolic links.
	FileErrorLoop
	// FileErrorNospc: no space left on device; write operation on a file failed
	// because the disk is full.
	FileErrorNospc
	// FileErrorNOMEM: no memory available. The system cannot allocate more
	// virtual memory because its capacity is full.
	FileErrorNOMEM
	// FileErrorMfile: current process has too many files open and can't open
	// any more. Duplicate descriptors do count toward this limit.
	FileErrorMfile
	// FileErrorNfile: there are too many distinct file openings in the entire
	// system.
	FileErrorNfile
	// FileErrorBadf: bad file descriptor; for example, I/O on a descriptor that
	// has been closed or reading from a descriptor open only for writing (or
	// vice versa).
	FileErrorBadf
	// FileErrorInval: invalid argument. This is used to indicate various kinds
	// of problems with passing the wrong argument to a library function.
	FileErrorInval
	// FileErrorPipe: broken pipe; there is no process reading from the other
	// end of a pipe. Every library function that returns this error code also
	// generates a 'SIGPIPE' signal; this signal terminates the program if not
	// handled or blocked. Thus, your program will never actually see this code
	// unless it has handled or blocked 'SIGPIPE'.
	FileErrorPipe
	// FileErrorAgain: resource temporarily unavailable; the call might work if
	// you try again later.
	FileErrorAgain
	// FileErrorIntr: interrupted function call; an asynchronous signal occurred
	// and prevented completion of the call. When this happens, you should try
	// the call again.
	FileErrorIntr
	// FileErrorIO: input/output error; usually used for physical read or write
	// errors. i.e. the disk or other physical device hardware is returning
	// errors.
	FileErrorIO
	// FileErrorPerm: operation not permitted; only the owner of the file (or
	// other resource) or processes with special privileges can perform the
	// operation.
	FileErrorPerm
	// FileErrorNosys: function not implemented; this indicates that the system
	// is missing some functionality.
	FileErrorNosys
	// FileErrorFailed does not correspond to a UNIX error code; this is the
	// standard "failed for unspecified reason" error code present in all
	// #GError error code enumerations. Returned if no specific code applies.
	FileErrorFailed
)

// String returns the name in string for FileError.
func (f FileError) String() string {
	switch f {
	case FileErrorExist:
		return "Exist"
	case FileErrorIsdir:
		return "Isdir"
	case FileErrorAcces:
		return "Acces"
	case FileErrorNametoolong:
		return "Nametoolong"
	case FileErrorNoent:
		return "Noent"
	case FileErrorNotdir:
		return "Notdir"
	case FileErrorNxio:
		return "Nxio"
	case FileErrorNodev:
		return "Nodev"
	case FileErrorRofs:
		return "Rofs"
	case FileErrorTxtbsy:
		return "Txtbsy"
	case FileErrorFault:
		return "Fault"
	case FileErrorLoop:
		return "Loop"
	case FileErrorNospc:
		return "Nospc"
	case FileErrorNOMEM:
		return "NOMEM"
	case FileErrorMfile:
		return "Mfile"
	case FileErrorNfile:
		return "Nfile"
	case FileErrorBadf:
		return "Badf"
	case FileErrorInval:
		return "Inval"
	case FileErrorPipe:
		return "Pipe"
	case FileErrorAgain:
		return "Again"
	case FileErrorIntr:
		return "Intr"
	case FileErrorIO:
		return "IO"
	case FileErrorPerm:
		return "Perm"
	case FileErrorNosys:
		return "Nosys"
	case FileErrorFailed:
		return "Failed"
	default:
		return fmt.Sprintf("FileError(%d)", f)
	}
}

// FileSetContentsFlags flags to pass to g_file_set_contents_full() to affect
// its safety and performance.
type FileSetContentsFlags int

const (
	// FileSetContentsNone: no guarantees about file consistency or durability.
	// The most dangerous setting, which is slightly faster than other settings.
	FileSetContentsNone FileSetContentsFlags = 0b0
	// FileSetContentsConsistent: guarantee file consistency: after a crash,
	// either the old version of the file or the new version of the file will be
	// available, but not a mixture. On Unix systems this equates to an fsync()
	// on the file and use of an atomic rename() of the new version of the file
	// over the old.
	FileSetContentsConsistent FileSetContentsFlags = 0b1
	// FileSetContentsDurable: guarantee file durability: after a crash, the new
	// version of the file will be available. On Unix systems this equates to an
	// fsync() on the file (if G_FILE_SET_CONTENTS_CONSISTENT is unset), or the
	// effects of G_FILE_SET_CONTENTS_CONSISTENT plus an fsync() on the
	// directory containing the file after calling rename().
	FileSetContentsDurable FileSetContentsFlags = 0b10
	// FileSetContentsOnlyExisting: only apply consistency and durability
	// guarantees if the file already exists. This may speed up file operations
	// if the file doesn’t currently exist, but may result in a corrupted
	// version of the new file if the system crashes while writing it.
	FileSetContentsOnlyExisting FileSetContentsFlags = 0b100
)

// String returns the names in string for FileSetContentsFlags.
func (f FileSetContentsFlags) String() string {
	if f == 0 {
		return "FileSetContentsFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(96)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FileSetContentsNone:
			builder.WriteString("None|")
		case FileSetContentsConsistent:
			builder.WriteString("Consistent|")
		case FileSetContentsDurable:
			builder.WriteString("Durable|")
		case FileSetContentsOnlyExisting:
			builder.WriteString("OnlyExisting|")
		default:
			builder.WriteString(fmt.Sprintf("FileSetContentsFlags(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FileSetContentsFlags) Has(other FileSetContentsFlags) bool {
	return (f & other) == other
}

// FileTest: test to perform on a file using g_file_test().
type FileTest int

const (
	// FileTestIsRegular: TRUE if the file is a regular file (not a directory).
	// Note that this test will also return TRUE if the tested file is a symlink
	// to a regular file.
	FileTestIsRegular FileTest = 0b1
	// FileTestIsSymlink: TRUE if the file is a symlink.
	FileTestIsSymlink FileTest = 0b10
	// FileTestIsDir: TRUE if the file is a directory.
	FileTestIsDir FileTest = 0b100
	// FileTestIsExecutable: TRUE if the file is executable.
	FileTestIsExecutable FileTest = 0b1000
	// FileTestExists: TRUE if the file exists. It may or may not be a regular
	// file.
	FileTestExists FileTest = 0b10000
)

// String returns the names in string for FileTest.
func (f FileTest) String() string {
	if f == 0 {
		return "FileTest(0)"
	}

	var builder strings.Builder
	builder.Grow(85)

	for f != 0 {
		next := f & (f - 1)
		bit := f - next

		switch bit {
		case FileTestIsRegular:
			builder.WriteString("IsRegular|")
		case FileTestIsSymlink:
			builder.WriteString("IsSymlink|")
		case FileTestIsDir:
			builder.WriteString("IsDir|")
		case FileTestIsExecutable:
			builder.WriteString("IsExecutable|")
		case FileTestExists:
			builder.WriteString("Exists|")
		default:
			builder.WriteString(fmt.Sprintf("FileTest(0b%b)|", bit))
		}

		f = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if f contains other.
func (f FileTest) Has(other FileTest) bool {
	return (f & other) == other
}

// Basename gets the name of the file without any leading directory components.
// It returns a pointer into the given file name string.
//
// Deprecated: Use g_path_get_basename() instead, but notice that
// g_path_get_basename() allocates new memory for the returned string, unlike
// this function which returns a pointer into the argument.
func Basename(fileName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_basename(_arg1)
	runtime.KeepAlive(fileName)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// BuildFilenamev behaves exactly like g_build_filename(), but takes the path
// elements as a string array, instead of varargs. This function is mainly meant
// for language bindings.
func BuildFilenamev(args []string) string {
	var _arg1 **C.gchar // out
	var _cret *C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(args)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(args)+1)
			var zero *C.gchar
			out[len(args)] = zero
			for i := range args {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(args[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.g_build_filenamev(_arg1)
	runtime.KeepAlive(args)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// BuildPathv behaves exactly like g_build_path(), but takes the path elements
// as a string array, instead of varargs. This function is mainly meant for
// language bindings.
func BuildPathv(separator string, args []string) string {
	var _arg1 *C.gchar  // out
	var _arg2 **C.gchar // out
	var _cret *C.gchar  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(separator)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.malloc(C.ulong(len(args)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(args)+1)
			var zero *C.gchar
			out[len(args)] = zero
			for i := range args {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(args[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.g_build_pathv(_arg1, _arg2)
	runtime.KeepAlive(separator)
	runtime.KeepAlive(args)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// CanonicalizeFilename gets the canonical file name from filename. All triple
// slashes are turned into single slashes, and all .. and .s resolved against
// relative_to.
//
// Symlinks are not followed, and the returned path is guaranteed to be
// absolute.
//
// If filename is an absolute path, relative_to is ignored. Otherwise,
// relative_to will be prepended to filename to make it absolute. relative_to
// must be an absolute path, or NULL. If relative_to is NULL, it'll fallback to
// g_get_current_dir().
//
// This function never fails, and will canonicalize file paths even if they
// don't exist.
//
// No file system I/O is done.
func CanonicalizeFilename(filename string, relativeTo string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	if relativeTo != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(relativeTo)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.g_canonicalize_filename(_arg1, _arg2)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(relativeTo)

	var _ret string // out

	_ret = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _ret
}

// FileErrorFromErrno gets a Error constant based on the passed-in err_no. For
// example, if you pass in EEXIST this function returns FILE_ERROR_EXIST. Unlike
// errno values, you can portably assume that all Error values will exist.
//
// Normally a Error value goes into a #GError returned from a function that
// manipulates files. So you would use g_file_error_from_errno() when
// constructing a #GError.
func FileErrorFromErrno(errNo int) FileError {
	var _arg1 C.gint       // out
	var _cret C.GFileError // in

	_arg1 = C.gint(errNo)

	_cret = C.g_file_error_from_errno(_arg1)
	runtime.KeepAlive(errNo)

	var _fileError FileError // out

	_fileError = FileError(_cret)

	return _fileError
}

// FileGetContents reads an entire file into allocated memory, with good error
// checking.
//
// If the call was successful, it returns TRUE and sets contents to the file
// contents and length to the length of the file contents in bytes. The string
// stored in contents will be nul-terminated, so for text files you can pass
// NULL for the length argument. If the call was not successful, it returns
// FALSE and sets error. The error domain is FILE_ERROR. Possible error codes
// are those in the Error enumeration. In the error case, contents is set to
// NULL and length is set to zero.
func FileGetContents(filename string) ([]byte, error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // in
	var _arg3 C.gsize   // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_get_contents(_arg1, &_arg2, &_arg3, &_cerr)
	runtime.KeepAlive(filename)

	var _contents []byte // out
	var _goerr error     // out

	defer C.free(unsafe.Pointer(_arg2))
	_contents = make([]byte, _arg3)
	copy(_contents, unsafe.Slice((*byte)(unsafe.Pointer(_arg2)), _arg3))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _contents, _goerr
}

// FileOpenTmp opens a file for writing in the preferred directory for temporary
// files (as returned by g_get_tmp_dir()).
//
// tmpl should be a string in the GLib file name encoding containing a sequence
// of six 'X' characters, as the parameter to g_mkstemp(). However, unlike these
// functions, the template should only be a basename, no directory components
// are allowed. If template is NULL, a default template is used.
//
// Note that in contrast to g_mkstemp() (and mkstemp()) tmpl is not modified,
// and might thus be a read-only literal string.
//
// Upon success, and if name_used is non-NULL, the actual name used is returned
// in name_used. This string should be freed with g_free() when not needed any
// longer. The returned name is in the GLib file name encoding.
func FileOpenTmp(tmpl string) (string, int, error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // in
	var _cret C.gint    // in
	var _cerr *C.GError // in

	if tmpl != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(tmpl)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.g_file_open_tmp(_arg1, &_arg2, &_cerr)
	runtime.KeepAlive(tmpl)

	var _nameUsed string // out
	var _gint int        // out
	var _goerr error     // out

	_nameUsed = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_gint = int(_cret)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _nameUsed, _gint, _goerr
}

// FileReadLink reads the contents of the symbolic link filename like the POSIX
// readlink() function. The returned string is in the encoding used for
// filenames. Use g_filename_to_utf8() to convert it to UTF-8.
func FileReadLink(filename string) (string, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_read_link(_arg1, &_cerr)
	runtime.KeepAlive(filename)

	var _ret string  // out
	var _goerr error // out

	_ret = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ret, _goerr
}

// FileSetContents writes all of contents to a file named filename. This is a
// convenience wrapper around calling g_file_set_contents_full() with flags set
// to G_FILE_SET_CONTENTS_CONSISTENT | G_FILE_SET_CONTENTS_ONLY_EXISTING and
// mode set to 0666.
func FileSetContents(filename string, contents []byte) error {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 C.gssize
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg3 = (C.gssize)(len(contents))
	if len(contents) > 0 {
		_arg2 = (*C.gchar)(unsafe.Pointer(&contents[0]))
	}

	C.g_file_set_contents(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(contents)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// FileSetContentsFull writes all of contents to a file named filename, with
// good error checking. If a file called filename already exists it will be
// overwritten.
//
// flags control the properties of the write operation: whether it’s atomic, and
// what the tradeoff is between returning quickly or being resilient to system
// crashes.
//
// As this function performs file I/O, it is recommended to not call it anywhere
// where blocking would cause problems, such as in the main loop of a graphical
// application. In particular, if flags has any value other than
// G_FILE_SET_CONTENTS_NONE then this function may call fsync().
//
// If G_FILE_SET_CONTENTS_CONSISTENT is set in flags, the operation is atomic in
// the sense that it is first written to a temporary file which is then renamed
// to the final name.
//
// Notes:
//
// - On UNIX, if filename already exists hard links to filename will break. Also
// since the file is recreated, existing permissions, access control lists,
// metadata etc. may be lost. If filename is a symbolic link, the link itself
// will be replaced, not the linked file.
//
// - On UNIX, if filename already exists and is non-empty, and if the system
// supports it (via a journalling filesystem or equivalent), and if
// G_FILE_SET_CONTENTS_CONSISTENT is set in flags, the fsync() call (or
// equivalent) will be used to ensure atomic replacement: filename will contain
// either its old contents or contents, even in the face of system power loss,
// the disk being unsafely removed, etc.
//
// - On UNIX, if filename does not already exist or is empty, there is a
// possibility that system power loss etc. after calling this function will
// leave filename empty or full of NUL bytes, depending on the underlying
// filesystem, unless G_FILE_SET_CONTENTS_DURABLE and
// G_FILE_SET_CONTENTS_CONSISTENT are set in flags.
//
// - On Windows renaming a file will not remove an existing file with the new
// name, so on Windows there is a race condition between the existing file being
// removed and the temporary file being renamed.
//
// - On Windows there is no way to remove a file that is open to some process,
// or mapped into memory. Thus, this function will fail if filename already
// exists and is open.
//
// If the call was successful, it returns TRUE. If the call was not successful,
// it returns FALSE and sets error. The error domain is FILE_ERROR. Possible
// error codes are those in the Error enumeration.
//
// Note that the name for the temporary file is constructed by appending up to 7
// characters to filename.
//
// If the file didn’t exist before and is created, it will be given the
// permissions from mode. Otherwise, the permissions of the existing file may be
// changed to mode depending on flags, or they may remain unchanged.
func FileSetContentsFull(filename string, contents []byte, flags FileSetContentsFlags, mode int) error {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 C.gssize
	var _arg4 C.GFileSetContentsFlags // out
	var _arg5 C.int                   // out
	var _cerr *C.GError               // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg3 = (C.gssize)(len(contents))
	if len(contents) > 0 {
		_arg2 = (*C.gchar)(unsafe.Pointer(&contents[0]))
	}
	_arg4 = C.GFileSetContentsFlags(flags)
	_arg5 = C.int(mode)

	C.g_file_set_contents_full(_arg1, _arg2, _arg3, _arg4, _arg5, &_cerr)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(contents)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(mode)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// TestFile returns TRUE if any of the tests in the bitfield test are TRUE. For
// example, (G_FILE_TEST_EXISTS | G_FILE_TEST_IS_DIR) will return TRUE if the
// file exists; the check whether it's a directory doesn't matter since the
// existence test is TRUE. With the current set of available tests, there's no
// point passing in more than one test at a time.
//
// Apart from G_FILE_TEST_IS_SYMLINK all tests follow symbolic links, so for a
// symbolic link to a regular file g_file_test() will return TRUE for both
// G_FILE_TEST_IS_SYMLINK and G_FILE_TEST_IS_REGULAR.
//
// Note, that for a dangling symbolic link g_file_test() will return TRUE for
// G_FILE_TEST_IS_SYMLINK and FALSE for all other flags.
//
// You should never use g_file_test() to test whether it is safe to perform an
// operation, because there is always the possibility of the condition changing
// before you actually perform the operation. For example, you might think you
// could use G_FILE_TEST_IS_SYMLINK to know whether it is safe to write to a
// file without being tricked into writing into a different location. It doesn't
// work!
//
//    // DON'T DO THIS
//    if (!g_file_test (filename, G_FILE_TEST_IS_SYMLINK))
//      {
//        fd = g_open (filename, O_WRONLY);
//        // write to fd
//      }
//
// Another thing to note is that G_FILE_TEST_EXISTS and
// G_FILE_TEST_IS_EXECUTABLE are implemented using the access() system call.
// This usually doesn't matter, but if your program is setuid or setgid it means
// that these tests will give you the answer for the real user ID and group ID,
// rather than the effective user ID and group ID.
//
// On Windows, there are no symlinks, so testing for G_FILE_TEST_IS_SYMLINK will
// always return FALSE. Testing for G_FILE_TEST_IS_EXECUTABLE will just check
// that the file exists and its name indicates that it is executable, checking
// for well-known extensions and those listed in the PATHEXT environment
// variable.
//
// This type has been renamed from file_test.
func TestFile(filename string, test FileTest) bool {
	var _arg1 *C.gchar    // out
	var _arg2 C.GFileTest // out
	var _cret C.gboolean  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GFileTest(test)

	_cret = C.g_file_test(_arg1, _arg2)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(test)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GetCurrentDir gets the current directory.
//
// The returned string should be freed when no longer needed. The encoding of
// the returned string is system defined. On Windows, it is always UTF-8.
//
// Since GLib 2.40, this function will return the value of the "PWD" environment
// variable if it is set and it happens to be the same as the current directory.
// This can make a difference in the case that the current directory is the
// target of a symbolic link.
func GetCurrentDir() string {
	var _cret *C.gchar // in

	_cret = C.g_get_current_dir()

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// MkdirWithParents: create a directory if it doesn't already exist. Create
// intermediate parent directories as needed, too.
func MkdirWithParents(pathname string, mode int) int {
	var _arg1 *C.gchar // out
	var _arg2 C.gint   // out
	var _cret C.gint   // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pathname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint(mode)

	_cret = C.g_mkdir_with_parents(_arg1, _arg2)
	runtime.KeepAlive(pathname)
	runtime.KeepAlive(mode)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PathGetBasename gets the last component of the filename.
//
// If file_name ends with a directory separator it gets the component before the
// last slash. If file_name consists only of directory separators (and on
// Windows, possibly a drive letter), a single separator is returned. If
// file_name is empty, it gets ".".
func PathGetBasename(fileName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_path_get_basename(_arg1)
	runtime.KeepAlive(fileName)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// PathGetDirname gets the directory components of a file name. For example, the
// directory component of /usr/bin/test is /usr/bin. The directory component of
// / is /.
//
// If the file name has no directory components "." is returned. The returned
// string should be freed when no longer needed.
func PathGetDirname(fileName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_path_get_dirname(_arg1)
	runtime.KeepAlive(fileName)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _filename
}

// PathIsAbsolute returns TRUE if the given file_name is an absolute file name.
// Note that this is a somewhat vague concept on Windows.
//
// On POSIX systems, an absolute file name is well-defined. It always starts
// from the single root directory. For example "/usr/local".
//
// On Windows, the concepts of current drive and drive-specific current
// directory introduce vagueness. This function interprets as an absolute file
// name one that either begins with a directory separator such as "\Users\tml"
// or begins with the root on a drive, for example "C:\Windows". The first case
// also includes UNC paths such as "\\\\myserver\docs\foo". In all cases, either
// slashes or backslashes are accepted.
//
// Note that a file name relative to the current drive root does not truly
// specify a file uniquely over time and across processes, as the current drive
// is a per-process value and can be changed.
//
// File names relative the current directory on some specific drive, such as
// "D:foo/bar", are not interpreted as absolute by this function, but they
// obviously are not relative to the normal current directory as returned by
// getcwd() or g_get_current_dir() either. Such paths should be avoided, or need
// to be handled using Windows-specific code.
func PathIsAbsolute(fileName string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_path_is_absolute(_arg1)
	runtime.KeepAlive(fileName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PathSkipRoot returns a pointer into file_name after the root component, i.e.
// after the "/" in UNIX or "C:\" under Windows. If file_name is not an absolute
// path it returns NULL.
func PathSkipRoot(fileName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_path_skip_root(_arg1)
	runtime.KeepAlive(fileName)

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}
