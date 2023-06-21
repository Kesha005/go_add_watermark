package main


import (
    "image"
    "image/draw"
    "image/jpeg"
    "os"
)

func main(){
	imgb, _ := os.Open("1.jpg")
    img, _ := jpeg.Decode(imgb)
    defer imgb.Close()

    wmb, _ := os.Open("2.jpg")
    watermark, _ := jpeg.Decode(wmb)
    defer wmb.Close()

    offset := image.Pt(200, 200)
    b := img.Bounds()
    m := image.NewRGBA(b)
    draw.Draw(m, b, img, image.ZP, draw.Src)
    draw.Draw(m, watermark.Bounds().Add(offset), watermark, image.ZP, draw.Over)

    imgw, _ := os.Create("watermarked.jpg")
    jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
    defer imgw.Close()
}