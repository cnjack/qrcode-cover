package qrcode_cover

import (
	"github.com/nfnt/resize"
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
)

func New(content string, level qrcode.RecoveryLevel, size int, cover image.Image) (image.Image, error) {
	q, err := qrcode.New(content, level)
	if err != nil {
		return nil, err
	}
	q.ForegroundColor = color.RGBA{67, 172, 67, 1}
	img := q.Image(size)
	cover_new := resize.Resize(uint(size/100*20), uint(size/100*20), cover, resize.NearestNeighbor)
	offset := image.Pt(img.Bounds().Dx()/2-cover_new.Bounds().Dx()/2, img.Bounds().Dy()/2-cover_new.Bounds().Dy()/2)
	canvas := image.NewRGBA(image.Rect(0, 0, size, size))
	draw.Draw(canvas, canvas.Bounds(), img, image.ZP, draw.Src)
	draw.Draw(canvas, cover_new.Bounds().Add(offset), cover_new, image.ZP, draw.Over)
	return canvas, nil
}
