#pragma once

#include "../pchheader.h"

class Texture {
public:
	/*! \brief Bind the texture to a texture mapping slot
	 *  \param texmapid What slot to bind the texture to
	 */
	void Bind(unsigned short texmapid);

protected:
	friend class TextureResource;

	GXTexObj object = {0};
	unsigned char* data;
	unsigned short format;
	unsigned int width, height;
	unsigned short minlod, maxlod;
	unsigned char wrapS, wrapT;
	unsigned char filterMode;
};