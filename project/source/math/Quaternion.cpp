#include "Quaternion.h"
#include "Matrix.h"
#include "Vector.h"
#include <math.h>

void Quaternion::Normalize() {
    const float sqrLength = SqrMagnitude();
    if ( sqrLength == 0.0f) {
        // ERROR
        x = y = z = w = 0.0f;
    } else {
        const float scale = 1.0f / sqrtf(sqrLength);
        x *= scale;
        y *= scale;
        z *= scale;
        w *= scale;
    }
}

Quaternion Quaternion::Normalized() const {
    Quaternion normalized = *this;
    normalized.Normalize();
    return normalized;
}

float Quaternion::SqrMagnitude() const {
    return (x * x) + (y * y) + (z * z) + (w * w);
}

Matrix Quaternion::ToMatrix() const {
    return Matrix({
        1.0f - 2.0f * y * y - 2.0f * z * z,
        2.0f * x * y + 2.0f * z * w,
        2.0f * x * z - 2.0f * y * w,
        0.0f,

        2.0f * x * y - 2.0f * z * w,
        1.0f - 2.0f * x * x - 2.0f * z * z,
        2.0f * z * y + 2.0f * x * w,
        0.0f,

        2.0f * x * z + 2.0f * y * w,
        2.0f * z * y - 2.0f * x * w,
        1.0f - 2.0f * x * x - 2.0f * y * y,
        0.0f
    });
}

Quaternion Quaternion::operator* (const Quaternion& other) const {
    return Quaternion(
        w * other.x + x * other.w - y * other.z - z * other.y,
        w * other.y + y * other.w - z * other.x - x * other.z,
        w * other.z + z * other.w - x * other.y - y * other.x,
        w * other.w - x * other.x - y * other.y - z * other.z
    );
}

Quaternion Quaternion::FromEuler(const Vector& rotation) {
    const Vector halfrot = rotation * 0.5f;

	const float s1 = sinf(halfrot.x);
	const float c1 = cosf(halfrot.x);
	const float s2 = sinf(halfrot.y);
	const float c2 = cosf(halfrot.y);
	const float s3 = sinf(halfrot.z);
	const float c3 = cosf(halfrot.z);
	const float c1c2 = c1 * c2;
	const float s1s2 = s1 * s3;

	return Quaternion(
	    (c1c2 * s3) + (s1s2 * c3),
	    (s1 * c2 * c3) + (c1 * s2 * s3),
	    (c1 * s2 * c3) - (s1 * c2 * s3),
    	(c1c2 * c3) - (s1s2 * s3)
    );
}