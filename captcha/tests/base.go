package tests

import (
	"image"
	"io/ioutil"

	"github.com/PandaManPMC/pic_captcha/captcha/base/codec"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
)

func loadPng(p string) (image.Image, error) {
	imgBytes, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return codec.DecodeByteToPng(imgBytes)
}

func loadFont(p string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return freetype.ParseFont(fontBytes)
}
