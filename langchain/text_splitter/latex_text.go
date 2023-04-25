package text_splitter

func LatexTextSplitter(text string, maxChunkSize int, lengthFunc LengthFunc) []string {
	separators := []string{
		"\n\\chapter{",
		"\n\\section{",
		"\n\\subsection{",
		"\n\\subsubsection{",
		"\n\\begin{enumerate}",
		"\n\\begin{itemize}",
		"\n\\begin{description}",
		"\n\\begin{list}",
		"\n\\begin{quote}",
		"\n\\begin{quotation}",
		"\n\\begin{verse}",
		"\n\\begin{verbatim}",
		"\n\\begin{align}",
		"$$",
		"$",
		" ",
		"",
	}
	return RecursiveCharacterTextSplitter(text, maxChunkSize, separators, lengthFunc)
}
