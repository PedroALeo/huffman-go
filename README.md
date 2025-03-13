# huffman-go

This repository provides a Go implementation of Huffman coding, an algorithm for data compression. The package includes functions for creating a Huffman code, compressing data with it, and decompressing the data back to its original form.

## Features

- **Create Huffman Code**: Generates a Huffman tree and creates a corresponding code map.
- **Compress Data**: Compresses a string using the generated Huffman code.
- **Decompress Data**: Decodes the compressed data back into its original string.
- **Print Huffman Code**: Displays the Huffman codes for each character in the input string.
- **Compression/Decompression with Map Code**: Compress and decompress strings using pre-generated Huffman code maps.
- **Observation**: The CompressWithMapCode function returns not only the compressed byte slice but also the original length of the string (in bits). This is important for decompression because the length of the original string is required to accurately decode the compressed data. This means that the same Huffman code map can be used to compress and decompress different strings, as long as the strings consist of the same set of characters. The original length ensures that the decompression process knows when to stop.

## Installation

`go get github.com/PedroALeo/huffman-go`
