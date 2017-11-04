#include "Camera.h"
#include "Graphics.h"

Camera* Camera::activeCamera = nullptr;

Camera::Camera(guVector position, guVector target) {
    this->position = position;
    this->target = target;

	guVector up = { 0, 1, 0 };
	guLookAt(cameraMtx, &position, &up, &target);

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

void Camera::SetActive() {
    GX_SetScissor(viewport.offsetLeft, viewport.offsetTop, viewport.width, viewport.height);
    GX_SetViewport(viewport.offsetLeft, viewport.offsetTop, viewport.width, viewport.height, 0, 1);
    GX_LoadProjectionMtx(perspectiveMtx, GX_PERSPECTIVE);

    activeCamera = this;
}

Mtx* Camera::GetActiveMtx() {
    return &activeCamera->cameraMtx;
}