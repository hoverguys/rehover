#pragma once
#include <entityx/entityx.h>
#include "../components/Rigidbody.h"

class Vector;

namespace ex = entityx;
namespace cp = Components;

class PhysicsSystem : public ex::System<PhysicsSystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
	void step(ex::EntityManager& es, ex::EventManager& events, cp::Rigidbody& body, const Vector& delta);
};