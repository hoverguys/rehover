package main

import (
	"fmt"
	"hash/fnv"
	"io"
	"os"
)

// PackerOptions is a list of options that can change the packer's behavior
type PackerOptions struct {
	IgnoreConflicts bool // Ignore conflicts (does not give a fatal error on conflict)
}

// Packer encodes a list of files into a single GCR file
type Packer struct {
	handle  io.WriteSeeker
	options PackerOptions
}

// NewPacker returns an instance of a packer that writes to a specified file
func NewPacker(writer io.WriteSeeker, options PackerOptions) *Packer {
	return &Packer{
		handle:  writer,
		options: options,
	}
}

// Pack packs all the files in a GCR file
// Note: if called multiple times it will overwrite the old content!
func (p *Packer) Pack(files []ResourceFile) error {
	header := make(ResourceHeader, len(files))
	headersize := getHeaderSize(len(files))

	// Skip headers for now (write zeroes)
	empty := make([]byte, headersize, headersize)
	_, err := p.handle.Write(empty)
	if err != nil {
		return fmt.Errorf("Error while writing padding data for header: %s", err.Error())
	}

	// Make hasher
	hash := fnv.New32()

	// Make conflict map
	hashmap := map[uint32]string{}

	// Start embedding files
	offset := headersize
	for index, res := range files {
		// Open file
		file, err := os.Open(res.File)
		if err != nil {
			return fmt.Errorf("Error while opening %s: %s", res.File, err.Error())
		}

		// Get metadata
		info, err := file.Stat()
		if err != nil {
			return fmt.Errorf("Error while getting metadata of %s: %s", res.File, err.Error())
		}

		// Get file length and padded length (to 4 byte boundaries)
		length := uint32(info.Size())
		paddedlen := 4 * ((length-1)/4 + 1)

		// Write file
		_, err = io.CopyN(p.handle, file, int64(length))
		if err != nil {
			return fmt.Errorf("Error while copying data from %s: %s", res.File, err.Error())
		}

		// Write padding
		_, err = p.handle.Write(empty[:paddedlen-length])
		if err != nil {
			return fmt.Errorf("Error while copying padding after %s: %s", res.File, err.Error())
		}

		// Hash identifier
		hash.Reset()
		_, err = fmt.Fprint(hash, res.Identifier)
		if err != nil {
			return fmt.Errorf("Error while hashing file identifier %s: %s", res.Identifier, err.Error())
		}
		filehash := hash.Sum32()

		// Check for conflicts
		if !p.options.IgnoreConflicts {
			orig, ok := hashmap[filehash]
			if ok {
				return fmt.Errorf("Hash conflict detected between the following files:\n    [%8x] %s\n    [%8x] %s", filehash, orig, filehash, res.Identifier)
			}
			hashmap[filehash] = res.Identifier
		}

		// Make header
		header[index] = FileHeader{
			Hash:   filehash,
			Offset: offset,
			Length: length,
		}

		// Set new offset
		offset += paddedlen
	}

	// Write header
	_, err = p.handle.Seek(0, io.SeekStart)
	if err != nil {
		return fmt.Errorf("Error while seeking to top: %s", err.Error())
	}

	_, err = p.handle.Write(header.Bytes())
	if err != nil {
		return fmt.Errorf("Error while writing GCR header: %s", err.Error())
	}

	return nil
}
