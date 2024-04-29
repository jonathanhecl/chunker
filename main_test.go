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
The earliest recorded human presence in modern-day Argentina dates back to the 
Paleolithic period.[13] The Inca Empire expanded to the northwest of the country in 
Pre-Columbian times. The country has its roots in Spanish colonization of the region 
during the 16th century.[14] Argentina rose as the successor state of the Viceroyalty 
of the Río de la Plata,[15] a Spanish overseas viceroyalty founded in 1776. The 
declaration and fight for independence (1810–1818) was followed by an extended civil 
war that lasted until 1861, culminating in the country's reorganization as a federation. 
The country thereafter enjoyed relative peace and stability, with several waves of 
European immigration, mainly Italians and Spaniards, influencing its culture and 
demography.[16][17][18][19]
Following the death of President Juan Perón in 1974, his widow and vice president, 
Isabel Perón, ascended to the presidency, before being overthrown in 1976. The 
following military junta, which was supported by the United States, persecuted 
and murdered thousands of political critics, activists, and leftists in the 
Dirty War, a period of state terrorism and civil unrest that lasted until the 
election of Raúl Alfonsín as president in 1983.
Argentina is a regional power, and retains its historic status as a middle power 
in international affairs.[20][21][22] A major non-NATO ally of the United States,[23] 
Argentina is a developing country with the second-highest HDI (human development index) 
in Latin America after Chile.[24] It maintains the second-largest economy in South 
America, and is a member of G-15 and G20. Argentina is also a founding member of the 
United Nations, World Bank, World Trade Organization, Mercosur, Community of Latin 
American and Caribbean States and the Organization of Ibero-American States.
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
			wantChunks: 34,
		},
	}
	c := NewChunker(100, 20, []string{"\n\n", "\n", ", ", " "}, true)
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
