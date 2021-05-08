package ImageTools

import (
	"errors"
	"math"
	"sync"
)

/*
 * Calculates the pixelwise sum of two images
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range
 */
func AddImage(a [][]float32, b [][]float32, normalise bool) ([][]float32, error) {

	// Check that dimensions match
	if len(a) != len(b) {
		return nil, errors.New("Width mismatch")
	}
	if len(a[0]) != len(b[0]) {
		return nil, errors.New("Height mismatch")
	}
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Add corresponding pixel values
				outputImage[i][j] = a[i][j] + b[i][j]
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the sum of an image and a scalar
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range, if the normalise arg is true
 */
func AddScalar(a [][]float32, b float32, normalise bool) ([][]float32, error) {

	// Get dimensions
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Add pixels
				outputImage[i][j] = a[i][j] + b
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the pixelwise difference of two images
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range
 */
func SubtractImage(a [][]float32, b [][]float32, normalise bool) ([][]float32, error) {

	// Check that dimensions match
	if len(a) != len(b) {
		return nil, errors.New("Width mismatch")
	}
	if len(a[0]) != len(b[0]) {
		return nil, errors.New("Height mismatch")
	}
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Subtract corresponding pixel values
				outputImage[i][j] = a[i][j] - b[i][j]
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the difference of an image and a scalar
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range, if the normalise arg is true
 */
func SubtractScalar(a [][]float32, b float32, normalise bool) ([][]float32, error) {

	// Get dimensions
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Subtract pixels
				outputImage[i][j] = a[i][j] - b
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the pixelwise product of two images
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range
 */
func MultiplyImage(a [][]float32, b [][]float32, normalise bool) ([][]float32, error) {

	// Check that dimensions match
	if len(a) != len(b) {
		return nil, errors.New("Width mismatch")
	}
	if len(a[0]) != len(b[0]) {
		return nil, errors.New("Height mismatch")
	}
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Multiply corresponding pixel values
				outputImage[i][j] = a[i][j] * b[i][j]
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the product of an image and a scalar
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range, if the normalise arg is true
 */
func MultiplyScalar(a [][]float32, b float32, normalise bool) ([][]float32, error) {

	// Get dimensions
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Multiply pixels
				outputImage[i][j] = a[i][j] * b
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the pixelwise quotient of two images
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range
 */
func DivideImage(a [][]float32, b [][]float32, normalise bool) ([][]float32, error) {

	// Check that dimensions match
	if len(a) != len(b) {
		return nil, errors.New("Width mismatch")
	}
	if len(a[0]) != len(b[0]) {
		return nil, errors.New("Height mismatch")
	}
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Divide corresponding pixel values (making sure we don't try and divide by zero!!)
				if b[i][j] == 0 {
					outputImage[i][j] = 0
					break
				}
				outputImage[i][j] = a[i][j] / b[i][j]
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the quotient of an image and a scalar
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range, if the normalise arg is true
 */
func DivideScalar(a [][]float32, b float32, normalise bool) ([][]float32, error) {

	// Get dimensions
	imageWidth, imageHeight := Dimensions(a)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Divide pixels
				outputImage[i][j] = a[i][j] / b
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage), nil
	}
	return outputImage, nil
}

/*
 * Calculates the pixelwise square root of an image
 * The output is normalised in the range 0-1 (inclusive), while preserving dynamic range
 */
func Sqrt(image [][]float32, normalise bool) [][]float32 {

	imageWidth, imageHeight := Dimensions(image)

	// Create output image
	outputImage := make([][]float32, imageWidth)
	for j := range outputImage {
		outputImage[j] = make([]float32, imageHeight)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(imageHeight)

	// Iterate over columns
	for j := 0; j < imageHeight; j++ {

		// Process each row on its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over row
			for i := 0; i < imageWidth; i++ {

				// Square root each pixel
				outputImage[i][j] = float32(math.Sqrt(math.Abs(float64(image[i][j]))))
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Return result
	if normalise {
		return Normalise(outputImage)
	}
	return outputImage
}