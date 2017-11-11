#pragma once

#include "../rendering/Mesh.h"
#include "Resource.h"

struct MeshResourceHeader {
	unsigned short vcount;  /*< Vertex count        */
	unsigned short ncount;  /*< Normal count        */
	unsigned short vtcount; /*< UV Coordinate count */
	unsigned short fcount;  /*< Face/Index count    */
};

class MeshResource : public Resource {
public:
	MeshResource(void* base, unsigned int size) : Resource(base, size) {}
	Mesh* Load();
	void Initialize() override;

private:
	MeshResourceHeader* header;
	Mesh* internal;
	bool loaded;
};
