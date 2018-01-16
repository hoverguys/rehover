#include <ogc/lwp_watchdog.h>
#include <fat.h>
#include "pchheader.h"

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"
#include "resources/TextureResource.h"

#include "Game.h"

int main() {
	// DEBUG: Enable USBGecko
	CON_EnableGecko(1, 0);

	Graphics::Init();
	fatInitDefault();
	ResourceLoader::LoadPack("rehover_data.gcr");

	Game game;
	game.init();

	while (!SYS_ResetButtonDown()) {
		// Logic
		auto updateStart = gettime();
		game.update(1.f / Graphics::GetFramerate());
		auto updateEnd = gettime();

		// Render to XFB
		Graphics::Done();
		auto graphicsEnd = gettime();

		// Metrics
		auto updateDelta = ticks_to_nanosecs(diff_ticks(updateStart, updateEnd));
		auto frameDelta = ticks_to_nanosecs(diff_ticks(updateEnd, graphicsEnd));
		std::printf(":frametime frame:%llu9 logic:%llu9\n", frameDelta, updateDelta);

		Graphics::Wait();
	}

	return 0;
}
