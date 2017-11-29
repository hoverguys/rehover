#pragma once
#include "../rendering/Mesh.h"

namespace Components {
struct MeshCollider {
	explicit MeshCollider(const std::shared_ptr<Mesh>& mesh) : mesh(mesh) {}

	std::shared_ptr<Mesh> mesh;
};
} // namespace Components