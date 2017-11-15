#pragma once

#ifdef EMBED_RESOURCES
#include <rehover_data_gcr.h>
const unsigned char* ResourceLoader::embedded = rehover_data_gcr_txt;
#endif