// Code generated by "stringer -type=MMModem3gppNetworkAvailability -trimprefix=MmModem3gppNetworkAvailability"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModem3gppNetworkAvailabilityUnknown-0]
	_ = x[MmModem3gppNetworkAvailabilityAvailable-1]
	_ = x[MmModem3gppNetworkAvailabilityCurrent-2]
	_ = x[MmModem3gppNetworkAvailabilityForbidden-3]
}

const _MMModem3gppNetworkAvailability_name = "UnknownAvailableCurrentForbidden"

var _MMModem3gppNetworkAvailability_index = [...]uint8{0, 7, 16, 23, 32}

func (i MMModem3gppNetworkAvailability) String() string {
	if i >= MMModem3gppNetworkAvailability(len(_MMModem3gppNetworkAvailability_index)-1) {
		return "MMModem3gppNetworkAvailability(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MMModem3gppNetworkAvailability_name[_MMModem3gppNetworkAvailability_index[i]:_MMModem3gppNetworkAvailability_index[i+1]]
}
