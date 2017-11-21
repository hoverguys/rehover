#include "Vector.h"
#include <math.h>

void Vector::Normalize() {
    float rlength = 1.0f / Magnitude();
    x *= rlength;
    y *= rlength;
    z *= rlength;
}

Vector Vector::Normalized() {
    Vector normalized = *this;
    normalized.Normalize();
    return normalized;
}

Vector Vector::Cross(const Vector& other) {
    return Vector{
        (y * other.z) - (z * other.y),
        (z * other.x) - (x * other.z),
        (x * other.y) - (y * other.x)
    };
}

float Vector::Dot(const Vector& other){
    return (x * other.x) + (y * other.y) + (z * other.z);
}

float Vector::Magnitude() {
    return sqrtf(x * x + y * y + z * z);
}

float Vector::SqrMagnitude() {
    return x * x + y * y + z * z;
}

Vector Vector::operator* (const float& scale) {
    return Vector(x * scale, y * scale, z * scale);
}

Vector Vector::operator+ (const Vector& other) {
    return Vector(x + other.x, y + other.y, z + other.z);
}

Vector Vector::operator- (const Vector& other) {
    return Vector(x - other.x, y - other.y, z - other.z);
}

Vector::operator guVector() {
    return guVector{x,y,z};
}