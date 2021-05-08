package kernels

import (
	"math"
	"sync"
)

/*
 * Generates a Gaussian kernel of a given size and standard deviation
 */
func Gaussian(size int, sigma float32) [][]float32 {

	// Create empty kernel
	kernel := make([][]float32, size)
	for j := range kernel {
		kernel[j] = make([]float32, size)
	}

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(size)

	// Iterate over columns
	for j := 0; j < size; j++ {

		// Process each row in its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over every pixel in the row
			for i := 0; i < size; i++ {

				a := 1 / (2 * math.Pi * float64(sigma) * float64(sigma))
				b := math.Pow(math.E, -1 * ((float64(i) * float64(i) + float64(j) * float64(j)) / (2 * float64(sigma) * float64(sigma))))

				kernel[i][j] = float32(a * b)
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	return NormaliseKernel(kernel)
}

/*
 * Nornalises a kernel by ensuring that all of its elements sum to 1
 */
func NormaliseKernel(kernel [][]float32) [][]float32 {

	// Calculate sum
	sum := float64(0)
	for j := 0; j < len(kernel[0]); j++ {
		for i := 0; i < len(kernel); i++ {
			sum += float64(kernel[i][j])
		}
	}
	inverseSum := 1 / sum

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(kernel[0]))

	// Iterate over columns
	for j := 0; j < len(kernel[0]); j++ {

		// Process each row in its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over every pixel in the row
			for i := 0; i < len(kernel); i++ {

				// Multiply every element by the inverse of the sum
				newElement := float64(kernel[i][j]) * inverseSum
				kernel[i][j] = float32(newElement)
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	return kernel
}