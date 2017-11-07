package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	conf := flag.String("list", "resources.txt", "Path to resource list")
	outpath := flag.String("out", "-", "Output file (- for stdout)")
	flag.Parse()

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
