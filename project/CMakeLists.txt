# Project setup
cmake_minimum_required(VERSION 3.1)
set(CMAKE_TOOLCHAIN_FILE ${CMAKE_CURRENT_LIST_DIR}/devkitPPC.cmake)
set(CMAKE_CXX_STANDARD 14)
project(rehover C CXX ASM)
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/cmake)
include(SetupOGC)
include(Tools)
include(PrecompiledHeader)
include(GitVersion)

# Options
set(EMBED_RESOURCES true CACHE BOOL "Embed resource pack into the binary")

# Sources
file(GLOB_RECURSE SOURCE_FILES source/*)
add_multi_target(rehover "${SOURCE_FILES}")
add_multi_precompiled_header(rehover source/pchheader.h)

# Dependencies
add_default_libraries(rehover ogc)
add_port_libraries(rehover entityx)

# Create resource pack
include(../Assets)
add_resource_pack(rehover-data ${ASSETS_PREFIX} ${ASSETS})

if(EMBED_RESOURCES)
	get_property(RESOURCEFILE TARGET rehover-data PROPERTY OUTPUT_NAME)
	embed_resources(rehover ${RESOURCEFILE})
	add_definitions(-DEMBED_RESOURCES)
else()
	add_multi_dependency(rehover rehover-data)
endif()

# Create .dol
add_dol_targets(rehover)
