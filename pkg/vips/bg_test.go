package vips_test

import (
	"fmt"
	"testing"

	"github.com/davidbyttow/govips/pkg/vips"
)

func TestBackground(t *testing.T) {
	for i := 0; i < 200; i++ {
		b := vips.Color{255, 255, 255}
		_, _, err := vips.NewTransform().LoadFile("/Users/realityone/Codes/goprojects/src/github.com/davidbyttow/govips/assets/fixtures/aaa.png").BackgroundColor(b).Format(vips.ImageTypeJPEG).Lossless().Apply()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i)
		// time.Sleep(time.Millisecond * 200)
	}
}
