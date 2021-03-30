/**
GIF动画生成
go run 5.go web
*/

package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	//"strconv"
	"time"
)

var palette = []color.Color{
	color.RGBA{0, 100, 0, 255},
	color.Black,
	color.RGBA{250, 250, 210, 255},
	color.RGBA{255, 140, 0, 255}}

const (
	whiteIndex = 0 //画板中第一种颜色
	blackIndex = 1 //画板中的下一种颜色
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	if len(os.Args) > 1 && os.Args[1] == "web" {
		hander := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w, r)
		}
		http.HandleFunc("/", hander)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
}

func lissajous(out io.Writer, req *http.Request) {
	const (
		//cycles  = 5     // 完整的x振荡器变化的个数
		res     = 0.001 // 角度分辨率
		size    = 100   // 图像画布包含 [-size..+size]
		nframes = 64    // 动画中的帧数
		delay   = 8     // 以10ms为单位的帧间延迟
	)

	query := req.URL.Query()
	cycles, err := strconv.ParseFloat(query.Get("cycles"), 64)
	if err != nil {
		cycles = 5
	}

	freq := rand.Float64() * 3.0 // y振荡器的相对频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t + freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(len(palette)-1)+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意: 忽略编码错误
}
