#pragma once
#include "../pchheader.h"
#include "Resource.h"

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

	template <typename T> static constexpr std::shared_ptr<T> Load(const char* path) {
		static_assert(std::is_base_of<Resource, T>::value, "Must inherit of type Resource");
		return Load<T>(fnv1_hash(path));
	}

	static void UnloadUnused() {
#ifdef EMBED_RESOURCES
		auto removed = 0;
		for (auto it = cache.cbegin(); it != cache.cend();) {
			if (it->second.use_count() <= 1) {
				// Removing it from the cache should get rid of the last reference
				// and cause the destructor to be called.
				it = cache.erase(it);
				removed++;
			} else {
				++it;
			}
		}
		std::printf("Unloaded %d resources\n", removed);
#else
		std::printf("Asset unloading not supported in embedded mode\n");
#endif

	}

private:

	static constexpr unsigned int fnv1_hash(const char* buffer) {
		const unsigned int fnv_prime32 = 16777619;
		unsigned int result = 2166136261;
		int i=0;
		while(buffer[i] != '\0') {
			result *= fnv_prime32;
			result ^= (unsigned int)buffer[i++];
		}
		return result;
	}

	typedef std::map<FileHash, std::pair<unsigned int, unsigned int>> FileMap;
	static FileMap files;

	typedef std::map<FileHash, std::shared_ptr<Resource>> ResourceMap;
	static ResourceMap cache;

	static const char* packfile;

#ifdef EMBED_RESOURCES
	static const unsigned char* embedded;
#endif
};