package chunker

type Chunker struct {
	ChunkSize     int
	Overlap       int
	Breakpoint    []string
	RemoveNewline bool
}

func NewChunker(chunkSize, overlap int, breakpoint []string, removeNewline bool) *Chunker {
	return &Chunker{
		ChunkSize:     chunkSize,
		Overlap:       overlap,
		Breakpoint:    breakpoint,
		RemoveNewline: removeNewline,
	}
}

func (c *Chunker) Chunk(data []byte) [][]byte {
	var chunks [][]byte
	var chunk []byte
	var lastBreakpoint int
	for i, b := range data {
		if i-lastBreakpoint >= c.ChunkSize {
			if c.RemoveNewline {
				chunk = cleanChunk(chunk)
			}
			chunks = append(chunks, chunk)
			lastBreakpoint = i
			chunk = nil
		}
		chunk = append(chunk, b)
		if i-lastBreakpoint >= c.ChunkSize-c.Overlap {
			for _, bp := range c.Breakpoint {
				if len(chunk) >= len(bp) {
					if string(chunk[len(chunk)-len(bp):]) == bp {
						if c.RemoveNewline {
							chunk = cleanChunk(chunk)
						}
						chunks = append(chunks, chunk)
						lastBreakpoint = i
						chunk = nil
						break
					}
				}
			}
		}
	}
	if len(chunk) > 0 {
		if c.RemoveNewline {
			chunk = cleanChunk(chunk)
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func cleanChunk(chunk []byte) []byte {
	cleanChunk := make([]byte, 0)
	for i := range chunk {
		if chunk[i] == '\n' {
			continue
		}
		cleanChunk = append(cleanChunk, chunk[i])
	}
	return cleanChunk
}
