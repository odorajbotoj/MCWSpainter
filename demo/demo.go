// demo.go
package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
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
	if err != nil {
		panic(err)
	}
	wx, err := strconv.ParseInt(worldX, 10, 64)
	if err != nil {
		panic(err)
	}
	wy, err := strconv.Atoi(worldY)
	if err != nil {
		panic(err)
	}
	wz, err := strconv.ParseInt(worldZ, 10, 64)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	p := painter.MapImage(img, painter.DefaultColorSpace, wy, 10*time.Millisecond)
	size := p.GetSize()
	for i := 0; i < size; i++ {
		ax, _, az, bn, dv := p.Next()
		// 此处禁用y轴变化, 即导入平面地图画
		fmt.Printf("setblock %d %d %d %s %d\n", wx+ax, wy, wz+az, bn, dv)
	}
	fmt.Println("say ====TASK FINISH====")
}
