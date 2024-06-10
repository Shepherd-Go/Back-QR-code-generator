package entity

import "image"

type QrImage struct {
	ImgFile image.Image
	Serial  string
	Zone    string
}
