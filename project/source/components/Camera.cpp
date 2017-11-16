#include "Camera.h"
#include "../rendering/Graphics.h"

namespace Components {
Camera::Camera() {
	SetViewport(0, 0, 1, 1);
	SetPerspective(60, 0.1f, 300.0f);
}

void Camera::SetViewport(float offsetLeft, float offsetTop, float width, float height) {
	GXRModeObj* rmode = Graphics::GetMode();
	viewport.width = rmode->viWidth * width;
	viewport.height = rmode->efbHeight * height;
	viewport.offsetLeft = rmode->viWidth * offsetLeft;
	viewport.offsetTop = rmode->efbHeight * offsetTop;
}

void Camera::SetPerspective(float fov, float nearPlane, float farPlane) {
	fieldOfView = fov;
	this->nearPlane = nearPlane;
	this->farPlane = farPlane;

	float aspectRatio = viewport.width / viewport.height;
	guPerspective(perspectiveMtx, fov, aspectRatio, nearPlane, farPlane);
}
} // namespace Components