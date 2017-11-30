#pragma once
#include "../math/Quaternion.h"
#include "../math/Matrix.h"
#include "../math/Vector.h"

namespace Components {
struct Transform {
	explicit Transform(const Vector& position = Vector(0, 0, 0), const Quaternion& rotation = Quaternion(0, 0, 0, 1)) : position(position), rotation(rotation) {}

	void SetRotation(Vector rotation);
	void SetRotation(Quaternion rotation);
	void Lookat(Vector target);
	void RotateAxisAngle(Vector axis, float angle);

	const Matrix& GetMatrix();

	Vector position;
	Vector scale = Vector(1, 1, 1);

	Vector forward;
	Vector right;
	Vector up;

private:
	Quaternion rotation;
	Matrix matrix;

	void Flush();
};
} // namespace Components