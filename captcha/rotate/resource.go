package rotate

import (
	"image"
)

// Resources .
type Resources struct {
	rangImages []image.Image
}

// NewResources .
func NewResources() *Resources {
	return &Resources{}
}

type Resource func(*Resources)

// WithImages is to set image
func WithImages(images []image.Image) Resource {
	return func(resources *Resources) {
		resources.rangImages = images
	}
}
