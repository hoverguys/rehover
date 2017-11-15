#include "GCController.h"

#include <cmath>
#include <ogc/pad.h>

inline float clamp(const float value, const float minVal, const float maxVal) {
	return value < minVal ? minVal : value > maxVal ? maxVal : value;
}

inline float normalize(const float raw, const float deadzone, const float threshold, const float multiplier) {
	if (std::fabs(raw) < deadzone) {
		return 0;
	}
	return clamp(raw, -threshold, threshold) * multiplier;
}

const bool GCController::IsDown(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsDown(controllerPort);
	return buttons & (btnid != 0);
}

const bool GCController::IsUp(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsUp(controllerPort);
	return buttons & (btnid != 0);
}

const bool GCController::IsHeld(const unsigned short btnid) const {
	auto buttons = PAD_ButtonsHeld(controllerPort);
	return buttons & (btnid != 0);
}

float GCController::AnalogX() const {
	return normalize(PAD_StickX(controllerPort), Deadzone, MaxStickThreshold, AnalogMultiplier);
}

float GCController::AnalogY() const {
	return normalize(PAD_StickY(controllerPort), Deadzone, MaxStickThreshold, AnalogMultiplier);
}

float GCController::CStickX() const {
	return normalize(PAD_SubStickX(controllerPort), Deadzone, MaxStickThreshold, AnalogMultiplier);
}

float GCController::CStickY() const {
	return normalize(PAD_SubStickY(controllerPort), Deadzone, MaxStickThreshold, AnalogMultiplier);
}

float GCController::TriggerL() const {
	return normalize(PAD_TriggerL(controllerPort), Deadzone, MaxTriggerThreshold, TriggerMultiplier);
}

float GCController::TriggerR() const {
	return normalize(PAD_TriggerR(controllerPort), Deadzone, MaxTriggerThreshold, TriggerMultiplier);
}
