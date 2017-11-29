#include "PhysicsSystem.h"

#include "../components/Transform.h"
#include "../components/Rigidbody.h"
#include "../components/MeshCollider.h"
#include "../math/Math.h"
#include "../math/Vector.h"

namespace cp = Components;

const Vector gravity = {0,-9.8f, 0};
const int steps = 4;
const float substep = 1.0f / steps;

void PhysicsSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Rigidbody>([&](ex::Entity entity, cp::Transform& transform, cp::Rigidbody& body) {
		Vector position = transform.position;
		Vector velocity = body.velocity;
		for (int i=0; i < 4; ++i) {
			// Gravity
			//velocity = velocity + (gravity * dt * substep);

			//Drag
			velocity = velocity + ((velocity * -1.0f) * 0.4f * dt * substep);
			position = step(es, events, position, velocity * (dt * substep));
		}

		transform.position = position;
		body.velocity = velocity;
	});
}

Vector PhysicsSystem::step(ex::EntityManager& es, ex::EventManager& events, Vector& position, const Vector& delta) {
	// STEP
	/*
		test = origin + delta
		test = test * inverse model matrix

		Do wall check
		Do floor check
		Do ceiling check

		result = test * model matrix
	*/
	return position + delta;
}