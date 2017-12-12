package main

import (
	"encoding/binary"
)

var __gx gxregdef
var wgPipe Fifo = NullFifo{}
var _gxtevcolid = [9]uint8{0, 1, 0, 1, 0, 1, 7, 5, 6}

func _SHIFTL(v, s, w uint32) uint32 {
	return (v & ((1 << w) - 1)) << s
}
func _SHIFTR(v, s, w uint32) uint32 {
	return (v >> s) & ((1 << w) - 1)
}

func GX_LOAD_BP_REG(x uint32) {
	wgPipe.U8(0x61)
	wgPipe.U32(x)
	if *debugmsg && recording {
		bytedata := make([]byte, 5)
		bytedata[0] = 0x61
		binary.BigEndian.PutUint32(bytedata[1:], x)
		logDataWrite(bytedata)
	}
}

func GX_LOAD_XF_REG(x, y uint32) {
	wgPipe.U8(0x10)
	wgPipe.U32(x & 0xffff)
	wgPipe.U32(y)
	if *debugmsg && recording {
		bytedata := make([]byte, 9)
		bytedata[0] = 0x10
		binary.BigEndian.PutUint32(bytedata[1:], x&0xffff)
		binary.BigEndian.PutUint32(bytedata[5:], y)
		logDataWrite(bytedata)
	}
}

func GX_LOAD_CP_REG(x uint8, y uint32) {
	wgPipe.U8(0x08)
	wgPipe.U8(x)
	wgPipe.U32(y)
	if *debugmsg && recording {
		bytedata := make([]byte, 6)
		bytedata[0] = 0x08
		bytedata[1] = x
		binary.BigEndian.PutUint32(bytedata[2:], y)
		logDataWrite(bytedata)
	}
}

func __GX_SetDirtyState() {
	if __gx.dirtyState&0x0001 != 0 {
		__GX_SetSUTexRegs()
	}
	if __gx.dirtyState&0x0002 != 0 {
		__GX_UpdateBPMask()
	}
	if __gx.dirtyState&0x0004 != 0 {
		__GX_SetGenMode()
	}
	/* Unrelated to TEV
	if __gx.dirtyState&0x0008 != 0 {
		__GX_SetVCD()
	}
	if __gx.dirtyState&0x0010 != 0 {
		__GX_SetVAT()
	}
	*/
	if __gx.dirtyState&(^uint32(0xff)) != 0 {
		if __gx.dirtyState&0x0f00 != 0 {
			__GX_SetChanColor()
		}
		if __gx.dirtyState&0x0100f000 != 0 {
			__GX_SetChanCntrl()
		}
		if __gx.dirtyState&0x02ff0000 != 0 {
			__GX_SetTexCoordGen()
		}
		if __gx.dirtyState&0x04000000 != 0 {
			__GX_SetMatrixIndex(0)
			__GX_SetMatrixIndex(5)
		}
	}
	__gx.dirtyState = 0
}

func __GX_SetGenMode() {
	GX_LOAD_BP_REG(__gx.genMode)
	__gx.xfFlush = 0
}

func __SetSURegs(texmap, texcoord uint8) {
	wd := __gx.texMapSize[texmap] & 0x3ff
	ht := _SHIFTR(__gx.texMapSize[texmap], 10, 10)
	wrap_s := __gx.texMapWrap[texmap] & 3
	wrap_t := _SHIFTR(__gx.texMapWrap[texmap], 2, 2)

	reg := (texcoord & 0x7)
	__gx.suSsize[reg] = (__gx.suSsize[reg] & (^uint32(0x0000ffff))) | wd
	__gx.suTsize[reg] = (__gx.suTsize[reg] & (^uint32(0x0000ffff))) | ht
	__gx.suSsize[reg] = (__gx.suSsize[reg] & (^uint32(0x00010000))) | (_SHIFTL(wrap_s, 16, 1))
	__gx.suTsize[reg] = (__gx.suTsize[reg] & (^uint32(0x00010000))) | (_SHIFTL(wrap_t, 16, 1))

	GX_LOAD_BP_REG(__gx.suSsize[reg])
	GX_LOAD_BP_REG(__gx.suTsize[reg])
}

func __GX_SetSUTexRegs() {
	var texcoord uint8
	var texmap uint8

	dirtev := uint8((_SHIFTR(__gx.genMode, 10, 4)) + 1)
	indtev := uint8(_SHIFTR(__gx.genMode, 16, 3))

	//indirect texture order
	for i := uint8(0); i < indtev; i++ {
		switch i {
		case GX_INDTEXSTAGE0:
			texmap = uint8(__gx.tevRasOrder[2] & 7)
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[2], 3, 3))
		case GX_INDTEXSTAGE1:
			texmap = uint8(_SHIFTR(__gx.tevRasOrder[2], 6, 3))
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[2], 9, 3))
		case GX_INDTEXSTAGE2:
			texmap = uint8(_SHIFTR(__gx.tevRasOrder[2], 12, 3))
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[2], 15, 3))
		case GX_INDTEXSTAGE3:
			texmap = uint8(_SHIFTR(__gx.tevRasOrder[2], 18, 3))
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[2], 21, 3))
		default:
			texmap = 0
			texcoord = 0
		}

		texcm := _SHIFTL(1, uint32(texcoord), 1)
		if (__gx.texCoordManually & texcm) == 0 {
			__SetSURegs(texmap, texcoord)
		}
	}

	//direct texture order
	for i := uint8(0); i < dirtev; i++ {
		tevreg := 3 + (_SHIFTR(uint32(i), 1, 3))
		texmap = uint8(__gx.tevTexMap[i])

		if i&1 != 0 {
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[tevreg], 15, 3))
		} else {
			texcoord = uint8(_SHIFTR(__gx.tevRasOrder[tevreg], 3, 3))
		}

		tevm := _SHIFTL(1, uint32(i), 1)
		texcm := _SHIFTL(1, uint32(texcoord), 1)
		if texmap != 0xff && (__gx.tevTexCoordEnable&tevm) != 0 && (__gx.texCoordManually&texcm) == 0 {
			__SetSURegs(texmap, texcoord)
		}
	}
}

// This function originally does nothing on Wii
// If we have Wii-only graphical issues this would be a good point to look into
func __GX_UpdateBPMask() {
	nbmp := _SHIFTR(__gx.genMode, 16, 3)

	var ntexmap uint32
	nres := uint8(0)
	for i := uint32(0); i < nbmp; i++ {
		switch uint8(i) {
		case GX_INDTEXSTAGE0:
			ntexmap = __gx.tevRasOrder[2] & 7
			break
		case GX_INDTEXSTAGE1:
			ntexmap = _SHIFTR(__gx.tevRasOrder[2], 6, 3)
			break
		case GX_INDTEXSTAGE2:
			ntexmap = _SHIFTR(__gx.tevRasOrder[2], 12, 3)
			break
		case GX_INDTEXSTAGE3:
			ntexmap = _SHIFTR(__gx.tevRasOrder[2], 18, 3)
			break
		default:
			ntexmap = 0
			break
		}
		nres |= (1 << ntexmap)
	}

	if uint8(__gx.tevIndMask) != nres {
		__gx.tevIndMask = (__gx.tevIndMask & (^uint32(0xff))) | uint32(nres)
		GX_LOAD_BP_REG(__gx.tevIndMask)
	}
}

func __GX_SetChanColor() {
	if __gx.dirtyState&0x0100 != 0 {
		GX_LOAD_XF_REG(0x100a, __gx.chnAmbColor[0])
	}
	if __gx.dirtyState&0x0200 != 0 {
		GX_LOAD_XF_REG(0x100b, __gx.chnAmbColor[1])
	}
	if __gx.dirtyState&0x0400 != 0 {
		GX_LOAD_XF_REG(0x100c, __gx.chnMatColor[0])
	}
	if __gx.dirtyState&0x0800 != 0 {
		GX_LOAD_XF_REG(0x100d, __gx.chnMatColor[1])
	}
}

func __GX_SetChanCntrl() {
	if __gx.dirtyState&0x01000000 != 0 {
		GX_LOAD_XF_REG(0x1009, (_SHIFTR(__gx.genMode, 4, 3)))
	}

	i := 0
	channel := uint32(0x100e)
	mask := _SHIFTR(__gx.dirtyState, 12, 4)
	for mask != 0 {
		if mask&0x0001 != 0 {
			GX_LOAD_XF_REG(channel, __gx.chnCntrl[i])
		}

		mask >>= 1
		channel++
		i++
	}
}

func __GX_SetTexCoordGen() {
	if __gx.dirtyState&0x02000000 != 0 {
		GX_LOAD_XF_REG(0x103f, (__gx.genMode & 0xf))
	}

	i := uint32(0)
	texcoord := uint32(0x1040)
	mask := _SHIFTR(__gx.dirtyState, 16, 8)
	for mask != 0 {
		if mask&0x0001 != 0 {
			GX_LOAD_XF_REG(texcoord, __gx.texCoordGen[i])
			GX_LOAD_XF_REG((texcoord + 0x10), __gx.texCoordGen2[i])
		}
		mask >>= 1
		texcoord++
		i++
	}
}

func __GX_SetMatrixIndex(mtx uint32) {
	if mtx < 5 {
		GX_LOAD_CP_REG(0x30, __gx.mtxIdxLo)
		GX_LOAD_XF_REG(0x1018, __gx.mtxIdxLo)
	} else {
		GX_LOAD_CP_REG(0x40, __gx.mtxIdxHi)
		GX_LOAD_XF_REG(0x1019, __gx.mtxIdxHi)
	}
}

func GX_Init() {
	__gx.gxFifoInited = 1
	__gx.tevIndMask = 0xff
	__gx.tevIndMask = (__gx.tevIndMask & (^uint32(0xff000000))) | (_SHIFTL(0x0f, 24, 8))

	i := 0
	re0 := uint32(0xc0)
	re1 := uint32(0xc1)
	for i < 16 {
		__gx.tevColorEnv[i] = (__gx.tevColorEnv[i] & (^uint32(0xff000000))) | (_SHIFTL(re0, 24, 8))
		__gx.tevAlphaEnv[i] = (__gx.tevAlphaEnv[i] & (^uint32(0xff000000))) | (_SHIFTL(re1, 24, 8))
		re0 += 2
		re1 += 2
		i++
	}

	__gx.texCoordManually = 0
	__gx.dirtyState = 0

	__gx.saveDLctx = 1
	__gx.gxFifoUnlinked = 0

	__gx.sciTLcorner = (__gx.sciTLcorner & (^uint32(0xff000000))) | (_SHIFTL(0x20, 24, 8))
	__gx.sciBRcorner = (__gx.sciBRcorner & (^uint32(0xff000000))) | (_SHIFTL(0x21, 24, 8))
	__gx.lpWidth = (__gx.lpWidth & (^uint32(0xff000000))) | (_SHIFTL(0x22, 24, 8))
	__gx.genMode = (__gx.genMode & (^uint32(0xff000000))) | (_SHIFTL(0x00, 24, 8))

	i = 0
	re0 = 0x30
	re1 = 0x31
	for i < 8 {
		__gx.suSsize[i] = (__gx.suSsize[i] & (^uint32(0xff000000))) | (_SHIFTL(re0, 24, 8))
		__gx.suTsize[i] = (__gx.suTsize[i] & (^uint32(0xff000000))) | (_SHIFTL(re1, 24, 8))
		re0 += 2
		re1 += 2
		i++
	}

	__gx.peZMode = (__gx.peZMode & (^uint32(0xff000000))) | (_SHIFTL(0x40, 24, 8))
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0xff000000))) | (_SHIFTL(0x41, 24, 8))
	__gx.peCMode1 = (__gx.peCMode1 & (^uint32(0xff000000))) | (_SHIFTL(0x42, 24, 8))
	__gx.peCntrl = (__gx.peCntrl & (^uint32(0xff000000))) | (_SHIFTL(0x43, 24, 8))

	i = 0
	re0 = 0x25
	for i < 11 {
		__gx.tevRasOrder[i] = (__gx.tevRasOrder[i] & (^uint32(0xff000000))) | (_SHIFTL(re0, 24, 8))
		re0++
		i++
	}

	i = 0
	re0 = 0xf6
	for i < 8 {
		__gx.tevSwapModeTable[i] = (__gx.tevSwapModeTable[i] & (^uint32(0xff000000))) | (_SHIFTL(re0, 24, 8))
		re0++
		i++
	}

	__gx.tevTexCoordEnable = 0
	__gx.cpPerfMode = 0

	i = 0
	for i < 16 {
		__gx.tevTexMap[i] = 0xff
		i++
	}

	__GX_InitGX()
}

func __GX_InitGX() {
	black := GXColor{0, 0, 0, 0}
	white := GXColor{255, 255, 255, 255}

	GX_SetNumTexGens(1)

	GX_SetCoPlanar(GX_DISABLE)
	GX_SetCullMode(GX_CULL_BACK)
	//GX_SetClipMode(GX_CLIP_ENABLE)

	GX_SetNumChans(0)

	GX_SetChanCtrl(GX_COLOR0A0, GX_DISABLE, GX_SRC_REG, GX_SRC_VTX, GX_LIGHTNULL, GX_DF_NONE, GX_AF_NONE)
	GX_SetChanAmbColor(GX_COLOR0A0, black)
	GX_SetChanMatColor(GX_COLOR0A0, white)

	GX_SetChanCtrl(GX_COLOR1A1, GX_DISABLE, GX_SRC_REG, GX_SRC_VTX, GX_LIGHTNULL, GX_DF_NONE, GX_AF_NONE)
	GX_SetChanAmbColor(GX_COLOR1A1, black)
	GX_SetChanMatColor(GX_COLOR1A1, white)

	GX_SetTevOrder(GX_TEVSTAGE0, GX_TEXCOORD0, GX_TEXMAP0, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE1, GX_TEXCOORD1, GX_TEXMAP1, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE2, GX_TEXCOORD2, GX_TEXMAP2, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE3, GX_TEXCOORD3, GX_TEXMAP3, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE4, GX_TEXCOORD4, GX_TEXMAP4, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE5, GX_TEXCOORD5, GX_TEXMAP5, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE6, GX_TEXCOORD6, GX_TEXMAP6, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE7, GX_TEXCOORD7, GX_TEXMAP7, GX_COLOR0A0)
	GX_SetTevOrder(GX_TEVSTAGE8, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE9, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE10, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE11, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE12, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE13, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE14, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetTevOrder(GX_TEVSTAGE15, GX_TEXCOORDNULL, GX_TEXMAP_NULL, GX_COLORNULL)
	GX_SetNumTevStages(1)
	GX_SetTevOp(GX_TEVSTAGE0, GX_REPLACE)
	GX_SetAlphaCompare(GX_ALWAYS, 0, GX_AOP_AND, GX_ALWAYS, 0)
	GX_SetZTexture(GX_ZT_DISABLE, GX_TF_Z8, 0)
	for i := uint8(0); i < GX_MAX_TEVSTAGE; i++ {
		GX_SetTevKColorSel(i, GX_TEV_KCSEL_1_4)
		GX_SetTevKAlphaSel(i, GX_TEV_KASEL_1)
		GX_SetTevSwapMode(i, GX_TEV_SWAP0, GX_TEV_SWAP0)
	}

	GX_SetTevSwapModeTable(GX_TEV_SWAP0, GX_CH_RED, GX_CH_GREEN, GX_CH_BLUE, GX_CH_ALPHA)
	GX_SetTevSwapModeTable(GX_TEV_SWAP1, GX_CH_RED, GX_CH_RED, GX_CH_RED, GX_CH_ALPHA)
	GX_SetTevSwapModeTable(GX_TEV_SWAP2, GX_CH_GREEN, GX_CH_GREEN, GX_CH_GREEN, GX_CH_ALPHA)
	GX_SetTevSwapModeTable(GX_TEV_SWAP3, GX_CH_BLUE, GX_CH_BLUE, GX_CH_BLUE, GX_CH_ALPHA)
	for i := uint8(0); i < GX_MAX_TEVSTAGE; i++ {
		GX_SetTevDirect(i)
	}

	GX_SetNumIndStages(0)
	GX_SetIndTexCoordScale(GX_INDTEXSTAGE0, GX_ITS_1, GX_ITS_1)
	GX_SetIndTexCoordScale(GX_INDTEXSTAGE1, GX_ITS_1, GX_ITS_1)
	GX_SetIndTexCoordScale(GX_INDTEXSTAGE2, GX_ITS_1, GX_ITS_1)
	GX_SetIndTexCoordScale(GX_INDTEXSTAGE3, GX_ITS_1, GX_ITS_1)

	//GX_SetFog(GX_FOG_NONE, 0, 1, 0.1, 1, black)
	//GX_SetFogRangeAdj(GX_DISABLE, 0, NULL)

	GX_SetBlendMode(GX_BM_NONE, GX_BL_SRCALPHA, GX_BL_INVSRCALPHA, GX_LO_CLEAR)
	GX_SetColorUpdate(GX_ENABLE)
	GX_SetAlphaUpdate(GX_ENABLE)
	GX_SetZMode(GX_ENABLE, GX_LEQUAL, GX_TRUE)
	GX_SetZCompLoc(GX_TRUE)
	//GX_SetDither(GX_ENABLE)
	//GX_SetDstAlpha(GX_DISABLE, 0)

	//GX_PokeColorUpdate(GX_TRUE)
	//GX_PokeAlphaUpdate(GX_TRUE)
	//GX_PokeDither(GX_FALSE)
	//GX_PokeBlendMode(GX_BM_NONE, GX_BL_ZERO, GX_BL_ONE, GX_LO_SET)
	//GX_PokeAlphaMode(GX_ALWAYS, 0)
	//GX_PokeAlphaRead(GX_READ_FF)
	//GX_PokeDstAlpha(GX_DISABLE, 0)
	//GX_PokeZMode(GX_TRUE, GX_ALWAYS, GX_TRUE)

	//GX_SetGPMetric(GX_PERF0_NONE, GX_PERF1_NONE)
	//GX_ClearGPMetric()
}

func GX_SetCullMode(mode uint8) {
	cm2hw := [4]uint8{0, 2, 1, 3}

	__gx.genMode = (__gx.genMode & (^uint32(0xC000))) | (_SHIFTL(uint32(cm2hw[mode]), 14, 2))
	__gx.dirtyState |= 0x0004
}

func GX_SetCoPlanar(enable uint8) {
	__gx.genMode = (__gx.genMode & (^uint32(0x80000))) | (_SHIFTL(uint32(enable), 19, 1))
	GX_LOAD_BP_REG(0xFE080000)
	GX_LOAD_BP_REG(__gx.genMode)
}

func GX_SetChanAmbColor(channel uint8, color GXColor) {
	var reg uint32
	val := (_SHIFTL(uint32(color.r), 24, 8)) | (_SHIFTL(uint32(color.g), 16, 8)) | (_SHIFTL(uint32(color.b), 8, 8)) | 0
	switch channel {
	case GX_COLOR0:
		reg = 0
		val |= (__gx.chnAmbColor[0] & 0xff)
		break
	case GX_COLOR1:
		reg = 1
		val |= (__gx.chnAmbColor[1] & 0xff)
		break
	case GX_ALPHA0:
		reg = 0
		val = ((__gx.chnAmbColor[0] & (^uint32(0xff))) | uint32(color.a&0xff))
		break
	case GX_ALPHA1:
		reg = 1
		val = ((__gx.chnAmbColor[1] & (^uint32(0xff))) | uint32(color.a&0xff))
		break
	case GX_COLOR0A0:
		reg = 0
		val |= uint32(color.a & 0xff)
		break
	case GX_COLOR1A1:
		reg = 1
		val |= uint32(color.a & 0xff)
		break
	default:
		return
	}

	__gx.chnAmbColor[reg] = val
	__gx.dirtyState |= (0x0100 << reg)
}

func GX_SetChanMatColor(channel uint8, color GXColor) {
	var reg uint32
	val := (_SHIFTL(uint32(color.r), 24, 8)) | (_SHIFTL(uint32(color.g), 16, 8)) | (_SHIFTL(uint32(color.b), 8, 8)) | 0x00
	switch channel {
	case GX_COLOR0:
		reg = 0
		val |= (__gx.chnMatColor[0] & 0xff)
		break
	case GX_COLOR1:
		reg = 1
		val |= (__gx.chnMatColor[1] & 0xff)
		break
	case GX_ALPHA0:
		reg = 0
		val = ((__gx.chnMatColor[0] & (^uint32(0xff))) | uint32(color.a&0xff))
		break
	case GX_ALPHA1:
		reg = 1
		val = ((__gx.chnMatColor[1] & (^uint32(0xff))) | uint32(color.a&0xff))
		break
	case GX_COLOR0A0:
		reg = 0
		val |= uint32(color.a & 0xff)
		break
	case GX_COLOR1A1:
		reg = 1
		val |= uint32(color.a & 0xff)
		break
	default:
		return
	}

	__gx.chnMatColor[reg] = val
	__gx.dirtyState |= (0x0400 << reg)
}

func GX_SetChanCtrl(channel, enable, ambsrc, matsrc, litmask, diff_fn, attn_fn uint8) {
	difffn := diff_fn
	if attn_fn == GX_AF_SPEC {
		difffn = GX_DF_NONE
	}
	// (attn_fn > 0)
	attnpos := uint32(0)
	if attn_fn > 0 {
		attnpos = 1
	}
	// ((GX_AF_NONE - attn_fn) > 0)
	attnpos2 := uint32(0)
	if (GX_AF_NONE - attn_fn) > 0 {
		attnpos2 = 1
	}
	val := uint32(matsrc&1) | (_SHIFTL(uint32(enable), 1, 1)) | (_SHIFTL(uint32(litmask), 2, 4)) | (_SHIFTL(uint32(ambsrc), 6, 1)) | (_SHIFTL(uint32(difffn), 7, 2)) | (_SHIFTL(attnpos2, 9, 1)) | (_SHIFTL(attnpos, 10, 1)) | (_SHIFTL((_SHIFTR(uint32(litmask), 4, 4)), 11, 4))

	reg := uint32(channel & 0x03)
	__gx.chnCntrl[reg] = val
	__gx.dirtyState |= (0x1000 << reg)

	if channel == GX_COLOR0A0 {
		__gx.chnCntrl[2] = val
		__gx.dirtyState |= 0x5000
	} else {
		__gx.chnCntrl[3] = val
		__gx.dirtyState |= 0xa000
	}
}

func GX_SetTevOrder(tevstage, texcoord uint8, texmap uint32, color uint8) {
	reg := uint32(3 + (_SHIFTR(uint32(tevstage), 1, 3)))

	__gx.tevTexMap[tevstage&0xf] = texmap

	var texc uint32
	texm := texmap & (^uint32(0x100))
	if texm >= GX_MAX_TEXMAP {
		texm = 0
	}
	if texcoord >= GX_MAXCOORD {
		texc = 0
		__gx.tevTexCoordEnable &= ^(_SHIFTL(1, uint32(tevstage), 1))
	} else {
		texc = uint32(texcoord)
		__gx.tevTexCoordEnable |= (_SHIFTL(1, uint32(tevstage), 1))
	}

	if tevstage&1 != 0 {
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x7000))) | (_SHIFTL(texm, 12, 3))
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x38000))) | (_SHIFTL(texc, 15, 3))

		colid := GX_ALPHA_BUMP
		if color != GX_COLORNULL {
			colid = _gxtevcolid[color]
		}
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x380000))) | (_SHIFTL(uint32(colid), 19, 3))

		tmp := uint32(1)
		if texmap == GX_TEXMAP_NULL || texmap&0x100 != 0 {
			tmp = 0
		}
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x40000))) | (_SHIFTL(tmp, 18, 1))
	} else {
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x7))) | (texm & 0x7)
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x38))) | (_SHIFTL(texc, 3, 3))

		colid := GX_ALPHA_BUMP
		if color != GX_COLORNULL {
			colid = _gxtevcolid[color]
		}
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x380))) | (_SHIFTL(uint32(colid), 7, 3))

		tmp := uint32(1)
		if texmap == GX_TEXMAP_NULL || texmap&0x100 != 0 {
			tmp = 0
		}
		__gx.tevRasOrder[reg] = (__gx.tevRasOrder[reg] & (^uint32(0x40))) | (_SHIFTL(tmp, 6, 1))
	}
	GX_LOAD_BP_REG(__gx.tevRasOrder[reg])
	__gx.dirtyState |= 0x0001
}

func GX_SetNumTevStages(num uint8) {
	__gx.genMode = (__gx.genMode & (^uint32(0x3c00))) | (_SHIFTL(uint32(num-1), 10, 4))
	__gx.dirtyState |= 0x0004
}

func GX_SetNumChans(num uint8) {
	__gx.genMode = (__gx.genMode & (^uint32(0x70))) | (_SHIFTL(uint32(num), 4, 3))
	__gx.dirtyState |= 0x01000004
}

func GX_SetNumTexGens(nr uint32) {
	__gx.genMode = (__gx.genMode & (^uint32(0xf))) | (nr & 0xf)
	__gx.dirtyState |= 0x02000004
}

func GX_SetTevSwapModeTable(swapid, r, g, b, a uint8) {
	regA := _SHIFTL(uint32(swapid), 1, 3)
	regB := regA + 1

	__gx.tevSwapModeTable[regA] = (__gx.tevSwapModeTable[regA] & (^uint32(0x3))) | uint32(r&0x3)
	__gx.tevSwapModeTable[regA] = (__gx.tevSwapModeTable[regA] & (^uint32(0xC))) | (_SHIFTL(uint32(g), 2, 2))
	GX_LOAD_BP_REG(__gx.tevSwapModeTable[regA])

	__gx.tevSwapModeTable[regB] = (__gx.tevSwapModeTable[regB] & (^uint32(0x3))) | uint32(b&0x3)
	__gx.tevSwapModeTable[regB] = (__gx.tevSwapModeTable[regB] & (^uint32(0xC))) | (_SHIFTL(uint32(a), 2, 2))
	GX_LOAD_BP_REG(__gx.tevSwapModeTable[regB])
}

func GX_SetTevIndirect(tevstage, indtexid, format, bias, mtxid, wrap_s, wrap_t, addprev, utclod, a uint8) {
	val := (0x10000000 | (_SHIFTL(uint32(tevstage), 24, 4))) | uint32(indtexid&3) | (_SHIFTL(uint32(format), 2, 2)) | (_SHIFTL(uint32(bias), 4, 3)) | (_SHIFTL(uint32(a), 7, 2)) | (_SHIFTL(uint32(mtxid), 9, 4)) | (_SHIFTL(uint32(wrap_s), 13, 3)) | (_SHIFTL(uint32(wrap_t), 16, 3)) | (_SHIFTL(uint32(utclod), 19, 1)) | (_SHIFTL(uint32(addprev), 20, 1))
	GX_LOAD_BP_REG(val)
}

func GX_SetTevDirect(tevstage uint8) {
	GX_SetTevIndirect(tevstage, GX_INDTEXSTAGE0, GX_ITF_8, GX_ITB_NONE, GX_ITM_OFF, GX_ITW_OFF, GX_ITW_OFF, GX_FALSE, GX_FALSE, GX_ITBA_OFF)
}

func GX_SetNumIndStages(nstages uint8) {
	__gx.genMode = (__gx.genMode & (^uint32(0x70000))) | (_SHIFTL(uint32(nstages), 16, 3))
	__gx.dirtyState |= 0x0006
}

func GX_SetIndTexCoordScale(indtexid, scale_s, scale_t uint8) {
	switch indtexid {
	case GX_INDTEXSTAGE0:
		__gx.tevRasOrder[0] = (__gx.tevRasOrder[0] & (^uint32(0x0f))) | uint32(scale_s&0x0f)
		__gx.tevRasOrder[0] = (__gx.tevRasOrder[0] & (^uint32(0xF0))) | (_SHIFTL(uint32(scale_t), 4, 4))
		GX_LOAD_BP_REG(__gx.tevRasOrder[0])
		break
	case GX_INDTEXSTAGE1:
		__gx.tevRasOrder[0] = (__gx.tevRasOrder[0] & (^uint32(0xF00))) | (_SHIFTL(uint32(scale_s), 8, 4))
		__gx.tevRasOrder[0] = (__gx.tevRasOrder[0] & (^uint32(0xF000))) | (_SHIFTL(uint32(scale_t), 12, 4))
		GX_LOAD_BP_REG(__gx.tevRasOrder[0])
		break
	case GX_INDTEXSTAGE2:
		__gx.tevRasOrder[1] = (__gx.tevRasOrder[1] & (^uint32(0x0f))) | uint32(scale_s&0x0f)
		__gx.tevRasOrder[1] = (__gx.tevRasOrder[1] & (^uint32(0xF0))) | (_SHIFTL(uint32(scale_t), 4, 4))
		GX_LOAD_BP_REG(__gx.tevRasOrder[1])
		break
	case GX_INDTEXSTAGE3:
		__gx.tevRasOrder[1] = (__gx.tevRasOrder[1] & (^uint32(0xF00))) | (_SHIFTL(uint32(scale_s), 8, 4))
		__gx.tevRasOrder[1] = (__gx.tevRasOrder[1] & (^uint32(0xF000))) | (_SHIFTL(uint32(scale_t), 12, 4))
		GX_LOAD_BP_REG(__gx.tevRasOrder[1])
		break
	}
}

func GX_SetZCompLoc(before_tex uint8) {
	__gx.peCntrl = (__gx.peCntrl & (^uint32(0x40))) | (_SHIFTL(uint32(before_tex), 6, 1))
	GX_LOAD_BP_REG(__gx.peCntrl)
}

func GX_SetTevColorIn(tevstage, a, b, c, d uint8) {
	reg := tevstage & 0xf
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0xF000))) | (_SHIFTL(uint32(a), 12, 4))
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0xF00))) | (_SHIFTL(uint32(b), 8, 4))
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0xF0))) | (_SHIFTL(uint32(c), 4, 4))
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0xf))) | uint32(d&0xf)

	GX_LOAD_BP_REG(__gx.tevColorEnv[reg])
}

func GX_SetTevAlphaIn(tevstage, a, b, c, d uint8) {
	reg := tevstage & 0xf
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0xE000))) | (_SHIFTL(uint32(a), 13, 3))
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x1C00))) | (_SHIFTL(uint32(b), 10, 3))
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x380))) | (_SHIFTL(uint32(c), 7, 3))
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x70))) | (_SHIFTL(uint32(d), 4, 3))

	GX_LOAD_BP_REG(__gx.tevAlphaEnv[reg])
}

func GX_SetTevOp(tevstage, mode uint8) {
	defcolor := GX_CC_RASC
	defalpha := GX_CA_RASA

	if tevstage != GX_TEVSTAGE0 {
		defcolor = GX_CC_CPREV
		defalpha = GX_CA_APREV
	}

	switch mode {
	case GX_MODULATE:
		GX_SetTevColorIn(tevstage, GX_CC_ZERO, GX_CC_TEXC, defcolor, GX_CC_ZERO)
		GX_SetTevAlphaIn(tevstage, GX_CA_ZERO, GX_CA_TEXA, defalpha, GX_CA_ZERO)
	case GX_DECAL:
		GX_SetTevColorIn(tevstage, defcolor, GX_CC_TEXC, GX_CC_TEXA, GX_CC_ZERO)
		GX_SetTevAlphaIn(tevstage, GX_CA_ZERO, GX_CA_ZERO, GX_CA_ZERO, defalpha)
	case GX_BLEND:
		GX_SetTevColorIn(tevstage, defcolor, GX_CC_ONE, GX_CC_TEXC, GX_CC_ZERO)
		GX_SetTevAlphaIn(tevstage, GX_CA_ZERO, GX_CA_TEXA, defalpha, GX_CA_RASA)
	case GX_REPLACE:
		GX_SetTevColorIn(tevstage, GX_CC_ZERO, GX_CC_ZERO, GX_CC_ZERO, GX_CC_TEXC)
		GX_SetTevAlphaIn(tevstage, GX_CA_ZERO, GX_CA_ZERO, GX_CA_ZERO, GX_CA_TEXA)
	case GX_PASSCLR:
		GX_SetTevColorIn(tevstage, GX_CC_ZERO, GX_CC_ZERO, GX_CC_ZERO, defcolor)
		GX_SetTevAlphaIn(tevstage, GX_CC_A2, GX_CC_A2, GX_CC_A2, defalpha)
	}
	GX_SetTevColorOp(tevstage, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_TRUE, GX_TEVPREV)
	GX_SetTevAlphaOp(tevstage, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_TRUE, GX_TEVPREV)
}

func GX_SetTevColorOp(tevstage, tevop, tevbias, tevscale, clamp, tevregid uint8) {
	reg := tevstage & 0xf
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x40000))) | (_SHIFTL(uint32(tevop), 18, 1))
	if tevop <= GX_TEV_SUB {
		__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x300000))) | (_SHIFTL(uint32(tevscale), 20, 2))
		__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x30000))) | (_SHIFTL(uint32(tevbias), 16, 2))
	} else {
		__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x300000))) | ((_SHIFTL(uint32(tevop), 19, 4)) & 0x300000)
		__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x30000))) | 0x30000
	}
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0x80000))) | (_SHIFTL(uint32(clamp), 19, 1))
	__gx.tevColorEnv[reg] = (__gx.tevColorEnv[reg] & (^uint32(0xC00000))) | (_SHIFTL(uint32(tevregid), 22, 2))

	GX_LOAD_BP_REG(__gx.tevColorEnv[reg])
}

func GX_SetTevAlphaOp(tevstage, tevop, tevbias, tevscale, clamp, tevregid uint8) {
	reg := tevstage & 0xf
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x40000))) | (_SHIFTL(uint32(tevop), 18, 1))
	if tevop <= GX_TEV_SUB {
		__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x300000))) | (_SHIFTL(uint32(tevscale), 20, 2))
		__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x30000))) | (_SHIFTL(uint32(tevbias), 16, 2))
	} else {
		__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x300000))) | ((_SHIFTL(uint32(tevop), 19, 4)) & 0x300000)
		__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x30000))) | 0x30000
	}
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x80000))) | (_SHIFTL(uint32(clamp), 19, 1))
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0xC00000))) | (_SHIFTL(uint32(tevregid), 22, 2))

	GX_LOAD_BP_REG(__gx.tevAlphaEnv[reg])
}

func GX_SetTevKColorSel(tevstage, sel uint8) {
	reg := _SHIFTR(uint32(tevstage), 1, 3)

	if tevstage&1 != 0 {
		__gx.tevSwapModeTable[reg] = (__gx.tevSwapModeTable[reg] & (^uint32(0x7C000))) | (_SHIFTL(uint32(sel), 14, 5))
	} else {
		__gx.tevSwapModeTable[reg] = (__gx.tevSwapModeTable[reg] & (^uint32(0x1F0))) | (_SHIFTL(uint32(sel), 4, 5))
	}
	GX_LOAD_BP_REG(__gx.tevSwapModeTable[reg])
}

func GX_SetTevKAlphaSel(tevstage, sel uint8) {
	reg := _SHIFTR(uint32(tevstage), 1, 3)

	if tevstage&1 != 0 {
		__gx.tevSwapModeTable[reg] = (__gx.tevSwapModeTable[reg] & (^uint32(0xF80000))) | (_SHIFTL(uint32(sel), 19, 5))
	} else {
		__gx.tevSwapModeTable[reg] = (__gx.tevSwapModeTable[reg] & (^uint32(0x3E00))) | (_SHIFTL(uint32(sel), 9, 5))
	}
	GX_LOAD_BP_REG(__gx.tevSwapModeTable[reg])
}

func GX_SetTevSwapMode(tevstage, ras_sel, tex_sel uint8) {
	reg := tevstage & 0xf
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0x3))) | uint32(ras_sel&0x3)
	__gx.tevAlphaEnv[reg] = (__gx.tevAlphaEnv[reg] & (^uint32(0xC))) | (_SHIFTL(uint32(tex_sel), 2, 2))
	GX_LOAD_BP_REG(__gx.tevAlphaEnv[reg])
}

func GX_SetZTexture(op, fmt uint8, bias uint32) {
	if fmt == GX_TF_Z8 {
		fmt = 0
	} else if fmt == GX_TF_Z16 {
		fmt = 1
	} else {
		fmt = 2
	}

	val := (_SHIFTL(uint32(op), 2, 2)) | uint32(fmt&3)
	GX_LOAD_BP_REG(0xF4000000 | (bias & 0x00FFFFFF))
	GX_LOAD_BP_REG(0xF5000000 | (val & 0x00FFFFFF))
}

func GX_SetAlphaCompare(comp0, ref0, aop, comp1, ref1 uint8) {
	val := (_SHIFTL(uint32(aop), 22, 2)) | (_SHIFTL(uint32(comp1), 19, 3)) | (_SHIFTL(uint32(comp0), 16, 3)) | (_SHIFTL(uint32(ref1), 8, 8)) | uint32(ref0)
	GX_LOAD_BP_REG(0xf3000000 | val)
}

func GX_SetBlendMode(btype, src_fact, dst_fact, op uint8) {
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x1)))
	if btype == GX_BM_BLEND || btype == GX_BM_SUBTRACT {
		__gx.peCMode0 |= 0x1
	}

	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x800)))
	if btype == GX_BM_SUBTRACT {
		__gx.peCMode0 |= 0x800
	}

	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x2)))
	if btype == GX_BM_LOGIC {
		__gx.peCMode0 |= 0x2
	}

	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0xF000))) | (_SHIFTL(uint32(op), 12, 4))
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0xE0))) | (_SHIFTL(uint32(dst_fact), 5, 3))
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x700))) | (_SHIFTL(uint32(src_fact), 8, 3))

	GX_LOAD_BP_REG(__gx.peCMode0)
}

func GX_SetColorUpdate(enable uint8) {
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x8))) | (_SHIFTL(uint32(enable), 3, 1))
	GX_LOAD_BP_REG(__gx.peCMode0)
}

func GX_SetAlphaUpdate(enable uint8) {
	__gx.peCMode0 = (__gx.peCMode0 & (^uint32(0x10))) | (_SHIFTL(uint32(enable), 4, 1))
	GX_LOAD_BP_REG(__gx.peCMode0)
}

func GX_SetZMode(enable, zfunc, update_enable uint8) {
	__gx.peZMode = (__gx.peZMode & (^uint32(0x1))) | uint32(enable&1)
	__gx.peZMode = (__gx.peZMode & (^uint32(0xe))) | (_SHIFTL(uint32(zfunc), 1, 3))
	__gx.peZMode = (__gx.peZMode & (^uint32(0x10))) | (_SHIFTL(uint32(update_enable), 4, 1))
	GX_LOAD_BP_REG(__gx.peZMode)
}

func GX_Flush() {
	if __gx.dirtyState != 0 {
		__GX_SetDirtyState()
	}

	// Pad to 32B boundary
	wgPipe.Pad32()
}

// This is nasty!
var gxfunctions = map[string]interface{}{
	"GX_SetCullMode":         GX_SetCullMode,
	"GX_SetCoPlanar":         GX_SetCoPlanar,
	"GX_SetChanAmbColor":     GX_SetChanAmbColor,
	"GX_SetChanMatColor":     GX_SetChanMatColor,
	"GX_SetChanCtrl":         GX_SetChanCtrl,
	"GX_SetTevOrder":         GX_SetTevOrder,
	"GX_SetNumTevStages":     GX_SetNumTevStages,
	"GX_SetNumChans":         GX_SetNumChans,
	"GX_SetNumTexGens":       GX_SetNumTexGens,
	"GX_SetTevSwapModeTable": GX_SetTevSwapModeTable,
	"GX_SetTevIndirect":      GX_SetTevIndirect,
	"GX_SetTevDirect":        GX_SetTevDirect,
	"GX_SetNumIndStages":     GX_SetNumIndStages,
	"GX_SetIndTexCoordScale": GX_SetIndTexCoordScale,
	"GX_SetZCompLoc":         GX_SetZCompLoc,
	"GX_SetTevColorIn":       GX_SetTevColorIn,
	"GX_SetTevAlphaIn":       GX_SetTevAlphaIn,
	"GX_SetTevOp":            GX_SetTevOp,
	"GX_SetTevColorOp":       GX_SetTevColorOp,
	"GX_SetTevAlphaOp":       GX_SetTevAlphaOp,
	"GX_SetTevKColorSel":     GX_SetTevKColorSel,
	"GX_SetTevKAlphaSel":     GX_SetTevKAlphaSel,
	"GX_SetTevSwapMode":      GX_SetTevSwapMode,
	"GX_SetZTexture":         GX_SetZTexture,
	"GX_SetAlphaCompare":     GX_SetAlphaCompare,
	"GX_SetBlendMode":        GX_SetBlendMode,
	"GX_SetColorUpdate":      GX_SetColorUpdate,
	"GX_SetAlphaUpdate":      GX_SetAlphaUpdate,
	"GX_SetZMode":            GX_SetZMode,
	"GX_Flush":               GX_Flush,
}
