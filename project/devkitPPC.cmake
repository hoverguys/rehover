# Based on DevkitArm3DS.cmake by Lectem (https://github.com/Lectem/3ds-cmake)
# Wintermute, if you're reading this, I'm sorry but not really.
# At least I fixed the casing.

set(CMAKE_SYSTEM_NAME Generic)
set(CMAKE_SYSTEM_PROCESSOR powerpc)
set(GEKKO TRUE) # To be used for multiplatform projects

# Variables
set(GCN TRUE CACHE BOOL "Build Gamecube version")
set(WII FALSE CACHE BOOL "Build Wii version")

# devkitPro paths are broken on windows, so we have to fix those
macro(msys_to_cmake_path MsysPath ResultingPath)
	if(WIN32)
		string(REGEX REPLACE "^/([a-zA-Z])/" "\\1:/" ${ResultingPath} "${MsysPath}")
	else()
		set(${ResultingPath} "${MsysPath}")
	endif()
endmacro()

msys_to_cmake_path("$ENV{DEVKITPRO}" DEVKITPRO)
if(NOT IS_DIRECTORY ${DEVKITPRO})
    message(FATAL_ERROR "Please set DEVKITPRO in your environment")
endif()

msys_to_cmake_path("$ENV{DEVKITPPC}" DEVKITPPC)
if(NOT IS_DIRECTORY ${DEVKITPPC})
    message(FATAL_ERROR "Please set DEVKITPPC in your environment")
endif()

# Prefix detection only works with compiler id "GNU"
# CMake will look for prefixed g++, cpp, ld, etc. automatically
if(WIN32)
    set(CMAKE_C_COMPILER "${DEVKITPPC}/bin/powerpc-eabi-gcc.exe")
    set(CMAKE_CXX_COMPILER "${DEVKITPPC}/bin/powerpc-eabi-g++.exe")
    set(CMAKE_AR "${DEVKITPPC}/bin/powerpc-eabi-gcc-ar.exe" CACHE STRING "")
    set(CMAKE_RANLIB "${DEVKITPPC}/bin/powerpc-eabi-gcc-ranlib.exe" CACHE STRING "")
else()
    set(CMAKE_C_COMPILER "${DEVKITPPC}/bin/powerpc-eabi-gcc")
    set(CMAKE_CXX_COMPILER "${DEVKITPPC}/bin/powerpc-eabi-g++")
    set(CMAKE_AR "${DEVKITPPC}/bin/powerpc-eabi-gcc-ar" CACHE STRING "")
    set(CMAKE_RANLIB "${DEVKITPPC}/bin/powerpc-eabi-gcc-ranlib" CACHE STRING "")
endif()

set(CMAKE_FIND_ROOT_PATH ${DEVKITPPC} ${DEVKITPRO})
set(CMAKE_FIND_ROOT_PATH_MODE_PROGRAM NEVER)
set(CMAKE_FIND_ROOT_PATH_MODE_LIBRARY ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_INCLUDE ONLY)
set(CMAKE_FIND_ROOT_PATH_MODE_PACKAGE ONLY)

SET(BUILD_SHARED_LIBS OFF CACHE INTERNAL "Shared libs not available" )

set(ARCH "-mcpu=750 -meabi -mhard-float")
set(CMAKE_C_FLAGS "${ARCH} -ffast-math --pedantic" CACHE STRING "C flags")
set(CMAKE_CXX_FLAGS "${CMAKE_C_FLAGS} -fno-exceptions -fno-rtti" CACHE STRING "C++ flags")