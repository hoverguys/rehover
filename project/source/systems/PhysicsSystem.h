#pragma once
#include <entityx/entityx.h>

namespace ex = entityx;

class PhysicsSystem : public ex::System<PhysicsSystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
};