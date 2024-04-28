package chunker

type Chunker struct {
	ChunkSize int
	Overlap   int
}

func NewChunker(chunkSize, overlap int) *Chunker {
	return &Chunker{
		ChunkSize: chunkSize,
		Overlap:   overlap,
	}
}

func (c *Chunker) Chunk(data []byte) [][]byte {
	var chunks [][]byte
	for i := 0; i < len(data); i += c.ChunkSize - c.Overlap {
		end := i + c.ChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunks = append(chunks, data[i:end])
	}
	return chunks
}
