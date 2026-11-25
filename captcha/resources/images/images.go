package images

import (
	"fmt"
	"image"

	assets "github.com/PandaManPMC/pic_captcha/captcha/bindata/images"
	"github.com/PandaManPMC/pic_captcha/captcha/helper"
)

func GetImages(filePath string, imgCount int) ([]image.Image, error) {

	if "" == filePath {
		filePath = "sourcedata/images/image-%d.jpg"
		imgCount = 5
	}

	var images []image.Image
	var err error
	for i := 1; i <= imgCount; i++ {
		var img image.Image
		var asset []byte
		asset, err = assets.Asset(fmt.Sprintf(filePath, i))
		if err != nil {
			return images, err
		}

		img, err = helper.DecodeByteToJpeg(asset)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, nil
}
