#pragma once
#include <array>

class Quaternion;
class Vector;

class Matrix {
public:
    Matrix() {printf("*");};

    static Matrix Identity();
    static Matrix LookAt(Vector origin, Vector up, Vector target);
    static Matrix AxisAngle(Vector axis, float angle);

    void Scale(Vector scale);
    void Translate(Vector delta);

    //TODO: Swap func here?
    void Inverse();
    Matrix Inversed() const;
    void Transpose();
    Matrix Transposed() const;

    Vector Multiply(Vector vec) const;
    Vector MultiplySR(Vector vec) const;

    Matrix operator* (const Matrix& other);

    Quaternion ToQuaternion() const;
    void ToNative(Mtx matrix) const;
private:
    friend class Vector;
    friend class Quaternion;
    std::array<float, 3*4> internal;
};
