#include "RenderSystem.h"

#include "../components/Camera.h"
#include "../components/Light.h"
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
		guVector up = {0, 1, 0};
		guLookAt(cameraMatrix, &transform.position, &up, &target);

		// Setup lights
		SetupLights(cameraMatrix, es);

		// Render
		RenderScene(cameraMatrix, es, events, dt);
	});
}

void RenderSystem::SetupLights(Mtx& cameraMtx, ex::EntityManager& es) {
	unsigned short lightId = GX_LIGHT0;
	es.each<cp::Transform, cp::DirectionalLight>(
	    [&](ex::Entity entity, cp::Transform& transform, cp::DirectionalLight& light) {
		    // Too many lights?
		    if (lightId >= GX_MAXLIGHT) {
			    // We should give an error or something, at least on debug
			    return;
		    }

		    light.Setup(cameraMtx, transform);
		    light.Bind(lightId);

		    // Increase light id (it's a bitmask)
		    lightId = lightId << 1;
	    });
}

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

		    auto material = renderable.material;
		    if (material) {
			    auto shader = material->shader;
			    if (shader) {
				    // Setup shader
				    shader->Use();

				    // Setup shader uniforms
				    auto settings = material->uniforms;
				    GX_SetChanAmbColor(GX_COLOR0A0, GXColor{0x00, 0x00, 0x00, 0x00});
				    GX_SetChanMatColor(GX_COLOR0A0, settings.color0);
				    GX_SetChanMatColor(GX_COLOR1A1, settings.color1);
			    } else {
				    Shader::Default();
			    }

			    auto textures = material->textures;
			    for (unsigned int i = 0; i < textures.size(); i++) {
				    if (textures[i]) {
					    textures[i]->Bind(i);
				    }
			    }
		    }

		    renderable.mesh->Render();
	    });
}

void RenderSystem::SetupCamera(cp::Camera& camera) {
	GX_SetScissor(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height);
	GX_SetViewport(camera.viewport.offsetLeft, camera.viewport.offsetTop, camera.viewport.width, camera.viewport.height,
	               0, 1);
	GX_LoadProjectionMtx(camera.perspectiveMtx, GX_PERSPECTIVE);
}
