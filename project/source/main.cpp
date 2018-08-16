#include "pchheader.h"
#include <fat.h>
#include <ogc/lwp_watchdog.h>


#include "rendering/Graphics.h"
#include "resources/MeshResource.h"
#include "resources/ResourceLoader.h"
#include "resources/TextureResource.h"

#include "Game.h"

int main() {
	// DEBUG: Enable USBGecko
	CON_EnableGecko(1, 0);

	// Initialize graphics
	Graphics::Init();

	// Setup storage
	fatInitDefault();
#ifdef WII
	fatMountSimple("sd", &__io_wiisd);
#else
	fatMountSimple("sd", &__io_gcsdb);
#endif

	// Load assets
	ResourceLoader::LoadPack("sd:/rehover-data.gcr");

	// Start game
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
