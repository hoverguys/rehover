#include "Matrix.h"
#include "Vector.h"
#include "Quaternion.h"
#include <math.h>

Matrix Matrix::Identity() {
    Matrix identity;
    identity.internal = {
        1.0f, 0.0f, 0.0f, 0.0f,
        0.0f, 1.0f, 0.0f, 0.0f,
        0.0f, 0.0f, 1.0f, 0.0f,
    };

    return identity;
}

Matrix Matrix::LookAt(Vector cameraOrigin, Vector cameraUp, Vector cameraTarget) {
    Vector look = (cameraOrigin - cameraTarget).Normalized();
    Vector right = cameraUp.Cross(look).Normalized();
    Vector up = look.Cross(right);

    Matrix lookat;
    lookat.internal = {
        right.x,
        right.y,
        right.z,
        -( cameraOrigin.x * right.x + cameraOrigin.y * right.y + cameraOrigin.z * right.z ),

        up.x,
        up.y,
        up.z,
        -( cameraOrigin.x * up.x + cameraOrigin.y * up.y + cameraOrigin.z * up.z ),

        look.x,
        look.y,
        look.z,
        -( cameraOrigin.x * look.x + cameraOrigin.y * look.y + cameraOrigin.z * look.z )
    };

    return lookat;
}

Matrix Matrix::AxisAngle(Vector axis, float angle) {
    Matrix angleaxis;
	
	float s = sinf(angle);
	float c = cosf(angle);
	float t = 1.0f - c;
	
    axis.Normalize();
	float xSq = axis.x * axis.x;
	float ySq = axis.y * axis.y;
	float zSq = axis.z * axis.z;

    angleaxis.internal = {
        ( t * xSq )   + ( c ),
        ( t * axis.x * axis.y ) - ( s * axis.z ),
        ( t * axis.x * axis.z ) + ( s * axis.y ),
        0.0f,

        ( t * axis.x * axis.y ) + ( s * axis.z ),
        ( t * ySq ) + ( c ),
        ( t * axis.y * axis.z ) - ( s * axis.x ),
        0.0f,

        ( t * axis.x * axis.z ) - ( s * axis.y ),
        ( t * axis.y * axis.z ) + ( s * axis.x ),
        ( t * zSq )   + ( c ),
        0.0f
    };

    return angleaxis;
}

void Matrix::Scale(Vector scale) {
    for (int i=0; i < 4; ++i) {
        internal[i] *= scale.x;
    }

    for (int i=0; i < 4; ++i) {
        internal[4+i] *= scale.y;
    }

    for (int i=0; i < 4; ++i) {
        internal[8+i] *= scale.z;
    }
}

void Matrix::Translate(Vector delta) {
	internal[3] += delta.x;
	internal[7] += delta.y;
	internal[11] += delta.z;
}

void Matrix::Inverse() {
    float det = (internal[0] * internal[5] * internal[10]) +
                (internal[1] * internal[6] * internal[8]) +
                (internal[2] * internal[4] * internal[9]) -
                (internal[8] * internal[5] * internal[2]) -
                (internal[4] * internal[1] * internal[10]) -
                (internal[0] * internal[9] * internal[6]);
    
    // Early out
    if (det == 0.0f) {
        return;
    }

    float rcpdet = 1.0f / det;
    Matrix temp;

    temp.internal[0] =  (internal[5] * internal[10] - internal[9] * internal[6]) * rcpdet;
    temp.internal[1] = -(internal[1] * internal[10] - internal[9] * internal[2]) * rcpdet;
    temp.internal[2] =  (internal[1] * internal[6] - internal[5] * internal[2]) * rcpdet;

    temp.internal[4] = -(internal[4] * internal[10] - internal[8] * internal[6]) * rcpdet;
    temp.internal[5] =  (internal[0] * internal[10] - internal[8] * internal[2]) * rcpdet;
    temp.internal[6] = -(internal[0] * internal[6] - internal[4] * internal[2]) * rcpdet;

    temp.internal[8] =  (internal[4] * internal[9] - internal[8] * internal[5]) * rcpdet;
    temp.internal[9] = -(internal[0] * internal[9] - internal[8] * internal[1]) * rcpdet;
    temp.internal[10] = (internal[0] * internal[5] - internal[4] * internal[1]) * rcpdet;

    temp.internal[3]  = -temp.internal[0] * internal[3] - temp.internal[1] * internal[7] - temp.internal[2] * internal[11];
    temp.internal[7]  = -temp.internal[4] * internal[3] - temp.internal[5] * internal[7] - temp.internal[6] * internal[11];
    temp.internal[11] = -temp.internal[8] * internal[3] - temp.internal[9] * internal[7] - temp.internal[10] * internal[11];

    //TODO std::move?
    internal = temp.internal;
}

Matrix Matrix::Inversed() const {
    Matrix inversed = *this;
    inversed.Inverse();
    return inversed;
}

void Matrix::Transpose() {
    Matrix temp;
    temp.internal = {
        internal[0],
        internal[4],
        internal[8],
        0.0f,
        internal[1],
        internal[5],
        internal[9],
        0.0f,
        internal[2],
        internal[6],
        internal[10],
        0.0f
    };

    //TODO std::move?
    internal = temp.internal;
}

Vector Matrix::Multiply(Vector vec) const {
	Vector result;
    result.x = internal[0] * vec.x + internal[1] * vec.y + internal[2] * vec.z + internal[3];
    result.y = internal[4] * vec.x + internal[5] * vec.y + internal[6] * vec.z + internal[7];
    result.z = internal[8] * vec.x + internal[9] * vec.y + internal[10] * vec.z + internal[11];

    return result;
}

Vector Matrix::MultiplySR(Vector vec) const {
    Vector result;
    result.x = internal[0] * vec.x + internal[1] * vec.y + internal[2] * vec.z;
    result.y = internal[4] * vec.x + internal[5] * vec.y + internal[6] * vec.z;
    result.z = internal[8] * vec.x + internal[9] * vec.y + internal[10] * vec.z;

    return result;
}

Matrix Matrix::operator* (const Matrix& o) {
    Matrix result;
    result.internal = {
        internal[0] * o.internal[0] + internal[1] * o.internal[4] + internal[2] * o.internal[8],
        internal[0] * o.internal[1] + internal[1] * o.internal[5] + internal[2] * o.internal[9],
        internal[0] * o.internal[2] + internal[1] * o.internal[6] + internal[2] * o.internal[10],
        internal[0] * o.internal[3] + internal[1] * o.internal[7] + internal[2] * o.internal[11] + internal[3],
 
        internal[4] * o.internal[0] + internal[5] * o.internal[4] + internal[6] * o.internal[8],
        internal[4] * o.internal[1] + internal[5] * o.internal[5] + internal[6] * o.internal[9],
        internal[4] * o.internal[2] + internal[5] * o.internal[6] + internal[6] * o.internal[10],
        internal[4] * o.internal[3] + internal[5] * o.internal[7] + internal[6] * o.internal[11] + internal[7],
 
        internal[8] * o.internal[0] + internal[9] * o.internal[4] + internal[10] * o.internal[8],
        internal[8] * o.internal[1] + internal[9] * o.internal[5] + internal[10] * o.internal[9],
        internal[8] * o.internal[2] + internal[9] * o.internal[6] + internal[10] * o.internal[10],
        internal[8] * o.internal[3] + internal[9] * o.internal[7] + internal[10] * o.internal[11] + internal[11]
    }; 
 
    return result;
}

Quaternion Matrix::ToQuaternion() const {
    Quaternion result;

    const float diag = internal[0] + internal[5] + internal[10] + 1;
    const float scale = sqrtf(diag) * 2.0f;

	if(diag > 0.0f) {
		result.x = (internal[9] - internal[6]) / scale;
		result.y = (internal[2] - internal[8]) / scale;
		result.z = (internal[4] - internal[1]) / scale;
		result.w = 0.25f * scale;
	} else {
		if(internal[0] > internal[5] && internal[0] > internal[10]) {
			result.x = 0.25f * scale;
			result.y = (internal[1] + internal[4]) / scale;
			result.z = (internal[8] + internal[2]) / scale;
			result.w = (internal[9] - internal[6]) / scale;
		} else if(internal[5] > internal[10]) {
			result.x = (internal[1] + internal[4]) / scale;
			result.y = 0.25f * scale;
			result.z = (internal[6] + internal[9]) / scale;
			result.w = (internal[2] - internal[8]) / scale;
		} else {
			result.x = (internal[2] + internal[8]) / scale;
			result.y = (internal[6] + internal[9]) / scale;
			result.z = 0.25f * scale;
			result.w = (internal[4] - internal[1]) / scale;
		}
	}

    return result.Normalized();
}

void Matrix::ToNative(Mtx matrix) const {
    for (int i=0; i < 4; ++i) {
        matrix[0][i] = internal[i];
    }

    for (int i=0; i < 4; ++i) {
        matrix[1][i] = internal[4+i];
    }

    for (int i=0; i < 4; ++i) {
        matrix[2][i] = internal[8+i];
    }
}