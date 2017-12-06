#pragma once
#include "Math.h"

/*! \brief Quaternion class
 */
class Quaternion {
public:
	float x; /*< X value */
	float y; /*< Y value */
	float z; /*< Z value */
	float w; /*< W value */

	/*! \brief Create a quaternion from given values
	 *  \param x X value
	 *  \param y Y value
	 *  \param z Z value
	 *  \param w W value
	 */
	explicit Quaternion(float x, float y, float z, float w) : x(x), y(y), z(z), w(w) {}

	/*! \brief Create a quaternion with no angle or magnitude
	 */
	explicit Quaternion() : Quaternion(0, 0, 0, 1) {}

	/*! \brief Normalize the quaternion in place */
	void Normalize();

	/*! \brief Get a normalized copy of the quaternion
	 *  \return Normalized quaternion
	 */
	Quaternion Normalized() const;

	/*! \brief Get the squared magnitude (length) of the quaternion
	 *  \return Squared magnitude/length of the quaternion
	 */
	float SqrMagnitude() const;

	/*! \brief Get the rotation matrix for the quaternion
	 *  \return Rotation matrix
	 */
	Matrix ToMatrix() const;

	Quaternion operator*(const Quaternion& other) const;

	/*! \brief Create a quaternion from euler angles
	 *  \param rotation Euler angles
	 *  \return Quaternion that matches the euler angles
	 */
	static Quaternion FromEuler(const Vector& rotation);
};