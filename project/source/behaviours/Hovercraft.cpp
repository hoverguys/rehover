#include "Hovercraft.h"
#include "../components/Transform.h"

namespace cp = Components;

namespace Behaviours {
void HovercraftSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<Hovercraft>([&](ex::Entity entity, Hovercraft& hovercraft) {
		ex::ComponentHandle<cp::Transform> transform = entity.component<cp::Transform>();

        // Move hovercraft
		transform->position.x += hovercraft.controller->GetAxis(HovercraftController::MotionTurn) * 0.02f * dt;
        transform->position.y += hovercraft.controller->GetAxis(HovercraftController::MotionPitch) * 0.02f * dt;

        // Have camera track hovercraft
        ex::ComponentHandle<cp::Transform> camera_trans = hovercraft.camera.component<cp::Transform>();
        camera_trans->Lookat(transform->position);
	});
};
}; // namespace Behaviours