package ImageTools

import (
	"ImageTools/kernels"
	"math"
	"sync"
)

/*
 * Applies a kernel convolution to an image
 */
func Convolution(image [][]float32, kernel [][]float32, normalise bool) [][]float32 {

	imageWidth, imageHeight := Dimensions(image)
	kernelWidth, kernelHeight := Dimensions(kernel)
	halfKernelWidth, halfKernelHeight := int(math.Ceil(0.5 * float64(kernelWidth))), int(math.Ceil(0.5 * float64(kernelHeight)))

	// Pad image with a border of zeros big enough to prevent the kernel from going over the edge
	paddedImage := make([][]float32, imageWidth + halfKernelWidth + halfKernelWidth)
	for j := range paddedImage {
		paddedImage[j] = make([]float32, imageHeight + halfKernelHeight + halfKernelHeight)
	}
	for j := 0; j < imageHeight; j++ {
		for i := 0; i < imageWidth; i++ {
			paddedImage[i + halfKernelWidth][j + halfKernelHeight] = image[i][j]
		}
	}

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

		// Process each row in its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over every pixel in the row
			for i := 0; i < imageWidth; i++ {

				// Calculate dot product of kernel and local pixels at every pixel
				accumulator := float64(0)
				for kJ := 0; kJ < kernelHeight; kJ++ {
					for kI := 0; kI < kernelWidth; kI++ {
						imageValue := float64(paddedImage[i+(kI-halfKernelWidth)+halfKernelWidth][j+(kJ-halfKernelHeight)+halfKernelHeight])
						kernelValue := float64(kernel[kI][kJ])
						accumulator += imageValue * kernelValue
					}
				}
				outputImage[i][j] = float32(accumulator)
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	if normalise {
		return Normalise(outputImage)
	}
	return outputImage
}

/*
 * Applies a separated kernel convolution to an image
 */
func SepConvolution(image [][]float32, kernelA [][]float32, kernelB [][]float32, normalise bool) [][]float32 {

	// Apply first kernel
	image = Convolution(image, kernelA, false)

	// Apply second kernel
	image = Convolution(image, kernelB, false)

	if normalise {
		return Normalise(image)
	}
	return image
}

/*
 * Calculates the gradient magnitude at each pixel in an image
 */
func GradientMagnitude(image [][]float32) [][]float32 {

	// Create channels
	ch1 := make(chan [][]float32)
	ch2 := make(chan [][]float32)

	// Make copy of image
	cpImage := image

	// Apply Sobel filters
	go func() { ch1 <- SepConvolution(image, kernels.SepSobelXPt1, kernels.SepSobelXPt2, true) }()
	go func() { ch1 <- SepConvolution(cpImage, kernels.SepSobelYPt1, kernels.SepSobelYPt2, true) }()
	a := <- ch1
	b := <- ch1

	// Calculate cross products
	go func() {
		crossProduct, _ := MultiplyImage(a, a, false)
		ch2 <- crossProduct
	}()
	go func() {
		crossProduct, _ := MultiplyImage(b, b, false)
		ch2 <- crossProduct
	}()
	a = <- ch2
	b = <- ch2

	// Sum cross products
	a, _ = AddImage(a, b, false)

	// Square root sum
	b = Sqrt(a, false)


	// Normalise in the range 0-1 (inclusive) and return
	return Normalise(b)
}

/*
 * Calculates the orientation of each pixel in an image
 */
func PixelOrientation(image [][]float32) [][]float32 {

	// Create channels
	ch1 := make(chan [][]float32)
	ch2 := make(chan [][]float32)

	// Apply Sobel filters
	go func() { ch1 <- SepConvolution(image, kernels.SepSobelXPt1, kernels.SepSobelXPt2, true) }()
	go func() { ch2 <- SepConvolution(image, kernels.SepSobelYPt1, kernels.SepSobelYPt2, true) }()
	gx := <- ch1
	gy := <- ch2

	// Calculate quotient
	image, _ = DivideImage(gy, gx, false)

	// Calculate arctangent
	image = Atan(image)

	// Normalise in the range 0-1 (inclusive) and return
	return Normalise(image)
}