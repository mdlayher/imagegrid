// Package imagegrid enables composing one or more images into a single
// image, using a tiled grid pattern.
package imagegrid

import (
	"errors"
	"image"
	"image/draw"
)

// Draw combines a slice of image.Images into a single image, tiling the
// images from left-to-right, then top-to-bottom.
//
// n determines how many images will be drawn in each row or column.  n
// must be greater than zero.
//
// At least one image must be passed in the slice. Each image must have
// identical dimensions, and must be perfectly square.
func Draw(n int, images []image.Image) (image.Image, error) {
	if n < 1 {
		return nil, errors.New("n must be greater than zero")
	}
	if len(images) < 1 {
		return nil, errors.New("at least one image must be provided")
	}

	// All images are expected to be perfectly square, and width and height
	// must equal this length.
	length := images[0].Bounds().Dx()
	for _, img := range images {
		if length != img.Bounds().Dx() || length != img.Bounds().Dy() {
			return nil, errors.New("all images must be perfectly square and have same dimensions")
		}
	}

	out := image.NewRGBA(image.Rect(0, 0, length*n, length*n))

	var idx int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			r := image.Rect(
				// Use j for X-axis to draw images from left to right, then
				// top to bottom.
				j*length,
				i*length,
				(j*length)+length,
				(i*length)+length,
			)

			// Reset index if out of range.
			if idx == len(images) {
				idx = 0
			}

			draw.Draw(out, r, images[idx], image.ZP, draw.Over)
			idx++
		}
	}

	return out, nil
}
