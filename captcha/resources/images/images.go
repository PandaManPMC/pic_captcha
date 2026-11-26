package images

import (
	"fmt"
	"image"
	"os"

	assets "github.com/PandaManPMC/pic_captcha/captcha/bindata/images"
	"github.com/PandaManPMC/pic_captcha/captcha/helper"
)

func GetImages(filePath string, imgCount int) ([]image.Image, error) {
	embedded := filePath == ""

	if embedded {
		filePath = "sourcedata/images/image-%d.jpg"
		imgCount = 5
	}

	var images []image.Image

	for i := 1; i <= imgCount; i++ {
		asset, err := readImage(filePath, i, embedded)
		if err != nil {
			return images, err
		}

		img, err := helper.DecodeByteToJpeg(asset)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, nil
}

func readImage(filePath string, index int, embedded bool) ([]byte, error) {
	path := fmt.Sprintf(filePath, index)

	if embedded {
		return assets.Asset(path)
	}

	return os.ReadFile(path)
}
