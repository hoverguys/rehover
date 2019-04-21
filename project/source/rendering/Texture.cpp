#include "Texture.h"

void Texture::Bind(unsigned short texmapid, unsigned short tlutid) {
	GX_LoadTexObj(&object, texmapid);
	if (palette != nullptr) {
		GX_LoadTlut(&paletteobj, tlutid);
		GX_InitTexObjTlut(&object, tlutid);
	}
}