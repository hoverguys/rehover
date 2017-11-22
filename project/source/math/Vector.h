#pragma once
#include "ogc/gu.h"

class Vector {
    public:
    float x, y, z;

    Vector(float x, float y, float z) : x(x), y(y), z(z) {}
    Vector() : Vector(0,0,0) {}

    void Normalize();
    Vector Normalized() const;
    Vector Cross(Vector other) const;
    float Dot(Vector other) const;
    float Magnitude() const;
    float SqrMagnitude() const;

    Vector operator* (const float& scale);
    Vector operator+ (const Vector& other);
    Vector operator- (const Vector& other);
    operator guVector() const;
};