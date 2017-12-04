#pragma once
#include "Shader.h"
#include "Texture.h"

/*! \brief Shader parameters */
struct ShaderVars {
	/*! Color channels */
	GXColor color0, color1;
};

/*! Mesh Material */
struct Material {
	/*! Textures used by the material/shader */
	std::array<std::shared_ptr<Texture>, 8> textures;

	/*! \brief Shader object
	 *  If null, a default one will be used when rendering
	 */
	std::shared_ptr<Shader> shader;

	/*! Shader variables */
	ShaderVars uniforms;
};