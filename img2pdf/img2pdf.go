package img2pdf

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/jung-kurt/gofpdf"
)

// Generator is the struct of PDF generator.
type Generator struct {
	imageProcessingEvent func(id int, filename string)
}

// CreatePDF creates PDF file with the the name, containing image files given.
func (g Generator) CreatePDF(name string, imageFiles []string) error {
	// orientation: portrait, unit: pt
	pdf := gofpdf.New("P", "pt", "", "")

	// iterate over each image file
	for i, file := range imageFiles {
		width, height := getImageDimension(file)
		pdf.AddPageFormat("", gofpdf.SizeType{Wd: float64(width), Ht: float64(height)})
		pdf.ImageOptions(file, 0, 0, float64(width), float64(height), false,
			gofpdf.ImageOptions{}, 0, "")

		// trigger event handler
		if g.imageProcessingEvent != nil {
			g.imageProcessingEvent(i, file)
		}
	}
	return pdf.OutputFileAndClose(name)
}

// ImageProcessingEvent attaches an event handler when a file is processed.
// The first parameter of the event handler is the ID of the image files, starting from 0.
// The second parameter is the filename with full path.
func (g *Generator) ImageProcessingEvent(fn func(id int, filename string)) {
	g.imageProcessingEvent = fn
}

func getImageDimension(file string) (width int, height int) {
	f, err := os.Open(file)
	if err != nil {
		return 0, 0
	}
	defer func() { _ = f.Close() }()

	imgConfig, _, err := image.DecodeConfig(f)
	if err != nil {
		return 0, 0
	}
	return imgConfig.Width, imgConfig.Height
}
