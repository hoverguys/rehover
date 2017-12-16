set(ASSETS_PREFIX assets/)
set(ASSETS
	MODEL
		models/hovercraft.obj
#		models/plane.obj
#		models/terrain.obj
		models/testmap.obj
#		models/testplane.obj
#		models/testdip.obj
	TEXTURE
		IA8 CLAMP  BILINEAR textures/hovercraftGlobal.png
		I8  CLAMP  BILINEAR textures/hovercraftShade.png
#		I4  REPEAT NEAR     textures/checkerboard.png
		RGB565 REPEAT NEAR  textures/testmap.png
		RGB5A3 CLAMP BILINEAR sprites/logo.png
	SHADER
		shaders/hovercraft.tev
)