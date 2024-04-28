package chunker

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChunker_Chunk(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "Test 1",
			args: args{
				data: []byte(`hello 
				world, this is a test.
				A test to see if the chunker works.`),
			},
			want: [][]byte{
				[]byte(`hello
				world, this is a test.`),
				[]byte(`world, this is a test.
				A test to see if the chunker works.`),
			},
		},
	}
	c := NewChunker(30, 5)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.Chunk(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				for i, chunk := range got {
					fmt.Println(i, string(chunk))
				}
				t.Errorf("Chunker.Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
