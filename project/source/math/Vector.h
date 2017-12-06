#pragma once
#include "Math.h"

/*! \brief Vector with 3 elements for 3D positioning and math */
class Vector {
public:
	float x; /*< X value */
	float y; /*< Y value */
	float z; /*< Z value */

	/*! \brief Create a Vector from 3 values (one for each axis)
	 *  \param x X value
	 *  \param y Y valye
	 *  \param z Z value
	 */
	Vector(float x, float y, float z) : x(x), y(y), z(z) {}

	/*! \brief Create a Vector at the origin (0,0,0)
	 */
	Vector() : Vector(0, 0, 0) {}

	/*! \brief Normalize the vector in place */
	void Normalize();

	/*! \brief Get a normalized copy of the vector
	 *  \return The normalized vector
	 */
	Vector Normalized() const;

	/*! \brief Get the cross product (perpendicular vector) between two vectors (current and another one)
	 *  \param other Vector to calculate cross product with
	 *  \return Result of the cross product
	 */
	Vector Cross(const Vector& other) const;

	/*! \brief Get the dot product (inner product) between two vectors (current and another one)
	 *  \param other Vector to calculate dot product with
	 *  \return Result of the dot product
	 */
	float Dot(const Vector& other) const;

	/*! \brief Get the magnitude (length) of the vector
	 *  \return Magnitude/Length of the vector
	 */
	float Magnitude() const;

	/*! \brief Get the squared magnitude (length) of the vector
	 *  This is a bit faster than Magnitude() since it doesn't perform the square root step required in
	 *  calculating the euclidean norm of the vector
	 *  \return Squared magnitude/length of the vector
	 */
	float SqrMagnitude() const;

	Vector operator*(const float& scale) const;
	Vector operator*(const Vector& scale) const;
	Vector operator+(const Vector& other) const;
	Vector operator-(const Vector& other) const;

	/*! \brief Return the vector as a guVector (libogc format) */
	operator guVector() const;
};