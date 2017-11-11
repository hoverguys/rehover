#include "RenderSystem.h"

#include "../components/Camera.h"
#include "../components/Renderable.h"
#include "../components/Transform.h"

namespace cp = Components;

void RenderSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Camera>([&](ex::Entity entity, cp::Transform& transform, cp::Camera& camera) {
		// Setup camera
		SetupCamera(camera);
		Mtx& cameraMatrix = transform.GetMatrix();

		// Render
		RenderScene(cameraMatrix, es, events, dt);
	});
};

void RenderSystem::RenderScene(Mtx& cameraMtx, ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Renderable>(
	    [&](ex::Entity entity, cp::Transform& transform, cp::Renderable& renderable) {
		    Mtx& objectMtx = transform.GetMatrix();

		    // Positional matrix with camera
		    Mtx modelviewMtx, modelviewInverseMtx;
		    guMtxConcat(cameraMtx, objectMtx, modelviewMtx);
		    GX_LoadPosMtxImm(modelviewMtx, GX_PNMTX0);

		    // Normals
		    guMtxInverse(modelviewMtx, modelviewInverseMtx);
		    guMtxTranspose(modelviewInverseMtx, modelviewMtx);
		    GX_LoadNrmMtxImm(modelviewMtx, GX_PNMTX0);

		    renderable.mesh->Render();
	    });
}

void RenderSystem::SetupCamera(cp::Camera& camera) {
	GX_SetScissor(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height);
	GX_SetViewport(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height,
	               0, 1);
	GX_LoadProjectionMtx(camera.perspectiveMtx, GX_PERSPECTIVE);
}
