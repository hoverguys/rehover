#include "PhysicsSystem.h"

#include "../components/MeshCollider.h"
#include "../components/Rigidbody.h"
#include "../components/Transform.h"
#include "../math/Math.h"
#include "../math/Vector.h"

namespace cp = Components;

const Vector gravity = {0, -9.8f, 0};
const int steps = 4;
const float substep = 1.0f / steps;

void handleFloors(const Mesh& mesh, PhysicsStep& step);
void handleWalls(const Mesh& mesh, PhysicsStep& step);

void PhysicsSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Rigidbody>([&](ex::Entity entity, cp::Transform& transform, cp::Rigidbody& body) {

		PhysicsStep stepData = {transform.position, body.velocity};

		for (int i = 0; i < 4; ++i) {
			// Gravity
			stepData.velocity = stepData.velocity + (gravity * dt * substep);

			// Drag
			stepData.velocity = stepData.velocity + ((stepData.velocity * -1.0f) * 0.4f * dt * substep);
			step(es, events, stepData, stepData.velocity * (dt * substep));
		}

		// Sync transform with body
		transform.position = stepData.position;
		body.velocity = stepData.velocity;
	});
}

void PhysicsSystem::step(ex::EntityManager& es, ex::EventManager& events, PhysicsStep& step, const Vector& delta) {
	// STEP
	// Apply delta
	step.position = step.position + delta;

	// Correct position through collision
	es.each<cp::Transform, cp::MeshCollider>(
		[&](ex::Entity entity, cp::Transform& transform, cp::MeshCollider& collider) {
			const Matrix& modelMtx = transform.GetMatrix();
			const Matrix& inversedMtx = modelMtx.Inversed();

			// Move player into model space
			step.position = inversedMtx.Multiply(step.position);

			const Mesh& mesh = *collider.mesh;

			handleFloors(mesh, step);
			handleWalls(mesh, step);

			// Move player into world space
			step.position = modelMtx.Multiply(step.position);
		});
}

void handleFloors(const Mesh& mesh, PhysicsStep& step) {
	for (int f = 0; f < mesh.faceCount; ++f) {
		// Get face indices
		const int faceOffset = f * 3;
		const MeshIndex& i0 = mesh.indexArray[faceOffset + 0];
		const MeshIndex& i1 = mesh.indexArray[faceOffset + 1];
		const MeshIndex& i2 = mesh.indexArray[faceOffset + 2];

		// Get points and normal from face
		const Vector& normal = mesh.normalArray[i0.normal];

		// Skip if not a floor
		if (normal.y <= 0.01f) {
			continue;
		}

		const Vector& v0 = mesh.vertexArray[i0.vertex];
		const Vector& v1 = mesh.vertexArray[i1.vertex];
		const Vector& v2 = mesh.vertexArray[i2.vertex];

		const float alpha = 0.5f * (-v1.z * v2.x + v0.z * (-v1.x + v2.x) + v0.x * (v1.z - v2.z) + v1.x * v2.z);
		const float sign = alpha < 0.0f ? -1.0f : 1.0f;
		const float s =
			(v0.z * v2.x - v0.x * v2.z + (v2.z - v0.z) * step.position.x + (v0.x - v2.x) * step.position.z) * sign;
		const float t =
			(v0.x * v1.z - v0.z * v1.x + (v0.z - v1.z) * step.position.x + (v1.x - v0.x) * step.position.z) * sign;

		if (s < 0 || t < 0 || (s + t) >= 2.0f * alpha * sign) {
			continue;
		}

		// Skip if we are above the hitbox
		const Vector deltaTop = step.position - (v0 + Math::worldUp * 0.0f);
		if (normal.Dot(deltaTop) > 0.0f) {
			continue;
		}

		// Skip if we are below the hitbox
		const Vector deltaBottom = step.position - (v0 + Math::worldUp * -1.0f);
		if (normal.Dot(deltaBottom) < 0.0f) {
			continue;
		}

		// In triangle, snap
		const float d = normal.Dot(Math::worldUp);
		const float offset = (v0 - step.position).Dot(normal) / d;
		step.position.y += offset;

		// Not sure about this thing
		step.velocity.y = 0;
	}
}

void handleWalls(const Mesh& mesh, PhysicsStep& step) {
	const float wallFront = 0.0f;
	const float wallBack = -1.0f;

	for (int faceID = 0; faceID < mesh.faceCount; ++faceID) {
		// Get face indices
		const int faceOffset = faceID * 3;
		const MeshIndex& i0 = mesh.indexArray[faceOffset + 0];
		const MeshIndex& i1 = mesh.indexArray[faceOffset + 1];
		const MeshIndex& i2 = mesh.indexArray[faceOffset + 2];

		// Get points and normal from face
		const Vector normal = mesh.normalArray[i0.normal];

		// Skip if not a wall
		if (normal.y <= -0.01f || normal.y >= 0.01f) {
			continue;
		}

		const Vector skewnormal = (normal * Vector(1.0f, 0.0f, 1.0f)).Normalized();
		const Vector& v0 = mesh.vertexArray[i0.vertex];
		const Vector& v1 = mesh.vertexArray[i1.vertex];
		const Vector& v2 = mesh.vertexArray[i2.vertex];

		// Check if point is in triangle
		// Compute vectors
		const Vector edge1 = v1 - v0;
		const Vector edge2 = v2 - v0;
		const Vector rayDirection = skewnormal * -1.0f;
		const Vector h = rayDirection.Cross(edge2);
		const float f = 1.0f / edge1.Dot(h);
		const Vector s = step.position - v0;
		const float u = f * s.Dot(h);

		if (u < 0.0f || u > 1.0f) {
			continue;
		}

		const Vector q = s.Cross(edge1);
		const float v = f * rayDirection.Dot(q);
		if (v < 0.0f || u + v > 1.0f) {
			continue;
		}

		// Cool we intersected! lets see if we are in the hitbox
		float t = f * edge2.Dot(q);
		if (t > wallFront || t < wallBack) {
			continue;
		}

		// We hit bricks
		step.position = step.position - (skewnormal * t);
		step.velocity = step.velocity - (skewnormal * skewnormal.Dot(step.velocity));
	}
}