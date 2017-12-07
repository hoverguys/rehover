package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	outpath := flag.String("out", "-", "Output file (- for stdout)")
	flag.Parse()

	// Get output writer
	out := os.Stdout
	if *outpath != "-" {
		file, err := os.Create(*outpath)
		checkErr(err, "Cannot create output file")
		defer file.Close()
		out = file
	}

	wgPipe = NewFifo(binary.BigEndian)

	/* BEGIN SHADER CODE */

	GX_SetChanCtrl(GX_COLOR0A0, GX_ENABLE, GX_SRC_REG, GX_SRC_REG, GX_LIGHT0, GX_DF_CLAMP, GX_AF_NONE)
	GX_SetNumTevStages(2)
	GX_SetNumChans(1)
	GX_SetNumTexGens(1)

	// No indirect stages
	GX_SetNumIndStages(0)
	GX_SetTevDirect(GX_TEVSTAGE0)
	GX_SetTevDirect(GX_TEVSTAGE1)

	// Stage 1: Multiply color with brightness map, ignore alpha
	GX_SetTevOrder(GX_TEVSTAGE0, GX_TEXCOORD0, GX_TEXMAP1, GX_COLOR0A0)
	GX_SetTevColorIn(GX_TEVSTAGE0, GX_CC_ZERO, GX_CC_RASC, GX_CC_TEXC, GX_CC_ZERO)
	GX_SetTevColorOp(GX_TEVSTAGE0, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_FALSE, GX_TEVPREV)

	// Stage 2: Add colored brightness map to global map, use alpha from global map
	GX_SetTevOrder(GX_TEVSTAGE1, GX_TEXCOORD0, GX_TEXMAP0, GX_COLORNULL)
	GX_SetTevColorIn(GX_TEVSTAGE1, GX_CC_CPREV, GX_CC_ZERO, GX_CC_ZERO, GX_CC_TEXC)
	GX_SetTevColorOp(GX_TEVSTAGE1, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_FALSE, GX_TEVPREV)
	GX_SetTevAlphaIn(GX_TEVSTAGE1, GX_CA_ZERO, GX_CA_ZERO, GX_CA_ZERO, GX_CA_TEXA)
	GX_SetTevAlphaOp(GX_TEVSTAGE1, GX_TEV_ADD, GX_TB_ZERO, GX_CS_SCALE_1, GX_FALSE, GX_TEVPREV)
	GX_SetZTexture(GX_ZT_DISABLE, GX_TF_I4, 0)

	// Set alpha blending
	GX_SetAlphaCompare(GX_ALWAYS, 0, GX_AOP_AND, GX_ALWAYS, 0)
	GX_SetZCompLoc(GX_TRUE)

	/* END SHADER CODE */

	GX_Flush()

	io.Copy(out, wgPipe.Buffer())
}

func checkErr(err error, msg string, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[FAIL] "+msg+":\n    ", args...)
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
