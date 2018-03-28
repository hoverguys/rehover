#pragma once

#include "../pchheader.h"

#include "../rendering/Shader.h"
#include "Resource.h"

/*! \brief Shader resource loader
 *  Loads a shader from a shader file in memory
 */
class ShaderResource : public Resource {
public:
	/*!
	 * \brief Load a shader from a shader file
	 *
	 * \param base Pointer to the shader file
	 * \param size Size of the shader file
	 */
	ShaderResource(void* base, unsigned int size) : Resource(base, size) {}

	/*!
	 * \brief Load and get the shader
	 *
	 * \return Pointer to the loaded shader
	 */
	std::shared_ptr<Shader> Load();

	long int ReferenceCount() override {
		return internal == NULL ? 0 : internal.use_count();
	}

	void Initialize() override;

private:
	std::shared_ptr<Shader> internal;
};
