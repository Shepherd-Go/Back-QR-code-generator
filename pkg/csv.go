package pkg

import (
	"encoding/csv"
	"mime/multipart"

	"github.com/jszwec/csvutil"
)

func BindFile(file *multipart.FileHeader, dest interface{}) (err error) {
	data, err := file.Open()
	if err != nil {
		return
	}

	defer data.Close()

	reader := csv.NewReader(data)
	header, err := csvutil.Header(dest, "csv")
	if err != nil {
		return
	}
	dec, err := csvutil.NewDecoder(reader, header...)
	if err != nil {
		return
	}

	if err = dec.Decode(dest); err != nil {
		return
	}

	return
}
