#include "Texture.h"

void Texture::Bind(unsigned short texmapid) { GX_LoadTexObj(&object, texmapid); }