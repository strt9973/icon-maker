package main

import (
	"image"
	"image/color"
	"math"

	"golang.org/x/image/draw"
)
func resizeImage(img image.Image, size int) image.Image{
    newImage := image.NewRGBA(image.Rect(0, 0, size, size))
    draw.CatmullRom.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

    return newImage
}

func addMarginImage(img image.Image, size int) image.Image{
    newImage := image.NewRGBA(image.Rect(0, 0, size, size))

	contentSize := int(float64(size) * 0.8)
    margin := (size - contentSize) / 2
    contentImage := image.NewRGBA(image.Rect(0, 0, contentSize, contentSize))
    draw.CatmullRom.Scale(contentImage, contentImage.Bounds(), img, img.Bounds(), draw.Over, nil)
	offsetRect := image.Rect(margin, margin, margin+contentSize, margin+contentSize)
    draw.Draw(newImage, offsetRect, contentImage, image.Point{}, draw.Over)

    return newImage
}

func alphaForDistance(d, radius, margin float64) uint8 {
	if d <= radius-margin {
		return 255
	}
	if d >= radius {
		return 0
	}
	return uint8(255 * (radius - d) / margin)
}

func createRoundedMask(size int) *image.Alpha {
	radius := int(size) / 5
    margin := 1.0 
	mask := image.NewAlpha(image.Rect(0, 0, size, size))

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			alpha := 255 
			// 上左
			if x < radius && y < radius {
				d := math.Hypot(float64(radius-x-1), float64(radius-y-1))
				alpha = int(alphaForDistance(d, float64(radius), margin))
			}
			// 上右
			if x >= size-radius && y < radius {
				d := math.Hypot(float64(x-(size-radius)), float64(radius-y-1))
				alpha = int(alphaForDistance(d, float64(radius), margin))
			}
			// 下左
			if x < radius && y >= size-radius {
				d := math.Hypot(float64(radius-x-1), float64(y-(size-radius)))
				alpha = int(alphaForDistance(d, float64(radius), margin))
			}
			// 下右
			if x >= size-radius && y >= size-radius {
				d := math.Hypot(float64(x-(size-radius)), float64(y-(size-radius)))
				alpha = int(alphaForDistance(d, float64(radius), margin))
			}
			mask.SetAlpha(x, y, color.Alpha{A: uint8(alpha)})
		}
	}
	return mask
}

func applyMask(src image.Image, mask *image.Alpha) image.Image {
    bounds := src.Bounds()
    dst := image.NewNRGBA(bounds)
    draw.Draw(dst, bounds, src, bounds.Min, draw.Src)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            c := dst.NRGBAAt(x, y)
            maskAlpha := mask.AlphaAt(x, y).A
            c.A = uint8((uint16(c.A) * uint16(maskAlpha)) / 255)
            dst.SetNRGBA(x, y, c)
        }
    }
    return dst
}

func editImage(img image.Image, size int) image.Image {
	img = resizeImage(img, size)
	mask := createRoundedMask(size)
	img = applyMask(img, mask)
	img = addMarginImage(img, size)
	return img
}