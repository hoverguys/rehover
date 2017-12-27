#pragma once

#include "../math/Rect.h"

/*! \brief All entries in a atlas */
struct Atlas {
	std::map<unsigned int, Rect> coordinates;
};