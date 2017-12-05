#pragma once
#include "../pchheader.h"

namespace ex = entityx;

class UISystem : public ex::System<UISystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;

private:
	/*! Setup rendering for 2D sprites */
	void Setup2D();
};