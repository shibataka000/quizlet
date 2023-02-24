package quizlet

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		text string
		word Word
	}{
		{
			text: `significantly(adv)   大いに、著しく、かなり/ in a statistically significant way; hugely, notably, considerably, remarkably e.g. increase significantly, change significantly`,
			word: Word{
				English:     "significantly",
				Kind:        "adv",
				Japanese:    "大いに、著しく、かなり",
				Description: "in a statistically significant way; hugely, notably, considerably, remarkably",
				Example:     "increase significantly, change significantly",
			},
		},
		{
			text: `municipal(n) /mjunísəpəl/    地方自治体の/ of or pertaining to a city, town, etc., or its local government  e.g. municipal elections, municipal policy`,
			word: Word{
				English:     "municipal",
				Kind:        "n",
				Japanese:    "地方自治体の",
				Description: "of or pertaining to a city, town, etc., or its local government",
				Example:     "municipal elections, municipal policy",
			},
		},
		{
			text: `  ※would rather … : …する方がよい、むしろ…したい`,
			word: Word{
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
			require := require.New(t)
			word, err := Parse(tt.text)
			require.NoError(err)
			require.Equal(tt.word, word)
		})
	}
}

func TestSkip(t *testing.T) {
	tests := []struct {
		text string
		skip bool
	}{
		{
			text: `significantly(adv)   大いに、著しく、かなり/ in a statistically significant way; hugely, notably, considerably, remarkably e.g. increase significantly, change significantly`,
			skip: false,
		},
		{
			text: `municipal(n) /mjunísəpəl/    地方自治体の/ of or pertaining to a city, town, etc., or its local government  e.g. municipal elections, municipal policy`,
			skip: false,
		},
		{
			text: `  ※would rather … : …する方がよい、むしろ…したい`,
			skip: false,
		},
		{
			text: `★Useful Words and Phrases`,
			skip: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.text, func(t *testing.T) {
			require := require.New(t)
			skip := ShouldSkip(tt.text)
			require.Equal(tt.skip, skip)
		})
	}
}
