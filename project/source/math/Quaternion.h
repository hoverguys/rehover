#pragma once

class Matrix;

class Quaternion {
public:
    float x, y, z, w;

    void Normalize();
    Quaternion Normalized() const;
    float SqrMagnitude() const;

    Matrix ToMatrix() const;

    Quaternion operator* (const Quaternion& other);
};