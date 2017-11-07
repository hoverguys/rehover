package main

import (
	"encoding/binary"
	"fmt"
)

// ResourceHeader is the header of a GCR
type ResourceHeader []FileHeader

// Bytes returns the header as bytes to be included at the top of the GCR header
func (r ResourceHeader) Bytes() []byte {
	filesize := getHeaderSize(len(r))
	out := make([]byte, filesize)
	binary.BigEndian.PutUint32(out, uint32(len(r)))
	offset := 4
	for _, file := range r {
		bytes := file.Bytes()
		copy(out[offset:], bytes[:])
		offset += FileHeaderSize
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

// FileHeaderSize is the size of a single file entry in the GCR header
const FileHeaderSize = 12

// Bytes returns the file entry header to be included in the GCR header
func (f FileHeader) Bytes() (out [FileHeaderSize]byte) {
	binary.BigEndian.PutUint32(out[0:], f.Hash)
	binary.BigEndian.PutUint32(out[4:], f.Offset)
	binary.BigEndian.PutUint32(out[8:], f.Length)
	return
}

func getHeaderSize(numfiles int) uint32 {
	// count(uint32) + each file entry(3*uint32)
	return uint32(4 + FileHeaderSize*numfiles)
}
