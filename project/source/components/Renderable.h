#pragma once
#include "../rendering/Material.h"
#include "../rendering/Mesh.h"


namespace Components {
/*! \brief Renderable component
 *  Gives the entity a mesh and optional material to be rendered by the render system
 *  NOTE: A renderable by itself is not enough to get rendered, a Transform component is also required
 */
struct Renderable {
	/*! \brief Create a renderable using just a mesh
	 *  \param mesh Mesh to render the entity with
	 */
	explicit Renderable(const std::shared_ptr<Mesh>& mesh) : Renderable(mesh, nullptr) {}

	/*! \brief Create a renderable using a mesh and material
	 *  \param mesh Mesh to render the entity with
	 *  \param material Material (texture+shader) to use to render the mesh
	 */
	explicit Renderable(const std::shared_ptr<Mesh>& mesh, const std::shared_ptr<Material>& material)
		: mesh(mesh), material(material) {}

	/*! Mesh */
	std::shared_ptr<Mesh> mesh;

	/*! Material */
	std::shared_ptr<Material> material;
};
} // namespace Components