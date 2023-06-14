package main

import "testing"

func Test_findAnagram(t *testing.T) {
	input := []struct {
		in    []string
		expected map[string][]string
	}{
		{
			in:    []string{"пятак", "пятка" , "тяпка", "листок", "слиток" , "столик", "один", "дино", "два"},
			expected: map[string][]string{"листок": {"листок", "слиток", "столик"},
				"один": {"дино", "один"}, "пятак": {"пятак", "пятка", "тяпка"}},
		},
		{
			in:    []string{},
			expected: map[string][]string{},
		},
	}

	for _, d := range input {
		t.Run("find anagram", func(t *testing.T) {
			result := findAnagrams(d.in)

			if len(result) != len(d.expected) {
				t.Fatal("No similar len")
			}

			for key, a1 := range d.expected {
				a2, ok := result[key]
				if !ok {
					t.Fatalf("Expected %s, got %s", d.expected, result)
				}

				if len(a1) != len(a2) {
					t.Fatalf("Expected %s, got %s", d.expected, result)
				}

				for i := 0; i < len(a1) && i < len(a2); i++ {
					if a1[i] != a2[i] {
						t.Fatalf("Expected %s, got %s", d.expected, result)
					}
				}
			}
		})
	}
}