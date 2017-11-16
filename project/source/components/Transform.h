#pragma once
#include <ogc/gu.h>

namespace Components {
struct Transform {
	Transform(guVector position) : Transform(position, {0, 0, 0, 1}) {}
	Transform(guVector position, guQuaternion rotation) : position(position), scale({1, 1, 1}), rotation(rotation) {}

	void SetRotation(guVector rotation);
	void SetRotation(guQuaternion rotation);
	void Lookat(guVector target);
	void RotateAxisAngle(guVector axis, float angle);

	Mtx& GetMatrix();

	guVector position;
	guVector scale;

	guVector forward;
	guVector right;
	guVector up;

private:
	guQuaternion rotation;
	Mtx matrix;

	void Flush();
};
} // namespace Components