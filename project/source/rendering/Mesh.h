#pragma once

struct MeshIndex {
    short unsigned int vertex;
    short unsigned int uv;
    short unsigned int normal;
};

class Mesh {
public:
    Mesh();
    ~Mesh();

    void Render();

protected:
    friend class MeshResource;

    float* positionArray;
    float* normalArray;
    float* uvArray;
    MeshIndex* indexArray;

    void* displayList;
    unsigned int displayListSize;
};