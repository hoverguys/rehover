#pragma once
#include "../math/Vector.h"

struct MeshIndex {
	const short unsigned int vertex;
	const short unsigned int uv;
	const short unsigned int normal;
};

struct Mesh {
	explicit Mesh(const Vector* vertices, const Vector* normals, const float* uvs, const MeshIndex* indices, const unsigned short int facecount) :
		vertexArray(vertices), normalArray(normals), uvArray(uvs), indexArray(indices), faceCount(facecount) {}
	~Mesh();

	void Render();

	const Vector* vertexArray;
	const Vector* normalArray;
	const float* uvArray;
	const MeshIndex* indexArray;
	const unsigned short int faceCount;

	void* displayList = nullptr;
	unsigned int displayListSize = 0;
};