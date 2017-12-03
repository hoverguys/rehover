#include "UISystem.h"

#include "../rendering/Graphics.h"

#include "../components/Sprite.h"
#include "../components/Transform.h"

namespace cp = Components;

void UISystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	// Setup 2D rendering
	Setup2D();

	es.each<cp::Transform, cp::Sprite>([&](ex::Entity entity, cp::Transform& transform, cp::Sprite& sprite) {
		const Matrix& spriteMtx = transform.GetMatrix();

		// Positional matrix with camera
		Mtx nativeTemp;
		spriteMtx.ToNative(nativeTemp);
		GX_LoadPosMtxImm(nativeTemp, GX_PNMTX0);

		auto material = sprite.material;
		if (material) {
			material->Use();
		}

		auto uv = sprite.bounds.Bounds();

		GX_Begin(GX_QUADS, GX_VTXFMT0, 4);

		/* Top left */
		GX_Position2f32(0, 0);
		GX_TexCoord2f32(uv.first.x, uv.first.y);

		/* Bottom left */
		GX_Position2f32(0, sprite.size.y);
		GX_TexCoord2f32(uv.first.x, uv.second.y);

		/* Bottom right */
		GX_Position2f32(sprite.size.x, sprite.size.y);
		GX_TexCoord2f32(uv.second.x, uv.second.y);

		/* Top right */
		GX_Position2f32(sprite.size.x, 0);
		GX_TexCoord2f32(uv.second.x, uv.first.y);

		GX_End();
	});
}

void UISystem::Setup2D() {
	// Use orthogonal
	Graphics::Set2DMode();

	// Use identity matrix for normals
	Mtx dummy;
	guMtxIdentity(dummy);
	GX_LoadNrmMtxImm(dummy, GX_PNMTX0);
}