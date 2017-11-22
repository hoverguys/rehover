#pragma once

class Matrix;

class Quaternion {
public:
    float x, y, z, w;

    Quaternion(float x, float y, float z, float w) : x(x), y(y), z(z), w(w) {}
    Quaternion() : Quaternion(0, 0, 0, 1) {}

    void Normalize();
    Quaternion Normalized() const;
    float SqrMagnitude() const;

    Matrix ToMatrix() const;

    Quaternion operator* (const Quaternion& other) const;
};