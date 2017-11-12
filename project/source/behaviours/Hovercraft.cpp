#include "Hovercraft.h"
#include "../components/Transform.h"

namespace cp = Components;

namespace Behaviours {
void Hovercraft::Tick(ex::Entity entity, ex::TimeDelta dt) {
		ex::ComponentHandle<cp::Transform> transform = entity.component<cp::Transform>();

		// Move hovercraft
		transform->position.x += controller->GetAxis(HovercraftController::Motion::Turn) * 0.02f * dt;
		transform->position.y += controller->GetAxis(HovercraftController::Motion::Pitch) * 0.02f * dt;

		// Have camera track hovercraft
		ex::ComponentHandle<cp::Transform> camera_trans = camera.component<cp::Transform>();
		camera_trans->Lookat(transform->position);
}
} // namespace Behaviours