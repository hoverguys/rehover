package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
)

var debugmsg *bool
var recording = false

func main() {
	inpath := flag.String("in", "-", "Input shader file (- for stdin)")
	outpath := flag.String("out", "-", "Output file (- for stdout)")
	debugmsg = flag.Bool("debug", false, "Enable debug messages on stderr")
	dump := flag.Bool("dump", false, "Dump a textual version of the parsed shader instead of generating the TDL")
	flag.Parse()

	// Get input reader
	in := os.Stdin
	if *inpath != "-" {
		file, err := os.Open(*inpath)
		checkErr(err, "Cannot open input file")
		defer file.Close()
		in = file
	}

	// Parse input file
	functions, err := Parse(in)
	checkErr(err, "Error while parsing shader")

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

	if *dump {
		for _, fn := range functions {
			fmt.Fprintln(out, fn)
		}
	}

	GX_Init()

	// Rehover tweaks
	GX_SetCullMode(GX_CULL_FRONT)

	GX_Flush()

	fifo := NewFifo(out, binary.BigEndian)
	wgPipe = fifo
	recording = true

	for _, fn := range functions {
		checkErr(fn.Call(), "Error while calling \"%s\"", fn)
	}

	GX_Flush()
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
