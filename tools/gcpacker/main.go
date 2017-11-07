package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

// ResourceFile is an entry int he resources list file containing the identifies and compiled resource paths
type ResourceFile struct {
	Identifier string
	File       string
}

func main() {
	list := flag.String("list", "resources.txt", "Path to resource list")
	outpath := flag.String("out", "-", "Output file (- for stdout)")
	ignoreconflict := flag.Bool("ignoreconflict", false, "Ignore conflicts (why would you do that)")
	flag.Parse()

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

	file, err := os.Open(*list)
	checkErr(err, "Could not open file list")

	resources := []ResourceFile{}
	rows, err := csv.NewReader(file).ReadAll()
	checkErr(err, "Could not parse file list")

	for lno, record := range rows {
		if len(record) < 2 {
			fmt.Fprintf(os.Stderr, "[WARN] Line no. %d has invalid resource record (not enough fields): %v\n", lno, record)
			continue
		}
		resources = append(resources, ResourceFile{
			Identifier: record[0],
			File:       record[1],
		})
	}

	packer := NewPacker(out, PackerOptions{
		IgnoreConflicts: *ignoreconflict,
	})
	checkErr(packer.Pack(resources), "Error while generating GCR file")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
