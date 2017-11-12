#pragma once
#include "rendering/Mesh.h"
#include "systems/InputSystem.h"
#include <entityx/entityx.h>
#include <memory>

namespace ex = entityx;

class Game : public ex::EntityX {
public:
	explicit Game();
	void init();
	void update(ex::TimeDelta dt);
};