package main

import (
	"fmt"
	"sort"

	"github.com/pd0mz/go-rig"
)

func main() {
	r, err := rig.Dial("tcp", "localhost:4532")
	if err != nil {
		panic(err)
	}

	var (
		freq  rig.Frequency
		hz    int
		mode  rig.Mode
		split rig.SplitMode
		flag  bool
		val   string
		vfo   rig.VFO
	)

	fmt.Printf("rig......: %s\n", r)
	if freq, err = r.Frequency(); err != nil {
		panic(err)
	}
	fmt.Printf("frequency: %s\n", freq)
	if freq, err = r.SplitFrequency(); err != nil {
		fmt.Printf("split....: error: %v\n", err)
	} else {
		fmt.Printf("split....: %s\n", freq)
	}
	if freq, err = r.RepeaterOffset(); err != nil {
		fmt.Printf("offset...: error: %v\n", err)
	} else {
		fmt.Printf("offset...: %s\n", freq)
	}
	if freq, err = r.TuningStep(); err != nil {
		fmt.Printf("step.....: error: %v\n", err)
	} else {
		fmt.Printf("step.....: %s\n", freq)
	}
	if mode, err = r.Mode(); err != nil {
		panic(err)
	}
	fmt.Printf("mode.....: %s\n", mode)
	if vfo, err = r.VFO(); err != nil {
		fmt.Printf("vfo......: error: %v\n", err)
	} else {
		fmt.Printf("vfo......: %s\n", vfo)
	}
	if split, err = r.SplitMode(); err != nil {
		fmt.Printf("splitmode: error: %v\n", err)
	} else {
		fmt.Printf("splitmode: %s\n", split)
	}
	if hz, err = r.RIT(); err != nil {
		fmt.Printf("rit......: error: %v\n", err)
	} else {
		fmt.Printf("rit......: %d Hz\n", hz)
	}
	if hz, err = r.XIT(); err != nil {
		fmt.Printf("xit......: error: %v\n", err)
	} else {
		fmt.Printf("xit......: %d Hz\n", hz)
	}
	if flag, err = r.PTT(); err != nil {
		fmt.Printf("ptt......: error: %v\n", err)
	} else {
		fmt.Printf("ptt......: %v\n", flag)
	}
	if flag, err = r.Squelch(); err != nil {
		fmt.Printf("squelch..: error: %v\n", err)
	} else {
		fmt.Printf("squelch..: %v\n", flag)
	}
	if val, err = r.CTCSS(); err != nil {
		fmt.Printf("ctcss tx.: error: %v\n", err)
	} else {
		fmt.Printf("ctcss tx.: %v\n", val)
	}
	if val, err = r.SquelchCTCSS(); err != nil {
		fmt.Printf("ctcss rx.: error: %v\n", err)
	} else {
		fmt.Printf("ctcss rx.: %v\n", val)
	}
	if val, err = r.DCS(); err != nil {
		fmt.Printf("dcs rx...: error: %v\n", err)
	} else {
		fmt.Printf("dcs rx...: %v\n", val)
	}
	if val, err = r.SquelchDCS(); err != nil {
		fmt.Printf("dcs tx...: error: %v\n", err)
	} else {
		fmt.Printf("dcs tx...: %v\n", val)
	}

	var (
		can []string
		cap = r.Capabilities()
	)
	for k := range cap {
		can = append(can, k)
	}
	sort.Sort(sort.StringSlice(can))
	for _, k := range can {
		fmt.Printf("can %s: %v\n", k, cap[k])
	}
}
