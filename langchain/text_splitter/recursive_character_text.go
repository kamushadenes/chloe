package text_splitter

func RecursiveCharacterTextSplitter(text string, maxChunkSize int, separators []string, lengthFunc LengthFunc) []string {
	if maxChunkSize <= 0 {
		panic("maxChunkSize must be greater than 0")
	}

	if lengthFunc == nil {
		lengthFunc = defaultLengthFunc
	}

	var splitter func(text string, maxChunkSize int, separators []string, separatorIndex int, lengthFunc LengthFunc) []string
	splitter = func(text string, maxChunkSize int, separators []string, separatorIndex int, lengthFunc LengthFunc) []string {
		if lengthFunc(text) <= maxChunkSize || separatorIndex >= len(separators) {
			return []string{text}
		}

		separator := separators[separatorIndex]
		chunks := CharacterTextSplitter(text, maxChunkSize, separator, lengthFunc)

		if len(chunks) <= 1 {
			return splitter(text, maxChunkSize, separators, separatorIndex+1, lengthFunc)
		}

		var result []string
		for _, chunk := range chunks {
			if lengthFunc(chunk) > maxChunkSize {
				subChunks := splitter(chunk, maxChunkSize, separators, separatorIndex+1, lengthFunc)
				result = append(result, subChunks...)
			} else {
				result = append(result, chunk)
			}
		}

		return result
	}

	return splitter(text, maxChunkSize, separators, 0, lengthFunc)
}
