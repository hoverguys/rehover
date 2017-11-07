#pragma once

class Resource {
public:
    Resource(void* address, unsigned int size) : address(address), size(size) {};
    virtual void Initialize() = 0;
protected:
    void* address;
    unsigned int size;
 };