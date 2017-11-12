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
	ColorFmtRGB565 ColorFmt = "RGB565"
	ColorFmtRGB5A3 ColorFmt = "RGB5A3"
	ColorFmtRGBA8  ColorFmt = "RGBA8"
)

// Color formats IDs (try to match libogc)
var fmtid = map[ColorFmt]uint8{
	ColorFmtRGB565: 0x4,
	ColorFmtRGB5A3: 0x5,
	ColorFmtRGBA8:  0x6,
}

// Description of all supported formats
var allfmt = map[ColorFmt]string{
	ColorFmtRGB565: "Limited color (R5G6B5)",
	ColorFmtRGB5A3: "Limited color and alpha (R4B4G4A3|R5B5G5)",
	ColorFmtRGBA8:  "Full 8bpc color (A8R8G8B8)",
}

type colorFmtEncoder func(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error)

var fmtEncoders = map[ColorFmt]colorFmtEncoder{
	ColorFmtRGB565: encodeRGB565,
	ColorFmtRGB5A3: encodeRGB5A3,
	ColorFmtRGBA8:  encodeRGBA8,
}

func encodeRGB565(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	size := tex.Bounds().Size()
	width := size.X
	height := size.Y

	rows := int(math.Ceil(float64(height) / 4))
	cols := int(math.Ceil(float64(width) / 4))

	// Iterate through each block
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Iterate through each pixel in the block
			for y := 0; y < 4; y++ {
				for x := 0; x < 4; x++ {
					// Get color (ignore alpha)
					r, g, b, _ := tex.At(col*4+x, row*4+y).RGBA()

					// Get only the required bits
					red := r >> 11
					green := g >> 10
					blue := b >> 11

					// Set pixel color
					color := uint16(blue | (green << 5) | (red << 11))

					// Write block data
					err := binary.Write(out, options.Endianess, color)
					if err != nil {
						return 0, err
					}
				}
			}
		}
	}

	return 0, nil
}

func encodeRGB5A3(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	size := tex.Bounds().Size()
	width := size.X
	height := size.Y

	rows := int(math.Ceil(float64(height) / 4))
	cols := int(math.Ceil(float64(width) / 4))

	// Iterate through each block
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			// Iterate through each pixel in the block
			for y := 0; y < 4; y++ {
				for x := 0; x < 4; x++ {
					// Get color
					r, g, b, a := tex.At(col*4+x, row*4+y).RGBA()

					// Decide between RGB5 and RGB4A3 depending on alpha
					var color uint16
					if a < 0xff00 {
						// RGB4A3
						red := r >> 12
						green := g >> 12
						blue := b >> 12
						alpha := a >> 13
						color = uint16(blue | (green << 4) | (red << 8) | (alpha << 12))
					} else {
						// RGB5
						red := r >> 11
						green := g >> 11
						blue := b >> 11
						color = uint16(blue | (green << 5) | (red << 10) | 0x8000)
					}

					// Write block data
					err := binary.Write(out, options.Endianess, color)
					if err != nil {
						return 0, err
					}
				}
			}
		}
	}

	return 0, nil
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
