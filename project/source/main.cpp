/* SDK Libraries */
#include <gccore.h>

#include "rendering/Graphics.h"
#include "resources/MeshResource.h"

#include <hovercraft_obj.h>

bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();

	MeshResource* resource = new MeshResource((unsigned char*)hovercraft_obj_txt, hovercraft_obj_txt_size);
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