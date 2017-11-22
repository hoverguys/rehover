#pragma once
#include <array>

class Quaternion;
class Vector;

class Matrix {
public:
    Matrix() {};
    Matrix(std::array<float, 3*4> data) : internal(data) {}

    static Matrix Identity();
    static Matrix LookAt(Vector origin, Vector up, Vector target);
    static Matrix AxisAngle(Vector axis, float angle);

    void Scale(const Vector& scale);
    void Translate(const Vector& delta);

    //TODO: Swap func here?
    void Inverse();
    Matrix Inversed() const;
    void Transpose();
    Matrix Transposed() const;

    Vector Multiply(const Vector& vec) const;
    Vector MultiplySR(const Vector& vec) const;

    Matrix operator* (const Matrix& other) const;

    Quaternion ToQuaternion() const;
    void ToNative(Mtx matrix) const;
private:
    friend class Vector;
    friend class Quaternion;
    std::array<float, 3*4> internal;
};
