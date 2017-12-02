#pragma once
#include "Math.h"

class Vector {
    public:
    float x, y, z;

    Vector(float x, float y, float z) : x(x), y(y), z(z) {}
    Vector() : Vector(0,0,0) {}

    void Normalize();
    Vector Normalized() const;
    Vector Cross(const Vector& other) const;
    float Dot(const Vector& other) const;
    float Magnitude() const;
    float SqrMagnitude() const;

    Vector operator* (const float& scale) const;
    Vector operator* (const Vector& scale) const;
    Vector operator+ (const Vector& other) const;
    Vector operator- (const Vector& other) const;
    operator guVector() const;
};