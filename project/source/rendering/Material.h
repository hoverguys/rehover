#pragma once
#include "Shader.h"
#include "Texture.h"

struct ShaderVars {
	GXColor color0, color1;
};

struct Material {
	std::array<std::shared_ptr<Texture>, 8> textures;
	std::shared_ptr<Shader> shader;
	ShaderVars uniforms;
};