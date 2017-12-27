#pragma once
#include "../pchheader.h"

#include "Atlas.h"
#include "Resource.h"

/*! \brief Atlas header
 */
struct AtlasResourceHeader {
	unsigned int entryCount; /*< Entry count */
};

/*! \brief Atlas entry
 */
struct AtlasEntry {
	unsigned int spriteName; /*< Path to sprite */
	unsigned short startX;   /*< X coordinate of the top left corner */
	unsigned short startY;   /*< Y coordinate of the top left corner */
	unsigned short sizeX;    /* Width */
	unsigned short sizeY;    /* Height */
};

/*!
 * \brief Atlas texture resource loader
 * Loads an atlas from a file in memory
 */
class AtlasResource : public Resource {
public:
	/*!
	 * \brief Loads a atlas in memory
	 *
	 * \param base Pointer to atlas file data
	 * \param size Size of the atlas file
	 */
	AtlasResource(void* base, unsigned int size) : Resource(base, size) {}

	/*!
	 * \brief Load and get the atlas
	 *
	 * \return Pointer to the loaded atlas
	 */
	std::shared_ptr<Atlas> Load();

	void Initialize() override;

private:
	AtlasResourceHeader* header = nullptr;
	bool loaded = false;
	std::shared_ptr<Atlas> internal;
};
