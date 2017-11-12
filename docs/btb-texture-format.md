# BTB (Binary Texture Blob) format

BTB is a GCR-friendly alternative to [TPL](https://devkitpro.org/wiki/libogc/GX#Textures). A BTB file contains a single texture data and its metadata (size, color format, mipmap settings) but using GCR as container instead of having its own.

The texture data is swizzled and tiled according to the color format in order to be uploaded to Gamecube with zero processing by the game code.

### Limits

| Field | Limit |
|-------|-------|
| Max texture size | 65536 x 65536 |
| Max color depth | 8 bpc (using RGBA8) |

## Color formats

### RGBA8

Blocks are 4x4 pixels

Each block is defines as follow:

```
LEGEND:
Rx (red) Gx (green) Bx (blue) Ax (alpha)
'x' is the xth pixel in the block, assuming each pixel is defined by X+(Y*4)

ENCODING:
A1 R1 A2 R2 .. A15 R15 A16 R16
G1 B1 G2 B2 .. G15 B15 G16 B16
```

## File format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | Header | BTB header | 14 |
| 0x20 | Texture data | texture data | ? |
| ? | Palette data | palette data (optional) | ? |

## Header format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | Texture width | uint16 | 2 |
| 0x2 | Texture height | uint16 | 2 |
| 0x4 | Color format | uint8 | 1 |
| 0x5 | Mipmaps | byte | 1 |
| 0x6 | Wrap + Filters | byte | 1 |
| 0x8 | Texture data offset | uint32 | 4 |
| 0xA | Texture palette offset (optional) | uint32 | 4 |

### Color format

| ID | Format name |
|----|-------------|
| 0x0 | I4 |
| 0x1 | I8 |
| 0x2 | IA4 |
| 0x3 | IA8 |
| 0x4 | RGB565 |
| 0x5 | RGB5A3 |
| 0x6 | RGBA8 |
| 0x7 | A8 |
| 0x8 | CI4 |
| 0x9 | CI8 |
| 0xA | CI14 |
| 0xE | CMPR (DXT1) |

### Wrap + Filters

| Bit(s) | Field |
|------|-------|
| 0x0 - 0x2 | S Wrap |
| 0x3 - 0x5 | T Wrap |
| 0x6 | Texture filter |
| 0x7 | Mipmap filter |

Valid values for wrapping modes:
| Value | Wrap mode |
|-------|-----------|
|0x0|Clamp|
|0x1|Repeat|
|0x2|Mirror|

Valid values for filters:
| Value | Wrap mode |
|-------|-----------|
|0x0|Near|
|0x1|Linear|

### Mipmaps

| Part | Field | Values |
|------|-------|-------|
| Upper 4 bits | Minimum LOD level | 0 to 10 (inclusive) |
| Lower 4 bits | Maximum LOD level | 0 to 10 (inclusive) |

Setting mipmaps at 0 will disable mipmapping

## References

- [Color formats](http://wiki.tockdom.com/wiki/Image_Formats)