package barcode

import (
	"bytes"
	"image/png"
	"net/http"
	"os"
	"strconv"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
)

type Barcode struct {
	Code   string
	Width  int
	Height int
}

func (b *Barcode) Generate(location string) {
	filename := location + "/" + b.Code + "-" + strconv.Itoa(b.Width) + "x" + strconv.Itoa(b.Height) + ".png"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		bc, _ := code128.Encode(b.Code)
		bc, _ = barcode.Scale(bc, b.Width, b.Height)
		f, _ := os.Create(filename)
		defer f.Close()
		png.Encode(f, bc)
	}
}

func (b *Barcode) WriteImage(rw http.ResponseWriter) {
	buffer := new(bytes.Buffer)

	bc, _ := code128.Encode(b.Code)
	bc, _ = barcode.Scale(bc, b.Width, b.Height)

	png.Encode(buffer, bc)

	rw.Header().Set("Content-Type", "image/png")
	rw.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	rw.Write(buffer.Bytes())
}
