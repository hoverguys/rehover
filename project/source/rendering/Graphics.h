#pragma once
#include "../pchheader.h"

class Graphics {
public:
	/*! \brief Initialize the GX subsystem
	 */
	static void Init();

	/*! \brief Get video mode
	 *  \return Currently preferred/used video mode
	 */
	static GXRModeObj* GetMode();

	/*! \brief Get current aspect ratio
	 *  \return Aspect ratio as width/height
	 */
	static f32 GetAspectRatio();

	/*! \brief Finish rendering and swap buffers
	 */
	static void Done();

	/*! \brief Wait for vblank
	 */
	static void Wait();

	/*! \brief Load 2D ortho matrix for sprite/font rendering
	 */
	static void Set2DMode();

	/*! \brief Wrapper for changing viewport (adds scissors and 2d update)
	 */
	static void SetViewport(f32 xOrig, f32 yOrig, f32 wd, f32 ht, f32 nearZ, f32 farZ);

	/*! \brief Get frame rate depending on current TV mode
	 */
	static u32 GetFramerate();

	/*! \brief Get a simpler video mode (for PAL vs NTSC checks)
	 */
	static u32 GetGenericVideoMode();

private:
	Graphics() {}

	/* Frame time (1/60 or 1/50 depending on video mode) */
	static f32 frameTime;
	static const u32 DEFAULT_FIFO_SIZE = 256 * 1024;

	static void* xfb[2];
	static u32 fbi;

	static GXRModeObj* rmode;
	static bool first_frame;
	static void* gpfifo;
	static f32 aspectRatio;
	static Mtx44 orthographicMatrix;
};
