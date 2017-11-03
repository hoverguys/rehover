#pragma once
#include "../rendering/Mesh.h"

struct MeshResourceHeader {
	unsigned int vcount;  /*< Vertex count        */
	unsigned int ncount;  /*< Normal count        */
	unsigned int vtcount; /*< UV Coordinate count */
	unsigned int fcount;  /*< Face/Index count    */
};

class MeshResource {
public:
	MeshResource(unsigned char* base, unsigned int size);
	Mesh* Load();

private:
    MeshResourceHeader* header;
	Mesh* internal;
	bool loaded;
};
