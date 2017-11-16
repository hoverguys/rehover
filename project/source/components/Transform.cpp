#include "Transform.h"
#include "../utils/math.h"
#include <math.h>

namespace Components {
Mtx& Transform::GetMatrix() {
	Flush();
	return matrix;
}

void Transform::SetRotation(guVector rotation) { this->rotation = Math::EulerToQuaternion(rotation); }

void Transform::SetRotation(guQuaternion rotation) { this->rotation = rotation; }

void Transform::Lookat(guVector target) {
	Mtx temp;
	guLookAt(temp, &position, &Math::worldUp, &target);
	c_guQuatMtx(&rotation, temp);
}

void Transform::RotateAxisAngle(guVector axis, float angle) {
	Mtx angleaxis;
	guQuaternion deltaq;
	guMtxRotAxisRad(angleaxis, &axis, angle);
	c_guQuatMtx(&deltaq, angleaxis);
	guQuatMultiply(&deltaq, &rotation, &rotation);
	
	Flush();
}

void Transform::Flush() {
	// Flush local matrix
	guMtxIdentity(matrix);
	guQuatNormalize(&rotation, &rotation);
	c_guMtxQuat(matrix, &rotation);
	guMtxScaleApply(matrix, matrix, scale.x, scale.y, scale.z);
	guMtxTransApply(matrix, matrix, position.x, position.y, position.z);

	// Update direction vectors
	guVecMultiplySR(matrix, &Math::worldUp, &up);
	guVecMultiplySR(matrix, &Math::worldForward, &forward);
	guVecMultiplySR(matrix, &Math::worldRight, &right);

	guVecNormalize(&up);
	guVecNormalize(&forward);
	guVecNormalize(&right);
}
} // namespace Components
