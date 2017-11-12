#include "Hovercraft.h"
#include "../components/Transform.h"

namespace cp = Components;

namespace Behaviours {
void HovercraftSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<Hovercraft>([&](ex::Entity entity, Hovercraft& hovercraft) {
		ex::ComponentHandle<cp::Transform> position = entity.component<cp::Transform>();
		position->position.x += hovercraft.controller->GetAxis(HovercraftController::MotionTurn) * 0.02f * dt;
	});
};
}; // namespace Behaviours