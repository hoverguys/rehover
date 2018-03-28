#pragma once

#include "../rendering/Mesh.h"
#include "Resource.h"

/*! \brief BMB file header
 */
struct MeshResourceHeader {
	unsigned short vcount;  /*< Vertex count        */
	unsigned short ncount;  /*< Normal count        */
	unsigned short vtcount; /*< UV Coordinate count */
	unsigned short fcount;  /*< Face/Index count    */
};

/*!
 * \brief Mesh resource loader
 * Loads a mesh from a BMB file in memory
 */
class MeshResource : public Resource {
public:
	/*!
	 * \brief Loads a mesh from a BMB file in memory
	 *
	 * \param base Pointer to BMB file data
	 * \param size Size of the BMB file
	 */
	MeshResource(void* base, unsigned int size) : Resource(base, size) {}

	/*!
	 * \brief Load and get the mesh
	 *
	 * \return Pointer to the loaded mesh
	 */
	std::shared_ptr<Mesh> Load();

	long int ReferenceCount() override {
		return internal == NULL ? 0 : internal.use_count();
	}

	void Initialize() override;

private:
	MeshResourceHeader* header = nullptr;
	bool loaded = false;
	std::shared_ptr<Mesh> internal;
};
