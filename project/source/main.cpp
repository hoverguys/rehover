/* SDK Libraries */
#include <gccore.h>
#include <stdio.h>
#include <gctypes.h>

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "rendering/Camera.h"

#include <hovercraft_bmb.h>


bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();

	//DEBUG: Load hardcoded model
	MeshResource* resource = new MeshResource((unsigned char*)hovercraft_bmb_txt, hovercraft_bmb_txt_size);
	Mesh* mesh = resource->Load();

	//DEBUG Camera
	Camera* camera = new Camera( { 0, 0, -10 }, { 0, 0, 100 });
	camera->SetActive();

	//DEBUG Render matrices
	Mtx modelviewMtx, modelviewInverseMtx, objectMtx;
	guMtxIdentity(objectMtx);
	guMtxConcat(*camera->GetActiveMtx(), objectMtx, modelviewMtx);
	GX_LoadPosMtxImm(modelviewMtx, GX_PNMTX0);

	guMtxInverse(modelviewMtx, modelviewInverseMtx);
	guMtxTranspose(modelviewInverseMtx, modelviewMtx);

	GX_LoadNrmMtxImm(modelviewMtx, GX_PNMTX0);

	isRunning = TRUE;
	while (isRunning) {
		mesh->Render();

		// Render here
		Graphics::Done();
		printf("Console doesnt work yet");
	}

	return 0;
}

void OnResetCalled() {
	isRunning = FALSE;
}