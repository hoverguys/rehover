#pragma once
#include "ogc/gu.h"

namespace Math {
    extern guVector worldUp;
    extern guVector worldForward;
    extern guVector worldRight;
    guQuaternion EulerToQuaternion(guVector rotation);
}