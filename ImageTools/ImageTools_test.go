package ImageTools

import (
	"ImageTools/kernels"
	"fmt"
	"testing"
)

func TestSignatureDifference(t *testing.T) {
	imgA, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}
	imgB, err := LoadImage("test-images/01-original.jpg")
	if err != nil {
		t.Fatal()
	}
	imgC, err := LoadImage("test-images/02-original.jpg")
	if err != nil {
		t.Fatal()
	}
	imgD, err := LoadImage("test-images/03-original.jpg")
	if err != nil {
		t.Fatal()
	}

	// Calculate signatures
	sigA := SignatureVector(imgA)
	sigB := SignatureVector(imgB)
	sigC := SignatureVector(imgC)
	sigD := SignatureVector(imgD)

	// A signature compared with itself should be zero
	if SignatureDifference(sigA, sigA) != 0 {
		fmt.Println("A and A not equal to 0")
		t.Fail()
	}
	if SignatureDifference(sigB, sigB) != 0 {
		fmt.Println("B and B not equal to 0")
		t.Fail()
	}
	if SignatureDifference(sigC, sigC) != 0 {
		fmt.Println("C and C not equal to 0")
		t.Fail()
	}
	if SignatureDifference(sigD, sigD) != 0 {
		fmt.Println("D and D not equal to 0")
		t.Fail()
	}

	// The difference between A and B
	ab, ba := SignatureDifference(sigA, sigB), SignatureDifference(sigB, sigA)
	fmt.Println("Signature A:", sigA)
	fmt.Println("Signature B:", sigB)
	fmt.Println("Difference:", (ab + ba) / 2)

	// The difference between C and D
	cd, dc := SignatureDifference(sigC, sigD), SignatureDifference(sigD, sigC)
	fmt.Println("Signature C:", sigC)
	fmt.Println("Signature D:", sigD)
	fmt.Println("Difference:", (cd + dc) / 2)

	// The difference between A and B should be greater than the difference between C and D
	if !(ab > cd) {
		fmt.Println("The difference between A and B should be greater than the difference between C and D")
		fmt.Println(ab, cd)
		t.Fail()
	}
}

func TestSubImage(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	got := SubImage(img, 1000, 1000, 1000, 1000)
	err = SaveImage("test-images/TestSubImage__00-within-bounds.jpg", got)
	if err != nil {
		t.Fail()
	}

	//width, height := Dimensions(img)
	got = SubImage(img, -100, -200, 5000, 5000)
	err = SaveImage("test-images/TestSubImage__01-out-of-bounds.jpg", got)
	if err != nil {
		t.Fail()
	}
}

func TestNormalise(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	err = SaveImage("test-images/TestNormalise__00-not-normalised.jpg", img)
	if err != nil {
		t.Fail()
	}

	got := Normalise(img)
	err = SaveImage("test-images/TestNormalise__01-normalised.jpg", got)
	if err != nil {
		t.Fail()
	}
}

func TestSepConvolution(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	convolution := Convolution(img, kernels.SobelX, true)
	sepConvolution := SepConvolution(img, kernels.SepSobelXPt1, kernels.SepSobelXPt2, true)

	// Check that they are the same (or close enough)
	_, mae, _ := AbsoluteError(convolution, sepConvolution)
	if mae > 0.01 {
		fmt.Println("MAE:", mae)
		t.Fail()
	}
}

func TestConvolution(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	laplacian := Convolution(img, kernels.Laplacian, true)
	err = SaveImage("test-images/TestConvolution__00-laplacian.jpg", laplacian)
	if err != nil {
		t.Fail()
	}
	mean, std := MeanStd(laplacian)

	got := SingleThreshold(laplacian, mean)
	err = SaveImage("test-images/TestConvolution__01-laplacian-single-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}

	got = DualThreshold(laplacian, mean - 0.5*std, mean + 0.5*std)
	err = SaveImage("test-images/TestConvolution__02-laplacian-dual-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}

	blur := Convolution(img, kernels.Gaussian(5, 8), true)
	err = SaveImage("test-images/TestConvolution__03-gaussian.jpg", blur)
	if err != nil {
		t.Fail()
	}
	mean, std = MeanStd(laplacian)

	got = SingleThreshold(blur, mean)
	err = SaveImage("test-images/TestConvolution__04-gaussian-single-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}

	got = DualThreshold(blur, mean - 0.5*std, mean + 0.5*std)
	err = SaveImage("test-images/TestConvolution__05-gaussian-dual-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}
}

func TestGradientMagnitude(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	gm := GradientMagnitude(img)
	err = SaveImage("test-images/TestGradientMagnitude__00-gradient-magnitude.jpg", gm)
	if err != nil {
		t.Fail()
	}
	mean, std := MeanStd(gm)

	got := SingleThreshold(gm, mean)
	err = SaveImage("test-images/TestGradientMagnitude__01-single-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}

	got = DualThreshold(gm, mean - 0.5*std, mean + 0.5*std)
	err = SaveImage("test-images/TestGradientMagnitude__02-dual-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}
}

func TestPixelOrientation(t *testing.T) {
	img, err := LoadImage("test-images/00-original.jpg")
	if err != nil {
		t.Fatal()
	}

	po := PixelOrientation(img)
	err = SaveImage("test-images/TestPixelOrientation__00-pixel-orientation.jpg", po)
	if err != nil {
		t.Fail()
	}
	mean, std := MeanStd(po)

	got := SingleThreshold(po, mean)
	err = SaveImage("test-images/TestPixelOrientation__01-single-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}

	got = DualThreshold(po, mean - 0.5*std, mean + 0.5*std)
	err = SaveImage("test-images/TestPixelOrientation__02-dual-threshold.jpg", got)
	if err != nil {
		t.Fail()
	}
}