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

// ColorFmt represents a color format (for GC textures)
type ColorFmt string

// All supported color formats
const (
	ColorFmtRGBA8 ColorFmt = "RGBA8"
)

// Description of all supported formats
var allfmt = map[ColorFmt]string{
	ColorFmtRGBA8: "Full 8bpc color (A8R8G8B8)",
}

// TextureOptions contains all the available options to tune the output texture
type TextureOptions struct {
	Endianess Endianess
	Format    ColorFmt
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

	// Write header
	binary.Write()

	switch options.Format {
	case ColorFmtRGBA8:
		// ...
	default:
		return fmt.Errorf("Unknown color format: %s (see -h for available formats)", options.Format)
	}

	// ...

	return nil
}
