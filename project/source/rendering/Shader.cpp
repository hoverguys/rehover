#include "Shader.h"
#include <ogc/gx.h>

void defaultTEV() {
	// 1 TEV Stage, 1 Texture (current)
	GX_SetNumIndStages(0);
	GX_SetNumTevStages(1);
	GX_SetNumTexGens(1);
	GX_SetNumChans(1);

	// Stage 1: Standard blending
	GX_SetTevOrder(GX_TEVSTAGE0, GX_TEXCOORD0, GX_TEXMAP0, GX_COLOR0A0);
	GX_SetTevColorIn(GX_TEVSTAGE0, GX_CC_ZERO, GX_CC_TEXC, GX_CC_RASC, GX_CC_ZERO);
	GX_SetTevColorOp(GX_TEVSTAGE0, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_FALSE, GX_TEVPREV);
	GX_SetTevAlphaIn(GX_TEVSTAGE0, GX_CA_ZERO, GX_CA_ZERO, GX_CA_ZERO, GX_CA_TEXA);
	GX_SetTevAlphaOp(GX_TEVSTAGE0, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_FALSE, GX_TEVPREV);

	// Reset color to #fff
	GX_SetChanMatColor(GX_COLOR0A0, GXColor{0xff, 0xff, 0xff, 0xff});
}

void Shader::Use() {
	GX_CallDispList(data, size);
}

void Shader::Default() {
	defaultTEV();

	// Set light on
	GX_SetChanCtrl(GX_COLOR0A0, GX_ENABLE, GX_SRC_REG, GX_SRC_REG, GX_MAXLIGHT - 1, GX_DF_CLAMP, GX_AF_NONE);
}

void Shader::DefaultUnlit() {
	defaultTEV();

	// Set light off
	GX_SetChanCtrl(GX_COLOR0A0, GX_DISABLE, GX_SRC_REG, GX_SRC_REG, GX_LIGHT0, GX_DF_CLAMP, GX_AF_NONE);
}