#include "ShaderResource.h"

void ShaderResource::Initialize() {
	unsigned char* data = static_cast<unsigned char*>(address);
	auto s = std::make_shared<Shader>();
	s->data = data;
	s->size = size;

	internal = s;
}

std::shared_ptr<Shader> ShaderResource::Load() { return internal; }