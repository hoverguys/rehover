#pragma once
#include <ogc/gu.h>

namespace Components {

/*! Camera viewport in screen space */
struct CameraViewport {
	float width;      /*< Width of the viewport rect */
	float height;     /*< Height of the viewport rect */
	float offsetLeft; /*< X position relative to the left of the frame */
	float offsetTop;  /*< Y position relative to the top of the frame */
};

/*! \brief Camera component
 *  Gives an entity a camera that renders to a viewport (on screen)
 */
class Camera {
public:
	/*! \brief Create a camera
	 *  \param offsetLeft Viewport X position (relative to left)
	 *  \param offsetTop  Viewport Y position (relative to top)
	 *  \param width      Viewport width
	 *  \param height     Viewport height
	 */
	explicit Camera(float offsetLeft = 0, float offsetTop = 0, float width = 1, float height = 1);

	/*! \brief Set camera viewport
	 *  \param offsetLeft Viewport X position (relative to left)
	 *  \param offsetTop  Viewport Y position (relative to top)
	 *  \param width      Viewport width
	 *  \param height     Viewport height
	 */
	void SetViewport(float offsetLeft, float offsetTop, float width, float height);

	/*! \brief Set camera perspective projection matrix
	 *  \param fov       Field of view
	 *  \param nearPlane Near clipping plane distance
	 *  \param farPlane  Far clipping plane distance
	 */
	void SetPerspective(float fov, float nearPlane, float farPlane);

	/*! Field of view */
	float fieldOfView;

	/*! Near clipping plane distance */
	float nearPlane;

	/*! Far clipping plane distance */
	float farPlane;

	/*! Camera viewport */
	CameraViewport viewport;

	/*! Perspective projection matrix */
	Mtx44 perspectiveMtx;
};
} // namespace Components