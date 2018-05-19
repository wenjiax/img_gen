package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/golang/freetype"
)

const (
	fontSize  float64 = 14              // font size
	fontWidth float64 = 4.92            // font fontWidth
	spacing   float64 = 1.5             // line spacing
	fontFile  string  = "./FZDBSJW.TTF" // font file
	dpi       float64 = 72              // screen resolution in dots per inch
)

var (
	// dpi = flag.Float64("dpi", 72, "sets the screen resolution in dots per inch")
	// fontFile = flag.String("font-file", "./FZDBSJW.TTF", "filename of the ttf font")
	imgTmp  = flag.String("t", "./temp1.png", "set image template")
	texts   = flag.String("s", "", "image content")
	outPath = flag.String("o", "./out.png", "set out image path")
)

func checkError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if *texts == "" {
		checkError(errors.New("texts is empty"))
	}
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontFile)
	checkError(err)
	font, err := freetype.ParseFont(fontBytes)
	checkError(err)

	// 读取图片模版
	bgfile, err := os.Open(*imgTmp)
	checkError(err)
	defer bgfile.Close()
	bgimg, err := png.Decode(bgfile)
	checkError(err)

	// 初始化前、背景颜色
	fg, bg := image.Black, image.White

	// 初始化上下文,设置属性
	c := freetype.NewContext()
	c.SetDPI(dpi)
	c.SetFont(font)
	c.SetFontSize(fontSize)

	// 获取图片模版宽、高
	bgImgWidth := bgimg.Bounds().Dx()
	bgImgHeight := bgimg.Bounds().Dy()

	// 计算每行文字数量
	width := bgImgWidth + (20 * 2)
	context := *texts
	var text []string
	// 每行字数
	lineCount1 := float64(width) / float64(fontWidth)
	lineCount := int(lineCount1)
	i, l := 0, len(context)
	for ; i < l-lineCount; i += lineCount {

		text = append(text, context[i:i+lineCount])
	}
	if i < l {
		text = append(text, context[i:])
	}
	// 计算输出图片的宽、高
	height := bgImgHeight + (c.PointToFixed(fontSize*spacing).Ceil()*len(text) + 50)

	// 创建输出图片
	rgba := image.NewRGBA(image.Rect(0, 0, width, height))

	// 画出背景和图片模版
	draw.Draw(rgba, rgba.Bounds(), bg, image.ZP, draw.Src)
	draw.Draw(rgba, rgba.Bounds(), bgimg, image.Point{X: -20, Y: 0}, draw.Src)
	// 设置属性
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	c.SetSrc(fg)

	// 生成标题
	now := time.Now()
	title := fmt.Sprintf("%d年%d月%d日  星期%d  %02d:%02d", now.Year(), now.Month(), now.Day(), now.Weekday(), now.Hour(), now.Minute())

	// 计算标题的位置
	x := (width - (len(title) * 5)) / 2
	tpt := freetype.Pt(x, 240+int(c.PointToFixed(fontSize)>>8))

	// 画出标题文字
	_, err = c.DrawString(title, tpt)
	// 画出文字内容
	pt := freetype.Pt(10, 265+int(c.PointToFixed(fontSize)>>8))
	for _, s := range text {
		_, err = c.DrawString(s, pt)
		checkError(err)
		pt.Y += c.PointToFixed(fontSize * spacing)
	}

	// 保存图片
	f, err := os.Create(*outPath)
	checkError(err)
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, rgba)
	checkError(err)
	err = b.Flush()
	checkError(err)
	fmt.Printf("successful!	OutPath:%s\n", *outPath)
}
