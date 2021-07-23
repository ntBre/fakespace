package main

import (
	crypt "crypto/rand"
	"flag"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

var (
	YMAX = flag.Int("y", 500, "height of image in pixels")
	XMAX = flag.Int("x", int(math.Round(3.25/1.75))**YMAX,
		"width of image in pixels; defaults to ratio for JPCA TOC graphic")
	RED_MAX = flag.Int("red", 10, "number of large red stars")
	BLU_MAX = flag.Int("blue", 10, "number of large blue stars")
	rad     = flag.Int("size", 2, "radius of large stars in pixels")
)

func AddStar(img *image.NRGBA, col color.NRGBA) {
	rxl, _ := crypt.Int(crypt.Reader, big.NewInt(int64(*XMAX)))
	ryl, _ := crypt.Int(crypt.Reader, big.NewInt(int64(*YMAX)))
	x, y := int(rxl.Int64()), int(ryl.Int64())
	for rx := x - *rad; rx <= x+*rad && rx <= *XMAX; rx++ {
		for ry := y - *rad; ry <= y+*rad && ry <= *YMAX; ry++ {
			yy := ry - y
			xx := rx - x
			if *rad**rad >= xx*xx+yy*yy {
				img.Set(rx, ry, col)
			}
		}
	}
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		panic("not enough arguments")
	}
	rand.Seed(int64(time.Now().Nanosecond()))
	img := image.NewNRGBA(image.Rect(0, 0, *XMAX, *YMAX))
	f, err := os.Create(args[0])
	if err != nil {
		panic(err)
	}
	for x := 0; x <= *XMAX; x++ {
		for y := 0; y <= *YMAX; y++ {
			r := rand.Float64()
			if r < 0.004 {
				img.Set(x, y, color.NRGBA{255, 255, 255, 255})
			} else {
				img.Set(x, y, color.NRGBA{0, 0, 0, 255})
			}
		}
	}
	// red ones
	for i := 0; i < *RED_MAX; i++ {
		AddStar(img, color.NRGBA{
			R: 255,
			G: 112,
			B: 3,
			A: 150 | uint8(rand.Intn(255)),
		})
	}
	// blue ones
	for i := 0; i < *BLU_MAX; i++ {
		AddStar(img, color.NRGBA{
			R: 0,
			G: 112,
			B: 255,
			A: 150 | uint8(rand.Intn(255)),
		})
	}
	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
