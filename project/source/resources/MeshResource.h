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
	std::shared_ptr<Mesh> Load();
	void Initialize() override;

private:
	MeshResourceHeader* header;
	std::shared_ptr<Mesh> internal;
	bool loaded;
};
