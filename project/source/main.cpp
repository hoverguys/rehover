/* SDK Libraries */
#include <gccore.h>
#include <gctypes.h>
#include <math.h>
#include <stdio.h>

#include "rendering/Graphics.h"
#include "resources/ResourceLoader.h"
#include "resources/MeshResource.h"
#include "resources/TextureResource.h"

#include "Game.h"

bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();
	ResourceLoader::LoadPack("rehover_data.gcr");

	// DEBUG: Load hardcoded model
	auto meshresource = ResourceLoader::Load<MeshResource>("assets/models/hovercraft.obj");
	Mesh* mesh = meshresource->Load();

	// DEBUG: Load hardcoded texture
	auto texresource = ResourceLoader::Load<TextureResource>("assets/textures/hovercraftGlobal.png");
	Texture* texture = texresource->Load();
	texture->Bind(GX_TEXMAP0);

	Game game;
	game.init(mesh);

	isRunning = true;
	while (isRunning) {
		game.update(1.f / Graphics::GetFramerate());

		// Render here
		Graphics::Done();
	}

	return 0;
}

void OnResetCalled() {
	isRunning = false;
}