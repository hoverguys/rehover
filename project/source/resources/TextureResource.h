#pragma once

#include "../pchheader.h"

#include "../rendering/Texture.h"
#include "Resource.h"

/*! \brief BTB file header
 */
struct TextureResourceHeader {
	unsigned short width;       /*< Texture width  */
	unsigned short height;      /*< Texture height */
	unsigned char format;       /*< Color format */
	unsigned short maxlod : 4;  /*< Max LOD (0-10) */
	unsigned short minlod : 4;  /*< Min LOD (0-10) */
	unsigned short wrapS : 3;   /*< Wrap S */
	unsigned short wrapT : 3;   /*< Wrap T */
	unsigned short filter : 2;  /*< Texture filtering mode */
	unsigned int dataOffset;    /*< Offset to texture data */
	unsigned int paletteOffset; /*< Offset to palette data (if applicable) */
} __attribute__((packed));

/*!
 * \brief Texture resource loader
 * Loads a texture from a BTB file loaded in memory
 */
class TextureResource : public Resource {
public:
	/*!
	 * \brief Load the texture from a BTB file
	 *
	 * \param base Pointer to BTB file
	 * \param size Size of the BTB file
	 */
	TextureResource(void* base, unsigned int size) : Resource(base, size) {}

	/*!
	 * \brief Load and get the texture
	 *
	 * \return Pointer to the loaded texture
	 */
	std::shared_ptr<Texture> Load();

	long int ReferenceCount() override {
		return internal == NULL ? 0 : internal.use_count();
	}

	void Initialize() override;

private:
	TextureResourceHeader* header = nullptr;
	bool loaded = false;
	std::shared_ptr<Texture> internal;
};
