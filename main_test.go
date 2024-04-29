package chunker

import (
	"fmt"
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
Patagonian Ice Field, and a part of Antarctica.

The earliest recorded human presence in modern-day Argentina dates back to the Paleolithic
period.[c] The country has its roots in Spanish colonization of the region during the 16th
century. Argentina rose as the successor state of the Viceroyalty of the Río de la Plata,[d]
a Spanish overseas viceroyalty founded in 1776. The declaration and fight for independence
(from Spain) was in 1816. The country thereafter enjoyed relative peace and stability, with
several waves of European.
`

func TestChunker_Chunk(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name       string
		args       args
		wantChunks int
		maxSize    int
	}{
		{
			name: "Test example",
			args: args{
				data: exampleText,
			},
			wantChunks: 9,
			maxSize:    150,
		},
	}
	c := NewChunker(150, 30, DefaultSeparators, true)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := c.Chunk(tt.args.data)

			for i, chunk := range got {
				fmt.Println("Chunk ", i+1, " `"+chunk+"` [ Length", len(chunk), "]")

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
		{
			name: "Test find last separator",
			args: args{
				chunk: `the Southern 
Patagonian Ice Field, and a part of Antarctica.

The earliest recorded human presence in modern-day Argentina dates back to the Paleoli`,
				separators: DefaultSeparators,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := findLastSeparator(tt.args.chunk, tt.args.separators); got != tt.want {
				t.Errorf("findLastSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
