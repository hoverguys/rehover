package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-fs/fat"

	fs "github.com/mitchellh/go-fs"
)

const DEFAULTSIZE = 24 * 1024 * 1024 // 24MiB

func main() {
	var files multiString
	flag.Var(&files, "in", "File to put in the SD card (can be specified multiple times)")
	size := flag.Int64("size", DEFAULTSIZE, "Filesystem size")
	output := flag.String("out", "sd.raw", "Output file containing FAT16 filesystem")
	label := flag.String("label", "sd", "Disk label")
	flag.Parse()

	// Print recap
	fmt.Fprintf(os.Stderr, "Building SD card:\n  File: %s\n  Size: %d bytes\n  Label: %s\n\n", *output, *size, *label)

	// Create output file
	out, err := os.Create(*output)
	checkErr(err, "Could not create output file")
	err = out.Truncate(*size)
	checkErr(err, "Could not grow output file to %d", *size)

	// Create and format to FAT16
	device, err := fs.NewFileDisk(out)
	checkErr(err, "Could not create FAT device")

	err = fat.FormatSuperFloppy(device, &fat.SuperFloppyConfig{
		FATType: fat.FAT16,
		Label:   *label,
		OEMName: *label,
	})
	checkErr(err, "Could not format file to FAT16")

	// Create usable FAT16 instance from formatted "device"
	fatFs, err := fat.New(device)
	checkErr(err, "Could not create FAT16 instance from FAT16 file")

	root, err := fatFs.RootDir()
	checkErr(err, "Could not get to root dir")

	// Read each file from input and put them in the filesystem
	for _, filename := range files {
		basename := filepath.Base(filename)

		// Write message to stderr
		fmt.Fprintf(os.Stderr, "Adding: %s -> /%s\n", filename, basename)

		// Open input file
		srcFile, err := os.Open(filename)
		checkErr(err, "Could not open input file %s", filename)

		// Create new entry in the filesystem
		entry, err := root.AddFile(basename)
		checkErr(err, "Could not create file entry for %s", filename)

		// Get entry file handle
		fatFile, err := entry.File()
		checkErr(err, "Could not get handle for file entry %s", entry.Name())

		// Copy all data from input file to entry
		_, err = io.Copy(fatFile, srcFile)
		checkErr(err, "Could not copy data from input file %s to entry %s", filename, entry.Name())

		// Close input file
		err = srcFile.Close()
		checkErr(err, "Could not close input file %s", filename)
	}

	// All done, close output file
	err = out.Close()
	checkErr(err, "Could not close output file")

	fmt.Fprintf(os.Stderr, "\nAll done\n")
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

type multiString []string

func (i *multiString) String() string {
	return strings.Join(*i, ",")
}

func (i *multiString) Set(value string) error {
	*i = append(*i, value)
	return nil
}
