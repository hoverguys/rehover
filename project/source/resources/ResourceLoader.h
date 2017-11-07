#pragma once
#include <map>
#include "Resource.h"
#include "../utils/fnv.h"

typedef unsigned int FileHash;

class ResourceLoader {
public:
    static void LoadPack(const char* path);
    
    static std::shared_ptr<Resource> Load(FileHash hash);

    template<typename T>
    static constexpr std::shared_ptr<T> Load(const char* path) {
        static_assert(std::is_base_of<Resource, T>::value, "Must inherit of type Resource");
        return std::static_pointer_cast<T>(Load(fnv1_hash(path)));
    }

private:
    typedef std::map<FileHash, std::pair<unsigned int, unsigned int>> FileMap;
    static FileMap files;

    typedef std::map<FileHash, std::shared_ptr<Resource>> ResourceMap;
    static ResourceMap cache;

    static const char* packfile;
};