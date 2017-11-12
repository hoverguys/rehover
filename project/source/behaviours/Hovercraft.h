#pragma once
#include "../input/HovercraftController.h"
#include <entityx/entityx.h>

namespace ex = entityx;

namespace Behaviours {
struct Hovercraft {
	std::shared_ptr<HovercraftController> controller;
};

class HovercraftSystem : public ex::System<HovercraftSystem> {
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
};
}; // namespace Behaviours