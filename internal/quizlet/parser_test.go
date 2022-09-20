package quizlet

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParse(t *testing.T) {
	tests := []struct {
		text string
		word *Word
		skip bool
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
			skip: false,
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
			skip: false,
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
			skip: false,
		},
		{
			text: `★Useful Words and Phrases`,
			word: nil,
			skip: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			word, skip, err := Parse(tt.text)
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(tt.skip, skip); diff != "" {
				t.Fatalf(diff)
			}
			if !tt.skip {
				if diff := cmp.Diff(tt.word, word); diff != "" {
					t.Fatalf(diff)
				}
			}
		})
	}
}
