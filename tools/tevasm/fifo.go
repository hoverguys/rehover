package main

import (
	"encoding/binary"
	"io"
)

type Fifo interface {
	U8(x uint8)
	U32(x uint32)
	Pad32()
}

type NullFifo struct{}

func (f NullFifo) U8(x uint8) {
	// Ignore
}

func (f NullFifo) U32(x uint32) {
	// Ignore
}

func (f NullFifo) Pad32() {
	// Ignore
}

type FifoWriter struct {
	Endianess binary.ByteOrder

	writer  io.Writer
	written int
}

func NewFifo(writer io.Writer, endianess binary.ByteOrder) *FifoWriter {
	return &FifoWriter{
		Endianess: endianess,
		writer:    writer,
		written:   0,
	}
}

func (f *FifoWriter) U8(x uint8) {
	binary.Write(f.writer, f.Endianess, x)
	f.written++
}

func (f *FifoWriter) U32(x uint32) {
	binary.Write(f.writer, f.Endianess, x)
	f.written += 4
}

func (f *FifoWriter) Pad32() {
	empty := make([]byte, 32, 32)
	paddedlen := 32 * ((f.written-1)/32 + 1)
	f.writer.Write(empty[:paddedlen-f.written])
}
