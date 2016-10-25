package qrcode_cover_test

import (
	"bytes"
	qrcodecover "github.com/cnjack/qrcode-cover"
	"image/jpeg"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	f, err := os.Open("logo.jpg")
	if err != nil {
		log.Println("1err:", err.Error())
		return
	}
	c, err := jpeg.Decode(f)
	if err != nil {
		log.Println("2err:", err.Error())
		return
	}
	q, err := qrcodecover.New("http://m.qschou.com", 3, 500, c)
	if err != nil {
		log.Println("3err:", err.Error())
		return
	}
	if q == nil {
		if err != nil {
			log.Println("4err: empty image")
			return
		}
	}
	var b bytes.Buffer
	err = jpeg.Encode(&b, q, &jpeg.Options{100})
	if err != nil {
		log.Println("5err:", err.Error())
		return
	}
	err = ioutil.WriteFile("test.jpg", b.Bytes(), os.FileMode(0644))
	if err != nil {
		log.Println("6err:", err.Error())
		return
	}
}
