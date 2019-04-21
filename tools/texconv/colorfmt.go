package main

import (
	"encoding/binary"
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
)

// FormatOptions contains all the required options to encode a picture to a specific color format
type FormatOptions struct {
	Endianess  binary.ByteOrder
	PaletteFmt ColorFmt
}

type PaletteData struct {
	Offset     uint32
	NumEntries uint16
}

// ColorFmt represents a color format (for GC textures)
type ColorFmt string

// All supported color formats
const (
	ColorFmtNone   ColorFmt = "" // Only used for checks
	ColorFmtI4     ColorFmt = "I4"
	ColorFmtI8     ColorFmt = "I8"
	ColorFmtIA4    ColorFmt = "IA4"
	ColorFmtIA8    ColorFmt = "IA8"
	ColorFmtRGB565 ColorFmt = "RGB565"
	ColorFmtRGB5A3 ColorFmt = "RGB5A3"
	ColorFmtRGBA8  ColorFmt = "RGBA8"
	ColorFmtA8     ColorFmt = "A8"
	ColorFmtCI4    ColorFmt = "CI4"
	ColorFmtCI8    ColorFmt = "CI8"
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
	ColorFmtCI4:    0x8,
	ColorFmtCI8:    0x9,
}

// Palette color formats IDs (try to match libogc)
var palfmtid = map[ColorFmt]uint8{
	ColorFmtIA8:    0x0,
	ColorFmtRGB565: 0x1,
	ColorFmtRGB5A3: 0x2,
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
	ColorFmtCI4:    "Palette-indexed 4bit (CI4)",
	ColorFmtCI8:    "Palette-indexed 8bit (CI8)",
}

type colorFmtEncoder func(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error)

var fmtEncoders = map[ColorFmt]colorFmtEncoder{
	ColorFmtI4:     encodeI4,
	ColorFmtI8:     encodeI8,
	ColorFmtIA4:    encodeIA4,
	ColorFmtIA8:    encodeIA8,
	ColorFmtRGB565: encodeRGB565,
	ColorFmtRGB5A3: encodeRGB5A3,
	ColorFmtRGBA8:  encodeRGBA8,
	ColorFmtA8:     encodeA8,
	ColorFmtCI4:    encodeCI4,
	ColorFmtCI8:    encodeCI8,
}

func encodeI4(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 8, 8, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
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

				block[baseIndex] = (grey1 << 4) | (grey2 & 0x0f)
			}
		}
		return block
	})
}

func encodeI8(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
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

func encodeA8(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
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

func encodeIA4(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get color
				col := tex.At(tilePoint.X+x, tilePoint.Y+y)
				grey, _, _, alpha := rgba8(color.GrayModel.Convert(col))

				block[baseIndex] = (grey << 4) | (alpha & 0x0f)
			}
		}
		return block
	})
}

func encodeIA8(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Set pixel color
				color := getIA8(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}
		return block
	})
}

func encodeRGB565(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Set pixel color
				color := getRGB565(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}
		return block
	})
}

func encodeRGB5A3(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		// Prepare block data
		block := make([]byte, 32, 32)

		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get converted color
				color := getRGB5A3(tex.At(tilePoint.X+x, tilePoint.Y+y))

				// Write block data
				options.Endianess.PutUint16(block[baseIndex*2:], color)
			}
		}

		return block
	})
}

func encodeRGBA8(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	return PaletteData{}, writeTiles(tex, out, options, 4, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
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

func encodeCI4(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	// Retrieve palette and color indices
	pim, ok := tex.(image.PalettedImage)
	if !ok {
		return PaletteData{}, fmt.Errorf("input image doesn't seem to be paletted")
	}
	palette, ok := tex.ColorModel().(color.Palette)
	if !ok {
		return PaletteData{}, fmt.Errorf("input image doesn't seem to be paletted")
	}

	if len(palette) > 16384 {
		return PaletteData{}, fmt.Errorf("too many palette entries (max 16384)")
	}

	maxid := uint16(0)

	err := writeTiles(tex, out, options, 8, 8, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 8; y++ {
			for x := 0; x < 4; x++ {
				// Calculate base index
				baseIndex := x + y*4

				// Get colors
				idx1 := pim.ColorIndexAt(tilePoint.X+x*2, tilePoint.Y+y)
				idx2 := pim.ColorIndexAt(tilePoint.X+x*2+1, tilePoint.Y+y)

				if uint16(idx1) > maxid {
					maxid = uint16(idx1)
				}
				if uint16(idx2) > maxid {
					maxid = uint16(idx2)
				}

				block[baseIndex] = (idx1 << 4) | (idx2 & 0x0f)
			}
		}
		return block
	})
	if err != nil {
		return PaletteData{}, err
	}

	if maxid > 16 {
		return PaletteData{}, fmt.Errorf("image needs more than 16 colors (max for CI4), try CI8")
	}

	paletteoffset, err := out.Seek(0, io.SeekCurrent)
	if err != nil {
		return PaletteData{}, err
	}

	if paletteoffset%32 > 0 {
		paddedlen := ((paletteoffset/32)+1)*32 - paletteoffset
		empty := make([]byte, 32, 32)
		_, err = out.Write(empty[:paddedlen-paletteoffset])
		if err != nil {
			return PaletteData{}, fmt.Errorf("error while copying padding for padding: %s", err.Error())
		}

		paletteoffset += paddedlen
	}

	// Serialize palette and write it
	palettebytes := makePaletteBlock(palette, options, maxid+1)
	_, err = out.Write(palettebytes)
	return PaletteData{uint32(paletteoffset), maxid + 1}, err
}

func encodeCI8(tex image.Image, out io.WriteSeeker, options FormatOptions) (PaletteData, error) {
	// Retrieve palette and color indices
	pim, ok := tex.(image.PalettedImage)
	if !ok {
		return PaletteData{}, fmt.Errorf("input image doesn't seem to be paletted")
	}
	palette, ok := tex.ColorModel().(color.Palette)
	if !ok {
		return PaletteData{}, fmt.Errorf("input image doesn't seem to be paletted")
	}

	if len(palette) > 16384 {
		return PaletteData{}, fmt.Errorf("too many palette entries (max 16384)")
	}

	maxid := uint16(0)

	// Write indices as data
	err := writeTiles(tex, out, options, 8, 4, func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte {
		block := make([]byte, 32, 32)
		// Iterate through each pixel in the block
		for y := 0; y < 4; y++ {
			for x := 0; x < 8; x++ {
				// Calculate base index
				baseIndex := x + y*8

				// Set color index
				block[baseIndex] = pim.ColorIndexAt(tilePoint.X+x, tilePoint.Y+y)

				if uint16(block[baseIndex]) > maxid {
					maxid = uint16(block[baseIndex])
				}
			}
		}
		return block
	})
	if err != nil {
		return PaletteData{}, err
	}

	paletteoffset, err := out.Seek(0, io.SeekCurrent)
	if err != nil {
		return PaletteData{}, err
	}

	if paletteoffset%32 > 0 {
		paddedlen := ((paletteoffset/32)+1)*32 - paletteoffset
		empty := make([]byte, 32, 32)
		_, err = out.Write(empty[:paddedlen-paletteoffset])
		if err != nil {
			return PaletteData{}, fmt.Errorf("error while copying padding for padding: %s", err.Error())
		}

		paletteoffset += paddedlen
	}

	// Serialize palette and write it
	palettebytes := makePaletteBlock(palette, options, maxid+1)
	_, err = out.Write(palettebytes)
	return PaletteData{uint32(paletteoffset), maxid + 1}, err
}

type tileFunc func(tex image.Image, tilePoint image.Point, options FormatOptions) []byte

func writeTiles(tex image.Image, out io.WriteSeeker, options FormatOptions, tileWidth, tileHeight int, tilefn tileFunc) error {
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

func makePaletteBlock(palette color.Palette, options FormatOptions, maxPalette uint16) []byte {
	// Sanity check (probably not needed)
	if maxPalette > uint16(len(palette)) {
		maxPalette = uint16(len(palette))
	}
	// Create buffer for palette
	palettedata := make([]byte, 2*maxPalette)
	for index, col := range palette {
		if uint16(index) >= maxPalette {
			break
		}
		// Get color converted to chosen palette color format
		var colordata uint16
		switch options.PaletteFmt {
		case ColorFmtIA8:
			colordata = getIA8(col)
		case ColorFmtRGB565:
			colordata = getRGB565(col)
		case ColorFmtRGB5A3:
			colordata = getRGB5A3(col)
		}
		// Write into the palette data block
		options.Endianess.PutUint16(palettedata[index*2:], colordata)
	}
	return palettedata
}

func rgba8(c color.Color) (r, g, b, a uint8) {
	r32, g32, b32, a32 := c.RGBA()
	return uint8(r32 >> 8), uint8(g32 >> 8), uint8(b32 >> 8), uint8(a32 >> 8)
}

func getRGB5A3(c color.Color) uint16 {
	// Get color
	r, g, b, a := rgba8(c)

	// Decide between RGB5 and RGB4A3 depending on alpha
	if a < 0xff {
		// RGB4A3
		red := r >> 4
		green := g >> 4
		blue := b >> 4
		alpha := a >> 5
		return uint16(uint16(blue) | (uint16(green) << 4) | (uint16(red) << 8) | (uint16(alpha) << 12))
	}

	// RGB5
	red := r >> 3
	green := g >> 3
	blue := b >> 3
	return uint16(uint16(blue) | (uint16(green) << 5) | (uint16(red) << 10) | 0x8000)
}

func getRGB565(c color.Color) uint16 {
	// Get color (ignore alpha)
	r, g, b, _ := rgba8(c)

	// Get only the required bits
	red := r >> 3
	green := g >> 2
	blue := b >> 3

	return uint16(uint16(blue) | (uint16(green) << 5) | (uint16(red) << 11))
}

func getIA8(c color.Color) uint16 {
	grey, _, _, alpha := rgba8(color.GrayModel.Convert(c))

	// Set pixel color
	return uint16(uint16(grey) | (uint16(alpha) << 8))
}
