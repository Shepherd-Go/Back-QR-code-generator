package entity

import "image"

type QrImage struct {
	ImgFile  image.Image
	PathName string
	Serial   string
	Zone     string
}
