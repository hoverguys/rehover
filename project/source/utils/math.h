#pragma once
#include "ogc/gu.h"
#include "../math/Vector.h"
#include "../math/Quaternion.h"
namespace Math {
    extern Vector worldUp;
    extern Vector worldForward;
    extern Vector worldRight;
    Quaternion EulerToQuaternion(Vector rotation);
}