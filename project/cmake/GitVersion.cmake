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