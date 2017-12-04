#pragma once

/*! Vector with 2 elements for 2D positions and math */
struct Vector2D {
	float x; /*< X value */
	float y; /*< Y value */

	Vector2D(const float x, const float y) : x(x), y(y) {}
	Vector2D() : Vector2D(0, 0) {}

	Vector2D operator*(const float& scale) const;
	Vector2D operator*(const Vector2D& scale) const;
	Vector2D operator+(const Vector2D& other) const;
	Vector2D operator-(const Vector2D& other) const;

	Vector2D& operator*=(const float& scale);
	Vector2D& operator*=(const Vector2D& scale);
	Vector2D& operator+=(const Vector2D& other);
	Vector2D& operator-=(const Vector2D& other);
};