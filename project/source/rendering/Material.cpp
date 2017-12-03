#include "Material.h"

void Material::Use() {
	if (shader) {
		// Setup shader
		shader->Use();

		// Setup shader uniforms
		auto settings = uniforms;
		GX_SetChanAmbColor(GX_COLOR0A0, GXColor{0x00, 0x00, 0x00, 0x00});
		GX_SetChanMatColor(GX_COLOR0A0, settings.color0);
		GX_SetChanMatColor(GX_COLOR1A1, settings.color1);
	} else {
		Shader::Default();
	}

	for (unsigned int i = 0; i < textures.size(); i++) {
		if (textures[i]) {
			textures[i]->Bind(i);
		}
	}
}