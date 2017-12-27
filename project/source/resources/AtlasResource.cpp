#include "AtlasResource.h"

void AtlasResource::Initialize() {
	header = static_cast<AtlasResourceHeader*>(address);
	printf("%d entries in atlas\n", header->entryCount);

	internal = std::make_shared<Atlas>();

	// Load each entry into the atlas
	auto offset = address + sizeof(AtlasResourceHeader);
	for (unsigned int i = 0; i < header->entryCount; ++i) {
		auto entry = static_cast<AtlasEntry*>(offset + sizeof(AtlasEntry) * i);
		auto rect = Rect(entry->startX, entry->startY, entry->sizeX, entry->sizeY);
		internal->coordinates[entry->spriteName] = rect;
		printf("Entry %d has coords %f %f size %f %f\n", entry->spriteName, rect.start.x, rect.start.y, rect.size.x,
			   rect.size.y);
	}
}

std::shared_ptr<Atlas> AtlasResource::Load() {
	return internal;
}