package main

import (
	"encoding/binary"
	"image"
	"io"
	"math"
)

// FormatOptions contains all the required options to encode a picture to a specific color format
type FormatOptions struct {
	Endianess binary.ByteOrder
}

// ColorFmt represents a color format (for GC textures)
type ColorFmt string

// All supported color formats
const (
	ColorFmtRGBA8 ColorFmt = "RGBA8"
)

// Color formats IDs (try to match libogc)
var fmtid = map[ColorFmt]uint8{
	ColorFmtRGBA8: 6,
}

// Description of all supported formats
var allfmt = map[ColorFmt]string{
	ColorFmtRGBA8: "Full 8bpc color (A8R8G8B8)",
}

type colorFmtEncoder func(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error)

var fmtEncoders = map[ColorFmt]colorFmtEncoder{
	ColorFmtRGBA8: encodeRGBA8,
}

func encodeRGBA8(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	size := tex.Bounds().Size()
	width := size.X
	height := size.Y

	rows := int(math.Ceil(float64(height) / 4))
	cols := int(math.Ceil(float64(width) / 4))

	// Iterate through each block
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Prepare block data
			var block [64]byte

			// Iterate through each pixel in the block
			for y := 0; y < 4; y++ {
				for x := 0; x < 4; x++ {
					// Calculate base index
					baseIndex := x + y*4

					// Get color
					r, g, b, a := tex.At(col*4+x, row*4+y).RGBA()

					// Set colors
					block[baseIndex*2] = byte(a >> 8)
					block[baseIndex*2+1] = byte(r >> 8)
					block[32+baseIndex*2] = byte(g >> 8)
					block[32+baseIndex*2+1] = byte(b >> 8)
				}
			}

			// Write block data
			_, err := out.Write(block[:])
			if err != nil {
				return 0, err
			}
		}
	}

	return 0, nil
}
