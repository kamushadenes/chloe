package text_splitter

import "strings"

func CharacterTextSplitter(text string, maxChunkSize int, separator string, lengthFunc LengthFunc) []string {
	if maxChunkSize <= 0 {
		panic("maxChunkSize must be greater than 0")
	}

	if lengthFunc == nil {
		lengthFunc = defaultLengthFunc
	}

	var chunks []string
	startIndex := 0
	endIndex := 0
	textLen := lengthFunc(text)

	for endIndex < textLen {
		if endIndex+maxChunkSize < textLen {
			endIndex += maxChunkSize
			splitIndex := strings.LastIndex(text[startIndex:endIndex], separator)

			if splitIndex != -1 {
				endIndex = startIndex + splitIndex + len(separator)
			} else {
				endIndex = startIndex + maxChunkSize
			}

			chunks = append(chunks, text[startIndex:endIndex])
			startIndex = endIndex
		} else {
			chunks = append(chunks, text[startIndex:])
			break
		}
	}

	return chunks
}
