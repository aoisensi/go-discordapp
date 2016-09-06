package discord

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/jpeg"
	"io"
)

type AvatarData struct {
	buf *bytes.Buffer
}

func NewAvatarData(img image.Image) *AvatarData {
	d := new(AvatarData)
	d.buf = new(bytes.Buffer)
	jpeg.Encode(d.buf, img, nil)
	return d
}

func NewAvatarDataFromJPEG(jpg io.Reader) *AvatarData {
	d := new(AvatarData)
	d.buf = new(bytes.Buffer)
	d.buf.ReadFrom(jpg)
	return d
}

func (d *AvatarData) toString() string {
	data := bytes.NewBufferString("data:image/jpeg;base64,")
	data.ReadFrom(base64.NewDecoder(base64.StdEncoding, d.buf))
	return data.String()
}
