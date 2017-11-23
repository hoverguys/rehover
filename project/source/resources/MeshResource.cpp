#include "MeshResource.h"

#include <malloc.h>
#include <memory>
#include <ogc/gx.h>
#include <stdio.h>
#include <string.h>

void MeshResource::Initialize() {
	header = static_cast<MeshResourceHeader*>(address);
	unsigned char* base = static_cast<unsigned char*>(address);

	const unsigned int posOffset = sizeof(MeshResourceHeader);
	const unsigned int nrmOffset = posOffset + (sizeof(float) * 3 * header->vcount);
	const unsigned int texOffset = nrmOffset + (sizeof(float) * 3 * header->ncount);
	const unsigned int indOffset = texOffset + (sizeof(float) * 2 * header->vtcount);

	auto m = std::make_shared<Mesh>();

	// TODO: Different casting?
	m->positionArray = reinterpret_cast<float*>(base + posOffset);
	m->normalArray = reinterpret_cast<float*>(base + nrmOffset);
	m->uvArray = reinterpret_cast<float*>(base + texOffset);
	m->indexArray = reinterpret_cast<MeshIndex*>(base + indOffset);

	loaded = false;
	internal = m;
}

std::shared_ptr<Mesh> MeshResource::Load() {
	// Early out
	if (loaded) {
		return internal;
	}

	// Calculate cost
	const unsigned int indicesCount = header->fcount * 3;
	const unsigned int indicesSize =
	    indicesCount * sizeof(MeshIndex); // 3 indices per vertex index (p,n,t) that are u16 in size
	const unsigned int callSize = 89;     // Size of setup var

	// Round up to nearest 32 multiplication
	const unsigned int dispSize = (((indicesSize + callSize + 63) >> 5) + 1) << 5;

	// Allocate display list
	internal->displayList = memalign(32, dispSize);
	memset(internal->displayList, 0, dispSize);
	DCInvalidateRange(internal->displayList, dispSize);

	// Build display list
	GX_BeginDispList(internal->displayList, dispSize);

	GX_ClearVtxDesc();
	GX_SetVtxDesc(GX_VA_POS, GX_INDEX16);
	GX_SetVtxDesc(GX_VA_NRM, GX_INDEX16);
	GX_SetVtxDesc(GX_VA_TEX0, GX_INDEX16);

	GX_SetVtxAttrFmt(GX_VTXFMT0, GX_VA_POS, GX_POS_XYZ, GX_F32, 0);
	GX_SetVtxAttrFmt(GX_VTXFMT0, GX_VA_NRM, GX_NRM_XYZ, GX_F32, 0);
	GX_SetVtxAttrFmt(GX_VTXFMT0, GX_VA_TEX0, GX_TEX_ST, GX_F32, 0);

	GX_SetArray(GX_VA_POS, (void*)internal->positionArray, 3 * sizeof(float));
	GX_SetArray(GX_VA_NRM, (void*)internal->normalArray, 3 * sizeof(float));
	GX_SetArray(GX_VA_TEX0, (void*)internal->uvArray, 2 * sizeof(float));

	/* Fill the list with indices */
	GX_Begin(GX_TRIANGLES, GX_VTXFMT0, indicesCount);
	for (unsigned int i = 0; i < indicesCount; i++) {
		MeshIndex index = internal->indexArray[i];
		GX_Position1x16(index.vertex);
		GX_Normal1x16(index.normal);
		GX_TexCoord1x16(index.uv);
	}
	GX_End();

	/* Close display list */
	internal->displayListSize = GX_EndDispList();
	if (internal->displayListSize == 0) {
		printf("Error: Display list is wrong size [%u]\n", dispSize);
		return NULL;
	}

	loaded = true;
	return internal;
}