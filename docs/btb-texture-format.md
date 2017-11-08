# BTB (Binary Texture Blob) format

BTB is a GCR-friendly alternative to [TPL](https://devkitpro.org/wiki/libogc/GX#Textures). A BTB file contains a single texture data and its metadata (size, color format, mipmap settings) but using GCR as container instead of having its own.

The texture data is swizzled and tiled according to the color format in order to be uploaded to Gamecube with zero processing by the game code.

### Limits

| Field | Limit |
|-------|-------|
| Max texture size | 65536 x 65536 |
| Max color depth | 8 bpc (using RGBA8) |

## File format

| Name | Data type | Length (bytes) |
|------|-----------|----------------|
| Header | BTB header | 8 |
| Texture data | texture data | ? |

## Header format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | Texture width | uint16 | 2 |
| 0x2 | Texture height | uint16 | 2 |
| 0x4 | Mipmaps | byte | 1 |
| 0x5 | *reserved* | ? | 3 |

### Mipmaps

| Part | Field | Values |
|------|-------|-------|
| Upper 4 bits | Minimum LOD level | 0 to 10 (inclusive) |
| Lower 4 bits | Maximum LOD level | 0 to 10 (inclusive) |

Setting mipmaps at 0 will disable mipmapping