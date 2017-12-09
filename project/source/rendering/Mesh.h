#pragma once
#include "../math/Vector.h"

/*! \brief Index entry for each vertex in a face
 */
struct MeshIndex {
	const short unsigned int vertex; /*< Vertex position index */
	const short unsigned int uv;     /*< Texture coordinate index */
	const short unsigned int normal; /*< Vertex normal index */
};

/*!
 * \brief Mesh resource
 * A mesh that can be used in a Renderable and contains all its data
 */
struct Mesh {
	/*!
	 * \brief Create a new mesh from loaded data
	 *
	 * \param vertices All vertex positions
	 * \param normals All vertex normals
	 * \param uvs All texture coordinates
	 * \param indices All indices used for faces (every 3 indices are a triangle)
	 * \param facecount Number of triangles in the mesh
	 */
	explicit Mesh(const Vector* vertices, const Vector* normals, const float* uvs, const MeshIndex* indices,
				  const unsigned short int facecount)
		: vertexArray(vertices), normalArray(normals), uvArray(uvs), indexArray(indices), faceCount(facecount) {}
	~Mesh();

	/*!
	 * \brief Call the display list that renders the model on screen
	 */
	void Render();

	/*! List of vertex positions */
	const Vector* vertexArray;

	/*! List of vertex normals */
	const Vector* normalArray;

	/*! List of texture coordinates */
	const float* uvArray;

	/*! List of contiguous indices */
	const MeshIndex* indexArray;

	/*! Number of triangles */
	const unsigned short int faceCount;

	/*! Display list that renders the model on screen */
	void* displayList = nullptr;

	/*! Display list size */
	unsigned int displayListSize = 0;
};