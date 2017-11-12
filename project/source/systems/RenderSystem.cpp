#include "RenderSystem.h"

#include "../components/Camera.h"
#include "../components/Renderable.h"
#include "../components/Transform.h"

namespace cp = Components;

void RenderSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Camera>([&](ex::Entity entity, cp::Transform& transform, cp::Camera& camera) {
		// Setup camera
		SetupCamera(camera);
		Mtx& lookat = transform.GetMatrix();
		guVector target = {0, 0, -1};
		guVecMultiply(lookat, &target, &target);

		// Create proper look at matrix
		Mtx cameraMatrix;
		guVector up = {0,1,0};
		guLookAt(cameraMatrix, &transform.position, &up, &target);

		// Render
		RenderScene(cameraMatrix, es, events, dt);
	});
};

void RenderSystem::RenderScene(Mtx& cameraMtx, ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Renderable>(
	    [&](ex::Entity entity, cp::Transform& transform, cp::Renderable& renderable) {
		    Mtx& modelMtx = transform.GetMatrix();

		    // Positional matrix with camera
		    Mtx modelviewMtx, modelviewInverseMtx;
		    guMtxConcat(cameraMtx, modelMtx, modelviewMtx);
		    GX_LoadPosMtxImm(modelviewMtx, GX_PNMTX0);

		    // Normals
		    guMtxInverse(modelviewMtx, modelviewInverseMtx);
		    guMtxTranspose(modelviewInverseMtx, modelviewInverseMtx);
		    GX_LoadNrmMtxImm(modelviewInverseMtx, GX_PNMTX0);

		    renderable.mesh->Render();
	    });
}

void RenderSystem::SetupCamera(cp::Camera& camera) {
	GX_SetScissor(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height);
	GX_SetViewport(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height,
	               0, 1);
	GX_LoadProjectionMtx(camera.perspectiveMtx, GX_PERSPECTIVE);
}
