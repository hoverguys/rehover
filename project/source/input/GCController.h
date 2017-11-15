#pragma once

/*! Gamecube controller */
class GCController {
private:
	/*! Gamecube controller port (0-3) */
	unsigned short controllerPort;

	const float Deadzone = 10;             // Ignore all analog values under this absolute value
	const float MaxStickThreshold = 70;    // Clamp all analog values to this maximum absolute value
	const float MaxTriggerThreshold = 200; // Clamp all trigger values to this maximum absolute value

	const float AnalogMultiplier = 1.0f / MaxStickThreshold;
	const float TriggerMultiplier = 1.0f / MaxTriggerThreshold;

public:
	GCController(const unsigned short port) : controllerPort(port) {}

	const bool IsUp(const unsigned short btn) const;
	const bool IsDown(const unsigned short btn) const;
	const bool IsHeld(const unsigned short btn) const;
	float AnalogX() const;
	float AnalogY() const;
	float CStickX() const;
	float CStickY() const;
	float TriggerL() const;
	float TriggerR() const;
};