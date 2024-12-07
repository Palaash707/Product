package imageutils

import (
	"image/jpeg"
	"os"

	"github.com/Palaash707/Product/internal/logging"
	"github.com/disintegration/imaging"
	"go.uber.org/zap"
)

// CompressImage compresses the uploaded image
func CompressImage(inputPath, outputPath string) error {
    // Log start of image processing
    logging.Logger.Info("Starting image compression",
        zap.String("input_path", inputPath),
        zap.String("output_path", outputPath),
    )

    // Open the image
    img, err := imaging.Open(inputPath)
    if err != nil {
        logging.Logger.Error("Failed to open image",
            zap.String("input_path", inputPath),
            zap.Error(err),
        )
        return err
    }

    // Resize and compress the image
    img = imaging.Resize(img, 800, 600, imaging.Lanczos)

    // Create the output file for the compressed image
    outFile, err := os.Create(outputPath)
    if err != nil {
        logging.Logger.Error("Failed to create output file",
            zap.String("output_path", outputPath),
            zap.Error(err),
        )
        return err
    }
    defer outFile.Close()

    // Save as JPEG with quality set to 80
    err = jpeg.Encode(outFile, img, &jpeg.Options{Quality: 80})
    if err != nil {
        logging.Logger.Error("Failed to compress image",
            zap.String("output_path", outputPath),
            zap.Error(err),
        )
        return err
    }

    logging.Logger.Info("Image compression successful",
        zap.String("input_path", inputPath),
        zap.String("output_path", outputPath),
    )

    return nil
}
