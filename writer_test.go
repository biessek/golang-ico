package ico

import (
	"testing"
	"os"
	"path/filepath"
	"image"
	"image/png"
)

func TestEncode(t *testing.T) {
	t.Parallel()
	origfile := "testdata/golang.ico"
	file := "testdata/golang_test.ico"

	f, err := os.Open("testdata/golang.png")
	img, err := png.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	var newFile *os.File
	if newFile, err = os.Create(filepath.Join(file)); err != nil {
		t.Error(err)
	}
	err = Encode(newFile, img)
	if err != nil {
		t.Error(err)
	}
	newFile.Close()

	f, err = os.Open(origfile)
	if err != nil {
		t.Error(err)
	}
	origICO, err := Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()

	newFile, err = os.Open(file)
	if err != nil {
		t.Error(err)
	}
	newICO, err := Decode(newFile)
	if err != nil {
		t.Error(err)
	}
	newFile.Close()

	inrgba, ok := origICO.(*image.NRGBA)
	if !ok {
		t.Fatal("not nrgba")
	}
	pnrgba, ok := newICO.(*image.NRGBA)
	if !ok {
		t.Fatal("new not nrgba")
	}
	if b, err := fastCompare(inrgba, pnrgba); err != nil || b != 0 {
		t.Fatalf("pix differ %d %v\n", b, err)
	}

}
