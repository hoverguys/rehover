#pragma once
#include "GCController.h"

#include <memory>

class HovercraftController {
protected:
	HovercraftController(){};

public:
	/*! All actions that can be performed while controlling an hovercraft */
	enum Action {
		ActionJump /*< Jump action */
	};

	/*! All analog motions that can be performed */
	enum Motion {
		MotionThrottle, /*< Acceleration */
		MotionPitch,	/*< Pitch */
		MotionBrake,    /*< Deceleration */
		MotionTurn      /*< Turning */
	};

	/*! \brief Check if the player is trying to perform an action
	 *  \param action What action to check for
	 *  \return true if the player did press a button to perform such action, false otherwise
	 */
	virtual const bool GetAction(const Action action) const = 0;

	/*! \brief Check if the player is trying to trying to perform some time of motion
	 *  \param axis What axis of motion to check for
	 *  \return true if the player is trying to move along that axis, false otherwise
	 */
	virtual short GetAxis(const Motion axis) const = 0;
};

class GCHovercraftController : public HovercraftController {
private:
	std::shared_ptr<GCController> controller;

public:
	GCHovercraftController(std::shared_ptr<GCController> input) : controller(input){};
	const bool GetAction(const Action action) const override;
	short GetAxis(const Motion axis) const override;
};