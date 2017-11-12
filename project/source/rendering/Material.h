#pragma once
#include "Texture.h"

struct Material {
    std::array<std::shared_ptr<Texture>, 8> textures;
};