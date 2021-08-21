// demo.go
package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"time"

	painter "github.com/odorajbotoj/MCWSpainter"
)

func main() {
	filename := os.Args[1]
	worldX := os.Args[2]
	worldY := os.Args[3]
	worldZ := os.Args[4]
	file, err := os.Open(filename)
	wx := strconv.ParseInt(worldX, 10, 64)
	wz := strconv.ParseInt(worldZ, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	img, err := image.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := MapImage(img, DefaultColorSpace, worldY, 10*time.Millisecond)
	for i := 0; i < p.GetSize(); i++ {
		ax, ay, az, bn, dv := p.Next()
		fmt.Printf("setblock %d %d %d %s %d\n", wx+ax, ay, wz+az, bn, dv)
	}
	fmt.Println("END")
}
