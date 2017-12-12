set(TOOLPATH ${CMAKE_SOURCE_DIR}/../tools) 
set(TOOLBIN ${TOOLPATH}/bin)

message(STATUS "Finding tools in ${TOOLBIN}")

# Check for objconv
find_program(OBJCONV objconv ${TOOLBIN})
if(NOT OBJCONV)
    message(WARNING "Could not find objconv")
endif()

# Check for bento
find_program(BENTO bento ${TOOLBIN})
if(NOT BENTO)
    message(WARNING "Could not find bento")
endif()

# Check for gcpacker
find_program(GCPACKER gcpacker ${TOOLBIN})
if(NOT GCPACKER)
    message(WARNING "Could not find gcpacker")
endif()

# Check for texconv
find_program(TEXCONV texconv ${TOOLBIN})
if(NOT TEXCONV)
    message(WARNING "Could not find texconv")
endif()

# Check for tevasm
find_program(TEVASM tevasm ${TOOLBIN})
if(NOT TEVASM)
    message(WARNING "Could not find tevasm")
endif()

include(FindPackageHandleStandardArgs)
find_package_handle_standard_args(TOOLS DEFAULT_MSG
                                  OBJCONV BENTO GCPACKER TEXCONV TEVASM)

mark_as_advanced(OBJCONV BENTO GCPACKER TEXCONV TEVASM TOOLBIN)

if(TOOLS_FOUND)
    message(STATUS "All tools found")
else()
    message(FATAL_ERROR "Could not find one or more required tools! Run \"tools/build.cmd\" or \"tools/build.sh\" depending on your OS to fix this")
endif()

# Convert one or more OBJ models to BMB
# The <output> variable contains the list of the converted models
# Usage:
#     convert_models(<output> <obj1> [<obj2> ..])
function(convert_models output)
    # Make directory
    set(MODEL_BMB_PATH ${CMAKE_CURRENT_BINARY_DIR}/models_bmb)
    set(MODELS "")
    file(MAKE_DIRECTORY ${MODEL_BMB_PATH})

    # Process all the given models
    foreach(__file ${ARGN})
        # Get output filename
        get_filename_component(__file_wd ${__file} NAME)
        string(REGEX REPLACE ".obj$" ".bmb" __BMB_FILE_NAME ${__file_wd})
        set(TARGET_FILE ${MODEL_BMB_PATH}/${__BMB_FILE_NAME})
        # Schedule objconv to run
        add_custom_command(OUTPUT ${TARGET_FILE}
            COMMAND ${OBJCONV} -in ${__file} -out ${TARGET_FILE}
            DEPENDS ${__file}
            WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
        set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES ${TARGET_FILE})
        # Append new file to output array
        list(APPEND MODELS ${TARGET_FILE})
    endforeach()
    set(${output} ${MODELS} PARENT_SCOPE)
endfunction()

# Convert one or more image files to BTB
# The <output> variable contains the list of the converted textures
# Usage:
#     convert_textures(<output> <fmt> <wrap> <filter> <tex1> [<tex2> ..])
function(convert_textures output fmt wrap filter)
    # Make directory
    set(TEXTURE_BTB_PATH ${CMAKE_CURRENT_BINARY_DIR}/textures_btb)
    set(TEXTURES "")
    file(MAKE_DIRECTORY ${TEXTURE_BTB_PATH})

    string(TOLOWER ${wrap} WRAP_L)
    string(TOLOWER ${filter} FILTER_L)

    # Process all the given textures
    foreach(__file ${ARGN})
        # Get output filename
        get_filename_component(__file_wd ${__file} NAME)
        string(REGEX REPLACE ".[^.]+$" ".btb" __BTB_FILE_NAME ${__file_wd})
        set(TARGET_FILE ${TEXTURE_BTB_PATH}/${__BTB_FILE_NAME})
        # Schedule objconv to run
        add_custom_command(OUTPUT ${TARGET_FILE}
            COMMAND ${TEXCONV} -in ${__file} -fmt ${fmt} -wrap ${WRAP_L} -filter ${FILTER_L} -out ${TARGET_FILE}
            DEPENDS ${__file}
            WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
        set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES ${TARGET_FILE})
        # Append new file to output array
        list(APPEND TEXTURES ${TARGET_FILE})
    endforeach()
    set(${output} ${TEXTURES} PARENT_SCOPE)
endfunction()

# Convert one or more TEV shaders to TDL
# The <output> variable contains the list of the converted shaders
# Usage:
#     convert_shaders(<output> <shader1> [<shader2> ..])
function(convert_shaders output)
    # Make directory
    set(SHADER_TDL_PATH ${CMAKE_CURRENT_BINARY_DIR}/shaders_tdl)
    set(SHADERS "")
    file(MAKE_DIRECTORY ${SHADER_TDL_PATH})

    # Process all the given shaders
    foreach(__file ${ARGN})
        # Get output filename
        get_filename_component(__file_wd ${__file} NAME)
        string(REGEX REPLACE ".[^.]+$" ".tdl" __TDL_FILE_NAME ${__file_wd})
        set(TARGET_FILE ${SHADER_TDL_PATH}/${__TDL_FILE_NAME})
        # Schedule tevasm to run
        add_custom_command(OUTPUT ${TARGET_FILE}
            COMMAND ${TEVASM} -in ${__file} -out ${TARGET_FILE}
            DEPENDS ${__file}
            WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
        set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES ${TARGET_FILE})
        # Append new file to output array
        list(APPEND SHADERS ${TARGET_FILE})
    endforeach()
    set(${output} ${SHADERS} PARENT_SCOPE)
endfunction()


# Embed arbitrary files into the final binaries
# <target> *must* be a multi-target
# Usage:
#     embed_resources(<target> <res1> [<res2> ..])
function(embed_resources target)
    # Make directories
    set(RES_OBJ_PATH ${CMAKE_CURRENT_BINARY_DIR}/resources_asm)
    file(MAKE_DIRECTORY ${RES_OBJ_PATH})
    set(RES_HEADER_PATH ${CMAKE_CURRENT_BINARY_DIR}/resources_header)
    file(MAKE_DIRECTORY ${RES_HEADER_PATH})

    # Add as include path
    if(GCN)
        target_include_directories(${target}_gcn PUBLIC ${RES_HEADER_PATH})
    endif()
    if(WII)
        target_include_directories(${target}_wii PUBLIC ${RES_HEADER_PATH})
    endif()

    # Process all the given resources
    foreach(__file ${ARGN})

        # Get C-friendly name
        get_filename_component(__file_wd ${__file} NAME)
        string(REGEX REPLACE "^([0-9])" "_\\1" __BIN_FILE_NAME ${__file_wd}) # add '_' if the file name starts by a number
        string(REGEX REPLACE "[-./]" "_" __BIN_FILE_NAME ${__BIN_FILE_NAME})
       
        # Call bento
        add_custom_command(OUTPUT ${RES_OBJ_PATH}/${__BIN_FILE_NAME}.s
                           COMMAND ${BENTO} -in ${__file} -align 32 -name ${__BIN_FILE_NAME} -headerpath ${RES_HEADER_PATH} -objectpath ${RES_OBJ_PATH}
                           DEPENDS ${__file}
                           WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
        add_library(${__BIN_FILE_NAME} ${RES_OBJ_PATH}/${__BIN_FILE_NAME}.s)
        set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES
                                 ${RES_OBJ_PATH}/${__BIN_FILE_NAME}.s
                                 ${RES_HEADER_PATH}/${__BIN_FILE_NAME}.h)
 
        # Add asm file as ASM library to either or both targets
        if(GCN)
            target_link_libraries(${target}_gcn ${__BIN_FILE_NAME})
        endif()
        if(WII)
            target_link_libraries(${target}_wii ${__BIN_FILE_NAME})
        endif()
    endforeach()
endfunction()

# Create a resource pack with one or more files inside
# <prefix> is the assets folder, including trailing slash
# <target> is the resulting resource pack target
# Usage:
#     add_resource_pack(<target> <prefix> <type> <res1> [[<type2> <res2> ..])
# All supported types:
#     BIN - Binary, embed as it is
#   MODEL - Model, use convert_models(..)
function(add_resource_pack target prefix)
    set(_filelist "${CMAKE_CURRENT_BINARY_DIR}/${target}_res.txt")
    set(_fname ${target}.gcr)
    set(_resfile "${CMAKE_CURRENT_BINARY_DIR}/${_fname}")
    set(_filetype "BIN")
    set(_txtfmt "RGBA8")
    set(_txtwrap "CLAMP")
    set(_txtfilter "BILINEAR")
    set(_depends ${_filelist})

    # Create resource list
    file(WRITE "${_filelist}" "")

    # Process all the given resources
    foreach(_name ${ARGN})
        if(_name MATCHES "BIN|MODEL|TEXTURE|SHADER")
            set(_filetype "${_name}")
        elseif(_name MATCHES "I4|I8|IA4|IA8|RGB565|RGB5A3|RGBA8|A8|CI4|CI8|CI14|CMPR")
            set(_txtfmt "${_name}")
        elseif(_name MATCHES "CLAMP|REPEAT|MIRROR")
            set(_txtwrap "${_name}")
        elseif(_name MATCHES "NEAR|BILINEAR|TRILINEAR")
            set(_txtfilter "${_name}")
        else()
            # Check what type is currently active
            if(_filetype STREQUAL "BIN")
                # Include file as it is, fix relative path
                file(APPEND "${_filelist}" "${_name},${CMAKE_CURRENT_LIST_DIR}/${prefix}${_name}\n")
                list(APPEND _depends "${CMAKE_CURRENT_LIST_DIR}/${prefix}${_name}")
            elseif(_filetype STREQUAL "MODEL")
                # Call convert_models(..) and add target path
                convert_models(MODEL "${prefix}${_name}")
                file(APPEND "${_filelist}" "${_name},${MODEL}\n")
                list(APPEND _depends ${MODEL})
            elseif(_filetype STREQUAL "TEXTURE")
                # Call convert_textures(..) and add target path
                convert_textures(TEXTURE ${_txtfmt} ${_txtwrap} ${_txtfilter} "${prefix}${_name}")
                file(APPEND "${_filelist}" "${_name},${TEXTURE}\n")
                list(APPEND _depends ${TEXTURE})
            elseif(_filetype STREQUAL "SHADER")
                # Call convert_shaders(..) and add target path
                convert_shaders(SHADER "${prefix}${_name}")
                file(APPEND "${_filelist}" "${_name},${SHADER}\n")
                list(APPEND _depends ${SHADER})
            endif()
        endif()
    endforeach()

    # Create resource pack target
    add_custom_command(OUTPUT ${_resfile}
                       COMMAND ${GCPACKER} -verbose -list ${_filelist} -out ${_resfile}
                       DEPENDS ${_depends}
                       COMMENT "Generating resource pack ${_fname}"
                       VERBATIM)
    add_custom_target(${target} DEPENDS ${_resfile})
    set_target_properties(${target} PROPERTIES OUTPUT_NAME ${_resfile})
    set_property(DIRECTORY APPEND PROPERTY ADDITIONAL_MAKE_CLEAN_FILES
                             ${_resfile})
endfunction()