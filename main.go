package chunker

import (
	"strings"
)

type Chunker struct {
	ChunkSize            int
	Overlap              int
	Separators           []string
	OutputWithoutNewline bool
}

var (
	DefaultSeparators = []string{"\n\n", " ", "\n"}
)

func NewChunker(chunkSize, overlap int, separators []string, outputWithoutNewline bool) *Chunker {
	return &Chunker{
		ChunkSize:            chunkSize,
		Overlap:              overlap,
		Separators:           separators,
		OutputWithoutNewline: outputWithoutNewline,
	}
}

func (c *Chunker) Chunk(data string) []string {
	var chunks []string

	var i int = 0
	for {
		if i == 0 {
			if len(data) < c.ChunkSize {
				possibleChunk := data
				if c.OutputWithoutNewline {
					possibleChunk = removeNewlineInChunk(possibleChunk)
				}

				chunks = append(chunks, possibleChunk)
				break
			}

			possibleChunk := data[:c.ChunkSize]
			if c.OutputWithoutNewline {
				possibleChunk = removeNewlineInChunk(possibleChunk)
			}
			lastSeparator, ss := findLastSeparator(possibleChunk, c.Separators)
			chunks = append(chunks, possibleChunk[:lastSeparator])
			i = lastSeparator + ss - c.Overlap
		} else {
			if len(data)-i < c.ChunkSize {
				possibleChunk := data[i:]
				if c.OutputWithoutNewline {
					possibleChunk = removeNewlineInChunk(possibleChunk)
				}
				firstSeparator := findFirstSeparator(possibleChunk, c.Separators)
				if firstSeparator > c.Overlap {
					firstSeparator = 0
				}
				chunks = append(chunks, possibleChunk[firstSeparator:])
				break
			}

			possibleChunk := data[i : i+c.ChunkSize]
			if c.OutputWithoutNewline {
				possibleChunk = removeNewlineInChunk(possibleChunk)
			}
			firstSeparator := findFirstSeparator(possibleChunk, c.Separators)
			if firstSeparator > c.Overlap {
				firstSeparator = 0
			}
			lastSeparator, ss := findLastSeparator(possibleChunk, c.Separators)
			if lastSeparator < firstSeparator {
				lastSeparator = len(possibleChunk)
			}

			chunks = append(chunks, possibleChunk[firstSeparator:lastSeparator])
			i += lastSeparator + ss - c.Overlap
		}
	}

	return chunks
}

func findFirstSeparator(chunk string, separators []string) (offset int) {
	for _, sp := range separators {
		if len(chunk) >= len(sp) {
			firstPos := strings.Index(chunk, sp)
			if firstPos != -1 {
				return firstPos + len(sp)
			}
		}
	}
	return 0
}

func findLastSeparator(chunk string, separators []string) (offset, separatorSize int) {
	for _, sp := range inverted(separators) {
		if len(chunk) >= len(sp) {
			lastPos := strings.LastIndex(chunk, sp)
			if lastPos != -1 {
				return lastPos, len(sp)
			}
		}
	}
	return 0, 0
}

func removeNewlineInChunk(chunk string) string {
	// remove /n from the beginning of the chunk
	if chunk[0] == '\n' {
		chunk = chunk[1:]
	}

	// remove /n from the end of the chunk
	if chunk[len(chunk)-1] == '\n' {
		chunk = chunk[:len(chunk)-1]
	}

	// remove /n in the middle of the chunk, replace with space if it is not followed by a space
	chunk = strings.ReplaceAll(chunk, "\n ", " ")
	chunk = strings.ReplaceAll(chunk, " \n", " ")
	chunk = strings.ReplaceAll(chunk, "\n", " ")

	return chunk
}

func inverted(s []string) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[len(s)-1-i] = v
	}
	return r
}
