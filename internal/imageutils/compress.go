package imageutils

import (
	"image/jpeg"
	"os"

	"github.com/disintegration/imaging"
)

// CompressImage compresses the uploaded image
func CompressImage(inputPath, outputPath string) error {
    img, err := imaging.Open(inputPath)
    if err != nil {
        return err
    }

    // Resize and compress the image
    img = imaging.Resize(img, 800, 600, imaging.Lanczos)

    outFile, err := os.Create(outputPath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    // Save as JPEG with quality set to 80
    err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 80})
    return err
}
