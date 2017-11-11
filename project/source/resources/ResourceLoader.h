#pragma once
#include "../utils/fnv.h"
#include "Resource.h"
#include <map>

typedef unsigned int FileHash;

class ResourceLoader {
public:
	static void LoadPack(const char* path);

	template <typename T> static std::shared_ptr<T> Load(FileHash hash) {
		auto entry = cache.find(hash);
		if (entry != cache.end()) {
			return std::static_pointer_cast<T>(entry->second);
		}

		auto file = files.find(hash);
		if (file == files.end()) {
			printf("File %08x not found\n", hash);
		}
		auto info = file->second;

#ifndef EMBED_RESOURCES
		// Allocate buffer for resource
		// Load from pack
		printf("Loading from file");
		return nullptr;
#else
		unsigned char* address = (unsigned char*)embedded;
		address += info.first; // Add offset
		printf("Loading from memory @ 0x%08x\n", address);
#endif

		auto resource = std::make_shared<T>((void*)address, (unsigned int)info.second);
		resource->Initialize();

		// Cache loaded resource
		cache.emplace(hash, resource);

		return resource;
	}

	template <typename T> static constexpr std::shared_ptr<T> Load(const char* path) {
		static_assert(std::is_base_of<Resource, T>::value, "Must inherit of type Resource");
		return Load<T>(fnv1_hash(path));
	}

private:
	typedef std::map<FileHash, std::pair<unsigned int, unsigned int>> FileMap;
	static FileMap files;

	typedef std::map<FileHash, std::shared_ptr<Resource>> ResourceMap;
	static ResourceMap cache;

	static const char* packfile;

#ifdef EMBED_RESOURCES
	static const unsigned char* embedded;
#endif
};