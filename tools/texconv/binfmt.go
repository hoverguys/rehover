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

// WrapStrategy represents a wrapping strategy for both direction
type WrapStrategy string

// All supported wrap strategies
const (
	WrapClamp  WrapStrategy = "clamp"
	WrapRepeat WrapStrategy = "repeat"
	WrapMirror WrapStrategy = "mirror"
)

var wrapvalue = map[WrapStrategy]uint8{
	WrapClamp:  0,
	WrapRepeat: 1,
	WrapMirror: 2,
}

// Filter represents a filter to apply to the texture
type Filter string

// All supported filters
const (
	FilterNear      Filter = "near"
	FilterBilinear  Filter = "bilinear"
	FilterTrilinear Filter = "trilinear"
)

// TextureOptions contains all the available options to tune the output texture
type TextureOptions struct {
	Endianess Endianess
	Format    ColorFmt
	MaxLOD    int
	MinLOD    int
	Wrap      WrapStrategy
	Filter    Filter
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
		return fmt.Errorf("Unknown endianess: %s (see -h for available formats)", options.Endianess)
	}

	size := tex.Bounds().Size()
	width := size.X
	height := size.Y
	mipmap := byte((options.MaxLOD << 8) | (options.MinLOD & 0xf))

	wrap, ok := wrapvalue[options.Wrap]
	if !ok {
		return fmt.Errorf("Unknown wrap mode: %s (see -h for available formats)", options.Wrap)
	}

	var filter uint8
	switch options.Filter {
	case FilterNear:
		filter = 0
	case FilterBilinear:
		filter = 2
	case FilterTrilinear:
		filter = 3
	default:
		return fmt.Errorf("Unknown filter: %s (see -h for available formats)", options.Wrap)
	}

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
	header[6] = filter | wrap<<2 | wrap<<5

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
	endianess.PutUint32(header[7:], uint32(headerlen))

	// Add palette offset if color fmt has it
	//switch (options.Format) {
	//	case <FORMAT>
	//	endianess.PutUint32(header[10:], uint32(headerlen)+paletteoffset)
	//}
	_ = paletteoffset

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
