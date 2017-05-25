// Command gophergrid creates a PNG image with a grid of tiled gophers.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/mdlayher/imagegrid"
)

func main() {
	var (
		flagN = flag.Int("n", 5, "number of images per row or column")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nBy default, %s will print a PNG image to stdout.\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\nIt is recommended to redirect stdout to a file or pipe.")
	}

	flag.Parse()
	log.SetOutput(os.Stderr)

	bufs := [][]byte{
		gopherBlackWhite,
		gopherColor,
	}

	images := make([]image.Image, 0, len(bufs))
	for _, b := range bufs {
		img, err := png.Decode(bytes.NewReader(b))
		if err != nil {
			log.Fatalf("failed to decode PNG image: %v", err)
		}

		images = append(images, img)
	}

	out, err := imagegrid.Draw(*flagN, images)
	if err != nil {
		log.Fatalf("failed to draw grid: %v", err)
	}

	if err := png.Encode(os.Stdout, out); err != nil {
		log.Fatalf("failed to encode PNG image: %v", err)
	}
}
