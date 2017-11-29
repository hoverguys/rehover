set(ASSETS_PREFIX assets/)
set(ASSETS
	MODEL
		models/hovercraft.obj
		models/plane.obj
		models/terrain.obj
		models/testmap.obj
	TEXTURE
		IA8 CLAMP  BILINEAR textures/hovercraftGlobal.png
		IA8 CLAMP  BILINEAR textures/hovercraftShade.png
		I4  REPEAT NEAR     textures/checkerboard.png
		RGB565 REPEAT BILINEAR textures/testmap.png
	BIN
		shaders/hovercraft.bin
)