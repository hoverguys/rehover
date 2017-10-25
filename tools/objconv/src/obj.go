package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Coord represents a 3D (+1) coordinate for OBJ structures
type Coord struct {
	X, Y, Z, W float32
}

// UV represents a 2D (+1) coordinate for OBJ structures
type UV struct {
	U, V, W float32
}

// Face represents a single face
type Face struct {
	//TODO
}

// Object is an OBJ Object
type Object struct {
	Name          string
	Vertices      []Coord
	TextureCoords []UV
	VertexNormals []Coord
	Faces         []Face
}

// Mesh is a collection of objects from an OBJ file
type Mesh []Object

// ParseOBJ parses an OBJ file from a reader and returns a mesh and an optional error, if any
func ParseOBJ(in io.Reader) (Mesh, error) {
	//TODO MTL support??
	var currentObject Object
	var mesh Mesh
	reader := bufio.NewReader(in)
	linenum := 0
	for {
		linenum++

		// Get next line from reader
		line, err := reader.ReadString('\n')
		if err != nil {
			//TODO Should probably check for other errors and not assume EOF
			break
		}

		// Trim extra whitespace (like windows' \r)
		line = strings.TrimSpace(line)

		// Skip empty lines and comments
		if len(line) < 0 || line[0] == '#' {
			continue
		}

		// Get first atom
		space := strings.IndexRune(line, ' ')
		if space < 1 {
			// No space?!
			err = fmt.Errorf("Weird line (%d): %s", linenum, line)
			return nil, err
		}

		ftype := line[:space]
		rest := line[space+1:] // NOTE: line is guaranteed to be longer than 'space' or that whitespace would have been trimmed

		switch ftype {
		// Object (and name)
		case "o":
			// Only add the current object if valid (ie. not the empty one before the first one)
			if currentObject.Valid() {
				mesh = append(mesh, currentObject)
			}
			currentObject = Object{
				Name: rest,
			}
		// Vertex
		case "v":
			coord, err := parseCoord(rest)
			if err != nil {
				return nil, err
			}
			currentObject.Vertices = append(currentObject.Vertices)
		// UV coordinate
		case "vt":
		// Vertex normal
		case "vn":
			coord, err := parseCoord(rest)
			if err != nil {
				return nil, err
			}
			currentObject.VertexNormals = append(currentObject.Vertices)
		// Face
		case "f":
			//TODO
		// Ignore the following until we support them properly
		case "g", "usemtl", "mtllib", "s", "vp":
			// nothing
		default:
			// Unknown stuff
			err = fmt.Errorf("Weird line (%d): %s", linenum, line)
			return nil, err
		}
	}
	return mesh, err
}

// Valid checks wether the object is valid or not
func (o Object) Valid() bool {
	return o.Name != "" && len(o.Vertices) > 0 && len(o.Faces) > 0
}

func parseCoord(line string) (c Coord, err error) {
	// On XYZW, W defaults to 1
	c.W = 1

	// Read as much as possible
	_, err = fmt.Sscan(line, &c.X, &c.Y, &c.Z, &c.W)
	return
}

func parseUV(line string) (u UV, err error) {
	// Read as much as possible
	_, err = fmt.Sscan(line, &u.U, &u.V, &u.W)
	return
}
