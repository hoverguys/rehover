#pragma once

struct Point2D {
	float x, y;

	Point2D(const float x, const float y) : x(x), y(y) {}
	Point2D() : Point2D(0, 0) {}

	Point2D operator*(const float& scale) const;
	Point2D operator*(const Point2D& scale) const;
	Point2D operator+(const Point2D& other) const;
	Point2D operator-(const Point2D& other) const;

	Point2D& operator*=(const float& scale);
	Point2D& operator*=(const Point2D& scale);
	Point2D& operator+=(const Point2D& other);
	Point2D& operator-=(const Point2D& other);
};