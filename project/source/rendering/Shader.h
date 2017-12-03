#pragma once

class Shader {
public:
	/*! \brief Calls the display list to set the TEV up accordingly
	 */
	void Use();

	/*! \brief Reset to a default shader
	 */
	static void Default();

	/*! \brief Reset to a default unlit shader
	 */
	static void DefaultUnlit();

protected:
	friend class ShaderResource;

	unsigned char* data;
	unsigned int size;
};