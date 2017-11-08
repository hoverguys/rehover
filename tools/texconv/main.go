package main

import (
	"flag"
	"fmt"
	"image"
	"os"

	// Image formats
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nSupported color formats:\n")
		for format, desc := range allfmt {
			fmt.Fprintf(os.Stderr, "    %-8s %s\n", format, desc)
		}
	}
	path := flag.String("in", "-", "Image file to convert (- for STDIN)")
	outpath := flag.String("out", "-", "Output file (- for STDOUT)")
	endianess := flag.String("endianess", "big", "Endianess of values (valid values: big, small)")
	format := flag.String("fmt", "RGBA8", "Output color format (see below for full list)")
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

	// Read image from input
	img, srcformat, err := image.Decode(in)
	checkErr(err, "Could not decode input image")

	fmt.Fprintf(os.Stderr, "Image detected as: %s\n", srcformat)

	checkErr(SaveTexture(img, out, TextureOptions{
		Endianess: Endianess(*endianess),
		Format:    ColorFmt(*format),
	}), "Error while saving output texture")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
