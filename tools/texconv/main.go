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

var quiet *bool

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
		fmt.Fprint(os.Stderr, "\nSupported color formats:\n")
		for format, desc := range allfmt {
			fmt.Fprintf(os.Stderr, "    %-8s %s\n", format, desc)
		}
		fmt.Fprint(os.Stderr, "\nSupported palette color formats:\n")
		fmt.Fprintf(os.Stderr, "    %-8s %s\n", ColorFmtIA8, allfmt[ColorFmtIA8])
		fmt.Fprintf(os.Stderr, "    %-8s %s\n", ColorFmtRGB565, allfmt[ColorFmtRGB565])
		fmt.Fprintf(os.Stderr, "    %-8s %s\n", ColorFmtRGB5A3, allfmt[ColorFmtRGB5A3])
	}
	path := flag.String("in", "-", "Image file to convert (- for STDIN)")
	outpath := flag.String("out", "-", "Output file (- for STDOUT)")
	endianess := flag.String("endianess", string(EndianessBig), "Endianess of values (valid values: big, small)")
	wrap := flag.String("wrap", string(WrapClamp), "Wrapping strategy (valid values: clamp, repeat, mirror)")
	filter := flag.String("filter", string(FilterTrilinear), "Filter (valid values: near, bilinear, trilinear)")
	format := flag.String("fmt", string(ColorFmtRGBA8), "Output color format (see below for full list)")
	palettefmt := flag.String("palfmt", string(ColorFmtRGB5A3), "Palette color format (see below for full list)")
	maxlod := flag.Int("maxlod", 0, "Maximum mipmap level (0-10)")
	minlod := flag.Int("minlod", 0, "Minimum mipmap level (0-10)")
	quiet = flag.Bool("quiet", false, "Don't write info messages")
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

	if !*quiet {
		fmt.Fprintf(os.Stderr, "Image detected as: %s\n", srcformat)
		size := img.Bounds().Size()
		fmt.Fprintf(os.Stderr, "Size: %d x %d\n", size.X, size.Y)
		fmt.Fprintf(os.Stderr, "Encoding as: %s\n", *format)
	}

	checkErr(SaveTexture(img, out, TextureOptions{
		Endianess:  Endianess(*endianess),
		Format:     ColorFmt(*format),
		PaletteFmt: ColorFmt(*palettefmt),
		MaxLOD:     *maxlod,
		MinLOD:     *minlod,
		Wrap:       WrapStrategy(*wrap),
		Filter:     Filter(*filter),
	}), "Error while saving output texture")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
