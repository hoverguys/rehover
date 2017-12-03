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
				Shader::DefaultUnlit();
			}

			auto textures = material->textures;
			for (unsigned int i = 0; i < textures.size(); i++) {
				if (textures[i]) {
					textures[i]->Bind(i);
				}
			}
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

	// Setup vertex descriptors to use immediate mode
	GX_ClearVtxDesc();
	GX_SetVtxDesc(GX_VA_POS, GX_DIRECT);
	GX_SetVtxDesc(GX_VA_TEX0, GX_DIRECT);
	GX_SetVtxAttrFmt(GX_VTXFMT0, GX_VA_POS, GX_POS_XY, GX_F32, 0);
	GX_SetVtxAttrFmt(GX_VTXFMT0, GX_VA_TEX0, GX_TEX_ST, GX_F32, 0);
}