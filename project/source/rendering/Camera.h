#pragma once
#include <gccore.h>

struct CameraViewport {
    float width;
    float height;
    float offsetLeft;
    float offsetTop;
};

class Camera {
public:
    Camera(guVector position, guVector target);
    void SetViewport(float width, float height, float offsetLeft, float offsetTop);
    void SetPerspective(float fov, float nearPlane, float farPlane);
    void SetActive();

    static Mtx* GetActiveMtx();

private:
    static Camera* activeCamera;

    guVector position;
    guVector target;

    float fieldOfView;
    float nearPlane;
    float farPlane;

    CameraViewport viewport;
    Mtx44 perspectiveMtx;
    Mtx cameraMtx;
};