package main

import (
	"bufio"
	"fmt"
	"github.com/nfnt/resize"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

func main() {
	file, err := os.Open("D:\\df\\ce.png") //此处自行更改自己所需要识别的图片路径
	defer file.Close()
	if err != nil {
		panic(err)
	}
	//解析图片
	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	//得到新的图片
	newImage := Number(img)
	witerFile(newImage, "10")
	//将新图片二虚化
	new1Image := Binarization(newImage)
	witerFile(new1Image, "11")
	//切割图片
	new2Image := CutImage(new1Image)
	witerFile(new2Image, "12")
	//得到身份证号码图片集
	images := SplitImage(new2Image)
	idcard := NumberDistinguish(images)

	fmt.Printf("身份证号是:%s", idcard)
}

func witerFile(src image.Image, name string) {
	outFile, err := os.Create("D:\\df\\" + name + ".png")
	defer outFile.Close()
	if err != nil {
		panic(err)
	}
	b := bufio.NewWriter(outFile)
	err = png.Encode(b, src)
	if err != nil {
		panic(err)
	}
	err = b.Flush()
	if err != nil {
		panic(err)
	}
}

// Number 号码定位
func Number(src image.Image) image.Image {
	rect := src.Bounds() // 获取图片的大小
	// 左上角的坐标 x: w(总宽度)*测量x轴的位置/测量总x的宽度
	// 左上角的坐标 y: h(总高度)*测量y轴的位置/测量总y的高度
	//此处图片的尺寸需要根据所需识别的图片进行确定
	log.Println(rect.Dx())
	log.Println(rect.Dy())
	left := image.Point{X: rect.Dx() * 220 / 620, Y: rect.Dy() * 325 / 385}
	// 右下角的坐标 x: w(总宽度)*测量x轴的位置/测量总x的宽度
	// 右下角的坐标 y: h(总高度)*测量y轴的位置/测量总y的高度
	//此处图片的尺寸需要根据所需识别的图片进行确定
	right := image.Point{X: rect.Dx() * 540 / 620, Y: rect.Dy() * 345 / 385}
	newReact := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: right.X - left.X, Y: right.Y - left.Y},
	} // 创建一个新的矩形 ,将原图切割后的图片保存在该矩形中
	newImage := image.NewRGBA(newReact)                 // 创建一个新的图片
	draw.Draw(newImage, newReact, src, left, draw.Over) // 将原图绘制到新图片中
	return newImage
}

// Binarization 将图片二值化
func Binarization(src image.Image) image.Image {
	//将图片灰化
	dst := image.NewGray16(src.Bounds())                          // 创建一个新的灰度图
	draw.Draw(dst, dst.Bounds(), src, src.Bounds().Min, draw.Src) // 将原图绘制到新图片中

	//遍历像素点，实现二值化
	for x := 0; x < src.Bounds().Dx(); x++ {
		for y := 0; y < src.Bounds().Dy(); y++ {
			r, _, _, _ := src.At(x, y).RGBA() //取出每个像素的r,g,b,a
			if r < 0x5555 {
				dst.Set(x, y, color.White) //将灰度值小于0x5555的像素置为0
			} else {
				dst.Set(x, y, color.Black)
			}
		}
	}
	return dst
}

// CutImage 寻找边缘坐标更加细致的切割图片
func CutImage(src image.Image) image.Image {
	var left, right image.Point //左上角右下角坐标
	//寻找左边边缘白点的x坐标
	for x := 0; x < src.Bounds().Dx(); x++ {
		for y := 0; y < src.Bounds().Dy(); y++ {
			r, _, _, _ := src.At(x, y).RGBA()
			if r == 0xFFFF {
				left.X = x
				x = src.Bounds().Dx() //使外层循环结束
				break
			}
		}
	}
	//寻找左边边缘白点的y坐标
	for y := 0; y < src.Bounds().Dy(); y++ {
		for x := 0; x < src.Bounds().Dx(); x++ {
			r, _, _, _ := src.At(x, y).RGBA()
			if r == 0xFFFF {
				left.Y = y
				y = src.Bounds().Dy() //使外层循环结束
				break
			}
		}
	}
	//寻找右边边缘白点的x坐标
	for x := src.Bounds().Dx(); x > 0; x-- {
		for y := src.Bounds().Dy(); y > 0; y-- {
			r, _, _, _ := src.At(x, y).RGBA()
			if r == 0xFFFF {
				right.X = x + 1
				x = 0 //使外层循环结束
				break
			}
		}
	}
	//寻找右边边缘白点的y坐标
	for y := src.Bounds().Dy() - 1; y > 0; y-- {
		for x := src.Bounds().Dx() - 1; x > 0; x-- {
			r, _, _, _ := src.At(x, y).RGBA()
			if r == 0xFFFF {
				right.Y = y + 1
				y = 0 //使外层循环结束
				break
			}
		}
	}
	//按照坐标点将图像精准切割
	newReact := image.Rect(0, 0, right.X-left.X+1,
		right.Y-left.Y+2) // 创建一个新的矩形 ,将原图切割后的图片保存在该矩形中
	//log.Println(left, right)
	//log.Println(src.Bounds(), newReact)
	dst := image.NewRGBA(newReact)
	draw.Draw(dst, dst.Bounds(), src, left, draw.Over)
	return dst
}

// SplitImage 将每一个数字切割出来
func SplitImage(src image.Image) []image.Image {
	var dsts []image.Image
	leftX := 0
	for x := 0; x < src.Bounds().Dx(); x++ {
		temp := false
		for y := 0; y < src.Bounds().Dy(); y++ {
			r, _, _, _ := src.At(x, y).RGBA()
			if r == 0xFFFF {
				temp = true
				break
			}
		}
		if temp {
			continue
		}
		dst := image.NewGray16(image.Rect(0, 0, x-leftX, src.Bounds().Dy()))
		draw.Draw(dst, dst.Bounds(), src, image.Point{X: leftX, Y: 0}, draw.Src)
		//下一个起点
		for x1 := x + 1; x1 < src.Bounds().Dx(); x1++ {
			temp := false
			for y := 0; y < src.Bounds().Dy(); y++ {
				r, _, _, _ := src.At(x1, y).RGBA()
				if r == 0xFFFF {
					temp = true
					break
				}
			}
			if temp {
				leftX = x1
				x = x1
				break
			}
		}
		img := resize.Resize(8, 8, dst, resize.Lanczos3)
		dsts = append(dsts, img)
	}
	//fmt.Println(len(dsts))
	return dsts
}

var Data = map[string]string{
	"0": "0111110011111110000000001000000010000000100000100111111000011000",
	"1": "0100000001000000010000001100000011000000111111101111111011111110",
	"2": "0000001010000110000010001000100010011000000100001110000001100000",
	"3": "0000000000000000000000000001000000110000011100101101001011001110",
	"4": "0000110000011100001001000100010000000100000111100000110000000100",
	"5": "0000000011100000001000000000000000000010000100000001011000011100",
	"6": "0000110000111110001100100110000010000000100100100001111000001100",
	"7": "0000000000000000000011100001111000010000001000001100000011000000",
	"8": "0100111011111010100100100001000000010000101100100110111000000100",
	"9": "0010000001110000100110000000101000001110100011000111100001100000",
}

// NumberDistinguish 对每个数字图片进行遍历，得到的结果与库进行比较到处结果
func NumberDistinguish(src []image.Image) string {
	id := ""
	fmt.Printf("src的长度%d", src[0].Bounds().Dx())
	for i := 0; i < len(src); i++ {
		// 获取图片的指纹
		sign := ""
		for x := 0; x < src[i].Bounds().Dx(); x++ {
			for y := 0; y < src[i].Bounds().Dy(); y++ {
				r, _, _, _ := src[i].At(x, y).RGBA()
				if r > 0x7777 {
					sign += "1"
				} else {
					sign += "0"
				}
			}
		}

		// 对比指纹
		number := ""
		//对比相似率
		percent := 0.0
		for k, v := range Data {
			sum := 0
			for i := 0; i < 64; i++ {
				if v[i:i+1] == sign[i:i+1] {
					sum++
				}
			}
			//不断比较当匹配率达到最大时，就是此时所对应的数字
			if float64(sum)/64 > percent {
				number = k
				percent = float64(sum) / 64
			}
		}

		//log.Println(sign, number, percent)
		id += number
	}
	return id
}
