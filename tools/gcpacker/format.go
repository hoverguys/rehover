package main

import (
	"encoding/binary"
	"fmt"
)

// ResourceHeader is the header of a GCR
type ResourceHeader []FileHeader

// Bytes returns the header as bytes to be included at the top of the GCR header
func (r ResourceHeader) Bytes() []byte {
	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, uint32(len(r)))
	for _, file := range r {
		bytes := file.Bytes()
		out = append(out, bytes[:]...)
	}
	return out
}

// FileHeader is a single file entry in the GCR header
type FileHeader struct {
	Hash   uint32
	Offset uint32
	Length uint32
}

func (f FileHeader) String() string {
	return fmt.Sprintf("File %d (offset: %x | %d bytes)", f.Hash, f.Offset, f.Length)
}

// Bytes returns the file entry header to be included in the GCR header
func (f FileHeader) Bytes() (out [12]byte) {
	binary.BigEndian.PutUint32(out[0:], f.Hash)
	binary.BigEndian.PutUint32(out[4:], f.Offset)
	binary.BigEndian.PutUint32(out[8:], f.Length)
	return
}
