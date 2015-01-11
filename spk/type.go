package spk

//go:generate stringer -type=Type

// Type is the SPK data type.
// See "Supported Data Types" in "SPK required reading" for details.
type Type int32

// TODO remove, just use interfaces.
const (
	ModDiffArrays       Type = 1
	TypChebyshevPosOnly Type = 2
	TypChebyshev        Type = 3
	Discrete            Type = 5
	LagrangeEq          Type = 8
	LagrangeUneq        Type = 9
	SpaceCommTwoLine    Type = 10
	HermiteUniform      Type = 12
	HermiteNonUniform   Type = 13
	ChebyshevNonUniform Type = 14
	ConicProp           Type = 15
	EquinoctialElements Type = 17
	CebyshevVelOnly     Type = 20
	ExtModDiffArrays    Type = 21
)
