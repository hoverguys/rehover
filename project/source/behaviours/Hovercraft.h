#pragma once
#include "../input/HovercraftController.h"
#include <entityx/entityx.h>

namespace ex = entityx;

namespace Behaviours {
struct Hovercraft {
	std::shared_ptr<HovercraftController> controller;
	ex::Entity camera;

	void Tick(ex::Entity entity, ex::TimeDelta dt);
};
} // namespace Behaviours

