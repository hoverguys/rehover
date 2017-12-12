package main

import (
	"encoding/binary"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"

	"github.com/adinfinit/texpack/maxrect"
)

type TexPackerOptions struct {
	MaxBounds   image.Point
	WriteHeader bool
}

type TexPacker struct {
	outfile string
	options TexPackerOptions
	images  []imageInfo
}

type imageInfo struct {
	Image  image.Image
	Path   string
	Coords Rect
}

func NewTexPacker(outfile string, options TexPackerOptions) *TexPacker {
	return &TexPacker{
		outfile: outfile,
		options: options,
	}
}

// Add specifies a new texture to be packed.
func (packer *TexPacker) Add(texpath string) error {
	input, err := os.Open(texpath)
	checkErr(err, "Cannot open input file: "+texpath)
	defer input.Close()

	img, fmt, err := image.Decode(input)
	if err != nil {
		return err
	}
	if fmt != "png" && fmt != "jpeg" && fmt != "gif" {
		return errors.New("Unsupported input format: " + fmt)
	}
	packer.images = append(packer.images, imageInfo{
		Image: img,
		Path:  texpath,
	})
	return nil
}

// Save packs all the given textures into one and writes the result into `output`
func (packer *TexPacker) Save() error {
	output, err := packer.getOutput()
	if err != nil {
		return err
	}
	defer output.Close()

	outtex, err := packer.pack()
	if err != nil {
		return err
	}
	if packer.options.WriteHeader {
		if err = packer.writeHeader(output); err != nil {
			return err
		}
	}
	return png.Encode(output, outtex)
}

// getOutput opens the output file and returns its handle
func (packer *TexPacker) getOutput() (io.WriteCloser, error) {
	// Get output writer
	out := os.Stdout
	if packer.outfile != "-" {
		file, err := os.Create(packer.outfile)
		if err != nil {
			return nil, errors.New("Cannot create output file")
		}
		out = file
	}
	return out, nil
}

// pack runs the maxrect algorithms on the input textures, then writes the result
// in the output texture, which is returned
func (packer *TexPacker) pack() (image.Image, error) {

	points := getImageSizes(packer.images)

	outsize, positions, ok := minimizeFit(packer.options.MaxBounds, points)
	if !ok {
		return nil, errors.New("Couldn't pack all images!")
	}

	// Write packed texture
	outtex := image.NewRGBA(image.Rect(0, 0, outsize.X, outsize.Y))
	for i, imginfo := range packer.images {
		pos := positions[i]
		if err := writeImageAt(outtex, imginfo.Image, pos.Min.X, pos.Min.Y); err != nil {
			return nil, err
		}
		packer.images[i].Coords = Rect{
			Start: Point{uint16(pos.Min.X), uint16(pos.Min.Y)},
			Size:  Point{uint16(pos.Dx()), uint16(pos.Dy())},
		}
		// No need to keep this anymore
		packer.images[i].Image = nil
	}

	return outtex, nil
}

// writeHeader outputs binary metadata on the packed textures to the given Writer.
func (packer *TexPacker) writeHeader(output io.Writer) error {

	nEntries := len(packer.images)

	// Output file format is:
	// ParentTexture Hash [4B]
	// Entry Count        [4B]
	// Entry0             [12B]
	// ...

	// Used to check hash collisions
	hashes := make(map[FileHash]string, nEntries)

	// Parent Texture Hash
	ptHash := ToFileHash(packer.outfile)
	hashes[ptHash] = packer.outfile
	if _, err := output.Write(ptHash.Bytes()); err != nil {
		return err
	}

	// Entry Count
	countbuf := make([]byte, 4)
	binary.BigEndian.PutUint32(countbuf, uint32(nEntries))
	if _, err := output.Write(countbuf); err != nil {
		return err
	}

	// Entries
	for _, imginfo := range packer.images {
		hash := ToFileHash(imginfo.Path)
		if orig, collides := hashes[hash]; collides {
			return fmt.Errorf("Hash conflict detected between the following files:\n    [%8x] %s\n    [%8x] %s", hash, orig, hash, imginfo.Path)
		}
		hashes[hash] = imginfo.Path
		entry := Entry{
			TexPath: hash,
			Coords:  imginfo.Coords,
		}
		if _, err := output.Write(entry.Bytes()); err != nil {
			return err
		}
	}

	return nil
}

func getImageSizes(images []imageInfo) []image.Point {
	points := make([]image.Point, len(images))
	for i, imginfo := range images {
		points[i] = imginfo.Image.Bounds().Size()
	}
	return points
}

// writeImageAt adds all pixels of `src` to `dst`, starting from pixel at (`x`,`y`).
// It returns an error if the whole source image couldn't be added to the target.
func writeImageAt(dst *image.RGBA, src image.Image, x, y int) error {
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
