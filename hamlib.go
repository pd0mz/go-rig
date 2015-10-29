package rig

import (
	"errors"
	"strings"
)

var (
	ErrHamlibReport = errors.New("did not receive a report from hamlib")
	ErrHamlib       = []error{
		nil,
		errors.New("hamlib: invalid parameter"),
		errors.New("hamlib: invalid configuration"),
		errors.New("hamlib: memory shortage"),
		errors.New("hamlib: function not implemented"),
		errors.New("hamlib: communication timed out"),
		errors.New("hamlib: I/O error"),
		errors.New("hamlib: internal error"),
		errors.New("hamlib: protocol error"),
		errors.New("hamlib: command rejected by rig"),
		errors.New("hamlib: command performed but truncated"),
		errors.New("hamlib: function not available"),
		errors.New("hamlib: VFO not targetable"),
		errors.New("hamlib: bus error"),
		errors.New("hamlib: bus busy"),
		errors.New("hamlib: invalid pointer"),
		errors.New("hamlib: invalid VFO"),
		errors.New("hamlib: argument out of domain"),
	}
)

type hamlibRig struct {
	modelName      string
	manufacturer   string
	backendVersion string
	rigType        string
	ctcss          []string
	dcs            []string
	can            map[string]bool
}

func (r *hamlibRig) parseCapabilities(c []string) error {
	if r.can == nil {
		r.can = make(map[string]bool)
	}

	for _, line := range c {
		part := strings.SplitN(line, ":", 2)
		if len(part) != 2 {
			continue
		}

		key := strings.ToLower(part[0])
		switch {
		case key == "model name":
			r.modelName = strings.TrimSpace(part[1])
		case key == "mfg name":
			r.manufacturer = strings.TrimSpace(part[1])
		case key == "backend version":
			r.backendVersion = strings.TrimSpace(part[1])
		case strings.HasPrefix(key, "can "):
			r.can[key[4:]] = part[1][len(part[1])-1] == 'Y'
		}
	}

	return nil
}
