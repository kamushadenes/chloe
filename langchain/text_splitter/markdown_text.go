package text_splitter

func MarkdownTextSplitter(text string, maxChunkSize int, lengthFunc LengthFunc) []string {
	separators := []string{
		"\n## ",
		"\n### ",
		"\n#### ",
		"\n##### ",
		"\n###### ",
		"```\n\n",
		"\n\n***\n\n",
		"\n\n---\n\n",
		"\n\n___\n\n",
		"\n\n",
		"\n",
		" ",
		"",
	}
	return RecursiveCharacterTextSplitter(text, maxChunkSize, separators, lengthFunc)
}
