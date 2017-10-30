#include "graphics.h"

/* SDK libs */
#include <string.h>
#include <malloc.h>

/* GX vars */
#define DEFAULT_FIFO_SIZE	(256*1024)
static void *xfb[2] = { NULL, NULL };
static u32 fbi = 0;
static GXRModeObj *rmode = NULL;
BOOL first_frame = FALSE;
void *gpfifo = NULL;
f32 aspectRatio;
Mtx44 orthographicMatrix;

void Graphics::Init() {
	VIDEO_Init();

	/* Get render mode */
	rmode = VIDEO_GetPreferredMode(NULL); //&TVPal528Int;
	
	/* Try to get the framerate */
	Graphics::frameTime = 1.f / Graphics::GetFramerate();

	/* Allocate frame buffers */
	xfb[0] = (u32 *)MEM_K0_TO_K1(SYS_AllocateFramebuffer(rmode));
	xfb[1] = (u32 *)MEM_K0_TO_K1(SYS_AllocateFramebuffer(rmode));

	/* Clean buffers */
	VIDEO_ClearFrameBuffer(rmode, xfb[0], COLOR_BLACK);
	VIDEO_ClearFrameBuffer(rmode, xfb[1], COLOR_BLACK);

	VIDEO_Configure(rmode);
	VIDEO_SetNextFramebuffer(xfb[fbi]);
	VIDEO_SetBlack(TRUE);
	VIDEO_Flush();
	VIDEO_WaitVSync();
	if (rmode->viTVMode & VI_NON_INTERLACE) {
		VIDEO_WaitVSync();
	} else {
		while (VIDEO_GetNextField()) {
			VIDEO_WaitVSync();
		}
	}

	/* Set aspect ratio */
	aspectRatio = 4.f / 3.f;

#ifdef WII
	/* If 16:9 we need some hacks */
	if (CONF_GetAspectRatio()) {
		rmode->viWidth = 678;
		rmode->viXOrigin = (VI_MAX_WIDTH_PAL - 678) / 2;
		aspectRatio = 16.f / 9.f;
	}
#endif


	/* Swap frames */
	fbi ^= 1;

	/* Init flipper */
	gpfifo = MEM_K0_TO_K1(memalign(32, DEFAULT_FIFO_SIZE));
	memset(gpfifo, 0, DEFAULT_FIFO_SIZE);
	GX_Init(gpfifo, DEFAULT_FIFO_SIZE);

	/* Clear the background to black and clear the Z buf */
	GXColor background = { 0xa0, 0xe0, 0xf0, 0xff };
	GX_SetCopyClear(background, GX_MAX_Z24);

	GX_SetDispCopyYScale(GX_GetYScaleFactor(rmode->efbHeight, rmode->xfbHeight));
	GX_SetDispCopySrc(0, 0, rmode->fbWidth, rmode->efbHeight);
	GX_SetDispCopyDst(rmode->fbWidth, rmode->xfbHeight);
	GX_SetCopyFilter(rmode->aa, rmode->sample_pattern, GX_TRUE, rmode->vfilter);
	GX_SetFieldMode(rmode->field_rendering, ((rmode->viHeight == 2 * rmode->xfbHeight) ? GX_ENABLE : GX_DISABLE));

	GX_SetPixelFmt(rmode->aa ? GX_PF_RGB565_Z16 : GX_PF_RGB8_Z24, GX_ZC_LINEAR);

	GX_SetCullMode(GX_CULL_FRONT);
	GX_CopyDisp(xfb[fbi], GX_TRUE);
	GX_SetDispCopyGamma(GX_GM_1_0);

	/* Clear texture cache */
	GX_InvalidateTexAll();

	GX_SetNumTexGens(1);
	GX_SetTexCoordGen(GX_TEXCOORD0, GX_TG_MTX2x4, GX_TG_TEX0, GX_IDENTITY);
	GX_SetTevOrder(GX_TEVSTAGE0, GX_TEXCOORD0, GX_TEXMAP0, GX_COLOR0A0);

	first_frame = TRUE;

	Graphics::SetViewport(0, 0, rmode->viWidth, rmode->viHeight, 0, 1);
}

void Graphics::Done() {
	/* Finish up rendering */
	GX_SetZMode(GX_TRUE, GX_LEQUAL, GX_TRUE);
	GX_SetColorUpdate(GX_TRUE);
	GX_CopyDisp(xfb[fbi], GX_TRUE);

	GX_DrawDone();

	/* Flush and swap buffers */
	VIDEO_SetNextFramebuffer(xfb[fbi]);
	if (first_frame) {
		first_frame = 0;
		VIDEO_SetBlack(FALSE);
	}

	VIDEO_Flush();
	VIDEO_WaitVSync();
	fbi ^= 1;
}

GXRModeObj* Graphics::GetMode() {
	return rmode;
}

f32 Graphics::GetAspectRatio() {
	return aspectRatio;
}

void Graphics::Set2DMode() {
	GX_LoadProjectionMtx(orthographicMatrix, GX_ORTHOGRAPHIC);
}

void Graphics::SetViewport(f32 xOrig, f32 yOrig, f32 wd, f32 ht, f32 nearZ, f32 farZ) {
	GX_SetScissor(xOrig, yOrig, wd, ht);
	GX_SetViewport(xOrig, yOrig, wd, ht, nearZ, farZ);

	guOrtho(orthographicMatrix, 0, ht, 0, wd, 0.1f, 300.0f);
}

u32 Graphics::GetFramerate() {
	u32 tvmode = rmode->viTVMode >> 2;
	switch (tvmode) {
	case VI_NTSC:
	case VI_EURGB60:
	case VI_DEBUG:
	case VI_MPAL:
		return 60;
	case VI_PAL:
	case VI_DEBUG_PAL:
	default:
		return 50;
	}
}

f32 Graphics::frameTime = 0;