#pragma once
#include "../rendering/Mesh.h"
#include "../rendering/Material.h"

namespace Components {
struct Renderable {
	Renderable(const std::shared_ptr<Mesh>& mesh) : Renderable(mesh, nullptr) {}
	Renderable(const std::shared_ptr<Mesh>& mesh, const std::shared_ptr<Material>& material) : mesh(mesh), material(material) {}

	std::shared_ptr<Mesh> mesh;
	std::shared_ptr<Material> material;
};
} // namespace Components