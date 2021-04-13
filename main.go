package main

import (
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
	size   int
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
	str = 0
	for i := 5; i > 1; i-- {
		str += int(im.txt[i]) * int(math.Pow(256, float64(i)-2))
	}
	im.size = str
}

func (im *image) multiplyImage(n int) image {
	var newIm image
	newIm.txt = append(newIm.txt, im.txt[0])
	newIm.txt = append(newIm.txt, im.txt[1])
	//s:= formatIntTo4Byte(im.size+(im.size*im.width)*(n-1))
	s := formatIntTo4Byte(im.width*im.height*n*n*3 + 55)
	for i := 2; i < 6; i++ {
		newIm.txt = append(newIm.txt, s[i-2])
	}
	for i := 6; i < 18; i++ {
		newIm.txt = append(newIm.txt, im.txt[i])
	}
	//copy(newIm.txt[:18], im.txt[:18])
	w := formatIntTo4Byte(im.width * n)
	for i := 18; i < 22; i++ {
		newIm.txt = append(newIm.txt, w[i-18])
	}
	h := formatIntTo4Byte(im.height * n)
	for i := 22; i < 26; i++ {
		//copy(newIm.txt[18:22], formatIntTo4Byte(newWidth))
		newIm.txt = append(newIm.txt, h[i-22])
	}
	//copy(newIm.txt[22:26], formatIntTo4Byte(newHeight))
	for i := 26; i < 29; i++ {
		newIm.txt = append(newIm.txt, im.txt[i])
	}
	for i := 29; i < 54; i++ {
		newIm.txt = append(newIm.txt, 0)
	}
	//fmt.Println(im.txt[54:])
	//copy(newIm.txt[26:54], im.txt[26:54])
	for i := 54; i < len(im.txt)-2; i += im.width * 3 {
		//fmt.Println(">",i)
		for p := 0; p < n; p++ {
			for j := i; j < i+im.width*3-2 && j < len(im.txt)-2; j += 3 {
				///fmt.Println("j:",j,"i:",i,len(im.txt),i+im.width*3-2)
				//fmt.Println(im.txt)
				for k := 0; k < n; k++ {
					newIm.txt = append(newIm.txt, im.txt[j])
					newIm.txt = append(newIm.txt, im.txt[j+1])
					newIm.txt = append(newIm.txt, im.txt[j+2])
				}
			}
		}
	}
	newIm.txt = append(newIm.txt, 0)
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
	im.open("1.bmp")
	im.sizeCalc()
	im2 := im.multiplyImage(10)
	im2.sizeCalc()
	im2.create("2.bmp")
	/*var im3 image
	im3.open("2.bmp")
	im3.sizeCalc()
	/*for i := 0; i < len(im3.txt); i++ {
		if i >= len(im.txt) {
			fmt.Println(i, im3.txt[i], ">")
		} else {
			fmt.Println(i, im3.txt[i], ">", im.txt[i])
		}
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
