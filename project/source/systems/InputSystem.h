#pragma once
#include "../pchheader.h"

#include "../input/GCController.h"

namespace ex = entityx;

class InputSystem : public ex::System<InputSystem> {
private:
	/*! All connected pads, as bitset */
	unsigned long gcConnectedPads;

	/*! All gamecube pads instances */
	std::shared_ptr<GCController> gcControllers[PAD_CHANMAX];

public:
	InputSystem() : gcConnectedPads(0) {}

	/*! \brief Provide a gamecube controller
	 *  \param id Pad slot #
	 *  \return Shared pointer of an GCController instance
	 */
	std::shared_ptr<GCController> GetController(unsigned short padId) const;

	// EntityX methods
	void configure(ex::EntityManager& entities, ex::EventManager& events) override;
	void update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) override;
};