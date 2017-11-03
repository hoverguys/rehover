#include "Mesh.h"
#include <malloc.h>

Mesh::Mesh() {
    displayList = nullptr;
}

Mesh::~Mesh() {
    if (displayList != nullptr) {
        free(displayList);
    }
}