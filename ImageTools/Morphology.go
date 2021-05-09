package ImageTools

import "ImageTools/kernels"

/*
 * Performs morphological erosion of a binarised image
 */
func BinaryErosion(image [][]float32, size int) [][]float32 {

	// Create structuring element of desired size
	structuringElement := kernels.BinaryErosionDilationStructuringElement(size)

	// Apply structuring element
	summed := Convolution(image, structuringElement, false)

	// Apply threshold
	return SingleThreshold(summed, float32(size) * float32(size) * 0.75)
}

/*
 * Performs morphological dilation of a binarised image
 */
func BinaryDilation(image [][]float32, size int) [][]float32 {

	// Create structuring element of desired size
	structuringElement := kernels.BinaryErosionDilationStructuringElement(size)

	// Apply structuring element
	return Convolution(image, structuringElement, true)
}

/*
 * Performs morphological opening of a binarised image
 */
func BinaryOpening(image [][]float32, size int) [][]float32 {

	// Erode
	image = BinaryErosion(image, size)

	// Dilate
	return BinaryDilation(image, size)
}

/*
 * Performs morphological closing of a binarised image
 */
func BinaryClosing(image [][]float32, size int) [][]float32 {

	// Dilate
	image = BinaryDilation(image, size)

	// Erode
	return BinaryErosion(image, size)
}