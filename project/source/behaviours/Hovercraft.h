#pragma once
#include <entityx/entityx.h>

namespace ex = entityx;

namespace Behaviours {
struct Hovercraft {

};

class HovercraftSystem : public ex::System<HovercraftSystem> {
    void update(ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt) override;
};
};