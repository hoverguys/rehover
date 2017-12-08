package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
)

func logDataWrite(data []byte) {
	lines := strings.Split(string(debug.Stack()), "\n")
	// Get function call hierarchy
	const prefix = "main."
	const plen = len(prefix)
	var functions []string
	for _, line := range lines {
		// Skip file paths
		if !strings.HasPrefix(line, prefix) {
			continue
		}
		// Trim package
		funcname := strings.TrimSpace(line[plen:])

		// Do not log obvious functions
		if strings.HasPrefix(funcname, "logDataWrite(") || strings.HasPrefix(funcname, "main(") || strings.HasPrefix(funcname, "GX_LOAD_") {
			continue
		}

		// Append to list of functions
		functions = append(functions, funcname)
	}
	// Format byte dump
	bytestr := ""
	for _, byt := range data {
		bytestr += fmt.Sprintf("%02x ", byt)
	}
	// Format stack
	stack := strings.Join(functions, " ← ")

	// Print data write
	fmt.Fprintf(os.Stderr, "%-27s | %s\n", bytestr, stack)
}
