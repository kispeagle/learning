package main

import (
	"design-pattern/structural/proxy"
	"fmt"
)

func main() {
	image := "http://github.com/kispeagle/avatar.jpg"

	proxyImage := proxy.ProxyImage{}

	fmt.Println("This is first call")
	proxyImage.ShowImage(image)

	fmt.Println("This is second call")
	proxyImage.ShowImage(image)

}
