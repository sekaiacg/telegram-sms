// Code generated by "stringer -type=MMModemBand -trimprefix=MmModemBand"; DO NOT EDIT.

package modemmanager

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MmModemBandUnknown-0]
	_ = x[MmModemBandEgsm-1]
	_ = x[MmModemBandDcs-2]
	_ = x[MmModemBandPcs-3]
	_ = x[MmModemBandG850-4]
	_ = x[MmModemBandUtran1-5]
	_ = x[MmModemBandUtran3-6]
	_ = x[MmModemBandUtran4-7]
	_ = x[MmModemBandUtran6-8]
	_ = x[MmModemBandUtran5-9]
	_ = x[MmModemBandUtran8-10]
	_ = x[MmModemBandUtran9-11]
	_ = x[MmModemBandUtran2-12]
	_ = x[MmModemBandUtran7-13]
	_ = x[MmModemBandG450-14]
	_ = x[MmModemBandG480-15]
	_ = x[MmModemBandG750-16]
	_ = x[MmModemBandG380-17]
	_ = x[MmModemBandG410-18]
	_ = x[MmModemBandG710-19]
	_ = x[MmModemBandG810-20]
	_ = x[MmModemBandEutran1-31]
	_ = x[MmModemBandEutran2-32]
	_ = x[MmModemBandEutran3-33]
	_ = x[MmModemBandEutran4-34]
	_ = x[MmModemBandEutran5-35]
	_ = x[MmModemBandEutran6-36]
	_ = x[MmModemBandEutran7-37]
	_ = x[MmModemBandEutran8-38]
	_ = x[MmModemBandEutran9-39]
	_ = x[MmModemBandEutran10-40]
	_ = x[MmModemBandEutran11-41]
	_ = x[MmModemBandEutran12-42]
	_ = x[MmModemBandEutran13-43]
	_ = x[MmModemBandEutran14-44]
	_ = x[MmModemBandEutran17-47]
	_ = x[MmModemBandEutran18-48]
	_ = x[MmModemBandEutran19-49]
	_ = x[MmModemBandEutran20-50]
	_ = x[MmModemBandEutran21-51]
	_ = x[MmModemBandEutran22-52]
	_ = x[MmModemBandEutran23-53]
	_ = x[MmModemBandEutran24-54]
	_ = x[MmModemBandEutran25-55]
	_ = x[MmModemBandEutran26-56]
	_ = x[MmModemBandEutran27-57]
	_ = x[MmModemBandEutran28-58]
	_ = x[MmModemBandEutran29-59]
	_ = x[MmModemBandEutran30-60]
	_ = x[MmModemBandEutran31-61]
	_ = x[MmModemBandEutran32-62]
	_ = x[MmModemBandEutran33-63]
	_ = x[MmModemBandEutran34-64]
	_ = x[MmModemBandEutran35-65]
	_ = x[MmModemBandEutran36-66]
	_ = x[MmModemBandEutran37-67]
	_ = x[MmModemBandEutran38-68]
	_ = x[MmModemBandEutran39-69]
	_ = x[MmModemBandEutran40-70]
	_ = x[MmModemBandEutran41-71]
	_ = x[MmModemBandEutran42-72]
	_ = x[MmModemBandEutran43-73]
	_ = x[MmModemBandEutran44-74]
	_ = x[MmModemBandEutran45-75]
	_ = x[MmModemBandEutran46-76]
	_ = x[MmModemBandEutran47-77]
	_ = x[MmModemBandEutran48-78]
	_ = x[MmModemBandEutran49-79]
	_ = x[MmModemBandEutran50-80]
	_ = x[MmModemBandEutran51-81]
	_ = x[MmModemBandEutran52-82]
	_ = x[MmModemBandEutran53-83]
	_ = x[MmModemBandEutran54-84]
	_ = x[MmModemBandEutran55-85]
	_ = x[MmModemBandEutran56-86]
	_ = x[MmModemBandEutran57-87]
	_ = x[MmModemBandEutran58-88]
	_ = x[MmModemBandEutran59-89]
	_ = x[MmModemBandEutran60-90]
	_ = x[MmModemBandEutran61-91]
	_ = x[MmModemBandEutran62-92]
	_ = x[MmModemBandEutran63-93]
	_ = x[MmModemBandEutran64-94]
	_ = x[MmModemBandEutran65-95]
	_ = x[MmModemBandEutran66-96]
	_ = x[MmModemBandEutran67-97]
	_ = x[MmModemBandEutran68-98]
	_ = x[MmModemBandEutran69-99]
	_ = x[MmModemBandEutran70-100]
	_ = x[MmModemBandEutran71-101]
	_ = x[MmModemBandCdmaBc0-128]
	_ = x[MmModemBandCdmaBc1-129]
	_ = x[MmModemBandCdmaBc2-130]
	_ = x[MmModemBandCdmaBc3-131]
	_ = x[MmModemBandCdmaBc4-132]
	_ = x[MmModemBandCdmaBc5-134]
	_ = x[MmModemBandCdmaBc6-135]
	_ = x[MmModemBandCdmaBc7-136]
	_ = x[MmModemBandCdmaBc8-137]
	_ = x[MmModemBandCdmaBc9-138]
	_ = x[MmModemBandCdmaBc10-139]
	_ = x[MmModemBandCdmaBc11-140]
	_ = x[MmModemBandCdmaBc12-141]
	_ = x[MmModemBandCdmaBc13-142]
	_ = x[MmModemBandCdmaBc14-143]
	_ = x[MmModemBandCdmaBc15-144]
	_ = x[MmModemBandCdmaBc16-145]
	_ = x[MmModemBandCdmaBc17-146]
	_ = x[MmModemBandCdmaBc18-147]
	_ = x[MmModemBandCdmaBc19-148]
	_ = x[MmModemBandUtran10-210]
	_ = x[MmModemBandUtran11-211]
	_ = x[MmModemBandUtran12-212]
	_ = x[MmModemBandUtran13-213]
	_ = x[MmModemBandUtran14-214]
	_ = x[MmModemBandUtran19-219]
	_ = x[MmModemBandUtran20-220]
	_ = x[MmModemBandUtran21-221]
	_ = x[MmModemBandUtran22-222]
	_ = x[MmModemBandUtran25-225]
	_ = x[MmModemBandUtran26-226]
	_ = x[MmModemBandUtran32-232]
	_ = x[MmModemBandAny-256]
}

const (
	_MMModemBand_name_0 = "UnknownEgsmDcsPcsG850Utran1Utran3Utran4Utran6Utran5Utran8Utran9Utran2Utran7G450G480G750G380G410G710G810"
	_MMModemBand_name_1 = "Eutran1Eutran2Eutran3Eutran4Eutran5Eutran6Eutran7Eutran8Eutran9Eutran10Eutran11Eutran12Eutran13Eutran14"
	_MMModemBand_name_2 = "Eutran17Eutran18Eutran19Eutran20Eutran21Eutran22Eutran23Eutran24Eutran25Eutran26Eutran27Eutran28Eutran29Eutran30Eutran31Eutran32Eutran33Eutran34Eutran35Eutran36Eutran37Eutran38Eutran39Eutran40Eutran41Eutran42Eutran43Eutran44Eutran45Eutran46Eutran47Eutran48Eutran49Eutran50Eutran51Eutran52Eutran53Eutran54Eutran55Eutran56Eutran57Eutran58Eutran59Eutran60Eutran61Eutran62Eutran63Eutran64Eutran65Eutran66Eutran67Eutran68Eutran69Eutran70Eutran71"
	_MMModemBand_name_3 = "CdmaBc0CdmaBc1CdmaBc2CdmaBc3CdmaBc4"
	_MMModemBand_name_4 = "CdmaBc5CdmaBc6CdmaBc7CdmaBc8CdmaBc9CdmaBc10CdmaBc11CdmaBc12CdmaBc13CdmaBc14CdmaBc15CdmaBc16CdmaBc17CdmaBc18CdmaBc19"
	_MMModemBand_name_5 = "Utran10Utran11Utran12Utran13Utran14"
	_MMModemBand_name_6 = "Utran19Utran20Utran21Utran22"
	_MMModemBand_name_7 = "Utran25Utran26"
	_MMModemBand_name_8 = "Utran32"
	_MMModemBand_name_9 = "Any"
)

var (
	_MMModemBand_index_0 = [...]uint8{0, 7, 11, 14, 17, 21, 27, 33, 39, 45, 51, 57, 63, 69, 75, 79, 83, 87, 91, 95, 99, 103}
	_MMModemBand_index_1 = [...]uint8{0, 7, 14, 21, 28, 35, 42, 49, 56, 63, 71, 79, 87, 95, 103}
	_MMModemBand_index_2 = [...]uint16{0, 8, 16, 24, 32, 40, 48, 56, 64, 72, 80, 88, 96, 104, 112, 120, 128, 136, 144, 152, 160, 168, 176, 184, 192, 200, 208, 216, 224, 232, 240, 248, 256, 264, 272, 280, 288, 296, 304, 312, 320, 328, 336, 344, 352, 360, 368, 376, 384, 392, 400, 408, 416, 424, 432, 440}
	_MMModemBand_index_3 = [...]uint8{0, 7, 14, 21, 28, 35}
	_MMModemBand_index_4 = [...]uint8{0, 7, 14, 21, 28, 35, 43, 51, 59, 67, 75, 83, 91, 99, 107, 115}
	_MMModemBand_index_5 = [...]uint8{0, 7, 14, 21, 28, 35}
	_MMModemBand_index_6 = [...]uint8{0, 7, 14, 21, 28}
	_MMModemBand_index_7 = [...]uint8{0, 7, 14}
)

func (i MMModemBand) String() string {
	switch {
	case i <= 20:
		return _MMModemBand_name_0[_MMModemBand_index_0[i]:_MMModemBand_index_0[i+1]]
	case 31 <= i && i <= 44:
		i -= 31
		return _MMModemBand_name_1[_MMModemBand_index_1[i]:_MMModemBand_index_1[i+1]]
	case 47 <= i && i <= 101:
		i -= 47
		return _MMModemBand_name_2[_MMModemBand_index_2[i]:_MMModemBand_index_2[i+1]]
	case 128 <= i && i <= 132:
		i -= 128
		return _MMModemBand_name_3[_MMModemBand_index_3[i]:_MMModemBand_index_3[i+1]]
	case 134 <= i && i <= 148:
		i -= 134
		return _MMModemBand_name_4[_MMModemBand_index_4[i]:_MMModemBand_index_4[i+1]]
	case 210 <= i && i <= 214:
		i -= 210
		return _MMModemBand_name_5[_MMModemBand_index_5[i]:_MMModemBand_index_5[i+1]]
	case 219 <= i && i <= 222:
		i -= 219
		return _MMModemBand_name_6[_MMModemBand_index_6[i]:_MMModemBand_index_6[i+1]]
	case 225 <= i && i <= 226:
		i -= 225
		return _MMModemBand_name_7[_MMModemBand_index_7[i]:_MMModemBand_index_7[i+1]]
	case i == 232:
		return _MMModemBand_name_8
	case i == 256:
		return _MMModemBand_name_9
	default:
		return "MMModemBand(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
