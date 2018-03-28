#pragma once
#include "pchheader.h"

namespace ex = entityx;

class Game : public ex::EntityX {
public:
	explicit Game();
	void init();
	void update(ex::TimeDelta dt);
};