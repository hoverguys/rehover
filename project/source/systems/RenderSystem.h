#pragma once
#include <entityx/entityx.h>
#include "../components/Camera.h"

namespace ex = entityx;

class RenderSystem : public ex::System<RenderSystem> {
    void update(ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt) override;
    void RenderScene(Mtx& cameraMtx, ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt);
    static void SetupCamera(Components::Camera& camera);
};