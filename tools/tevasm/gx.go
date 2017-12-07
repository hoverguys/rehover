package main

const (
	GX_FALSE   uint8 = 0
	GX_TRUE    uint8 = 1
	GX_DISABLE uint8 = 0
	GX_ENABLE  uint8 = 1

	GX_COLOR0      uint8 = 0
	GX_COLOR1      uint8 = 1
	GX_ALPHA0      uint8 = 2
	GX_ALPHA1      uint8 = 3
	GX_COLOR0A0    uint8 = 4
	GX_COLOR1A1    uint8 = 5
	GX_COLORZERO   uint8 = 6
	GX_ALPHA_BUMP  uint8 = 7
	GX_ALPHA_BUMPN uint8 = 8
	GX_COLORNULL   uint8 = 0xff

	GX_INDTEXSTAGE0    uint8 = 0
	GX_INDTEXSTAGE1    uint8 = 1
	GX_INDTEXSTAGE2    uint8 = 2
	GX_INDTEXSTAGE3    uint8 = 3
	GX_MAX_INDTEXSTAGE uint8 = 4

	GX_ITF_8        uint8 = 0
	GX_ITF_5        uint8 = 1
	GX_ITF_4        uint8 = 2
	GX_ITF_3        uint8 = 3
	GX_MAX_ITFORMAT uint8 = 4

	GX_ITB_NONE   uint8 = 0
	GX_ITB_S      uint8 = 1
	GX_ITB_T      uint8 = 2
	GX_ITB_ST     uint8 = 3
	GX_ITB_U      uint8 = 4
	GX_ITB_SU     uint8 = 5
	GX_ITB_TU     uint8 = 6
	GX_ITB_STU    uint8 = 7
	GX_MAX_ITBIAS uint8 = 8

	GX_ITM_OFF uint8 = 0
	GX_ITM_0   uint8 = 1
	GX_ITM_1   uint8 = 2
	GX_ITM_2   uint8 = 3
	GX_ITM_S0  uint8 = 5
	GX_ITM_S1  uint8 = 6
	GX_ITM_S2  uint8 = 7
	GX_ITM_T0  uint8 = 9
	GX_ITM_T1  uint8 = 10
	GX_ITM_T2  uint8 = 11

	GX_ITW_OFF    uint8 = 0
	GX_ITW_256    uint8 = 1
	GX_ITW_128    uint8 = 2
	GX_ITW_64     uint8 = 3
	GX_ITW_32     uint8 = 4
	GX_ITW_16     uint8 = 5
	GX_ITW_0      uint8 = 6
	GX_MAX_ITWRAP uint8 = 7

	GX_ITBA_OFF     uint8 = 0
	GX_ITBA_S       uint8 = 1
	GX_ITBA_T       uint8 = 2
	GX_ITBA_U       uint8 = 3
	GX_MAX_ITBALPHA uint8 = 4

	GX_ITS_1       uint8 = 0
	GX_ITS_2       uint8 = 1
	GX_ITS_4       uint8 = 2
	GX_ITS_8       uint8 = 3
	GX_ITS_16      uint8 = 4
	GX_ITS_32      uint8 = 5
	GX_ITS_64      uint8 = 6
	GX_ITS_128     uint8 = 7
	GX_ITS_256     uint8 = 8
	GX_MAX_ITSCALE uint8 = 9

	GX_TEVSTAGE0    uint8 = 0
	GX_TEVSTAGE1    uint8 = 1
	GX_TEVSTAGE2    uint8 = 2
	GX_TEVSTAGE3    uint8 = 3
	GX_TEVSTAGE4    uint8 = 4
	GX_TEVSTAGE5    uint8 = 5
	GX_TEVSTAGE6    uint8 = 6
	GX_TEVSTAGE7    uint8 = 7
	GX_TEVSTAGE8    uint8 = 8
	GX_TEVSTAGE9    uint8 = 9
	GX_TEVSTAGE10   uint8 = 10
	GX_TEVSTAGE11   uint8 = 11
	GX_TEVSTAGE12   uint8 = 12
	GX_TEVSTAGE13   uint8 = 13
	GX_TEVSTAGE14   uint8 = 14
	GX_TEVSTAGE15   uint8 = 15
	GX_MAX_TEVSTAGE uint8 = 16

	GX_AF_SPEC uint8 = 0
	GX_AF_SPOT uint8 = 1
	GX_AF_NONE uint8 = 2

	GX_DF_NONE   uint8 = 0
	GX_DF_SIGNED uint8 = 1
	GX_DF_CLAMP  uint8 = 2

	GX_SRC_REG uint8 = 0
	GX_SRC_VTX uint8 = 1

	GX_LIGHT0    uint8 = 0x001
	GX_LIGHT1    uint8 = 0x002
	GX_LIGHT2    uint8 = 0x004
	GX_LIGHT3    uint8 = 0x008
	GX_LIGHT4    uint8 = 0x010
	GX_LIGHT5    uint8 = 0x020
	GX_LIGHT6    uint8 = 0x040
	GX_LIGHT7    uint8 = 0x080
	GX_LIGHTNULL uint8 = 0x000

	GX_CC_CPREV uint8 = 0
	GX_CC_APREV uint8 = 1
	GX_CC_C0    uint8 = 2
	GX_CC_A0    uint8 = 3
	GX_CC_C1    uint8 = 4
	GX_CC_A1    uint8 = 5
	GX_CC_C2    uint8 = 6
	GX_CC_A2    uint8 = 7
	GX_CC_TEXC  uint8 = 8
	GX_CC_TEXA  uint8 = 9
	GX_CC_RASC  uint8 = 10
	GX_CC_RASA  uint8 = 11
	GX_CC_ONE   uint8 = 12
	GX_CC_HALF  uint8 = 13
	GX_CC_KONST uint8 = 14
	GX_CC_ZERO  uint8 = 15

	GX_CA_APREV uint8 = 0
	GX_CA_A0    uint8 = 1
	GX_CA_A1    uint8 = 2
	GX_CA_A2    uint8 = 3
	GX_CA_TEXA  uint8 = 4
	GX_CA_RASA  uint8 = 5
	GX_CA_KONST uint8 = 6
	GX_CA_ZERO  uint8 = 7

	GX_TG_POS       uint8 = 0
	GX_TG_NRM       uint8 = 1
	GX_TG_BINRM     uint8 = 2
	GX_TG_TANGENT   uint8 = 3
	GX_TG_TEX0      uint8 = 4
	GX_TG_TEX1      uint8 = 5
	GX_TG_TEX2      uint8 = 6
	GX_TG_TEX3      uint8 = 7
	GX_TG_TEX4      uint8 = 8
	GX_TG_TEX5      uint8 = 9
	GX_TG_TEX6      uint8 = 10
	GX_TG_TEX7      uint8 = 11
	GX_TG_TEXCOORD0 uint8 = 12
	GX_TG_TEXCOORD1 uint8 = 13
	GX_TG_TEXCOORD2 uint8 = 14
	GX_TG_TEXCOORD3 uint8 = 15
	GX_TG_TEXCOORD4 uint8 = 16
	GX_TG_TEXCOORD5 uint8 = 17
	GX_TG_TEXCOORD6 uint8 = 18
	GX_TG_COLOR0    uint8 = 19
	GX_TG_COLOR1    uint8 = 20

	GX_TEXCOORD0    uint8 = 0
	GX_TEXCOORD1    uint8 = 1
	GX_TEXCOORD2    uint8 = 2
	GX_TEXCOORD3    uint8 = 3
	GX_TEXCOORD4    uint8 = 4
	GX_TEXCOORD5    uint8 = 5
	GX_TEXCOORD6    uint8 = 6
	GX_TEXCOORD7    uint8 = 7
	GX_MAXCOORD     uint8 = 8
	GX_TEXCOORDNULL uint8 = 0xff

	GX_TEXMAP0        uint32 = 0
	GX_TEXMAP1        uint32 = 1
	GX_TEXMAP2        uint32 = 2
	GX_TEXMAP3        uint32 = 3
	GX_TEXMAP4        uint32 = 4
	GX_TEXMAP5        uint32 = 5
	GX_TEXMAP6        uint32 = 6
	GX_TEXMAP7        uint32 = 7
	GX_MAX_TEXMAP     uint32 = 8
	GX_TEXMAP_NULL    uint32 = 0xff
	GX_TEXMAP_DISABLE uint32 = 0x100

	GX_TEV_ADD           uint8 = 0
	GX_TEV_SUB           uint8 = 1
	GX_TEV_COMP_R8_GT    uint8 = 8
	GX_TEV_COMP_R8_EQ    uint8 = 9
	GX_TEV_COMP_GR16_GT  uint8 = 10
	GX_TEV_COMP_GR16_EQ  uint8 = 11
	GX_TEV_COMP_BGR24_GT uint8 = 12
	GX_TEV_COMP_BGR24_EQ uint8 = 13
	GX_TEV_COMP_RGB8_GT  uint8 = 14
	GX_TEV_COMP_RGB8_EQ  uint8 = 15
	GX_TEV_COMP_A8_GT    uint8 = GX_TEV_COMP_RGB8_GT
	GX_TEV_COMP_A8_EQ    uint8 = GX_TEV_COMP_RGB8_EQ

	GX_TB_ZERO     uint8 = 0
	GX_TB_ADDHALF  uint8 = 1
	GX_TB_SUBHALF  uint8 = 2
	GX_MAX_TEVBIAS uint8 = 3

	GX_CS_SCALE_1   uint8 = 0
	GX_CS_SCALE_2   uint8 = 1
	GX_CS_SCALE_4   uint8 = 2
	GX_CS_DIVIDE_2  uint8 = 3
	GX_MAX_TEVSCALE uint8 = 4

	GX_TEVPREV    uint8 = 0
	GX_TEVREG0    uint8 = 1
	GX_TEVREG1    uint8 = 2
	GX_TEVREG2    uint8 = 3
	GX_MAX_TEVREG uint8 = 4

	_GX_TF_ZTF uint8 = 0x10
	_GX_TF_CTF uint8 = 0x20

	GX_TF_I4     uint8 = 0x0
	GX_TF_I8     uint8 = 0x1
	GX_TF_IA4    uint8 = 0x2
	GX_TF_IA8    uint8 = 0x3
	GX_TF_RGB565 uint8 = 0x4
	GX_TF_RGB5A3 uint8 = 0x5
	GX_TF_RGBA8  uint8 = 0x6
	GX_TF_CI4    uint8 = 0x8
	GX_TF_CI8    uint8 = 0x9
	GX_TF_CI14   uint8 = 0xa
	GX_TF_CMPR   uint8 = 0xE

	GX_TL_IA8    uint8 = 0x00
	GX_TL_RGB565 uint8 = 0x01
	GX_TL_RGB5A3 uint8 = 0x02

	GX_CTF_R4    uint8 = (0x0 | _GX_TF_CTF)
	GX_CTF_RA4   uint8 = (0x2 | _GX_TF_CTF)
	GX_CTF_RA8   uint8 = (0x3 | _GX_TF_CTF)
	GX_CTF_YUVA8 uint8 = (0x6 | _GX_TF_CTF)
	GX_CTF_A8    uint8 = (0x7 | _GX_TF_CTF)
	GX_CTF_R8    uint8 = (0x8 | _GX_TF_CTF)
	GX_CTF_G8    uint8 = (0x9 | _GX_TF_CTF)
	GX_CTF_B8    uint8 = (0xA | _GX_TF_CTF)
	GX_CTF_RG8   uint8 = (0xB | _GX_TF_CTF)
	GX_CTF_GB8   uint8 = (0xC | _GX_TF_CTF)

	GX_TF_Z8    uint8 = (0x1 | _GX_TF_ZTF)
	GX_TF_Z16   uint8 = (0x3 | _GX_TF_ZTF)
	GX_TF_Z24X8 uint8 = (0x6 | _GX_TF_ZTF)

	GX_CTF_Z4   uint8 = (0x0 | _GX_TF_ZTF | _GX_TF_CTF)
	GX_CTF_Z8M  uint8 = (0x9 | _GX_TF_ZTF | _GX_TF_CTF)
	GX_CTF_Z8L  uint8 = (0xA | _GX_TF_ZTF | _GX_TF_CTF)
	GX_CTF_Z16L uint8 = (0xC | _GX_TF_ZTF | _GX_TF_CTF)

	GX_TF_A8 uint8 = GX_CTF_A8

	GX_ZT_DISABLE uint8 = 0
	GX_ZT_ADD     uint8 = 1
	GX_ZT_REPLACE uint8 = 2
	GX_MAX_ZTEXOP uint8 = 3

	GX_NEVER   uint8 = 0
	GX_LESS    uint8 = 1
	GX_EQUAL   uint8 = 2
	GX_LEQUAL  uint8 = 3
	GX_GREATER uint8 = 4
	GX_NEQUAL  uint8 = 5
	GX_GEQUAL  uint8 = 6
	GX_ALWAYS  uint8 = 7

	GX_AOP_AND     uint8 = 0
	GX_AOP_OR      uint8 = 1
	GX_AOP_XOR     uint8 = 2
	GX_AOP_XNOR    uint8 = 3
	GX_MAX_ALPHAOP uint8 = 4

	GX_KCOLOR0    uint8 = 0
	GX_KCOLOR1    uint8 = 1
	GX_KCOLOR2    uint8 = 2
	GX_KCOLOR3    uint8 = 3
	GX_KCOLOR_MAX uint8 = 4

	GX_TEV_KCSEL_1    uint8 = 0x00
	GX_TEV_KCSEL_7_8  uint8 = 0x01
	GX_TEV_KCSEL_3_4  uint8 = 0x02
	GX_TEV_KCSEL_5_8  uint8 = 0x03
	GX_TEV_KCSEL_1_2  uint8 = 0x04
	GX_TEV_KCSEL_3_8  uint8 = 0x05
	GX_TEV_KCSEL_1_4  uint8 = 0x06
	GX_TEV_KCSEL_1_8  uint8 = 0x07
	GX_TEV_KCSEL_K0   uint8 = 0x0C
	GX_TEV_KCSEL_K1   uint8 = 0x0D
	GX_TEV_KCSEL_K2   uint8 = 0x0E
	GX_TEV_KCSEL_K3   uint8 = 0x0F
	GX_TEV_KCSEL_K0_R uint8 = 0x10
	GX_TEV_KCSEL_K1_R uint8 = 0x11
	GX_TEV_KCSEL_K2_R uint8 = 0x12
	GX_TEV_KCSEL_K3_R uint8 = 0x13
	GX_TEV_KCSEL_K0_G uint8 = 0x14
	GX_TEV_KCSEL_K1_G uint8 = 0x15
	GX_TEV_KCSEL_K2_G uint8 = 0x16
	GX_TEV_KCSEL_K3_G uint8 = 0x17
	GX_TEV_KCSEL_K0_B uint8 = 0x18
	GX_TEV_KCSEL_K1_B uint8 = 0x19
	GX_TEV_KCSEL_K2_B uint8 = 0x1A
	GX_TEV_KCSEL_K3_B uint8 = 0x1B
	GX_TEV_KCSEL_K0_A uint8 = 0x1C
	GX_TEV_KCSEL_K1_A uint8 = 0x1D
	GX_TEV_KCSEL_K2_A uint8 = 0x1E
	GX_TEV_KCSEL_K3_A uint8 = 0x1F

	GX_TEV_KASEL_1    uint8 = 0x00
	GX_TEV_KASEL_7_8  uint8 = 0x01
	GX_TEV_KASEL_3_4  uint8 = 0x02
	GX_TEV_KASEL_5_8  uint8 = 0x03
	GX_TEV_KASEL_1_2  uint8 = 0x04
	GX_TEV_KASEL_3_8  uint8 = 0x05
	GX_TEV_KASEL_1_4  uint8 = 0x06
	GX_TEV_KASEL_1_8  uint8 = 0x07
	GX_TEV_KASEL_K0_R uint8 = 0x10
	GX_TEV_KASEL_K1_R uint8 = 0x11
	GX_TEV_KASEL_K2_R uint8 = 0x12
	GX_TEV_KASEL_K3_R uint8 = 0x13
	GX_TEV_KASEL_K0_G uint8 = 0x14
	GX_TEV_KASEL_K1_G uint8 = 0x15
	GX_TEV_KASEL_K2_G uint8 = 0x16
	GX_TEV_KASEL_K3_G uint8 = 0x17
	GX_TEV_KASEL_K0_B uint8 = 0x18
	GX_TEV_KASEL_K1_B uint8 = 0x19
	GX_TEV_KASEL_K2_B uint8 = 0x1A
	GX_TEV_KASEL_K3_B uint8 = 0x1B
	GX_TEV_KASEL_K0_A uint8 = 0x1C
	GX_TEV_KASEL_K1_A uint8 = 0x1D
	GX_TEV_KASEL_K2_A uint8 = 0x1E
	GX_TEV_KASEL_K3_A uint8 = 0x1F

	GX_TEV_SWAP0   uint8 = 0
	GX_TEV_SWAP1   uint8 = 1
	GX_TEV_SWAP2   uint8 = 2
	GX_TEV_SWAP3   uint8 = 3
	GX_MAX_TEVSWAP uint8 = 4
)

type gxregdef struct {
	cpSRreg           uint16
	cpCRreg           uint16
	cpCLreg           uint16
	xfFlush           uint16
	xfFlushExp        uint16
	xfFlushSafe       uint16
	gxFifoInited      uint32
	vcdClear          uint32
	VATTable          uint32
	mtxIdxLo          uint32
	mtxIdxHi          uint32
	texCoordManually  uint32
	vcdLo             uint32
	vcdHi             uint32
	vcdNrms           uint32
	dirtyState        uint32
	perf0Mode         uint32
	perf1Mode         uint32
	cpPerfMode        uint32
	VAT0reg           [8]uint32
	VAT1reg           [8]uint32
	VAT2reg           [8]uint32
	texMapSize        [8]uint32
	texMapWrap        [8]uint32
	sciTLcorner       uint32
	sciBRcorner       uint32
	lpWidth           uint32
	genMode           uint32
	suSsize           [8]uint32
	suTsize           [8]uint32
	tevTexMap         [16]uint32
	tevColorEnv       [16]uint32
	tevAlphaEnv       [16]uint32
	tevSwapModeTable  [8]uint32
	tevRasOrder       [11]uint32
	tevTexCoordEnable uint32
	tevIndMask        uint32
	texCoordGen       [8]uint32
	texCoordGen2      [8]uint32
	dispCopyCntrl     uint32
	dispCopyDst       uint32
	dispCopyTL        uint32
	dispCopyWH        uint32
	texCopyCntrl      uint32
	texCopyDst        uint32
	texCopyTL         uint32
	texCopyWH         uint32
	peZMode           uint32
	peCMode0          uint32
	peCMode1          uint32
	peCntrl           uint32
	chnAmbColor       [2]uint32
	chnMatColor       [2]uint32
	chnCntrl          [4]uint32
	texRegion         [24]GXTexRegion
	tlutRegion        [20]GXTlutRegion
	saveDLctx         uint8
	gxFifoUnlinked    uint8
	texCopyZTex       uint8
}

type gx_texreg struct {
	val [4]uint32
}
type GXTexRegion gx_texreg

type gx_tlutreg struct {
	val [4]uint32
}
type GXTlutRegion gx_tlutreg

var __gx gxregdef
var wgPipe *FifoWriter
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
}

func GX_LOAD_XF_REG(x, y uint32) {
	wgPipe.U8(0x10)
	wgPipe.U32(x & 0xffff)
	wgPipe.U32(y)
}

func GX_LOAD_CP_REG(x uint8, y uint32) {
	wgPipe.U8(0x08)
	wgPipe.U8(x)
	wgPipe.U32(y)
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

func GX_Flush() {
	if __gx.dirtyState != 0 {
		__GX_SetDirtyState()
	}

	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
	wgPipe.U32(0)
}
