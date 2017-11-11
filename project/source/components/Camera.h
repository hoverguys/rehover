#pragma once
#include <ogc/gu.h>

namespace Components {
struct CameraViewport {
	float width;
	float height;
	float offsetLeft;
	float offsetTop;
};

class Camera {
public:
	Camera();
	void SetViewport(float width, float height, float offsetLeft, float offsetTop);
	void SetPerspective(float fov, float nearPlane, float farPlane);

	float fieldOfView;
	float nearPlane;
	float farPlane;

	CameraViewport viewport;
	Mtx44 perspectiveMtx;
};
}; // namespace Components