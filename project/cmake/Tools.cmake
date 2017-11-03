set(TOOLPATH ${CMAKE_SOURCE_DIR}/../tools) 
set(TOOLBIN ${TOOLPATH}/bin)

message(STATUS "Finding tools in ${TOOLBIN}")

# Check for objconv
find_program(OBJCONV objconv ${TOOLBIN})
if(NOT OBJCONV)
    message(WARNING "Could not find objconv")
endif()

# Check for objconv
find_program(BENTO bento ${TOOLBIN})
if(NOT BENTO)
    message(WARNING "Could not find bento")
endif()

include(FindPackageHandleStandardArgs)
find_package_handle_standard_args(TOOLS DEFAULT_MSG
                                  OBJCONV BENTO)

mark_as_advanced(OBJCONV BENTO TOOLBIN)

if(TOOLS_FOUND)
    message(STATUS "All tools found")
else()
    message(FATAL_ERROR "Could not find one or more required tools! Run \"tools/build.cmd\" or \"tools/build.sh\" depending on your OS to fix this")
endif()

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
                           COMMAND ${BENTO} -in ${__file} -name ${__BIN_FILE_NAME} -headerpath ${RES_HEADER_PATH} -objectpath ${RES_OBJ_PATH}
                           DEPENDS ${__file}
                           WORKING_DIRECTORY ${CMAKE_CURRENT_LIST_DIR})
        add_library(${__BIN_FILE_NAME} ${RES_OBJ_PATH}/${__BIN_FILE_NAME}.s)

        # Add asm file as ASM library to either or both targets
        if(GCN)
            target_link_libraries(${target}_gcn ${__BIN_FILE_NAME})
        endif()
        if(WII)
            target_link_libraries(${target}_wii ${__BIN_FILE_NAME})
        endif()
    endforeach()
endfunction()
