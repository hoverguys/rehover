#include "HovercraftController.h"

#include <ogc/pad.h>

/*
 * Gamecube controller
 *
 * Default mapping (not customizable atm):
 *    A (Button) -> Jump
 *    L (Analog) -> Brake
 *    R (Analog) -> Accelerate
 *    CtrlSt. X  -> Turn
 *
 */

const bool GCHovercraftController::GetAction(HovercraftController::Action action) const {
	// TODO Make button mapping configurable in the future
	switch (action) {
	case Action::Jump: return controller->IsDown(PAD_BUTTON_A);
	default: return false;
	}
}

float GCHovercraftController::GetAxis(HovercraftController::Motion axis) const {
	// TODO Make axis mapping configurable in the future
	switch (axis) {
	case Motion::Throttle: return controller->TriggerR();
	case Motion::Brake: return controller->TriggerL();
	case Motion::Turn: return controller->AnalogX();
	case Motion::Pitch: return controller->AnalogY();
	default: return 0;
	}
}