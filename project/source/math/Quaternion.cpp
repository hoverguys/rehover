#include "Quaternion.h"
#include "Matrix.h"
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
    Matrix mtx = Matrix({
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

    return mtx;
}

Quaternion Quaternion::operator* (const Quaternion& other) const {
    Quaternion result = {
        w * other.x + x * other.w - y * other.z - z * other.y,
        w * other.y + y * other.w - z * other.x - x * other.z,
        w * other.z + z * other.w - x * other.y - y * other.x,
        w * other.w - x * other.x - y * other.y - z * other.z
    };

    return result;
}