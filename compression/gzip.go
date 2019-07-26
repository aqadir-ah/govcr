package compression

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
)

// Compress data and return the result
func Compress(data []byte) ([]byte, error) {
	var out bytes.Buffer

	w := gzip.NewWriter(&out)
	if _, err := io.Copy(w, bytes.NewBuffer(data)); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// Decompress data and return the result
func Decompress(data []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	return data, nil
}