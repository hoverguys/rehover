package main

import (
	"bytes"
	"encoding/binary"
)

type FifoWriter struct {
	Endianess binary.ByteOrder

	internal *bytes.Buffer
}

func NewFifo(endianess binary.ByteOrder) *FifoWriter {
	return &FifoWriter{
		Endianess: endianess,
		internal:  new(bytes.Buffer),
	}
}

func (f *FifoWriter) U8(x uint8) {
	binary.Write(f.internal, f.Endianess, x)
}

func (f *FifoWriter) U32(x uint32) {
	binary.Write(f.internal, f.Endianess, x)
}

func (f *FifoWriter) Buffer() *bytes.Buffer {
	return f.internal
}
