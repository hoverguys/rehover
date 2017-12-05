#pragma once
#include "../pchheader.h"

#include "../input/HovercraftController.h"

namespace ex = entityx;

namespace Behaviours {

/*! \brief Hovercraft behavior
 *  Allows the hovercraft entity to be controlled by players
 */
struct Hovercraft {
	/*! Controller used to control the hovercraft */
	std::shared_ptr<HovercraftController> controller;

	/*! Camera following the hovercraft */
	ex::Entity camera;

	/*! Update method */
	void Tick(ex::Entity entity, ex::TimeDelta dt);
};
} // namespace Behaviours
