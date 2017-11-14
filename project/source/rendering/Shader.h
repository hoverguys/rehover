#pragma once

class Shader {
public:
	/*! \brief Calls the display list to set the TEV up accordingly
	 */
	void Use();

protected:
	friend class ShaderResource;

	unsigned char* data;
	unsigned int size;
};