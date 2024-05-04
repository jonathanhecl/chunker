package chunker

import (
	"testing"
)

var exampleText = `
Argentina,[a] officially the Argentine Republic,[b] is a country in the southern half of 
South America. Argentina covers an area of 2,780,400 km2 (1,073,500 sq mi),[B] making it 
the second-largest country in South America after Brazil, the fourth-largest country in 
the Americas, and the eighth-largest country in the world. It shares the bulk of the 
Southern Cone with Chile to the west, and is also bordered by Bolivia and Paraguay to 
the north, Brazil to the northeast, Uruguay and the South Atlantic Ocean to the east, 
and the Drake Passage to the south. Argentina is a federal state subdivided into 
twenty-three provinces, and one autonomous city, which is the federal capital and 
largest city of the nation, Buenos Aires. The provinces and the capital have their 
own constitutions, but exist under a federal system. Argentina claims sovereignty 
over the Falkland Islands, South Georgia and the South Sandwich Islands, the Southern 
Patagonian Ice Field, and a part of Antarctica.`

func TestChunker_Chunk(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		chunker    *Chunker
		name       string
		args       args
		wantChunks int
		maxSize    int
	}{
		{
			chunker: NewChunker(40, 10, DefaultSeparators, true),
			name:    "Test demo",
			args: args{
				data: "This is a test string. It is used to test the chunker. It is a very simple chunker.",
			},
			wantChunks: 3,
			maxSize:    40,
		},
		{
			chunker: NewChunker(150, 30, DefaultSeparators, true),
			name:    "Example with wikipedia text",
			args: args{
				data: exampleText,
			},
			wantChunks: 9,
			maxSize:    150,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.chunker.Chunk(tt.args.data)

			for _, chunk := range got {
				// fmt.Println("Chunk `"+chunk+"` [ Length", len(chunk), "]")

				if len(chunk) > tt.maxSize {
					t.Errorf("Chunker.Chunk() = %v, want %v", len(chunk), tt.maxSize)
				}
			}

			if len(got) != tt.wantChunks {
				t.Errorf("Chunker.Chunk() = %v, want %v", len(got), tt.wantChunks)
			}
		})
	}
}

func Test_findLastSeparator(t *testing.T) {
	type args struct {
		chunk      string
		separators []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test find last separator",
			args: args{
				chunk:      "Testing the logic of findLastSeparator",
				separators: DefaultSeparators,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := findLastSeparator(tt.args.chunk, tt.args.separators, 0); got != tt.want {
				t.Errorf("findLastSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkChunk_Example1KB(b *testing.B) {
	characters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "\n"}

	size := int64(1024)
	content := make([]byte, size)
	for i := 0; i < len(content); i++ {
		content[i] = characters[i%len(characters)][0]
	}

	chunker := NewChunker(256, 32, DefaultSeparators, true)
	for i := 0; i < b.N; i++ {
		chunker.Chunk(string(content))
	}
}

func BenchmarkChunk_Example1MB(b *testing.B) {
	characters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "\n"}

	size := int64(1024 * 1024)
	content := make([]byte, size)
	for i := 0; i < len(content); i++ {
		content[i] = characters[i%len(characters)][0]
	}

	chunker := NewChunker(512, 64, DefaultSeparators, true)
	for i := 0; i < b.N; i++ {
		chunker.Chunk(string(content))
	}
}

func BenchmarkChunk_Example5MB(b *testing.B) {
	characters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "\n"}

	size := int64(5 * 1024 * 1024)
	content := make([]byte, size)
	for i := 0; i < len(content); i++ {
		content[i] = characters[i%len(characters)][0]
	}

	chunker := NewChunker(512, 64, DefaultSeparators, true)
	for i := 0; i < b.N; i++ {
		chunker.Chunk(string(content))
	}
}
