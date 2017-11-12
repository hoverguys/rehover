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

				grey1, _, _, _ := color.GrayModel.Convert(color1).RGBA()
				grey2, _, _, _ := color.GrayModel.Convert(color2).RGBA()

				block[baseIndex] = byte(((grey1 >> 12) << 4) | (grey2 >> 12))
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
				baseIndex := x + y*4

				// Get color
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				grey, _, _, _ := color.GrayModel.Convert(col).RGBA()

				block[baseIndex] = byte(grey >> 8)
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
				grey, _, _, alpha := color.GrayModel.Convert(col).RGBA()

				block[baseIndex] = byte((grey>>12)<<4 | (alpha >> 12))
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
				grey, _, _, alpha := color.GrayModel.Convert(col).RGBA()

				// Set pixel color
				color := uint16(grey>>8 | ((alpha >> 8) << 8))

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
				r, g, b, _ := tex.At(tilePoint.X+x, tilePoint.Y+y).RGBA()

				// Get only the required bits
				red := r >> 11
				green := g >> 10
				blue := b >> 11

				// Set pixel color
				color := uint16(blue | (green << 5) | (red << 11))

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
				r, g, b, a := tex.At(tilePoint.X+x, tilePoint.Y+y).RGBA()

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
				r, g, b, a := tex.At(tilePoint.X+x, tilePoint.Y+y).RGBA()

				// Set colors
				block[baseIndex*2] = byte(a >> 8)
				block[baseIndex*2+1] = byte(r >> 8)
				block[32+baseIndex*2] = byte(g >> 8)
				block[32+baseIndex*2+1] = byte(b >> 8)
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
