#pragma once
#include "../math/Vector.h"

struct MeshIndex {
	short unsigned int vertex;
	short unsigned int uv;
	short unsigned int normal;
};

class Mesh {
public:
	Mesh();
	~Mesh();

	void Render();

protected:
	friend class MeshResource;

	Vector* positionArray = nullptr;
	Vector* normalArray = nullptr;
	float* uvArray = nullptr;
	MeshIndex* indexArray = nullptr;

	void* displayList = nullptr;
	unsigned int displayListSize = 0;
};