/* SDK Libraries */
#include <gccore.h>

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"

#include <hovercraft_bmb.h>

bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();

	MeshResource* resource = new MeshResource((unsigned char*)hovercraft_bmb_txt, hovercraft_bmb_txt_size);
	Mesh* mesh = resource->Load();

	isRunning = TRUE;
	while (isRunning) {
		// Render here
		Graphics::Done();
	}

	return 0;
}

void OnResetCalled() {
	isRunning = FALSE;
}