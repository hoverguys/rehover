/* SDK Libraries */
#include <gccore.h>
#include <stdio.h>
#include <gctypes.h>
#include <math.h>

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
	Camera* camera = new Camera( { 0, 0, -10 }, { 0, 0, 0 });
	camera->SetActive();

	isRunning = TRUE;
	unsigned int frame = 0;
	while (isRunning) {
		
		//DEBUG Move Camera
		camera->Move({cos(frame * 0.01f) * 10, 5, sin(frame * 0.01f) * 10}, { 0, 0, 0 });

		//DEBUG Render matrices
		Mtx modelviewMtx, modelviewInverseMtx, objectMtx;
		guMtxIdentity(objectMtx);
		guMtxConcat(*camera->GetActiveMtx(), objectMtx, modelviewMtx);
		GX_LoadPosMtxImm(modelviewMtx, GX_PNMTX0);

		guMtxInverse(modelviewMtx, modelviewInverseMtx);
		guMtxTranspose(modelviewInverseMtx, modelviewMtx);

		GX_LoadNrmMtxImm(modelviewMtx, GX_PNMTX0);

		mesh->Render();

		// Render here
		Graphics::Done();
		printf("Frame %d\n", frame);
		frame++;
	}

	return 0;
}

void OnResetCalled() {
	isRunning = FALSE;
}