#pragma once
#include "../pchheader.h"

#include "../math/Matrix.h"
#include "../math/Quaternion.h"
#include "../math/Vector.h"

namespace Components {

/*! \brief Transform component
 *  Gives the entity a position/rotation/scale in the 3D world
 */
struct Transform {
	/*! \brief Create a new transform
	 *  \param position Initial position
	 *  \param rotation Initial rotation
	 */
	explicit Transform(const Vector& position = Vector(0, 0, 0), const Quaternion& rotation = Quaternion(0, 0, 0, 1))
		: position(position), rotation(rotation) {}

	/*! \brief Set rotation using euler angles
	 *  \param rotation Rotation to set
	 */
	void SetRotation(Vector rotation);

	/*! \brief Set rotation using a quaternion
	 *  \param rotation Rotation to set
	 */
	void SetRotation(Quaternion rotation);

	/*! \brief Rotate transform to look at target (forward)
	 *  \param target Position to look towards
	 */
	void Lookat(Vector target);

	/*! \brief Rotate of a certain angle along an axis
	 *  \param axis  Axis to rotate along (should be normalized)
	 *  \param angle Angle of rotation
	 */
	void RotateAxisAngle(Vector axis, float angle);

	/*! \brief Get transform matrix */
	const Matrix& GetMatrix();

	/*! Entity position */
	Vector position;

	/*! Entity scale */
	Vector scale = Vector(1, 1, 1);

	/*! Forward direction */
	Vector forward;

	/*! Right direction */
	Vector right;

	/*! Up direction */
	Vector up;

private:
	Quaternion rotation;
	Matrix matrix;

	/*! \brief Regenerate transform matrix */
	void Flush();
};
} // namespace Components