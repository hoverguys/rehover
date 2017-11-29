#include "Transform.h"
#include "../math/Math.h"
#include <math.h>

namespace Components {
const Matrix& Transform::GetMatrix() {
	Flush();
	return matrix;
}

void Transform::SetRotation(Vector rotation) { this->rotation = Quaternion::FromEuler(rotation); }

void Transform::SetRotation(Quaternion rotation) { this->rotation = rotation; }

void Transform::Lookat(Vector target) {
	Matrix temp = Matrix::LookAt(position, Math::worldUp, target);
	rotation = temp.ToQuaternion();
	
	Flush();
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

	// Update direction vectors with rotation
	up = matrix.MultiplySR(Math::worldUp);
	forward = matrix.MultiplySR(Math::worldForward);
	right = matrix.MultiplySR(Math::worldRight);

	// Scale and translate
	matrix.Scale(scale);
	matrix.Translate(position);
}
} // namespace Components
