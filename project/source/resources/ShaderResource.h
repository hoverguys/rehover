#pragma once

#include <memory>

#include "../rendering/Shader.h"
#include "Resource.h"

class ShaderResource : public Resource {
public:
	ShaderResource(void* base, unsigned int size) : Resource(base, size) {}
	std::shared_ptr<Shader> Load();
	void Initialize() override;

private:
	std::shared_ptr<Shader> internal;
};
