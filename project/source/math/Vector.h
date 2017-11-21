#pragma once
#include "ogc/gu.h"

class Vector {
    public:
    float x, y, z;

    Vector(float x, float y, float z) : x(x), y(y), z(z) {}
    Vector() : Vector(0,0,0) {}

    void Normalize();
    Vector Normalized();
    Vector Cross(const Vector& other);
    float Dot(const Vector& other);
    float Magnitude();
    float SqrMagnitude();

    Vector operator* (const float& scale);
    Vector operator+ (const Vector& other);
    Vector operator- (const Vector& other);
    operator guVector();
};