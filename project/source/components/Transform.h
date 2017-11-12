#pragma once
#include <ogc/gu.h>

namespace Components {
struct Transform {
	Transform(guVector position) : Transform(position, {0, 0, 0, 1}) {}
	Transform(guVector position, guQuaternion rotation) : position(position), rotation(rotation), scale({1, 1, 1}) {}

	void SetRotation(guVector rotation);
	void SetRotation(guQuaternion rotation);
	void Lookat(guVector target);

	Mtx& GetMatrix();

	guVector scale;
	guVector position;

private:
	guQuaternion rotation;
	Mtx matrix;
};
}; // namespace Components