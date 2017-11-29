#pragma once
#include "../components/Camera.h"
#include <entityx/entityx.h>
#include "../math/Matrix.h"

namespace ex = entityx;

class RenderSystem : public ex::System<RenderSystem> {
public:
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
private:
	void RenderScene(const Matrix& cameraMtx, ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt);
	void SetupLights(const Matrix& cameraMtx, ex::EntityManager& es);
	static void SetupCamera(Components::Camera& camera);
};