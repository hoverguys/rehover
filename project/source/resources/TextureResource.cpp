#include "TextureResource.h"

#include <malloc.h>
#include <ogc/gx.h>
#include <stdio.h>
#include <string.h>

void TextureResource::Initialize() {
	header = static_cast<TextureResourceHeader*>(address);
	unsigned char* data = static_cast<unsigned char*>(address) + header->dataOffset;
	printf("Loading texture (%dx%d fmt %d wrap S%x T%x filter %x) at offset %x\n", header->width, header->height,
	       header->format, header->wrapS, header->wrapT, header->filter, data);

	auto t = std::make_shared<Texture>();

	t->width = header->width;
	t->height = header->height;
	t->format = header->format;
	t->wrapS = header->wrapS;
	t->wrapT = header->wrapT;
	t->filterMode = header->filter;
	t->maxlod = header->maxlod;
	t->minlod = header->minlod;
	t->data = data;

	loaded = false;
	internal = t;
}

std::shared_ptr<Texture> TextureResource::Load() {
	if (loaded) {
		return internal;
	}

	auto& t = internal;

	auto useMipmaps = t->minlod + t->maxlod != 0;
	auto filterTex = GX_LINEAR;
	auto filterMip = GX_LINEAR;
	auto useTrilinear = useMipmaps && t->filterMode == 3 ? GX_TRUE : GX_FALSE;

	switch (t->filterMode) {
	case 0:
		filterTex = useMipmaps ? GX_NEAR_MIP_NEAR : GX_NEAR;
		filterMip = GX_NEAR;
		break;
	case 1:
		filterTex = useMipmaps ? GX_NEAR_MIP_LIN : GX_NEAR;
		filterMip = GX_LINEAR;
		break;
	case 2:
		filterTex = useMipmaps ? GX_LIN_MIP_NEAR : GX_LINEAR;
		filterMip = GX_NEAR;
		break;
	case 3:
		filterTex = GX_LIN_MIP_LIN;
		filterMip = GX_LINEAR;
		break;
	}

	GX_InitTexObj(&t->object, t->data, t->width, t->height, t->format, t->wrapS, t->wrapT, useTrilinear);
	GX_InitTexObjLOD(&t->object, filterTex, filterMip, t->minlod, t->maxlod, 0, 0, 0, GX_ANISO_1);

	loaded = true;
	return internal;
}