package main

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
)

var (
	ErrUnsupportedFile = errors.New("the required argument was not found")
)

func main() {
	args := os.Args

	if len(args) < 2 {
		slog.Error(ErrUnsupportedFile.Error())
		os.Exit(1)
	}

	environment, err := ReadDir(os.Args[1])
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	for key, val := range environment {
		if !val.NeedRemove {
			_, ok := os.LookupEnv(key)
			if ok {
				err := os.Unsetenv(key)
				if err != nil {
					slog.Error(fmt.Errorf("unset env err: %w", err).Error())
				}
			}

			err := os.Setenv(key, val.Value)
			if err != nil {
				slog.Error(fmt.Errorf("set env err: %w", err).Error())
			}
		} else {
			err := os.Unsetenv(key)
			if err != nil {
				slog.Error(fmt.Errorf("unset env err: %w", err).Error())
			}
		}
	}

	for _, val := range os.Environ() {
		fmt.Println(val)
	}
}
