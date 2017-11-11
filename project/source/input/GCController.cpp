#include "GCController.h"

#include <ogc/pad.h>

bool GCController::IsDown(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsDown(controllerPort);
	return buttons & btnid != 0;
}

bool GCController::IsUp(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsUp(controllerPort);
	return buttons & btnid != 0;
}

bool GCController::IsHeld(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsHeld(controllerPort);
	return buttons & btnid != 0;
}

short GCController::AnalogX() { return PAD_StickX(controllerPort); }
short GCController::AnalogY() { return PAD_StickY(controllerPort); }
short GCController::CStickX() { return PAD_SubStickX(controllerPort); }
short GCController::CStickY() { return PAD_SubStickY(controllerPort); }
short GCController::TriggerL() { return PAD_TriggerL(controllerPort); }
short GCController::TriggerR() { return PAD_TriggerR(controllerPort); }
