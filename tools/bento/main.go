package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

var otpldata = `
.section .rodata
.balign {{.Alignment}}
.global {{.CName}}_txt
.global {{.CName}}_txt_end
.global {{.CName}}_txt_size

{{.CName}}_txt:
{{range .Bytes | split}}
.byte  {{. | commasep}}
{{end}}

{{.CName}}_txt_end:

.align
{{.CName}}_txt_size: .int {{.Bytes | len}}
`

var ctpldata = `
extern const unsigned char {{.CName}}_txt[];
extern const unsigned char {{.CName}}_txt_end[];
extern const unsigned int {{.CName}}_txt_size;
`

type FileData struct {
	CName string
	Bytes []byte
}

func main() {
	alignment := flag.Int("align", 4, "Boundary alignment, in bytes")
	maxlen := flag.Int("line-length", 16, "Maximum number of bytes per line")
	name := flag.String("name", "", "Name of the output struct (filename, if empty)")
	path := flag.String("in", "-", "Input file (- for stdin)")
	outpath := flag.String("out", "-", "Output file (- for stdout)")
	flag.Parse()

	// Get input reader
	in := os.Stdin
	if *path != "-" {
		file, err := os.Open(*path)
		checkErr(err, "Cannot open input file")
		defer file.Close()
		in = file
		if *name == "" {
			*name = filepath.Base(*path)
		}
	} else {
		// cannot use stdin without name parameter
		if *name == "" {
			fmt.Println("Cannot use stdin as input without a -name")
			flag.Usage()
			os.Exit(1)
		}
	}

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

	bytes, err := ioutil.ReadAll(in)
	checkErr(err, "Cannot read input file")

	filedata := FileData{
		CName: strings.Map(maprune, *name),
		Bytes: bytes,
	}
}

// Replace any non alphanumeric char with _
func maprune(r rune) rune {
	switch {
	case unicode.IsLetter(r), unicode.IsNumber(r):
		return r
	default:
		return '_'
	}
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
