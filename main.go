package chunker

import (
	"fmt"
	"sort"
	"strings"
)

type Chunker struct {
	ChunkSize            int
	Overlap              int
	Separators           []string
	OutputWithoutNewline bool
	Debug                bool
	// internal
	chunks []string
}

var (
	DefaultSeparators = []string{"\n\n", " ", "\n"}
)

func NewChunker(chunkSize, overlap int, separators []string, outputWithoutNewline, debug bool) *Chunker {
	if chunkSize <= 0 {
		chunkSize = 150
	}
	if overlap <= 0 {
		overlap = 30
	}
	if overlap >= chunkSize {
		overlap = int(chunkSize / 4)
	}
	if len(separators) == 0 {
		separators = DefaultSeparators
	}

	return &Chunker{
		ChunkSize:            chunkSize,
		Overlap:              overlap,
		Separators:           separators,
		OutputWithoutNewline: outputWithoutNewline,
		Debug:                debug,
	}
}

func (c *Chunker) Chunk(data string) []string {
	c.ClearChunks()

	var i int = 0
	for {
		if c.Debug {
			fmt.Println("i: ", i, "len(data): ", len(data))
			fmt.Println("chunks: ", len(c.chunks))
		}
		if i == 0 {
			if len(data) < c.ChunkSize {
				possibleChunk := data

				c.addChunk(possibleChunk)
				break
			}

			possibleChunk := data[:c.ChunkSize]
			lastSeparator, ss := findLastSeparator(possibleChunk, c.Separators, 0)

			possibleChunk = possibleChunk[:lastSeparator]
			c.addChunk(possibleChunk)

			i = lastSeparator + ss - c.Overlap
		} else {
			if len(data)-i < c.ChunkSize {
				possibleChunk := data[i:]
				firstSeparator := findFirstSeparator(possibleChunk, c.Separators)
				if firstSeparator >= c.Overlap {
					firstSeparator = 0
				}

				possibleChunk = possibleChunk[firstSeparator:]
				c.addChunk(possibleChunk)
				break
			}

			possibleChunk := data[i : i+c.ChunkSize]
			firstSeparator := findFirstSeparator(possibleChunk, c.Separators)
			if firstSeparator >= c.Overlap {
				firstSeparator = 0
			}
			lastSeparator, ss := findLastSeparator(possibleChunk, c.Separators, firstSeparator)

			possibleChunk = possibleChunk[firstSeparator:lastSeparator]
			c.addChunk(possibleChunk)

			if len(possibleChunk) > c.Overlap {
				i += lastSeparator + ss - c.Overlap
			} else {
				if lastSeparator == 0 && ss == 0 {
					i += c.Overlap
				} else {
					i += lastSeparator + ss
				}
			}
		}
	}

	return c.GetChunks()
}

func (c *Chunker) addChunk(chunk string) {
	if c.OutputWithoutNewline {
		chunk = removeNewlineInChunk(chunk)
	}

	chunk = strings.TrimSpace(chunk)
	if len(chunk) == 0 {
		return
	}

	c.chunks = append(c.chunks, chunk)
}

func (c *Chunker) ClearChunks() {
	c.chunks = make([]string, 0)
}

func (c *Chunker) GetChunkSize() int {
	return c.ChunkSize
}

func (c *Chunker) GetChunks() []string {
	return c.chunks
}

func findFirstSeparator(chunk string, separators []string) (offset int) {
	pos := [][]int{}
	for _, sp := range separators {
		if len(chunk) >= len(sp) {
			firstPos := strings.Index(chunk, sp)
			if firstPos != -1 {
				pos = append(pos, []int{firstPos, len(sp)})
				continue
			}
		}
	}
	if len(pos) == 0 {
		return 0
	}

	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0] < pos[j][0]
	})

	return pos[0][0] + pos[0][1]
}

func findLastSeparator(chunk string, separators []string, from int) (offset, separatorSize int) {
	pos := [][]int{}
	for _, sp := range separators {
		if len(chunk) >= len(sp) {
			lastPos := strings.LastIndex(chunk, sp)
			if lastPos != -1 && lastPos > from {
				pos = append(pos, []int{lastPos, len(sp)})
				continue
			}
		}
	}

	if len(pos) == 0 {
		return len(chunk), 0
	}

	sort.Slice(pos, func(i, j int) bool {
		return pos[i][0] > pos[j][0]
	})

	return pos[0][0], pos[0][1]
}

func removeNewlineInChunk(chunk string) string {
	if len(chunk) == 0 {
		return chunk
	}

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
