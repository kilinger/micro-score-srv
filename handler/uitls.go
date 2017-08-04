package handler

import (
	"bytes"
	"image/png"

	"strings"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"

	"code.xxxxxx.com/micro/scores-srv/store"
)

func encode(content string, level qr.ErrorCorrectionLevel, size int) ([]byte, error) {
	code, err := qr.Encode(content, level, qr.Unicode)
	if err != nil {
		return nil, err
	}

	code, err = barcode.Scale(code, size, size)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = png.Encode(&b, code)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func generate(name, context string, size int, store store.Store) (string, error) {
	png, err := encode(context, qr.M, size)
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(name, ".ping") {
		name = name + ".png"
	}

	url, err := store.Save(name, bytes.NewReader(png), "image/png")
	if err != nil {
		return "", err
	}

	return url, nil
}
