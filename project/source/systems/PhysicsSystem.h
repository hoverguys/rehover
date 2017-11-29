#pragma once
#include <entityx/entityx.h>

class Vector;

namespace ex = entityx;

class PhysicsSystem : public ex::System<PhysicsSystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
	Vector step(ex::EntityManager& es, ex::EventManager& events, Vector origin, const Vector& velocity);
};