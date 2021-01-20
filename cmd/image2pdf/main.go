package main

import (
	"fmt"
	"image2pdf/fileutils"
	"image2pdf/img2pdf"
	"os"
	"path/filepath"
)

const usage = `Usage: %s <directory that contains images> [<directory2 that contains images> ...]
You can drag and drop folders onto the script.
`

func main() {
	if len(os.Args) > 1 {
		for _, directory := range os.Args[1:] {
			processDirectory(directory)
		}
		fmt.Print("Press any key to continue...")
		_, _ = fmt.Scanln()
	} else {
		printUsage()
	}
}

func printError(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, a...)
}

func processDirectory(directory string) {
	images, err := fileutils.GetImages(directory)
	if err != nil {
		printError(err.Error())
		return
	}

	if len(images) > 0 {
		pdf := fileutils.GeneratePDFName(directory)
		gen := img2pdf.Generator{}
		gen.ImageProcessingEvent(func(id int, filename string) {
			fmt.Printf("%d: %s\n", id+1, filepath.Base(filename))
		})
		err := gen.CreatePDF(pdf, images)
		if err != nil {
			printError("An error occurred when creating PDF file: %s\n%v\n", pdf, err)
			return
		}
		fmt.Printf("PDF file created: %s with %d image file(s).\n", pdf, len(images))
	}
}

func printUsage() {
	printError(usage, os.Args[0])
}
