// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 52F592                  one-file-pdf/utest/[clean.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_Clean_ is the unit test for PDF.Clean()
func Test_PDF_Clean_(t *testing.T) {
	fmt.Println("Test PDF.Clean()")

	// calling Clean() multiple times on a non-initialized PDF:
	// (you should not do this normally, use NewPDF() to create a PDF)
	// - should not panic
	// - length of Errors() should be zero
	// - Errors() should be []error{}, not nil
	func() {
		var doc pdf.PDF // uninitialized PDF
		doc.Clean().Clean().Clean()
		//
		//        got                want
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// same as above for a PDF properly initialized with NewPDF()
	// (also, call Clean() without chaining)
	func() {
		doc := pdf.NewPDF("A4")
		doc.Clean()
		doc.Clean()
		doc.Clean()
		//        got                want
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

	// create a new PDF with an unknown page size, then call Clean()
	// first, Errors should have 1 error
	// after Clean(), Errors should be zero-length
	func() {
		doc := pdf.NewPDF("Parchment")
		//        got                want
		TEqual(t, len(doc.Errors()), 1)
		doc.Clean()
		//        got                want
		TEqual(t, len(doc.Errors()), 0)
		TEqual(t, doc.Errors(), []error{})
	}()

} //                                                             Test_PDF_Clean_

//end
