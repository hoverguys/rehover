#include "Texture.h"
#include <malloc.h>

Texture::Texture() { object = nullptr; }

Texture::~Texture() {
	if (object != nullptr) {
		free(object);
	}
}

void Texture::Bind(unsigned short texmapid) { GX_LoadTexObj(object, texmapid); }