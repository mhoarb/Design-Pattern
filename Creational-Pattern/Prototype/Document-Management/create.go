package main

import "fmt"

func CreateDocumant(format string) (Document, error) {
	switch format {
	case "txt":
		return &TextDocument{}, nil
	case "pdf":
		return &PDFDocument{}, nil
	case "xlsx":
		return &SpreadSheet{}, nil
	default:
		return nil, fmt.Errorf("Unsupported document format: %s\n", format)

	}
}
