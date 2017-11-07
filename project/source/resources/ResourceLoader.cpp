#include "ResourceLoader.h"

#include <stdio.h>

struct PackHeader {
    unsigned int fileCount;
};

struct PackEntry {
    FileHash hash;
    unsigned int offset;
    unsigned int size;
};

#ifdef EMBED_RESOURCES
#include <rehover_data_gcr.h>
#endif

ResourceLoader::FileMap ResourceLoader::files;
ResourceLoader::ResourceMap ResourceLoader::cache;
const char* ResourceLoader::packfile = nullptr;

void ResourceLoader::LoadPack(const char* path) {
#ifndef EMBED_RESOURCES
    //Load header from file into allocated memory
    //libfat for loating
    packfile = path;
    char* address = 0;
    printf("Loading pack from file %s\n", path);
#else
    char* address = (char*)rehover_data_gcr_txt;
    printf("Loading pack from memory @ 0x%08x\n", address);
#endif

    //Set header pointer
    PackHeader* header = reinterpret_cast<PackHeader*>(address);
    printf("Pack contains %d file(s)\n", header->fileCount);

    address += sizeof(PackHeader);
    for(int i = 0; i < header->fileCount; ++i, address += sizeof(PackEntry)) {
        PackEntry* entry = reinterpret_cast<PackEntry*>(address);

        files.emplace(entry->hash, std::pair<unsigned int, unsigned int>(entry->offset, entry->size));
        printf("Found file %08x at %u, size %u bytes", entry->hash, entry->offset, entry->size);
    }
}

std::shared_ptr<Resource> ResourceLoader::Load(FileHash hash) {
    auto entry = cache.find(hash);
    if (entry != cache.end()) {
        return entry->second;
    }

    auto file = files.find(hash);
    assert(file != files.end());
    auto info = file->second;

#ifndef EMBED_RESOURCES
    // Allocate buffer for resource
    // Load from pack
#else
    char* address = (char*)rehover_data_gcr_txt;
    address += info.first; //Add offset

    auto resource = new Resource((void*)address, info.second);
#endif



    return nullptr;
}