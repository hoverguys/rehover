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
