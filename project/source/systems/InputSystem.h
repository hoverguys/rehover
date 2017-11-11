#pragma once
#include "../input/GCController.h"
#include <entityx/entityx.h>
#include <map>

namespace ex = entityx;

class InputSystem : public ex::System<InputSystem> {
	/*! All connected pads, as bitset */
	unsigned long gcConnectedPads;

	/*! All gamecube pads instances */
	std::map<unsigned short, GCController> gcControllers;

	// EntityX methods
	void configure(ex::EntityManager& entities, ex::EventManager& events) override;
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
};