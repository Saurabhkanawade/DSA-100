package main

type DocumentReader interface {
	Read() string
}
type DocumentWriter interface {
	Write() string
}
type DocumentPrinter interface {
	Print() string
}

type TextDocumentReader struct {
}
type TextDocument struct {
}

func (t *TextDocumentReader) Write() string {
	return "printing the document"
}
func (t *TextDocument) Print() string {
	return "printing the document"
}

func main() {
	td := TextDocument{}
	tdR := TextDocumentReader{}
	td.Print()
	tdR.Write()
}
