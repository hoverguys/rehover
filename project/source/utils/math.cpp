#include "./math.h"
#include <math.h>

namespace Math {

guVector worldUp = { 0, 1, 0 };
guVector worldForward = { 0, 0, -1 };
guVector worldRight = { 1, 0, 0 };

guQuaternion EulerToQuaternion(guVector rotation) {
	guVecScale(&rotation, &rotation, 0.5f);

	float c1, s1, c2, s2, c3, s3;

	s1 = sinf(rotation.x);
	c1 = cosf(rotation.x);
	s2 = sinf(rotation.y);
	c2 = cosf(rotation.y);
	s3 = sinf(rotation.z);
	c3 = cosf(rotation.z);
	float c1c2 = c1 * c2;
	float s1s2 = s1 * s3;

	guQuaternion result;
	result.w = c1c2 * c3 - s1s2 * s3;
	result.x = c1c2 * s3 + s1s2 * c3;
	result.y = s1 * c2 * c3 + c1 * s2 * s3;
	result.z = c1 * s2 * c3 - s1 * c2 * s3;
	return result;
}
} // namespace Math