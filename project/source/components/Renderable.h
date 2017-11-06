#pragma once
#include "rendering/Mesh.h"

namespace Components {
    struct Renderable {
        Renderable(Mesh* mesh) : mesh(mesh) {}
        
        Mesh* mesh;
    };
};