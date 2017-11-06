#pragma once
#include <entityx/entityx.h>

namespace ex = entityx;

class RenderSystem : public ex::System<RenderSystem> {
    void update(ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt) override;
};