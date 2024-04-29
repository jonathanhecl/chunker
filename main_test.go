package chunker

import (
	"fmt"
	"testing"
)

var exampleText = []byte(`
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
Patagonian Ice Field, and a part of Antarctica.
`)

func TestChunker_Chunk(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name       string
		args       args
		wantChunks int
	}{
		{
			name: "Test example",
			args: args{
				data: exampleText,
			},
			wantChunks: 12,
		},
	}
	c := NewChunker(100, 20, DefaultSeparators, true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Chunk(tt.args.data)

			for i, chunk := range got {
				fmt.Println("Chunk ", i+1, " `"+string(chunk)+"`")
			}

			if len(got) != tt.wantChunks {
				t.Errorf("Chunker.Chunk() = %v, want %v", len(got), tt.wantChunks)
			}
		})
	}
}
