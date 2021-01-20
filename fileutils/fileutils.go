package fileutils

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/pkg/errors"
)

var supportedExtNames = []string{".jpg", ".png"}

// GetImages gets paths of all image files under the specified path.
func GetImages(path string) ([]string, error) {
	result := make([]string, 0)
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			// ignore sub-directories
			return nil
		}

		if isImageType(info.Name()) {
			result = append(result, path)
		}

		return nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "an error occurred when walking directory: "+path)
	}

	sort.Strings(result)
	return result, nil
}

func isImageType(filename string) bool {
	filename = strings.ToLower(filename)
	for _, ext := range supportedExtNames {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

// GeneratePDFName generates full path and filename of the output PDF based on the path given.
// For example, if the path containing image files is "/a/b/c/d", the generated PDF file will be "/a/b/c/d.pdf".
func GeneratePDFName(path string) string {
	return filepath.Join(filepath.Dir(path), filepath.Base(path)+".pdf")
}
