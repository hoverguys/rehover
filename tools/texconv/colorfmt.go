package main

import (
	"encoding/binary"
	"image"
	"image/color"
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
	ColorFmtI4     ColorFmt = "I4"
	ColorFmtI8     ColorFmt = "I8"
	ColorFmtIA4    ColorFmt = "IA4"
	ColorFmtIA8    ColorFmt = "IA8"
	ColorFmtRGB565 ColorFmt = "RGB565"
	ColorFmtRGB5A3 ColorFmt = "RGB5A3"
	ColorFmtRGBA8  ColorFmt = "RGBA8"
	ColorFmtA8     ColorFmt = "A8"
)

// Color formats IDs (try to match libogc)
var fmtid = map[ColorFmt]uint8{
	ColorFmtI4:     0x0,
	ColorFmtI8:     0x1,
	ColorFmtIA4:    0x2,
	ColorFmtIA8:    0x3,
	ColorFmtRGB565: 0x4,
	ColorFmtRGB5A3: 0x5,
	ColorFmtRGBA8:  0x6,
	ColorFmtA8:     0x7,
}

// Description of all supported formats
var allfmt = map[ColorFmt]string{
	ColorFmtI4:     "Grayscale 4bit (I4)",
	ColorFmtI8:     "Grayscale 8bit (I8)",
	ColorFmtIA4:    "Grayscale + Alpha 4bit (I4A4)",
	ColorFmtIA8:    "Grayscale + Alpha 8bit (I8A8)",
	ColorFmtRGB565: "Limited color (R5G6B5)",
	ColorFmtRGB5A3: "Limited color and alpha (R4B4G4A3|R5B5G5)",
	ColorFmtRGBA8:  "Full 8bpc color (A8R8G8B8)",
	ColorFmtA8:     "Alpha-only 8bit (A8)",
}

type colorFmtEncoder func(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error)

var fmtEncoders = map[ColorFmt]colorFmtEncoder{
	ColorFmtI4:     encodeI4,
	ColorFmtI8:     encodeI8,
	ColorFmtIA4:    encodeIA4,
	ColorFmtIA8:    encodeIA8,
	ColorFmtRGB565: encodeRGB565,
	ColorFmtRGB5A3: encodeRGB5A3,
	ColorFmtRGBA8:  encodeRGBA8,
	ColorFmtA8:     encodeA8,
}

func encodeI4(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 8, 8, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 8; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get colors
				color1 := tex.At(tilePoint.X+x*2, tilePoint.Y+y)
				color2 := tex.At(tilePoint.X+x*2+1, tilePoint.Y+y)

				grey1, _, _, _ := rgba8(color.GrayModel.Convert(color1))
				grey2, _, _, _ := rgba8(color.GrayModel.Convert(color2))

				block[baseIndex] = (grey1 & 0xf0) | (grey2 >> 4)
			}
		}
		return block
	})
}

func encodeI8(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				// Calculate base index
				baseIndex := x + y*8

				// Get color
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				grey, _, _, _ := rgba8(color.GrayModel.Convert(col))

				block[baseIndex] = grey
			}
		}
		return block
	})
}

func encodeA8(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				// Calculate base index
				baseIndex := x + y*8

				// Get alpha
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				_, _, _, alpha := rgba8(col)

				block[baseIndex] = alpha
			}
		}
		return block
	})
}

func encodeIA4(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				grey, _, _, alpha := rgba8(color.GrayModel.Convert(col))

				block[baseIndex] = (grey & 0xf0) | (alpha >> 4)
			}
		}
		return block
	})
}

func encodeIA8(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color (ignore alpha)
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				grey, _, _, alpha := rgba8(color.GrayModel.Convert(col))

				// Set pixel color
				color := uint16(uint16(grey) | (uint16(alpha) << 8))

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}
		return block
	})
}

func encodeRGB565(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color (ignore alpha)
				r, g, b, _ := rgba8(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Get only the required bits
				red := r >> 3
				green := g >> 2
				blue := b >> 3

				// Set pixel color
				color := uint16(uint16(blue) | (uint16(green) << 5) | (uint16(red) << 11))

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}
		return block
	})
}

func encodeRGB5A3(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color
				r, g, b, a := rgba8(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Decide between RGB5 and RGB4A3 depending on alpha
				var color uint16
				if a < 0xff {
					// RGB4A3
					red := r >> 4
					green := g >> 4
					blue := b >> 4
					alpha := a >> 5
					color = uint16(uint16(blue) | (uint16(green) << 4) | (uint16(red) << 8) | (uint16(alpha) << 12))
				} else {
					// RGB5
					red := r >> 3
					green := g >> 3
					blue := b >> 3
					color = uint16(uint16(blue) | (uint16(green) << 5) | (uint16(red) << 10) | 0x8000)
				}

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}

		return block
	})
}

func encodeRGBA8(tex image.Image, out io.Writer, options FormatOptions) (paletteOffset uint32, err error) {
	return 0, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 64, 64)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color
				r, g, b, a := rgba8(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Set colors
				block[baseIndex*2] = a
				block[baseIndex*2+1] = r
				block[32+baseIndex*2] = g
				block[32+baseIndex*2+1] = b
			}
		}

		return block
	})
}

type tileFunc func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte

func writeTiles(tex image.Image, out io.Writer, options FormatOptions, tileWidth, tileHeight int, tilefn tileFunc) error {
	size := tex.Bounds().Size()
	width := size.X
	height := size.Y

	rows := int(math.Ceil(float64(height) / float64(tileHeight)))
	cols := int(math.Ceil(float64(width) / float64(tileWidth)))

	// Iterate through each block
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			sx := col * tileWidth
			sy := row * tileHeight
			tileData := tilefn(tex, image.Point{sx, sy}, options)
			_, err := out.Write(tileData)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func rgba8(c color.Color) (r, g, b, a uint8) {
	r32, g32, b32, a32 := c.RGBA()
	return uint8(r32 >> 8), uint8(g32 >> 8), uint8(b32 >> 8), uint8(a32 >> 8)
}
