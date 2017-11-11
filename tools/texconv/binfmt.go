package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"io"
)

// Endianess represents a byte order
type Endianess string

// All supported byte orders
const (
	EndianessLittle Endianess = "little" // Little endian (x86)
	EndianessBig    Endianess = "big"    // Big endian (PPC)
)

// TextureOptions contains all the available options to tune the output texture
type TextureOptions struct {
	Endianess Endianess
	Format    ColorFmt
	MaxLOD    int
	MinLOD    int
}

// SaveTexture takes a parsed image file and writes it in binary format using a provided byte order
func SaveTexture(tex image.Image, out io.Writer, options TextureOptions) error {
	var endianess binary.ByteOrder
	switch options.Endianess {
	case EndianessLittle:
		endianess = binary.LittleEndian
	case EndianessBig:
		endianess = binary.BigEndian
	default:
		return fmt.Errorf("Unknown endianess: %s (supported: big, little)", options.Endianess)
	}

	bounds := tex.Bounds()
	width := bounds.Max.X - bounds.Min.X
	height := bounds.Max.Y - bounds.Min.X
	mipmap := byte((options.MaxLOD << 8) | (options.MinLOD & 0xf))

	fmtfn, ok := fmtEncoders[options.Format]
	if !ok {
		return fmt.Errorf("Unknown color format: %s (see -h for available formats)", options.Format)
	}

	// Write header
	const headerSize = 6
	binary.Write(out, endianess, uint16(width))
	binary.Write(out, endianess, uint16(height))
	binary.Write(out, endianess, fmtid[options.Format])
	binary.Write(out, endianess, mipmap)

	// Pad to 32B
	padlen := 32 - headerSize
	padding := make([]byte, padlen, padlen)
	binary.Write(out, endianess, padding)

	fmtfn(tex, out, FormatOptions{
		Endianess: endianess,
	})

	return nil
}
