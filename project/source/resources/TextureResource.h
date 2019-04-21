#pragma once

#include "../pchheader.h"

#include "../rendering/Texture.h"
#include "Resource.h"

/*! \brief BTB file header
 */
struct TextureResourceHeader {
	u16 width;         /*< Texture width  */
	u16 height;        /*< Texture height */
	u8 format;         /*< Color format */
	u8 palfmt;         /*< Palette color format */
	u16 maxlod : 4;    /*< Max LOD (0-10) */
	u16 minlod : 4;    /*< Min LOD (0-10) */
	u16 wrapS : 3;     /*< Wrap S */
	u16 wrapT : 3;     /*< Wrap T */
	u16 filter : 2;    /*< Texture filtering mode */
	u32 dataOffset;    /*< Offset to texture data */
	u32 paletteOffset; /*< Offset to palette data (if applicable) */
	u16 paletteLength; /*< How many palette entries (max 16384) */
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

	long int ReferenceCount() override { return internal == NULL ? 0 : internal.use_count(); }

	void Initialize() override;

private:
	TextureResourceHeader* header = nullptr;
	bool loaded = false;
	std::shared_ptr<Texture> internal;
};
