#pragma once

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