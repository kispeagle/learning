package proxy

import (
	"fmt"
)

type IImage interface {
	ShowImage()
}

type RealImage struct {
	Image string
}

func NewRealImage(image string) *RealImage {
	return &RealImage{image}
}

func (r RealImage) ShowImage() {
	fmt.Println(r.Image)
}

type ProxyImage struct {
	Image IImage
}

func (p *ProxyImage) ShowImage(url string) {
	if p.Image == nil {
		fmt.Print("Loading image: ")
		p.Image = NewRealImage(url)
	} else {
		fmt.Print("Image already existed: ")
	}
	p.Image.ShowImage()
}

// func main() {
// 	image := "http://github.com/kispeagle/avatar.jpg"

// 	proxyImage := proxy.ProxyImage{}

// 	fmt.Println("This is first call")
// 	proxyImage.ShowImage(image)

// 	fmt.Println("This is second call")
// 	proxyImage.ShowImage(image)

// }
