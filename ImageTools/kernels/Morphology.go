package kernels

/*
 * Generates a structuring element for use in binary morphological operations
 * (It's just a load of 1s in a matrix of a given size)
 */
func BinaryErosionDilationStructuringElement(size int) [][]float32 {

	// Create structuring element of desired size
	structuringElement := make([][]float32, size)
	for j := range structuringElement {
		structuringElement[j] = make([]float32, size)
		for i := 0; i < size; i++ {
			structuringElement[j][i] = 1.0
		}
	}

	return structuringElement
}