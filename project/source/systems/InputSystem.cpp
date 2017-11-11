#include "InputSystem.h"

#include <ogc/pad.h>

void InputSystem::configure(ex::EntityManager& entities, ex::EventManager& events) {
	// Initialize pad subsystem(s)
	PAD_Init();
#ifdef WII
	WPAD_Init();
	WPAD_SetDataFormat(WPAD_CHAN_ALL, WPAD_FMT_BTNS_ACC);
#endif

	gcConnectedPads = PAD_ScanPads();
	printf("Pad matrix: %x\n", gcConnectedPads);
	for (unsigned short padid = 0; padid < PAD_CHANMAX; ++padid) {
		// Check if the pad is connected
		if (gcConnectedPads & (1 << padid) != 0) {
			// If so, initialize the gamepad controller object
			gcControllers.emplace(padid, GCController{padid});
		}
	}
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