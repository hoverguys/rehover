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
    message(STATUS "setting LIBOGC to ${LIBOGC}")

    # Add paths it to toolchain
    include_directories(${LIBOGC_INCLUDE_DIR})
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
    add_custom_target(${target}_dol ALL
                      ${ELF2DOL} $<TARGET_FILE:${target}> ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol
                      DEPENDS ${CMAKE_CURRENT_BINARY_DIR}/${target}.elf
                      COMMENT "Generating .dol files from compiled .elf"
                      VERBATIM)
    set_directory_properties(PROPERTIES ADDITIONAL_MAKE_CLEAN_FILES ${CMAKE_CURRENT_BINARY_DIR}/${target}.dol)
endfunction()

# Apply add_dol_target to a multi-target
function(add_dol_targets target)
    if(GCN)
        add_dol_target(${target}_gcn)
    endif()
    if(WII)
        add_dol_target(${target}_wii)
    endif()
endfunction()

# Add library to all available targets
function(add_default_library target libname)
    if(GCN)
        target_link_libraries(${target}_gcn "${LIBOGC_LIBRARY_DIR_GCN}/lib${libname}.a")
    endif()
    if(WII)
        target_link_libraries(${target}_wii "${LIBOGC_LIBRARY_DIR_WII}/lib${libname}.a")
    endif()
endfunction()

# Create both GCN and WII version
function(add_multi_target target sources)
    if(GCN)
        add_executable(${target}_gcn ${sources})
        target_compile_options(${target}_gcn PUBLIC -DGEKKO -mogc)
        target_link_libraries(${target}_gcn "-mogc")
    endif()
    if(WII)
        add_executable(${target}_wii ${sources})
        target_compile_options(${target}_wii PUBLIC -DWII -mrvl)
        target_link_libraries(${target}_wii "-mrvl")
    endif()
endfunction()