#include "Mesh.h"

#include "../pchheader.h"

Mesh::~Mesh() {
	if (displayList != nullptr) {
		free(displayList);
	}
}

void Mesh::Render() { GX_CallDispList(displayList, displayListSize); }