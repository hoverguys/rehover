package main

import (
	"errors"
	"image"
	"image/png"
	"io"
)

type TexPacker struct {
	images []image.Image // List of textures to pack
}

func NewTexPacker() *TexPacker {
	return &TexPacker{}
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
	outbounds := packer.calcOutBounds()

	// Write packed texture
	outtex := image.NewRGBA(outbounds)
	var x, y int
	for _, img := range packer.images {
		if err := addImageAt(outtex, img, x, y); err != nil {
			return nil, err
		}
		y += img.Bounds().Size().Y
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
