package click

import "github.com/PandaManPMC/pic_captcha/captcha/base/imagedata"

// CaptchaData .
type CaptchaData interface {
	GetData() map[int]*Dot
	GetMasterImage() imagedata.JPEGImageData
	GetThumbImage() imagedata.PNGImageData
}

// CaptData .
type CaptData struct {
	dots        map[int]*Dot
	masterImage imagedata.JPEGImageData
	thumbImage  imagedata.PNGImageData
}

var _ CaptchaData = (*CaptData)(nil)

// GetData is to get dot
func (c CaptData) GetData() map[int]*Dot {
	return c.dots
}

// GetMasterImage is to get master image
func (c CaptData) GetMasterImage() imagedata.JPEGImageData {
	return c.masterImage
}

// GetThumbImage is to get thumbnail image
func (c CaptData) GetThumbImage() imagedata.PNGImageData {
	return c.thumbImage
}
