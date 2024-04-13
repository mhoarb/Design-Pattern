package main

import "fmt"

func main() {
	originalDoc, err := CreateDocumant("txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Original Document Format:", originalDoc.GetFormat())

	// Clone the text document
	clonedDoc := originalDoc.Clone()

	// Modify the cloned document
	clonedDoc.(*TextDocument).Content = "This is the modified content."

	fmt.Println("Original Document Content:", originalDoc.(*TextDocument).Content) // Unchanged
	fmt.Println("Cloned Document Content:", clonedDoc.(*TextDocument).Content)     // Modified

	// Create a new PDF document
	pdfDoc, err := CreateDocumant("pdf")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("PDF Document Format:", pdfDoc.GetFormat())

	// Modify the original PDF document (demonstrates deep copy)
	pdfDoc.(*PDFDocument).Pages = append(pdfDoc.(*PDFDocument).Pages, "Page 2")

	// Print the number of pages in the cloned document (should still be 1)

}
