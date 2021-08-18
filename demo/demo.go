// demo.go
package main

import (
	"fmt"
	"image"
	"os"
	"time"

	painter "github.com/odorajbotoj/mcws-painter"
)

func main() {
	fmt.Println("Hello World!")
	file, err := os.Open("./grass.png")
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
	p := painter.MapImage(img, painter.DefaultColorSpace, 61, 10*time.Millisecond)
}
