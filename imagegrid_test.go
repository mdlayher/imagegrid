package imagegrid

import (
	"image"
	"image/color"
	"testing"
)

var (
	red   = color.RGBA{255, 0, 0, 255}
	green = color.RGBA{0, 255, 0, 255}
	blue  = color.RGBA{0, 0, 255, 255}
	white = color.RGBA{255, 255, 255, 255}
)

func TestDrawOK(t *testing.T) {
	const (
		imageLength = 1
		gridLength  = 2
		numImages   = 4
	)

	out, err := Draw(gridLength, makeImages(numImages, imageLength))
	if err != nil {
		t.Fatalf("failed to draw grid: %v", err)
	}

	// Expect colors to be tiled as:
	//
	//  red | green
	// -------------
	// blue | white
	colors := []color.Color{
		red,
		green,
		blue,
		white,
	}

	var n int
	for y := 0; y < out.Bounds().Dy(); y++ {
		for x := 0; x < out.Bounds().Dx(); x++ {
			if want, got := colors[n], out.At(x, y); want != got {
				t.Fatalf("unexpected color at (%d,%d):\n- want: %v\n-  got: %v",
					x, y, want, got)
			}
			n++
		}
	}
}

func makeImage(length int) func() image.Image {
	var i int
	return func() image.Image {
		defer func() { i++ }()

		img := image.NewRGBA(image.Rect(0, 0, length, length))

		var col color.Color
		switch {
		case i%4 == 0:
			col = red
		case i%4 == 1:
			col = green
		case i%4 == 2:
			col = blue
		case i%4 == 3:
			col = white
		}

		for j := 0; j < length; j++ {
			for k := 0; k < length; k++ {
				img.Set(j, k, col)
			}
		}

		return img
	}
}

func makeImages(n, length int) []image.Image {
	img := makeImage(length)

	images := make([]image.Image, 0, n)
	for i := 0; i < n; i++ {
		images = append(images, img())
	}

	return images
}
