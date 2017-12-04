#pragma once

/*! Gamecube controller */
class GCController {
private:
	/*! Gamecube controller port (0-3) */
	unsigned short controllerPort;

	/*! Controller's deadzone
	 *  Ignore all analog values under this absolute value to avoid jitter
	 *  \todo This should be made customizable in the future as not all controllers are the same
	 */
	const float Deadzone = 10;

	/*! Max range for analog sticks
	 *  Clamp all analog values to this maximum absolute value as most analog sticks
	 *  don't reach the top value
	 *  \todo This should be made customizable in the future as not all controllers are the same
	 */
	const float MaxStickThreshold = 70;

	/*! Max range for triggers
	 *  Clamp all trigger values to this maximum absolute value as most triggers don't
	 *  reach the top value (at least, not easily)
	 *  \todo This should be made customizable in the future as not all controllers are the same
	 */
	const float MaxTriggerThreshold = 200;

	const float AnalogMultiplier = 1.0f / MaxStickThreshold;
	const float TriggerMultiplier = 1.0f / MaxTriggerThreshold;

public:
	/*! \brief Set-up GC controller reader
	 *  \param port What port to read from (0-3)
	 */
	explicit GCController(const unsigned short port) : controllerPort(port) {}

	/*! \brief Check if button is up (not pressed)
	 *  \param  btn Button to check for
	 *  \return true if the button is not being pressed, false otherwise
	 */
	const bool IsUp(const unsigned short btn) const;

	/*! \brief Check if button is down (just pressed)
	 *  \param  btn Button to check for
	 *  \return true if the button is being pressed, false otherwise
	 */
	const bool IsDown(const unsigned short btn) const;

	/*! \brief Check if button is held (being pressed for more than one frame)
	 *  \param  btn Button to check for
	 *  \return true if the button is not being held down, false otherwise
	 */
	const bool IsHeld(const unsigned short btn) const;

	/*! \brief Get analog stick X axis value */
	float AnalogX() const;

	/*! \brief Get analog stick Y axis value */
	float AnalogY() const;

	/*! \brief Get C-stick X axis value */
	float CStickX() const;

	/*! \brief Get C-stick Y axis value */
	float CStickY() const;

	/*! \brief Get left trigger analog value */
	float TriggerL() const;

	/*! \brief Get right trigger analog value */
	float TriggerR() const;
};