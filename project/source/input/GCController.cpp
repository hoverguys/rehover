#include "GCController.h"

#include <ogc/pad.h>

const bool GCController::IsDown(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsDown(controllerPort);
	return buttons & btnid != 0;
}

const bool GCController::IsUp(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsUp(controllerPort);
	return buttons & btnid != 0;
}

const bool GCController::IsHeld(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsHeld(controllerPort);
	return buttons & btnid != 0;
}

short GCController::AnalogX() const { return PAD_StickX(controllerPort); }
short GCController::AnalogY() const { return PAD_StickY(controllerPort); }
short GCController::CStickX() const { return PAD_SubStickX(controllerPort); }
short GCController::CStickY() const { return PAD_SubStickY(controllerPort); }
short GCController::TriggerL() const { return PAD_TriggerL(controllerPort); }
short GCController::TriggerR() const { return PAD_TriggerR(controllerPort); }
