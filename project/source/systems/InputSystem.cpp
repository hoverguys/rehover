#include "InputSystem.h"

std::shared_ptr<GCController> InputSystem::GetController(unsigned short padId) const {
	if (padId >= PAD_CHANMAX) {
		return nullptr;
	}
	return gcControllers[padId];
}

void InputSystem::configure(ex::EntityManager& entities, ex::EventManager& events) {
	// Initialize pad subsystem(s)
	PAD_Init();
#ifdef WII
	WPAD_Init();
	WPAD_SetDataFormat(WPAD_CHAN_ALL, WPAD_FMT_BTNS_ACC);
#endif

	// Initialize GC controllers
	for (unsigned short id = PAD_CHAN0; id < PAD_CHANMAX; ++id) {
		gcControllers[id] = std::make_shared<GCController>(id);
	}

	gcConnectedPads = PAD_ScanPads();
}

void InputSystem::update(ex::EntityManager& es, ex::EventManager& events, ex::TimeDelta dt) {
	auto newConnectedPads = PAD_ScanPads();
	// Check if the pad mapping changed (added/removed pads)
	if (newConnectedPads != gcConnectedPads) {
		// TODO Manage this
	}

#ifdef WII
// TODO Wiimote support
#endif
};