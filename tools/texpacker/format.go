package main

import (
	"encoding/binary"
	"fmt"
	"hash/fnv"
)

// 4 bytes
type Point struct {
	X, Y uint16
}

// 8 bytes
type Rect struct {
	Start, Size Point
}

type FileHash uint32

// 4 + 4 = 8 bytes
type Header struct {
	ParentTexture FileHash
	EntryCount    int
}

// An Entry is a single texture's path + coordinates. 4 + 8 = 12 bytes
type Entry struct {
	TexPath FileHash
	Coords  Rect
}

var hash = fnv.New32()

func ToFileHash(s string) FileHash {
	hash.Reset()
	_, err := fmt.Fprintf(hash, s)
	if err != nil {
		panic(err)
	}
	return FileHash(hash.Sum32())
}

func (f FileHash) Bytes() []byte {
	out := make([]byte, 4)
	binary.BigEndian.PutUint32(out, uint32(f))
	return out
}

func (e Entry) Bytes() []byte {
	out := make([]byte, 12)
	binary.BigEndian.PutUint32(out, uint32(e.TexPath))
	binary.BigEndian.PutUint16(out, e.Coords.Start.X)
	binary.BigEndian.PutUint16(out, e.Coords.Start.Y)
	binary.BigEndian.PutUint16(out, e.Coords.Size.X)
	binary.BigEndian.PutUint16(out, e.Coords.Size.Y)
	return out
}
