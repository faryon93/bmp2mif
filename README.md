#bmp2mif

bmp2mif is a small tool, written in Go, which convertes a default 24bit bitmap file into a mif file. The mif file can be used to initialize a XILINX blockram.

## Install

Though bmp2mif is written in Go the advantages of its toolchain can be applied to this toll too. To install bmp2mif use the following shell command:

```
$: go get github.com/faryon93/bmp2mif/...
```

## Usage

The tool outputs the binary representation of the bitmap to stdout. Use stream redirection to copy it to a file.

```
$: bmp2mif test.bmp > test.mif
```