set(ASSETS_PREFIX assets/)

make_atlas(gamehud GAMEHUD_TEXTURE GAMEHUD_ATLAS ${ASSETS_PREFIX} 256
	sprites/LBPosition1.png
	sprites/LBPosition2.png
	sprites/LBPosition3.png
	sprites/LBPosition4.png
)

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
		RGB5A3 CLAMP BILINEAR ABS ${GAMEHUD_TEXTURE}
	SHADER
		shaders/hovercraft.tev
	BIN
		ABS ${GAMEHUD_ATLAS}
)