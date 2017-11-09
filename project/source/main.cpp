/* SDK Libraries */
#include <gccore.h>
#include <gctypes.h>
#include <math.h>
#include <stdio.h>

#include "Game.h"
#include "rendering/Camera.h"
#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"
#include "resources/TextureResource.h"

#include <entityx/entityx.h>

namespace ex = entityx;

bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();
	ResourceLoader::LoadPack("rehover_data.gcr");

	Game game;

	// DEBUG: Load hardcoded model
	auto meshresource = ResourceLoader::Load<MeshResource>("assets/models/hovercraft.obj");
	Mesh* mesh = meshresource->Load();

	game.init(mesh);

	// DEBUG: Load hardcoded texture
	auto texresource = ResourceLoader::Load<TextureResource>("assets/textures/hovercraftGlobal.png");
	Texture* texture = texresource->Load();

	texture->Bind(GX_TEXMAP0);

	// DEBUG Camera
	Camera* camera = new Camera({0, 0, -10}, {0, 0, 0});
	camera->SetActive();

	isRunning = TRUE;
	unsigned int frame = 0;
	while (isRunning) {

		// DEBUG Move Camera
		camera->Move({cos(frame * 0.01f) * 10, 5, sin(frame * 0.01f) * 10}, {0, 0, 0});

		// DEBUG Render matrices
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
		frame++;
	}

	return 0;
}

void OnResetCalled() { isRunning = FALSE; }