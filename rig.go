package rig

import (
	"fmt"
	"strconv"
	"strings"
)

type Frequency int64

const (
	Hz  Frequency = 1
	KHz Frequency = 10e2
	MHz Frequency = 10e5
	GHz Frequency = 10e8
	THz Frequency = 10e11
	PHz Frequency = 10e14
)

func (f Frequency) String() string {
	var v, s string
	switch {
	case f < KHz:
		v = strconv.Itoa(int(f))
		s = "Hz"
	case f < MHz:
		v = fmt.Sprintf("%f", float32(f)/float32(KHz))
		s = "kHz"
	case f < GHz:
		v = strconv.FormatFloat(float64(f)/float64(MHz), 'f', -1, 64)
		s = "MHz"
	case f < THz:
		v = strconv.FormatFloat(float64(f)/float64(GHz), 'f', -1, 64)
		s = "MHz"
	case f < PHz:
		v = strconv.FormatFloat(float64(f)/float64(THz), 'f', -1, 64)
		s = "THz"
	default:
		return "overflow"
	}

	v = strings.TrimSuffix(v, "0")
	return v + " " + s
}

type Rig interface {
	String() string
	Model() string
	Manufacturer() string
	Capabilities() map[string]bool

	Channel() (int, error)
	Frequency() (Frequency, error)
	SplitFrequency() (Frequency, error)
	RepeaterShift() (Frequency, error)
	RepeaterOffset() (Frequency, error)
	TuningStep() (Frequency, error)
	Mode() (Mode, error)
	VFO() (VFO, error)
	SplitMode() (SplitMode, error)

	RIT() (int, error)
	XIT() (int, error)
	PTT() (bool, error)
	Squelch() (bool, error)
	CTCSS() (string, error)
	SquelchCTCSS() (string, error)
	DCS() (string, error)
	SquelchDCS() (string, error)
}
