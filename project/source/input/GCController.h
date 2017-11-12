#pragma once

/*! Gamecube controller */
class GCController {
private:
	/*! Gamecube controller port (0-3) */
	unsigned short controllerPort;

public:
	GCController(const unsigned short port) : controllerPort(port) {}

	const bool IsUp(const unsigned short btn) const;
	const bool IsDown(const unsigned short btn) const;
	const bool IsHeld(const unsigned short btn) const;
	short AnalogX() const;
	short AnalogY() const;
	short CStickX() const;
	short CStickY() const;
	short TriggerL() const;
	short TriggerR() const;
};