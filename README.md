# hexylon

Hexylon `[hˈɛksɪlən]` is [the geek god](https://en.wikipedia.org/wiki/Null_pointer) of human-friendly binary data. Write human readable and documented binary data. Cat them together with hexylon.

## Getting started

Create new 0x file:
```
0x7f 0x45 0x4c 0x46 // ELF magic number
0x02 0x01 0x01 0x00 // version and ABI stuff
00 00 00 00         // no 0x prefix necessary
00 00 00 00         // some more zeros
```

Run hexylon: `hexylon elf.0x out.bin`

Dump via hexdump: `hexdump -C out.bin`
```
00000000  7f 45 4c 46 02 01 01 00  00 00 00 00 00 00 00 00  |.ELF............|
00000010
```
