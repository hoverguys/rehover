/* SDK Libraries */
#include <gccore.h>
#include <stdio.h>
#include <math.h>
#include <gctypes.h>

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "rendering/Camera.h"
#include "Game.h"
#include "utils/fnv.h"

#include <hovercraft_bmb.h>

#include <entityx/entityx.h>

namespace ex = entityx;


bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();

	Game game;

	//DEBUG: Load hardcoded model
	MeshResource* resource = new MeshResource((unsigned char*)hovercraft_bmb_txt, hovercraft_bmb_txt_size);
	Mesh* mesh = resource->Load();
	
	game.init(mesh);

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

		game.update(1.f / Graphics::GetFramerate());

		// Render here
		Graphics::Done();
		printf("%u\n", fnv1_hash("hovercraft.obj"));
		frame++;
	}

	return 0;
}

void OnResetCalled() {
	isRunning = FALSE;
}