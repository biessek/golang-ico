package ico

import (
	"testing"
	"os"
	"image/png"
	"image"
	"fmt"
	"math"
)

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}
func fastCompare(img1, img2 *image.NRGBA) (int64, error) {
	if img1.Bounds() != img2.Bounds() {
		return 0, fmt.Errorf("image bounds not equal: %+v, %+v", img1.Bounds(), img2.Bounds())
	}

	accumError := int64(0)

	for i := 0; i < len(img1.Pix); i++ {
		accumError += int64(sqDiffUInt8(img1.Pix[i], img2.Pix[i]))
	}

	return int64(math.Sqrt(float64(accumError))), nil
}

func aTestDecodeConfig(t *testing.T) {
	t.Parallel()
	file := "testdata/golang.ico"
	copyFile := "testdata/golang.png"
	reader, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	icoImage, err := DecodeConfig(reader)
	reader.Close()
	if err != nil {
		t.Fatal(err)
	}
	reader, err = os.Open(copyFile)
	if err != nil {
		t.Fatal(err)
	}
	pngImage, err := png.DecodeConfig(reader)
	reader.Close()
	if err != nil {
		t.Fatal(err)
	}

	if icoImage != pngImage {
		t.Errorf("%v - %v", icoImage, pngImage)
	}

}

func TestDecode(t *testing.T) {
	t.Parallel()
	file := "testdata/golang.ico"
	copyFile := "testdata/golang.png"
	reader, err := os.Open(file)
	if err != nil {
		t.Fatal(err)
	}
	icoImage, err := Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	reader.Close()

	reader, err = os.Open(copyFile)
	if err != nil {
		t.Fatal(err)
	}
	pngImage, err := png.Decode(reader)
	if err != nil {
		t.Fatal(err)
	}
	reader.Close()

	if icoImage == nil || !icoImage.Bounds().Eq(pngImage.Bounds()) {
		t.Fatal("bounds differ")
	}
	inrgba, ok := icoImage.(*image.NRGBA)
	if !ok {
		t.Fatal("not nrgba")
	}
	pnrgba, ok := pngImage.(*image.NRGBA)
	if !ok {
		t.Fatal("png not nrgba")
	}

	if b, err := fastCompare(inrgba, pnrgba); err != nil || b != 0 {
		t.Fatalf("pix differ %d %v\n", b, err)
	}
}
