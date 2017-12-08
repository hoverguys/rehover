package main

type GXColor struct {
	r, g, b, a uint8
}

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
