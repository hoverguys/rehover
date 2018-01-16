#include "ResourceLoader.h"
#include "EmbeddedResource.h"

#include <cstdio>

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

#ifndef EMBED_RESOURCES
void const* ResourceLoader::packfile = nullptr;
#endif

void ResourceLoader::LoadPack(const char* path) {
#ifndef EMBED_RESOURCES
	// Load header from file into allocated memory
	std::FILE* fp = std::fopen(path, "rb");
    assert(fp);
 
	// Get file size
    std::fseek(fp, 0, SEEK_END);
    std::size_t filesize = std::ftell(fp);
    std::fseek(fp, 0, SEEK_SET);

	// Copy file into aligned memory
	packfile = memalign(32, filesize);
	std::fread(packfile, filesize, 1, fp);
	std::fclose(fp);

	unsigned char* address = packfile;
	std::printf("Loading pack from file %s\n", path);
#else
	unsigned char* address = (unsigned char*)rehover_data_gcr_txt;
	std::printf("Loading pack from memory @ %p\n", address);
#endif

	// Set header pointer
	PackHeader* header = reinterpret_cast<PackHeader*>(address);
	std::printf("Pack contains %u file(s)\n", header->fileCount);

	address += sizeof(PackHeader);
	for (unsigned int i = 0; i < header->fileCount; ++i, address += sizeof(PackEntry)) {
		PackEntry* entry = reinterpret_cast<PackEntry*>(address);

		auto info = std::pair<unsigned int, unsigned int>(entry->offset, entry->size);
		files.emplace(entry->hash, info);
		std::printf("Found file %08x at %u, size %u bytes\n", entry->hash, info.first, info.second);
	}
}