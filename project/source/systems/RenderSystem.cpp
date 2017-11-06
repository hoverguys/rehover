#include "RenderSystem.h"

#include "../components/Renderable.h"

void RenderSystem::update(ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt) {
    const float fdt = static_cast<float>(dt);
    es.each<Components::Renderable>([fdt](ex::Entity entity, Components::Renderable& renderable) {
        renderable.mesh->Render();
    });
};
