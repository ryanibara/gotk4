// Code generated by girgen. DO NOT EDIT.

package pango

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pango.h>
import "C"

// GravityGetForMatrix finds the gravity that best matches the rotation
// component in a `PangoMatrix`.
func GravityGetForMatrix(matrix *Matrix) Gravity {
	var arg1 *C.PangoMatrix

	arg1 = (*C.PangoMatrix)(matrix.Native())

	ret := C.pango_gravity_get_for_matrix(arg1)

	var ret0 Gravity

	ret0 = Gravity(ret)

	return ret0
}

// GravityGetForScript returns the gravity to use in laying out a `PangoItem`.
//
// The gravity is determined based on the script, base gravity, and hint.
//
// If @base_gravity is PANGO_GRAVITY_AUTO, it is first replaced with the
// preferred gravity of @script. To get the preferred gravity of a script, pass
// PANGO_GRAVITY_AUTO and PANGO_GRAVITY_HINT_STRONG in.
func GravityGetForScript(script Script, baseGravity Gravity, hint GravityHint) Gravity {
	var arg1 C.PangoScript
	var arg2 C.PangoGravity
	var arg3 C.PangoGravityHint

	arg1 = (C.PangoScript)(script)
	arg2 = (C.PangoGravity)(baseGravity)
	arg3 = (C.PangoGravityHint)(hint)

	ret := C.pango_gravity_get_for_script(arg1, arg2, arg3)

	var ret0 Gravity

	ret0 = Gravity(ret)

	return ret0
}

// GravityGetForScriptAndWidth returns the gravity to use in laying out a single
// character or `PangoItem`.
//
// The gravity is determined based on the script, East Asian width, base
// gravity, and hint,
//
// This function is similar to [type_func@Pango.Gravity.get_for_script] except
// that this function makes a distinction between narrow/half-width and
// wide/full-width characters also. Wide/full-width characters always stand
// *upright*, that is, they always take the base gravity, whereas
// narrow/full-width characters are always rotated in vertical context.
//
// If @base_gravity is PANGO_GRAVITY_AUTO, it is first replaced with the
// preferred gravity of @script.
func GravityGetForScriptAndWidth(script Script, wide bool, baseGravity Gravity, hint GravityHint) Gravity {
	var arg1 C.PangoScript
	var arg2 C.gboolean
	var arg3 C.PangoGravity
	var arg4 C.PangoGravityHint

	arg1 = (C.PangoScript)(script)
	if wide {
		arg2 = C.TRUE
	}
	arg3 = (C.PangoGravity)(baseGravity)
	arg4 = (C.PangoGravityHint)(hint)

	ret := C.pango_gravity_get_for_script_and_width(arg1, arg2, arg3, arg4)

	var ret0 Gravity

	ret0 = Gravity(ret)

	return ret0
}

// GravityToRotation converts a Gravity value to its natural rotation in
// radians.
//
// Note that [method@Pango.Matrix.rotate] takes angle in degrees, not radians.
// So, to call [method@Pango.Matrix,rotate] with the output of this function you
// should multiply it by (180. / G_PI).
func GravityToRotation(gravity Gravity) float64 {
	var arg1 C.PangoGravity

	arg1 = (C.PangoGravity)(gravity)

	ret := C.pango_gravity_to_rotation(arg1)

	var ret0 float64

	ret0 = float64(ret)

	return ret0
}
