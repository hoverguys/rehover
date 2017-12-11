package main

import (
	"errors"
	"image"
	"image/png"
	"io"

	"github.com/adinfinit/texpack/maxrect"
)

type TexPacker struct {
	maxBounds image.Point
	images    []image.Image // List of textures to pack
}

func NewTexPacker(maxBounds image.Point) *TexPacker {
	return &TexPacker{
		maxBounds: maxBounds,
	}
}

// Add specifies a new texture to be packed.
func (packer *TexPacker) Add(input io.Reader) error {
	img, fmt, err := image.Decode(input)
	if err != nil {
		return err
	}
	if fmt != "png" && fmt != "jpeg" && fmt != "gif" {
		return errors.New("Unsupported input format: " + fmt)
	}
	packer.images = append(packer.images, img)
	return nil
}

// Save packs all the given textures into one and writes the result into `output`
func (packer *TexPacker) Save(output io.Writer) error {
	outtex, err := packer.pack()
	if err != nil {
		return err
	}
	return png.Encode(output, outtex)
}

func (packer *TexPacker) pack() (image.Image, error) {
	// Calculate output bounds (TODO)

	points := make([]image.Point, len(packer.images))
	for i, img := range packer.images {
		points[i] = img.Bounds().Size()
	}
	outsize, positions, ok := minimizeFit(packer.maxBounds, points)
	if !ok {
		return nil, errors.New("Couldn't pack all images!")
	}

	// Write packed texture
	outtex := image.NewRGBA(image.Rect(0, 0, outsize.X, outsize.Y))
	for i, img := range packer.images {
		pos := positions[i]
		if err := addImageAt(outtex, img, pos.Min.X, pos.Min.Y); err != nil {
			return nil, err
		}
	}
	return outtex, nil
}

func (packer *TexPacker) calcOutBounds() image.Rectangle {
	var w, h int
	for _, img := range packer.images {
		sz := img.Bounds().Size()
		if sz.X > w {
			w = sz.X
		}
		h += sz.Y
	}
	return image.Rect(0, 0, w, h)
}

// addImageAt adds all pixels of `src` to `dst`, starting from pixel at (`x`,`y`).
// It returns an error if the whole source image couldn't be added to the target.
func addImageAt(dst *image.RGBA, src image.Image, x, y int) error {
	srcsize := src.Bounds().Size()
	dstsize := dst.Bounds().Size()
	if srcsize.X > dstsize.X-x || srcsize.Y > dstsize.Y-y {
		return errors.New("The given image couldn't fit into the destination image")
	}

	for py := y; py < y+srcsize.Y; py++ {
		for px := x; px < x+srcsize.X; px++ {
			dst.Set(px, py, src.At(px-x, py-y))
		}
	}

	return nil
}

// Taken from https://github.com/adinfinit/texpack/blob/master/pack/fit.go
func minimizeFit(maxContextSize image.Point, sizes []image.Point) (contextSize image.Point, rects []image.Rectangle, ok bool) {

	try := func(size image.Point) ([]image.Rectangle, bool) {
		context := maxrect.New(size)
		return context.Adds(sizes...)
	}

	contextSize = maxContextSize
	rects, ok = try(contextSize)
	if !ok {
		return
	}

	shrunk, shrinkX, shrinkY := true, true, true
	for shrunk {
		shrunk = false
		if shrinkX {
			trySize := image.Point{contextSize.X - 128, contextSize.Y}
			tryRects, tryOk := try(trySize)
			if tryOk {
				contextSize = trySize
				rects = tryRects
				shrunk = true
			} else {
				shrinkX = false
			}
		}

		if shrinkY {
			trySize := image.Point{contextSize.X, contextSize.Y - 128}
			tryRects, tryOk := try(trySize)
			if tryOk {
				contextSize = trySize
				rects = tryRects
				shrunk = true
			} else {
				shrinkX = false
			}
		}
	}

	return
}
