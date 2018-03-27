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
const char* ResourceLoader::packfile = nullptr;
#endif

void ResourceLoader::LoadPack(const char* path) {
#ifndef EMBED_RESOURCES
	// Load header from file into allocated memory
	packfile = path;
	std::FILE* fp = std::fopen(packfile, "rb");
    assert(fp);
 
	// Read header
	PackHeader tmp;
	std::fread(&tmp, sizeof(PackHeader), 1, fp);

	// Calculate header size and read in
	auto headerSize = sizeof(PackHeader) + sizeof(PackEntry) * tmp.fileCount;
	void* packHeader = malloc(headerSize);

	// Copy header into memory
	std::fseek(fp, 0, SEEK_SET);
	std::fread(packHeader, headerSize, 1, fp);
	std::fclose(fp);

	unsigned char* address = (unsigned char*)packHeader;
	std::printf("Loading pack from file %s\n", packfile);
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
#ifndef EMBED_RESOURCES
	free(packHeader);
#endif
}