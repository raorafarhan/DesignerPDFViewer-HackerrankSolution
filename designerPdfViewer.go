func alphabet(word string) int {
	var wordIndex int
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	for i, v := range alphabet {
		if v == word {
			return i
		}
	}
	return wordIndex
}

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	var max int32

	for _, wordvalue := range word {
		wordIndex := alphabet(string(wordvalue))
		if max < h[int32(wordIndex)] {
			max = h[int32(wordIndex)]
		}

	}
	return max * int32(len(word))

}