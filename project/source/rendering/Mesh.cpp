#include "Mesh.h"
#include <malloc.h>
#include <ogc/gx.h>

Mesh::Mesh() {
    displayList = nullptr;
}

Mesh::~Mesh() {
    if (displayList != nullptr) {
        free(displayList);
    }
}

void Mesh::Render() {
    GX_CallDispList(displayList, displayListSize);
}