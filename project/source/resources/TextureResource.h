#pragma once

#include <memory>

#include "../rendering/Texture.h"
#include "Resource.h"

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

class TextureResource : public Resource {
public:
	TextureResource(void* base, unsigned int size) : Resource(base, size) {}
	std::shared_ptr<Texture> Load();
	void Initialize() override;

private:
	TextureResourceHeader* header = nullptr;
	bool loaded = false;
	std::shared_ptr<Texture> internal;
};
