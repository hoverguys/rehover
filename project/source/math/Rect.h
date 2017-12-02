#pragma once

#include "Point2D.h"

/*! \brief 2D Rectangle */
class Rect {
public:
	/*! Top-Left corner of the rectangle */
	Point2D start;

	/*! Bottom-Right corner of the rectangle */
	Point2D end;

	/*! \brief Create a rectangle providing the top-left and bottom-right corner coordinates
	 *
	 *  \param start Top left corner
	 *  \param end   Bottom right corner
	 */
	Rect(const Point2D& start, const Point2D& end) : start(start), end(end) {}

	/*! \brief Create a rectangle providing top left coordinates and size
	 *
	 *  \param x      Left coordinate
	 *  \param y      Top coordinate
	 *  \param width  Width
	 *  \param height Height
	 */
	Rect(const float x, const float y, const float width, const float height)
	    : Rect(Point2D(x, y), Point2D(x + width, y + height)) {}

	Rect() : Rect(Point2D(), Point2D()) {}

	/*! \brief Get width of the rectangle
	 */
	float Width() const { return end.x - start.x; }

	/*! \brief Get height of the rectangle
	 */
	float Height() const { return end.y - start.y; }

	/*! \brief Move the rectangle around using a relative point
	 */
	void Move(const Point2D& delta);
	void Move(const float x, const float y) { Move(Point2D(x, y)); }

	/*! \brief Resize the rectangle of a certain width and height
	 */
	void Resize(const Point2D& size);
	void Resize(const float width, const float height) { Resize(Point2D(width, height)); }
};