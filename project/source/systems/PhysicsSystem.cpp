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

Vector PhysicsSystem::step(ex::EntityManager& es, ex::EventManager& events, Vector position, const Vector& delta) {
	// STEP
	/*
		test = origin + delta
		test = test * inverse model matrix

		Do wall check
		Do floor check
		Do ceiling check

		result = test * model matrix
	*/
	es.each<cp::Transform, cp::MeshCollider>([&](ex::Entity entity, cp::Transform& transform, cp::MeshCollider& collider) {
		const Matrix& modelMtx = transform.GetMatrix();
		const Matrix& inversedMtx = modelMtx.Inversed();

		// Move player into model space
		Vector localPosition = inversedMtx.Multiply(position);
		
		const Mesh& mesh = *collider.mesh;

		int below = 0;
		int box = 0;
		int hit = 0;

		for (int f = 0; f < mesh.faceCount; ++f) {
			// Get face indices
			const MeshIndex& i0 = mesh.indexArray[0];
			const MeshIndex& i1 = mesh.indexArray[1];
			const MeshIndex& i2 = mesh.indexArray[2];

			// Get points and normal from face
			const Vector& normal = mesh.normalArray[i0.normal];
			const Vector& v0 = mesh.positionArray[i0.vertex];
			const Vector& v1 = mesh.positionArray[i1.vertex];
			const Vector& v2 = mesh.positionArray[i2.vertex];

			const Vector deltaTop = localPosition - (v0);
			if (normal.Dot(deltaTop) > 0) {
				continue;
			}

			below++;

			const Vector deltaBottom = localPosition - (v0 + Math::worldUp * -100);
			if (normal.Dot(deltaBottom) < 0) {
				continue;
			}

			box++;

			const float alpha = 0.5f * (-v1.z * v2.x + v0.z * (-v1.x + v2.x) + v0.x * (v1.z - v2.z) + v1.x * v2.z);
			const float sign = alpha < 0.0f ? -1.0f : 1.0f;
			const float s = (v0.z * v2.x - v0.x * v2.z + (v2.z - v0.z) * localPosition.x + (v0.x - v2.x) * localPosition.z) * sign;
    		const float t = (v0.x * v1.z - v0.z * v1.x + (v0.z - v1.z) * localPosition.x + (v1.x - v0.x) * localPosition.z) * sign;

			if (s > 0 && t > 0 && (s + t) < 2 * alpha * sign) {
				hit++;
			}
		}

		//printf("Detected total:%d low:%d box:%d hit:%d\n", mesh.faceCount, below, box, hit);

		// Move player into world space
		position = modelMtx.Multiply(localPosition);
	});

	return position + delta;
}