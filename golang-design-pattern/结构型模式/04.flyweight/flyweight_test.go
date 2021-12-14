package flyweight

import "testing"

func TestNewFlyWeight(t *testing.T) {
	v1 := NewImageWeightViewer("image1.png")
	v2 := NewImageWeightViewer("image1.png")
	if v1.ImageFlyWeight != v2.ImageFlyWeight {
		t.Fatal()
	}
}

func ExampleNewFlyWeight() {
	viewer := NewImageWeightViewer("image1.png")
	viewer.Display()
	// Output:
	// Display:image data image1.png
}

