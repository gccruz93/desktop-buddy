package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/gif"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var loadedGifs map[string][]*ebiten.Image

func loadGif(name string) {
	if _, ok := loadedGifs[name]; ok {
		return
	}
	file, err := os.ReadFile("./assets/" + name + ".gif")
	if err != nil {
		log.Println("loadGif ERR: " + name)
		var frames []*ebiten.Image
		frame := image.NewRGBA(image.Rect(0, 0, 192, 192))
		placeholder, _, err := image.Decode(bytes.NewReader(icon))
		if err != nil {
			log.Println(err)
			return
		}
		draw.Draw(frame, frame.Bounds(), placeholder, image.Point{}, draw.Over)
		loadedGifs[name] = append(frames, ebiten.NewImageFromImage(frame))
	} else {
		loadedGif, _ := gif.DecodeAll(bytes.NewReader(file))
		loadedGifs[name] = splitAnimatedGIF(loadedGif)
	}
}

func splitAnimatedGIF(gif *gif.GIF) []*ebiten.Image {
	var frames []*ebiten.Image
	imgWidth, imgHeight := getGifDimensions(gif)

	for _, srcImg := range gif.Image {
		frame := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))
		draw.Draw(frame, frame.Bounds(), srcImg, image.Point{}, draw.Over)
		frames = append(frames, ebiten.NewImageFromImage(frame))
	}

	return frames
}
func getGifDimensions(gif *gif.GIF) (x, y int) {
	var lowestX, lowestY, highestX, highestY int
	for _, img := range gif.Image {
		if img.Rect.Min.X < lowestX {
			lowestX = img.Rect.Min.X
		}
		if img.Rect.Min.Y < lowestY {
			lowestY = img.Rect.Min.Y
		}
		if img.Rect.Max.X > highestX {
			highestX = img.Rect.Max.X
		}
		if img.Rect.Max.Y > highestY {
			highestY = img.Rect.Max.Y
		}
	}
	return highestX - lowestX, highestY - lowestY
}
