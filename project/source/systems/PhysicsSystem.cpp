#include "PhysicsSystem.h"

#include "../components/Transform.h"
#include "../components/Rigidbody.h"
#include "../components/MeshCollider.h"
#include "../math/Math.h"

namespace cp = Components;

void PhysicsSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Rigidbody>([&](ex::Entity entity, cp::Transform& transform, cp::Rigidbody& body) {

	});

	es.each<cp::Transform, cp::MeshCollider>([&](ex::Entity entity, cp::Transform& transform, cp::MeshCollider& collider) {

	});
}