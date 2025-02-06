package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	var err error

	handleErr := func(file *os.File) {
		if closeErr := file.Close(); closeErr != nil {
			if err != nil {
				err = closeErr
			} else {
				err = fmt.Errorf("%w; %v", err, closeErr)
			}
		}
	}

	f1, err := os.OpenFile(fromPath, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer handleErr(f1)

	fi, err := f1.Stat()
	if err != nil {
		return err
	}

	fSize := fi.Size()
	if fSize < 0 {
		return ErrUnsupportedFile
	}

	if fSize < offset {
		return ErrOffsetExceedsFileSize
	}

	f2, err := os.Create(toPath)
	if err != nil {
		return err
	}
	defer handleErr(f2)

	_, err = f1.Seek(offset, 0)
	if err != nil {
		return err
	}

	if limit == 0 {
		limit = fSize
	}

	_, err = io.CopyN(f2, f1, limit)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}

	return nil
}
