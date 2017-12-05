#pragma once
#include "../pchheader.h"

#include "../math/Vector.h"

namespace ex = entityx;

struct PhysicsStep {
	Vector position;
	Vector velocity;
};

class PhysicsSystem : public ex::System<PhysicsSystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
	void step(ex::EntityManager& es, ex::EventManager& events, PhysicsStep& body, const Vector& delta);
};