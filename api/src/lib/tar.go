package lib

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
)

type tarContent struct {
	header *tar.Header
	body   []byte
}

func targz(contents []*tarContent) ([]byte, error) {
	result := bytes.Buffer{}
	writer := tar.NewWriter(&result)
	for _, content := range contents {
		if err := writer.WriteHeader(content.header); err != nil {
			return nil, err
		}
		if _, err := writer.Write(content.body); err != nil {
			return nil, err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return gz(result.Bytes())
}

func gz(content []byte) ([]byte, error) {
	result := bytes.Buffer{}
	writer, e := gzip.NewWriterLevel(&result, gzip.BestCompression)
	if e != nil {
		return nil, e
	}
	if _, err := writer.Write(content); err != nil {
		return nil, err
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
