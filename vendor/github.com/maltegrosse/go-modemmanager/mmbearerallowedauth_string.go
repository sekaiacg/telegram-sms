// Code generated by "stringer -type=MMBearerAllowedAuth -trimprefix=MmBearerAllowedAuth"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmBearerAllowedAuthUnknown-0]
	_ = x[MmBearerAllowedAuthNone-1]
	_ = x[MmBearerAllowedAuthPap-2]
	_ = x[MmBearerAllowedAuthChap-4]
	_ = x[MmBearerAllowedAuthMschap-8]
	_ = x[MmBearerAllowedAuthMschapv2-16]
	_ = x[MmBearerAllowedAuthEap-32]
}

const (
	_MMBearerAllowedAuth_name_0 = "UnknownNonePap"
	_MMBearerAllowedAuth_name_1 = "Chap"
	_MMBearerAllowedAuth_name_2 = "Mschap"
	_MMBearerAllowedAuth_name_3 = "Mschapv2"
	_MMBearerAllowedAuth_name_4 = "Eap"
)

var (
	_MMBearerAllowedAuth_index_0 = [...]uint8{0, 7, 11, 14}
)

func (i MMBearerAllowedAuth) String() string {
	switch {
	case i <= 2:
		return _MMBearerAllowedAuth_name_0[_MMBearerAllowedAuth_index_0[i]:_MMBearerAllowedAuth_index_0[i+1]]
	case i == 4:
		return _MMBearerAllowedAuth_name_1
	case i == 8:
		return _MMBearerAllowedAuth_name_2
	case i == 16:
		return _MMBearerAllowedAuth_name_3
	case i == 32:
		return _MMBearerAllowedAuth_name_4
	default:
		return "MMBearerAllowedAuth(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
