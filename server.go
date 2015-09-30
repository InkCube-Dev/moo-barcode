package main

import (
	"net/http"

	"github.com/golang/glog"

	"github.com/MooCommerce/moo-barcode/barcode"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {
	defer glog.Flush()

	goji.Get("/barcode/:code", barcodeHandler)
	goji.Serve()
}

func barcodeHandler(c web.C, rw http.ResponseWriter, req *http.Request) {
	b := &barcode.Barcode{c.URLParams["code"], 300, 50}
	b.WriteImage(rw)
}
