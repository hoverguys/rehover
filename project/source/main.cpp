/* SDK Libraries */
#include <gccore.h>

#include "graphics.h"

bool isRunning;
void OnResetCalled();

int main() {

	/* Setup reset function */
	SYS_SetResetCallback(OnResetCalled);

	Graphics::Init();

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