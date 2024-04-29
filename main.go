package chunker

type Chunker struct {
	ChunkSize            int
	Overlap              int
	Separators           []string
	OutputWithoutNewline bool
}

var (
	DefaultSeparators = []string{"\n\n", "\n", " "}
)

func NewChunker(chunkSize, overlap int, separators []string, outputWithoutNewline bool) *Chunker {
	return &Chunker{
		ChunkSize:            chunkSize,
		Overlap:              overlap,
		Separators:           separators,
		OutputWithoutNewline: outputWithoutNewline,
	}
}

func (c *Chunker) Chunk(data []byte) [][]byte {
	var chunks [][]byte
	var chunk []byte
	var lastBreakpoint int
	for i, b := range data {
		if i-lastBreakpoint >= c.ChunkSize {
			if c.OutputWithoutNewline {
				chunk = removeNewlineInChunk(chunk)
			}
			chunks = append(chunks, chunk)
			lastBreakpoint = i
			chunk = nil
		}
		chunk = append(chunk, b)
		if i-lastBreakpoint >= c.ChunkSize-c.Overlap {
			for _, sp := range c.Separators {
				if len(chunk) >= len(sp) {
					if string(chunk[len(chunk)-len(sp):]) == sp {
						if c.OutputWithoutNewline {
							chunk = removeNewlineInChunk(chunk)
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
		if c.OutputWithoutNewline {
			chunk = removeNewlineInChunk(chunk)
		}
		chunks = append(chunks, chunk)
	}
	return chunks
}

func removeNewlineInChunk(chunk []byte) []byte {
	cleanChunk := make([]byte, 0)
	for i := range chunk {
		if i == 0 && chunk[i] == '\n' { // remove leading newline
			continue
		}
		if i == len(chunk)-1 && chunk[i] == '\n' { // remove trailing newline
			continue
		}
		if chunk[i] == '\n' { // remove newline
			if i+1 < len(chunk) && chunk[i+1] != ' ' { // add space if next character is not space
				cleanChunk = append(cleanChunk, ' ')
			}
			continue
		}
		cleanChunk = append(cleanChunk, chunk[i])
	}
	return cleanChunk
}
