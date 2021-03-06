// painter.go
package MCWSpainter

import (
	"image"
	"time"
)

// 求平方
func square(v uint8) uint32 {
	return uint32(v * v)
}

// Color: 用于存储RGB值
type Color struct {
	R uint8
	G uint8
	B uint8
}

// Block: 用于存储数据值与颜色的映射
type Block struct {
	Name      string
	DataValue int
	Color     Color
	DeltaY    int
}

// ColorSpace: 由Block构建的列表
type ColorSpace struct {
	Space []Block
}

// 用于匹配单个颜色, 返回Block{方块名称, 特殊值(int)和Y轴偏移量}
func (s *ColorSpace) MapColor(clr Color) Block {
	var maxDiff uint32 = 195075 // 255*255*3
	index := 0                  // 存储最接近的索引值
	for i, j := range s.Space {
		diff := square(j.Color.R-clr.R) + square(j.Color.G-clr.G) + square(j.Color.B-clr.B) // 求色差,但是不开平方
		if diff == 0 {
			return j
		} else if diff < maxDiff {
			maxDiff = diff
			index = i
		} else {
			continue
		}
	}
	return s.Space[index]
}

// CmdResponse: 坐标, 方块特殊值
type CmdResponse struct {
	X         int64
	DeltaY    int
	Z         int64
	Name      string
	DataValue int
}

// Painter: CmdResponse的集合, 加上延时
type Painter struct {
	Index int
	Resps []CmdResponse
	NowY  int
	Dealy time.Duration
}
type IPainter interface {
	Next() (int64, int64, int64, string, int)
}

// Next()不断获取一行新的指令参数. 还有, 我不懂迭代器
func (p *Painter) Next() (int64, int64, int64, string, int) {
	r := p.Resps[p.Index]
	rx := r.X
	p.NowY += r.DeltaY
	ry := p.NowY
	rz := r.Z
	rn := r.Name
	rd := r.DataValue
	p.Index++
	time.Sleep(p.Dealy)
	if p.Index == len(p.Resps)-1 {
		p.Index = 0
	}
	return rx, int64(ry), rz, rn, rd
}

// GetSize()返回需要调用Next()的次数.
func (p *Painter) GetSize() int {
	return len(p.Resps)
}

// 封装一个MapImage函数, 返回一个Painter(struct)
func MapImage(img image.Image, cs ColorSpace, y int, dealy time.Duration) Painter {
	bounds := img.Bounds()
	dx := bounds.Dx()
	dy := bounds.Dy()
	memory := make(map[Color]Block) // 怎么说呢, 作为一个优秀的算法, 它需要带脑子
	var ret = make([]CmdResponse, dx*dy)
	var index = 0
	// X循环
	for i := 0; i < dx; i++ {
		// Z循环, Z- 为北
		for j := 0; j < dy; j++ {
			colorRgb := img.At(i, j)
			r, g, b, a := colorRgb.RGBA()
			if a == 100 {
				ret[index] = CmdResponse{int64(i), 0, int64(j), "air", 0}
			}
			//转换为 255 值
			r_uint8 := uint8(r >> 8)
			g_uint8 := uint8(g >> 8)
			b_uint8 := uint8(b >> 8)
			pColor := Color{r_uint8, g_uint8, b_uint8}
			// 如果memory中检索到此值, 就直接存入返回值; 否则先保存
			blk, ok := memory[pColor]
			if ok {
				ret[index] = CmdResponse{int64(i), blk.DeltaY, int64(j), blk.Name, blk.DataValue}
			} else {
				blk = cs.MapColor(pColor)
				re := CmdResponse{int64(i), blk.DeltaY, int64(j), blk.Name, blk.DataValue}
				ret[index] = re
				memory[pColor] = blk
			}
			index++
		}
	}
	return Painter{0, ret, y, dealy}
}
