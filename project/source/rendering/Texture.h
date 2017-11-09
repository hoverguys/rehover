#pragma once

#include <ogc/gx.h>

class Texture {
public:
	/*! \brief Bind the texture to a texture mapping slot
	 *  \param texmapid What slot to bind the texture to
	 */
	void Bind(unsigned short texmapid);

protected:
	friend class TextureResource;

	GXTexObj object;
	unsigned char* data;
	unsigned short format;
	unsigned int width, height;
	bool mipmaps;
	unsigned short minlod, maxlod;
};