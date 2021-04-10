package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
)

type pixel struct {
	red   int8
	green int8
	blue  int8
}

type image struct {
	txt    []byte
	height int
	width  int
	img    [][]pixel
}

func (im *image) open(str string) {
	f, err := os.Open(str)
	if err != nil {
		log.Fatal("failed to open file")
		return
	}
	im.txt, _ = ioutil.ReadAll(f)
}

func (im *image) sizeCalc() {
	var str int
	for i := 21; i > 17; i-- {
		str += int(im.txt[i]) * int(math.Pow(256, float64(i)-18))
	}
	im.width = str
	str = 0
	for i := 25; i > 21; i-- {
		str += int(im.txt[i]) * int(math.Pow(256, float64(i)-22))
	}
	im.height = str
}

func (im *image) multiplyImage(n int) image {
	newWidth := im.width * n
	newHeight := im.height // * n
	var newIm image
	for i := 0; i < 18; i++ {
		newIm.txt = append(newIm.txt, im.txt[i])
	}
	//copy(newIm.txt[:18], im.txt[:18])
	w := formatIntTo4Byte(newWidth)
	for i := 18; i < 22; i++ {
		newIm.txt = append(newIm.txt, w[i-18])
	}
	h := formatIntTo4Byte(newHeight)
	for i := 22; i < 26; i++ {
		//copy(newIm.txt[18:22], formatIntTo4Byte(newWidth))
		newIm.txt = append(newIm.txt, h[i-22])
	}
	//copy(newIm.txt[22:26], formatIntTo4Byte(newHeight))
	for i := 26; i < 55; i++ {
		newIm.txt = append(newIm.txt, im.txt[i])
	}
	//copy(newIm.txt[26:54], im.txt[26:54])

	for i := 54; i < len(im.txt)-1; i += 3 {
		for j := 0; j < n; j++ {
			newIm.txt = append(newIm.txt, im.txt[i])
			newIm.txt = append(newIm.txt, im.txt[i+1])
			newIm.txt = append(newIm.txt, im.txt[i+2])
			//copy(newIm.txt[i:i+3], im.txt[i:i+3])
		}
	}
	return newIm
}

func (im *image) create(str string) {
	f, _ := os.Create(str)
	f.Write(im.txt)
}

func formatIntTo4Byte(n int) []byte {
	sl := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		temp := n / 256
		sl[i] = n - temp*256
		n = temp
	}
	return []byte{byte(sl[0]), byte(sl[1]), byte(sl[2]), byte(sl[3])}
}

func main() {
	var im image
	im.open("test.bmp")
	im.sizeCalc()
	im2 := im.multiplyImage(2)
	im2.create("res.bmp")
	var im3 image
	im3.open("res.bmp")
	im3.sizeCalc()
	for i := 0; i < 54; i++ {
		fmt.Println(i, ">", im.txt[i])
	}
	/*for i := 54; i < len(d)-1; i += 3 {
		f2.Write(d[i : i+3])
		f2.Write(d[i : i+3])
	}
	f2.Write([]byte(string(0)))
	f2.Close()
	f2, _ = os.Open("res.bmp")
	d2, _ := ioutil.ReadAll(f2)
	fmt.Println(d)
	fmt.Println(d2)*/
}
