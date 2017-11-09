#include "TextureResource.h"

#include <malloc.h>
#include <ogc/gx.h>
#include <stdio.h>
#include <string.h>

void TextureResource::Initialize() {
	header = static_cast<TextureResourceHeader*>(address);
	unsigned char* data = static_cast<unsigned char*>(address) + sizeof(TextureResourceHeader);

	Texture* t = new Texture();

	t->width = header->width;
	t->height = header->height;
	t->format = header->format;
	t->mipmaps = header->maxlod != 0 || header->minlod != 0;
	if (t->mipmaps) {
		t->maxlod = header->maxlod;
		t->minlod = header->minlod;
	}
	t->data = data;

	loaded = false;
	internal = t;
}

Texture* TextureResource::Load() {
	if (loaded) {
		return internal;
	}

	auto& t = internal;

	auto mipmap = t->mipmaps ? GX_TRUE : GX_FALSE;

	GX_InitTexObj(&t->object, t->data, t->width, t->height, t->format, GX_CLAMP, GX_CLAMP, mipmap);

	if (mipmap) {
		GX_InitTexObjLOD(&t->object, GX_LINEAR, GX_LINEAR, t->minlod, t->maxlod, 0, 0, 0, GX_ANISO_1);
	}

	loaded = true;
	return internal;
}