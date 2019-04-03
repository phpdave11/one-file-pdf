// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-03 10:27:56 85C87A            one-file-pdf/utest/[fill_circle.go]
// -----------------------------------------------------------------------------

package utest

import (
	"fmt"
	"testing"

	"github.com/balacode/one-file-pdf"
)

// Test_PDF_FillCircle_ is the unit test for
// PDF.FillCircle(x, y, radius float64, fill ...bool) *PDF
//
// Runs the test by drawing the flag of Japan using correct proportions
func Test_PDF_FillCircle_(t *testing.T) {
	fmt.Println("Test PDF.FillCircle()")
	var (
		doc    = pdf.NewPDF("30cm x 20cm")
		x, y   = 15.0, 10.0         // center of page
		radius = (20.0 * 3 / 5) / 2 // diameter = 3/5 of height
	)
	doc.SetCompression(false).
		SetUnits("cm").
		SetColor("#BC002D (close to #BE0032 CrimsonGlory)").
		FillCircle(x, y, radius)

	const want = `
	%PDF-1.4
	1 0 obj <</Type/Catalog/Pages 2 0 R>>
	endobj
	2 0 obj <</Type/Pages/Count 1/MediaBox[0 0 850 566]/Kids[3 0 R]>>
	endobj
	3 0 obj <</Type/Page/Parent 2 0 R/Contents 4 0 R>>
	endobj
	4 0 obj <</Length 267>> stream
	0.737 0.000 0.176 rg
	0.737 0.000 0.176 RG
	255.118 283.465 m
	255.118 377.396 331.265 453.543 425.197 453.543 c
	519.129 453.543 595.276 377.396 595.276 283.465 c
	595.276 189.533 519.129 113.386 425.197 113.386 c
	331.265 113.386 255.118 189.533 255.118 283.465 c
	b
	endstream
	endobj
	xref
	0 5
	0000000000 65535 f
	0000000010 00000 n
	0000000056 00000 n
	0000000130 00000 n
	0000000189 00000 n
	trailer
	<</Size 5/Root 1 0 R>>
	startxref
	507
	%%EOF
	`

	ComparePDF(t, doc.Bytes(), want)
} //                                                        Test_PDF_FillCircle_

//end
