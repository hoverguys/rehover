#include "ResourceLoader.h"
#include "EmbeddedResource.h"

#include <stdio.h>

struct PackHeader {
	unsigned int fileCount;
};

struct PackEntry {
	FileHash hash;
	unsigned int offset;
	unsigned int size;
};

ResourceLoader::FileMap ResourceLoader::files;
ResourceLoader::ResourceMap ResourceLoader::cache;
const char* ResourceLoader::packfile = nullptr;

void ResourceLoader::LoadPack(const char* path) {
#ifndef EMBED_RESOURCES
	// Load header from file into allocated memory
	// libfat for loating
	packfile = path;
	unsigned char* address = 0;
	printf("Loading pack from file %s\n", path);
#else
	unsigned char* address = (unsigned char*)rehover_data_gcr_txt;
	printf("Loading pack from memory @ %p\n", address);
#endif

	// Set header pointer
	PackHeader* header = reinterpret_cast<PackHeader*>(address);
	printf("Pack contains %d file(s)\n", header->fileCount);

	address += sizeof(PackHeader);
	for (unsigned int i = 0; i < header->fileCount; ++i, address += sizeof(PackEntry)) {
		PackEntry* entry = reinterpret_cast<PackEntry*>(address);

		auto info = std::pair<unsigned int, unsigned int>(entry->offset, entry->size);
		files.emplace(entry->hash, info);
		printf("Found file %08x at %u, size %u bytes\n", entry->hash, info.first, info.second);
	}
}