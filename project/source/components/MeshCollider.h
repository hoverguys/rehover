#pragma once
#include "../rendering/Mesh.h"

namespace Components {

/*! \brief Mesh collider component
 *  Gives the entity a collision mesh for rigidbodies to interact with
 */
struct MeshCollider {
	/*! \brief Initialize mesh collider
	 *  \param mesh Mesh to use as collider
	 */
	explicit MeshCollider(const std::shared_ptr<Mesh>& mesh) : mesh(mesh) {}

	/*! Collision mesh */
	std::shared_ptr<Mesh> mesh;
};
} // namespace Components