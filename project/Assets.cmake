set(ASSETS_PREFIX assets/)
set(ASSETS
	MODEL
		models/hovercraft.obj
		models/plane.obj
		models/terrain.obj
	TEXTURE
		IA8 CLAMP  BILINEAR textures/hovercraftGlobal.png
		IA8 CLAMP  BILINEAR textures/hovercraftShade.png
		I4  REPEAT NEAR     textures/checkerboard.png
		RGB565 CLAMP BILINEAR textures/terrain.png
	BIN
		shaders/hovercraft.bin
)