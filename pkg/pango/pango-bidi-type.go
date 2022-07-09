// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/girepository"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gobject-2.0
// #include <stdlib.h>
// #include <glib.h>
// #include <glib-object.h>
import "C"

// GTypeBidiType returns the GType for the type BidiType.
//
// This function has the side effect of registering a GValue marshaler
// globally. Use this if you need that for any reason. The function is
// concurrently safe to use.
func GTypeBidiType() coreglib.Type {
	gtype := coreglib.Type(girepository.MustFind("Pango", "BidiType").RegisteredGType())
	coreglib.RegisterGValueMarshaler(gtype, marshalBidiType)
	return gtype
}

// BidiType: PangoBidiType represents the bidirectional character type of a
// Unicode character as specified by the <ulink
// url="http://www.unicode.org/reports/tr9/">Unicode bidirectional
// algorithm</ulink>.
//
// Deprecated: Use fribidi for this information.
type BidiType C.gint

const (
	// BidiTypeL: left-to-Right.
	BidiTypeL BidiType = iota
	// BidiTypeLre: left-to-Right Embedding.
	BidiTypeLre
	// BidiTypeLro: left-to-Right Override.
	BidiTypeLro
	// BidiTypeR: right-to-Left.
	BidiTypeR
	// BidiTypeAl: right-to-Left Arabic.
	BidiTypeAl
	// BidiTypeRLE: right-to-Left Embedding.
	BidiTypeRLE
	// BidiTypeRlo: right-to-Left Override.
	BidiTypeRlo
	// BidiTypePDF: pop Directional Format.
	BidiTypePDF
	// BidiTypeEn: european Number.
	BidiTypeEn
	// BidiTypeES: european Number Separator.
	BidiTypeES
	// BidiTypeEt: european Number Terminator.
	BidiTypeEt
	// BidiTypeAn: arabic Number.
	BidiTypeAn
	// BidiTypeCs: common Number Separator.
	BidiTypeCs
	// BidiTypeNsm: nonspacing Mark.
	BidiTypeNsm
	// BidiTypeBn: boundary Neutral.
	BidiTypeBn
	// BidiTypeB: paragraph Separator.
	BidiTypeB
	// BidiTypeS: segment Separator.
	BidiTypeS
	// BidiTypeWs: whitespace.
	BidiTypeWs
	// BidiTypeOn: other Neutrals.
	BidiTypeOn
)

func marshalBidiType(p uintptr) (interface{}, error) {
	return BidiType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for BidiType.
func (b BidiType) String() string {
	switch b {
	case BidiTypeL:
		return "L"
	case BidiTypeLre:
		return "Lre"
	case BidiTypeLro:
		return "Lro"
	case BidiTypeR:
		return "R"
	case BidiTypeAl:
		return "Al"
	case BidiTypeRLE:
		return "RLE"
	case BidiTypeRlo:
		return "Rlo"
	case BidiTypePDF:
		return "PDF"
	case BidiTypeEn:
		return "En"
	case BidiTypeES:
		return "ES"
	case BidiTypeEt:
		return "Et"
	case BidiTypeAn:
		return "An"
	case BidiTypeCs:
		return "Cs"
	case BidiTypeNsm:
		return "Nsm"
	case BidiTypeBn:
		return "Bn"
	case BidiTypeB:
		return "B"
	case BidiTypeS:
		return "S"
	case BidiTypeWs:
		return "Ws"
	case BidiTypeOn:
		return "On"
	default:
		return fmt.Sprintf("BidiType(%d)", b)
	}
}

// GetMirrorChar returns the mirrored character of a Unicode character.
//
// Mirror characters are determined by the Unicode mirrored property.
//
// Use g_unichar_get_mirror_char() instead; the docs for that function provide
// full details.
//
// The function takes the following parameters:
//
//    - ch: unicode character.
//    - mirroredCh: location to store the mirrored character.
//
// The function returns the following values:
//
//    - ok: TRUE if ch has a mirrored character and mirrored_ch is filled in,
//      FALSE otherwise.
//
func GetMirrorChar(ch uint32, mirroredCh *uint32) bool {
	var _args [2]girepository.Argument

	*(*C.gunichar)(unsafe.Pointer(&_args[0])) = C.gunichar(ch)
	*(**C.void)(unsafe.Pointer(&_args[1])) = (*C.void)(unsafe.Pointer(mirroredCh))

	_gret := girepository.MustFind("Pango", "get_mirror_char").Invoke(_args[:], nil)
	_cret = *(*C.gboolean)(unsafe.Pointer(&_gret))

	runtime.KeepAlive(ch)
	runtime.KeepAlive(mirroredCh)

	var _ok bool // out

	if *(*C.gboolean)(unsafe.Pointer(&_cret)) != 0 {
		_ok = true
	}

	return _ok
}
