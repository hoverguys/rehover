# GCR (GameCube Resource) format

GCR is a format for packing multiple files in a single archive 

Files are identified by a unique uint32 identifier (currently obtained by hashing the filename using FNV32)

File data is padded to the nearest 32 bit.

### Limits

| Field | Limit |
|-------|-------|
| # of files | x < 4294967296 |
| Single file size | 4 GiB (4294967296 bytes) | 
| Total GCR size | 4 GiB (4294967296 bytes) |

## File format

| Name | Data type | Length (bytes) |
|------|-----------|----------------|
| Header | GCR header | 4 + 12 * # of files |
| Files | file data array | ? |

## Header format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | # of files | uint32 | 4 |
| 0x4 | File entries | file entry array | 12 * # of files

### File entry format

| Offset | Name | Data type | Length (bytes) |
|--------|------|-----------|----------------|
| 0x0 | File identifier | uint32 | 4 |
| 0x4 | File offset | uint32 | 4 |
| 0x8 | File size | uint32 | 4 |