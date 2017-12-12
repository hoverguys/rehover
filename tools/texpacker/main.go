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
		fmt.Fprintf(os.Stderr, "Usage: %s [-o outfile] <file1> [<file2> ...]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nSupported input formats:\n")
		for _, format := range []string{"JPEG", "GIF", "PNG"} {
			fmt.Fprintf(os.Stderr, "    %s\n", format)
		}
	}
	outpath := flag.String("o", "-", "Output file (- for STDOUT)")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "[FAIL] No input files were specified\n")
		os.Exit(1)
	}

	// All positional arguments are input files
	inputFiles := flag.Args()

	maxBounds := image.Point{1 << 16, 1 << 16} // TODO
	packer := NewTexPacker(*outpath, TexPackerOptions{
		MaxBounds: maxBounds,
	})
	// Read images from input
	for _, path := range inputFiles {
		packer.Add(path)
	}
	if err := packer.Save(); err != nil {
		panic(err)
	}
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
