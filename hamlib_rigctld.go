package rig

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type RigCtld struct {
	net.Conn
	rig hamlibRig
}

func Dial(network, address string) (Rig, error) {
	c, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}

	r := &RigCtld{Conn: c}
	return r, r.getCapabilities()
}

func (r *RigCtld) command(c string) error {
	b := new(bytes.Buffer)
	b.Write([]byte(c + "\n"))

	_, err := b.WriteTo(r.Conn)
	if err != nil {
		return err
	}

	return r.commandReply()
}

func (r *RigCtld) getCommand(c string) (string, error) {
	b := new(bytes.Buffer)
	b.Write([]byte(c + "\n"))

	_, err := b.WriteTo(r.Conn)
	if err != nil {
		return "", err
	}

	s := bufio.NewScanner(r.Conn)
	if s.Scan() {
		d := s.Text()
		if strings.HasPrefix(d, "RPRT ") {
			return "", r.parseReport(d)
		}
		return d, nil
	}

	return "", io.EOF
}

func (r *RigCtld) extendedCommand(c string) ([]string, error) {
	b := new(bytes.Buffer)
	b.Write([]byte(c + "\n"))

	_, err := b.WriteTo(r.Conn)
	if err != nil {
		return nil, err
	}

	d := make([]string, 0)
	s := bufio.NewScanner(r.Conn)
	var report string

scanner:
	for s.Scan() {
		line := strings.TrimSuffix(s.Text(), "\n")
		//fmt.Printf("<<< %q\n", line)
		if strings.HasPrefix(line, "RPRT") {
			report = line
			break scanner
		}
		d = append(d, line)
	}

	return d, r.parseReport(report)
}

func (r *RigCtld) commandReply() error {
	s := bufio.NewScanner(r.Conn)
	if s.Scan() {
		return r.parseReport(s.Text())
	}

	return ErrHamlibReport
}

func (r *RigCtld) parseReport(line string) error {
	var status int
	if _, err := fmt.Sscanf(line, "RPRT %d\n", &status); err != nil {
		return err
	}
	if status < 0 {
		status *= -1
		if status > len(ErrHamlib) {
			return fmt.Errorf("hamlib: error %d", status)
		}
		return ErrHamlib[status]
	}

	return nil
}

func (r *RigCtld) getCapabilities() error {
	c, err := r.extendedCommand("\\dump_caps")
	if err != nil {
		return err
	}

	return r.rig.parseCapabilities(c)
}

func (r *RigCtld) Model() string                 { return r.rig.modelName }
func (r *RigCtld) Manufacturer() string          { return r.rig.manufacturer }
func (r *RigCtld) String() string                { return r.Manufacturer() + " " + r.Model() }
func (r *RigCtld) Capabilities() map[string]bool { return r.rig.can }

func (r *RigCtld) Channel() (int, error) {
	return 0, nil
}

func (r *RigCtld) Frequency() (Frequency, error) {
	d, err := r.getCommand("\\get_freq")
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}
	return Frequency(f), nil
}

func (r *RigCtld) SplitFrequency() (Frequency, error) {
	d, err := r.getCommand("\\get_split_freq")
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}
	return Frequency(f), nil
}

func (r *RigCtld) RepeaterShift() (Frequency, error) {
	d, err := r.getCommand("\\get_rptr_shift")
	if err != nil {
		return 0, err
	}
	switch strings.ToLower(d) {
	case "+":
		return 1, nil
	case "-":
		return -1, nil
	}
	return 0, nil
}

func (r *RigCtld) RepeaterOffset() (Frequency, error) {
	d, err := r.getCommand("\\get_rptr_offs")
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}
	return Frequency(f), nil
}

func (r *RigCtld) TuningStep() (Frequency, error) {
	d, err := r.getCommand("\\get_ts")
	if err != nil {
		return 0, err
	}
	f, err := strconv.ParseInt(d, 10, 64)
	if err != nil {
		return 0, err
	}
	return Frequency(f), nil
}

func (r *RigCtld) Mode() (Mode, error) {
	d, err := r.getCommand("\\get_mode")
	if err != nil {
		return ModeNone, err
	}
	if m, ok := ModeValue[d]; ok {
		return m, nil
	}
	return ModeNone, fmt.Errorf("unsupported mode %q reported", d)
}

func (r *RigCtld) VFO() (VFO, error) {
	d, err := r.getCommand("\\get_vfo")
	if err != nil {
		return VFONone, err
	}
	if m, ok := VFOValue[d]; ok {
		return m, nil
	}
	return VFONone, fmt.Errorf("unsupported VFO %q reported", d)
}

func (r *RigCtld) SplitMode() (SplitMode, error) {
	d, err := r.getCommand("\\get_split_vfo")
	if err != nil {
		return SplitModeNone, err
	}
	if m, ok := SplitModeValue[d]; ok {
		return m, nil
	}
	return SplitModeNone, fmt.Errorf("unsupported split mode %q reported", d)
}

func (r *RigCtld) RIT() (int, error) {
	d, err := r.getCommand("\\get_rit")
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(d)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (r *RigCtld) XIT() (int, error) {
	d, err := r.getCommand("\\get_xit")
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(d)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func (r *RigCtld) PTT() (bool, error) {
	d, err := r.getCommand("\\get_ptt")
	if err != nil {
		return false, err
	}
	return d == "1", nil
}

func (r *RigCtld) Squelch() (bool, error) {
	d, err := r.getCommand("\\get_dcd")
	if err != nil {
		return false, err
	}
	return d == "1", nil
}

func (r *RigCtld) CTCSS() (string, error) {
	d, err := r.getCommand("\\get_ctcss_tone")
	if err != nil {
		return d, err
	}
	return d, nil
}

func (r *RigCtld) SquelchCTCSS() (string, error) {
	d, err := r.getCommand("\\get_ctcss_sql")
	if err != nil {
		return d, err
	}
	return d, nil
}

func (r *RigCtld) DCS() (string, error) {
	d, err := r.getCommand("\\get_dcs_code")
	if err != nil {
		return d, err
	}
	return d, nil
}

func (r *RigCtld) SquelchDCS() (string, error) {
	d, err := r.getCommand("\\get_dcs_sql")
	if err != nil {
		return d, err
	}
	return d, nil
}
