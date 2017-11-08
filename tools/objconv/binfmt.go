package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

// Endianess represents a byte order
type Endianess string

// All supported byte orders
const (
	EndianessLittle Endianess = "little" // Little endian (x86)
	EndianessBig    Endianess = "big"    // Big endian (PPC)
)

const MaxUint16 = ^uint16(0)

// SaveModel takes a parsed model and writes it in binary format using a provided byte order
func SaveModel(mesh Mesh, out io.Writer, boType Endianess) error {
	var endianess binary.ByteOrder
	switch boType {
	case EndianessLittle:
		endianess = binary.LittleEndian
	case EndianessBig:
		endianess = binary.BigEndian
	default:
		return fmt.Errorf("Unknown endianess: %s (supported: big, little)", boType)
	}

	//TODO Support multiple objects?
	object := mesh.Objects[0]

	// Chech that the object can fit in a BMB file
	if len(mesh.Vertices) >= int(MaxUint16) {
		return fmt.Errorf("Model has too many vertex positions (has %d, max is %d)", len(mesh.Vertices), MaxUint16)
	}
	if len(mesh.VertexNormals) >= int(MaxUint16) {
		return fmt.Errorf("Model has too many vertex normals (has %d, max is %d)", len(mesh.VertexNormals), MaxUint16)
	}
	if len(mesh.TextureCoords) >= int(MaxUint16) {
		return fmt.Errorf("Model has too many texture coordinates (has %d, max is %d)", len(mesh.TextureCoords), MaxUint16)
	}
	if len(object.Faces) >= int(MaxUint16) {
		return fmt.Errorf("Model has too many faces (has %d, max is %d)", len(object.Faces), MaxUint16)
	}

	// Write header
	binary.Write(out, endianess, uint16(len(mesh.Vertices)))
	binary.Write(out, endianess, uint16(len(mesh.VertexNormals)))
	binary.Write(out, endianess, uint16(len(mesh.TextureCoords)))
	binary.Write(out, endianess, uint16(len(object.Faces)))

	// Write vertices
	for _, vertex := range mesh.Vertices {
		binary.Write(out, endianess, vertex.X)
		binary.Write(out, endianess, vertex.Y)
		binary.Write(out, endianess, vertex.Z)
	}

	// Write normals
	for _, normals := range mesh.VertexNormals {
		binary.Write(out, endianess, normals.X)
		binary.Write(out, endianess, normals.Y)
		binary.Write(out, endianess, normals.Z)
	}

	// Write texture Coordinates
	for _, uv := range mesh.TextureCoords {
		binary.Write(out, endianess, uv.U)
		binary.Write(out, endianess, 1.0-uv.V)
	}

	// Write faces
	for _, face := range object.Faces {
		for _, vcombo := range face {
			binary.Write(out, endianess, vcombo.Vertex)
			binary.Write(out, endianess, vcombo.TexCoord)
			binary.Write(out, endianess, vcombo.Normal)
		}
	}

	return nil
}
