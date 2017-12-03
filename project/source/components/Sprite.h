#pragma once

#include <memory>

#include "../math/Rect.h"
#include "../rendering/Material.h"

namespace Components {
/* \brief 2D sprite */
struct Sprite {
	/*! Base size (unscaled) */
	Vector2D size;

	/*! Sprite material (texture + optional shader) */
	std::shared_ptr<Material> material;

	/*! Texture offset (for atlasing) */
	Rect bounds;

	explicit Sprite(const Vector2D& size, const std::shared_ptr<Material>& material,
					const Rect& bounds = Rect(0, 0, 1, 1))
		: size(size), material(material), bounds(bounds){};
};
} // namespace Components