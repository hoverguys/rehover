#include "RenderSystem.h"

#include "../components/Camera.h"
#include "../components/Light.h"
#include "../components/Renderable.h"
#include "../components/Transform.h"
#include "../math/Math.h"

namespace cp = Components;

void RenderSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Camera>([&](ex::Entity entity, cp::Transform& transform, cp::Camera& camera) {
		// Setup camera
		SetupCamera(camera);

		const Matrix& lookat = transform.GetMatrix();
		const Vector target = lookat.Multiply(Math::worldForward);

		// Create proper look at matrix
		const Matrix cameraMatrix = Matrix::LookAt(transform.position, Math::worldUp, target);

		// Setup lights
		SetupLights(cameraMatrix, es);

		// Render
		RenderScene(cameraMatrix, es, events, dt);
	});
}

void RenderSystem::SetupLights(const Matrix& cameraMtx, ex::EntityManager& es) {
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
		}
	);
}

void RenderSystem::RenderScene(const Matrix& cameraMtx, ex::EntityManager& es, ex::EventManager& events,
							   ex::TimeDelta dt) {
	es.each<cp::Transform, cp::Renderable>(
		[&](ex::Entity entity, cp::Transform& transform, cp::Renderable& renderable) {
			const Matrix& modelMtx = transform.GetMatrix();

			// Positional matrix with camera
			Mtx nativeTemp;
			const Matrix modelviewMtx = cameraMtx * modelMtx;
			modelviewMtx.ToNative(nativeTemp);
			GX_LoadPosMtxImm(nativeTemp, GX_PNMTX0);

			// Normals
			Matrix modelviewInverseMtx = modelviewMtx.Inversed();
			modelviewInverseMtx.Transpose();
			modelviewInverseMtx.ToNative(nativeTemp);
			GX_LoadNrmMtxImm(nativeTemp, GX_PNMTX0);

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

	// Disable alpha blending
	/// \todo Maybe this SHOULD be on, but in a different way than sprites
	GX_SetBlendMode(GX_BM_BLEND, GX_BL_SRCALPHA, GX_BL_INVSRCALPHA, GX_LO_CLEAR);
}
