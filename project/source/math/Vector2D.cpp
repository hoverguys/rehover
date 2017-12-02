#include "Vector2D.h"

Vector2D Vector2D::operator*(const float& scale) const {
	return Vector2D(x * scale, y * scale);
}

Vector2D Vector2D::operator*(const Vector2D& other) const {
	return Vector2D(x * other.x, y * other.y);
}

Vector2D Vector2D::operator+(const Vector2D& other) const {
	return Vector2D(x + other.x, y + other.y);
}

Vector2D Vector2D::operator-(const Vector2D& other) const {
	return Vector2D(x - other.x, y - other.y);
}

Vector2D& Vector2D::operator*=(const float& scale) {
	x *= scale;
	y *= scale;
	return *this;
}

Vector2D& Vector2D::operator*=(const Vector2D& other) {
	x *= other.x;
	y *= other.y;
	return *this;
}

Vector2D& Vector2D::operator+=(const Vector2D& other) {
	x += other.x;
	y += other.y;
	return *this;
}

Vector2D& Vector2D::operator-=(const Vector2D& other) {
	x -= other.x;
	y -= other.y;
	return *this;
}