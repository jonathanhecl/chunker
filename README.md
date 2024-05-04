# Chunker [WIP]
A simple chunker for text data. It is useful for chunking text data into smaller parts for processing.

## Features
- Chunk text data into smaller parts
- Don't split words in the middle
- Limit the length of each chunk, and overlap between chunks
- Clean up the chunks by removing leading and trailing whitespaces
- Can remove new lines from the chunks

## Installation
```go get github.com/jonathanhecl/chunker```

## Usage
```go
c := chunker.NewChunker(40, 10, chunker.DefaultSeparators, true)
chunks := c.Chunk("This is a test string. It is used to test the chunker. It is a very simple chunker.")
```

## Result
```text
Chunk  1  `This is a test string. It is used to` [ Length 36 ]
Chunk  2  `used to test the chunker. It is a` [ Length 33 ]
Chunk  3  `It is a very simple chunker.` [ Length 28 ]
```

## Benchmark
```text
goos: windows
goarch: amd64
pkg: github.com/jonathanhecl/chunker
cpu: 13th Gen Intel(R) Core(TM) i7-13700K
BenchmarkChunk_Example1KB/input_size_1024(256/32)-24              457161              2607 ns/op
BenchmarkChunk_Example1MB/input_size_1048576(512/64)-24              438           2747814 ns/op
BenchmarkChunk_Example5MB/input_size_5242880(512/64)-24               75          13781380 ns/op
BenchmarkChunk_Example10MB/input_size_10485760(1024/128)-24           45          25868804 ns/op
```