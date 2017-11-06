#include "RenderSystem.h"
#include "Mesh.h"

void RenderSystem::update(ex::EntityManager &es, ex::EventManager &events, ex::TimeDelta dt) {
    const float fdt = static_cast<float>(dt);
    es.each<Mesh*>([fdt](ex::Entity entity, Mesh*& m) {
        m->Render();
    });
};
