#include "Hovercraft.h"
#include "../components/Transform.h"
#include "../math/Math.h"

namespace cp = Components;

namespace Behaviours {
void Hovercraft::Tick(ex::Entity entity, ex::TimeDelta dt) {
	ex::ComponentHandle<cp::Transform> transform = entity.component<cp::Transform>();

	// Rotate hovercraft
	float deltaRotation = controller->GetAxis(HovercraftController::Motion::Turn) * 1.2f * dt;
	transform->RotateAxisAngle(Math::worldUp, deltaRotation);

	// Forward
	Vector throttle = transform->forward * (controller->GetAxis(HovercraftController::Motion::Throttle) * 5.2f * dt);
	transform->position = transform->position + throttle;

	// Have camera track hovercraft
	ex::ComponentHandle<cp::Transform> camera_trans = camera.component<cp::Transform>();

	const float targetHeight = 1.6f;
	const float cameraHeight = 2.0f;
	const float cameraDistance = -5.0f;
	const float t = 1.f / 5.f;

	/* Calculate camera position */
	const Vector targetCameraPos = transform->position + (transform->forward * cameraDistance) + Vector(0, cameraHeight, 0);

	/* Calculate camera target */
	const Vector targetPos = transform->position + (transform->up * targetHeight);

	/* Lerp between old camera position and target */
	camera_trans->position = camera_trans->position + (targetCameraPos - camera_trans->position) * t;
	camera_trans->Lookat(targetPos);
}

} // namespace Behaviours