set(ASSETS_PREFIX assets/)
set(ASSETS
	MODEL
		models/hovercraft.obj
		models/plane.obj
	TEXTURE
		IA8 CLAMP  BILINEAR textures/hovercraftGlobal.png
		IA8 CLAMP  BILINEAR textures/hovercraftShade.png
		CMPR CLAMP NEAR     textures/checkerboard.png
		RGBA8  CLAMP NEAR     textures/rainbow.png
	BIN
		shaders/hovercraft.bin
)