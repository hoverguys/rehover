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

	GX_MODULATE uint8 = 0
	GX_DECAL    uint8 = 1
	GX_BLEND    uint8 = 2
	GX_REPLACE  uint8 = 3
	GX_PASSCLR  uint8 = 4

	GX_CH_RED   uint8 = 0
	GX_CH_GREEN uint8 = 1
	GX_CH_BLUE  uint8 = 2
	GX_CH_ALPHA uint8 = 3

	GX_BM_NONE       uint8 = 0
	GX_BM_BLEND      uint8 = 1
	GX_BM_LOGIC      uint8 = 2
	GX_BM_SUBTRACT   uint8 = 3
	GX_MAX_BLENDMODE uint8 = 4

	GX_BL_ZERO        uint8 = 0
	GX_BL_ONE         uint8 = 1
	GX_BL_SRCCLR      uint8 = 2
	GX_BL_INVSRCCLR   uint8 = 3
	GX_BL_SRCALPHA    uint8 = 4
	GX_BL_INVSRCALPHA uint8 = 5
	GX_BL_DSTALPHA    uint8 = 6
	GX_BL_INVDSTALPHA uint8 = 7
	GX_BL_DSTCLR      uint8 = GX_BL_SRCCLR
	GX_BL_INVDSTCLR   uint8 = GX_BL_INVSRCCLR

	GX_LO_CLEAR   uint8 = 0
	GX_LO_AND     uint8 = 1
	GX_LO_REVAND  uint8 = 2
	GX_LO_COPY    uint8 = 3
	GX_LO_INVAND  uint8 = 4
	GX_LO_NOOP    uint8 = 5
	GX_LO_XOR     uint8 = 6
	GX_LO_OR      uint8 = 7
	GX_LO_NOR     uint8 = 8
	GX_LO_EQUIV   uint8 = 9
	GX_LO_INV     uint8 = 10
	GX_LO_REVOR   uint8 = 11
	GX_LO_INVCOPY uint8 = 12
	GX_LO_INVOR   uint8 = 13
	GX_LO_NAND    uint8 = 14
	GX_LO_SET     uint8 = 15

	GX_CULL_ALL   uint8 = 3
	GX_CULL_BACK  uint8 = 2
	GX_CULL_FRONT uint8 = 1
	GX_CULL_NONE  uint8 = 0
)
