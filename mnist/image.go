package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"os/exec"
)

func Bytes2Image(rows, cols int, b []byte) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, rows-1, cols-1))
	idx := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, color.Gray16{(255 - uint16(b[idx])) << 8})
			idx++
		}
	}
	return img
}

func CreatePNGAndShow(img *image.RGBA, lbl byte, fIdx int) {
	// Limit check
	if fIdx >= 20 {
		return
	}

	// Create file
	p, e := os.Create(fmt.Sprintf("%d-%d.png", lbl, fIdx))
	if e != nil {
		log.Panic(e)
	}
	defer p.Close()

	// Encode
	e = png.Encode(p, img)
	if e != nil {
		log.Panic(e)
	}

	// Preview it
	Show(p.Name())
}

func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
