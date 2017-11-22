#include "Hovercraft.h"
#include "../components/Transform.h"
#include "../utils/math.h"

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
	Vector posTemp;
	Vector targetCameraPos = {0, cameraHeight, 0};
	
	//guVecScale(&transform->forward, &posTemp, cameraDistance);
	//guVecAdd(&targetCameraPos, &posTemp, &targetCameraPos);
	//guVecAdd(&transform->position, &targetCameraPos, &targetCameraPos);
	posTemp = transform->forward * cameraDistance;
	targetCameraPos = targetCameraPos + posTemp;
	targetCameraPos = targetCameraPos + transform->position;

	/* Calculate camera target */
	Vector targetPos;
	//guVecScale(&transform->up, &targetPos, targetHeight);
	//guVecAdd(&targetPos, &transform->position, &targetPos);
	targetPos = transform->up * targetHeight;
	targetPos = targetPos + transform->position;

	/* Lerp between old camera position and target */
	Vector camPos;
	//guVecSub(&targetCameraPos, &camera_trans->position, &camPos);
	//guVecScale(&camPos, &camPos, t);
	//guVecAdd(&camera_trans->position, &camPos, &camera_trans->position);
	camPos = targetCameraPos - camera_trans->position;
	camPos = camPos * t;
	camera_trans->position = camera_trans->position + camPos;

	camera_trans->Lookat(targetPos);
}

} // namespace Behaviours