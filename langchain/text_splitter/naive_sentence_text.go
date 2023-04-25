package text_splitter

func NaiveSentenceTextSplitter(text string, maxChunkSize int, lengthFunc LengthFunc) []string {
	separators := []string{".", "!", "?"}
	return RecursiveCharacterTextSplitter(text, maxChunkSize, separators, lengthFunc)
}
