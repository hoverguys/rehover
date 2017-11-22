#include "Transform.h"
#include "../utils/math.h"
#include <math.h>

namespace Components {
Matrix Transform::GetMatrix() {
	Flush();
	return matrix;
}

void Transform::SetRotation(Vector rotation) { this->rotation = Math::EulerToQuaternion(rotation); }

void Transform::SetRotation(Quaternion rotation) { this->rotation = rotation; }

void Transform::Lookat(Vector target) {
	Matrix temp = Matrix::LookAt(position, Math::worldUp, target);
	rotation = temp.ToQuaternion();
}

void Transform::RotateAxisAngle(Vector axis, float angle) {
	Matrix angleaxis = Matrix::AxisAngle(axis, angle);
	Quaternion deltaq = angleaxis.ToQuaternion();
	rotation = deltaq * rotation;
	
	Flush();
}

void Transform::Flush() {
	// Flush local matrix
	rotation.Normalize();
	matrix = rotation.ToMatrix();
	matrix.Scale(scale);
	matrix.Translate(position);

	// Update direction vectors
	up = matrix.MultiplySR(Math::worldUp).Normalized();
	forward = matrix.MultiplySR(Math::worldForward).Normalized();
	right = matrix.MultiplySR(Math::worldRight).Normalized();
}
} // namespace Components
