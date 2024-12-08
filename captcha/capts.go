package captcha

import (
	"github.com/PandaManPMC/pic_captcha/captcha/click"
	"github.com/PandaManPMC/pic_captcha/captcha/rotate"
	"github.com/PandaManPMC/pic_captcha/captcha/slide"
)

// NewClickBuilder .
func NewClickBuilder(opts ...click.Option) click.Builder {
	return click.NewBuilder(opts...)
}

// NewSlideBuilder .
func NewSlideBuilder(opts ...slide.Option) slide.Builder {
	return slide.NewBuilder(opts...)
}

// NewRotateBuilder .
func NewRotateBuilder(opts ...rotate.Option) rotate.Builder {
	return rotate.NewBuilder(opts...)
}
