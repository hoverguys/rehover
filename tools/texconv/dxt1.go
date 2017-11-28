package main

import (
	"image"
	"image/color"
	"math"
)

// Range fit DXT1 (w/ no alpha yet)
func encodeDXT1Block(tex image.Image, blockPoint image.Point, options FormatOptions) (out [8]byte) {
	// List all colors to compress
	colors := [16]color.Color{}
	for x := 0; x < 4; x++ {
		for y := 0; y < 4; y++ {
			colors[x+y*4] = tex.At(blockPoint.X+x, blockPoint.Y+y)
		}
	}

	// Get two opposite colors
	r1, g1, b1, _ := colors[0].RGBA()
	r2, g2, b2, _ := colors[15].RGBA()

	min := [3]uint8{
		uint8(r1 >> 8),
		uint8(g1 >> 8),
		uint8(b1 >> 8),
	}
	max := [3]uint8{
		uint8(r2 >> 8),
		uint8(g2 >> 8),
		uint8(b2 >> 8),
	}

	// Clamp extremes to RGB565
	min = [3]uint8{
		min[0] & 0xf8,
		min[1] & 0xfc,
		min[2] & 0xf8,
	}
	max = [3]uint8{
		max[0] & 0xf8,
		max[1] & 0xfc,
		max[2] & 0xf8,
	}

	// Use extremes as palette
	colS := color.RGBA{
		min[0],
		min[1],
		min[2],
		0xff,
	}
	colE := color.RGBA{
		max[0],
		max[1],
		max[2],
		0xff,
	}

	// Write palette to start of block
	palbytes := dxt1palette(colS, colE, options)
	copy(out[0:4], palbytes[:])

	// Calculate intermediate colors
	col3 := color.RGBA{
		lerp(min[0], max[0], 0.333),
		lerp(min[1], max[1], 0.333),
		lerp(min[2], max[2], 0.333),
		0xff,
	}
	col4 := color.RGBA{
		lerp(min[0], max[0], 0.666),
		lerp(min[1], max[1], 0.666),
		lerp(min[2], max[2], 0.666),
		0xff,
	}

	// Set-up palette for bit map generation
	palette := [4]color.RGBA{colS, col3, col4, colE}
	bitmap := uint32(0)

	// Fill bitmap with indexes
	for idx, col := range colors {
		// Calculate euclidean distance to get best palette color for the current one
		// TODO Might not be the best way to find closest color, our eyes are weird
		mincol := byte(0)
		mindist := 9999999.0
		for palidx, palcol := range palette {
			dist := distance(col, palcol)
			if dist < mindist {
				mincol = byte(palidx)
				mindist = dist
			}
		}

		bitmap |= (uint32(mincol) & 0x3) << (30 - uint(2*idx))
	}

	options.Endianess.PutUint32(out[4:8], bitmap)
	return
}

func lerp(a, b uint8, t float32) uint8 {
	af := float32(a)
	bf := float32(b)
	return uint8(af + t*(bf-af))
}

func distance(a, b color.Color) float64 {
	r1, g1, b1, _ := a.RGBA()
	r2, g2, b2, _ := b.RGBA()
	rdist := math.Pow(float64(r2)-float64(r1), 2)
	gdist := math.Pow(float64(g2)-float64(g1), 2)
	bdist := math.Pow(float64(b2)-float64(b1), 2)
	return math.Sqrt(rdist + gdist + bdist)
}

func dxt1palette(col1, col2 color.RGBA, options FormatOptions) (out [4]byte) {
	// Start color
	red := uint16(col1.R >> 3)
	green := uint16(col1.G >> 2)
	blue := uint16(col1.B >> 3)
	options.Endianess.PutUint16(out[0:], uint16(blue|(green<<5)|(red<<11)))

	// End color
	red = uint16(col2.R >> 3)
	green = uint16(col2.G >> 2)
	blue = uint16(col2.B >> 3)
	options.Endianess.PutUint16(out[2:], uint16(blue|(green<<5)|(red<<11)))

	return
}

//TODO Single color compression (as explained in http://sjbrown.co.uk/2006/01/19/dxt-compression-techniques/)
