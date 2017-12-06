#pragma once
#include "Math.h"

/*! \brief 4x3 Matrix
 */
class Matrix {
public:
	/*! \brief Create a zero matrix
	 *  This should never be used, if you want an empty matrix use Matrix::Identity
	 */
	Matrix(){};

	/*! \brief Create a matrix from an array of values
	 *  \param data 12 (4*3) values to fill the matrix with
	 */
	explicit Matrix(std::array<float, 3 * 4> data) : internal(data) {}

	/*! \brief Get identity matrix
	 *  | 1 0 0 0 |
	 *  | 0 1 0 0 |
	 *  | 0 0 1 0 |
	 *  \return Identity matrix
	 */
	static Matrix Identity();

	/*! \brief Get transform matrix that looks towards a certain point from a certain point
	 *  \param origin Position to look from
	 *  \param up     Up vector
	 *  \param target Point to look towards
	 *  \return Transform matrix at origin that looks at target
	 */
	static Matrix LookAt(Vector origin, Vector up, Vector target);

	/*! \brief Get rotation matrix that's rotated a certain amount on a axis
	 *  \param axis Rotation axis
	 *  \param angle Rotation amount
	 *  \return Rotation matrix
	 */
	static Matrix AxisAngle(Vector axis, float angle);

	/*! \brief Scale matrix in place by a certain amount (non-uniform)
	 *  \param scale Scale factor (in all 3 directions)
	 */
	void Scale(const Vector& scale);

	/*! \brief Translate matrix in place by a certain amount
	 *  \param delta Movement delta (in all 3 directions)
	 */
	void Translate(const Vector& delta);

	/*! \brief Inverse matrix in place
		\todo Swap func here?
	 */
	void Inverse();

	/*! \brief Get an inversed copy of the matrix
	 *  \return Inversed matrix
	 */
	Matrix Inversed() const;

	/*! \brief Transpose the matrix in place
	 */
	void Transpose();

	/*! \brief Get a transposed copy of the matrix
	 *  \return Transposed matrix
	 */
	Matrix Transposed() const;

	/*! \brief Multiply matrix with a column vector
	 *  \param vec Column vector to multiply with
	 *  \return Matrix product
	 */
	Vector Multiply(const Vector& vec) const;

	/*! \brief Multiply square matrix (3x3) with a column vector
	 *  This is pretty much the same as Multiply but it ignores the 4th column
	 *  \param vec Column vector to multiply with
	 *  \return Matrix product
	 */
	Vector MultiplySR(const Vector& vec) const;

	/*! \brief Perform matrix product between two matrices
	 *  \param other Other matrix to multiply with
	 *  \return Matrix product
	 */
	Matrix operator*(const Matrix& other) const;

	/*! \brief Convert rotation matrix to quaternion
	 *  \return Quatertnion from rotation matrix
	 */
	Quaternion ToQuaternion() const;

	/*! \brief Copy matrix data to libogc's native format (Mtx)
	 *  \param matrix Mtx object to fill with data
	 */
	void ToNative(Mtx matrix) const;

private:
	friend class Vector;
	friend class Quaternion;
	std::array<float, 3 * 4> internal;
};
