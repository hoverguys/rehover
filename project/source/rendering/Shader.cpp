#include "Shader.h"
#include <ogc/gx.h>

void Shader::Use() { GX_CallDispList(data, size); }