/* SDK Libraries */
#include <gccore.h>
#include <gctypes.h>
#include <math.h>
#include <stdio.h>

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"
#include "resources/TextureResource.h"

#include "Game.h"

int main() {
	// DEBUG: Enable USBGecko
	CON_EnableGecko(1, 0);

	Graphics::Init();
	ResourceLoader::LoadPack("rehover_data.gcr");

	Game game;
	game.init();

	while (!SYS_ResetButtonDown ()) {
		game.update(1.f / Graphics::GetFramerate());

		// Render here
		Graphics::Done();
	}

	return 0;
}
