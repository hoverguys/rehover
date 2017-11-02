package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"unicode"
)

var otpldata = `    .section .rodata
    .balign {{.Alignment}}
    .global {{.CName}}_txt
    .global {{.CName}}_txt_end
    .global {{.CName}}_txt_size

{{.CName}}_txt:
{{range .Bytes | split}}    .byte  {{. | commasep}}
{{end}}

{{.CName}}_txt_end:

    .align
{{.CName}}_txt_size: .int {{.Bytes | len}}
`

var ctpldata = `extern const unsigned char {{.CName}}_txt[];
extern const unsigned char {{.CName}}_txt_end[];
extern const unsigned int {{.CName}}_txt_size;
`

type filedata struct {
	Alignment int
	CName     string
	Bytes     []byte
}

var funcs = template.FuncMap{
	"split":    tSplit,
	"commasep": tCommasep,
}

var maxlen int

func main() {
	path := flag.String("in", "-", "Input file (- for stdin)")
	headerpath := flag.String("headerpath", ".", "Where to put the output header file")
	objectpath := flag.String("objectpath", ".", "Where to put the output asm file")
	alignment := flag.Int("align", 4, "Boundary alignment, in bytes")
	name := flag.String("name", "", "Name of the output files (filename, if empty)")
	flag.IntVar(&maxlen, "line-length", 16, "Maximum number of bytes per line")
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

	bytes, err := ioutil.ReadAll(in)
	checkErr(err, "Cannot read input file")

	data := filedata{
		Alignment: *alignment,
		CName:     strings.Map(maprune, *name),
		Bytes:     bytes,
	}

	// Create templates
	otpl := template.Must(template.New("object").Funcs(funcs).Parse(otpldata))
	ctpl := template.Must(template.New("header").Parse(ctpldata))

	// Open output files
	objfile, err := os.Create(filepath.Join(*objectpath, data.CName+".s"))
	checkErr(err, "Cannot create output object file")
	defer objfile.Close()

	headfile, err := os.Create(filepath.Join(*headerpath, data.CName+".h"))
	checkErr(err, "Cannot create output header file")
	defer headfile.Close()

	// Run the templates
	checkErr(otpl.Execute(objfile, data), "Error while writing object file")
	checkErr(ctpl.Execute(headfile, data), "Error while writing header file")
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

// Make C-friendly name
func cname(str string) string {
	// Add number to the start
	if len(str) > 0 && unicode.IsNumber(rune(str[0])) {
		str = "_" + str
	}
	// Change unfriendly characters
	return strings.Map(maprune, str)
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func tSplit(in []byte) (out [][]byte) {
	start := 0
	for start < len(in) {
		end := start + maxlen
		if end > len(in) {
			end = len(in)
		}
		out = append(out, in[start:end])
		start = end
	}
	return
}

func tCommasep(in []byte) string {
	strarr := make([]string, len(in))
	for idx, byt := range in {
		strarr[idx] = strconv.FormatUint(uint64(byt), 10)
	}
	return strings.Join(strarr, ",")
}
