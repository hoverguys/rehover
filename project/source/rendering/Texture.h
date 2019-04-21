#pragma once

#include "../pchheader.h"

class Texture {
public:
	/*! \brief Bind the texture to a texture mapping slot
	 *  \param texmapid What slot to bind the texture to
	 *  \param tlut What slot to bind the texture lookup table (palette) to, if necessary
	 */
	void Bind(unsigned short texmapid, unsigned short tlut = GX_TLUT0);

protected:
	friend class TextureResource;

	GXTexObj object = {0};
	GXTlutObj paletteobj = {0};
	unsigned char *data, *palette;
	u32 width, height;
	u16 paletteCount, minlod, maxlod;
	u8 format, palfmt, wrapS, wrapT, filterMode;
};