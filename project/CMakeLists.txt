# Project setup
cmake_minimum_required(VERSION 3.1)
set(CMAKE_TOOLCHAIN_FILE ${CMAKE_CURRENT_LIST_DIR}/devkitPPC.cmake)
set(CMAKE_CXX_STANDARD 11)
project(rehover C CXX)
list(APPEND CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/cmake)
include(SetupOGC)

# Get version if possible
include(FindGit OPTIONAL)
if(GIT_FOUND)
	execute_process(COMMAND
	  "${GIT_EXECUTABLE}" describe --tags --always --dirty
	  WORKING_DIRECTORY "${CMAKE_SOURCE_DIR}"
	  OUTPUT_VARIABLE GIT_VERSION
	  ERROR_QUIET OUTPUT_STRIP_TRAILING_WHITESPACE)
else()
	set(GIT_VERSION "unknown")
endif()
add_definitions(-DGIT_VERSION="${GIT_VERSION}")

# Sources
file(GLOB_RECURSE SOURCE_FILES source/*)
add_executable(rehover ${SOURCE_FILES})

# Dependencies
target_link_libraries(rehover ogc)

# Create .dol
add_dol_target(rehover)
