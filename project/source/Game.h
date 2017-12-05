#pragma once
#include "pchheader.h"

#include "rendering/Mesh.h"
#include "systems/InputSystem.h"

namespace ex = entityx;

class Game : public ex::EntityX {
public:
	explicit Game();
	void init();
	void update(ex::TimeDelta dt);
};