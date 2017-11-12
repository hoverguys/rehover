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
func SaveTexture(tex image.Image, out io.WriteSeeker, options TextureOptions) error {
	var endianess binary.ByteOrder
	switch options.Endianess {
	case EndianessLittle:
		endianess = binary.LittleEndian
	case EndianessBig:
		endianess = binary.BigEndian
	default:
		return fmt.Errorf("Unknown endianess: %s (supported: big, little)", options.Endianess)
	}

	size := tex.Bounds().Size()
	width := size.X
	height := size.Y
	mipmap := byte((options.MaxLOD << 8) | (options.MinLOD & 0xf))

	fmtfn, ok := fmtEncoders[options.Format]
	if !ok {
		return fmt.Errorf("Unknown color format: %s (see -h for available formats)", options.Format)
	}

	// Make header
	header := make([]byte, 32, 32)
	endianess.PutUint16(header[0:], uint16(width))
	endianess.PutUint16(header[2:], uint16(height))
	header[4] = fmtid[options.Format]
	header[5] = mipmap

	// Write padding for now
	headerlen := len(header)
	padding := make([]byte, headerlen, headerlen)
	binary.Write(out, endianess, padding)

	paletteoffset, err := fmtfn(tex, out, FormatOptions{
		Endianess: endianess,
	})
	if err != nil {
		return fmt.Errorf("Error encoding and writing texture data: %s", err.Error())
	}

	// Add data offsets to header
	endianess.PutUint32(header[6:], uint32(headerlen))

	// Add palette offset if color fmt has it
	//switch (options.Format) {
	//	case <FORMAT>
	//	endianess.PutUint32(header[10:], uint32(headerlen)+paletteoffset)
	//}

	// Write header to file
	// Write header
	_, err = out.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("Error while seeking to top: %s", err.Error())
	}

	_, err = out.Write(header)
	if err != nil {
		return fmt.Errorf("Error while writing GCR header: %s", err.Error())
	}

	return nil
}
