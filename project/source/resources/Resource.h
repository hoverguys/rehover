#pragma once

 class Resource {
 protected:
    Resource(void* address, unsigned int size) : address(address), size(size) {};
    virtual void Initialize() = 0;

    void* address;
    unsigned int size;
 };