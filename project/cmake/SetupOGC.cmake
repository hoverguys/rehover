# DevkitPro paths are broken on windows, so we have to fix those
macro(msys_to_cmake_path MsysPath ResultingPath)
    string(REGEX REPLACE "^/([a-zA-Z])/" "\\1:/" ${ResultingPath} "${MsysPath}")
endmacro()

if(NOT DEVKITPRO)
    msys_to_cmake_path("$ENV{DEVKITPRO}" DEVKITPRO)
endif()

set(LIBOGC_PATHS $ENV{LIBOGC} libogc ${DEVKITPRO}/libogc)

# Find libogc
find_path(LIBOGC_INCLUDE_DIR gccore.h
          PATHS ${LIBOGC_PATHS}
          PATH_SUFFIXES include libogc/include )

if(WII)
       find_path(LIBOGC_LIBRARY_DIR NAMES libogc.a
                 PATHS ${LIBOGC_PATHS}
                 PATH_SUFFIXES lib/wii libogc/lib/wii )
else()
       find_path(LIBOGC_LIBRARY_DIR NAMES libogc.a
                 PATHS ${LIBOGC_PATHS}
                 PATH_SUFFIXES lib/cube libogc/lib/cube )
endif()

include(FindPackageHandleStandardArgs)
# handle the QUIETLY and REQUIRED arguments and set LIBOGC_FOUND to TRUE
# if all listed variables are TRUE
find_package_handle_standard_args(LIBOGC  DEFAULT_MSG
                                  LIBOGC_LIBRARY_DIR LIBOGC_INCLUDE_DIR)

mark_as_advanced(LIBOGC_INCLUDE_DIR LIBOGC_LIBRARY_DIR)
if(LIBOGC_FOUND)
    set(LIBOGC ${LIBOGC_INCLUDE_DIR}/..)
    message(STATUS "setting LIBOGC to ${LIBOGC}")

    # Add paths it to toolchain
    include_directories(${LIBOGC_INCLUDE_DIR})
    link_directories(${LIBOGC_LIBRARY_DIR})
else()
    message(FATAL_ERROR "Could not find libogc")
endif()

# Find elf2dol (required for add_dol_target)
if(NOT ELF2DOL)
    find_program(ELF2DOL elf2dol ${DEVKITPPC}/bin)
    if(ELF2DOL)
        message(STATUS "elf2dol: ${ELF2DOL} - found")
    else()
        message(WARNING "elf2dol - not found")
    endif()
endif()

# Function to make .elf and .dols
function(add_dol_target target)
    # First, make sure the file has the proper extension (otherwise Dolphin won't load it)
    set_target_properties(${target} PROPERTIES SUFFIX ".elf")

    add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol
                       COMMAND ${ELF2DOL} $<TARGET_FILE:${target}> ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol
                       DEPENDS ${target}
                       COMMENT "Generating .dol files from compiled .elf"
                       VERBATIM)
    add_custom_target(${target}_dol ALL DEPENDS ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol)
endfunction()