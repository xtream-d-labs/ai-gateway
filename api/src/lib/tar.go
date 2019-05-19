package lib

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
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

func archive(dirName, outName string) error {
	tarfile, err := os.Create(outName)
	if err != nil {
		return err
	}
	defer tarfile.Close()

	var writer io.WriteCloser = tarfile
	if strings.HasSuffix(outName, ".gz") {
		writer = gzip.NewWriter(tarfile)
		defer writer.Close()
	}
	tarfileWriter := tar.NewWriter(writer)
	defer tarfileWriter.Close()

	err = filepath.Walk(dirName,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			header := new(tar.Header)
			header.Name = strings.Replace(file.Name(), dirName, "", -1)
			header.Size = info.Size()
			header.Mode = int64(info.Mode())
			header.ModTime = info.ModTime()

			if err = tarfileWriter.WriteHeader(header); err != nil {
				return err
			}
			if _, err = io.Copy(tarfileWriter, file); err != nil {
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}
