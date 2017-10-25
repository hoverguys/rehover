package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("in", "-", "OBJ file to convert (- for STDIN)")
	outpath := flag.String("out", "-", "Output file (- for STDOUT)")
	flag.Parse()

	// Get input reader
	in := os.Stdin
	if *path != "-" {
		file, err := os.Open(*path)
		checkErr(err, "Cannot open input file")
		defer file.Close()
		in = file
	}

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

	// Read input and get mesh data
	meshdata, err := ParseOBJ(in)
	checkErr(err, "Error while parsing input file")

	//TODO Save
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":", args)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
