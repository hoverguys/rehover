package main

import (
	"flag"
	"fmt"
	"image"
	"os"
	"path/filepath"

	// Image formats
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [flags] <file1> [<file2> ...]\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nSupported input formats:\n")
		for _, format := range []string{"JPEG", "GIF", "PNG"} {
			fmt.Fprintf(os.Stderr, "    %s\n", format)
		}
	}
	outpath := flag.String("o", "-", "Output file (- for STDOUT)")
	noheader := flag.Bool("noheader", false, "Don't write header, output raw PNG")
	maxW := flag.Int("maxwidth", 1<<16, "Max width of output texture")
	maxH := flag.Int("maxheight", 1<<16, "Max height of output texture")
	cwd, err := os.Getwd()
	checkErr(err, "Couldn't get working directory")
	stripPfx := flag.String("prefix", cwd, "Prefix to strip from file paths for hashing")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "[FAIL] No input files were specified\n")
		os.Exit(1)
	}

	// All positional arguments are input files
	inputFiles := flag.Args()

	maxBounds := image.Point{*maxW, *maxH}
	absoutpath, err := filepath.Abs(*outpath)
	checkErr(err, "Error converting outpath %s to absolute", *outpath)
	packer := NewTexPacker(absoutpath, TexPackerOptions{
		MaxBounds:   maxBounds,
		WriteHeader: !*noheader,
		StripPrefix: *stripPfx,
	})
	// Read images from input
	for _, path := range inputFiles {
		// Convert file path to absolute
		abspath, err := filepath.Abs(path)
		checkErr(err, "Error converting file path %s to absolute", path)
		packer.Add(abspath)
	}
	checkErr(packer.Save(), "Failed to save packed texture")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
