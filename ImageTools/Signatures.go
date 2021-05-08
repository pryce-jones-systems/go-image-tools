package ImageTools

import (
	"math"
	"sync"
)

/*
 * Computes a signature of an image using the algorithm described in https://doi.org/10.1109/ICIP.2002.1038047
 */
func SignatureVector(image [][]float32) []int {

	// Create an 11x11 matrix to represent ROI averages.
	// Has additional rows and columns of zeros so that the 8-neighbourhood can be computed for every ROI
	average := make([][]float32, 11)
	for j := range average {
		average[j] = make([]float32, 11)
	}

	// Compute the distance between ROIs, such that there are 81 of them evenly spaced
	width, height := Dimensions(image)
	xDistance := int(math.Floor(float64(width) / 11))
	yDistance := int(math.Floor(float64(height) / 11))

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(9)

	// Iterate over columns
	for j := 1; j < 10; j++ {

		// Process each row in its own goroutine
		go func(j int) {
			defer waitGroup.Done()

			// Iterate over the row
			for i := 1; i < 10; i++ {

				// Get ROI
				roi := SubImage(image, (xDistance * i) - 2, (yDistance * j) - 2, 5, 5)

				// Average all pixels in ROI
				average[i][j], _ = MeanStd(roi)
			}
		} (j)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Create empty signature vector
	var signature []int

	// Iterate over the averages
	for aJ := 1; aJ < 10; aJ++ {
		for aI := 1; aI < 10; aI++ {

			// Iterate over the 8-neighbourhood
			for nJ := -1; nJ < 2; nJ++ {
				for nI := -1; nI < 2; nI++ {

					// TODO make this bit branchless for a bit of a speed boost

					// Make sure we don't compare an average with itself
					if !((nI == 0) && (nJ == 0)) {

						// Classify the relationship between neighbouring averages based on how much darker/lighter the averages are
						difference := average[aI][aJ] - average[aI + nI][aJ + nJ]
						if difference < -2 {
							signature = append(signature, -2) // Much darker
						} else if -2 <= difference && difference < 0 {
							signature = append(signature, -1) // Darker
						} else if 0 < difference && difference <= 2 {
							signature = append(signature, 1) // Lighter
						} else if 2 < difference {
							signature = append(signature, 2) // Much darker
						} else {
							signature = append(signature, 0) // The same
						}
					}

					/*
					// Make sure we don't compare an average with itself
					if nI == 0 && nJ == 0 {
						break
					}

					// Classify the relationship between neighbouring averages based on how much darker/lighter the averages are
					difference := average[aI][aJ] - average[aI + nI][aJ + nJ]
					if difference < -2 {
						signature = append(signature, -2) // Much darker
						break
					} else if -2 <= difference && difference < 0 {
						signature = append(signature, -1) // Darker
						break
					} else if 0 < difference && difference <= 2 {
						signature = append(signature, 1) // Lighter
						break
					} else if 2 < difference {
						signature = append(signature, 2) // Much darker
						break
					}
					signature = append(signature, 0) // The same
					 */
				}
			}
		}
	}

	return signature
}

/*
 * Computes the L2 Norm of a signature vector
 */
func L2Norm(signature []int) float32 {

	// Create wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(signature))

	// Iterate the signature vector
	accumulator := float64(0)
	for i := 0; i < len(signature); i++ {

		// Process each element on its own goroutine
		go func(i int) {
			defer waitGroup.Done()

			// Add the square of each element to the accumulator
			element := float64(signature[i])
			accumulator += element * element
		} (i)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	// Square root the sum of the squares
	return float32(math.Sqrt(accumulator))
}

/*
 * Calculates the normalised difference between two signature vectors
 * Distance from A to B is always the same as distance from B to A
 */
func SignatureDifference(sigA []int, sigB []int) float32 {

	// Calculate distance between A and B on its own goroutine
	abChan := make(chan  float64)
	go func(ch chan float64) {
		ch <- signatureDifference(sigA, sigB)
	} (abChan)

	// Calculate distance between B and A on its own goroutine
	baChan := make(chan  float64)
	go func(ch chan float64) {
		ch <- signatureDifference(sigB, sigA)
	} (baChan)

	// Wait for distances to be calculated
	ab, ba := <- abChan, <- baChan

	// Average the two distances
	return float32((ab + ba) / 2)
}

/*
 * Calculates the normalised difference between two signature vectors
 * Distance from A to B is not necessarily the same as distance from B to A
 */
func signatureDifference(sigA []int, sigB []int) float64 {

	// Make sure both signatures have the same number of dimensions
	if len(sigA) != len(sigB) {
		return math.MaxFloat32
	}

	// Calculate the difference between every element of sigA and sigB
	var differenceVector []int
	for i := 0; i < len(sigA); i++ {
		differenceVector = append(differenceVector, sigA[i] - sigB[i])
	}

	// Calculate || sigA - sigB || on its own goroutine
	differenceChan := make(chan float64)
	go func(ch chan float64) {
		ch <- float64(L2Norm(differenceVector))
	} (differenceChan)


	// Calculate || sigA || on its own goroutine
	aChan := make(chan float64)
	go func(ch chan float64) {
		ch <- float64(L2Norm(sigA))
	} (aChan)

	// Calculate || sigB || on its own goroutine
	bChan := make(chan float64)
	go func(ch chan float64) {
		ch <- float64(L2Norm(sigB))
	} (bChan)

	// Wait for all the L2 Norms to be calculated
	differenceL2, sigAL2, sigBL2 := <- differenceChan, <- aChan, <- bChan

	// Calculate || sigA - sigB || / (|| sigA || + || sigB ||)
	delta := differenceL2 / (sigAL2 + sigBL2)

	// Make sure the result is a number
	if math.IsNaN(delta) {
		return 0
	}
	return delta
}