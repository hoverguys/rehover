#pragma once
#include "../rendering/Mesh.h"

namespace Components {
struct Renderable {
	Renderable(std::shared_ptr<Mesh> mesh) : mesh(mesh) {}

	std::shared_ptr<Mesh> mesh;
};
} // namespace Components