Super hacky solution but put this in main.cpp

```cpp
const auto dsize = 0x10000;
void* displayList = memalign(32, dsize);
memset(displayList, 0, dsize);
DCInvalidateRange(displayList, dsize);

// Build display list
GX_BeginDispList(displayList, dsize);

/* PUT SHADER CODE HERE */

auto size = GX_EndDispList();

printf("Size: %d\nData: ", size);
auto data = static_cast<unsigned char*>(displayList);
for (int i = 0; i < size; ++i) {
	printf("%02x ", data[i]);
}
printf("\n");
```

You might need to add these includes:
```cpp
#include <malloc.h>
#include <string.h>
```