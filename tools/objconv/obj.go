package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// OBJSettings contains constraints on OBJ parsing
type OBJSettings struct {
	AllowNgons      bool // Allow faces with more than 3 vertices
	PartialFaces    bool // Allow faces without texcoord or normal indices
	MultipleObjects bool // Allow multiple objects
	Optimize        bool // Try to perform some space-saving optimizations
}

// Coord represents a 3D (+1) coordinate for OBJ structures
type Coord struct {
	X, Y, Z, W float32
}

// UV represents a 2D (+1) coordinate for OBJ structures
type UV struct {
	U, V, W float32
}

// Face represents a single face
type Face []VertexCombo

// VertexCombo represents a vertex index, a texture coordinate and a vertex normal indices
type VertexCombo struct {
	Vertex, TexCoord, Normal uint16
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
func ParseOBJ(in io.Reader, settings OBJSettings) (Mesh, error) {
	//TODO MTL support??

	// Current object and mesh
	var currentObject Object
	var mesh Mesh

	// Optmization structures (deduplication)
	remapVertex := make(map[uint16]uint16)
	remapUV := make(map[uint16]uint16)
	remapNormal := make(map[uint16]uint16)

	// Start reading from file
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
		if len(line) < 1 || line[0] == '#' {
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
				return nil, fmt.Errorf("Error on line %d: %s", linenum, err.Error())
			}

			// Try to deduplicate
			if settings.Optimize {
				found := false
				for oindex, ocoord := range currentObject.Vertices {
					if coord == ocoord {
						remapVertex[uint16(len(currentObject.Vertices))] = uint16(oindex)
						found = true
					}
				}
				if !found {
					currentObject.Vertices = append(currentObject.Vertices, coord)
				}
			} else {
				// Just add
				currentObject.Vertices = append(currentObject.Vertices, coord)
			}

		// UV coordinate
		case "vt":
			uv, err := parseUV(rest)
			if err != nil {
				return nil, fmt.Errorf("Error on line %d: %s", linenum, err.Error())
			}

			// Try to deduplicate
			if settings.Optimize {
				found := false
				for oindex, ouv := range currentObject.TextureCoords {
					if uv == ouv {
						remapUV[uint16(len(currentObject.TextureCoords))] = uint16(oindex)
						found = true
					}
				}
				if !found {
					currentObject.TextureCoords = append(currentObject.TextureCoords, uv)
				}
			} else {
				// Just add
				currentObject.TextureCoords = append(currentObject.TextureCoords, uv)
			}

		// Vertex normal
		case "vn":
			coord, err := parseCoord(rest)
			if err != nil {
				return nil, fmt.Errorf("Error on line %d: %s", linenum, err.Error())
			}

			// Try to deduplicate
			if settings.Optimize {
				found := false
				for oindex, ocoord := range currentObject.VertexNormals {
					if coord == ocoord {
						remapNormal[uint16(len(currentObject.VertexNormals))] = uint16(oindex)
						found = true
					}
				}
				if !found {
					currentObject.VertexNormals = append(currentObject.VertexNormals, coord)
				}
			} else {
				// Just add
				currentObject.VertexNormals = append(currentObject.VertexNormals, coord)
			}

		// Face
		case "f":
			face, err := parseFace(rest, settings.PartialFaces)
			if err != nil {
				return nil, fmt.Errorf("Error on line %d: %s", linenum, err.Error())
			}
			// Check for ngons
			if len(face) > 3 && !settings.AllowNgons {
				return nil, fmt.Errorf("Face on line %d is an Ngon (%d vertices), and that's not allowed (needs -allowngons)", linenum, len(line))
			}

			// Handle remapping
			if settings.Optimize {
				for vcomboid := range face {
					// Check for deduplicated normal
					remappednormal, ok := remapNormal[face[vcomboid].Normal]
					if ok {
						face[vcomboid].Normal = remappednormal
					}
					// Check for deduplicated texture coordinate
					remappeduv, ok := remapUV[face[vcomboid].TexCoord]
					if ok {
						face[vcomboid].TexCoord = remappeduv
					}
				}
			}

			currentObject.Faces = append(currentObject.Faces, face)
		// Ignore the following until we support them properly
		case "g", "usemtl", "mtllib", "s", "vp":
			// nothing
		default:
			// Unknown stuff
			err = fmt.Errorf("Weird line (%d): %s", linenum, line)
			return nil, err
		}
	}

	// Add current object to mesh
	mesh = append(mesh, currentObject)

	// Check for multiple objects, if forbidden
	if len(mesh) > 1 && !settings.MultipleObjects {
		return nil, fmt.Errorf("This file contains %d objects, but only one is allowed (needs -allowmultiple)", len(mesh))
	}

	return mesh, nil
}

// Valid checks wether the object is valid or not
func (o Object) Valid() bool {
	return len(o.Vertices) > 0 && len(o.Faces) > 0
}

func parseCoord(line string) (c Coord, err error) {
	// On XYZW, W defaults to 1
	c.W = 1

	// Read as much as possible
	_, err = fmt.Sscan(line, &c.X, &c.Y, &c.Z, &c.W)
	// Ignore EOF
	if err == io.EOF {
		err = nil
	}
	return
}

func parseUV(line string) (u UV, err error) {
	// Read as much as possible
	_, err = fmt.Sscan(line, &u.U, &u.V, &u.W)
	// Ignore EOF
	if err == io.EOF {
		err = nil
	}
	return
}

func parseFace(line string, allowPartial bool) (f Face, err error) {
	var n int
	combos := strings.Fields(line)
	for idx, combo := range combos {
		vcombo := VertexCombo{}
		combo := strings.Replace(combo, "/", " ", -1)
		n, err = fmt.Sscan(combo, &vcombo.Vertex, &vcombo.TexCoord, &vcombo.Normal)
		if err != nil && err != io.EOF {
			return
		}
		// Check for partial match
		if n < 3 {
			// Partial matches not allowed? return error
			if !allowPartial {
				err = fmt.Errorf("Block #%d has a partial face index (%s) and that's not allowed (needs -allowpartialfaces)", idx, combos[idx])
				return
			}
			// Swap texcoord with normal if the format was "a//b"
			if strings.Index(combos[idx], "//") > 0 {
				vcombo.Normal, vcombo.TexCoord = vcombo.TexCoord, vcombo.Normal
			}
		}
		f = append(f, vcombo)
	}

	return
}
