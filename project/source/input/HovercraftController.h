#pragma once
#include "GCController.h"

#include <memory>

class HovercraftController {
protected:
	HovercraftController(){};

public:
	/*! All actions that can be performed while controlling an hovercraft */
	enum class Action {
		Jump /*< Jump action */
	};

	/*! All analog motions that can be performed */
	enum class Motion {
		Throttle, /*< Acceleration */
		Pitch,    /*< Pitch */
		Brake,    /*< Deceleration */
		Turn      /*< Turning */
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
	virtual float GetAxis(const Motion axis) const = 0;
};

class GCHovercraftController : public HovercraftController {
private:
	std::shared_ptr<GCController> controller;

public:
	explicit GCHovercraftController(const std::shared_ptr<GCController>& input) : controller(input){};
	const bool GetAction(const Action action) const override;
	float GetAxis(const Motion axis) const override;
};