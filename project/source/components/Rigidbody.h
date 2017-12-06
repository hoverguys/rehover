#pragma once
#include "../math/Vector.h"

namespace Components {
/*! \brief Rigid body component
 *  Allows the entity to interact with colliders
 */
struct Rigidbody {
	/*! \brief Create a rigidbody with an initial velocity
	 *  \param velocity Initial velocity
	 */
	explicit Rigidbody(const Vector& velocity = {0, 0, 0}) : velocity(velocity) {}

	/*! Current velocity */
	Vector velocity;

	/*! Is the rigidbody affected by gravity? */
	bool useGravity = true;
};
} // namespace Components