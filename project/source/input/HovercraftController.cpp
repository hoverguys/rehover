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
	case ActionJump: return controller->IsDown(PAD_BUTTON_A);
	}
}

short GCHovercraftController::GetAxis(HovercraftController::Motion axis) const {
	// TODO Make axis mapping configurable in the future
	switch (axis) {
	case MotionThrottle: return controller->TriggerR();
	case MotionBrake: return controller->TriggerL();
	case MotionTurn: return controller->AnalogX();
	case MotionPitch: return controller->AnalogY();
	}
}