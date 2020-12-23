package steg

import "image"

type imageType int

const (
	jpeg imageType = iota
	png
)

func (it imageType) Extension() string {
	return [...]string{"jpg", "png"}[it]
}

func processImage(filePath string) (image.Image, error) {
	return nil, nil
}

func exportImage(img image.Image, filePath string) error {
	return nil
}

func identifyFileType(filePath string) (imageType, error) {
	return jpeg, nil
}