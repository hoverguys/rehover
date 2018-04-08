package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	path := flag.String("in", "-", "OBJ file to convert (- for STDIN)")
	outpath := flag.String("out", "-", "Output file (- for STDOUT)")
	allowngons := flag.Bool("allowngons", false, "Allow NGons (not currently supported)")
	allowpartial := flag.Bool("allowpartialfaces", false, "Allow faces without texcoord or normal indices (not recomended)")
	allowmultiple := flag.Bool("allowmultiple", false, "Allow multiple meshes in one file (not currently supported)")
	dumpjson := flag.Bool("dumpjson", false, "Dump JSON from parser instead of binary (for debugging parser bugs)(")
	endianess := flag.String("endianess", "big", "Endianess of values (valid values: big, small)")
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
	meshdata, err := ParseOBJ(in, OBJSettings{
		AllowNgons:      *allowngons,
		PartialFaces:    *allowpartial,
		MultipleObjects: *allowmultiple,
	})
	checkErr(err, "Error while parsing input file")

	// Preprocess data
	Preprocess(&meshdata)

	// Dump to JSON is requested
	if *dumpjson {
		enc := json.NewEncoder(out)
		enc.SetIndent("", "  ")
		checkErr(enc.Encode(meshdata), "Error while encoding to JSON")
		return
	}

	// Save to binary
	checkErr(SaveModel(meshdata, out, Endianess(*endianess)), "Error saving model")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
