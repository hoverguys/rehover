# BMB model format

A BMB file is a model composed by a collection of vertex positions, vertex normals, texture coordinates (UV) and faces. BMB currently supports only one mesh per file.

Since the target platform runs on PowerPC, assume all values are big-endian.

Just like OBJ (original format), position, normals and coordinates are uncoupled and can be reused more than once.

## File format

| Name | Data type | Length (bytes) |
|------|-----------|----------------|
| Header | BMB header | 8 |
| Vertex positions | vertex position array | 12 * # of vertex positions |
| Vertex normals | vertex normal array | 12 * # of vertex normals |
| Texture coordinates | texture coordinate array | 8 * # of texture coords |
| Faces | face data array | 18 * # of faces

## Header format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | # of vertex positions | uint16 | 2 |
| 0x2 | # of vertex normals | uint16 | 2 |
| 0x4 | # of texture coordinates | uint16 | 2 |
| 0x6 | # of faces | uint16 | 2 |

## Vertex position/normal format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | X | float32 | 4 |
| 0x4 | Y | float32 | 4 |
| 0x8 | Z | float32 | 4 |

## Texture coordinate format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | U | float32 | 4 |
| 0x4 | V | float32 | 4 |

## Face format

Each face is made of 3 vertices forming a triangle:

### Face data

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | Vertex 1 | vertex data (see below) | 6 |
| 0x6 | Vertex 2 | vertex data (see below) | 6 |
| 0xC | Vertex 3 | vertex data (see below) | 6 |

### Vertex data

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | Vertex position index | uint16 | 2 |
| 0x2 | Texture coordinate index | uint16 | 2 |
| 0x4 | Vertex normal index | uint16 | 2 |