package main

var _ Document = &TextDocument{}

type TextDocument struct {
	Content string
}

func (t *TextDocument) GetFormat() string {
	return "txt"
}

func (t *TextDocument) Clone() Document {
	newDoc := *t
	return &newDoc
}
