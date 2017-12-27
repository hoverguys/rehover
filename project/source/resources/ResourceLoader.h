#pragma once
#include "../pchheader.h"
#include "Resource.h"

#include "AtlasResource.h"

/*! FNV1 hash of file paths */
typedef unsigned int FileHash;

/*! Atlas texture path and coordinate to a specific sprite */
typedef std::pair<FileHash, Rect> SpriteLocation;

class ResourceLoader {
public:
	static void LoadPack(const char* path);

	template <typename T>
	static std::shared_ptr<T> Load(const FileHash hash) {
		auto entry = cache.find(hash);
		if (entry != cache.end()) {
			return std::static_pointer_cast<T>(entry->second);
		}

		auto file = files.find(hash);
		if (file == files.end()) {
			std::printf("File %08x not found\n", hash);
			return nullptr;
		}
		auto info = file->second;

#ifndef EMBED_RESOURCES
		// Allocate buffer for resource
		// Load from pack
		std::printf("Loading from file");
		return nullptr;
#else
		unsigned char* address = (unsigned char*)embedded;
		address += info.first; // Add offset
		std::printf("Loading file %08x from memory at address: %p\n", hash, address);
#endif

		auto resource = std::make_shared<T>((void*)address, (unsigned int)info.second);
		resource->Initialize();

		// Cache loaded resource
		cache.emplace(hash, resource);

		return resource;
	}

	template <typename T>
	static constexpr std::shared_ptr<T> Load(const char* path) {
		static_assert(std::is_base_of<Resource, T>::value, "Must inherit of type Resource");
		return Load<T>(fnv1_hash(path));
	}

	static std::shared_ptr<Atlas> LoadAtlas(const FileHash hash) {
		auto atlasResource = Load<AtlasResource>(hash);
		auto atlas = atlasResource->Load();
		// TODO
		return atlas;
	}

	static constexpr std::shared_ptr<Atlas> LoadAtlas(const char* path) { return LoadAtlas(fnv1_hash(path)); }

	static SpriteLocation GetSprite(const FileHash hash) {
		auto it = sprites.find(hash);
		if (it != sprites.end()) {
			return it->second;
		}
		/*! \todo Fallback: load the sprite as a texture */
		std::printf("Sprite %08x not found\n", hash);
		return SpriteLocation{};
	}

	static constexpr SpriteLocation GetSprite(const char* path) { return GetSprite(fnv1_hash(path)); }

private:
	static constexpr unsigned int fnv1_hash(const char* buffer) {
		const unsigned int fnv_prime32 = 16777619;
		unsigned int result = 2166136261;
		int i = 0;
		while (buffer[i] != '\0') {
			result *= fnv_prime32;
			result ^= (unsigned int)buffer[i++];
		}
		return result;
	}

	typedef std::map<FileHash, std::pair<unsigned int, unsigned int>> FileMap;
	static FileMap files;

	typedef std::map<FileHash, std::shared_ptr<Resource>> ResourceMap;
	static ResourceMap cache;

	typedef std::map<FileHash, SpriteLocation> AtlasMap;
	static AtlasMap sprites;

	static const char* packfile;

#ifdef EMBED_RESOURCES
	static const unsigned char* embedded;
#endif
};