#pragma once

/*! Gamecube controller */
class GCController {
private:
	/*! Gamecube controller port (0-3) */
	unsigned short controllerPort;

public:
	GCController(const unsigned short port) : controllerPort(port) { printf("Created pad with id %d\n", port); }

	bool IsUp(const unsigned short btn) const;
	bool IsDown(const unsigned short btn) const;
	bool IsHeld(const unsigned short btn) const;
	short AnalogX();
	short AnalogY();
	short CStickX();
	short CStickY();
	short TriggerL();
	short TriggerR();
};