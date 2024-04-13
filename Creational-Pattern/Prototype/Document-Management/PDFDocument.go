package main

var _ Document = &PDFDocument{}

type PDFDocument struct {
	Pages []string
}

func (P *PDFDocument) GetFormat() string {
	return "pdf"
}

func (p *PDFDocument) Clone() Document {
	newDoc := PDFDocument{Pages: make([]string, len(p.Pages))}
	copy(newDoc.Pages, p.Pages)
	return &newDoc
}
