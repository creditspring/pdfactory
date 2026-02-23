package main

import (
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func GeneratePDF(html string, dpi uint) ([]byte, error) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	pdfg.Dpi.Set(dpi)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	page := wkhtmltopdf.NewPageReader(strings.NewReader(html))

	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}
