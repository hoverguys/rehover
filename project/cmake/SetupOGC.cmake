# DevkitPro paths are broken on windows, so we have to fix those
macro(msys_to_cmake_path MsysPath ResultingPath)
    string(REGEX REPLACE "^/([a-zA-Z])/" "\\1:/" ${ResultingPath} "${MsysPath}")
endmacro()

if(NOT DEVKITPRO)
    msys_to_cmake_path("$ENV{DEVKITPRO}" DEVKITPRO)
endif()

set(LIBOGC_PATHS $ENV{LIBOGC} libogc ${DEVKITPRO}/libogc)

# Set default portlib
set(PORTLIB_PATH ${DEVKITPRO}/portlibs/ppc CACHE STRING "Path to portlibs (if used)")
set(PORT_INCLUDE_DIR ${PORTLIB_PATH}/include)
set(PORT_LIBRARY_DIR_GCN ${PORTLIB_PATH}/lib/cube)
set(PORT_LIBRARY_DIR_WII ${PORTLIB_PATH}/lib/wii)
include_directories(${PORT_INCLUDE_DIR})

# Find libogc
find_path(LIBOGC_INCLUDE_DIR gccore.h
          PATHS ${LIBOGC_PATHS}
          PATH_SUFFIXES include libogc/include )

find_path(LIBOGC_LIBRARY_DIR_GCN NAMES libogc.a
          PATHS ${LIBOGC_PATHS}
          PATH_SUFFIXES lib/cube libogc/lib/cube )

find_path(LIBOGC_LIBRARY_DIR_WII NAMES libogc.a
          PATHS ${LIBOGC_PATHS}
          PATH_SUFFIXES lib/wii libogc/lib/wii )

include(FindPackageHandleStandardArgs)
# handle the QUIETLY and REQUIRED arguments and set LIBOGC_FOUND to TRUE
# if all listed variables are TRUE
find_package_handle_standard_args(LIBOGC  DEFAULT_MSG
                                  LIBOGC_LIBRARY_DIR_GCN LIBOGC_LIBRARY_DIR_WII LIBOGC_INCLUDE_DIR)

mark_as_advanced(LIBOGC_INCLUDE_DIR_GCN LIBOGC_LIBRARY_DIR_WII LIBOGC_LIBRARY_DIR)
if(LIBOGC_FOUND)
    set(LIBOGC ${LIBOGC_INCLUDE_DIR}/..)
    message(STATUS "Setting LIBOGC to ${LIBOGC}")

    # Add paths it to toolchain
    include_directories(${LIBOGC_INCLUDE_DIR})
else()
    message(FATAL_ERROR "Could not find libogc")
endif()

# Find elf2dol (required for add_dol_target)
if(NOT ELF2DOL)
    find_program(ELF2DOL elf2dol ${DEVKITPPC}/bin ${DEVKITPRO}/tools/bin)
    if(ELF2DOL)
        message(STATUS "elf2dol: ${ELF2DOL} - found")
    else()
        message(WARNING "elf2dol - not found")
    endif()
endif()
mark_as_advanced(ELF2DOL)

# Function to make .elf and .dols
# Usage:
#     add_dol_target(<target>)
function(add_dol_target target)
    # First, make sure the file has the proper extension (otherwise Dolphin won't load it)
    set_target_properties(${target} PROPERTIES SUFFIX ".elf")
    add_custom_target(${target}_dol ALL
                      ${ELF2DOL} $<TARGET_FILE:${target}> ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol
                      DEPENDS ${CMAKE_CURRENT_BINARY_DIR}/${target}.elf
                      COMMENT "Generating .dol files from compiled .elf"
                      VERBATIM)
    set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol)
endfunction()

# Apply add_dol_target to a multi-target
# Usage:
#     add_dol_targets(<target>)
function(add_dol_targets target)
    if(GCN)
        add_dol_target(${target}_gcn)
    endif()
    if(WII)
        add_dol_target(${target}_wii)
    endif()
endfunction()

# Add one or more libraries from dkP to all available targets
# Usage:
#     add_default_libraries(<target> <lib1> [<lib2> ..])
function(add_default_libraries target)
    foreach(libname ${ARGN})
        if(GCN)
            target_link_libraries(${target}_gcn "${LIBOGC_LIBRARY_DIR_GCN}/lib${libname}.a")
        endif()
        if(WII)
            target_link_libraries(${target}_wii "${LIBOGC_LIBRARY_DIR_WII}/lib${libname}.a")
        endif()
    endforeach()
endfunction()

# Add one or more libraries from ppc-portlibs to all available targets
# Usage:
#     add_default_libraries(<target> <lib1> [<lib2> ..])
function(add_port_libraries target)
    foreach(libname ${ARGN})
        if(GCN)
            target_link_libraries(${target}_gcn "${PORT_LIBRARY_DIR_GCN}/lib${libname}.a")
        endif()
        if(WII)
            target_link_libraries(${target}_wii "${PORT_LIBRARY_DIR_WII}/lib${libname}.a")
        endif()
    endforeach()
endfunction()

# Create both GCN and WII version
# Usage:
#     add_multi_target(<target>)
function(add_multi_target target)
    if(GCN)
        add_executable(${target}_gcn ${ARGN})
        target_compile_options(${target}_gcn PUBLIC -DGEKKO -mogc)
        target_link_libraries(${target}_gcn "-mogc")
    endif()
    if(WII)
        add_executable(${target}_wii ${ARGN})
        target_compile_options(${target}_wii PUBLIC -DWII -mrvl)
        target_link_libraries(${target}_wii "-mrvl")
    endif()
endfunction()

# Add a dependency to all targets inside a multi-target
# Usage:
#     add_multi_target(<target> <dep1> [<dep2> ..])
function(add_multi_dependency target)
    if(GCN)
        add_dependencies(${target}_gcn ${ARGN})
    endif()
    if(WII)
        add_dependencies(${target}_wii ${ARGN})
    endif()
endfunction()