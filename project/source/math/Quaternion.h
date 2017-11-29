#pragma once
#include "Math.h"

class Quaternion {
public:
    float x, y, z, w;

    explicit Quaternion(float x, float y, float z, float w) : x(x), y(y), z(z), w(w) {}
    explicit Quaternion() : Quaternion(0, 0, 0, 1) {}

    void Normalize();
    Quaternion Normalized() const;
    float SqrMagnitude() const;

    Matrix ToMatrix() const;

    Quaternion operator* (const Quaternion& other) const;

    static Quaternion FromEuler(const Vector& rotation);
};