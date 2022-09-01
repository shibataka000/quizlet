package quizlet

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		text string
		word *Word
	}{
		{
			text: `significantly(adv)   大いに、著しく、かなり/ in a statistically significant way; hugely, notably, considerably, remarkably e.g. increase significantly, change significantly`,
			word: &Word{
				English:     "significantly",
				Kind:        "adv",
				Japanese:    "大いに、著しく、かなり",
				Description: "in a statistically significant way; hugely, notably, considerably, remarkably",
				Example:     "increase significantly, change significantly",
			},
		},
		{
			text: `municipal(n) /mjunísəpəl/    地方自治体の/ of or pertaining to a city, town, etc., or its local government  e.g. municipal elections, municipal policy`,
			word: &Word{
				English:     "municipal",
				Kind:        "n",
				Japanese:    "地方自治体の",
				Description: "of or pertaining to a city, town, etc., or its local government",
				Example:     "municipal elections, municipal policy",
			},
		},
		{
			text: `  ※would rather … : …する方がよい、むしろ…したい`,
			word: &Word{
				English:     "would rather …",
				Kind:        "expression",
				Japanese:    "…する方がよい、むしろ…したい",
				Description: "",
				Example:     "",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			word, err := Parse(tt.text)
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(word, tt.word) {
				t.Fatalf("Expected is %v but actual is %v", tt.word, word)
			}
		})
	}
}
