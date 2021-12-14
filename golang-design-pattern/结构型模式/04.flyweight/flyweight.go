package flyweight

import "fmt"

type ImageFlyWeight struct {
	data string
}

func NewFlyWeight(name string) *ImageFlyWeight {
	data := fmt.Sprintf("image data %s",name)
	return &ImageFlyWeight{
		data: data,
	}
}

func (f *ImageFlyWeight) Data() string {
	return f.data
}

type ImageFlyWeightFactory struct {
	maps map[string] *ImageFlyWeight
}

var mageFlyWeightFactory *ImageFlyWeightFactory

func NewImageFlyWeightFactory() *ImageFlyWeightFactory {
	if mageFlyWeightFactory == nil {
		mageFlyWeightFactory = &ImageFlyWeightFactory{
			maps: map[string]*ImageFlyWeight{},
		}
	}
	return mageFlyWeightFactory
}

func (f ImageFlyWeightFactory) Get(name string) *ImageFlyWeight  {
	image := f.maps[name]
	if image == nil {
		image = NewFlyWeight(name)
		f.maps[name] = image
	}
	return image
}

type ImageWeightViewer struct {
	*ImageFlyWeight
}

func (i *ImageWeightViewer) Display()  {
	fmt.Printf("Display:%s\n",i.Data())
}

func NewImageWeightViewer(f string) *ImageWeightViewer {
	image := NewImageFlyWeightFactory().Get(f)
	return &ImageWeightViewer{
		ImageFlyWeight:image,
	}
}

