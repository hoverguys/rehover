#pragma once
#include "rendering/Mesh.h"
#include <entityx/entityx.h>

namespace ex = entityx;

class Game : public ex::EntityX {
public:
	explicit Game();
	void init(Mesh* mesh);
	void update(ex::TimeDelta dt);

	ex::Entity hovercraft;
};