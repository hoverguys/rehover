#pragma once

#include "Resource.h"
#include "../rendering/Mesh.h"

struct MeshResourceHeader {
	unsigned int vcount;  /*< Vertex count        */
	unsigned int ncount;  /*< Normal count        */
	unsigned int vtcount; /*< UV Coordinate count */
	unsigned int fcount;  /*< Face/Index count    */
};

class MeshResource : public Resource {
public:
	Mesh* Load();
protected:
	void Initialize() override;
private:
    MeshResourceHeader* header;
	Mesh* internal;
	bool loaded;
};
