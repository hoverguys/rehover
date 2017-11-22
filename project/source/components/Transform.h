#pragma once
#include "../math/Quaternion.h"
#include "../math/Matrix.h"
#include "../math/Vector.h"

namespace Components {
struct Transform {
	Transform(Vector position) : Transform(position, {0, 0, 0, 1}) {}
	Transform(Vector position, Quaternion rotation) : position(position), scale({1, 1, 1}), rotation(rotation) {}

	void SetRotation(Vector rotation);
	void SetRotation(Quaternion rotation);
	void Lookat(Vector target);
	void RotateAxisAngle(Vector axis, float angle);

	Matrix GetMatrix();

	Vector position;
	Vector scale;

	Vector forward;
	Vector right;
	Vector up;

private:
	Quaternion rotation;
	Matrix matrix;

	void Flush();
};
} // namespace Components