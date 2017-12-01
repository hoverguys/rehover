#pragma once
#include "../math/Vector.h"

namespace Components {
struct Rigidbody {
	explicit Rigidbody(const Vector& velocity = {0, 0, 0}) : velocity(velocity) {}

	Vector position = {0, 0, 0};
	Vector velocity;
    bool useGravity = true;
};
} // namespace Components