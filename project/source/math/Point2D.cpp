#include "Point2D.h"

Point2D Point2D::operator*(const float& scale) const {
	return Point2D(x * scale, y * scale);
}

Point2D Point2D::operator*(const Point2D& other) const {
	return Point2D(x * other.x, y * other.y);
}

Point2D Point2D::operator+(const Point2D& other) const {
	return Point2D(x + other.x, y + other.y);
}

Point2D Point2D::operator-(const Point2D& other) const {
	return Point2D(x - other.x, y - other.y);
}

Point2D& Point2D::operator*=(const float& scale) {
	x *= scale;
	y *= scale;
	return *this;
}

Point2D& Point2D::operator*=(const Point2D& other) {
	x *= other.x;
	y *= other.y;
	return *this;
}

Point2D& Point2D::operator+=(const Point2D& other) {
	x += other.x;
	y += other.y;
	return *this;
}

Point2D& Point2D::operator-=(const Point2D& other) {
	x -= other.x;
	y -= other.y;
	return *this;
}