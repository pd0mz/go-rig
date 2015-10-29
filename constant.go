package rig

import "strings"

type Mode uint32

const (
	ModeNone    Mode = 0
	ModeAM      Mode = 1 << iota // Amplitude Modulation
	ModeCW                       // Continuous Wave
	ModeUSB                      // Upper Side Band
	ModeLSB                      // Lower Side Band
	ModeRTTY                     // Radio Teletype
	ModeFM                       // Narrow band FM
	ModeWFM                      // Broadcast wide band FM
	ModeCWR                      // Continuous Wave Reverse
	ModeRTTYR                    // Radio Teletype Reverse
	ModeAMS                      // Amplitude Modulation Synchronous
	ModePKTLSB                   // Packet/Digital LSB
	ModePKTUSB                   // Packet/Digital USB
	ModePKTFM                    // Packet/Digital FM
	ModeECSSUSB                  // Exalted Carrier Single Side Band USB
	ModeECSSLSB                  // Exalted Carrier Single Side Band LSB
	ModeFAX                      // Facsimile
	ModeSAM                      // Synchronous AM double side band
	ModeSAL                      // Synchronous AM lower side band
	ModeSAH                      // Synchronous AM upper side band
	ModeDSB                      // Double side band suppressed carrier
)

var (
	ModeName = map[Mode]string{
		ModeNone:    "none",
		ModeAM:      "AM",
		ModeCW:      "CW",
		ModeUSB:     "USB",
		ModeLSB:     "LSB",
		ModeRTTY:    "RTTY",
		ModeFM:      "FM",
		ModeWFM:     "WFM",
		ModeCWR:     "CWR",
		ModeRTTYR:   "RTTYR",
		ModeAMS:     "AMS",
		ModePKTLSB:  "PKTLSB",
		ModePKTUSB:  "PKTUSB",
		ModePKTFM:   "PKTFM",
		ModeECSSUSB: "ECSSUSB",
		ModeECSSLSB: "ECSSLSB",
		ModeFAX:     "FAX",
		ModeSAM:     "SAM",
		ModeSAL:     "SAL",
		ModeSAH:     "SAH",
		ModeDSB:     "DSB",
	}
	ModeValue = map[string]Mode{}
)

func (m Mode) String() string {
	if s, ok := ModeName[m]; ok {
		return s
	}
	return "unknown"
}

type ITURegion uint8

const (
	ITURegion1 ITURegion = iota + 1
	ITURegion2
	ITURegion3
)

type VFO uint32

const (
	VFONone VFO = 0
	VFOA    VFO = 1 << iota
	VFOB
	VFOC
	VFOSub     VFO = 1 << 25
	VFOMain    VFO = 1 << 26
	VFOLast    VFO = 1 << 27
	VFOMemory  VFO = 1 << 28
	VFOCurrent VFO = 1 << 29
)

func (v VFO) String() string {
	if s, ok := VFOName[v]; ok {
		return s
	}
	return "unknown"
}

var (
	VFOName = map[VFO]string{
		VFONone:    "none",
		VFOA:       "VFO A",
		VFOB:       "VFO B",
		VFOC:       "VFO C",
		VFOSub:     "sub",
		VFOMain:    "main",
		VFOMemory:  "memory",
		VFOCurrent: "current",
	}
	VFOValue = map[string]VFO{}
)

type SplitMode uint8

const (
	SplitModeNone SplitMode = iota
	SplitModeOff
	SplitModeOn
	SplitModeTX
)

func (m SplitMode) String() string {
	if s, ok := SplitModeName[m]; ok {
		return s
	}
	return "unknown"
}

var (
	SplitModeName = map[SplitMode]string{
		SplitModeNone: "none",
		SplitModeOff:  "off",
		SplitModeOn:   "on",
		SplitModeTX:   "tx",
	}
	SplitModeValue = map[string]SplitMode{}
)

func init() {
	for m, n := range ModeName {
		ModeValue[n] = m
	}
	for v, n := range VFOName {
		VFOValue[strings.Replace(n, " ", "", -1)] = v
	}
	for m, n := range SplitModeName {
		SplitModeValue[n] = m
	}
}
