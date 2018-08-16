package main

// Preprocess does the necessary processing to the file to output a valid BMB
func Preprocess(meshdata *Mesh) {
	// Re-order faces to put walls at the end (for collisions)
	for meshid := range meshdata.Objects {
		faces := []Face{}
		walls := []Face{}
		for _, face := range meshdata.Objects[meshid].Faces {
			// Get normal
			normal := meshdata.VertexNormals[face[0].Normal]

			// Check if it's a wall
			if normal.Y >= -0.01 && normal.Y <= 0.01 {
				walls = append(walls, face)
			} else {
				faces = append(faces, face)
			}
		}
		// Add walls at the end of the faces and set as new set
		faces = append(faces, walls...)
		meshdata.Objects[meshid].Faces = faces
	}
}
