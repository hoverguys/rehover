#include "Vector.h"

void Vector::Normalize() {
    float rlength = 1.0f / Magnitude();
    x *= rlength;
    y *= rlength;
    z *= rlength;
}

Vector Vector::Normalized() const {
    Vector normalized = *this;
    normalized.Normalize();
    return normalized;
}

Vector Vector::Cross(const Vector& other) const {
    return Vector(
        (y * other.z) - (z * other.y),
        (z * other.x) - (x * other.z),
        (x * other.y) - (y * other.x)
    );
}

float Vector::Dot(const Vector& other) const {
    return (x * other.x) + (y * other.y) + (z * other.z);
}

float Vector::Magnitude() const {
    return sqrt(x * x + y * y + z * z);
}

float Vector::SqrMagnitude() const {
    return x * x + y * y + z * z;
}

Vector Vector::operator* (const float& scale) const {
    return Vector(x * scale, y * scale, z * scale);
}

Vector Vector::operator* (const Vector& scale) const {
    return Vector(x * scale.x, y * scale.y, z * scale.z);
}

Vector Vector::operator+ (const Vector& other) const {
    return Vector(x + other.x, y + other.y, z + other.z);
}

Vector Vector::operator- (const Vector& other) const {
    return Vector(x - other.x, y - other.y, z - other.z);
}

Vector::operator guVector() const {
    return guVector{x,y,z};
}