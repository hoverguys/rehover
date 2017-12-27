#pragma once

#include "../pchheader.h"
#include "Vector2D.h"

/*! \brief 2D Rectangle */
class Rect {
public:
	/*! Top-Left corner of the rectangle */
	Vector2D start;

	/*! Size of the rectangle */
	Vector2D size;

	/*! \brief Create a rectangle providing the top-left and bottom-right corner coordinates
	 *
	 *  \param start  Origin (top left corner)
	 *  \param size   Size of the rectangle
	 */
	constexpr Rect(const Vector2D& start, const Vector2D& size) : start(start), size(size) {}

	/*! \brief Create a rectangle providing top left coordinates and size
	 *
	 *  \param x      Left coordinate
	 *  \param y      Top coordinate
	 *  \param width  Width
	 *  \param height Height
	 */
	constexpr Rect(const float x, const float y, const float width, const float height)
		: Rect(Vector2D(x, y), Vector2D(width, height)) {}

	constexpr Rect() : Rect(Vector2D(), Vector2D()) {}

	/*! \brief Get width of the rectangle
	 */
	float Width() const { return size.x; }

	/*! \brief Get height of the rectangle
	 */
	float Height() const { return size.y; }

	/*! \brief Get the bounds of the rectangle
	 *
	 *  \return The top-left and bottom-right corners as a pair
	 */
	std::pair<Vector2D, Vector2D> Bounds();

	/*! \brief Move the rectangle around using a relative point
	 */
	void Move(const Vector2D& delta);
	void Move(const float x, const float y) { Move(Vector2D(x, y)); }

	/*! \brief Resize the rectangle of a certain width and height
	 */
	void Resize(const Vector2D& size);
	void Resize(const float width, const float height) { Resize(Vector2D(width, height)); }
};